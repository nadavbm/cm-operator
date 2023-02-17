/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"time"

	"github.com/nadavbm/cm-operator/api/v1alpha1"
	opconfigmapv1alpha1 "github.com/nadavbm/cm-operator/api/v1alpha1"
	"github.com/nadavbm/zlog"
	"go.uber.org/zap"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

// OpConfigMapReconciler reconciles a OpConfigMap object
type OpConfigMapReconciler struct {
	Logger *zlog.Logger
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=opconfigmap.example.com,resources=opconfigmaps,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=opconfigmap.example.com,resources=opconfigmaps/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=opconfigmap.example.com,resources=opconfigmaps/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the OpConfigMap object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.0/pkg/reconcile
func (r *OpConfigMapReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	r.Logger.Info("v1alpha1.PgDeployer changed. Start reconcile", zap.String("namespace", req.NamespacedName.Namespace))

	var opcm v1alpha1.OpConfigMap
	if err := r.Client.Get(context.Background(), req.NamespacedName, &opcm); err != nil {
		if errors.IsNotFound(err) {
			r.Logger.Info("cmop not found, probably deleted. skipping..", zap.String("namespace", req.Namespace))
			return ctrl.Result{Requeue: false, RequeueAfter: 0}, nil
		}
		r.Logger.Error("could not fetch resource, check if crd applied in the cluster..")
		return ctrl.Result{Requeue: true, RequeueAfter: time.Minute}, err
	}

	object := metav1.Object(buildConfigMap(req.Namespace, opcm.Spec))

	if err := controllerutil.SetControllerReference(&opcm, object, r.Scheme); err != nil {
		return ctrl.Result{}, err
	}

	r.Logger.Info("create object", zap.String("namespace", req.Namespace), zap.String("object", object.GetName()))
	if err := r.Create(ctx, object.(client.Object)); err != nil {
		if errors.IsAlreadyExists(err) {
			if err := r.Update(ctx, object.(client.Object)); err != nil {
				if errors.IsInvalid(err) {
					r.Logger.Error("invalid update", zap.String("object", object.GetName()))
				} else {
					r.Logger.Error("unable to update", zap.String("object", object.GetName()))
				}
			}
		} else {
			r.Logger.Error("could not create object", zap.String("object", object.GetName()), zap.Error(err))
			return ctrl.Result{}, err
		}
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *OpConfigMapReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&opconfigmapv1alpha1.OpConfigMap{}).
		Complete(r)
}

// buildConfigMap will build a kubernetes config map from the specifications given for file data and name
func buildConfigMap(ns string, cmspec v1alpha1.OpConfigMapSpec) *v1.ConfigMap {
	cm := &v1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "batch/v1/beta1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      cmspec.ConfigMapName,
			Namespace: ns,
		},
		Data: map[string]string{
			cmspec.FileName: cmspec.FileData,
		},
	}
	return cm
}
