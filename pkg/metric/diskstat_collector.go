package metric

import (
	"bufio"
	"fmt"
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	"io"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/record"
	"os"
	"strconv"
	"strings"
	"sync"
)

var (
	diskStatLabelNames = []string{"namespace", "pvc", "device", "type"}
)

var (
	//4 - reads completed successfully
	diskReadsCompletedDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "read_completed_total"),
		"The total number of reads completed successfully.",
		diskStatLabelNames, nil,
	)
	//5 - reads merged
	diskReadsMergeDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "read_merged_total"),
		"The total number of reads merged.",
		diskStatLabelNames,
		nil,
	)
	//6 - sectors read
	diskReadBytesDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "read_bytes_total"),
		"The total number of bytes read successfully.",
		diskStatLabelNames, nil,
	)
	//7 - time spent reading (ms)
	diskReadTimeMilliSecondsDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "read_time_milliseconds_total"),
		"The total number of seconds spent by all reads.",
		diskStatLabelNames,
		nil,
	)
	//8 - writes completed
	diskWritesCompletedDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "write_completed_total"),
		"The total number of writes completed successfully.",
		diskStatLabelNames, nil,
	)
	//9 - writes merged
	diskWriteMergeDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "write_merged_total"),
		"The number of writes merged.",
		diskStatLabelNames,
		nil,
	)
	//10 - sectors written
	diskWrittenBytesDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "write_bytes_total"),
		"The total number of bytes written successfully.",
		diskStatLabelNames, nil,
	)
	//11 - time spent writing (ms)
	diskWriteTimeMilliSecondsDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "write_time_milliseconds_total"),
		"This is the total number of seconds spent by all writes.",
		diskStatLabelNames,
		nil,
	)
	//12 - I/Os currently in progress
	diskIONowDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "io_now"),
		"The number of I/Os currently in progress.",
		diskStatLabelNames,
		nil,
	)
	//13 - time spent doing I/Os (ms)
	diskIOTimeSecondsDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "io_time_seconds_total"),
		"Total seconds spent doing I/Os.",
		diskStatLabelNames, nil,
	)
	//13 - capacity available
	diskCapacityAvailableDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "capacity_bytes_available"),
		"The number of available size(bytes).",
		diskStatLabelNames,
		nil,
	)

	//14 - capacity total
	diskCapacityTotalDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "capacity_bytes_total"),
		"The number of total size(bytes).",
		diskStatLabelNames,
		nil,
	)

	//15 - capacity used
	diskCapacityUsedDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "capacity_bytes_used"),
		"The number of used size(bytes).",
		diskStatLabelNames,
		nil,
	)
	//16 - inode available
	diskInodesAvailableDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "inodes_available"),
		"The number of available inodes.",
		diskStatLabelNames,
		nil,
	)
	//17 - inode total
	diskInodesTotalDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "inodes_total"),
		"The number of total inodes.",
		diskStatLabelNames,
		nil,
	)

	//18 - inode used
	diskInodesUsedDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "inodes_used"),
		"The number of used inodes.",
		diskStatLabelNames,
		nil,
	)
)

type diskInfo struct {
	PvcNamespace    string
	PvcName         string
	DiskID          string
	DeviceName      string
	VolDataPath     string
	GlobalMountPath string
}

type diskStatCollector struct {
	alertSwtichSet               *hashset.Set
	milliSecondsLatencyThreshold float64 //Unit: milliseconds
	capacityPercentageThreshold  float64
	descs                        []typedFactorDesc
	lastPvDiskInfoMap            map[string]diskInfo
	lastPvStatsMap               sync.Map
	clientSet                    *kubernetes.Clientset
	recorder                     record.EventRecorder
	nodeName                     string
}

func init() {
	registerCollector("diskstat", NewDiskStatCollector)
}

