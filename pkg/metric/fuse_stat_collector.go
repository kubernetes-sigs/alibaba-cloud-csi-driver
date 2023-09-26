package metric

import (
	"os"
	"strconv"
	"strings"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

var (
	fsClientPathPrefix      = "/host/var/run/"
	fsClientTypeArray       = []string{"ossfs", "efc"}
	podInfo                 = "pod_info"
	mountPointInfo          = "mount_point_info"
	counterTypeArray        = []string{"capacity_counter", "inodes_counter", "throughput_counter", "iops_counter", "latency_counter", "posix_counter", "oss_object_counter"}
	backendCounterTypeArray = []string{"backend_throughput_counter", "backend_iops_counter", "backend_latency_counter", "backend_meta_qps_ounter"}
	hotSpotArray            = []string{"hot_spot_read_file_top", "hot_spot_write_file_top", "hot_spot_head_file_top"}
)

var (
	usFsStatLabelNames = []string{"client_name", "backend_storage", "bucket_name", "namespace", "pod", "pv", "mount_point", "file_name"}
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
	hotSpotReadFileTopDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "hot_spot_read_file_top"),
		".",
		usFsStatLabelNames, nil,
	)
	hotSpotWriteFileTopDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "hot_spot_write_file_top"),
		".",
		usFsStatLabelNames, nil,
	)
	hotSpotHeadFileTopDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "hot_spot_head_file_top"),
		".",
		usFsStatLabelNames, nil,
	)
)

var (
	backendReadBytesTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "backend_read_bytes_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	backendWriteBytesTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "backend_write_bytes_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	backendReadCompletedTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "backend_read_completed_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	backendWriteCompletedTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "backend_write_completed_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	backendReadTimeMillisecondsTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "backend_read_time_milliseconds_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	backendWriteTimeMillisecondsTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "backend_write_time_milliseconds_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	backendPosixGetAttrTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "backend_posix_getattr_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	backendPosixGetModeTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "backend_posix_getmode_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	backendPosixAccessTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "backend_posix_access_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	backendPosixLookupTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "backend_posix_lookup_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	backendPosixMknodTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "backend_posix_mknod_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	backendPosixRemoveTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "backend_posix_remove_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	backendPosixSetAttrTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "backend_posix_setattr_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	backendPosixLinkTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "backend_posix_link_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	backendPosixReadLinkTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "backend_posix_readlink_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	backendPosixStatfsTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "backend_posix_statfs_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	backendPosixRenameTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "backend_posix_rename_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	backendPosixReaddirTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "backend_posix_readdir_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
)

type fuseInfo struct {
	ClientName     string
	BackendStorage string
	BucketName     string
	Namespace      string
	PodName        string
	PodUID         string
	PvName         string
	MountPoint     string
}

type usFsStatCollector struct {
	hotSpotReadFileTop                    *prometheus.Desc
	hotSpotWriteFileTop                   *prometheus.Desc
	hotSpotHeadFileTop                    *prometheus.Desc
	capacityBytesCounterDesc              capacityBytesCounterDesc
	inodeBytesCounterDesc                 inodeBytesCounterDesc
	throughputBytesCounterDesc            throughputBytesCounterDesc
	iopsCompletedCounterDesc              iopsCompletedCounterDesc
	latencyMillisecondsCounterDesc        latencyMillisecondsCounterDesc
	posixCounterDesc                      posixCounterDesc
	ossObjectCounterDesc                  ossObjectCounterDesc
	backendThroughputBytesCounterDesc     backendThroughputBytesCounterDesc
	backendIOPSCompletedCounterDesc       backendIOPSCompletedCounterDesc
	backendLatencyMillisecondsCounterDesc backendLatencyMillisecondsCounterDesc
	backendPosixCounterDesc               backendPosixCounterDesc
}

type capacityBytesCounterDesc struct {
	descs []typedFactorDesc
}
type inodeBytesCounterDesc struct {
	descs []typedFactorDesc
}

type throughputBytesCounterDesc struct {
	descs []typedFactorDesc
}

type iopsCompletedCounterDesc struct {
	descs []typedFactorDesc
}

type latencyMillisecondsCounterDesc struct {
	descs []typedFactorDesc
}

type posixCounterDesc struct {
	descs []typedFactorDesc
}

type ossObjectCounterDesc struct {
	descs []typedFactorDesc
}

