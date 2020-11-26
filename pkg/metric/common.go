package metric

import (
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/procfs"
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"strings"
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
)

const (
	latencyTooHigh    = "LatencyTooHigh"
	capacityNotEnough = "NotEnoughDiskSpace"
)

var (
	metricType       string
	nodeMetricSet    = hashset.New("diskstat", "pfsblockstat")
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

type storageInfo struct {
	PvcNamespace    string
	PvcName         string
	DiskID          string
	DeviceName      string
	VolDataPath     string
	GlobalMountPath string
}

func (d *typedFactorDesc) mustNewConstMetric(value float64, labels ...string) prometheus.Metric {
	if d.factor != 0 {
		value *= d.factor
	}
	return prometheus.MustNewConstMetric(d.desc, d.valueType, value, labels...)
}

func updateMap(clientSet *kubernetes.Clientset, lastPvStorageInfoMap *map[string]storageInfo, jsonPaths []string, deriverName string, keyword string) {
	thisPvStorageInfoMap := make(map[string]storageInfo, 0)
	cmd := "mount | grep csi | grep " + keyword
	line, err := utils.Run(cmd)
	if err != nil {
		logrus.Errorf("Execute cmd %s is failed, err: %s", cmd, err)
		return
	}
	for _, path := range jsonPaths {
		//Get disk pvName
		pvName, diskID, err := getVolumeInfoByJSON(path, deriverName)
		if err != nil {
			logrus.Errorf("Get volume info by path %s is failed, err:%s", path, err)
			continue
		}

		if !strings.Contains(line, pvName) {
			continue
		}

		deviceName, err := getDeviceByVolumeID(pvName, diskID)
		if err != nil {
			logrus.Errorf("Get dev name by diskID %s is failed, err:%s", diskID, err)
			continue
		}
		strorageInfo := storageInfo{
			DiskID:      diskID,
			DeviceName:  deviceName,
			VolDataPath: path,
		}
		thisPvStorageInfoMap[pvName] = strorageInfo
	}

	//If there is a change: add, modify, delete
	updateStorageInfoMap(clientSet, thisPvStorageInfoMap, lastPvStorageInfoMap)
}

func updateStorageInfoMap(clientSet *kubernetes.Clientset, thisPvStorageInfoMap map[string]storageInfo, lastPvStorageInfoMap *map[string]storageInfo) {
	for pv, thisInfo := range thisPvStorageInfoMap {
		lastInfo, ok := (*lastPvStorageInfoMap)[pv]
		// add and modify
		if !ok || thisInfo.VolDataPath != lastInfo.VolDataPath {
			pvcNamespace, pvcName, err := getPvcByPvName(clientSet, pv)
			if err != nil {
				logrus.Errorf("Get pvc by pv %s is failed, err:%s", pv, err.Error())
				continue
			}
			updateInfo := storageInfo{
				DiskID:       thisInfo.DiskID,
				VolDataPath:  thisInfo.VolDataPath,
				DeviceName:   thisInfo.DeviceName,
				PvcName:      pvcName,
				PvcNamespace: pvcNamespace,
			}
			(*lastPvStorageInfoMap)[pv] = updateInfo
		}
	}
	//if pv exist thisPvStorageInfoMap and not exist lastPvStorageInfoMap, pv should be deleted
	for lastPv := range *lastPvStorageInfoMap {
		_, ok := thisPvStorageInfoMap[lastPv]
		if !ok {
			delete(*lastPvStorageInfoMap, lastPv)
		}
	}
}