func parseDiskThreshold(defaultLatencyThreshold float64, defaultCapacityPercentageThreshold float64) (*hashset.Set, float64, float64) {
	alertSet := hashset.New()

	diskLantencyThreshold := strings.ToLower(strings.Trim(os.Getenv("DISK_LATENCY_THRESHOLD"), " "))
	if len(diskLantencyThreshold) != 0 {
		alertSet.Add(latencySwitch)
		defaultLatencyThreshold, _ = parseLantencyThreshold(diskLantencyThreshold, defaultLatencyThreshold)
	}

	diskCapacityThreshold := strings.ToLower(strings.Trim(os.Getenv("DISK_CAPACITY_THRESHOLD_PERCENTAGE"), " "))
	if len(diskCapacityThreshold) != 0 {
		alertSet.Add(capacitySwitch)
		defaultCapacityPercentageThreshold, _ = parseCapacityThreshold(diskCapacityThreshold, defaultCapacityPercentageThreshold)
	}

	return alertSet, defaultLatencyThreshold, defaultCapacityPercentageThreshold
}

// NewDiskStatCollector returns a new Collector exposing disk stats.
func NewDiskStatCollector() (Collector, error) {
	alertSet, latencyThreshold, capacityPercentageThreshold := parseDiskThreshold(diskDefaultsLantencyThreshold, diskDefaultsCapacityPercentageThreshold)
	recorder := utils.NewEventRecorder()
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	nodeName := os.Getenv("KUBE_NODE_NAME")
	return &diskStatCollector{
		descs: []typedFactorDesc{
			//4 - reads completed successfully
			{desc: diskReadsCompletedDesc, valueType: prometheus.CounterValue},
			//5 - reads merged
			{desc: diskReadsMergeDesc, valueType: prometheus.CounterValue},
			//6 - sectors read
			{desc: diskReadBytesDesc, valueType: prometheus.CounterValue, factor: diskSectorSize},
			//7 - time spent reading (ms)
			{desc: diskReadTimeMilliSecondsDesc, valueType: prometheus.CounterValue, factor: .001},
			//8 - writes completed
			{desc: diskWritesCompletedDesc, valueType: prometheus.CounterValue},
			//9 - writes merged
			{desc: diskWriteMergeDesc, valueType: prometheus.CounterValue},
			//10 - sectors written
			{desc: diskWrittenBytesDesc, valueType: prometheus.CounterValue, factor: diskSectorSize},
			//11 - time spent writing (ms)
			{desc: diskWriteTimeMilliSecondsDesc, valueType: prometheus.CounterValue, factor: .001},
			//12 - I/Os currently in progress
			{desc: diskIONowDesc, valueType: prometheus.GaugeValue},
			//13 - time spent doing I/Os (ms)
			{desc: diskIOTimeSecondsDesc, valueType: prometheus.CounterValue, factor: .001},
			//14 - capacity available
			{desc: diskCapacityAvailableDesc, valueType: prometheus.CounterValue},
			//15 - capacity total
			{desc: diskCapacityTotalDesc, valueType: prometheus.CounterValue},
			//16 - capacity used
			{desc: diskCapacityUsedDesc, valueType: prometheus.CounterValue},
			//17 - inode available
			{desc: diskInodesAvailableDesc, valueType: prometheus.CounterValue},
			//18 - inode total
			{desc: diskInodesTotalDesc, valueType: prometheus.CounterValue},
			//19 - inode used
			{desc: diskInodesUsedDesc, valueType: prometheus.CounterValue},
		},
		lastPvDiskInfoMap:            make(map[string]diskInfo, 0),
		lastPvStatsMap:               sync.Map{},
		clientSet:                    clientset,
		milliSecondsLatencyThreshold: latencyThreshold,
		capacityPercentageThreshold:  capacityPercentageThreshold,
		alertSwtichSet:               alertSet,
		recorder:                     recorder,
		nodeName:                     nodeName,
	}, nil
}

