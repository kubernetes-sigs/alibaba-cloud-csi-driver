package metric

import (
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
)

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

var (
	execCountName   = "execution_count"
	execCountFQName = prometheus.BuildFQName(csiNamespace, grpcSubsystem, execCountName)
	execCountHelp   = "CSI grpc execution count."

	execTimeName   = "execution_time_total"
	execTimeFQName = prometheus.BuildFQName(csiNamespace, grpcSubsystem, execTimeName)
	execTimeHelp   = "CSI grpc execution time in total."
)

var CsiGrpcExecTimeCollector = csiGrpcExecTimeCollector{
	ExecCountMetric: prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: csiNamespace,
		Subsystem: grpcSubsystem,
		Name:      execCountName,
		Help:      execCountHelp,
	}, csiGrpcExecTimeLabels),
	ExecTimeTotalMetric: prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: csiNamespace,
		Subsystem: grpcSubsystem,
		Name:      execTimeName,
		Help:      execTimeHelp,
	}, csiGrpcExecTimeLabels),
}

func init() {
	registerCollector(CsiGrpcExecTimeCollectorName, GetCsiGrpcExecTimeCollector)
}

func GetCsiGrpcExecTimeCollector() (Collector, error) {
	return &CsiGrpcExecTimeCollector, nil
}

func (c *csiGrpcExecTimeCollector) Get() []*Metric {
	countMetrics := extractMetricsFromMetricVec(execCountFQName, execCountHelp, c.ExecCountMetric, prometheus.CounterValue)
	timeMetrics := extractMetricsFromMetricVec(execTimeFQName, execTimeHelp, c.ExecTimeTotalMetric, prometheus.CounterValue)
	return append(countMetrics, timeMetrics...)
}

func (c *csiGrpcExecTimeCollector) Update(ch chan<- prometheus.Metric) error {
	metrics := c.Get()
	for _, metric := range metrics {
		ch <- prometheus.MustNewConstMetric(metric.Desc, metric.ValueType, metric.Value, convertLabelsToString(metric.VariableLabelPairs)...)
	}
	return nil
}

// InitGRPC fills each possible method with OK code and 0 value
//
// Tell users which metrics are available. And make rate() work for the first invoke.
func (c *csiGrpcExecTimeCollector) InitGRPC(desc grpc.ServiceDesc, driver string) {
	for _, method := range desc.Methods {
		labels := prometheus.Labels{
			CsiGrpcExecTimeLabelMethod: method.MethodName,
			CsiGrpcExecTimeLabelType:   driver,
			CsiGrpcExecTimeLabelCode:   "OK",
		}
		c.ExecCountMetric.With(labels).Add(0)
		c.ExecTimeTotalMetric.With(labels).Add(0)
	}
}
