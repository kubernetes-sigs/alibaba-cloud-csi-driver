package metric

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/prometheus/client_golang/prometheus"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/record"
	"k8s.io/klog/v2"
	"k8s.io/mount-utils"
)

var (
	nfsStatLabelNames = []string{"namespace", "pvc", "server", "type"}
)

const (
	NFSMetricsCount = 16
	GiBSize         = 1024 * 1024 * 1024
)

func nfsMetricDesc(name, help string) *MetaDesc {
	return NewMetaDesc(nodeNamespace, volumeSubsystem, name, help, nfsStatLabelNames, nil)
}

var (
	nfsReadsCompletedDesc                = nfsMetricDesc("read_completed_total", "The total number of reads completed successfully.")
	nfsReadsTransDesc                    = nfsMetricDesc("read_transmissions_total", "How many transmissions of this op type have been sent.")
	nfsReadsTimeoutDesc                  = nfsMetricDesc("read_timeouts_total", "How many timeouts of this op type have occurred.")
	nfsReadsSentBytesDesc                = nfsMetricDesc("read_sent_bytes_total", "How many bytes have been sent for this op type.")
	nfsReadsRecvBytesDesc                = nfsMetricDesc("read_bytes_total", "The total number of bytes read successfully.")
	nfsReadsQueueTimeMilliSecondsDesc    = nfsMetricDesc("read_queue_time_milliseconds_total", "How long ops of this type have waited in queue before being transmitted (microsecond).")
	nfsReadsRttTimeMilliSecondsDesc      = nfsMetricDesc("read_rtt_time_milliseconds_total", "How long the client waited to receive replies of this op type from the server (microsecond).")
	nfsReadsExecuteTimeMilliSecondsDesc  = nfsMetricDesc("read_time_milliseconds_total", "The total number of milliseconds spent by all reads.")
	nfsWritesCompletedDesc               = nfsMetricDesc("write_completed_total", "The total number of writes completed successfully.")
	nfsWritesTransDesc                   = nfsMetricDesc("write_transmissions_total", "How many transmissions of this op type have been sent.")
	nfsWritesTimeoutDesc                 = nfsMetricDesc("write_timeouts_total", "How many timeouts of this op type have occurred.")
	nfsWritesSentBytesDesc               = nfsMetricDesc("write_bytes_total", "The total number of bytes written successfully.")
	nfsWritesRecvBytesDesc               = nfsMetricDesc("write_recv_bytes_total", "How many bytes have been received for this op type.")
	nfsWritesQueueTimeMilliSecondsDesc   = nfsMetricDesc("write_queue_time_milliseconds_total", "How long ops of this type have waited in queue before being transmitted (microsecond).")
	nfsWritesRttTimeMilliSecondsDesc     = nfsMetricDesc("write_rtt_time_milliseconds_total", "How long the client waited to receive replies of this op type from the server (microsecond).")
	nfsWritesExecuteTimeMilliSecondsDesc = nfsMetricDesc("write_time_milliseconds_total", "The total number of milliseconds spent by all writes.")
)

type nfsInfo struct {
	PvcNamespace string
	PvcName      string
	VolDataPath  string
}

type nfsCapacityInfo struct {
	TotalInodeCount  int64  `json:"totalInodeCount"`
	UsedInodeCount   int64  `json:"usedInodeCount"`
	TotalSize        int64  `json:"totalSize"`
	TotalSizeInBytes *int64 `json:"totalSizeInBytes"`
	UsedSize         int64  `json:"usedSize"`
	UsedSizeInBytes  *int64 `json:"usedSizeInBytes"`
}

type nfsStatCollector struct {
	descs            []typedFactorDesc
	pvInfoLock       sync.Mutex
	lastPvNfsInfoMap map[string]nfsInfo
	lastPvStatsMap   sync.Map
	client           kubernetes.Interface
	recorder         record.EventRecorder
	mounter          mount.Interface
}

func init() {
	registerCollector("nfs_stat", NewNfsStatCollector, nasDriverName)
}

