package metric

import (
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/procfs"
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
)

const (
	//OssStorageName represents the storage type name of Oss
	ossStorageName string = "oss"
	//NasStorageName represents the storage type name of Nas
	nasStorageName string = "nas"
	//diskStorageName represents the storage type name of Disk
	diskStorageName string = "disk"
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
	diskSectorSize = 512
)

const (
	diskStatsFileName = "diskstats"
)

var (
	metricType       string
	nodeMetricSet    = hashset.New("diskstat","rdsrawblockstat")
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

func updateMapping(clientSet *kubernetes.Clientset, lastPvPathMapping *map[string]string, lastPvPvcMapping *map[string][]string, deriverNamer string)(map[string]string, error){
	volDataJSONPath, err := findVolDataJSONFileByPattern(podsRootPath)
	if err != nil {
		return nil, err
	}

	pvDeviceNameMapping := make(map[string]string, 0)
	thisPvPathMapping := make(map[string]string, 0)
	for _, path := range volDataJSONPath {
		//Get disk pvName
		pvName, err := getVolumeIDByJSON(path, deriverNamer)
		if err != nil {
			continue
		}
		pvDevice, err := getDeviceByVolumeID(pvName)
		if err != nil {
			continue
		}
		thisPvPathMapping[pvName] = path
		pvDeviceNameMapping[pvDevice] = pvName
	}

	//If there is a change:add, modify, delete
	updateLastPvcMapping(clientSet, thisPvPathMapping, lastPvPathMapping, lastPvPvcMapping)
	return thisPvPathMapping, nil
}

func updateLastPvcMapping(clientSet *kubernetes.Clientset, thisPvPathMapping map[string]string, lastPvPathMapping *map[string]string,  lastPvPvcMapping *map[string][]string) {
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
