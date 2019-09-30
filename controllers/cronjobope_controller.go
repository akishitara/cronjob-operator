/*

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
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/akishitara/cronjob-monitor/pkg/debugger"
	akishitarav1 "github.com/akishitara/cronjob-operator/api/v1"
)

// CronjobOpeReconciler reconciles a CronjobOpe object
type CronjobOpeReconciler struct {
	client.Client
	Log logr.Logger
}

// +kubebuilder:rbac:groups=akishitara.akishitara.cronjob-operator,resources=cronjobopes,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=akishitara.akishitara.cronjob-operator,resources=cronjobopes/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=batch,resources=cronjob,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=batch,resources=cronjob/status,verbs=get
// +kubebuilder:rbac:groups=batch,resources=job,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=batch,resources=job/status,verbs=get

// Reconcile aaa
func (r *CronjobOpeReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {

	var err error
	ctx := context.Background()
	log := r.Log.WithValues("cronjobope", req.NamespacedName)

	// your logic here
	debugger.YamlPrint(MakeCronjobSample())
	cronjob := MakeCronjobSample()

	err = r.Get(ctx, req.NamespacedName, &cronjob)
	if err != nil {
		log.Info("Does Exist", "cronjob", cronjob)
		return ctrl.Result{}, nil
	}

	err = r.Client.Create(ctx, &cronjob)
	if err != nil {
		log.Error(err, "Error des", "cronjob", cronjob)
	}

	return ctrl.Result{}, nil
}

// SetupWithManager aaa
func (r *CronjobOpeReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&akishitarav1.CronjobOpe{}).
		Complete(r)
}
