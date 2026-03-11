package metric

import "github.com/prometheus/client_golang/prometheus"

const (
	VolumeStatsCollectorName = "volume_stats"
	VolumeStatsLabelType     = "type"
)

var volumeStatLabels = []string{VolumeStatsLabelType}

type VolumeStatType uint8

type volumeStatCollector struct {
	AttachmentCountMetric     *prometheus.CounterVec
	AttachmentTimeTotalMetric *prometheus.CounterVec
}

const VolumeAttachTimeStat VolumeStatType = 0

var (
	attachmentCountName   = "attachment_count"
	attachmentCountFQName = prometheus.BuildFQName(nodeNamespace, volumeSubsystem, attachmentCountName)
	attachmentCountHelp   = "Volume attachment count."

	attachmentTimeTotalName   = "attachment_time_total"
	attachmentTimeTotalFQName = prometheus.BuildFQName(nodeNamespace, volumeSubsystem, attachmentTimeTotalName)
	attachmentTimeTotalHelp   = "Volume attachment time in total."
)

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

func (c *volumeStatCollector) Get() []*Metric {
	countMetrics := extractMetricsFromMetricVec(attachmentCountFQName, attachmentCountHelp, c.AttachmentCountMetric, prometheus.CounterValue)
	timeMetrics := extractMetricsFromMetricVec(attachmentTimeTotalFQName, attachmentTimeTotalHelp, c.AttachmentTimeTotalMetric, prometheus.CounterValue)
	return append(countMetrics, timeMetrics...)
}

func (c *volumeStatCollector) Update(ch chan<- prometheus.Metric) error {
	c.AttachmentCountMetric.Collect(ch)
	c.AttachmentTimeTotalMetric.Collect(ch)
	return nil
}
