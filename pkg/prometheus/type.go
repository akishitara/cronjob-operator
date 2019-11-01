package prometheus

import (
	"net/http"
	"time"

	k8scli "github.com/akishitara/cronjob-operator/pkg/client"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	JobStatusMetrics = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "successJobSatus",
		Help: "cronjob success job count",
	},
		[]string{"cronjob"},
	)

	JobDurationMetrics = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "successJob",
		Help: "cronjob success job",
	},
		[]string{"cronjob"},
	)

	CronjobErrorLogCountMetrics = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "ErrorLogCount",
		Help: "last pods error log",
	},
		[]string{"cronjob"},
	)
)

type PodStatus struct {
	Name       string
	LastStatus bool
	Success    int16
	Fail       int16
}

func (input PodStatus) increse(status bool) PodStatus {
	if status {
		return PodStatus{
			Name:       input.Name,
			LastStatus: input.LastStatus,
			Success:    input.Success + 1,
			Fail:       input.Fail,
		}
	} else {
		return PodStatus{
			Name:       input.Name,
			LastStatus: input.LastStatus,
			Success:    input.Success,
			Fail:       input.Fail + 1,
		}
	}
}

//func init() {
//	prometheus.MustRegister(JobStatusMetrics)
//}

func Run() error {
	go func() {
		var errorStrings = []string{
			"Error",
			"Fail",
			"false",
			"Exception",
		}
		for {
			cronjobs := k8scli.ActiveCronjobsAllGet()
			for _, cronjob := range cronjobs {
				errorLogCount := k8scli.LastJobErrorLogCount(cronjob, errorStrings)
				CronjobErrorLogCountMetrics.
					WithLabelValues(cronjob.Name).
					Set(float64(errorLogCount))
				time.Sleep(10 * time.Second)
			}
		}
	}()
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
	return errors.New("aaa")
}