type backendThroughputBytesCounterDesc struct {
	descs []typedFactorDesc
}
type backendIOPSCompletedCounterDesc struct {
	descs []typedFactorDesc
}
type backendLatencyMillisecondsCounterDesc struct {
	descs []typedFactorDesc
}
type backendPosixCounterDesc struct {
	descs []typedFactorDesc
}

func init() {
	registerCollector("fuse_stat", NewFuseStatCollector, ossDriverName, nasDriverName)
}

// NewUsFsStatCollector returns a new Collector exposing user space fs stats.
func NewFuseStatCollector() (Collector, error) {
	return &usFsStatCollector{
		hotSpotReadFileTop:  hotSpotReadFileTopDesc,
		hotSpotWriteFileTop: hotSpotWriteFileTopDesc,
		hotSpotHeadFileTop:  hotSpotHeadFileTopDesc,
		capacityBytesCounterDesc: capacityBytesCounterDesc{
			descs: []typedFactorDesc{
				{desc: capacityBytesUsedCounterDesc, valueType: prometheus.CounterValue},
				{desc: capacityBytesAvailableCounterDesc, valueType: prometheus.CounterValue},
				{desc: capacityBytesTotalCounterDesc, valueType: prometheus.CounterValue},
			},
		},
		inodeBytesCounterDesc: inodeBytesCounterDesc{
			descs: []typedFactorDesc{
				{desc: inodeBytesUsedCounterDesc, valueType: prometheus.CounterValue},
				{desc: inodeBytesAvailableCounterDesc, valueType: prometheus.CounterValue},
				{desc: inodeBytesTotalCounterDesc, valueType: prometheus.CounterValue},
			},
		},
		throughputBytesCounterDesc: throughputBytesCounterDesc{
			descs: []typedFactorDesc{
				{desc: readBytesTotalCounterDesc, valueType: prometheus.CounterValue},
				{desc: writeBytesTotalCounterDesc, valueType: prometheus.CounterValue},
			},
		},
		iopsCompletedCounterDesc: iopsCompletedCounterDesc{
			descs: []typedFactorDesc{
				{desc: readCompletedTotalCounterDesc, valueType: prometheus.CounterValue},
				{desc: writeCompletedTotalCounterDesc, valueType: prometheus.CounterValue},
			},
		},
		latencyMillisecondsCounterDesc: latencyMillisecondsCounterDesc{
			descs: []typedFactorDesc{
				{desc: readTimeMillisecondsTotalCounterDesc, valueType: prometheus.CounterValue, factor: .001},
				{desc: writeTimeMillisecondsTotalCounterDesc, valueType: prometheus.CounterValue, factor: .001},
			},
		},
		posixCounterDesc: posixCounterDesc{
			descs: []typedFactorDesc{
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
			},
		},
		ossObjectCounterDesc: ossObjectCounterDesc{
			descs: []typedFactorDesc{
				{desc: ossPutObjectTotalCounterDesc, valueType: prometheus.CounterValue},
				{desc: ossGetObjectTotalCounterDesc, valueType: prometheus.CounterValue},
				{desc: ossHeadObjectTotalCounterDesc, valueType: prometheus.CounterValue},
				{desc: ossDeleteObjectTotalCounterDesc, valueType: prometheus.CounterValue},
				{desc: ossPostObjectTotalCounterDesc, valueType: prometheus.CounterValue},
			},
		},
		backendThroughputBytesCounterDesc: backendThroughputBytesCounterDesc{
			descs: []typedFactorDesc{
				{desc: backendReadBytesTotalCounterDesc, valueType: prometheus.CounterValue},
				{desc: backendWriteBytesTotalCounterDesc, valueType: prometheus.CounterValue},
			},
		},
		backendIOPSCompletedCounterDesc: backendIOPSCompletedCounterDesc{
			descs: []typedFactorDesc{
				{desc: backendReadCompletedTotalCounterDesc, valueType: prometheus.CounterValue},
				{desc: backendWriteCompletedTotalCounterDesc, valueType: prometheus.CounterValue},
			},
		},
		backendLatencyMillisecondsCounterDesc: backendLatencyMillisecondsCounterDesc{
			descs: []typedFactorDesc{
				{desc: backendReadTimeMillisecondsTotalCounterDesc, valueType: prometheus.CounterValue, factor: .001},
				{desc: backendWriteTimeMillisecondsTotalCounterDesc, valueType: prometheus.CounterValue, factor: .001},
			},
		},
		backendPosixCounterDesc: backendPosixCounterDesc{
			descs: []typedFactorDesc{
				{desc: backendPosixGetAttrTotalCounterDesc, valueType: prometheus.CounterValue},
				{desc: backendPosixGetModeTotalCounterDesc, valueType: prometheus.CounterValue},
				{desc: backendPosixAccessTotalCounterDesc, valueType: prometheus.CounterValue},
				{desc: backendPosixLookupTotalCounterDesc, valueType: prometheus.CounterValue},
				{desc: backendPosixMknodTotalCounterDesc, valueType: prometheus.CounterValue},
				{desc: backendPosixRemoveTotalCounterDesc, valueType: prometheus.CounterValue},
				{desc: backendPosixSetAttrTotalCounterDesc, valueType: prometheus.CounterValue},
				{desc: backendPosixLinkTotalCounterDesc, valueType: prometheus.CounterValue},
				{desc: backendPosixReadLinkTotalCounterDesc, valueType: prometheus.CounterValue},
				{desc: backendPosixStatfsTotalCounterDesc, valueType: prometheus.CounterValue},
				{desc: backendPosixRenameTotalCounterDesc, valueType: prometheus.CounterValue},
				{desc: backendPosixReaddirTotalCounterDesc, valueType: prometheus.CounterValue},
			},
		},
	}, nil
}

