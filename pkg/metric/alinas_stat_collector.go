package metric

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/options"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	alinasStatLabelNames = []string{"namespace", "pvc", "server", "type"}
)

var (
	//0 - reads completed successfully
	alinasReadsCompletedDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "read_completed_total"),
		"The total number of reads completed successfully.",
		alinasStatLabelNames, nil,
	)
	//2 - read timeout
	alinasReadsTimeOutDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "read_timeouts_total"),
		"How many timeouts of this op type have occurred.",
		alinasStatLabelNames, nil,
	)
	//4 - read recv bytes
	alinasReadsBytesDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "read_bytes_total"),
		"The total number of bytes read successfully.",
		alinasStatLabelNames, nil,
	)
	//5 - read queue time
	alinasReadsQueueTimeMilliSecondsDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "read_queue_time_milliseconds_total"),
		"How long ops of this type have waited in queue before being transmitted (microsecond).",
		alinasStatLabelNames, nil,
	)
	//7 - time spent reading (ms)
	alinasReadTimeMilliSecondsDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "read_time_milliseconds_total"),
		"The total number of seconds spent by all reads.",
		alinasStatLabelNames,
		nil,
	)
	//8 - writes completed successfully
	alinasWritesCompletedDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "write_completed_total"),
		"The total number of writes completed successfully.",
		alinasStatLabelNames, nil,
	)
	//10 - writes timeout
	alinasWritesTimeOutDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "write_timeouts_total"),
		"How many timeouts of this op type have occurred.",
		alinasStatLabelNames, nil,
	)
	//11 - writes send bytes
	alinasWritesBytesDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "write_bytes_total"),
		"The total number of bytes written successfully.",
		alinasStatLabelNames, nil,
	)
	//13 - writes queue time
	alinasWritesQueueTimeMilliSecondsDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "write_queue_time_milliseconds_total"),
		"How long ops of this type have waited in queue before being transmitted (microsecond).",
		alinasStatLabelNames, nil,
	)
	//7 - time spent reading (ms)
	alinasWriteTimeMilliSecondsDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "write_time_milliseconds_total"),
		"The total number of seconds spent by all writes.",
		alinasStatLabelNames,
		nil,
	)
	//16 - capacity available
	alinasCapacityAvailableDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "capacity_bytes_available"),
		"The number of available size(bytes).",
		alinasStatLabelNames,
		nil,
	)
	//17 - capacity total
	alinasCapacityTotalDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "capacity_bytes_total"),
		"The number of total size(bytes).",
		alinasStatLabelNames,
		nil,
	)
	//18 - capacity used
	alinasCapacityUsedDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "capacity_bytes_used"),
		"The number of used size(bytes).",
		alinasStatLabelNames,
		nil,
	)
)

type alinasInfo struct {
	PvcNamespace string
	PvcName      string
	PvName       string
	ServerName   string
	UUID         string
}

type alinasCapacityInfo struct {
	TotalInodeCount int64 `json:"totalInodeCount"`
	UsedInodeCount  int64 `json:"usedInodeCount"`
	TotalSize       int64 `json:"totalSize"`
	UsedSize        int64 `json:"usedSize"`
}

type alinasStatCollector struct {
	descs                 []typedFactorDesc
	clientSet             *kubernetes.Clientset
	crdClient             dynamic.Interface
	lastMountPointInfoMap map[string]alinasInfo
}

func init() {
	registerCollector("alinasstat", NewAlinasStatCollector)
}