// NewNfsStatCollector returns a new Collector exposing nfs stats.
func NewNfsStatCollector() (Collector, error) {
	recorder := utils.NewEventRecorder()
	client, err := newK8sClient()
	if err != nil {
		return nil, err
	}

	return &nfsStatCollector{
		descs: []typedFactorDesc{
			//read
			{MetaDesc: nfsReadsCompletedDesc, valueType: prometheus.CounterValue},
			{MetaDesc: nfsReadsTransDesc, valueType: prometheus.CounterValue},
			{MetaDesc: nfsReadsTimeoutDesc, valueType: prometheus.CounterValue},
			{MetaDesc: nfsReadsSentBytesDesc, valueType: prometheus.CounterValue},
			{MetaDesc: nfsReadsRecvBytesDesc, valueType: prometheus.CounterValue},
			{MetaDesc: nfsReadsQueueTimeMilliSecondsDesc, valueType: prometheus.CounterValue},
			{MetaDesc: nfsReadsRttTimeMilliSecondsDesc, valueType: prometheus.CounterValue},
			{MetaDesc: nfsReadsExecuteTimeMilliSecondsDesc, valueType: prometheus.CounterValue},
			//write
			{MetaDesc: nfsWritesCompletedDesc, valueType: prometheus.CounterValue},
			{MetaDesc: nfsWritesTransDesc, valueType: prometheus.CounterValue},
			{MetaDesc: nfsWritesTimeoutDesc, valueType: prometheus.CounterValue},
			{MetaDesc: nfsWritesSentBytesDesc, valueType: prometheus.CounterValue},
			{MetaDesc: nfsWritesRecvBytesDesc, valueType: prometheus.CounterValue},
			{MetaDesc: nfsWritesQueueTimeMilliSecondsDesc, valueType: prometheus.CounterValue},
			{MetaDesc: nfsWritesRttTimeMilliSecondsDesc, valueType: prometheus.CounterValue},
			{MetaDesc: nfsWritesExecuteTimeMilliSecondsDesc, valueType: prometheus.CounterValue},
		},
		lastPvNfsInfoMap: make(map[string]nfsInfo, 0),
		lastPvStatsMap:   sync.Map{},
		client:           client,
		recorder:         recorder,
		mounter:          mount.NewWithoutSystemd(""),
	}, nil
}

func (p *nfsStatCollector) Update(ch chan<- prometheus.Metric) error {
	metrics := p.Get()
	for _, metric := range metrics {
		ch <- prometheus.MustNewConstMetric(metric.Desc, metric.ValueType, metric.Value, convertLabelsToString(metric.VariableLabelPairs)...)
	}
	return nil
}

func (p *nfsStatCollector) Get() (metrics []*Metric) {
	pvNameServerMap, pvNameStatsMap, err := getNfsStat()
	if len(pvNameStatsMap) == 0 {
		klog.V(2).InfoS("No nfs stats found")
		return
	}

	if err != nil {
		klog.ErrorS(err, "couldn't get nfsstats")
		return
	}
	volJSONPaths, err := findVolJSON(podsRootPath)
	if err != nil {
		klog.ErrorS(err, "couldn't find vol json", "podRootPath", podsRootPath)
		return
	}
	p.updateMap(&p.lastPvNfsInfoMap, volJSONPaths, nasDriverName)
	for pvName, stats := range pvNameStatsMap {
		nfsInfo := p.lastPvNfsInfoMap[pvName]
		metrics = append(metrics, p.getNfsMetrics(pvName, nfsInfo.PvcNamespace, nfsInfo.PvcName, pvNameServerMap[pvName], stats)...)
	}
	return
}

func (p *nfsStatCollector) getNfsMetrics(pvName string, pvcNamespace string, pvcName string, serverName string, stats []string) (metrics []*Metric) {
	defer p.lastPvStatsMap.Store(pvName, stats)
	for i, value := range stats {
		if i >= len(p.descs) {
			return
		}

		valueFloat64, err := strconv.ParseFloat(value, 64)
		if err != nil {
			klog.Errorf("Convert value %s to float64 is failed, err:%s, stat:%+v", value, err, stats)
			continue
		}

		metrics = append(metrics, MustNewMetricWithTypedFactorDesc(p.descs[i], valueFloat64, pvcNamespace, pvcName, serverName, nasStorageName))
	}
	return
}

func (p *nfsStatCollector) updateMap(lastPvNfsInfoMap *map[string]nfsInfo, jsonPaths []string, deriverName string) {
	thisPvNfsInfoMap := make(map[string]nfsInfo, 0)
	for _, path := range jsonPaths {
		//Get nfs pvName
		pvName, _, err := getVolumeInfoByJSON(path, deriverName)
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

		nfsInfo := nfsInfo{
			VolDataPath: path,
		}
		thisPvNfsInfoMap[pvName] = nfsInfo
	}

	//If there is a change: add, modify, delete
	p.updateNfsInfoMap(thisPvNfsInfoMap, lastPvNfsInfoMap)
}

