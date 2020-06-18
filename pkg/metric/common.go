package metric

import (
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/procfs"
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
)

const (
	OssStorageName     string = "oss"
	NasStorageName     string = "nas"
	DiskStorageName    string = "disk"
	unknownStorageName string = "unknown"
	ossDriverName      string = "ossplugin.csi.alibabacloud.com"
	nasDriverName      string = "nasplugin.csi.alibabacloud.com"
	diskDriverName     string = "diskplugin.csi.alibabacloud.com"
)

const (
	defaultNamespace string = "default"
	clusterNamespace string = "cluster"
	nodeNamespace    string = "node"
)

const (
	scrapeSubSystem                string = "scrape"
	volumeSubSystem                string = "volume"
	persistentVolumeSubSystem      string = "pv"
	persistentVolumeClaimSubSystem string = "pvc"
)

const (
	diskSectorSize  = 512
	counterInitSize = 1
)

const (
	diskStatsFileName = "diskstats"
)

var (
	shellType        = "bash"
	metricType       string
	nodeMetricSet    = hashset.New("diskstat")
	clusterMetricSet = hashset.New("pv", "pvc")
)

const (
	nodeServer    = "node-server"
	clusterServer = "cluster-server"
)

const (
	volDataFile      = "vol_data.json"
	csiMountKeyWords = "kubernetes.io~csi"
	procPath         = procfs.DefaultMountPoint + "/"
	podsRootPath     = "/var/lib/kubelet/pods"
)

const (
	NodeStageVolumeAction     = "NodeStageVolume"
	NodeUnstageVolumeAction   = "NodeUnstageVolume"
	NodePublishVolumeAction   = "NodePublishVolume"
	NodeUnPublishVolumeAction = "NodeUnPublishVolume"
	CreateVolumeAction        = "CreateVolume"
	DeleteVolumeAction        = "DeleteVolume"
)

var (
	diskStatLabelNames              = []string{"namespace", "persistentvolumeclaim", "device", "type"}
	nasStatLabelNames               = []string{"persistentvolumeclaim", "type"}
	persistentVolumeLabelNames      = []string{"type", "status"}
	persistentVolumeClaimLabelNames = []string{"type", "status"}
	actionLabelNames                = []string{"persistentvolumeclaim", "type", "action"}
)

var (
	ActionCollectorInstance *actionCollector = nil
	CSICollectorInstance    *CSICollector    = nil
	factories                                = make(map[string]func() (Collector, error))
)

var (
	scrapeDurationDesc = prometheus.NewDesc(
		prometheus.BuildFQName(clusterNamespace, scrapeSubSystem, "collector_duration_seconds"),
		"csi_metric: Duration of a collector scrape.",
		[]string{"collector"},
		nil,
	)
	scrapeSuccessDesc = prometheus.NewDesc(
		prometheus.BuildFQName(clusterNamespace, scrapeSubSystem, "collector_success"),
		"csi_metric: Whether a collector succeeded.",
		[]string{"collector"},
		nil,
	)
	persistentVolumeCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(clusterNamespace, persistentVolumeSubSystem, "counter"),
		"The total number of Persistent Volume",
		persistentVolumeLabelNames, nil,
	)
	persistentVolumeClaimCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(clusterNamespace, persistentVolumeClaimSubSystem, "counter"),
		"The total number of Persistent Volume Claim",
		persistentVolumeClaimLabelNames, nil,
	)
	//4 - reads completed successfully
	diskReadsCompletedDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "disk_read_complete_total"),
		"The total number of reads completed successfully.",
		diskStatLabelNames, nil,
	)
	//5 - reads merged
	diskReadsMergeDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "disk_read_merge_total"),
		"The total number of reads merged.",
		diskStatLabelNames,
		nil,
	)
	//6 - sectors read
	diskReadBytesDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "disk_read_byte_total"),
		"The total number of bytes read successfully.",
		diskStatLabelNames, nil,
	)
	//7 - time spent reading (ms)
	diskReadTimeSecondsDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "disk_read_time_second_total"),
		"The total number of seconds spent by all reads.",
		diskStatLabelNames,
		nil,
	)
	//8 - writes completed
	diskWritesCompletedDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "disk_write_complete_total"),
		"The total number of writes completed successfully.",
		diskStatLabelNames, nil,
	)
	//9 - writes merged
	diskWriteMergeDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "disk_write_merge_total"),
		"The number of writes merged.",
		diskStatLabelNames,
		nil,
	)
	//10 - sectors written
	diskWrittenBytesDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "disk_write_byte_total"),
		"The total number of bytes written successfully.",
		diskStatLabelNames, nil,
	)

	//11 - time spent writing (ms)
	diskWriteTimeSecondsDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "disk_write_time_second_total"),
		"This is the total number of seconds spent by all writes.",
		diskStatLabelNames,
		nil,
	)
	//12 - I/Os currently in progress
	diskIONowDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "disk_io_now"),
		"The number of I/Os currently in progress.",
		diskStatLabelNames,
		nil,
	)
	//disk action Duration
	actionTimeSecondsDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "action_time_second_total"),
		"This is the total number of seconds spent by this pv.",
		actionLabelNames,
		nil,
	)

	//4 - reads completed successfully
	nasReadsCompletedDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "nas_read_complete_total"),
		"The total number of reads completed successfully.",
		nasStatLabelNames, nil,
	)
	//6 - sectors read
	nasReadBytesDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "nas_read_byte_total"),
		"The total number of bytes read successfully.",
		nasStatLabelNames, nil,
	)
	//7 - time spent reading (ms)
	nasReadTimeSecondsDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "nas_read_time_second_total"),
		"The total number of seconds spent by all reads.",
		nasStatLabelNames,
		nil,
	)
	//8 - writes completed
	nasWritesCompletedDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "nas_writes_complete_total"),
		"The total number of writes completed successfully.",
		nasStatLabelNames, nil,
	)
	//10 - sectors written
	nasWrittenBytesDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "nas_written_byte_total"),
		"The total number of bytes written successfully.",
		nasStatLabelNames, nil,
	)
	//11 - time spent writing (ms)
	nasWriteTimeSecondsDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "nas_write_time_second_total"),
		"This is the total number of seconds spent by all writes.",
		nasStatLabelNames,
		nil,
	)
)

type typedFactorDesc struct {
	desc      *prometheus.Desc
	valueType prometheus.ValueType
	factor    float64
}

func (d *typedFactorDesc) mustNewConstMetric(value float64, labels ...string) prometheus.Metric {
	if d.factor != 0 {
		value *= d.factor
	}
	return prometheus.MustNewConstMetric(d.desc, d.valueType, value, labels...)
}

func updateLastPvcMapping(thisPvPathMapping map[string]string, lastPvPathMapping *map[string]string, clientSet *kubernetes.Clientset, lastPvPvcMapping *map[string][]string) {
	for thisKey, thisValue := range thisPvPathMapping {
		lastValue, ok := (*lastPvPathMapping)[thisKey]
		if !ok || thisValue != lastValue {
			pvcNamespace, pvcName, err := getPvcByPvName(clientSet, thisKey)
			if err != nil {
				logrus.Errorf("GetPvcByPvName err:%s", err.Error())
				continue
			}
			(*lastPvPathMapping)[thisKey] = thisValue
			(*lastPvPvcMapping)[thisKey] = []string{pvcNamespace, pvcName}
		}
	}
	for lastKey := range *lastPvPvcMapping {
		_, ok := thisPvPathMapping[lastKey]
		if !ok {
			delete(*lastPvPvcMapping, lastKey)
			delete(*lastPvPathMapping, lastKey)
		}
	}
}
