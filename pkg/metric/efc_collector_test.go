package metric

import (
	"net"
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	promdto "github.com/prometheus/client_model/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestEfcCollector_extractUniqueMountUUIDs(t *testing.T) {
	collector := &efcCollector{}

	t.Run("should extract unique mount UUIDs from mountinfo", func(t *testing.T) {
		tempDir := t.TempDir()
		mountInfoPath := filepath.Join(tempDir, "mountinfo")

		mountInfoContent := `1851 3839 0:245 / /run/alinas/bindroot/bindroot-ce80f-WPcW4L23 rw,nosuid,nodev,relatime shared:479 - fuse.aliyun-alinas-efc bindroot-ce80f-WPcW4L23:35e9fxxxxx-xxxxx.cn-shanghai.nas.aliyuncs.com:/ rw,user_id=0,xxx
1852 3839 0:246 / /run/alinas/bindroot/bindroot-ce80f-WPcW4L24 rw,nosuid,nodev,relatime shared:480 - fuse.aliyun-alinas-efc bindroot-ce80f-WPcW4L24:35e9fxxxxx-xxxxx.cn-hangzhou.nas.aliyuncs.com:/ rw,user_id=0,xxx
1853 3839 0:247 / /run/alinas/bindroot/bindroot-ce80f-WPcW4L25 rw,nosuid,nodev,relatime shared:481 - fuse.aliyun-alinas-efc bindroot-ce80f-WPcW4L25:35e9fxxxxx-xxxxx.cn-beijing.nas.aliyuncs.com:/ rw,user_id=0,xxx
1854 3839 0:248 / /run/alinas/bindroot/bindroot-ce80f-WPcW4L23 rw,nosuid,nodev,relatime shared:482 - ext4 /dev/sda1 rw,relatime shared:1 -`

		err := os.WriteFile(mountInfoPath, []byte(mountInfoContent), 0644)
		require.NoError(t, err)

		uuids, err := collector.extractUniqueMountUUIDs(mountInfoPath)
		require.NoError(t, err)
		assert.Len(t, uuids, 3)

		expectedUUIDs := []string{"bindroot-ce80f-WPcW4L23", "bindroot-ce80f-WPcW4L24", "bindroot-ce80f-WPcW4L25"}
		for _, uuid := range expectedUUIDs {
			_, exists := uuids[uuid]
			assert.True(t, exists, "UUID %s should be present", uuid)
		}
	})

	t.Run("should handle file not found error", func(t *testing.T) {
		nonExistentPath := filepath.Join(t.TempDir(), "non-existent")
		uuids, err := collector.extractUniqueMountUUIDs(nonExistentPath)
		require.Error(t, err)
		assert.Nil(t, uuids)
	})
}

func TestEfcCollector_scanMonitorSockets(t *testing.T) {
	collector := &efcCollector{}

	t.Run("should find monitor socket files for given UUIDs", func(t *testing.T) {
		tempDir := os.TempDir()

		socketFile1 := filepath.Join(tempDir, "mountuuid1.monitor.sock")
		socketFile2 := filepath.Join(tempDir, "mountuuid2.monitor.sock")
		nonMonitorFile := filepath.Join(tempDir, "mountuuid3.other.sock")
		socketFiles := []string{socketFile1, socketFile2, nonMonitorFile}

		for _, sock := range socketFiles {
			listener, err := net.Listen("unix", sock)
			if err != nil {
				t.Fatalf("Failed to create socket file %v: %v", sock, err)
			}
			defer func() {
				_ = listener.Close()
			}()
		}

		mountUUIDs := map[string]struct{}{
			"mountuuid1": {},
			"mountuuid2": {},
			"mountuuid3": {}, // This one won't be found as a monitor socket
		}

		socketPaths, err := collector.scanMonitorSockets(tempDir, mountUUIDs)
		require.NoError(t, err)
		assert.Len(t, socketPaths, 2)

		expectedPaths := []string{socketFile1, socketFile2}
		for _, expectedPath := range expectedPaths {
			assert.Contains(t, socketPaths, expectedPath)
		}
	})

	t.Run("should handle directory not found error", func(t *testing.T) {
		nonExistentDir := filepath.Join(t.TempDir(), "non-existent")
		mountUUIDs := map[string]struct{}{
			"bindroot-ce80f-WPcW4L23": {},
		}

		socketPaths, err := collector.scanMonitorSockets(nonExistentDir, mountUUIDs)
		require.Error(t, err)
		assert.Nil(t, socketPaths)
	})
}

