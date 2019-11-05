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

	akishitarav1 "github.com/akishitara/cronjob-operator/api/v1"
	"github.com/akishitara/cronjob-operator/pkg/debugger"
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

	ctx := context.Background()
	log := r.Log.WithValues("cronjobope", req.NamespacedName)

	var err error
	var cronjob akishitarav1.CronjobOpeList

	// your logic here
	err = r.List(ctx, &cronjob)
	if err != nil {
		log.Info("OK")
		return ctrl.Result{}, nil
	}

	debugger.YamlPrint(cronjob)

	for _, item := range cronjob.Items {
		var image string
		var restartPolicy string
		var successfullJobHistoryLimit *int32
		var failedJobHistoryLimit *int32
		var concurrencyPolicy string
		var parallelism *int32
		var completions *int32
		var backoffLimit *int32

		image = item.Spec.Image
		restartPolicy = item.Spec.RestartPolicy
		successfullJobHistoryLimit = item.Spec.SuccessfulJobsHistoryLimit
		failedJobHistoryLimit = item.Spec.FailedJobsHistoryLimit
		concurrencyPolicy = item.Spec.ConcurrencyPolicy
		parallelism = item.Spec.Parallelism
		completions = item.Spec.Completions
		backoffLimit = item.Spec.BackoffLimit

		for _, param := range item.Spec.Param1 {
			option := akishitarav1.CronOption{
				JobOption: akishitarav1.JobOption{
					Name:     param.Name,
					Schedule: param.Schedule,
					Cmd:      param.Cmd,
				},
				Image:                      image,
				RestartPolicy:              restartPolicy,
				SuccessfulJobsHistoryLimit: successfullJobHistoryLimit,
				FailedJobsHistoryLimit:     failedJobHistoryLimit,
				ConcurrencyPolicy:          concurrencyPolicy,
				Parallelism:                parallelism,
				Completions:                completions,
				BackoffLimit:               backoffLimit,
			}

			sampleData := MakeCronjobSample(item, option)
			debugger.YamlPrint(sampleData)
			err = r.Client.Create(ctx, &sampleData)
			if err != nil {
				log.Info("Create Fail")
			}
		}
	}

	return ctrl.Result{}, nil
}

// SetupWithManager aaa
func (r *CronjobOpeReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&akishitarav1.CronjobOpe{}).
		Complete(r)
}
