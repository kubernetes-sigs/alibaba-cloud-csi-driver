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

var VolumeStatCollector = volumeStatCollector{
	AttachmentCountMetric: prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: nodeNamespace,
		Subsystem: volumeSubsystem,
		Name:      "attachment_count",
		Help:      "Volume attachment count.",
	}, volumeStatLabels),
	AttachmentTimeTotalMetric: prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: nodeNamespace,
		Subsystem: volumeSubsystem,
		Name:      "attachment_time_total",
		Help:      "Volume attachment time in total.",
	}, volumeStatLabels),
}

func init() {
	registerCollector(VolumeStatsCollectorName, GetVolumeStatCollector, diskDriverName, nasDriverName, ossDriverName)
}

func GetVolumeStatCollector() (Collector, error) {
	return &VolumeStatCollector, nil
}

func (c *volumeStatCollector) Update(ch chan<- prometheus.Metric) error {
	c.AttachmentCountMetric.Collect(ch)
	c.AttachmentTimeTotalMetric.Collect(ch)
	return nil
}
