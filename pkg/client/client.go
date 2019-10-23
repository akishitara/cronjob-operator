package client

import (
	"fmt"
	"log"

	"github.com/akishitara/cronjob-operator/pkg/debugger"
	corev1 "k8s.io/api/core/v1"
)

//func init() {
//	prometheus.MustRegister(JobStatusMetrics)
//}

// Run Debug用
func Run() {
	fmt.Println("==========")

	debugger.YamlPrint("aaa")
	cronjob := MakeCronjobSample()
	batchclient := kubeClient().BatchV1beta1().CronJobs(corev1.NamespaceDefault)

	_, err := batchclient.Create(&cronjob)
	if err != nil {
		log.Fatal(err)
	}
	//	Cronjob := ActiveCronjobsAllGet()
	//	debugger.YamlPrint(Cronjob)
	//
	//	for i, _ := range Cronjob {
	//		debugger.YamlPrint(Cronjob[i].lastJobs())
	//		for j, _ := range Cronjob[i].Jobs {
	//			debugger.YamlPrint(Cronjob[i].Jobs[j].Status[0].Status)
	//		}
	//	}
	//
	//	go func() {
	//		for {
	//			JobStatusMetrics.With(prometheus.Labels{"cronjob": "crontest1"}).Set(1)
	//			time.Sleep(30 * time.Second)
	//		}
	//	}()
	//	http.Handle("/metrics", promhttp.Handler())
	//	log.Fatal(http.ListenAndServe(":8080", nil))
}
