package metric

import (
	"path/filepath"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	//NasStorageName represents the storage type name of Nas
	nasStorageName string = "nas"
	//diskStorageName represents the storage type name of Disk
	diskStorageName string = "disk"
	//ossDriverName represents the csi storage type name of Oss
	ossDriverName string = "ossplugin.csi.alibabacloud.com"
	//nasDriverName represents the csi storage type name of Nas
	nasDriverName    string = "nasplugin.csi.alibabacloud.com"
	bmcpfsDriverName string = "bmcpfsplugin.csi.alibabacloud.com"
	//diskDriverName represents the csi storage type name of Disk
	diskDriverName string = "diskplugin.csi.alibabacloud.com"
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
	diskDefaultsLatencyThreshold            = 10
	diskDefaultsCapacityPercentageThreshold = 85
	nfsDefaultsCapacityPercentageThreshold  = 85
	nfsStatsFileName                        = "/proc/self/mountstats"
	latencyTooHigh                          = "LatencyTooHigh"
	capacityNotEnough                       = "NotEnoughDiskSpace"
	ioHang                                  = "IOHang"
)

const (
	volDataFile      = "vol_data.json"
	csiMountKeyWords = "volumes/kubernetes.io~csi"
)

type collectorFactoryFunc = func() (Collector, error)
type collectorRegistryItem struct {
	Name           string
	Factory        collectorFactoryFunc
	RelatedDrivers []string
}

var (
	registry     = []collectorRegistryItem{}
	podsRootPath = filepath.Join(utils.KubeletRootDir, "/pods")
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
