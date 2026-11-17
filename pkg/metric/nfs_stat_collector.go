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

func nfsMetricDesc(name, help string) *MetaDesc {
	return NewMetaDesc(nodeNamespace, volumeSubsystem, name, help, nfsStatLabelNames, nil)
}

var (
	nfsReadsCompletedDesc                = nfsMetricDesc("read_completed_total", "The total number of reads completed successfully.")
	nfsReadsTransDesc                    = nfsMetricDesc("read_transmissions_total", "How many transmissions of this op type have been sent.")
	nfsReadsTimeOutDesc                  = nfsMetricDesc("read_timeouts_total", "How many timeouts of this op type have occurred.")
	nfsReadsSentBytesDesc                = nfsMetricDesc("read_sent_bytes_total", "How many bytes have been sent for this op type.")
	nfsReadsRecvBytesDesc                = nfsMetricDesc("read_bytes_total", "The total number of bytes read successfully.")
	nfsReadsQueueTimeMilliSecondsDesc    = nfsMetricDesc("read_queue_time_milliseconds_total", "How long ops of this type have waited in queue before being transmitted (microsecond).")
	nfsReadsRttTimeMilliSecondsDesc      = nfsMetricDesc("read_rtt_time_milliseconds_total", "How long the client waited to receive replies of this op type from the server (microsecond).")
	nfsReadsExecuteTimeMilliSecondsDesc  = nfsMetricDesc("read_time_milliseconds_total", "The total number of milliseconds spent by all reads.")
	nfsWritesCompletedDesc               = nfsMetricDesc("write_completed_total", "The total number of writes completed successfully.")
	nfsWritesTransDesc                   = nfsMetricDesc("write_transmissions_total", "How many transmissions of this op type have been sent.")
	nfsWritesTimeOutDesc                 = nfsMetricDesc("write_timeouts_total", "How many timeouts of this op type have occurred.")
	nfsWritesSentBytesDesc               = nfsMetricDesc("write_bytes_total", "The total number of bytes written successfully.")
	nfsWritesRecvBytesDesc               = nfsMetricDesc("write_recv_bytes_total", "How many bytes have been received for this op type.")
	nfsWritesQueueTimeMilliSecondsDesc   = nfsMetricDesc("write_queue_time_milliseconds_total", "How long ops of this type have waited in queue before being transmitted (microsecond).")
	nfsWritesRttTimeMilliSecondsDesc     = nfsMetricDesc("write_rtt_time_milliseconds_total", "How long the client waited to receive replies of this op type from the server (microsecond).")
	nfsWritesExecuteTimeMilliSecondsDesc = nfsMetricDesc("write_time_milliseconds_total", "The total number of milliseconds spent by all writes.")
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
			{MetaDesc: nfsReadsCompletedDesc, valueType: prometheus.CounterValue},
			{MetaDesc: nfsReadsTransDesc, valueType: prometheus.CounterValue},
			{MetaDesc: nfsReadsTimeOutDesc, valueType: prometheus.CounterValue},
			{MetaDesc: nfsReadsSentBytesDesc, valueType: prometheus.CounterValue},
			{MetaDesc: nfsReadsRecvBytesDesc, valueType: prometheus.CounterValue},
			{MetaDesc: nfsReadsQueueTimeMilliSecondsDesc, valueType: prometheus.CounterValue},
			{MetaDesc: nfsReadsRttTimeMilliSecondsDesc, valueType: prometheus.CounterValue},
			{MetaDesc: nfsReadsExecuteTimeMilliSecondsDesc, valueType: prometheus.CounterValue},
			//write
			{MetaDesc: nfsWritesCompletedDesc, valueType: prometheus.CounterValue},
			{MetaDesc: nfsWritesTransDesc, valueType: prometheus.CounterValue},
			{MetaDesc: nfsWritesTimeOutDesc, valueType: prometheus.CounterValue},
			{MetaDesc: nfsWritesSentBytesDesc, valueType: prometheus.CounterValue},
			{MetaDesc: nfsWritesRecvBytesDesc, valueType: prometheus.CounterValue},
			{MetaDesc: nfsWritesQueueTimeMilliSecondsDesc, valueType: prometheus.CounterValue},
			{MetaDesc: nfsWritesRttTimeMilliSecondsDesc, valueType: prometheus.CounterValue},
			{MetaDesc: nfsWritesExecuteTimeMilliSecondsDesc, valueType: prometheus.CounterValue},
		},
		lastPvNfsInfoMap: make(map[string]nfsInfo, 0),
		lastPvStatsMap:   sync.Map{},
		clientSet:        clientset,
		crdClient:        crdClient,
		mounter:          mount.NewWithoutSystemd(""),
	}, nil
}

func (p *nfsStatCollector) Get(ctx context.Context, pvcs sets.Set[string]) ([]*Metric, error) {
	pvNameStatsMap, err := getNfsStat()
	if len(pvNameStatsMap) == 0 {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("couldn't get nfsstats: %s", err)
	}
	volJSONPaths, err := findVolJSON(podsRootPath)
	if err != nil {
		return nil, err
	}
	p.updateMap(ctx, &p.lastPvNfsInfoMap, volJSONPaths, nasDriverName)
	var metrics []*Metric
	for pvName, stats := range pvNameStatsMap {
		if err := ctx.Err(); err != nil {
			return metrics, err
		}
		nfsInfo := p.lastPvNfsInfoMap[pvName]
		metrics = append(metrics, p.getNfsMetrics(pvName, nfsInfo.PvcNamespace, nfsInfo.PvcName, nfsInfo.ServerName, stats)...)
	}
	return metrics, nil
}

func (p *nfsStatCollector) getNfsMetrics(pvName string, pvcNamespace string, pvcName string, serverName string, stats []string) (metrics []*Metric) {
	defer p.lastPvStatsMap.Store(pvName, stats)
	for i, value := range stats {
		if i >= len(p.descs) {
			return metrics
		}
		valueFloat64, err := strconv.ParseFloat(value, 64)
		if err != nil {
			klog.Errorf("Convert value %s to float64 is failed, err:%s, stat:%+v", value, err, stats)
			continue
		}

		metrics = append(metrics, p.descs[i].mustNewMetric(valueFloat64, pvcNamespace, pvcName, serverName, nasStorageName))
	}
	return metrics
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
