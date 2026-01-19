package metric

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/prometheus/client_golang/prometheus"
	promdto "github.com/prometheus/client_model/go"
	"github.com/prometheus/procfs/blockdevice"
	"golang.org/x/sys/unix"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/record"
	"k8s.io/klog/v2"
	"k8s.io/mount-utils"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/options"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
)

var (
	diskStatLabelNames  = []string{"namespace", "pvc", "device"}
	diskStatConstLabels = prometheus.Labels{"type": diskStorageName}
	scalerPvcMap        *sync.Map
)

func diskMetricDesc(name, help string) *MetaDesc {
	return NewMetaDesc(nodeNamespace, volumeSubsystem, name, help, diskStatLabelNames, diskStatConstLabels)
}

func convertStringSliceToLabelPairs(labels []string) []*promdto.LabelPair {
	result := make([]*promdto.LabelPair, len(labels))
	for i, label := range labels {
		result[i] = &promdto.LabelPair{Name: &diskStatLabelNames[i], Value: &label}
	}
	return result
}

func convertLabelsToString(labelPairs []*promdto.LabelPair) []string {
	result := make([]string, len(labelPairs))
	for i, pair := range labelPairs {
		result[i] = *pair.Value
	}
	return result
}

// stats from /proc/diskstats
var (
	diskReadsCompletedDesc = diskMetricDesc("read_completed_total", "The total number of reads completed successfully.")

	diskReadsMergeDesc           = diskMetricDesc("read_merged_total", "The total number of reads merged.")
	diskReadBytesDesc            = diskMetricDesc("read_bytes_total", "The total number of bytes read successfully.")
	diskReadTimeMilliSecondsDesc = diskMetricDesc("read_time_milliseconds_total", "The total number of milliseconds spent by all reads.")

	diskWritesCompletedDesc       = diskMetricDesc("write_completed_total", "The total number of writes completed successfully.")
	diskWriteMergeDesc            = diskMetricDesc("write_merged_total", "The number of writes merged.")
	diskWrittenBytesDesc          = diskMetricDesc("write_bytes_total", "The total number of bytes written successfully.")
	diskWriteTimeMilliSecondsDesc = diskMetricDesc("write_time_milliseconds_total", "The total number of milliseconds spent by all writes.")

	diskIONowDesc         = diskMetricDesc("io_now", "The number of I/Os currently in progress.")
	diskIOTimeSecondsDesc = diskMetricDesc("io_time_seconds_total", "Total seconds spent doing I/Os.")
)

type capDescs struct {
	Available *MetaDesc
	Total     *MetaDesc
	Used      *MetaDesc
}

var capStatDescs = map[csi.VolumeUsage_Unit]capDescs{
	csi.VolumeUsage_BYTES: {
		Available: diskMetricDesc("capacity_bytes_available", "The number of available size(bytes)."),
		Total:     diskMetricDesc("capacity_bytes_total", "The number of total size(bytes)."),
		Used:      diskMetricDesc("capacity_bytes_used", "The number of used size(bytes)."),
	},
	csi.VolumeUsage_INODES: {
		Available: diskMetricDesc("inodes_available", "The number of available inodes."),
		Total:     diskMetricDesc("inodes_total", "The number of total inodes."),
		Used:      diskMetricDesc("inodes_used", "The number of used inodes."),
	},
}

type diskInfo struct {
	PVCRef      *v1.ObjectReference
	DiskID      string
	Dev         uint64 // as in syscall.Stat_t.Dev
	VolDataPath string
}

