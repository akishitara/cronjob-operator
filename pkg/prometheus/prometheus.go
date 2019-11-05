package prometheus

import (
	"net/http"
	"time"

	k8scli "github.com/akishitara/cronjob-operator/pkg/client"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

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
