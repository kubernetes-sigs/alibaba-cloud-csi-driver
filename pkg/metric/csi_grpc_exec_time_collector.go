package metric

import "github.com/prometheus/client_golang/prometheus"

const (
	CsiGrpcExecTimeCollectorName = "csi_grpc_exec_time"
	CsiGrpcExecTimeLabelMethod   = "method"
	CsiGrpcExecTimeLabelType     = "type"
	CsiGrpcExecTimeLabelCode     = "code"
)

var csiGrpcExecTimeLabels = []string{
	CsiGrpcExecTimeLabelMethod, CsiGrpcExecTimeLabelType, CsiGrpcExecTimeLabelCode,
}

type csiGrpcExecTimeCollector struct {
	ExecCountMetric     *prometheus.CounterVec
	ExecTimeTotalMetric *prometheus.CounterVec
}

var CsiGrpcExecTimeCollector = csiGrpcExecTimeCollector{
	ExecCountMetric: prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: csiNamespace,
		Subsystem: grpcSubsystem,
		Name:      "execution_count",
		Help:      "CSI grpc execution count.",
	}, csiGrpcExecTimeLabels),
	ExecTimeTotalMetric: prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: csiNamespace,
		Subsystem: grpcSubsystem,
		Name:      "execution_time_total",
		Help:      "CSI grpc execution time in total.",
	}, csiGrpcExecTimeLabels),
}

func init() {
	registerCollector(CsiGrpcExecTimeCollectorName, GetCsiGrpcExecTimeCollector)
}

func GetCsiGrpcExecTimeCollector() (Collector, error) {
	return &CsiGrpcExecTimeCollector, nil
}

func (c *csiGrpcExecTimeCollector) Update(ch chan<- prometheus.Metric) error {
	c.ExecCountMetric.Collect(ch)
	c.ExecTimeTotalMetric.Collect(ch)
	return nil
}
