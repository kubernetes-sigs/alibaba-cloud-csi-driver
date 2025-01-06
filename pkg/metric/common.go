package metric

import (
	"path/filepath"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/procfs"
)

const (
	//OssStorageName represents the storage type name of Oss
	ossStorageName string = "oss"
	//NasStorageName represents the storage type name of Nas
	nasStorageName string = "nas"
	//aliNasStorageName represents the storage type name of AliNas
	aliNasStorageName string = "alinas"
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
	diskDriverName  string = "diskplugin.csi.alibabacloud.com"
	localDriverName string = "localplugin.csi.alibabacloud.com"
	// unknown metric value
	UnknownValue string = "unknown"
)

const (
	clusterNamespace                        = "cluster"
	csiNamespace                            = "csi"
	nodeNamespace                           = "node"
	grpcSubsystem                           = "grpc"
	scrapeSubsystem                         = "scrape"
	volumeSubsystem                         = "volume"
	diskSectorSize                          = 512
	diskDefaultsLantencyThreshold           = 10
	diskDefaultsCapacityPercentageThreshold = 85
	nfsDefaultsCapacityPercentageThreshold  = 85
	float64EqualityThreshold                = 1e-9
	diskStatsFileName                       = "diskstats"
	nfsStatsFileName                        = "/proc/self/mountstats"
	latencyTooHigh                          = "LatencyTooHigh"
	capacityNotEnough                       = "NotEnoughDiskSpace"
	ioHang                                  = "IOHang"
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
)

// collectorFactoryFunc can return a nil Collector if it wants to be ignored
type collectorFactoryFunc = func() (Collector, error)

type collectorRegistryItem struct {
	Name           string
	Factory        collectorFactoryFunc
	RelatedDrivers []string
}

// csiCollectorInstance is a single instance of CSICollector
// Factories are the mapping between monitoring types and collectorFactoryFunc
var (
	csiCollectorInstance *CSICollector
	registry             = []collectorRegistryItem{}
	rawBlockRootPath     = filepath.Join(utils.KubeletRootDir, "/plugins/kubernetes.io/csi/volumeDevices/")
	podsRootPath         = filepath.Join(utils.KubeletRootDir, "/pods")
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
