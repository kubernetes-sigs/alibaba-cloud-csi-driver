package metric

import (
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

var (
	fsClientPathPrefix = "/host/var/run/"
	fsClientTypeArray  = []string{"ossfs"}
	counterTypeArray   = []string{"info", "capacity_counter", "inodes_counter", "throughput_counter", "iops_counter", "latency_counter", "posix_counter", "oss_object_counter", "frequently_access_file_top"}
)

var (
	usFsStatLabelNames = []string{"client_name", "backend_storage", "namespace", "pod", "pv", "pvc", "mount_point"}
)

var (
	capacityBytesUsedCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "capacity_bytes_used_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	capacityBytesAvailableCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "capacity_bytes_available_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	capacityBytesTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "capacity_bytes_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	inodeBytesUsedCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "inode_bytes_used_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	inodeBytesAvailableCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "inode_bytes_available_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	inodeBytesTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "inode_bytes_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	readBytesTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "read_bytes_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	writeBytesTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "write_bytes_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	readCompletedTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "read_completed_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	writeCompletedTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "write_completed_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	readTimeMillisecondsTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "read_time_milliseconds_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	writeTimeMillisecondsTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "write_time_milliseconds_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	posixMkdirTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "posix_mkdir_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	posixRmdirTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "posix_rmdir_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	posixOpendirTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "posix_opendir_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	posixReaddirTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "posix_readdir_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	posixWriteTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "posix_write_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	posixFlushTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "posix_flush_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	posixFsyncTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "posix_fsync_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	posixReleaseTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "posix_release_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	posixReadTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "posix_read_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	posixCreateTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "posix_create_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	posixOpenTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "posix_open_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	posixAccessTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "posix_access_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	posixRenameTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "posix_rename_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	posixChownTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "posix_chown_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	posixChmodTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "posix_chmod_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	posixTruncateTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "posix_truncate_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	ossPutObjectTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "oss_put_object_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	ossGetObjectTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "oss_get_object_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	ossHeadObjectTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "oss_head_object_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	ossDeleteObjectTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "oss_delete_object_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	ossPostObjectTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "oss_post_object_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	frequentlyReadFileTopNDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "frequently_read_file_top_N"),
		".",
		usFsStatLabelNames, nil,
	)
	frequentlyWriteFileTopNDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "frequently_write_file_top_N"),
		".",
		usFsStatLabelNames, nil,
	)
)

type usFsInfo struct {
	ClientName     string
	BackendStorage string
	Namespace      string
	PodName        string
	PodUID         string
	PvcName        string
	PvName         string
	MountPoint     string
	TopN           string
}

type usFsStatCollector struct {
	descs []typedFactorDesc
}

func init() {
	registerCollector("user_space_fs_stat", NewUsFsStatCollector)
}

// NewUsFsStatCollector returns a new Collector exposing user space fs stats.
func NewUsFsStatCollector() (Collector, error) {
	return &usFsStatCollector{
		descs: []typedFactorDesc{
			//0-2
			{desc: capacityBytesUsedCounterDesc, valueType: prometheus.CounterValue},
			{desc: capacityBytesAvailableCounterDesc, valueType: prometheus.CounterValue},
			{desc: capacityBytesTotalCounterDesc, valueType: prometheus.CounterValue},
			//3-5
			{desc: inodeBytesUsedCounterDesc, valueType: prometheus.CounterValue},
			{desc: inodeBytesAvailableCounterDesc, valueType: prometheus.CounterValue},
			{desc: inodeBytesTotalCounterDesc, valueType: prometheus.CounterValue},
			//6-7
			{desc: readBytesTotalCounterDesc, valueType: prometheus.CounterValue},
			{desc: writeBytesTotalCounterDesc, valueType: prometheus.CounterValue},
			//8-9
			{desc: readCompletedTotalCounterDesc, valueType: prometheus.CounterValue},
			{desc: writeCompletedTotalCounterDesc, valueType: prometheus.CounterValue},
			//10-11
			{desc: readTimeMillisecondsTotalCounterDesc, valueType: prometheus.CounterValue, factor: .001},
			{desc: writeTimeMillisecondsTotalCounterDesc, valueType: prometheus.CounterValue, factor: .001},
			//12-27
			{desc: posixMkdirTotalCounterDesc, valueType: prometheus.CounterValue},
			{desc: posixRmdirTotalCounterDesc, valueType: prometheus.CounterValue},
			{desc: posixOpendirTotalCounterDesc, valueType: prometheus.CounterValue},
			{desc: posixReaddirTotalCounterDesc, valueType: prometheus.CounterValue},
			{desc: posixWriteTotalCounterDesc, valueType: prometheus.CounterValue},
			{desc: posixFlushTotalCounterDesc, valueType: prometheus.CounterValue},
			{desc: posixFsyncTotalCounterDesc, valueType: prometheus.CounterValue},
			{desc: posixReleaseTotalCounterDesc, valueType: prometheus.CounterValue},
			{desc: posixReadTotalCounterDesc, valueType: prometheus.CounterValue},
			{desc: posixCreateTotalCounterDesc, valueType: prometheus.CounterValue},
			{desc: posixOpenTotalCounterDesc, valueType: prometheus.CounterValue},
			{desc: posixAccessTotalCounterDesc, valueType: prometheus.CounterValue},
			{desc: posixRenameTotalCounterDesc, valueType: prometheus.CounterValue},
			{desc: posixChownTotalCounterDesc, valueType: prometheus.CounterValue},
			{desc: posixChmodTotalCounterDesc, valueType: prometheus.CounterValue},
			{desc: posixTruncateTotalCounterDesc, valueType: prometheus.CounterValue},
			//28-32
			{desc: ossPutObjectTotalCounterDesc, valueType: prometheus.CounterValue},
			{desc: ossGetObjectTotalCounterDesc, valueType: prometheus.CounterValue},
			{desc: ossHeadObjectTotalCounterDesc, valueType: prometheus.CounterValue},
			{desc: ossDeleteObjectTotalCounterDesc, valueType: prometheus.CounterValue},
			{desc: ossPostObjectTotalCounterDesc, valueType: prometheus.CounterValue},
			{desc: frequentlyReadFileTopNDesc, valueType: prometheus.CounterValue},
			{desc: frequentlyWriteFileTopNDesc, valueType: prometheus.CounterValue},
		},
	}, nil
}

