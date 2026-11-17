package metric

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
	"k8s.io/apimachinery/pkg/util/sets"
)

const (
	VolumeStatsCollectorName = "volume_stats"
	VolumeStatsLabelType     = "type"

	attachmentCountName     = "attachment_count"
	attachmentCountHelp     = "Volume attachment count."
	attachmentTimeTotalName = "attachment_time_total"
	attachmentTimeTotalHelp = "Volume attachment time in total."
)

var volumeStatLabels = []string{VolumeStatsLabelType}

type VolumeStatType uint8

type volumeStatCollector struct {
	AttachmentCountMetric     *prometheus.CounterVec
	AttachmentTimeTotalMetric *prometheus.CounterVec
}

const VolumeAttachTimeStat VolumeStatType = 0

var VolumeStatCollector = volumeStatCollector{
	AttachmentCountMetric: prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: nodeNamespace,
		Subsystem: volumeSubsystem,
		Name:      attachmentCountName,
		Help:      attachmentCountHelp,
	}, volumeStatLabels),
	AttachmentTimeTotalMetric: prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: nodeNamespace,
		Subsystem: volumeSubsystem,
		Name:      attachmentTimeTotalName,
		Help:      attachmentTimeTotalHelp,
	}, volumeStatLabels),
}

func init() {
	registerCollector(VolumeStatsCollectorName, GetVolumeStatCollector, diskDriverName, nasDriverName, ossDriverName)
}

func GetVolumeStatCollector() (Collector, error) {
	return &VolumeStatCollector, nil
}

func (c *volumeStatCollector) Get(ctx context.Context, pvcs sets.Set[string]) ([]*Metric, error) {
	metrics := extractMetricsFromMetricVec(prometheus.BuildFQName(nodeNamespace, volumeSubsystem, attachmentCountName), attachmentCountHelp, c.AttachmentCountMetric, prometheus.CounterValue)
	metrics = append(metrics, extractMetricsFromMetricVec(prometheus.BuildFQName(nodeNamespace, volumeSubsystem, attachmentTimeTotalName), attachmentTimeTotalHelp, c.AttachmentTimeTotalMetric, prometheus.CounterValue)...)
	return metrics, nil
}
