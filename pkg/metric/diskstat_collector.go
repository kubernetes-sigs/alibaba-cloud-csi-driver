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
	readsCompletedDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "read_completed_total"),
		"The total number of reads completed successfully.",
		diskStatLabelNames, nil,
	)
	//5 - reads merged
	readsMergeDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "read_merged_total"),
		"The total number of reads merged.",
		diskStatLabelNames,
		nil,
	)
	//6 - sectors read
	readBytesDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "read_bytes_total"),
		"The total number of bytes read successfully.",
		diskStatLabelNames, nil,
	)
	//7 - time spent reading (ms)
	readTimeMilliSecondsDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "read_time_milliseconds_total"),
		"The total number of seconds spent by all reads.",
		diskStatLabelNames,
		nil,
	)
	//8 - writes completed
	writesCompletedDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "write_completed_total"),
		"The total number of writes completed successfully.",
		diskStatLabelNames, nil,
	)
	//9 - writes merged
	writeMergeDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "write_merged_total"),
		"The number of writes merged.",
		diskStatLabelNames,
		nil,
	)
	//10 - sectors written
	writtenBytesDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "write_bytes_total"),
		"The total number of bytes written successfully.",
		diskStatLabelNames, nil,
	)
	//11 - time spent writing (ms)
	writeTimeMilliSecondsDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "write_time_milliseconds_total"),
		"This is the total number of seconds spent by all writes.",
		diskStatLabelNames,
		nil,
	)
	//12 - I/Os currently in progress
	ioNowDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "io_now"),
		"The number of I/Os currently in progress.",
		diskStatLabelNames,
		nil,
	)
	//13 - capacity available
	capacityAvailableDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "capacity_bytes_available"),
		"The number of available size(bytes).",
		diskStatLabelNames,
		nil,
	)

	//14 - capacity total
	capacityTotalDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "capacity_bytes_total"),
		"The number of total size(bytes).",
		diskStatLabelNames,
		nil,
	)

	//15 - capacity used
	capacityUsedDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "capacity_bytes_used"),
		"The number of used size(bytes).",
		diskStatLabelNames,
		nil,
	)
	//16 - inode available
	inodesAvailableDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "inodes_available"),
		"The number of available inodes.",
		diskStatLabelNames,
		nil,
	)
	//17 - inode total
	inodesTotalDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "inodes_total"),
		"The number of total inodes.",
		diskStatLabelNames,
		nil,
	)

	//18 - inode used
	inodesUsedDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "inodes_used"),
		"The number of used inodes.",
		diskStatLabelNames,
		nil,
	)
)

type diskStatCollector struct {
	alertSwtichSet               *hashset.Set
	milliSecondsLatencyThreshold float64 //Unit: milliseconds
	capacityPercentageThreshold  float64
	descs                        []typedFactorDesc
	lastPvStorageInfoMap         map[string]storageInfo
	lastPvStatsMap               sync.Map
	clientSet                    *kubernetes.Clientset
	recorder                     record.EventRecorder
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
	var recorder record.EventRecorder
	if !alertSet.Empty() {
		recorder = utils.NewEventRecorder()
	}
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return &diskStatCollector{
		descs: []typedFactorDesc{
			//4 - reads completed successfully
			{desc: readsCompletedDesc, valueType: prometheus.CounterValue},
			//5 - reads merged
			{desc: readsMergeDesc, valueType: prometheus.CounterValue},
			//6 - sectors read
			{desc: readBytesDesc, valueType: prometheus.CounterValue, factor: diskSectorSize},
			//7 - time spent reading (ms)
			{desc: readTimeMilliSecondsDesc, valueType: prometheus.CounterValue, factor: .001},
			//8 - writes completed
			{desc: writesCompletedDesc, valueType: prometheus.CounterValue},
			//9 - writes merged
			{desc: writeMergeDesc, valueType: prometheus.CounterValue},
			//10 - sectors written
			{desc: writtenBytesDesc, valueType: prometheus.CounterValue, factor: diskSectorSize},
			//11 - time spent writing (ms)
			{desc: writeTimeMilliSecondsDesc, valueType: prometheus.CounterValue, factor: .001},
			//12 - I/Os currently in progress
			{desc: ioNowDesc, valueType: prometheus.GaugeValue},
			//13 - capacity available
			{desc: capacityAvailableDesc, valueType: prometheus.CounterValue},
			//14 - capacity total
			{desc: capacityTotalDesc, valueType: prometheus.CounterValue},
			//15 - capacity used
			{desc: capacityUsedDesc, valueType: prometheus.CounterValue},
			//16 - inode available
			{desc: inodesAvailableDesc, valueType: prometheus.CounterValue},
			//17 - inode total
			{desc: inodesTotalDesc, valueType: prometheus.CounterValue},
			//18 - inode used
			{desc: inodesUsedDesc, valueType: prometheus.CounterValue},
		},
		lastPvStorageInfoMap:         make(map[string]storageInfo, 0),
		lastPvStatsMap:               sync.Map{},
		clientSet:                    clientset,
		milliSecondsLatencyThreshold: latencyThreshold,
		capacityPercentageThreshold:  capacityPercentageThreshold,
		alertSwtichSet:               alertSet,
		recorder:                     recorder,
	}, nil
}

