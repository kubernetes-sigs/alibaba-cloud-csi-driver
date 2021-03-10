package metric

import (
	"fmt"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/record"
	"k8s.io/kubernetes/pkg/volume/util/fs"
	"os"
	"strconv"
	"strings"
	"sync"
)

var (
	nfsStatLabelNames = []string{"namespace", "pvc", "server", "type"}
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

type nfsStatCollector struct {
	descs            []typedFactorDesc
	lastPvNfsInfoMap map[string]nfsInfo
	lastPvStatsMap   sync.Map
	clientSet        *kubernetes.Clientset
	recorder         record.EventRecorder
}

func init() {
	registerCollector("nfsstat", NewNfsStatCollector)
}

// NewNfsStatCollector returns a new Collector exposing nfs stats.
func NewNfsStatCollector() (Collector, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	recorder := utils.NewEventRecorder()
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

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
			{desc: nfsCapacityTotalDesc, valueType: prometheus.CounterValue},
			{desc: nfsCapacityUsedDesc, valueType: prometheus.CounterValue},
			{desc: nfsCapacityAvailableDesc, valueType: prometheus.CounterValue},
		},
		lastPvNfsInfoMap: make(map[string]nfsInfo, 0),
		lastPvStatsMap:   sync.Map{},
		clientSet:        clientset,
		recorder:         recorder,
	}, nil
}

func (p *nfsStatCollector) Update(ch chan<- prometheus.Metric) error {
	//startTime := time.Now()
	pvNameStatsMap, err := getNfsStat()
	if len(pvNameStatsMap) == 0 || pvNameStatsMap == nil {
		return nil
	}

	if err != nil {
		return fmt.Errorf("couldn't get nfsstats: %s", err)
	}
	volJSONPaths, err := findVolJSON(podsRootPath)
	if err != nil {
		logrus.Errorf("Find nfs vol_data json is failed, err:%s", err)
		return err
	}
	p.updateMap(&p.lastPvNfsInfoMap, volJSONPaths, nasDriverName, "volumes")

	wg := sync.WaitGroup{}
	for pvName, stats := range pvNameStatsMap {
		nfsInfo := p.lastPvNfsInfoMap[pvName]
		getNfsCapacityStat(pvName, nfsInfo, &stats)
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
			logrus.Errorf("Convert value %s to float64 is failed, err:%s, stat:%+v", value, err, stats)
			continue
		}

		ch <- p.descs[i].mustNewConstMetric(valueFloat64, pvcNamespace, pvcName, serverName, nasStorageName)
	}
}

func (p *nfsStatCollector) updateMap(lastPvNfsInfoMap *map[string]nfsInfo, jsonPaths []string, deriverName string, keyword string) {
	thisPvNfsInfoMap := make(map[string]nfsInfo, 0)
	cmd := "mount | grep nfs | grep csi | grep " + keyword
	line, err := utils.Run(cmd)
	if err != nil && strings.Contains(err.Error(), "with out: , with error:") {
		p.updateNfsInfoMap(thisPvNfsInfoMap, lastPvNfsInfoMap)
		return
	}
	if err != nil {
		logrus.Errorf("Execute cmd %s is failed, err: %s", cmd, err)
		return
	}
	for _, path := range jsonPaths {
		//Get nfs pvName
		pvName, _, err := getVolumeInfoByJSON(path, deriverName)
		if err != nil {
			if err.Error() != "VolumeType is not the expected type" {
				logrus.Errorf("Get volume info by path %s is failed, err:%s", path, err)
			}
			continue
		}

		if !strings.Contains(line, "/"+pvName+"/") {
			continue
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
	for pv, thisInfo := range thisPvNfsInfoMap {
		lastInfo, ok := (*lastPvNfsInfoMap)[pv]
		// add and modify
		if !ok || thisInfo.VolDataPath != lastInfo.VolDataPath {
			pvcNamespace, pvcName, serverName, err := getPvcByPvNameByNas(p.clientSet, pv)
			if err != nil {
				logrus.Errorf("Get pvc by pv %s is failed, err:%s", pv, err.Error())
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

func getNfsCapacityStat(pvName string, info nfsInfo, stat *[]string) error {
	mountPath := strings.Replace(info.VolDataPath, "/vol_data.json", "", -1)
	mountPath = mountPath + "/mount"
	available, capacity, usage, _, _, _, err := fs.FsInfo(mountPath)
	if err != nil {
		logrus.Errorf("Get fs info %s is failed,err:%s", mountPath, err)
		return err
	}
	*stat = append(*stat, strconv.Itoa(int(capacity)))
	*stat = append(*stat, strconv.Itoa(int(usage)))
	*stat = append(*stat, strconv.Itoa(int(available)))
	return nil
}
