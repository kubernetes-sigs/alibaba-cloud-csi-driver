package metric

import (
	"fmt"
	"sort"

	"github.com/prometheus/client_golang/prometheus"
	promdto "github.com/prometheus/client_model/go"
	"google.golang.org/protobuf/proto"
)

type Metric struct {
	*MetaDesc
	Value              float64
	ValueType          prometheus.ValueType
	VariableLabelPairs []*promdto.LabelPair
}

type MetaDesc struct {
	*prometheus.Desc
	FQName          string
	Help            string
	ConstLabelPairs []*promdto.LabelPair
	VariableLabels  []string
}

func NewMetaDesc(namespace, subsystem, name, help string, variableLabels []string, constLabels prometheus.Labels) (res *MetaDesc) {
	fqName := prometheus.BuildFQName(namespace, subsystem, name)
	desc := prometheus.NewDesc(fqName, help, variableLabels, constLabels)
	res = &MetaDesc{
		Desc:           desc,
		FQName:         fqName,
		Help:           help,
		VariableLabels: variableLabels,
	}
	res.ConstLabelPairs = make([]*promdto.LabelPair, 0, len(constLabels))
	for k, v := range constLabels {
		res.ConstLabelPairs = append(res.ConstLabelPairs, &promdto.LabelPair{
			Name:  proto.String(k),
			Value: proto.String(v),
		})
	}
	sort.Sort(LabelPairSorter(res.ConstLabelPairs))
	return
}

func MustNewMetricWithTypedFactorDesc(desc typedFactorDesc, value float64, labelValues ...string) *Metric {
	if desc.factor != 0 {
		value *= desc.factor
	}
	return MustNewMetricWithMetaDesc(desc.MetaDesc, value, desc.valueType, labelValues...)
}

func MustNewMetricWithMetaDesc(desc *MetaDesc, value float64, valueType prometheus.ValueType, labelValues ...string) *Metric {
	return MustNewMetric(desc, value, valueType, labelValues...)
}

func MustNewMetric(desc *MetaDesc, value float64, valueType prometheus.ValueType, labelValues ...string) *Metric {
	metric, err := NewMetric(desc, value, valueType, labelValues...)
	if err != nil {
		panic(err)
	}
	return metric
}

func NewMetricWithTypedFactorDesc(desc typedFactorDesc, value float64, labelValues ...string) (*Metric, error) {
	if desc.factor != 0 {
		value *= desc.factor
	}
	return NewMetric(desc.MetaDesc, value, desc.valueType, labelValues...)
}

func NewMetric(desc *MetaDesc, value float64, valueType prometheus.ValueType, labelValues ...string) (*Metric, error) {
	if desc == nil {
		return nil, fmt.Errorf("desc cannot be nil")
	}
	if len(desc.VariableLabels) != len(labelValues) {
		return nil, fmt.Errorf("labels and labelValues must have the same length")
	}

	pairs, err := makeLabelPairs(desc.VariableLabels, labelValues...)
	if err != nil {
		return nil, err
	}

	return &Metric{
		MetaDesc:           desc,
		VariableLabelPairs: append(pairs),
		Value:              value,
		ValueType:          valueType,
	}, nil
}

func makeLabelPairs(labels []string, labelValues ...string) (pairs []*promdto.LabelPair, err error) {
	gatherer := prometheus.DefaultGatherer
	gatherer.Gather()
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

type LabelPairSorter []*promdto.LabelPair

func (s LabelPairSorter) Len() int {
	return len(s)
}

func (s LabelPairSorter) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s LabelPairSorter) Less(i, j int) bool {
	return s[i].GetName() < s[j].GetName()
}
