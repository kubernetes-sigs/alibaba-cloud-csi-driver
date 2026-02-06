package metric

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	promdto "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/expfmt"
	"k8s.io/klog/v2"
)

const (
	mountInfoPath     = "/proc/self/mountinfo"
	efcSocketBasePath = "/host/var/run/efc"
	efcMountType      = "fuse.aliyun-alinas-efc"
	efcMetricsPath    = "http://127.0.0.1/v1/metrics/prometheus"
)

type efcCollector struct{}

const efcCollectorName = "efc"

func init() {
	registerCollector(efcCollectorName, NewEfcCollector, bmcpfsDriverName, nasDriverName)
}

var _ Collector = &efcCollector{}

func NewEfcCollector() (Collector, error) {
	return &efcCollector{}, nil
}

func (c *efcCollector) Get() (metrics []*Metric) {
	mountUUIDs, err := c.extractUniqueMountUUIDs(mountInfoPath)
	if err != nil {
		klog.ErrorS(err, "failed to extract unique mount UUIDs")
		return
	}

	monitorSockets, err := c.scanMonitorSockets(efcSocketBasePath, mountUUIDs)
	if err != nil {
		klog.ErrorS(err, "failed to scan monitor sockets")
		return
	}

	for _, socketPath := range monitorSockets {
		socketMetrics, err := c.queryMetricsFromSocket(socketPath)
		if err != nil {
			klog.ErrorS(err, "failed to query metrics from socket", "socket", socketPath)
			continue
		}
		metrics = append(metrics, socketMetrics...)
	}

	return metrics
}

func (c *efcCollector) Update(ch chan<- prometheus.Metric) error {
	metrics := c.Get()
	for _, metric := range metrics {
		ch <- prometheus.MustNewConstMetric(metric.Desc, metric.ValueType, metric.Value, convertLabelsToString(metric.VariableLabelPairs)...)
	}
	return nil
}

// extractUniqueMountUUIDs reads /proc/self/mountinfo and extracts unique mount UUIDs
// for filesystems of type fuse.aliyun-alinas-efc
func (c *efcCollector) extractUniqueMountUUIDs(infoPath string) (map[string]struct{}, error) {
	file, err := os.Open(infoPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open mountinfo: %w", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			klog.ErrorS(err, "failed to close mountinfo file", "path", infoPath)
		}
	}()

	scanner := bufio.NewScanner(file)
	uniqueUUIDs := make(map[string]struct{})

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		fields := strings.Fields(line)
		// Find the filesystem type (comes after the '-' separator)
		// Format: 1851 3839 0:245 / /run/alinas/bindroot/bindroot-ce80f-WPcW4L23 rw,nosuid,nodev,relatime shared:479 - fuse.aliyun-alinas-efc bindroot-ce80f-WPcW4L23:35e9fxxxxx-xxxxx.cn-shanghai.nas.aliyuncs.com:/ rw,user_id=0,xxx
		for i, field := range fields {
			if field != "-" || i+2 >= len(fields) {
				continue
			}

			fsType := fields[i+1]
			if fsType != efcMountType {
				continue
			}

			deviceField := fields[i+2]
			uuidParts := strings.Split(deviceField, ":")
			if len(uuidParts) == 0 {
				continue
			}

			uuid := uuidParts[0]
			if uuid != "" {
				uniqueUUIDs[uuid] = struct{}{}
			}
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error scanning mountinfo: %w", err)
	}

	return uniqueUUIDs, nil
}

// scanMonitorSockets looks for {mountuuid}.monitor.sock files in efcSocketBasePath
func (c *efcCollector) scanMonitorSockets(socketDir string, mountUUIDs map[string]struct{}) ([]string, error) {
	var socketPaths []string

	entries, err := os.ReadDir(socketDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read efc socket directory: %w", err)
	}

	for _, entry := range entries {
		if entry.Type() != os.ModeSocket {
			continue
		}

		fileName := entry.Name()
		uuid := strings.TrimSuffix(fileName, ".monitor.sock")
		if _, ok := mountUUIDs[uuid]; ok {
			socketPath := filepath.Join(socketDir, fileName)
			socketPaths = append(socketPaths, socketPath)
		}
	}

	return socketPaths, nil
}

// queryMetricsFromSocket sends a request to the Unix socket and parses the returned metrics
func (c *efcCollector) queryMetricsFromSocket(socketPath string) (metrics []*Metric, err error) {
	transport := &http.Transport{
		Dial: func(proto, addr string) (net.Conn, error) {
			return net.Dial("unix", socketPath)
		},
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   10 * time.Second,
	}

	req, err := http.NewRequest("GET", efcMetricsPath, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request to socket %s: %w", socketPath, err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			klog.ErrorS(err, "failed to close response body", "socket", socketPath)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request to socket %s returned status code %d", socketPath, resp.StatusCode)
	}

	decoder := expfmt.NewDecoder(resp.Body, expfmt.ResponseFormat(resp.Header))

	for {
		var metricFamily promdto.MetricFamily
		err := decoder.Decode(&metricFamily)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, fmt.Errorf("failed to decode metric: %w", err)
		}

		convertedMetrics := convertPrometheusDtoToMetric(&metricFamily)
		metrics = append(metrics, convertedMetrics...)
	}

	return metrics, nil
}

// convertPrometheusDtoToMetric converts prometheus dto metrics to our internal Metric type
func convertPrometheusDtoToMetric(metricFamily *promdto.MetricFamily) (metrics []*Metric) {
	if metricFamily == nil {
		return
	}
	for _, dtoMetric := range metricFamily.Metric {
		var value float64
		var valueType prometheus.ValueType

		switch metricFamily.GetType() {
		case promdto.MetricType_COUNTER:
			if dtoMetric.Counter != nil {
				value = dtoMetric.Counter.GetValue()
				valueType = prometheus.CounterValue
			}
		case promdto.MetricType_GAUGE:
			if dtoMetric.Gauge != nil {
				value = dtoMetric.Gauge.GetValue()
				valueType = prometheus.GaugeValue
			}
		case promdto.MetricType_UNTYPED:
			if dtoMetric.Untyped != nil {
				value = dtoMetric.Untyped.GetValue()
				valueType = prometheus.GaugeValue // treat untyped metrics as gauges
				if strings.HasSuffix(metricFamily.GetName(), "_count") {
					valueType = prometheus.CounterValue
				}
			}
		default:
			continue // skip unsupported types
		}

		metric := NewMetricWithLabelPairs(metricFamily.GetName(), metricFamily.GetHelp(), value, valueType, dtoMetric.Label)
		metrics = append(metrics, metric)
	}

	return metrics
}