// NewAlinasStatCollector returns a new Collector exposing alinas stats.
func NewAlinasStatCollector() (Collector, error) {
	config, err := clientcmd.BuildConfigFromFlags(options.MasterURL, options.Kubeconfig)
	if err != nil {
		return nil, err
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	crdClient, err := dynamic.NewForConfig(config)
	if err != nil {
		log.Fatalf("Failed to create crd client: %v", err)
	}

	return &alinasStatCollector{
		descs: []typedFactorDesc{
			//read
			{desc: alinasReadsCompletedDesc, valueType: prometheus.CounterValue},
			{desc: alinasReadsTimeOutDesc, valueType: prometheus.CounterValue},
			{desc: alinasReadsBytesDesc, valueType: prometheus.CounterValue},
			{desc: alinasReadsQueueTimeMilliSecondsDesc, valueType: prometheus.CounterValue, factor: .001},
			{desc: alinasReadTimeMilliSecondsDesc, valueType: prometheus.CounterValue, factor: .001},
			//write
			{desc: alinasWritesCompletedDesc, valueType: prometheus.CounterValue},
			{desc: alinasWritesTimeOutDesc, valueType: prometheus.CounterValue},
			{desc: alinasWritesBytesDesc, valueType: prometheus.CounterValue},
			{desc: alinasWritesQueueTimeMilliSecondsDesc, valueType: prometheus.CounterValue, factor: .001},
			{desc: alinasWriteTimeMilliSecondsDesc, valueType: prometheus.CounterValue, factor: .001},
			//capacity
			{desc: alinasCapacityTotalDesc, valueType: prometheus.CounterValue},
			{desc: alinasCapacityUsedDesc, valueType: prometheus.CounterValue},
			{desc: alinasCapacityAvailableDesc, valueType: prometheus.CounterValue},
		},
		clientSet:             clientset,
		crdClient:             crdClient,
		lastMountPointInfoMap: make(map[string]alinasInfo, 0),
	}, nil
}

// 959684a3e6-ege66.cn-zhangjiakou.nas.aliyuncs.com:/:4x62bMXP on /var/lib/kubelet/pods/6f7c54d6-571d-4688-97a9-f74f94111e61/volumes/kubernetes.io~csi/nas-4c831575-fee8-452f-9fc9-4161825f55a9/mount type alifuse.aliyun-alinas-unasd (rw,nosuid,nodev,relatime,user_id=0,group_id=0,default_permissions,allow_other)
func getMountPointInfoMap() map[string]alinasInfo {
	m := make(map[string]alinasInfo, 0)
	s1 := "/kubernetes.io~csi/"
	s2 := "/mount"
	s3 := ":"
	cmd := "mount | grep alifuse.aliyun-alinas-eac | grep csi | grep volumes"
	line, err := utils.Run(cmd)
	if err != nil {
		return m
	}
	arr := strings.Split(line, "\n")
	for _, s := range arr {
		arr1 := strings.Split(s, " ")
		if len(arr1) < 3 {
			continue
		}
		if strings.Contains(arr1[2], s1) && strings.Contains(arr1[2], s2) {
			pvStart := strings.Index(arr1[2], s1)
			pvEnd := strings.Index(arr1[2], s2)
			uuidStart := strings.Index(arr1[0], s3)
			serverEnd := strings.Index(arr1[0], s3)
			alinasInfo := alinasInfo{
				PvName:     arr1[2][pvStart+len(s1) : pvEnd],
				UUID:       arr1[0][0:uuidStart],
				ServerName: arr1[0][0:serverEnd],
			}
			m[arr1[0]] = alinasInfo
		}
	}
	return m
}

func getMetric(respStr string) []string {
	s1 := "{"
	s2 := " "
	res := make([]string, 13)
	respArr := strings.Split(respStr, "\n")
	for _, resp := range respArr {
		nameEnd := strings.Index(resp, s1)
		numStart := strings.LastIndex(resp, s2)
		if nameEnd == -1 || numStart == -1 {
			continue
		}
		name := resp[0:nameEnd]
		num := resp[numStart+1:]
		switch name {
		case "read_completed_total":
			res[0] = num
		case "read_timeouts_total":
			res[1] = num
		case "read_bytes_total":
			res[2] = num
		case "read_queue_time_milliseconds_total":
			res[3] = num
		case "read_time_milliseconds_total":
			res[4] = num
		case "write_completed_total":
			res[5] = num
		case "write_timeouts_total":
			res[6] = num
		case "write_bytes_total":
			res[7] = num
		case "write_queue_time_milliseconds_total":
			res[8] = num
		case "write_time_milliseconds_total":
			res[9] = num
		case "capacity_bytes_total":
			res[10] = num
		case "capacity_bytes_available":
			res[12] = num
		}
	}
	capacityTotal, _ := strconv.ParseInt(res[11], 10, 64)
	capacityAvailable, _ := strconv.ParseInt(res[12], 10, 64)
	res[11] = strconv.FormatInt(capacityTotal-capacityAvailable, 10)
	return res
}

func (p *alinasStatCollector) Update(ch chan<- prometheus.Metric) error {
	thisMountPointInfoMap := getMountPointInfoMap()
	if thisMountPointInfoMap == nil || len(thisMountPointInfoMap) == 0 {
		return nil
	}
	p.updateMap(thisMountPointInfoMap)
	wg := sync.WaitGroup{}
	for _, info := range p.lastMountPointInfoMap {
		wg.Add(1)
		if len(info.UUID) != 0 {
			udsClient := NewUDSClient(fmt.Sprintf("/host/var/run/eac/%s.monitor.sock", info.UUID))
			resp, err := udsClient.Get("http://localhost/metric")
			if err != nil {
				continue
			}
			respStr := udsClient.ReadBody(resp)
			if err != nil {
				continue
			}
			defer udsClient.Close(resp)
			stats := getMetric(respStr)
			go func(pvNameArgs string, pvcNamespaceArgs string, pvcNameArgs string, serverNameArgs string, statsArgs []string) {
				defer wg.Done()
				p.setAliNasMetric(pvNameArgs, pvcNamespaceArgs, pvcNameArgs, serverNameArgs, statsArgs, ch)
			}(info.PvName, info.PvcNamespace, info.PvcName, info.ServerName, stats)
		}
	}
	wg.Wait()
	return nil
}

func (p *alinasStatCollector) setAliNasMetric(pvName string, pvcNamespace string, pvcName string, serverName string, stats []string, ch chan<- prometheus.Metric) {
	for i, value := range stats {
		if i >= len(p.descs) {
			return
		}

		valueFloat64, err := strconv.ParseFloat(value, 64)
		if err != nil {
			log.Errorf("Convert value %s to float64 is failed, err:%s, stat:%+v", value, err, stats)
			continue
		}

		ch <- p.descs[i].mustNewConstMetric(valueFloat64, pvcNamespace, pvcName, serverName, aliNasStorageName)
	}
}

func (p *alinasStatCollector) updateMap(thisMountPointInfoMap map[string]alinasInfo) {
	p.updateAliNasInfoMap(thisMountPointInfoMap)
}

func (p *alinasStatCollector) updateAliNasInfoMap(thisMountPointInfoMap map[string]alinasInfo) {
	for mp, thisInfo := range thisMountPointInfoMap {
		_, ok := (p.lastMountPointInfoMap)[mp]
		// not found
		if !ok {
			pvcNamespace, pvcName, serverName, err := getPvcByPvName(p.clientSet, p.crdClient, thisInfo.PvName)
			if err != nil {
				continue
			}
			updateInfo := alinasInfo{
				PvcName:      pvcName,
				PvcNamespace: pvcNamespace,
				ServerName:   serverName,
				UUID:         thisInfo.UUID,
			}
			(p.lastMountPointInfoMap)[mp] = updateInfo
		}
	}
	//if pv exist thisPvStorageInfoMap and not exist lastPvNfsInfoMap, pv should be deleted
	for lastMountPoint := range p.lastMountPointInfoMap {
		_, ok := thisMountPointInfoMap[lastMountPoint]
		if !ok {
			delete(p.lastMountPointInfoMap, lastMountPoint)
		}
	}
}
