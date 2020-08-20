/*
Copyright 2020 Alexei Ledenev.

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

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	chaosv1alpha1 "github.com/alexei-led/timon/apis/chaos/v1alpha1"
)

// KillContainerReconciler reconciles a KillContainer object
type KillContainerReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=chaos.pumba.dev,resources=killcontainers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=chaos.pumba.dev,resources=killcontainers/status,verbs=get;update;patch

func (r *KillContainerReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("killcontainer", req.NamespacedName)

	// your logic here

	return ctrl.Result{}, nil
}

func (r *KillContainerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&chaosv1alpha1.KillContainer{}).
		Complete(r)
}