func getPodUID(fsClientPathPrefix string, fsClientType string) ([]string, error) {
	fsClientPath := fsClientPathPrefix + fsClientType
	if !utils.IsFileExisting(fsClientPath) {
		_ = os.MkdirAll(fsClientPath, os.FileMode(0755))
	}
	return listDirectory(fsClientPathPrefix + fsClientType)
}

func (p *usFsStatCollector) postHotTopFileMetrics(hotSpotType string, fsClientInfo *fuseInfo, metricsArray []string, ch chan<- prometheus.Metric) {
	for _, metricsValue := range metricsArray {
		start := strings.LastIndex(metricsValue, ":")
		if start == -1 {
			continue
		}
		fileName := metricsValue[0:start]
		value := metricsValue[start+1:]
		valueFloat64, err := strconv.ParseFloat(value, 64)
		if err != nil {
			continue
		}
		switch hotSpotType {
		case "hot_spot_read_file_top":
			ch <- prometheus.MustNewConstMetric(p.hotSpotReadFileTop, prometheus.GaugeValue, valueFloat64, fsClientInfo.ClientName, fsClientInfo.BackendStorage, fsClientInfo.BucketName, fsClientInfo.Namespace, fsClientInfo.PodName, fsClientInfo.PvName, fsClientInfo.MountPoint, fileName)
		case "hot_spot_write_file_top":
			ch <- prometheus.MustNewConstMetric(p.hotSpotWriteFileTop, prometheus.GaugeValue, valueFloat64, fsClientInfo.ClientName, fsClientInfo.BackendStorage, fsClientInfo.BucketName, fsClientInfo.Namespace, fsClientInfo.PodName, fsClientInfo.PvName, fsClientInfo.MountPoint, fileName)
		case "hot_spot_head_file_top":
			ch <- prometheus.MustNewConstMetric(p.hotSpotHeadFileTop, prometheus.GaugeValue, valueFloat64, fsClientInfo.ClientName, fsClientInfo.BackendStorage, fsClientInfo.BucketName, fsClientInfo.Namespace, fsClientInfo.PodName, fsClientInfo.PvName, fsClientInfo.MountPoint, fileName)
		default:
			log.Errorf("Unknow hotSpotType:%s", hotSpotType)
		}
	}
}

