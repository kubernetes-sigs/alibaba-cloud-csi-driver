package metric

import "github.com/prometheus/client_golang/prometheus"

const (
	CsiGrpcExecTimeCollectorName = "csi_grpc_exec_time"
	CsiGrpcExecTimeLabelMethod   = "method"
	CsiGrpcExecTimeLabelType     = "type"
	CsiGrpcExecTimeLabelId       = "id"
	CsiGrpcExecTimeLabelCode     = "code"
)

var csiGrpcExecTimeLabels = []string{
	CsiGrpcExecTimeLabelMethod, CsiGrpcExecTimeLabelType, CsiGrpcExecTimeLabelId, CsiGrpcExecTimeLabelCode,
}

type csiGrpcExecTimeCollector struct {
	Metric *prometheus.HistogramVec
}

var CsiGrpcExecTimeCollector = csiGrpcExecTimeCollector{
	Metric: prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: csiNamespace,
		Subsystem: grpcSubsystem,
		Name:      "execution_time",
		Help:      "Csi grpc execution time.",
	}, csiGrpcExecTimeLabels),
}

func init() {
	registerCollector(CsiGrpcExecTimeCollectorName, NewCsiGrpcExecTimeCollector, diskDriverName, nasDriverName, ossDriverName)
}

func NewCsiGrpcExecTimeCollector() (Collector, error) {
	return &CsiGrpcExecTimeCollector, nil
}

func (c *csiGrpcExecTimeCollector) Update(ch chan<- prometheus.Metric) error {
	c.Metric.Collect(ch)
	return nil
}
