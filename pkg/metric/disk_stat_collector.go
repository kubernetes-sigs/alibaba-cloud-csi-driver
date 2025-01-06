package metric

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/record"
	"k8s.io/klog/v2"
	"k8s.io/mount-utils"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/options"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
)

var (
	diskStatLabelNames = []string{"namespace", "pvc", "device", "type"}
	scalerPvcMap       *sync.Map
)

var (
	// 4 - reads completed successfully
	diskReadsCompletedDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "read_completed_total"),
		"The total number of reads completed successfully.",
		diskStatLabelNames, nil,
	)
	// 5 - reads merged
	diskReadsMergeDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "read_merged_total"),
		"The total number of reads merged.",
		diskStatLabelNames,
		nil,
	)
	// 6 - sectors read
	diskReadBytesDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "read_bytes_total"),
		"The total number of bytes read successfully.",
		diskStatLabelNames, nil,
	)
	// 7 - time spent reading (ms)
	diskReadTimeMilliSecondsDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "read_time_milliseconds_total"),
		"The total number of seconds spent by all reads.",
		diskStatLabelNames,
		nil,
	)
	// 8 - writes completed
	diskWritesCompletedDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "write_completed_total"),
		"The total number of writes completed successfully.",
		diskStatLabelNames, nil,
	)
	//9 - writes merged
	diskWriteMergeDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "write_merged_total"),
		"The number of writes merged.",
		diskStatLabelNames,
		nil,
	)
	//10 - sectors written
	diskWrittenBytesDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "write_bytes_total"),
		"The total number of bytes written successfully.",
		diskStatLabelNames, nil,
	)
	//11 - time spent writing (ms)
	diskWriteTimeMilliSecondsDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "write_time_milliseconds_total"),
		"This is the total number of seconds spent by all writes.",
		diskStatLabelNames,
		nil,
	)
	//12 - I/Os currently in progress
	diskIONowDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "io_now"),
		"The number of I/Os currently in progress.",
		diskStatLabelNames,
		nil,
	)
	//13 - time spent doing I/Os (ms)
	diskIOTimeSecondsDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "io_time_seconds_total"),
		"Total seconds spent doing I/Os.",
		diskStatLabelNames, nil,
	)
	//13 - capacity available
	diskCapacityAvailableDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "capacity_bytes_available"),
		"The number of available size(bytes).",
		diskStatLabelNames,
		nil,
	)

	//14 - capacity total
	diskCapacityTotalDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "capacity_bytes_total"),
		"The number of total size(bytes).",
		diskStatLabelNames,
		nil,
	)

	//15 - capacity used
	diskCapacityUsedDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "capacity_bytes_used"),
		"The number of used size(bytes).",
		diskStatLabelNames,
		nil,
	)
	//16 - inode available
	diskInodesAvailableDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "inodes_available"),
		"The number of available inodes.",
		diskStatLabelNames,
		nil,
	)
	//17 - inode total
	diskInodesTotalDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "inodes_total"),
		"The number of total inodes.",
		diskStatLabelNames,
		nil,
	)

	//18 - inode used
	diskInodesUsedDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "inodes_used"),
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
	milliSecondsLatencyThreshold float64 //Unit: milliseconds
	capacityPercentageThreshold  float64
	descs                        []typedFactorDesc
	pvInfoLock                   sync.Mutex
	lastPvDiskInfoMap            map[string]diskInfo
	lastPvStatsMap               sync.Map
	clientSet                    *kubernetes.Clientset
	recorder                     record.EventRecorder
	mounter                      mount.Interface
	nodeName                     string
}

func init() {
	registerCollector("disk_stat", NewDiskStatCollector, diskDriverName)
}

func getDiskLatencyThreshold() float64 {
	latencyStr := strings.ToLower(strings.Trim(os.Getenv("DISK_LATENCY_THRESHOLD"), " "))
	if len(latencyStr) != 0 {
		latency, _ := parseLantencyThreshold(latencyStr, diskDefaultsLantencyThreshold)
		return latency
	}
	return 0
}

func getDiskCapacityThreshold() float64 {
	capacityStr := strings.ToLower(strings.Trim(os.Getenv("DISK_CAPACITY_THRESHOLD_PERCENTAGE"), " "))
	if len(capacityStr) != 0 {
		capacity, _ := parseCapacityThreshold(capacityStr, diskDefaultsCapacityPercentageThreshold)
		return capacity
	}
	return 0
}