func (p *diskStatCollector) Update(ch chan<- prometheus.Metric) error {
	//startTime := time.Now()
	deviceNameStatsMap, err := getDiskStats()
	if err != nil {
		return fmt.Errorf("couldn't get diskstats: %s", err)
	}
	volJSONPaths, err := findVolJSON(podsRootPath)
	if err != nil {
		logrus.Errorf("Find disk vol_data json is failed, err:%s", err)
		return err
	}
	p.updateMap(&p.lastPvDiskInfoMap, volJSONPaths, diskDriverName, "volumes")

	wg := sync.WaitGroup{}
	for deviceName, stats := range deviceNameStatsMap {
		for pvName, info := range p.lastPvDiskInfoMap {
			if info.DeviceName != deviceName {
				continue
			}
			stats, _ := getDiskCapacityMetric(pvName, &info, stats)

			wg.Add(1)
			go func(deviceNameArgs string, pvNameArgs string, pvcNamespaceArgs string, pvcNameArgs string, statsArgs []string) {
				defer wg.Done()
				p.setDiskMetric(deviceNameArgs, pvNameArgs, pvcNamespaceArgs, pvcNameArgs, statsArgs, ch)
			}(deviceName, pvName, info.PvcNamespace, info.PvcName, stats)
		}
	}
	wg.Wait()
	//elapsedTime := time.Since(startTime)
	//logrus.Info("DiskStat spent time:", elapsedTime)
	return nil
}

func isExceedLatencyThreshold(stats []string, lastStats []string, iopsIndex int, latencyIndex int, threshold float64) (float64, bool) {
	thisIOPS, _ := strconv.ParseFloat(stats[iopsIndex], 64)
	lastIOPS, _ := strconv.ParseFloat(lastStats[iopsIndex], 64)
	thisLatency, _ := strconv.ParseFloat(stats[latencyIndex], 64)
	lastLatency, _ := strconv.ParseFloat(lastStats[latencyIndex], 64)
	incrementLatency := thisLatency - lastLatency
	incrementIOPS := thisIOPS - lastIOPS
	if almostEqualFloat64(incrementIOPS, 0) {
		return 0, false
	}
	if (incrementLatency / incrementIOPS) > threshold {
		return incrementLatency / incrementIOPS, true
	}
	return incrementLatency / incrementIOPS, false
}

func isIOHang(stats []string, lastStats []string) bool {
	if len(stats) < 10 || len(lastStats) < 10 {
		logrus.Errorf("stats and last stats array length is less than 10")
		return false
	}
	if (stats[0] == lastStats[0]) &&
		(stats[4] == lastStats[4]) &&
		(stats[9] != lastStats[9]) {
		return true
	}
	return false
}

func (p *diskStatCollector) latencyEventAlert(pvName string, pvcName string, pvcNamespace string, stats []string, index int) {

	lastStats, ok := p.lastPvStatsMap.Load(pvName)
	if p.alertSwtichSet.Contains(latencySwitch) && ok {
		thisLatency, exceed := isExceedLatencyThreshold(stats, lastStats.([]string), index, index+3, p.milliSecondsLatencyThreshold)
		if exceed {
			ref := &v1.ObjectReference{
				Kind:      "PersistentVolumeClaim",
				Name:      pvcName,
				UID:       "",
				Namespace: pvcNamespace,
			}
			reason := fmt.Sprintf("Pvc %s latency load is too high, nodeName: %s, namespace: %s, latency:%.2f ms, threshold:%.2f ms", pvcName, p.nodeName, pvcNamespace, thisLatency, p.milliSecondsLatencyThreshold)
			utils.CreateEvent(p.recorder, ref, v1.EventTypeWarning, latencyTooHigh, reason)
		}
	}
}

