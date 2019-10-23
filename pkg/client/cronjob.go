package client

import (
	batch "k8s.io/api/batch/v1"
	batchv1 "k8s.io/api/batch/v1beta1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/rand"
)

// cronjob全取得
func cronjobList() (*batchv1.CronJobList, error) {
	cronjobClient := kubeClient().BatchV1beta1().CronJobs(apiv1.NamespaceDefault)
	return cronjobClient.List(metav1.ListOptions{})
}

// 名前からcronjobを取得
func cronjobGet(name string) *batchv1.CronJob {
	cronjobClient := kubeClient().BatchV1beta1().CronJobs(apiv1.NamespaceDefault)
	res, _ := cronjobClient.Get(name, metav1.GetOptions{})
	return res
}

// cronjob(公式)から内部typeに変換
func cronjobToActiveCronjob(cronjob batchv1.CronJob) ActiveCronjob {
	var res ActiveCronjob
	res = ActiveCronjob{
		Name:     cronjob.Name,
		UID:      string(cronjob.UID),
		Schedule: cronjob.Spec.Schedule,
		Image:    cronjob.Spec.JobTemplate.Spec.Template.Spec.Containers[0].Image,
		Command:  cronjob.Spec.JobTemplate.Spec.Template.Spec.Containers[0].Command,
		Jobs:     jobsToActivejobs(jobListFilterCronjobName(cronjob.Name)),
	}
	return res
}

// ActiveCronjobsAllGet 登録されてる全部のcronjobを内部Typeで取得
func ActiveCronjobsAllGet() []ActiveCronjob {
	var res []ActiveCronjob
	cronjobs, _ := cronjobList()
	for _, a := range cronjobs.Items {
		res = append(res, cronjobToActiveCronjob(a))
	}
	return res
}

func MakeCronjobSample() batchv1.CronJob { //*v1beta1.CronJob {
	res := batchv1.CronJob{
		TypeMeta: metav1.TypeMeta{
			Kind:       "CronJob",
			APIVersion: "batch/v1beta1",
		},
		ObjectMeta: metav1.ObjectMeta{Name: "aaa"},
		Spec: batchv1.CronJobSpec{
			Schedule: "*/1 * * * *",
			JobTemplate: batchv1.JobTemplateSpec{
				Spec: MakeJobSpec(),
			},
		},
	}
	return res
}

func TriggerCronJob(name string) error {
	cronjobClient := kubeClient().BatchV1beta1().CronJobs(apiv1.NamespaceDefault)
	cronJob, err := cronjobClient.Get(name, metav1.GetOptions{})

	if err != nil {
		return err
	}

	annotations := make(map[string]string)
	annotations["cronjob.kubernetes.io/instantiate"] = "manual"

	labels := make(map[string]string)
	for k, v := range cronJob.Spec.JobTemplate.Labels {
		labels[k] = v
	}

	//job name cannot exceed DNS1053LabelMaxLength (52 characters)
	var newJobName string
	if len(cronJob.Name) < 42 {
		newJobName = cronJob.Name + "-manual-" + rand.String(3)
	} else {
		newJobName = cronJob.Name[0:41] + "-manual-" + rand.String(3)
	}

	jobToCreate := &batch.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:        newJobName,
			Namespace:   apiv1.NamespaceDefault,
			Annotations: annotations,
			Labels:      labels,
		},
		Spec: cronJob.Spec.JobTemplate.Spec,
	}

	_, err = kubeClient().BatchV1().Jobs(apiv1.NamespaceDefault).Create(jobToCreate)

	if err != nil {
		return err
	}

	return nil
}
