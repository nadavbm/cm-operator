package kuber

import (
	"context"

	"github.com/nadavbm/cm-operator/api/v1alpha1"
	"github.com/nadavbm/zlog"
	"go.uber.org/zap"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type Kuber struct {
	logger *zlog.Logger
	client kubernetes.Clientset
}

// New will create a new instance of kuber
func New(logger *zlog.Logger) (*Kuber, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}

	k8sClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return &Kuber{
		logger: logger,
		client: *k8sClient,
	}, nil
}

// ApplyConfigMap will apply the configmap in kubernetes namespace
func (k *Kuber) ApplyConfigMap(ns string, cmspec v1alpha1.OpConfigMapSpec) (*v1.ConfigMap, error) {
	cm := buildConfigMap(ns, cmspec)
	cmInterface := k.client.CoreV1().ConfigMaps(cm.GetNamespace())

	cm, err := cmInterface.Create(context.TODO(), cm, metav1.CreateOptions{})
	if err == nil {
		k.logger.Info("configMap created", zap.String("namespace:", ns), zap.Any("configMap:", cm))
		return cm, nil
	} else if err != nil && !errors.IsAlreadyExists(err) {
		return nil, err
	}

	cmExists, err := cmInterface.Get(context.TODO(), cm.Name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	cmExists.Data = cm.Data
	cm, err = cmInterface.Update(context.TODO(), cmExists, metav1.UpdateOptions{})
	if err != nil {
		return nil, err
	}

	k.logger.Info("configMap updated", zap.String("namespace:", ns), zap.Any("configMap:", cm))
	return cm, nil
}

// buildConfigMap will build a kubernetes config map from the specifications given for file data and name
func buildConfigMap(ns string, cmspec v1alpha1.OpConfigMapSpec) *v1.ConfigMap {
	cm := &v1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "batch/v1/beta1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "cmop",
			Namespace: ns,
		},
		Data: map[string]string{
			cmspec.FileName: cmspec.FileData,
		},
	}
	return cm
}
