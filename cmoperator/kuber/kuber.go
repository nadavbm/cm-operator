package kuber

import (
	"github.com/nadavbm/cm-operator/api/v1alpha1"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type Kuber struct {
	client kubernetes.Clientset
}

func New() (*Kuber, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}

	k8sClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return &Kuber{
		client: *k8sClient,
	}, nil
}

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
