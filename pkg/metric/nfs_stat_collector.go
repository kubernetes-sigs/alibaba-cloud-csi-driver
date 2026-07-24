//go:build !windows

package metric

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/options"
	"github.com/prometheus/client_golang/prometheus"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"
	"k8s.io/mount-utils"
)

var (
	nfsStatLabelNames = []string{"namespace", "pvc", "server", "type"}
)

const (
	NFSMetricsCount = 16
)

var (
	//0 - reads completed successfully
	nfsReadsCompletedDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "read_completed_total"),
		"The total number of reads completed successfully.",
		nfsStatLabelNames, nil,
	)
	//1 - reads transmissions successfully
	nfsReadsTransDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "read_transmissions_total"),
		"How many transmissions of this op type have been sent.",
		nfsStatLabelNames, nil,
	)
	//2 - read timeout
	nfsReadsTimeOutDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "read_timeouts_total"),
		"How many timeouts of this op type have occurred.",
		nfsStatLabelNames, nil,
	)
	//3 - read send bytes
	nfsReadsSentBytesDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "read_sent_bytes_total"),
		"How many bytes have been sent for this op type.",
		nfsStatLabelNames, nil,
	)
	//4 - read recv bytes
	nfsReadsRecvBytesDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "read_bytes_total"),
		"The total number of bytes read successfully.",
		nfsStatLabelNames, nil,
	)
	//5 - read queue time
	nfsReadsQueueTimeMilliSecondsDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "read_queue_time_milliseconds_total"),
		"How long ops of this type have waited in queue before being transmitted (microsecond).",
		nfsStatLabelNames, nil,
	)
	//6 - read rtt time
	nfsReadsRttTimeMilliSecondsDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "read_rtt_time_milliseconds_total"),
		"How long the client waited to receive replies of this op type from the server (microsecond).",
		nfsStatLabelNames, nil,
	)
	//7 - read execute time
	nfsReadsExecuteTimeMilliSecondsDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "read_time_milliseconds_total"),
		"The total number of milliseconds spent by all reads.",
		nfsStatLabelNames, nil,
	)
	//8 - writes completed successfully
	nfsWritesCompletedDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "write_completed_total"),
		"The total number of writes completed successfully.",
		nfsStatLabelNames, nil,
	)
	//9 - writes transmissions successfully
	nfsWritesTransDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "write_transmissions_total"),
		"How many transmissions of this op type have been sent.",
		nfsStatLabelNames, nil,
	)
	//10 - writes timeout
	nfsWritesTimeOutDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "write_timeouts_total"),
		"How many timeouts of this op type have occurred.",
		nfsStatLabelNames, nil,
	)
	//11 - writes send bytes
	nfsWritesSentBytesDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "write_bytes_total"),
		"The total number of bytes written successfully.",
		nfsStatLabelNames, nil,
	)
	//12 - writes recv bytes
	nfsWritesRecvBytesDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "write_recv_bytes_total"),
		"How many bytes have been received for this op type.",
		nfsStatLabelNames, nil,
	)
	//13 - writes queue time
	nfsWritesQueueTimeMilliSecondsDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "write_queue_time_milliseconds_total"),
		"How long ops of this type have waited in queue before being transmitted (microsecond).",
		nfsStatLabelNames, nil,
	)
	//14 - writes rtt time
	nfsWritesRttTimeMilliSecondsDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "write_rtt_time_milliseconds_total"),
		"How long the client waited to receive replies of this op type from the server (microsecond).",
		nfsStatLabelNames, nil,
	)
	//15 - writes execute time
	nfsWritesExecuteTimeMilliSecondsDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "write_time_milliseconds_total"),
		"The total number of milliseconds spent by all writes.",
		nfsStatLabelNames, nil,
	)
)

type nfsInfo struct {
	PvcNamespace string
	PvcName      string
	ServerName   string
	VolDataPath  string
}

type nfsStatCollector struct {
	descs            []typedFactorDesc
	pvInfoLock       sync.Mutex
	lastPvNfsInfoMap map[string]nfsInfo
	lastPvStatsMap   sync.Map
	clientSet        *kubernetes.Clientset
	crdClient        dynamic.Interface
	mounter          mount.Interface
}

func init() {
	registerCollector("nfs_stat", NewNfsStatCollector, nasDriverName)
}

func warnIfNfsCapacityThresholdSet() {
	if v := strings.TrimSpace(os.Getenv("NFS_CAPACITY_THRESHOLD_PERCENTAGE")); v != "" {
		klog.Warningf("NFS_CAPACITY_THRESHOLD_PERCENTAGE=%s is set but no longer supported: NAS capacity metrics and the associated alerts have been removed, so this setting has no effect.", v)
	}
}

