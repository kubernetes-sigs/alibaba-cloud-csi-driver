package metric

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
	"google.golang.org/protobuf/proto"
	"k8s.io/klog/v2"
)

// MetaDesc carries the metadata of a prometheus.Desc so samples can be
// rendered without a Prometheus registry.
type MetaDesc struct {
	*prometheus.Desc
	FQName          string
	Help            string
	VariableLabels  []string
	ConstLabelPairs []*dto.LabelPair
}

func NewMetaDesc(namespace, subsystem, name, help string, variableLabels []string, constLabels prometheus.Labels) *MetaDesc {
	fqName := prometheus.BuildFQName(namespace, subsystem, name)
	pairs := make([]*dto.LabelPair, 0, len(constLabels))
	for k, v := range constLabels {
		pairs = append(pairs, &dto.LabelPair{Name: proto.String(k), Value: proto.String(v)})
	}
	slices.SortFunc(pairs, func(a, b *dto.LabelPair) int {
		return strings.Compare(a.GetName(), b.GetName())
	})
	return &MetaDesc{
		Desc:            prometheus.NewDesc(fqName, help, variableLabels, constLabels),
		FQName:          fqName,
		Help:            help,
		VariableLabels:  variableLabels,
		ConstLabelPairs: pairs,
	}
}

// Metric is a single sample: descriptor, value and variable label pairs.
type Metric struct {
	*MetaDesc
	Value              float64
	ValueType          prometheus.ValueType
	VariableLabelPairs []*dto.LabelPair
}

func NewMetric(desc *MetaDesc, value float64, valueType prometheus.ValueType, labelValues ...string) (*Metric, error) {
	if desc == nil {
		return nil, errors.New("desc cannot be nil")
	}
	if len(desc.VariableLabels) != len(labelValues) {
		return nil, fmt.Errorf("got %d label values for %d variable labels of %s", len(labelValues), len(desc.VariableLabels), desc.FQName)
	}
	pairs := make([]*dto.LabelPair, len(labelValues))
	for i, label := range desc.VariableLabels {
		pairs[i] = &dto.LabelPair{Name: proto.String(label), Value: proto.String(labelValues[i])}
	}
	return &Metric{
		MetaDesc:           desc,
		Value:              value,
		ValueType:          valueType,
		VariableLabelPairs: pairs,
	}, nil
}

func MustNewMetric(desc *MetaDesc, value float64, valueType prometheus.ValueType, labelValues ...string) *Metric {
	m, err := NewMetric(desc, value, valueType, labelValues...)
	if err != nil {
		panic(err)
	}
	return m
}

func (m *Metric) constMetric() prometheus.Metric {
	values := make([]string, len(m.VariableLabelPairs))
	for i, pair := range m.VariableLabelPairs {
		values[i] = pair.GetValue()
	}
	return prometheus.MustNewConstMetric(m.Desc, m.ValueType, m.Value, values...)
}

func sendMetrics(ch chan<- prometheus.Metric, metrics []*Metric) {
	for _, m := range metrics {
		ch <- m.constMetric()
	}
}

// extractMetricsFromMetricVec snapshots a CounterVec/GaugeVec as samples.
func extractMetricsFromMetricVec(fqName, help string, vec prometheus.Collector, valueType prometheus.ValueType) []*Metric {
	ch := make(chan prometheus.Metric)
	go func() {
		vec.Collect(ch)
		close(ch)
	}()

	var metrics []*Metric
	for m := range ch {
		var sample dto.Metric
		if err := m.Write(&sample); err != nil {
			klog.ErrorS(err, "Failed to write metric", "fqName", fqName)
			continue
		}
		value, err := metricValue(&sample, valueType)
		if err != nil {
			klog.ErrorS(err, "Failed to get metric value", "fqName", fqName)
			continue
		}
		metrics = append(metrics, &Metric{
			MetaDesc: &MetaDesc{
				Desc:   m.Desc(),
				FQName: fqName,
				Help:   help,
			},
			Value:              value,
			ValueType:          valueType,
			VariableLabelPairs: sample.Label,
		})
	}
	return metrics
}

func metricValue(m *dto.Metric, valueType prometheus.ValueType) (float64, error) {
	switch valueType {
	case prometheus.CounterValue:
		if m.Counter == nil || m.Counter.Value == nil {
			return 0, errors.New("nil counter value")
		}
		return m.Counter.GetValue(), nil
	case prometheus.GaugeValue:
		if m.Gauge == nil || m.Gauge.Value == nil {
			return 0, errors.New("nil gauge value")
		}
		return m.Gauge.GetValue(), nil
	case prometheus.UntypedValue:
		if m.Untyped == nil || m.Untyped.Value == nil {
			return 0, errors.New("nil untyped value")
		}
		return m.Untyped.GetValue(), nil
	default:
		return 0, fmt.Errorf("unsupported value type: %v", valueType)
	}
}
