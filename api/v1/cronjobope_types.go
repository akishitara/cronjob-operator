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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// JobOption
type JobOption struct {
	Name     string   `json:"name"`
	Schedule string   `json:"schedule"`
	Cmd      []string `json:"cmd"`
}

// CronExec Option
type CronOption struct {
	JobOption                  JobOption `json:"jobOption"`
	Image                      string    `json:"image"`
	RestartPolicy              string    `json:"restartPolicy"`
	SuccessfulJobsHistoryLimit *int32    `json:"successfullJobHistoryLimit"`
	FailedJobsHistoryLimit     *int32    `json:"failedJobsHistoryLimit"`
	ConcurrencyPolicy          string    `json:"concurrencyPolicy"`
	Parallelism                *int32    `json:"parallelism"`
	Completions                *int32    `json:"completions"`
	BackoffLimit               *int32    `json:"backoffLimit"`
}

// CronjobOpeSpec defines the desired state of CronjobOpe
type CronjobOpeSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Schedule Cron format
	Param1                     []JobOption `json:"param1"`
	Image                      string      `json:"image"`
	RestartPolicy              string      `json:"restartPolicy"`
	SuccessfulJobsHistoryLimit *int32      `json:"successfullJobHistoryLimit"`
	FailedJobsHistoryLimit     *int32      `json:"failedJobsHistoryLimit"`
	ConcurrencyPolicy          string      `json:"concurrencyPolicy"`
	Parallelism                *int32      `json:"parallelism"`
	Completions                *int32      `json:"completions"`
	BackoffLimit               *int32      `json:"backoffLimit"`
}

// CronjobOpeStatus defines the observed state of CronjobOpe
type CronjobOpeStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// CronjobOpe is the Schema for the cronjobopes API
type CronjobOpe struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CronjobOpeSpec   `json:"spec,omitempty"`
	Status CronjobOpeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CronjobOpeList contains a list of CronjobOpe
type CronjobOpeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CronjobOpe `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CronjobOpe{}, &CronjobOpeList{})
}
