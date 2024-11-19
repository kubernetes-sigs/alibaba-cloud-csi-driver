package metric

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/transport"
	"k8s.io/klog/v2"
	statsapi "k8s.io/kubelet/pkg/apis/stats/v1alpha1"
)

const (
	saTokenFile            = "/var/run/secrets/kubernetes.io/serviceaccount/token"
	kubeletStatsSummaryUrl = "https://localhost:10250/stats/summary"
	kubeletHttpTimeout     = time.Second * 5
)

type fsStatsMetric struct {
	desc     *prometheus.Desc
	getValue func(*statsapi.FsStats) *float64
}

func (m *fsStatsMetric) Metric(fsStats *statsapi.FsStats, labels ...string) prometheus.Metric {
	value := m.getValue(fsStats)
	if value == nil {
		return nil
	}
	return prometheus.MustNewConstMetric(m.desc, prometheus.GaugeValue, *value, labels...)
}

func uint64ToFloat64(value *uint64) *float64 {
	if value == nil {
		return nil
	}
	fValue := float64(*value)
	return &fValue
}

func generateFsStatsDescs(namespace, subsystem string, labels []string) []*fsStatsMetric {
	return []*fsStatsMetric{
		{
			desc:     prometheus.NewDesc(prometheus.BuildFQName(namespace, subsystem, "inodes_free"), "Number of available Inodes", labels, nil),
			getValue: func(fs *statsapi.FsStats) *float64 { return uint64ToFloat64(fs.InodesFree) },
		},
		{
			desc:     prometheus.NewDesc(prometheus.BuildFQName(namespace, subsystem, "inodes_total"), "Total number of Inodes", labels, nil),
			getValue: func(fs *statsapi.FsStats) *float64 { return uint64ToFloat64(fs.Inodes) },
		},
		{
			desc:     prometheus.NewDesc(prometheus.BuildFQName(namespace, subsystem, "inodes_used"), "Number of used Inodes", labels, nil),
			getValue: func(fs *statsapi.FsStats) *float64 { return uint64ToFloat64(fs.InodesUsed) },
		},
		{
			desc:     prometheus.NewDesc(prometheus.BuildFQName(namespace, subsystem, "limit_bytes"), "Number of bytes that can be consumed by the container on this filesystem", labels, nil),
			getValue: func(fs *statsapi.FsStats) *float64 { return uint64ToFloat64(fs.CapacityBytes) },
		},
		{
			desc:     prometheus.NewDesc(prometheus.BuildFQName(namespace, subsystem, "usage_bytes"), "Number of bytes that are consumed by the container on this filesystem", labels, nil),
			getValue: func(fs *statsapi.FsStats) *float64 { return uint64ToFloat64(fs.UsedBytes) },
		},
		{
			desc:     prometheus.NewDesc(prometheus.BuildFQName(namespace, subsystem, "available_bytes"), "Number of bytes that not consumed", labels, nil),
			getValue: func(fs *statsapi.FsStats) *float64 { return uint64ToFloat64(fs.AvailableBytes) },
		},
	}
}

var (
	rootfsMetrics           = generateFsStatsDescs("container", "fs", []string{"namespace", "pod", "pod_uid", "container"})
	ephemeralStorageMetrics = generateFsStatsDescs("ephemeral_storage", "pod", []string{"namespace", "pod", "pod_uid"})
)

func init() {
	registerCollector("kubelet_stats_summary", NewKubeletStatsSummaryCollector)
}

type kubeletStatsSummaryCollector struct {
	client *http.Client
}

func (c *kubeletStatsSummaryCollector) Update(ch chan<- prometheus.Metric) error {
	resp, err := c.client.Get(kubeletStatsSummaryUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		klog.V(4).InfoS("Failed to get kubelet stats summary", "status_code", resp.StatusCode, "body", string(body))
		return nil
	}
	var summary statsapi.Summary
	if err := json.NewDecoder(resp.Body).Decode(&summary); err != nil {
		return err
	}
	for _, pod := range summary.Pods {
		if pod.EphemeralStorage != nil {
			for _, m := range ephemeralStorageMetrics {
				metric := m.Metric(pod.EphemeralStorage, pod.PodRef.Namespace, pod.PodRef.Name, pod.PodRef.UID)
				if metric != nil {
					ch <- metric
				}
			}
		}

		for _, container := range pod.Containers {
			if container.Rootfs != nil {
				for _, m := range rootfsMetrics {
					metric := m.Metric(container.Rootfs, pod.PodRef.Namespace, pod.PodRef.Name, pod.PodRef.UID, container.Name)
					if metric != nil {
						ch <- metric
					}
				}
			}
		}
	}
	return nil
}

func NewKubeletStatsSummaryCollector() (Collector, error) {
	config := &transport.Config{
		UserAgent: rest.DefaultKubernetesUserAgent(),
		TLS: transport.TLSConfig{
			// kubelet cert SANs do not contains localhost or 127.0.0.1
			Insecure: true,
		},
		BearerTokenFile: saTokenFile,
	}
	tr, err := transport.New(config)
	if err != nil {
		return nil, err
	}

	collector := &kubeletStatsSummaryCollector{
		client: &http.Client{
			Transport: tr,
			Timeout:   kubeletHttpTimeout,
		},
	}
	return collector, nil
}
