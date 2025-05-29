package usecase

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	reportCreateCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "report_create_total",
		Help: "Total number of reports created",
	})
	reportCreateDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "report_create_duration_seconds",
		Help:    "Duration of report creation",
		Buckets: prometheus.DefBuckets,
	})
)

func init() {
	prometheus.MustRegister(reportCreateCounter)
	prometheus.MustRegister(reportCreateDuration)
}
