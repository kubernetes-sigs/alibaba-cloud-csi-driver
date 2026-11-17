package metric

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
	"k8s.io/apimachinery/pkg/util/sets"
)

const (
	CsiGrpcExecTimeCollectorName = "csi_grpc_exec_time"
	CsiGrpcExecTimeLabelMethod   = "method"
	CsiGrpcExecTimeLabelType     = "type"
	CsiGrpcExecTimeLabelCode     = "code"

	execCountName     = "execution_count"
	execCountHelp     = "CSI grpc execution count."
	execTimeTotalName = "execution_time_total"
	execTimeTotalHelp = "CSI grpc execution time in total."
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
		Name:      execCountName,
		Help:      execCountHelp,
	}, csiGrpcExecTimeLabels),
	ExecTimeTotalMetric: prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: csiNamespace,
		Subsystem: grpcSubsystem,
		Name:      execTimeTotalName,
		Help:      execTimeTotalHelp,
	}, csiGrpcExecTimeLabels),
}

func (c *csiGrpcExecTimeCollector) Get(ctx context.Context, pvcs sets.Set[string]) ([]*Metric, error) {
	metrics := extractMetricsFromMetricVec(prometheus.BuildFQName(csiNamespace, grpcSubsystem, execCountName), execCountHelp, c.ExecCountMetric, prometheus.CounterValue)
	metrics = append(metrics, extractMetricsFromMetricVec(prometheus.BuildFQName(csiNamespace, grpcSubsystem, execTimeTotalName), execTimeTotalHelp, c.ExecTimeTotalMetric, prometheus.CounterValue)...)
	return metrics, nil
}

func (c *csiGrpcExecTimeCollector) Collect(ch chan<- prometheus.Metric) {
	c.ExecCountMetric.Collect(ch)
	c.ExecTimeTotalMetric.Collect(ch)
}

func (c *csiGrpcExecTimeCollector) Describe(ch chan<- *prometheus.Desc) {
	c.ExecCountMetric.Describe(ch)
	c.ExecTimeTotalMetric.Describe(ch)
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
