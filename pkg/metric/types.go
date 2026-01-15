package metric

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	promdto "github.com/prometheus/client_model/go"
)

type Metric struct {
	Desc      *prometheus.Desc
	Labels    []*promdto.LabelPair
	Value     float64
	ValueType prometheus.ValueType
}

func MustNewMetricWithTypedFactorDesc(desc typedFactorDesc, value float64, labelValues ...string) *Metric {
	metric, err := NewMetricWithTypedFactorDesc(desc, value, labelValues...)
	if err != nil {
		panic(err)
	}
	return metric
}

func MustNewMetric(desc *prometheus.Desc, value float64, valueType prometheus.ValueType, labels []string, labelValues ...string) *Metric {
	metric, err := NewMetric(desc, value, valueType, labels, labelValues...)
	if err != nil {
		panic(err)
	}
	return metric
}

func NewMetricWithTypedFactorDesc(desc typedFactorDesc, value float64, labelValues ...string) (*Metric, error) {
	return NewMetric(desc.desc, value, desc.valueType, desc.labels, labelValues...)
}

func NewMetric(desc *prometheus.Desc, value float64, valueType prometheus.ValueType, labels []string, labelValues ...string) (*Metric, error) {
	if desc == nil {
		return nil, fmt.Errorf("desc cannot be nil")
	}
	if len(labels) != len(labelValues) {
		return nil, fmt.Errorf("labels and labelValues must have the same length")
	}

	pairs, err := makeLabelPairs(labels, labelValues...)
	if err != nil {
		return nil, err
	}

	return &Metric{
		Desc:      desc,
		Labels:    pairs,
		Value:     value,
		ValueType: valueType,
	}, nil
}

func makeLabelPairs(labels []string, labelValues ...string) (pairs []*promdto.LabelPair, err error) {
	if len(labels) != len(labelValues) {
		return nil, fmt.Errorf("labels and labelValues must have the same length")
	}
	for i, label := range labels {
		pairs = append(pairs, &promdto.LabelPair{
			Name:  &label,
			Value: &labelValues[i],
		})
	}
	return pairs, nil
}
