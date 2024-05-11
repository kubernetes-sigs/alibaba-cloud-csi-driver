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

	"github.com/emirpasic/gods/sets/hashset"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/options"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/prometheus/client_golang/prometheus"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
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

var (
	//0 - reads completed successfully
	nfsReadsCompletedDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "read_completed_total"),
		"The total number of reads completed successfully.",
		nfsStatLabelNames, nil,
	)
	//1 - reads transmissions successfully
	nfsReadsTransDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "read_transmissions_total"),
		"How many transmissions of this op type have been sent.",
		nfsStatLabelNames, nil,
	)
	//2 - read timeout
	nfsReadsTimeOutDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "read_timeouts_total"),
		"How many timeouts of this op type have occurred.",
		nfsStatLabelNames, nil,
	)
	//3 - read send bytes
	nfsReadsSentBytesDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "read_sent_bytes_total"),
		"How many bytes have been sent for this op type.",
		nfsStatLabelNames, nil,
	)
	//4 - read recv bytes
	nfsReadsRecvBytesDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "read_bytes_total"),
		"The total number of bytes read successfully.",
		nfsStatLabelNames, nil,
	)
	//5 - read queue time
	nfsReadsQueueTimeMilliSecondsDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "read_queue_time_milliseconds_total"),
		"How long ops of this type have waited in queue before being transmitted (microsecond).",
		nfsStatLabelNames, nil,
	)
	//6 - read rtt time
	nfsReadsRttTimeMilliSecondsDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "read_rtt_time_milliseconds_total"),
		"How long the client waited to receive replies of this op type from the server (microsecond).",
		nfsStatLabelNames, nil,
	)
	//7 - read execute time
	nfsReadsExecuteTimeMilliSecondsDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "read_time_milliseconds_total"),
		"The total number of seconds spent by all reads.",
		nfsStatLabelNames, nil,
	)
	//8 - writes completed successfully
	nfsWritesCompletedDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "write_completed_total"),
		"The total number of writes completed successfully.",
		nfsStatLabelNames, nil,
	)
	//9 - writes transmissions successfully
	nfsWritesTransDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "write_transmissions_total"),
		"How many transmissions of this op type have been sent.",
		nfsStatLabelNames, nil,
	)
	//10 - writes timeout
	nfsWritesTimeOutDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "write_timeouts_total"),
		"How many timeouts of this op type have occurred.",
		nfsStatLabelNames, nil,
	)
	//11 - writes send bytes
	nfsWritesSentBytesDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "write_bytes_total"),
		"The total number of bytes written successfully.",
		nfsStatLabelNames, nil,
	)
	//12 - writes recv bytes
	nfsWritesRecvBytesDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "write_recv_bytes_total"),
		"How many bytes have been received for this op type.",
		nfsStatLabelNames, nil,
	)
	//13 - writes queue time
	nfsWritesQueueTimeMilliSecondsDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "write_queue_time_milliseconds_total"),
		"How long ops of this type have waited in queue before being transmitted (microsecond).",
		nfsStatLabelNames, nil,
	)
	//14 - writes rtt time
	nfsWritesRttTimeMilliSecondsDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "write_rtt_time_milliseconds_total"),
		"How long the client waited to receive replies of this op type from the server (microsecond).",
		nfsStatLabelNames, nil,
	)
	//15 - writes execute time
	nfsWritesExecuteTimeMilliSecondsDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "write_time_milliseconds_total"),
		"This is the total number of seconds spent by all writes.",
		nfsStatLabelNames, nil,
	)
	//16 - capacity available
	nfsCapacityAvailableDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "capacity_bytes_available"),
		"The number of available size(bytes).",
		nfsStatLabelNames,
		nil,
	)

	//17 - capacity total
	nfsCapacityTotalDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "capacity_bytes_total"),
		"The number of total size(bytes).",
		nfsStatLabelNames,
		nil,
	)

	//18 - capacity used
	nfsCapacityUsedDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "capacity_bytes_used"),
		"The number of used size(bytes).",
		nfsStatLabelNames,
		nil,
	)
)

type nfsInfo struct {
	PvcNamespace string
	PvcName      string
	ServerName   string
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
	descs                       []typedFactorDesc
	pvInfoLock                  sync.Mutex
	lastPvNfsInfoMap            map[string]nfsInfo
	lastPvStatsMap              sync.Map
	clientSet                   *kubernetes.Clientset
	crdClient                   dynamic.Interface
	monitorClient               *StorageMonitorClient
	recorder                    record.EventRecorder
	alertSwtichSet              *hashset.Set
	capacityPercentageThreshold float64
}

func init() {
	registerCollector("nfs_stat", NewNfsStatCollector, nasDriverName)
}

func parseNfsThreshold(defaultCapacityPercentageThreshold float64) (*hashset.Set, float64) {
	alertSet := hashset.New()
	nfsCapacityThreshold := strings.ToLower(strings.Trim(os.Getenv("NFS_CAPACITY_THRESHOLD_PERCENTAGE"), " "))
	if len(nfsCapacityThreshold) != 0 {
		alertSet.Add(capacitySwitch)
		defaultCapacityPercentageThreshold, _ = parseCapacityThreshold(nfsCapacityThreshold, defaultCapacityPercentageThreshold)
	}

	return alertSet, defaultCapacityPercentageThreshold
}