func (p *diskStatCollector) Update(ch chan<- prometheus.Metric) error {
	//startTime := time.Now()
	deviceNameStatsMap, err := getDiskStats()
	if err != nil {
		return fmt.Errorf("couldn't get diskstats: %s", err)
	}
	volJSONPaths, err := findVolJSONByDisk(podsRootPath)
	if err != nil {
		logrus.Errorf("Find disk vol_data json is failed, err:%s", err)
		return err
	}
	updateMap(p.clientSet, &p.lastPvStorageInfoMap, volJSONPaths, diskDriverName, "volumes")

	wg := sync.WaitGroup{}
	for deviceName, stats := range deviceNameStatsMap {
		for pvName, info := range p.lastPvStorageInfoMap {
			if info.DeviceName != deviceName {
				continue
			}
			stats, _ := getCapacityMetric(pvName, &info, stats)

			wg.Add(1)
			go func(nodeNameArgs string, deviceNameArgs string, pvNameArgs string, pvcNamespaceArgs string, pvcNameArgs string, statsArgs []string) {
				defer wg.Done()
				p.setDiskMetric(nodeNameArgs, deviceNameArgs, pvNameArgs, pvcNamespaceArgs, pvcNameArgs, statsArgs, ch)
			}(info.NodeName, deviceName, pvName, info.PvcNamespace, info.PvcName, stats)
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
		logrus.Errorf("Convert incrementIOPS %f to int is failed, err:number is zero", incrementIOPS)
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

func (p *diskStatCollector) latencyEventAlert(nodeName string, pvName string, pvcName string, pvcNamespace string, stats []string, index int) {
	lastStats, ok := p.lastPvStatsMap.Load(pvName)
	if p.alertSwtichSet.Contains(latencySwitch) && ok {
		thisLatency, exceed := isExceedLatencyThreshold(stats, lastStats.([]string), index, index+3, p.milliSecondsLatencyThreshold)
		if exceed {
			ref := &v1.ObjectReference{
				Kind:      "Volume",
				Name:      pvcName,
				UID:       "",
				Namespace: pvcNamespace,
			}
			reason := fmt.Sprintf("Pvc %s latency load is too high, nodeName: %s, namespace: %s, latency:%.2f ms, threshold:%.2f ms", nodeName, pvcName, pvcNamespace, thisLatency, p.milliSecondsLatencyThreshold)
			utils.CreateEvent(p.recorder, ref, v1.EventTypeWarning, latencyTooHigh, reason)
		}
	}
}

func (p *diskStatCollector) ioHangEventAlert(nodeName string, devName string, pvName string, pvcName string, pvcNamespace string, stats []string) {
	lastStats, ok := p.lastPvStatsMap.Load(pvName)
	if ok {
		isHang := isIOHang(stats, lastStats.([]string))
		if isHang {
			ref := &v1.ObjectReference{
				Kind:      "Volume",
				Name:      pvcName,
				UID:       "",
				Namespace: pvcNamespace,
			}
			reason := fmt.Sprintf("IO Hang on Persistent Volume %s, nodeName:%s, diskID:%s, Device:%s", nodeName, pvName, p.lastPvStorageInfoMap[pvName].DiskID, devName)
			utils.CreateEvent(p.recorder, ref, v1.EventTypeWarning, ioHang, reason)
		}
	}
}

func (p *diskStatCollector) capacityEventAlert(nodeName string, valueFloat64 float64, pvcName string, pvcNamespace string, stats []string) {
	if p.alertSwtichSet.Contains(capacitySwitch) {
		capacityTotalFloat64, err := strconv.ParseFloat(stats[10], 64)
		if err != nil {
			logrus.Errorf("Convert diskCapacityTotalDesc %s to float64 is failed, err:%s", stats[10], err)
			return
		}
		if almostEqualFloat64(capacityTotalFloat64, 0) {
			logrus.Errorf("Equal capacityTotalFloat64 %s and zero is true", stats[10])
			return
		}
		usedPercentage := (valueFloat64 / capacityTotalFloat64) * 100
		if usedPercentage >= p.capacityPercentageThreshold {
			ref := &v1.ObjectReference{
				Kind:      "Volume",
				Name:      pvcName,
				UID:       "",
				Namespace: pvcNamespace,
			}
			reason := fmt.Sprintf("Pvc %s is not enough disk space, nodeName:%s, namespace: %s, usedPercentage:%.2f%%, threshold:%.2f%%", nodeName, pvcName, pvcNamespace, usedPercentage, p.capacityPercentageThreshold)
			utils.CreateEvent(p.recorder, ref, v1.EventTypeWarning, capacityNotEnough, reason)
		}
	}
}

func (p *diskStatCollector) setDiskMetric(nodeName string, devName string, pvName string, pvcNamespace string, pvcName string, stats []string, ch chan<- prometheus.Metric) {
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
			p.latencyEventAlert(nodeName, pvName, pvcName, pvcNamespace, stats, 0)
		}

		if i == 7 { //7：diskWriteTimeMilliSecondsDesc
			p.latencyEventAlert(nodeName, pvName, pvcName, pvcNamespace, stats, 4)
		}

		if i == 9 { //9: ioTimeSecondsDesc
			p.ioHangEventAlert(nodeName, devName, pvName, pvcName, pvcNamespace, stats)
		}

		if i == 11 { //11：diskCapacityUsedDesc
			p.capacityEventAlert(nodeName, valueFloat64, pvcName, pvcNamespace, stats)
		}
		ch <- p.descs[i].mustNewConstMetric(valueFloat64, pvcNamespace, pvcName, devName, diskStorageName)
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

func getGlobalMountPathByPvName(pvName string, info *storageInfo) {
	info.GlobalMountPath = fmt.Sprintf("/var/lib/kubelet/plugins/kubernetes.io/csi/pv/%s/globalmount", pvName)
}

func getCapacityMetric(pvName string, info *storageInfo, stat []string) ([]string, error) {
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
		diskStats[dev] = parts[3:12]
	}

	return diskStats, scanner.Err()
}