func getPodUID(fsClientPathPrefix string, fsClientType string) ([]string, error) {
	return listDirectory(fsClientPathPrefix + fsClientType)
}

func setStat(start int, end int, stat *[35]string, metricsArray []string) {
	if len(metricsArray) < end-start+1 {
		return
	}
	for i := start; i <= end; i++ {
		(*stat)[i] = metricsArray[i]
	}
}

func (p *usFsStatCollector) Update(ch chan<- prometheus.Metric) error {
	initFsClientFlag := false
	var stat = [35]string{}
	usFsClientInfo := new(usFsInfo)
	for _, fsClientType := range fsClientTypeArray {
		podUIDArray, err := getPodUID(fsClientPathPrefix, fsClientType)
		if err != nil {
			continue
		}
		for _, podUID := range podUIDArray {
			for _, counterType := range counterTypeArray {
				metricsReadLineArray, err := readLines(fsClientPathPrefix + fsClientType + "/" + podUID + "/" + counterType)
				if err != nil {
					continue
				}
				if len(metricsReadLineArray) == 0 {
					continue
				}
				metricsArray := strings.Split(metricsReadLineArray[0], " ")
				switch counterType {
				case "info":
					if len(metricsArray) < 9 {
						continue
					}
					usFsClientInfo.ClientName = metricsArray[0]
					usFsClientInfo.BackendStorage = metricsArray[1]
					usFsClientInfo.Namespace = metricsArray[2]
					usFsClientInfo.PodName = metricsArray[3]
					usFsClientInfo.PodUID = metricsArray[4]
					usFsClientInfo.PvName = metricsArray[5]
					usFsClientInfo.PvcName = metricsArray[6]
					usFsClientInfo.MountPoint = metricsArray[7]
					usFsClientInfo.TopN = metricsArray[8]
					initFsClientFlag = true
				case "capacity_counter":
					setStat(0, 2, &stat, metricsArray)
				case "inodes_counter":
					setStat(3, 5, &stat, metricsArray)
				case "throughput_counter":
					setStat(6, 7, &stat, metricsArray)
				case "iops_counter":
					setStat(8, 9, &stat, metricsArray)
				case "latency_counter":
					setStat(10, 11, &stat, metricsArray)
				case "posix_counter":
					setStat(12, 27, &stat, metricsArray)
				case "oss_object_counter":
					setStat(28, 32, &stat, metricsArray)
				case "frequently_access_file_top":
					break
				default:
					log.Errorf("Unknow counterType:%s", counterType)
				}
			}
		}
	}
	log.Infof("stat:%+v", stat)
	if initFsClientFlag {
		p.setMetric(usFsClientInfo, stat, ch)
	}
	return nil
}

func (p *usFsStatCollector) setMetric(usFsClientInfo *usFsInfo, stats [35]string, ch chan<- prometheus.Metric) {
	for i, value := range stats {
		if i >= len(p.descs) {
			return
		}

		valueFloat64, err := strconv.ParseFloat(value, 64)
		if err != nil {
			log.Errorf("Convert value %s to float64 is failed, err:%s, stat:%+v", value, err, stats)
			continue
		}

		ch <- p.descs[i].mustNewConstMetric(valueFloat64, usFsClientInfo.ClientName, usFsClientInfo.BackendStorage, usFsClientInfo.Namespace, usFsClientInfo.PodName, usFsClientInfo.PvcName, usFsClientInfo.PvcName, usFsClientInfo.MountPoint)
	}
}