// NewNfsStatCollector returns a new Collector exposing nfs stats.
func NewNfsStatCollector() (Collector, error) {
	config, err := clientcmd.BuildConfigFromFlags(options.MasterURL, options.Kubeconfig)
	if err != nil {
		return nil, err
	}
	recorder := utils.NewEventRecorder()
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	crdClient, err := dynamic.NewForConfig(config)
	if err != nil {
		klog.Fatalf("Failed to create crd client: %v", err)
	}

	alertSet, capacityPercentageThreshold := parseNfsThreshold(nfsDefaultsCapacityPercentageThreshold)

	return &nfsStatCollector{
		descs: []typedFactorDesc{
			//read
			{desc: nfsReadsCompletedDesc, valueType: prometheus.CounterValue},
			{desc: nfsReadsTransDesc, valueType: prometheus.CounterValue},
			{desc: nfsReadsTimeOutDesc, valueType: prometheus.CounterValue},
			{desc: nfsReadsSentBytesDesc, valueType: prometheus.CounterValue},
			{desc: nfsReadsRecvBytesDesc, valueType: prometheus.CounterValue},
			{desc: nfsReadsQueueTimeMilliSecondsDesc, valueType: prometheus.CounterValue, factor: .001},
			{desc: nfsReadsRttTimeMilliSecondsDesc, valueType: prometheus.CounterValue, factor: .001},
			{desc: nfsReadsExecuteTimeMilliSecondsDesc, valueType: prometheus.CounterValue, factor: .001},
			//write
			{desc: nfsWritesCompletedDesc, valueType: prometheus.CounterValue},
			{desc: nfsWritesTransDesc, valueType: prometheus.CounterValue},
			{desc: nfsWritesTimeOutDesc, valueType: prometheus.CounterValue},
			{desc: nfsWritesSentBytesDesc, valueType: prometheus.CounterValue},
			{desc: nfsWritesRecvBytesDesc, valueType: prometheus.CounterValue},
			{desc: nfsWritesQueueTimeMilliSecondsDesc, valueType: prometheus.CounterValue, factor: .001},
			{desc: nfsWritesRttTimeMilliSecondsDesc, valueType: prometheus.CounterValue, factor: .001},
			{desc: nfsWritesExecuteTimeMilliSecondsDesc, valueType: prometheus.CounterValue, factor: .001},
			// capacity
			{desc: nfsCapacityTotalDesc, valueType: prometheus.CounterValue},
			{desc: nfsCapacityUsedDesc, valueType: prometheus.CounterValue},
			{desc: nfsCapacityAvailableDesc, valueType: prometheus.CounterValue},
		},
		lastPvNfsInfoMap:            make(map[string]nfsInfo, 0),
		lastPvStatsMap:              sync.Map{},
		clientSet:                   clientset,
		crdClient:                   crdClient,
		recorder:                    recorder,
		monitorClient:               NewStorageMonitorClient(clientset),
		alertSwtichSet:              alertSet,
		capacityPercentageThreshold: capacityPercentageThreshold,
	}, nil
}

func (p *nfsStatCollector) Update(ch chan<- prometheus.Metric) error {
	//startTime := time.Now()
	pvNameStatsMap, err := getNfsStat()
	if pvNameStatsMap == nil || len(pvNameStatsMap) == 0 {
		return nil
	}

	if err != nil {
		return fmt.Errorf("couldn't get nfsstats: %s", err)
	}
	volJSONPaths, err := findVolJSON(podsRootPath)
	if err != nil {
		return err
	}
	//log.Infof("volJSONPaths:%+v", volJSONPaths)
	p.updateMap(&p.lastPvNfsInfoMap, volJSONPaths, nasDriverName)
	//log.Infof("lastPvNfsInfoMap:%+v", p.lastPvNfsInfoMap)
	wg := sync.WaitGroup{}
	for pvName, stats := range pvNameStatsMap {
		nfsInfo := p.lastPvNfsInfoMap[pvName]
		//klog.Infof("pv: %s, stats: %v, nfsInfo: %+v", pvName, stats, nfsInfo)
		capacityStats, err := getNfsCapacityStat(pvName, nfsInfo, p)
		if err != nil {
			//klog.Errorf("get capacity of PV %s: %v", pvName, err)
			stats = append(stats, UnknownValue, UnknownValue, UnknownValue)
		} else {
			stats = append(stats, capacityStats...)
		}

		wg.Add(1)
		go func(pvNameArgs string, pvcNamespaceArgs string, pvcNameArgs string, serverNameArgs string, statsArgs []string) {
			defer wg.Done()
			p.setNfsMetric(pvNameArgs, pvcNamespaceArgs, pvcNameArgs, serverNameArgs, statsArgs, ch)
		}(pvName, nfsInfo.PvcNamespace, nfsInfo.PvcName, nfsInfo.ServerName, stats)
	}
	wg.Wait()
	//elapsedTime := time.Since(startTime)
	//logrus.Info("Nfsstat spent time:", elapsedTime)
	return nil
}