func (p *usFsStatCollector) postCounterMetrics(counterType string, fsClientInfo *fuseInfo, metricsArray []string, ch chan<- prometheus.Metric) {
	if len(metricsArray) == 0 {
		return
	}

	for i, value := range metricsArray {
		if len(strings.TrimSpace(value)) == 0 {
			continue
		}
		valueFloat64, err := strconv.ParseFloat(value, 64)
		if err != nil {
			log.Errorf("Convert value %s to float64 is failed, err:%s, counterType:%s, metricsArray:%+v", value, err, counterType, metricsArray)
			continue
		}

		switch counterType {
		case "capacity_counter":
			if i >= len(p.capacityBytesCounterDesc.descs) {
				return
			}
			ch <- p.capacityBytesCounterDesc.descs[i].mustNewConstMetric(valueFloat64, fsClientInfo.ClientName, fsClientInfo.BackendStorage, fsClientInfo.BucketName, fsClientInfo.Namespace, fsClientInfo.PodName, fsClientInfo.PvName, fsClientInfo.MountPoint, "")
		case "inodes_counter":
			if i >= len(p.inodeBytesCounterDesc.descs) {
				return
			}
			ch <- p.inodeBytesCounterDesc.descs[i].mustNewConstMetric(valueFloat64, fsClientInfo.ClientName, fsClientInfo.BackendStorage, fsClientInfo.BucketName, fsClientInfo.Namespace, fsClientInfo.PodName, fsClientInfo.PvName, fsClientInfo.MountPoint, "")
		case "throughput_counter":
			if i >= len(p.throughputBytesCounterDesc.descs) {
				return
			}
			ch <- p.throughputBytesCounterDesc.descs[i].mustNewConstMetric(valueFloat64, fsClientInfo.ClientName, fsClientInfo.BackendStorage, fsClientInfo.BucketName, fsClientInfo.Namespace, fsClientInfo.PodName, fsClientInfo.PvName, fsClientInfo.MountPoint, "")
		case "iops_counter":
			if i >= len(p.iopsCompletedCounterDesc.descs) {
				return
			}
			ch <- p.iopsCompletedCounterDesc.descs[i].mustNewConstMetric(valueFloat64, fsClientInfo.ClientName, fsClientInfo.BackendStorage, fsClientInfo.BucketName, fsClientInfo.Namespace, fsClientInfo.PodName, fsClientInfo.PvName, fsClientInfo.MountPoint, "")
		case "latency_counter":
			if i >= len(p.latencyMillisecondsCounterDesc.descs) {
				return
			}
			ch <- p.latencyMillisecondsCounterDesc.descs[i].mustNewConstMetric(valueFloat64, fsClientInfo.ClientName, fsClientInfo.BackendStorage, fsClientInfo.BucketName, fsClientInfo.Namespace, fsClientInfo.PodName, fsClientInfo.PvName, fsClientInfo.MountPoint, "")
		case "posix_counter":
			if i >= len(p.posixCounterDesc.descs) {
				return
			}
			ch <- p.posixCounterDesc.descs[i].mustNewConstMetric(valueFloat64, fsClientInfo.ClientName, fsClientInfo.BackendStorage, fsClientInfo.BucketName, fsClientInfo.Namespace, fsClientInfo.PodName, fsClientInfo.PvName, fsClientInfo.MountPoint, "")
		case "oss_object_counter":
			if i >= len(p.ossObjectCounterDesc.descs) {
				return
			}
			ch <- p.ossObjectCounterDesc.descs[i].mustNewConstMetric(valueFloat64, fsClientInfo.ClientName, fsClientInfo.BackendStorage, fsClientInfo.BucketName, fsClientInfo.Namespace, fsClientInfo.PodName, fsClientInfo.PvName, fsClientInfo.MountPoint, "")
		default:
			log.Errorf("Unknow counterType:%s", counterType)
		}
	}
}

