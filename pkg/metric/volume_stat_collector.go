package metric

import "github.com/prometheus/client_golang/prometheus"

const (
	VolumeStatsCollectorName = "volume_stats"
	VolumeStatsLabelType     = "type"
	VolumeStatsLabelId       = "id"
	VolumeStatsLabelCode     = "error_code"
)

var volumeStatLabels = []string{VolumeStatsLabelType, VolumeStatsLabelId, VolumeStatsLabelCode}

type VolumeStatType uint8

type volumeStatCollector struct {
	Metrics map[VolumeStatType]*prometheus.HistogramVec
}

const VolumeAttachTimeStat VolumeStatType = 0

var VolumeStatCollector = volumeStatCollector{
	Metrics: map[VolumeStatType]*prometheus.HistogramVec{
		VolumeAttachTimeStat: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: nodeNamespace,
			Subsystem: volumeSubsystem,
			Name:      "attachment_time",
			Help:      "Volume attachment time.",
		}, volumeStatLabels),
	},
}

func init() {
	registerCollector(VolumeStatsCollectorName, NewVolumeStatCollector, diskDriverName, nasDriverName, ossDriverName)
}

func NewVolumeStatCollector() (Collector, error) {
	return &VolumeStatCollector, nil
}

func (c *volumeStatCollector) Update(ch chan<- prometheus.Metric) error {
	for _, metric := range c.Metrics {
		metric.Collect(ch)
	}
	return nil
}