// NewNfsStatCollector returns a new Collector exposing nfs stats.
func NewNfsStatCollector() (Collector, error) {
	warnIfNfsCapacityThresholdSet()
	config, err := options.GetRestConfig()
	if err != nil {
		return nil, err
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	crdCfg := options.GetRestConfigForCRD(*config)
	crdClient, err := dynamic.NewForConfig(crdCfg)
	if err != nil {
		klog.Fatalf("Failed to create crd client: %v", err)
	}

	return &nfsStatCollector{
		descs: []typedFactorDesc{
			//read
			{desc: nfsReadsCompletedDesc, valueType: prometheus.CounterValue},
			{desc: nfsReadsTransDesc, valueType: prometheus.CounterValue},
			{desc: nfsReadsTimeOutDesc, valueType: prometheus.CounterValue},
			{desc: nfsReadsSentBytesDesc, valueType: prometheus.CounterValue},
			{desc: nfsReadsRecvBytesDesc, valueType: prometheus.CounterValue},
			{desc: nfsReadsQueueTimeMilliSecondsDesc, valueType: prometheus.CounterValue},
			{desc: nfsReadsRttTimeMilliSecondsDesc, valueType: prometheus.CounterValue},
			{desc: nfsReadsExecuteTimeMilliSecondsDesc, valueType: prometheus.CounterValue},
			//write
			{desc: nfsWritesCompletedDesc, valueType: prometheus.CounterValue},
			{desc: nfsWritesTransDesc, valueType: prometheus.CounterValue},
			{desc: nfsWritesTimeOutDesc, valueType: prometheus.CounterValue},
			{desc: nfsWritesSentBytesDesc, valueType: prometheus.CounterValue},
			{desc: nfsWritesRecvBytesDesc, valueType: prometheus.CounterValue},
			{desc: nfsWritesQueueTimeMilliSecondsDesc, valueType: prometheus.CounterValue},
			{desc: nfsWritesRttTimeMilliSecondsDesc, valueType: prometheus.CounterValue},
			{desc: nfsWritesExecuteTimeMilliSecondsDesc, valueType: prometheus.CounterValue},
		},
		lastPvNfsInfoMap: make(map[string]nfsInfo, 0),
		lastPvStatsMap:   sync.Map{},
		clientSet:        clientset,
		crdClient:        crdClient,
		mounter:          mount.NewWithoutSystemd(""),
	}, nil
}

func (p *nfsStatCollector) Update(ctx context.Context, pvcs sets.Set[string], ch chan<- prometheus.Metric) error {
	//startTime := time.Now()
	pvNameStatsMap, err := getNfsStat()
	if len(pvNameStatsMap) == 0 {
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
	p.updateMap(ctx, &p.lastPvNfsInfoMap, volJSONPaths, nasDriverName)
	//log.Infof("lastPvNfsInfoMap:%+v", p.lastPvNfsInfoMap)
	wg := sync.WaitGroup{}
	for pvName, stats := range pvNameStatsMap {
		if err := ctx.Err(); err != nil {
			break
		}
		nfsInfo := p.lastPvNfsInfoMap[pvName]

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
		valueFloat64, err := strconv.ParseFloat(value, 64)
		if err != nil {
			klog.Errorf("Convert value %s to float64 is failed, err:%s, stat:%+v", value, err, stats)
			continue
		}

		ch <- p.descs[i].mustNewConstMetric(valueFloat64, pvcNamespace, pvcName, serverName, nasStorageName)
	}
}

func (p *nfsStatCollector) updateMap(ctx context.Context, lastPvNfsInfoMap *map[string]nfsInfo, jsonPaths []string, deriverName string) {
	thisPvNfsInfoMap := make(map[string]nfsInfo, 0)
	for _, path := range jsonPaths {
		if ctx.Err() != nil {
			break
		}
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
	p.updateNfsInfoMap(ctx, thisPvNfsInfoMap, lastPvNfsInfoMap)
}

func (p *nfsStatCollector) updateNfsInfoMap(ctx context.Context, thisPvNfsInfoMap map[string]nfsInfo, lastPvNfsInfoMap *map[string]nfsInfo) {
	p.pvInfoLock.Lock()
	defer p.pvInfoLock.Unlock()

	for pv, thisInfo := range thisPvNfsInfoMap {
		lastInfo, ok := (*lastPvNfsInfoMap)[pv]
		// add and modify
		if !ok || thisInfo.VolDataPath != lastInfo.VolDataPath {
			pvcNamespace, pvcName, serverName, err := getNasPvcByPvName(ctx, p.clientSet, p.crdClient, pv)
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
	if mountStatsFile == nil {
		return nil, fmt.Errorf("File %s is not found.", nfsStatsFileName)
	}
	if err != nil {
		return nil, fmt.Errorf("Open file %s is error, err:%s", nfsStatsFileName, err)
	}
	defer mountStatsFile.Close()
	mountArr, err := parseMountStats(mountStatsFile)
	if err != nil {
		return nil, fmt.Errorf("ParseMountStats %s is error, err:%s.", nfsStatsFileName, err)
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