func (p *usFsStatCollector) postBackendCounterMetrics(counterType string, fsClientInfo *fuseInfo, metricsArray []string, ch chan<- prometheus.Metric) {
	if len(metricsArray) == 0 {
		return
	}

	for i, value := range metricsArray {
		if len(strings.TrimSpace(value)) == 0 {
			continue
		}
		valueFloat64, err := strconv.ParseFloat(value, 64)
		if err != nil {
			log.Errorf("Convert value %s to float64 is failed, err:%s, counterType:%s, metricsArray:%+v", value, err, counterType, metricsArray)
			continue
		}

		switch counterType {
		case "backend_iops_counter":
			if i >= len(p.backendIOPSCompletedCounterDesc.descs) {
				return
			}
			ch <- p.backendIOPSCompletedCounterDesc.descs[i].mustNewConstMetric(valueFloat64, fsClientInfo.ClientName, fsClientInfo.BackendStorage, fsClientInfo.BucketName, fsClientInfo.Namespace, fsClientInfo.PodName, fsClientInfo.PvName, fsClientInfo.MountPoint, "")
		case "backend_latency_counter":
			if i >= len(p.backendLatencyMillisecondsCounterDesc.descs) {
				return
			}
			ch <- p.backendLatencyMillisecondsCounterDesc.descs[i].mustNewConstMetric(valueFloat64, fsClientInfo.ClientName, fsClientInfo.BackendStorage, fsClientInfo.BucketName, fsClientInfo.Namespace, fsClientInfo.PodName, fsClientInfo.PvName, fsClientInfo.MountPoint, "")
		case "backend_meta_qps_ounter":
			if i >= len(p.backendPosixCounterDesc.descs) {
				return
			}
			ch <- p.backendPosixCounterDesc.descs[i].mustNewConstMetric(valueFloat64, fsClientInfo.ClientName, fsClientInfo.BackendStorage, fsClientInfo.BucketName, fsClientInfo.Namespace, fsClientInfo.PodName, fsClientInfo.PvName, fsClientInfo.MountPoint, "")
		case "backend_throughput_counter":
			if i >= len(p.backendThroughputBytesCounterDesc.descs) {
				return
			}
			ch <- p.backendThroughputBytesCounterDesc.descs[i].mustNewConstMetric(valueFloat64, fsClientInfo.ClientName, fsClientInfo.BackendStorage, fsClientInfo.BucketName, fsClientInfo.Namespace, fsClientInfo.PodName, fsClientInfo.PvName, fsClientInfo.MountPoint, "")
		default:
			log.Errorf("Unknow counterType:%s", counterType)
		}
	}
}

func (p *usFsStatCollector) Update(ch chan<- prometheus.Metric) error {
	fsClientInfo := new(fuseInfo)
	// foreach fuse client type
	for _, fsClientType := range fsClientTypeArray {
		// get pod uid
		podUIDArray, err := getPodUID(fsClientPathPrefix, fsClientType)
		if err != nil {
			continue
		}
		//foreach pod uid
		for _, podUID := range podUIDArray {
			//get pod info
			podInfoArray, err := readFirstLines(fsClientPathPrefix + fsClientType + "/" + podUID + "/" + podInfo)
			if err != nil {
				continue
			}
			//namespace pod_name uid top_number
			if len(podInfoArray) < 4 {
				continue
			}
			fsClientInfo.Namespace = podInfoArray[0]
			fsClientInfo.PodName = podInfoArray[1]
			fsClientInfo.PodUID = podInfoArray[2]
			// list volume from pod
			volumeArray, err := listDirectory(fsClientPathPrefix + fsClientType + "/" + podUID + "/")
			if err != nil {
				continue
			}
			// foreach volume
			for _, volume := range volumeArray {
				mountPointInfoArray, err := readFirstLines(fsClientPathPrefix + fsClientType + "/" + podUID + "/" + volume + "/" + mountPointInfo)
				if err != nil {
					continue
				}
				//fuse_client storage_type filesystem_id pv_name mount_point
				if len(mountPointInfoArray) < 5 {
					continue
				}
				fsClientInfo.ClientName = mountPointInfoArray[0]
				fsClientInfo.BackendStorage = mountPointInfoArray[1]
				fsClientInfo.BucketName = mountPointInfoArray[2]
				fsClientInfo.PvName = mountPointInfoArray[3]
				fsClientInfo.MountPoint = mountPointInfoArray[4]
				// foreach counter metrics
				for _, counterType := range counterTypeArray {
					metricsArray, err := readFirstLines(fsClientPathPrefix + fsClientType + "/" + podUID + "/" + volume + "/" + counterType)
					if err != nil {
						continue
					}
					p.postCounterMetrics(counterType, fsClientInfo, metricsArray, ch)
				}
				// foreach hot_top_file metrics
				for _, hotSpotType := range hotSpotArray {
					metricsArray, err := readFirstLines(fsClientPathPrefix + fsClientType + "/" + podUID + "/" + volume + "/" + hotSpotType)
					if err != nil {
						continue
					}
					p.postHotTopFileMetrics(hotSpotType, fsClientInfo, metricsArray, ch)
				}
				// foreach backend counter metrics
				for _, backendCounterType := range backendCounterTypeArray {
					metricsArray, err := readFirstLines(fsClientPathPrefix + fsClientType + "/" + podUID + "/" + volume + "/" + backendCounterType)
					if err != nil {
						continue
					}
					p.postBackendCounterMetrics(backendCounterType, fsClientInfo, metricsArray, ch)
				}
			}
		}
	}
	return nil
}