func (p *nfsStatCollector) updateNfsInfoMap(thisPvNfsInfoMap map[string]nfsInfo, lastPvNfsInfoMap *map[string]nfsInfo) {
	p.pvInfoLock.Lock()
	defer p.pvInfoLock.Unlock()

	for pv, thisInfo := range thisPvNfsInfoMap {
		lastInfo, ok := (*lastPvNfsInfoMap)[pv]
		// add and modify
		if !ok || thisInfo.VolDataPath != lastInfo.VolDataPath {
			pvcNamespace, pvcName, err := getNasPvcByPvName(p.client, pv, thisInfo.VolDataPath)
			if err != nil {
				continue
			}
			updateInfo := nfsInfo{
				VolDataPath:  thisInfo.VolDataPath,
				PvcName:      pvcName,
				PvcNamespace: pvcNamespace,
			}
			(*lastPvNfsInfoMap)[pv] = updateInfo
		}
	}
	//if pv exist thisPvStorageInfoMap and not exist lastPvNfsInfoMap, pv should be deleted
	for lastPv := range *lastPvNfsInfoMap {
		_, ok := thisPvNfsInfoMap[lastPv]
		if !ok {
			delete(*lastPvNfsInfoMap, lastPv)
		}
	}
}

func getNfsStat() (map[string]string, map[string][]string, error) {
	pvServerMapping := make(map[string]string, 0)
	pvNameStatMapping := make(map[string][]string, 0)
	mountStatsFile, err := os.Open(nfsStatsFileName)
	if mountStatsFile == nil {
		return nil, nil, fmt.Errorf("File %s is not found.", nfsStatsFileName)
	}
	if err != nil {
		return nil, nil, fmt.Errorf("Open file %s is error, err:%s", nfsStatsFileName, err)
	}
	defer mountStatsFile.Close()
	mountArr, err := parseMountStats(mountStatsFile)
	if err != nil {
		return nil, nil, fmt.Errorf("ParseMountStats %s is error, err:%s.", nfsStatsFileName, err)
	}
	for _, mount := range mountArr {
		pvName := getPVName(mount)
		segments := strings.Split(mount.Device, ":")
		if len(segments) >= 2 {
			pvServerMapping[pvName] = segments[0]
		} else {
			pvServerMapping[pvName] = mount.Device
		}
		nfsOperationStats := mount.Stats.operationStats()
		for _, operation := range nfsOperationStats {
			addNfsStat(&pvNameStatMapping, pvName, operation, "READ")
		}
		for _, operation := range nfsOperationStats {
			addNfsStat(&pvNameStatMapping, pvName, operation, "WRITE")
		}
	}
	return pvServerMapping, pvNameStatMapping, nil
}

func getPVName(mount *Mount) string {
	if mount == nil {
		return ""
	}
	paths := strings.Split(mount.Mount, "/")
	if len(paths) < 2 {
		return ""
	}
	return paths[len(paths)-2]
}

func addNfsStat(pvNameStatMapping *map[string][]string, pvName string, operationStat NFSOperationStats, keyWord string) {
	if operationStat.Operation == keyWord {
		if len((*pvNameStatMapping)[pvName]) >= NFSMetricsCount {
			return
		}
		(*pvNameStatMapping)[pvName] = append((*pvNameStatMapping)[pvName], strconv.Itoa(int(operationStat.Requests)))
		(*pvNameStatMapping)[pvName] = append((*pvNameStatMapping)[pvName], strconv.Itoa(int(operationStat.Transmissions)))
		(*pvNameStatMapping)[pvName] = append((*pvNameStatMapping)[pvName], strconv.Itoa(int(operationStat.MajorTimeouts)))
		(*pvNameStatMapping)[pvName] = append((*pvNameStatMapping)[pvName], strconv.Itoa(int(operationStat.BytesSent)))
		(*pvNameStatMapping)[pvName] = append((*pvNameStatMapping)[pvName], strconv.Itoa(int(operationStat.BytesReceived)))
		(*pvNameStatMapping)[pvName] = append((*pvNameStatMapping)[pvName], strconv.Itoa(int(operationStat.CumulativeQueueMilliseconds)))
		(*pvNameStatMapping)[pvName] = append((*pvNameStatMapping)[pvName], strconv.Itoa(int(operationStat.CumulativeTotalResponseMilliseconds)))
		(*pvNameStatMapping)[pvName] = append((*pvNameStatMapping)[pvName], strconv.Itoa(int(operationStat.CumulativeTotalRequestMilliseconds)))
	}
}
