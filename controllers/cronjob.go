package controllers

import (
	batchv1 "k8s.io/api/batch/v1"
	batchv1beta1 "k8s.io/api/batch/v1beta1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MakeCronjobSample aaa
func MakeCronjobSample() batchv1beta1.CronJob { //*v1beta1.CronJob {
	return batchv1beta1.CronJob{
		TypeMeta: metav1.TypeMeta{
			Kind:       "CronJob",
			APIVersion: "batch/v1beta1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "aaa",
			Namespace: corev1.NamespaceDefault,
		},
		Spec: batchv1beta1.CronJobSpec{
			Schedule: "*/1 * * * *",
			JobTemplate: batchv1beta1.JobTemplateSpec{
				Spec: MakeJobSpec(),
			},
		},
	}
}

// MakeJobSpec aaa
func MakeJobSpec() batchv1.JobSpec {
	return batchv1.JobSpec{
		Template: corev1.PodTemplateSpec{
			ObjectMeta: metav1.ObjectMeta{
				Name: "bbb",
			},
			Spec: corev1.PodSpec{
				Containers:    []corev1.Container{{Name: "eee", Image: "nginx"}},
				RestartPolicy: "Never",
			},
		},
	}
}