func (p *diskStatCollector) ioHangEventAlert(devName string, pvName string, pvcName string, pvcNamespace string, stats []string) {
	lastStats, ok := p.lastPvStatsMap.Load(pvName)
	if ok {
		isHang := isIOHang(stats, lastStats.([]string))
		if isHang {
			ref := &v1.ObjectReference{
				Kind:      "PersistentVolumeClaim",
				Name:      pvcName,
				UID:       "",
				Namespace: pvcNamespace,
			}
			reason := fmt.Sprintf("IO Hang on Persistent Volume %s, nodeName:%s, diskID:%s, Device:%s", pvName, p.nodeName, p.lastPvDiskInfoMap[pvName].DiskID, devName)
			utils.CreateEvent(p.recorder, ref, v1.EventTypeWarning, ioHang, reason)
		}
	}
}

func (p *diskStatCollector) capacityEventAlert(valueFloat64 float64, pvcName string, pvcNamespace string, stats []string) {
	if p.alertSwtichSet.Contains(capacitySwitch) {
		capacityTotalFloat64, err := strconv.ParseFloat(stats[11], 64)
		if err != nil {
			logrus.Errorf("Convert diskCapacityTotalDesc %s to float64 is failed, err:%s", stats[10], err)
			return
		}
		if almostEqualFloat64(capacityTotalFloat64, 0) {
			return
		}
		usedPercentage := (valueFloat64 / capacityTotalFloat64) * 100
		if usedPercentage >= p.capacityPercentageThreshold {
			ref := &v1.ObjectReference{
				Kind:      "PersistentVolumeClaim",
				Name:      pvcName,
				UID:       "",
				Namespace: pvcNamespace,
			}
			reason := fmt.Sprintf("Pvc %s is not enough disk space, nodeName:%s, namespace: %s, usedPercentage:%.2f%%, threshold:%.2f%%", pvcName, p.nodeName, pvcNamespace, usedPercentage, p.capacityPercentageThreshold)
			utils.CreateEvent(p.recorder, ref, v1.EventTypeWarning, capacityNotEnough, reason)
		}
	}
}

func (p *diskStatCollector) setDiskMetric(devName string, pvName string, pvcNamespace string, pvcName string, stats []string, ch chan<- prometheus.Metric) {
	defer p.lastPvStatsMap.Store(pvName, stats)
	for i, value := range stats {
		if i >= len(p.descs) {
			return
		}

		valueFloat64, err := strconv.ParseFloat(value, 64)
		if err != nil {
			logrus.Errorf("Convert value %s to float64 is failed, err:%s", value, err)
			continue
		}
		if i == 3 { //3：diskReadTimeMilliSecondsDesc
			p.latencyEventAlert(pvName, pvcName, pvcNamespace, stats, 0)
		}

		if i == 7 { //7：diskWriteTimeMilliSecondsDesc
			p.latencyEventAlert(pvName, pvcName, pvcNamespace, stats, 4)
		}

		if i == 9 { //9: ioTimeSecondsDesc
			p.ioHangEventAlert(devName, pvName, pvcName, pvcNamespace, stats)
		}

		if i == 12 { //12：diskCapacityUsedDesc
			p.capacityEventAlert(valueFloat64, pvcName, pvcNamespace, stats)
		}

		ch <- p.descs[i].mustNewConstMetric(valueFloat64, pvcNamespace, pvcName, devName, diskStorageName)
	}

}

