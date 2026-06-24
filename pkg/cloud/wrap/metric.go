package wrap

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	APIRequestsTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "alibaba_cloud", Subsystem: "api", Name: "requests_total",
		Help: "Number of Alibaba Cloud API calls",
	}, []string{"product", "action", "code"})

	APIDurationSecondsTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "alibaba_cloud", Subsystem: "api", Name: "duration_seconds_total",
		Help: "Elapsed time for Alibaba Cloud API calls",
	}, []string{"product", "action", "code"})
)
