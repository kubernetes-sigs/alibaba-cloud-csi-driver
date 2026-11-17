package metric

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewMetaDesc(t *testing.T) {
	desc := NewMetaDesc("node", "volume", "test_total", "help", []string{"pv"}, prometheus.Labels{"z": "1", "a": "2"})
	assert.Equal(t, "node_volume_test_total", desc.FQName)
	assert.Equal(t, "help", desc.Help)
	require.Len(t, desc.ConstLabelPairs, 2)
	assert.Equal(t, "a", desc.ConstLabelPairs[0].GetName())
	assert.Equal(t, "z", desc.ConstLabelPairs[1].GetName())
}

func TestNewMetric(t *testing.T) {
	desc := NewMetaDesc("node", "volume", "test", "help", []string{"pv"}, nil)

	t.Run("nilDesc", func(t *testing.T) {
		_, err := NewMetric(nil, 1, prometheus.GaugeValue)
		assert.Error(t, err)
	})

	t.Run("labelCountMismatch", func(t *testing.T) {
		_, err := NewMetric(desc, 1, prometheus.GaugeValue)
		assert.Error(t, err)
	})

	t.Run("valid", func(t *testing.T) {
		m, err := NewMetric(desc, 1, prometheus.GaugeValue, "pv-1")
		require.NoError(t, err)
		require.Len(t, m.VariableLabelPairs, 1)
		assert.Equal(t, "pv", m.VariableLabelPairs[0].GetName())
		assert.Equal(t, "pv-1", m.VariableLabelPairs[0].GetValue())
	})
}

func TestMetricConstMetric(t *testing.T) {
	desc := NewMetaDesc("node", "volume", "test", "help", []string{"pv"}, prometheus.Labels{"node": "n1"})
	m := MustNewMetric(desc, 3, prometheus.GaugeValue, "pv-1")

	var sample dto.Metric
	require.NoError(t, m.constMetric().Write(&sample))
	assert.Equal(t, 3.0, sample.Gauge.GetValue())
	require.Len(t, sample.Label, 2)
	assert.Equal(t, "node", sample.Label[0].GetName())
	assert.Equal(t, "pv", sample.Label[1].GetName())
	assert.Equal(t, "pv-1", sample.Label[1].GetValue())
}

func TestSendMetrics(t *testing.T) {
	desc := NewMetaDesc("node", "volume", "test", "help", nil, nil)
	ch := make(chan prometheus.Metric, 2)
	sendMetrics(ch, []*Metric{
		MustNewMetric(desc, 1, prometheus.GaugeValue),
		MustNewMetric(desc, 2, prometheus.CounterValue),
	})
	close(ch)
	var count int
	for range ch {
		count++
	}
	assert.Equal(t, 2, count)
}

func TestMetricValue(t *testing.T) {
	t.Run("nilValues", func(t *testing.T) {
		for _, vt := range []prometheus.ValueType{prometheus.CounterValue, prometheus.GaugeValue, prometheus.UntypedValue} {
			_, err := metricValue(&dto.Metric{}, vt)
			assert.Error(t, err)
		}
	})

	t.Run("unsupportedType", func(t *testing.T) {
		_, err := metricValue(&dto.Metric{}, prometheus.ValueType(99))
		assert.Error(t, err)
	})

	t.Run("values", func(t *testing.T) {
		v := 2.5
		m := &dto.Metric{
			Counter: &dto.Counter{Value: &v},
			Gauge:   &dto.Gauge{Value: &v},
			Untyped: &dto.Untyped{Value: &v},
		}
		for _, vt := range []prometheus.ValueType{prometheus.CounterValue, prometheus.GaugeValue, prometheus.UntypedValue} {
			got, err := metricValue(m, vt)
			require.NoError(t, err)
			assert.Equal(t, v, got)
		}
	})
}

func TestExtractMetricsFromMetricVec(t *testing.T) {
	vec := prometheus.NewCounterVec(prometheus.CounterOpts{Name: "test_total", Help: "help"}, []string{"pv"})
	vec.WithLabelValues("pv-1").Add(2)

	metrics := extractMetricsFromMetricVec("test_total", "help", vec, prometheus.CounterValue)
	require.Len(t, metrics, 1)
	assert.Equal(t, "test_total", metrics[0].FQName)
	assert.Equal(t, 2.0, metrics[0].Value)
	assert.Equal(t, prometheus.CounterValue, metrics[0].ValueType)
	require.Len(t, metrics[0].VariableLabelPairs, 1)
	assert.Equal(t, "pv", metrics[0].VariableLabelPairs[0].GetName())
	assert.Equal(t, "pv-1", metrics[0].VariableLabelPairs[0].GetValue())
}
