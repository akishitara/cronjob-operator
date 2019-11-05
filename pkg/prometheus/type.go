package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
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