func (p *nfsStatCollector) setNfsMetric(pvName string, pvcNamespace string, pvcName string, serverName string, stats []string, ch chan<- prometheus.Metric) {
	defer p.lastPvStatsMap.Store(pvName, stats)
	for i, value := range stats {
		if i >= len(p.descs) {
			return
		}
		if value == UnknownValue {
			continue
		}

		valueFloat64, err := strconv.ParseFloat(value, 64)
		if err != nil {
			klog.Errorf("Convert value %s to float64 is failed, err:%s, stat:%+v", value, err, stats)
			continue
		}

		ch <- p.descs[i].mustNewConstMetric(valueFloat64, pvcNamespace, pvcName, serverName, nasStorageName)
	}
}

func (p *nfsStatCollector) updateMap(lastPvNfsInfoMap *map[string]nfsInfo, jsonPaths []string, deriverName string) {
	thisPvNfsInfoMap := make(map[string]nfsInfo, 0)
	mounter := mount.New("")
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
		notMounted, err := mounter.IsLikelyNotMountPoint(mountPoint)
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
			pvcNamespace, pvcName, serverName, err := getPvcByPvName(p.clientSet, p.crdClient, pv)
			if err != nil {
				continue
			}
			updateInfo := nfsInfo{
				VolDataPath:  thisInfo.VolDataPath,
				PvcName:      pvcName,
				PvcNamespace: pvcNamespace,
				ServerName:   serverName,
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

func getNfsStat() (map[string][]string, error) {
	pvNameStatMapping := make(map[string][]string, 0)
	mountStatsFile, err := os.Open(nfsStatsFileName)
	defer mountStatsFile.Close()
	if mountStatsFile == nil {
		return nil, errors.New(fmt.Sprintf("File %s is not found.", nfsStatsFileName))
	}
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Open file %s is error, err:%s", nfsStatsFileName, err))
	}
	mountArr, err := parseMountStats(mountStatsFile)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("ParseMountStats %s is error, err:%s.", nfsStatsFileName, err))
	}
	for _, mount := range mountArr {
		nfsOperationStats := mount.Stats.operationStats()
		for _, operation := range nfsOperationStats {
			addNfsStat(&pvNameStatMapping, mount.Mount, operation, "READ")
		}
		for _, operation := range nfsOperationStats {
			addNfsStat(&pvNameStatMapping, mount.Mount, operation, "WRITE")
		}
	}
	return pvNameStatMapping, nil
}

func addNfsStat(pvNameStatMapping *map[string][]string, mountPath string, operationStat NFSOperationStats, keyWord string) {
	if operationStat.Operation == keyWord {
		pathArr := strings.Split(mountPath, "/")
		pvName := pathArr[len(pathArr)-2]
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

func getNfsCapacityStat(pvName string, info nfsInfo, p *nfsStatCollector) ([]string, error) {
	capacityInfo, err := p.monitorClient.GetNasCapacityInfo(pvName)
	if err != nil {
		return nil, err
	}
	if capacityInfo != nil && capacityInfo.TotalSize != -1 && capacityInfo.UsedSize != -1 {
		total, used := getTotalAndUsedSize(capacityInfo)
		p.capacityEventAlert(total, used, pvName, info)
		return []string{
			strconv.FormatInt(total, 10),
			strconv.FormatInt(used, 10),
			strconv.FormatInt(total-used, 10),
		}, nil
	}
	// NOTE: The system is unable to extract the capacity statistics because the capacity of the NFS mount point is based
	// on the overall quota of the NAS filesystem, rather than the specific path that is mounted.
	return nil, errors.New("capacity metrics from storage-monitor missing or invalid")
}

func getTotalAndUsedSize(info *nfsCapacityInfo) (int64, int64) {
	if info.TotalSizeInBytes != nil && info.UsedSizeInBytes != nil {
		return *info.TotalSizeInBytes, *info.UsedSizeInBytes
	}
	return info.TotalSize * GiBSize, info.UsedSize * GiBSize
}

func (p *nfsStatCollector) capacityEventAlert(totalSize int64, usedSize int64, pvName string, info nfsInfo) {
	total, used, gibSize := float64(totalSize), float64(usedSize), float64(GiBSize)
	if p.alertSwtichSet.Contains(capacitySwitch) {
		usedPercentage := 100 * used / total
		if usedPercentage >= float64(p.capacityPercentageThreshold) {
			ref := &v1.ObjectReference{
				Kind:      "PersistentVolumeClaim",
				Name:      info.PvcName,
				UID:       "",
				Namespace: info.PvcNamespace,
			}
			reason := fmt.Sprintf("Pvc %s is running out of disk space, namespace: %s, totalSize:%fGi, usedSize:%fGi, usedPercentage:%.2f%%, threshold:%.2f%%",
				info.PvcName, info.PvcNamespace, total/gibSize, used/gibSize, usedPercentage, p.capacityPercentageThreshold)
			utils.CreateEvent(p.recorder, ref, v1.EventTypeWarning, capacityNotEnough, reason)
		}
	}
}
