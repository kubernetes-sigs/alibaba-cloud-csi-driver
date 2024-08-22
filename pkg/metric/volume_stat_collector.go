package metric

import "github.com/prometheus/client_golang/prometheus"

const (
	VolumeStatCollectorName    = "volume_stat"
	VolumeStatTypeLabelName    = "type"
	VolumeStatIdLabelName      = "id"
	VolumeStatErrCodeLabelName = "error_code"
)

var volumeStatLabels = []string{VolumeStatTypeLabelName, VolumeStatIdLabelName, VolumeStatErrCodeLabelName}

type VolumeStatType uint8

type volumeStatCollector struct {
	Metrics map[VolumeStatType]*prometheus.HistogramVec
}

const (
	VolumeAttachTimeStat VolumeStatType = iota
	ControllerPublishExecTimeStat
	ControllerUnpublishExecTimeStat
	NodeStageExecTimeStat
	NodeUnstageExecTimeStat
	NodePublishExecTimeStat
	NodeUnpublishExecTimeStat
)

var VolumeStatCollector = volumeStatCollector{
	Metrics: map[VolumeStatType]*prometheus.HistogramVec{
		VolumeAttachTimeStat: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: csiNamespace,
			Subsystem: execTimeSubsystem,
			Name:      "volume_attachment",
			Help:      "Volume attachment time.",
		}, volumeStatLabels),
		ControllerPublishExecTimeStat: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: csiNamespace,
			Subsystem: execTimeSubsystem,
			Name:      "controller_publish",
			Help:      "ControllerPublishVolume execution time.",
		}, volumeStatLabels),
		ControllerUnpublishExecTimeStat: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: csiNamespace,
			Subsystem: execTimeSubsystem,
			Name:      "controller_unpublish",
			Help:      "ControllerUnublishVolume execution time.",
		}, volumeStatLabels),
		NodeStageExecTimeStat: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: csiNamespace,
			Subsystem: execTimeSubsystem,
			Name:      "node_stage",
			Help:      "NodeStageVolume execution time.",
		}, volumeStatLabels),
		NodeUnstageExecTimeStat: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: csiNamespace,
			Subsystem: execTimeSubsystem,
			Name:      "node_unstage",
			Help:      "NodeUnstageVolume execution time.",
		}, volumeStatLabels),
		NodePublishExecTimeStat: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: csiNamespace,
			Subsystem: execTimeSubsystem,
			Name:      "node_publish",
			Help:      "NodePublishVolume execution time.",
		}, volumeStatLabels),
		NodeUnpublishExecTimeStat: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: csiNamespace,
			Subsystem: execTimeSubsystem,
			Name:      "node_unpublish",
			Help:      "NodeUnpublishVolume execution time.",
		}, volumeStatLabels),
	},
}

func init() {
	registerCollector(VolumeStatCollectorName, NewVolumeStatCollector, diskDriverName, nasDriverName, ossDriverName)
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
