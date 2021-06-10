package metric

import (
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/procfs"
)

const (
	//OssStorageName represents the storage type name of Oss
	ossStorageName string = "oss"
	//NasStorageName represents the storage type name of Nas
	nasStorageName string = "nas"
	//diskStorageName represents the storage type name of Disk
	diskStorageName string = "disk"
	pfsBlockName    string = "pfsblock"
	//unknownStorageName represents the storage type name of Unknown
	unknownStorageName string = "unknown"
	//ossDriverName represents the csi storage type name of Oss
	ossDriverName string = "ossplugin.csi.alibabacloud.com"
	//nasDriverName represents the csi storage type name of Nas
	nasDriverName string = "nasplugin.csi.alibabacloud.com"
	//diskDriverName represents the csi storage type name of Disk
	diskDriverName string = "diskplugin.csi.alibabacloud.com"
)

const (
	clusterNamespace string = "cluster"
	nodeNamespace    string = "node"
)

const (
	scrapeSubSystem string = "scrape"
	volumeSubSystem string = "volume"
)

const (
	latencySwitch  = "latency"
	capacitySwitch = "capacity"
)
const (
	diskSectorSize                          = 512
	diskDefaultsLantencyThreshold           = 10
	diskDefaultsCapacityPercentageThreshold = 85
	float64EqualityThreshold                = 1e-9
)

const (
	diskStatsFileName = "diskstats"
	nfsStatsFileName  = "/proc/self/mountstats"
)

const (
	latencyTooHigh    = "LatencyTooHigh"
	capacityNotEnough = "NotEnoughDiskSpace"
	ioHang            = "IOHang"
)

var (
	metricType       string
	nodeMetricSet    = hashset.New("diskstat", "pfsblockstat", "nfsstat")
	clusterMetricSet = hashset.New("")
)

const (
	// PluginService represents the csi-plugin type.
	pluginService = "plugin"
	// ProvisionerService represents the csi-provisioner type.
	provisionerService = "provisioner"
)

const (
	volDataFile      = "vol_data.json"
	csiMountKeyWords = "volumes/kubernetes.io~csi"
	procPath         = procfs.DefaultMountPoint + "/"
	rawBlockRootPath = "/var/lib/kubelet/plugins/kubernetes.io/csi/volumeDevices/"
	podsRootPath     = "/var/lib/kubelet/pods"
)

type collectorFactoryFunc = func() (Collector, error)

//csiCollectorInstance is a single instance of CSICollector
//Factories are the mapping between monitoring types and collectorFactoryFunc
var (
	csiCollectorInstance *CSICollector
	factories            = make(map[string]collectorFactoryFunc)
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