// NewDiskStatCollector returns a new Collector exposing disk stats.
func NewDiskStatCollector() (Collector, error) {
	recorder := utils.NewEventRecorder()
	config, err := clientcmd.BuildConfigFromFlags(options.MasterURL, options.Kubeconfig)
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
		milliSecondsLatencyThreshold: getDiskLatencyThreshold(),
		capacityPercentageThreshold:  getDiskCapacityThreshold(),
		recorder:                     recorder,
		mounter:                      mount.NewWithoutSystemd(""),
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
		return err
	}
	p.updateMap(&p.lastPvDiskInfoMap, volJSONPaths, diskDriverName)

	wg := sync.WaitGroup{}
	for pvName, info := range p.lastPvDiskInfoMap {
		stats, ok := deviceNameStatsMap[info.DeviceName]
		if !ok {
			continue
		}
		stats, _ = getDiskCapacityMetric(&info, stats)
		if scalerPvcMap != nil {
			if _, ok := scalerPvcMap.Load(info.PvcName); !ok {
				continue
			}
		}
		wg.Add(1)
		go func(pvName string, info diskInfo, stats []string) {
			defer wg.Done()
			p.setDiskMetric(pvName, info, stats, ch)
		}(pvName, info, stats)
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

func (p *diskStatCollector) latencyEventAlert(pvName string, pvcName string, pvcNamespace string, stats []string, index int) {
	lastStats, ok := p.lastPvStatsMap.Load(pvName)
	if ok && p.milliSecondsLatencyThreshold > 0 {
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

func (p *diskStatCollector) capacityEventAlert(valueFloat64 float64, pvcName string, pvcNamespace string, stats []string) {
	if p.capacityPercentageThreshold > 0 {
		capacityTotalFloat64, err := strconv.ParseFloat(stats[11], 64)
		if err != nil {
			klog.Errorf("Convert diskCapacityTotalDesc %s to float64 is failed, err:%s", stats[10], err)
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

func (p *diskStatCollector) setDiskMetric(pvName string, info diskInfo, stats []string, ch chan<- prometheus.Metric) {
	defer p.lastPvStatsMap.Store(pvName, stats)
	for i, value := range stats {
		if i >= len(p.descs) {
			return
		}

		valueFloat64, err := strconv.ParseFloat(value, 64)
		if err != nil {
			klog.Errorf("Convert value %s to float64 is failed, err:%s", value, err)
			continue
		}
		if i == 3 { //3: diskReadTimeMilliSecondsDesc
			p.latencyEventAlert(pvName, info.PvcName, info.PvcNamespace, stats, 0)
		}

		if i == 7 { //7: diskWriteTimeMilliSecondsDesc
			p.latencyEventAlert(pvName, info.PvcName, info.PvcNamespace, stats, 4)
		}

		if i == 12 { //12: diskCapacityUsedDesc
			p.capacityEventAlert(valueFloat64, info.PvcName, info.PvcNamespace, stats)
		}

		ch <- p.descs[i].mustNewConstMetric(valueFloat64, info.PvcNamespace, info.PvcName, info.DeviceName, diskStorageName)
	}

}

func (p *diskStatCollector) updateMap(lastPvDiskInfoMap *map[string]diskInfo, jsonPaths []string, driverName string) {
	thisPvDiskInfoMap := make(map[string]diskInfo, 0)
	for _, path := range jsonPaths {
		//Get disk pvName
		pvName, diskID, err := getVolumeInfoByJSON(path, driverName)
		if err != nil {
			if err != ErrUnexpectedVolumeType {
				klog.Errorf("Get volume info by path %s is failed, err:%s", path, err)
			}
			continue
		}

		mountPoint := filepath.Join(path, "../mount")
		notMounted, err := p.mounter.IsLikelyNotMountPoint(mountPoint)
		if err != nil {
			if !errors.Is(err, fs.ErrNotExist) {
				klog.Errorf("Check if %s is mount point failed: %v", mountPoint, err)
			}
			continue
		}
		if notMounted {
			continue // may be leaked volume
		}

		deviceName, err := getDeviceByVolumeID(pvName, diskID)
		if err != nil {
			klog.Errorf("Get dev name by diskID %s is failed, err:%s", diskID, err)
			continue
		}
		if deviceName == "" {
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
	p.pvInfoLock.Lock()
	defer p.pvInfoLock.Unlock()

	for pv, thisInfo := range thisPvDiskInfoMap {
		lastInfo, ok := (*lastPvDiskInfoMap)[pv]
		// add and modify
		if !ok || thisInfo.VolDataPath != lastInfo.VolDataPath {
			pvcNamespace, pvcName, err := getPvcByPvNameByDisk(p.clientSet, pv)
			if err != nil {
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

func getDiskCapacityMetric(info *diskInfo, stat []string) ([]string, error) {
	info.GlobalMountPath = getGlobalMountPathByDiskID(info.DiskID)
	response, err := utils.GetMetrics(info.GlobalMountPath)
	if err != nil {
		klog.Errorf("Get pv %s metrics from kubelet is failed, err: %s", info.GlobalMountPath, err)
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
