package controllers

import (
	akishitarav1 "github.com/akishitara/cronjob-operator/api/v1"
	batchv1 "k8s.io/api/batch/v1"
	batchv1beta1 "k8s.io/api/batch/v1beta1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MakeCronjobSample aaa
func MakeCronjobSample(option akishitarav1.CronOption) batchv1beta1.CronJob { //*v1beta1.CronJob {
	return batchv1beta1.CronJob{
		TypeMeta: metav1.TypeMeta{
			Kind:       "CronJob",
			APIVersion: "batch/v1beta1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      option.JobOption.Name,
			Namespace: corev1.NamespaceDefault,
		},
		Spec: batchv1beta1.CronJobSpec{
			Schedule: option.JobOption.Schedule,
			JobTemplate: batchv1beta1.JobTemplateSpec{
				Spec: MakeJobSpec(option),
			},
			ConcurrencyPolicy:          chooseConcurrencyPolicy(option.ConcurrencyPolicy),
			SuccessfulJobsHistoryLimit: option.SuccessfulJobsHistoryLimit,
			FailedJobsHistoryLimit:     option.FailedJobsHistoryLimit,
		},
	}
}

// MakeJobSpec aaa
func MakeJobSpec(option akishitarav1.CronOption) batchv1.JobSpec {
	return batchv1.JobSpec{
		Parallelism:  option.Parallelism,
		Completions:  option.Completions,
		BackoffLimit: option.BackoffLimit,
		Template: corev1.PodTemplateSpec{
			ObjectMeta: metav1.ObjectMeta{
				Name: option.JobOption.Name,
			},
			Spec: corev1.PodSpec{
				Containers: []corev1.Container{
					{
						Name:    option.JobOption.Name,
						Image:   option.Image,
						Command: option.JobOption.Cmd,
					},
				},
				RestartPolicy: chooseRestartPolicy(option.RestartPolicy),
			},
		},
	}
}

func chooseRestartPolicy(restartPolicy string) corev1.RestartPolicy {
	switch restartPolicy {
	case "Never":
		return corev1.RestartPolicyNever
	case "Always":
		return corev1.RestartPolicyAlways
	case "Onfailure":
		return corev1.RestartPolicyOnFailure
	default:
		return corev1.RestartPolicyNever
	}
}

func chooseConcurrencyPolicy(concurrencyPolicy string) batchv1beta1.ConcurrencyPolicy {
	switch concurrencyPolicy {
	case "Allow":
		return batchv1beta1.AllowConcurrent
	case "Forbid":
		return batchv1beta1.ForbidConcurrent
	case "Replace":
		return batchv1beta1.ReplaceConcurrent
	default:
		return batchv1beta1.AllowConcurrent
	}
}