type diskStatCollector struct {
	milliSecondsLatencyThreshold float64 //Unit: milliseconds
	capacityPercentageThreshold  float64
	pvInfoLock                   sync.Mutex
	lastPvDiskInfoMap            map[string]diskInfo
	lastPvStatsMap               atomic.Pointer[map[uint64]*blockdevice.Diskstats]
	diskStats                    *ProcDiskStats
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
		latency, _ := parseLatencyThreshold(latencyStr, diskDefaultsLatencyThreshold)
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
	config, err := options.GetRestConfig()
	if err != nil {
		return nil, err
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	diskStats, err := NewDefaultProcDiskStats()
	if err != nil {
		return nil, fmt.Errorf("couldn't get diskstats: %s", err)
	}
	diskStats.HungDuration = 30 * time.Second

	nodeName := os.Getenv("KUBE_NODE_NAME")
	return &diskStatCollector{
		lastPvDiskInfoMap:            make(map[string]diskInfo, 0),
		diskStats:                    diskStats,
		clientSet:                    clientset,
		milliSecondsLatencyThreshold: getDiskLatencyThreshold(),
		capacityPercentageThreshold:  getDiskCapacityThreshold(),
		recorder:                     recorder,
		mounter:                      mount.NewWithoutSystemd(""),
		nodeName:                     nodeName,
	}, nil
}

func (p *diskStatCollector) Get() (metrics []*Metric) {
	volJSONPaths, err := findVolJSON(podsRootPath)
	if err != nil {
		return
	}
	p.updateMap(&p.lastPvDiskInfoMap, volJSONPaths, diskDriverName)

	diskStats, err := p.diskStats.GetStats()
	if err != nil {
		klog.Errorf("couldn't get diskstats: %s", err)
		return
	}
	deviceNameStatsMap := make(map[uint64]*blockdevice.Diskstats, len(diskStats))
	for i, s := range diskStats {
		dev := unix.Mkdev(s.MajorNumber, s.MinorNumber)
		deviceNameStatsMap[dev] = &diskStats[i]
	}
	defer p.lastPvStatsMap.Store(&deviceNameStatsMap)

	hungDevs := sets.New[uint64]()
	for _, info := range p.diskStats.GetHungDisks() {
		hungDevs.Insert(unix.Mkdev(info.MajorNumber, info.MinorNumber))
	}

	var lastStatsMap map[uint64]*blockdevice.Diskstats
	if m := p.lastPvStatsMap.Load(); m != nil {
		lastStatsMap = *m
	}

	for pvName, info := range p.lastPvDiskInfoMap {
		if scalerPvcMap != nil {
			if _, ok := scalerPvcMap.Load(info.PVCRef.Name); !ok {
				continue
			}
		}
		stats, ok := deviceNameStatsMap[info.Dev]
		if !ok {
			continue
		}
		devPath := "/dev/" + stats.DeviceName
		labelValues := []string{info.PVCRef.Namespace, info.PVCRef.Name, devPath}
		metrics = append(metrics, p.getDiskStats(&stats.IOStats, labelValues...)...)
		if lastStats, ok := lastStatsMap[info.Dev]; ok {
			p.latencyEventAlert(&stats.IOStats, &lastStats.IOStats, info.PVCRef)
		}
		if hungDevs.Has(info.Dev) {
			p.recorder.Eventf(info.PVCRef, v1.EventTypeWarning, ioHang, "IO Hang on Persistent Volume %s, nodeName:%s, diskID:%s, Device:%s",
				pvName, p.nodeName, info.DiskID, devPath)
		}

		capStats, err := getDiskCapacityMetric(info.DiskID)
		if err != nil {
			klog.ErrorS(err, "Get disk capacity failed", "disk", info.DiskID, err)
			continue
		}
		metrics = append(metrics, p.getCapStats(capStats, labelValues)...)
		p.capacityEventAlert(capStats, info.PVCRef)
	}

	return metrics
}

func (p *diskStatCollector) Update(ch chan<- prometheus.Metric) error {
	metrics := p.Get()
	for _, metric := range metrics {
		ch <- prometheus.MustNewConstMetric(metric.Desc, metric.ValueType, metric.Value, convertLabelsToString(metric.VariableLabelPairs)...)
	}
	return nil
}

func (p *diskStatCollector) latencyEventAlert(stat, lastStat *blockdevice.IOStats, ref *v1.ObjectReference) {
	if p.milliSecondsLatencyThreshold <= 0 {
		return
	}

	alert := func(op string, dLatency, dIOPS uint64) {
		if dIOPS == 0 {
			return
		}
		l := float64(dLatency) / float64(dIOPS)
		if l <= p.milliSecondsLatencyThreshold {
			return
		}
		p.recorder.Eventf(ref, v1.EventTypeWarning, latencyTooHigh, "PVC %s/%s latency is too high, nodeName: %s, latency:%.2f ms, threshold:%.2f ms",
			ref.Namespace, ref.Name, op, p.nodeName, l, p.milliSecondsLatencyThreshold)
	}
	// Note that this `lastStat.ReadTicks-stat.ReadTicks` even works when the counter overflows and wraps around.
	alert("read", lastStat.ReadTicks-stat.ReadTicks, lastStat.ReadIOs-stat.ReadIOs)
	alert("write", lastStat.WriteTicks-stat.WriteTicks, lastStat.WriteIOs-stat.WriteIOs)
}

func (p *diskStatCollector) capacityEventAlert(usage []*csi.VolumeUsage, ref *v1.ObjectReference) {
	if p.capacityPercentageThreshold <= 0 {
		return
	}
	for _, stat := range usage {
		usedPercentage := (float64(stat.Used) / float64(stat.Total)) * 100
		if usedPercentage >= p.capacityPercentageThreshold {
			p.recorder.Eventf(ref, v1.EventTypeWarning, capacityNotEnough,
				"PVC %s/%s has not enough disk %v capacity, nodeName:%s, used:%.2f%%, threshold:%.2f%%",
				ref.Namespace, ref.Name, stat.Unit, p.nodeName, usedPercentage, p.capacityPercentageThreshold)
		}
	}
}

func (p *diskStatCollector) getDiskStats(stats *blockdevice.IOStats, labelValues ...string) []*Metric {
	return []*Metric{
		MustNewMetricWithMetaDesc(diskReadsCompletedDesc, float64(stats.ReadIOs), prometheus.CounterValue, labelValues...),
		MustNewMetricWithMetaDesc(diskReadsMergeDesc, float64(stats.ReadMerges), prometheus.CounterValue, labelValues...),
		MustNewMetricWithMetaDesc(diskReadBytesDesc, float64(stats.ReadSectors)*diskSectorSize, prometheus.CounterValue, labelValues...),
		MustNewMetricWithMetaDesc(diskReadTimeMilliSecondsDesc, float64(stats.ReadTicks), prometheus.CounterValue, labelValues...),
		MustNewMetricWithMetaDesc(diskWritesCompletedDesc, float64(stats.WriteIOs), prometheus.CounterValue, labelValues...),
		MustNewMetricWithMetaDesc(diskWriteMergeDesc, float64(stats.WriteMerges), prometheus.CounterValue, labelValues...),
		MustNewMetricWithMetaDesc(diskWrittenBytesDesc, float64(stats.WriteSectors)*diskSectorSize, prometheus.CounterValue, labelValues...),
		MustNewMetricWithMetaDesc(diskWriteTimeMilliSecondsDesc, float64(stats.WriteTicks), prometheus.CounterValue, labelValues...),
		MustNewMetricWithMetaDesc(diskIONowDesc, float64(stats.IOsInProgress), prometheus.GaugeValue, labelValues...),
		MustNewMetricWithMetaDesc(diskIOTimeSecondsDesc, float64(stats.IOsTotalTicks)/1000, prometheus.CounterValue, labelValues...),
	}
}

func (p *diskStatCollector) getCapStats(stats []*csi.VolumeUsage, labelValues []string) (metrics []*Metric) {
	for _, stat := range stats {
		descs := capStatDescs[stat.Unit]
		metrics = append(metrics, MustNewMetricWithMetaDesc(descs.Available, float64(stat.Available), prometheus.GaugeValue, labelValues...))
		metrics = append(metrics, MustNewMetricWithMetaDesc(descs.Total, float64(stat.Total), prometheus.GaugeValue, labelValues...))
		metrics = append(metrics, MustNewMetricWithMetaDesc(descs.Used, float64(stat.Used), prometheus.GaugeValue, labelValues...))
	}
	return
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

		var stat unix.Stat_t
		err = unix.Stat(mountPoint, &stat)
		if err != nil {
			klog.Errorf("Get stat of %s failed: %v", mountPoint, err)
			continue
		}
		diskInfo := diskInfo{
			DiskID:      diskID,
			Dev:         uint64(stat.Dev),
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
			pvcRef, err := getDiskPvcByPvName(p.clientSet, pv)
			if err != nil {
				continue
			}
			updateInfo := diskInfo{
				DiskID:      thisInfo.DiskID,
				VolDataPath: thisInfo.VolDataPath,
				Dev:         thisInfo.Dev,
				PVCRef:      pvcRef,
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

func getDiskCapacityMetric(diskID string) ([]*csi.VolumeUsage, error) {
	globalMountPath := getGlobalMountPathByDiskID(diskID)
	response, err := utils.GetMetrics(globalMountPath)
	if err != nil {
		return nil, err
	}
	return response.Usage, nil
}