func TestEfcCollector_convertPrometheusDtoToMetric(t *testing.T) {
	tests := []struct {
		name            string
		metricFamily    *promdto.MetricFamily
		expectedMetrics []*Metric
	}{
		{
			name: "nil metrics",
		},
		{
			name: "counter type metric family",
			metricFamily: &promdto.MetricFamily{
				Name: proto.String("test_fq"),
				Help: proto.String("test-help"),
				Type: (*promdto.MetricType)(proto.Int32(int32(promdto.MetricType_COUNTER))),
				Metric: []*promdto.Metric{
					{
						Label: []*promdto.LabelPair{
							{
								Name:  proto.String("label1"),
								Value: proto.String("value1"),
							},
						},
						Counter: &promdto.Counter{
							Value: proto.Float64(123.45),
						},
					},
					{
						Counter: &promdto.Counter{
							Value: proto.Float64(67.89),
						},
					},
				},
			},
			expectedMetrics: []*Metric{
				{
					MetaDesc: &MetaDesc{
						FQName: "test_fq",
						Help:   "test-help",
					},
					Value:     123.45,
					ValueType: prometheus.CounterValue,
					VariableLabelPairs: []*promdto.LabelPair{
						{
							Name:  proto.String("label1"),
							Value: proto.String("value1"),
						},
					},
				},
				{
					MetaDesc: &MetaDesc{
						FQName: "test_fq",
						Help:   "test-help",
					},
					Value:     67.89,
					ValueType: prometheus.CounterValue,
				},
			},
		},
		{
			name: "gauge type metric family",
			metricFamily: &promdto.MetricFamily{
				Name: proto.String("test_fq"),
				Help: proto.String("test-help"),
				Type: (*promdto.MetricType)(proto.Int32(int32(promdto.MetricType_GAUGE))),
				Metric: []*promdto.Metric{
					{
						Gauge: &promdto.Gauge{
							Value: proto.Float64(123.45),
						},
					},
				},
			},
			expectedMetrics: []*Metric{
				{
					MetaDesc: &MetaDesc{
						FQName: "test_fq",
						Help:   "test-help",
					},
					Value:     123.45,
					ValueType: prometheus.GaugeValue,
				},
			},
		},
		{
			name: "untyped type metric family",
			metricFamily: &promdto.MetricFamily{
				Name: proto.String("test_fq"),
				Help: proto.String("test-help"),
				Type: (*promdto.MetricType)(proto.Int32(int32(promdto.MetricType_UNTYPED))),
				Metric: []*promdto.Metric{
					{
						Untyped: &promdto.Untyped{
							Value: proto.Float64(123.45),
						},
					},
				},
			},
			expectedMetrics: []*Metric{
				{
					MetaDesc: &MetaDesc{
						FQName: "test_fq",
						Help:   "test-help",
					},
					Value:     123.45,
					ValueType: prometheus.CounterValue,
				},
			},
		},
		{
			name: "histogram type metric family",
			metricFamily: &promdto.MetricFamily{
				Name: proto.String("test_fq"),
				Help: proto.String("test-help"),
				Type: (*promdto.MetricType)(proto.Int32(int32(promdto.MetricType_HISTOGRAM))),
				Metric: []*promdto.Metric{
					{
						Histogram: &promdto.Histogram{
							SampleCount: proto.Uint64(123),
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			metrics := convertPrometheusDtoToMetric(tt.metricFamily)
			assert.Len(t, metrics, len(tt.expectedMetrics))
			for i := range len(tt.expectedMetrics) {
				assert.NotNil(t, tt.expectedMetrics[i])
				assert.NotNil(t, metrics[i])
				assert.Equal(t, tt.expectedMetrics[i].FQName, metrics[i].FQName)
				assert.Equal(t, tt.expectedMetrics[i].Help, metrics[i].Help)
				assert.Equal(t, tt.expectedMetrics[i].FQName, metrics[i].FQName)
				assert.Equal(t, tt.expectedMetrics[i].Help, metrics[i].Help)
				assert.Equal(t, tt.expectedMetrics[i].VariableLabelPairs, metrics[i].VariableLabelPairs)
			}
		})
	}
}

func TestEfcCollector_queryMetricsFromSocket(t *testing.T) {
	t.Run("should return error for non-existent socket", func(t *testing.T) {
		collector := &efcCollector{}

		nonExistentSocket := filepath.Join(t.TempDir(), "non-existent.sock")
		metrics, err := collector.queryMetricsFromSocket(nonExistentSocket)

		require.Error(t, err)
		assert.Nil(t, metrics)
		assert.Contains(t, err.Error(), "failed to execute request to socket")
	})

	t.Run("should query metrics from socket successfully", func(t *testing.T) {
		tempDir := os.TempDir()
		socketPath := filepath.Join(tempDir, "test.monitor.sock")

		listener, err := net.Listen("unix", socketPath)
		require.NoError(t, err)
		defer func() {
			_ = listener.Close()
		}()

		metricsResponse := `# HELP efc_test_counter Test counter metric
# TYPE efc_test_counter counter
efc_test_counter{label1="value1"} 123.45
# HELP efc_test_gauge Test gauge metric
# TYPE efc_test_gauge gauge
efc_test_gauge{label2="value2"} 67.89
`
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/v1/metrics/prometheus" {
				w.Header().Set("Content-Type", `text/plain; version=0.0.4`)
				w.WriteHeader(http.StatusOK)
				_, err := w.Write([]byte(metricsResponse))
				require.NoError(t, err)
			} else {
				w.WriteHeader(http.StatusNotFound)
			}
		})

		server := &http.Server{Handler: handler}
		go func() {
			_ = server.Serve(listener)
		}()
		defer func() {
			_ = server.Close()
		}()

		// Give the server a moment to start
		time.Sleep(10 * time.Millisecond)

		collector := &efcCollector{}
		metrics, err := collector.queryMetricsFromSocket(socketPath)

		require.NoError(t, err)
		assert.NotNil(t, metrics)
		assert.Len(t, metrics, 2)

		// Check first metric (counter)
		assert.Equal(t, "efc_test_counter", metrics[0].FQName)
		assert.Equal(t, "Test counter metric", metrics[0].Help)
		assert.Equal(t, prometheus.CounterValue, metrics[0].ValueType)
		assert.Equal(t, 123.45, metrics[0].Value)
		assert.Len(t, metrics[0].VariableLabelPairs, 1)
		assert.Equal(t, "label1", *metrics[0].VariableLabelPairs[0].Name)
		assert.Equal(t, "value1", *metrics[0].VariableLabelPairs[0].Value)

		// Check second metric (gauge)
		assert.Equal(t, "efc_test_gauge", metrics[1].FQName)
		assert.Equal(t, "Test gauge metric", metrics[1].Help)
		assert.Equal(t, prometheus.GaugeValue, metrics[1].ValueType)
		assert.Equal(t, 67.89, metrics[1].Value)
		assert.Len(t, metrics[1].VariableLabelPairs, 1)
		assert.Equal(t, "label2", *metrics[1].VariableLabelPairs[0].Name)
		assert.Equal(t, "value2", *metrics[1].VariableLabelPairs[0].Value)
	})
}