func (p *diskStatCollector) updateMap(lastPvDiskInfoMap *map[string]diskInfo, jsonPaths []string, deriverName string, keyword string) {
	thisPvDiskInfoMap := make(map[string]diskInfo, 0)
	cmd := "mount | grep csi | grep " + keyword
	line, err := utils.Run(cmd)
	if err != nil && strings.Contains(err.Error(), "with out: , with error:") {
		p.updateDiskInfoMap(thisPvDiskInfoMap, lastPvDiskInfoMap)
		return
	}
	if err != nil {
		logrus.Errorf("Execute cmd %s is failed, err: %s", cmd, err)
		return
	}
	for _, path := range jsonPaths {
		//Get disk pvName
		pvName, diskID, err := getVolumeInfoByJSON(path, deriverName)
		if err != nil {
			if err.Error() != "VolumeType is not the expected type" {
				logrus.Errorf("Get volume info by path %s is failed, err:%s", path, err)
			}
			continue
		}

		if !strings.Contains(line, "/"+pvName+"/") {
			continue
		}

		deviceName, err := getDeviceByVolumeID(pvName, diskID)
		if err != nil {
			logrus.Errorf("Get dev name by diskID %s is failed, err:%s", diskID, err)
			continue
		}
		diskInfo := diskInfo{
			DiskID:      diskID,
			DeviceName:  deviceName,
			VolDataPath: path,
		}
		thisPvDiskInfoMap[pvName] = diskInfo
	}

	//If there is a change: add, modify, delete
	p.updateDiskInfoMap(thisPvDiskInfoMap, lastPvDiskInfoMap)
}

func (p *diskStatCollector) updateDiskInfoMap(thisPvDiskInfoMap map[string]diskInfo, lastPvDiskInfoMap *map[string]diskInfo) {
	for pv, thisInfo := range thisPvDiskInfoMap {
		lastInfo, ok := (*lastPvDiskInfoMap)[pv]
		// add and modify
		if !ok || thisInfo.VolDataPath != lastInfo.VolDataPath {
			pvcNamespace, pvcName, err := getPvcByPvNameByDisk(p.clientSet, pv)
			if err != nil {
				logrus.Errorf("Get pvc by pv %s is failed, err:%s", pv, err.Error())
				continue
			}
			updateInfo := diskInfo{
				DiskID:       thisInfo.DiskID,
				VolDataPath:  thisInfo.VolDataPath,
				DeviceName:   thisInfo.DeviceName,
				PvcName:      pvcName,
				PvcNamespace: pvcNamespace,
			}
			(*lastPvDiskInfoMap)[pv] = updateInfo
		}
	}
	//if pv exist thisPvStorageInfoMap and not exist lastPvDiskInfoMap, pv should be deleted
	for lastPv := range *lastPvDiskInfoMap {
		_, ok := thisPvDiskInfoMap[lastPv]
		if !ok {
			delete(*lastPvDiskInfoMap, lastPv)
		}
	}
}

func getDiskStats() (map[string][]string, error) {
	diskStatsPath := procFilePath(diskStatsFileName)

	file, err := os.Open(diskStatsPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return parseDiskStats(file)
}

func getGlobalMountPathByPvName(pvName string, info *diskInfo) {
	info.GlobalMountPath = fmt.Sprintf("/var/lib/kubelet/plugins/kubernetes.io/csi/pv/%s/globalmount", pvName)
}

func getDiskCapacityMetric(pvName string, info *diskInfo, stat []string) ([]string, error) {
	getGlobalMountPathByPvName(pvName, info)
	response, err := utils.GetMetrics(info.GlobalMountPath)
	if err != nil {
		logrus.Errorf("Get pv %s metrics from kubelet is failed, err: %s", info.GlobalMountPath, err)
		return stat, err
	}
	for _, volumeUsage := range response.Usage {
		stat = append(stat, strconv.FormatInt(volumeUsage.Available, 10))
		stat = append(stat, strconv.FormatInt(volumeUsage.Total, 10))
		stat = append(stat, strconv.FormatInt(volumeUsage.Used, 10))
	}
	return stat, nil
}

func parseDiskStats(r io.Reader) (map[string][]string, error) {
	var (
		diskStats = map[string][]string{}
		scanner   = bufio.NewScanner(r)
	)

	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		if len(parts) < 4 { // we strip major, minor and dev
			return nil, fmt.Errorf("Invalid line in %s: %s", procFilePath(diskStatsFileName), scanner.Text())
		}
		dev := "/dev/" + parts[2]
		diskStats[dev] = parts[3:13]
	}

	return diskStats, scanner.Err()
}
