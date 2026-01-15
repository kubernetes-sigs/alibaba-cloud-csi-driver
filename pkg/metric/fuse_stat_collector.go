package metric

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/prometheus/client_golang/prometheus"
	"k8s.io/klog/v2"
)

var (
	fsClientPathPrefix      = "/host/var/run/"
	fsClientTypeArray       = []string{"ossfs", "efc"}
	counterTypeArray        = utils.CounterTypeMetricsArray
	backendCounterTypeArray = []string{"backend_throughput_counter", "backend_iops_counter", "backend_latency_counter", "backend_meta_qps_ounter"}
	hotSpotArray            = utils.HotSpotMetricsArray
	mountPointStatusArray   = utils.MountpointMetricsArray
)

var (
	usFsStatLabelNames = []string{
		"client_name",
		"backend_storage",
		"bucket_name",
		"namespace",
		"pod",
		"pv",
		"mount_point",
		"file_name",
		"exit_reason"}
)

var (
	capacityBytesUsedCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "capacity_bytes_used_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	capacityBytesAvailableCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "capacity_bytes_available_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	capacityBytesTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "capacity_bytes_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	inodeBytesUsedCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "inode_bytes_used_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	inodeBytesAvailableCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "inode_bytes_available_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	inodeBytesTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "inode_bytes_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	readBytesTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "read_bytes_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	writeBytesTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "write_bytes_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	readCompletedTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "read_completed_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	writeCompletedTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "write_completed_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	readTimeMillisecondsTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "read_time_milliseconds_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	writeTimeMillisecondsTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "write_time_milliseconds_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	posixMkdirTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "posix_mkdir_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	posixRmdirTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "posix_rmdir_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	posixOpendirTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "posix_opendir_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	posixReaddirTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "posix_readdir_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	posixWriteTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "posix_write_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	posixFlushTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "posix_flush_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	posixFsyncTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "posix_fsync_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	posixReleaseTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "posix_release_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	posixReadTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "posix_read_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	posixCreateTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "posix_create_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	posixOpenTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "posix_open_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	posixAccessTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "posix_access_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	posixRenameTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "posix_rename_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	posixChownTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "posix_chown_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	posixChmodTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "posix_chmod_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	posixTruncateTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "posix_truncate_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	ossPutObjectTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "oss_put_object_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	ossGetObjectTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "oss_get_object_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	ossHeadObjectTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "oss_head_object_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	ossDeleteObjectTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "oss_delete_object_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	ossPostObjectTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "oss_post_object_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	hotSpotReadFileTopDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "hot_spot_read_file_top"),
		".",
		usFsStatLabelNames, nil,
	)
	hotSpotWriteFileTopDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "hot_spot_write_file_top"),
		".",
		usFsStatLabelNames, nil,
	)
	hotSpotHeadFileTopDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "hot_spot_head_file_top"),
		".",
		usFsStatLabelNames, nil,
	)
)

var (
	backendReadBytesTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "backend_read_bytes_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	backendWriteBytesTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "backend_write_bytes_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	backendReadCompletedTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "backend_read_completed_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	backendWriteCompletedTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "backend_write_completed_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	backendReadTimeMillisecondsTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "backend_read_time_milliseconds_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	backendWriteTimeMillisecondsTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "backend_write_time_milliseconds_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	backendPosixGetAttrTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "backend_posix_getattr_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	backendPosixGetModeTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "backend_posix_getmode_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	backendPosixAccessTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "backend_posix_access_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	backendPosixLookupTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "backend_posix_lookup_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	backendPosixMknodTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "backend_posix_mknod_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	backendPosixRemoveTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "backend_posix_remove_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	backendPosixSetAttrTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "backend_posix_setattr_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	backendPosixLinkTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "backend_posix_link_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	backendPosixReadLinkTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "backend_posix_readlink_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	backendPosixStatfsTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "backend_posix_statfs_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	backendPosixRenameTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "backend_posix_rename_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	backendPosixReaddirTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, "backend_posix_readdir_total_counter"),
		".",
		usFsStatLabelNames, nil,
	)
	mountRetryTotalCounterDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, utils.MetricsMountRetryCount),
		".",
		usFsStatLabelNames, nil,
	)
	mountPointStatusDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, utils.MetricsMountPointStatus),
		".",
		usFsStatLabelNames, nil,
	)
	mountPointFailoverTotalCountDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, utils.MetricsMountPointFailoverCount),
		".",
		usFsStatLabelNames, nil,
	)
	lastFuseClientExitReasonDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubsystem, utils.MetricsLastFuseClientExitReason),
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

// getCommonLabels returns the common label values for metrics
func (p *usFsStatCollector) getCommonLabels(fsClientInfo *fuseInfo, fileName, exitReason string) []string {
	return []string{
		fsClientInfo.ClientName,
		fsClientInfo.BackendStorage,
		fsClientInfo.BucketName,
		fsClientInfo.Namespace,
		fsClientInfo.PodName,
		fsClientInfo.PvName,
		fsClientInfo.MountPoint,
		fileName,
		exitReason,
	}
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
	mountRetryTotalCounter                *typedFactorDesc
	mountPointStatus                      *typedFactorDesc
	mountPointFailoverTotalCounter        *typedFactorDesc
	lastFuseClientExitReason              *typedFactorDesc
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

// NewFuseStatCollector returns a new Collector exposing user space fs stats.
func NewFuseStatCollector() (Collector, error) {
	return &usFsStatCollector{
		hotSpotReadFileTop:  hotSpotReadFileTopDesc,
		hotSpotWriteFileTop: hotSpotWriteFileTopDesc,
		hotSpotHeadFileTop:  hotSpotHeadFileTopDesc,
		capacityBytesCounterDesc: capacityBytesCounterDesc{
			descs: []typedFactorDesc{
				{desc: capacityBytesUsedCounterDesc, valueType: prometheus.CounterValue, labels: usFsStatLabelNames},
				{desc: capacityBytesAvailableCounterDesc, valueType: prometheus.CounterValue, labels: usFsStatLabelNames},
				{desc: capacityBytesTotalCounterDesc, valueType: prometheus.CounterValue, labels: usFsStatLabelNames},
			},
		},
		inodeBytesCounterDesc: inodeBytesCounterDesc{
			descs: []typedFactorDesc{
				{desc: inodeBytesUsedCounterDesc, valueType: prometheus.CounterValue, labels: usFsStatLabelNames},
				{desc: inodeBytesAvailableCounterDesc, valueType: prometheus.CounterValue, labels: usFsStatLabelNames},
				{desc: inodeBytesTotalCounterDesc, valueType: prometheus.CounterValue, labels: usFsStatLabelNames},
			},
		},
		throughputBytesCounterDesc: throughputBytesCounterDesc{
			descs: []typedFactorDesc{
				{desc: readBytesTotalCounterDesc, valueType: prometheus.CounterValue, labels: usFsStatLabelNames},
				{desc: writeBytesTotalCounterDesc, valueType: prometheus.CounterValue, labels: usFsStatLabelNames},
			},
		},
		iopsCompletedCounterDesc: iopsCompletedCounterDesc{
			descs: []typedFactorDesc{
				{desc: readCompletedTotalCounterDesc, valueType: prometheus.CounterValue, labels: usFsStatLabelNames},
				{desc: writeCompletedTotalCounterDesc, valueType: prometheus.CounterValue, labels: usFsStatLabelNames},
			},
		},
		latencyMillisecondsCounterDesc: latencyMillisecondsCounterDesc{
			descs: []typedFactorDesc{
				{desc: readTimeMillisecondsTotalCounterDesc, valueType: prometheus.CounterValue, factor: .001, labels: usFsStatLabelNames},
				{desc: writeTimeMillisecondsTotalCounterDesc, valueType: prometheus.CounterValue, factor: .001, labels: usFsStatLabelNames},
			},
		},
		posixCounterDesc: posixCounterDesc{
			descs: []typedFactorDesc{
				{desc: posixMkdirTotalCounterDesc, valueType: prometheus.CounterValue, labels: usFsStatLabelNames},
				{desc: posixRmdirTotalCounterDesc, valueType: prometheus.CounterValue, labels: usFsStatLabelNames},
				{desc: posixOpendirTotalCounterDesc, valueType: prometheus.CounterValue, labels: usFsStatLabelNames},
				{desc: posixReaddirTotalCounterDesc, valueType: prometheus.CounterValue, labels: usFsStatLabelNames},
				{desc: posixWriteTotalCounterDesc, valueType: prometheus.CounterValue, labels: usFsStatLabelNames},
				{desc: posixFlushTotalCounterDesc, valueType: prometheus.CounterValue, labels: usFsStatLabelNames},
				{desc: posixFsyncTotalCounterDesc, valueType: prometheus.CounterValue, labels: usFsStatLabelNames},
				{desc: posixReleaseTotalCounterDesc, valueType: prometheus.CounterValue, labels: usFsStatLabelNames},
				{desc: posixReadTotalCounterDesc, valueType: prometheus.CounterValue, labels: usFsStatLabelNames},
				{desc: posixCreateTotalCounterDesc, valueType: prometheus.CounterValue, labels: usFsStatLabelNames},
				{desc: posixOpenTotalCounterDesc, valueType: prometheus.CounterValue, labels: usFsStatLabelNames},
				{desc: posixAccessTotalCounterDesc, valueType: prometheus.CounterValue, labels: usFsStatLabelNames},
				{desc: posixRenameTotalCounterDesc, valueType: prometheus.CounterValue, labels: usFsStatLabelNames},
				{desc: posixChownTotalCounterDesc, valueType: prometheus.CounterValue, labels: usFsStatLabelNames},
				{desc: posixChmodTotalCounterDesc, valueType: prometheus.CounterValue, labels: usFsStatLabelNames},
				{desc: posixTruncateTotalCounterDesc, valueType: prometheus.CounterValue, labels: usFsStatLabelNames},
			},
		},
		ossObjectCounterDesc: ossObjectCounterDesc{
			descs: []typedFactorDesc{
				{desc: ossPutObjectTotalCounterDesc, valueType: prometheus.CounterValue, labels: usFsStatLabelNames},
				{desc: ossGetObjectTotalCounterDesc, valueType: prometheus.CounterValue, labels: usFsStatLabelNames},
				{desc: ossHeadObjectTotalCounterDesc, valueType: prometheus.CounterValue, labels: usFsStatLabelNames},
				{desc: ossDeleteObjectTotalCounterDesc, valueType: prometheus.CounterValue, labels: usFsStatLabelNames},
				{desc: ossPostObjectTotalCounterDesc, valueType: prometheus.CounterValue, labels: usFsStatLabelNames},
			},
		},
		backendThroughputBytesCounterDesc: backendThroughputBytesCounterDesc{
			descs: []typedFactorDesc{
				{desc: backendReadBytesTotalCounterDesc, valueType: prometheus.CounterValue, labels: usFsStatLabelNames},
				{desc: backendWriteBytesTotalCounterDesc, valueType: prometheus.CounterValue, labels: usFsStatLabelNames},
			},
		},
		backendIOPSCompletedCounterDesc: backendIOPSCompletedCounterDesc{
			descs: []typedFactorDesc{
				{desc: backendReadCompletedTotalCounterDesc, valueType: prometheus.CounterValue, labels: usFsStatLabelNames},
				{desc: backendWriteCompletedTotalCounterDesc, valueType: prometheus.CounterValue, labels: usFsStatLabelNames},
			},
		},
		backendLatencyMillisecondsCounterDesc: backendLatencyMillisecondsCounterDesc{
			descs: []typedFactorDesc{
				{desc: backendReadTimeMillisecondsTotalCounterDesc, valueType: prometheus.CounterValue, factor: .001, labels: usFsStatLabelNames},
				{desc: backendWriteTimeMillisecondsTotalCounterDesc, valueType: prometheus.CounterValue, factor: .001, labels: usFsStatLabelNames},
			},
		},
		backendPosixCounterDesc: backendPosixCounterDesc{
			descs: []typedFactorDesc{
				{desc: backendPosixGetAttrTotalCounterDesc, valueType: prometheus.CounterValue, labels: usFsStatLabelNames},
				{desc: backendPosixGetModeTotalCounterDesc, valueType: prometheus.CounterValue, labels: usFsStatLabelNames},
				{desc: backendPosixAccessTotalCounterDesc, valueType: prometheus.CounterValue, labels: usFsStatLabelNames},
				{desc: backendPosixLookupTotalCounterDesc, valueType: prometheus.CounterValue, labels: usFsStatLabelNames},
				{desc: backendPosixMknodTotalCounterDesc, valueType: prometheus.CounterValue, labels: usFsStatLabelNames},
				{desc: backendPosixRemoveTotalCounterDesc, valueType: prometheus.CounterValue, labels: usFsStatLabelNames},
				{desc: backendPosixSetAttrTotalCounterDesc, valueType: prometheus.CounterValue, labels: usFsStatLabelNames},
				{desc: backendPosixLinkTotalCounterDesc, valueType: prometheus.CounterValue, labels: usFsStatLabelNames},
				{desc: backendPosixReadLinkTotalCounterDesc, valueType: prometheus.CounterValue, labels: usFsStatLabelNames},
				{desc: backendPosixStatfsTotalCounterDesc, valueType: prometheus.CounterValue, labels: usFsStatLabelNames},
				{desc: backendPosixRenameTotalCounterDesc, valueType: prometheus.CounterValue, labels: usFsStatLabelNames},
				{desc: backendPosixReaddirTotalCounterDesc, valueType: prometheus.CounterValue, labels: usFsStatLabelNames},
			},
		},
		mountRetryTotalCounter:         &typedFactorDesc{desc: mountRetryTotalCounterDesc, valueType: prometheus.CounterValue, labels: usFsStatLabelNames},
		mountPointStatus:               &typedFactorDesc{desc: mountPointStatusDesc, valueType: prometheus.GaugeValue, labels: usFsStatLabelNames},
		mountPointFailoverTotalCounter: &typedFactorDesc{desc: mountPointFailoverTotalCountDesc, valueType: prometheus.CounterValue, labels: usFsStatLabelNames},
		lastFuseClientExitReason:       &typedFactorDesc{desc: lastFuseClientExitReasonDesc, valueType: prometheus.GaugeValue, labels: usFsStatLabelNames},
	}, nil
}

func getSubDirArray(fsClientPathPrefix string, fsClientType string) ([]string, error) {
	fsClientPath := fsClientPathPrefix + fsClientType
	if !utils.IsFileExisting(fsClientPath) {
		_ = os.MkdirAll(fsClientPath, os.FileMode(0755))
	}
	return listDirectory(fsClientPathPrefix + fsClientType)
}

func (p *usFsStatCollector) postHotTopFileMetrics(hotSpotType string, fsClientInfo *fuseInfo, metricsArray []string) (metrics []*Metric) {
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
		labels := p.getCommonLabels(fsClientInfo, fileName, "")
		switch hotSpotType {
		case "hot_spot_read_file_top":
			metrics = append(metrics, MustNewMetric(p.hotSpotReadFileTop, valueFloat64, prometheus.GaugeValue, usFsStatLabelNames, labels...))
		case "hot_spot_write_file_top":
			metrics = append(metrics, MustNewMetric(p.hotSpotWriteFileTop, valueFloat64, prometheus.GaugeValue, usFsStatLabelNames, labels...))
		case "hot_spot_head_file_top":
			metrics = append(metrics, MustNewMetric(p.hotSpotHeadFileTop, valueFloat64, prometheus.GaugeValue, usFsStatLabelNames, labels...))
		default:
			klog.Errorf("Unknown hotSpotType:%s", hotSpotType)
		}
	}
	return
}

func (p *usFsStatCollector) postCounterMetrics(counterType string, fsClientInfo *fuseInfo, metricsArray []string) (metrics []*Metric) {
	if len(metricsArray) == 0 {
		return
	}

	for i, value := range metricsArray {
		if len(strings.TrimSpace(value)) == 0 {
			continue
		}
		valueFloat64, err := strconv.ParseFloat(value, 64)
		if err != nil {
			klog.Errorf("Convert value %s to float64 is failed, err:%s, counterType:%s, metricsArray:%+v", value, err, counterType, metricsArray)
			continue
		}

		labels := p.getCommonLabels(fsClientInfo, "", "")
		switch counterType {
		case "capacity_counter":
			if i >= len(p.capacityBytesCounterDesc.descs) {
				return
			}
			metrics = append(metrics, MustNewMetricWithTypedFactorDesc(p.capacityBytesCounterDesc.descs[i], valueFloat64, labels...))
		case "inodes_counter":
			if i >= len(p.inodeBytesCounterDesc.descs) {
				return
			}
			metrics = append(metrics, MustNewMetricWithTypedFactorDesc(p.inodeBytesCounterDesc.descs[i], valueFloat64, labels...))

		case "throughput_counter":
			if i >= len(p.throughputBytesCounterDesc.descs) {
				return
			}
			metrics = append(metrics, MustNewMetricWithTypedFactorDesc(p.throughputBytesCounterDesc.descs[i], valueFloat64, labels...))
		case "iops_counter":
			if i >= len(p.iopsCompletedCounterDesc.descs) {
				return
			}
			metrics = append(metrics, MustNewMetricWithTypedFactorDesc(p.iopsCompletedCounterDesc.descs[i], valueFloat64, labels...))
		case "latency_counter":
			if i >= len(p.latencyMillisecondsCounterDesc.descs) {
				return
			}
			metrics = append(metrics, MustNewMetricWithTypedFactorDesc(p.latencyMillisecondsCounterDesc.descs[i], valueFloat64, labels...))
		case "posix_counter":
			if i >= len(p.posixCounterDesc.descs) {
				return
			}
			metrics = append(metrics, MustNewMetricWithTypedFactorDesc(p.posixCounterDesc.descs[i], valueFloat64, labels...))
		case "oss_object_counter":
			if i >= len(p.ossObjectCounterDesc.descs) {
				return
			}
			metrics = append(metrics, MustNewMetricWithTypedFactorDesc(p.ossObjectCounterDesc.descs[i], valueFloat64, labels...))
		default:
			klog.Errorf("Unknown counterType:%s", counterType)
		}
	}
	return
}

func (p *usFsStatCollector) postBackendCounterMetrics(counterType string, fsClientInfo *fuseInfo, metricsArray []string) (metrics []*Metric) {
	if len(metricsArray) == 0 {
		return
	}

	for i, value := range metricsArray {
		if len(strings.TrimSpace(value)) == 0 {
			continue
		}
		valueFloat64, err := strconv.ParseFloat(value, 64)
		if err != nil {
			klog.Errorf("Convert value %s to float64 is failed, err:%s, counterType:%s, metricsArray:%+v", value, err, counterType, metricsArray)
			continue
		}

		labels := p.getCommonLabels(fsClientInfo, "", "")
		switch counterType {
		case "backend_iops_counter":
			if i >= len(p.backendIOPSCompletedCounterDesc.descs) {
				return
			}
			metrics = append(metrics, MustNewMetricWithTypedFactorDesc(p.backendIOPSCompletedCounterDesc.descs[i], valueFloat64, labels...))
		case "backend_latency_counter":
			if i >= len(p.backendLatencyMillisecondsCounterDesc.descs) {
				return
			}
			metrics = append(metrics, MustNewMetricWithTypedFactorDesc(p.backendLatencyMillisecondsCounterDesc.descs[i], valueFloat64, labels...))
		case "backend_meta_qps_ounter":
			if i >= len(p.backendPosixCounterDesc.descs) {
				return
			}
			metrics = append(metrics, MustNewMetricWithTypedFactorDesc(p.backendPosixCounterDesc.descs[i], valueFloat64, labels...))
		case "backend_throughput_counter":
			if i >= len(p.backendThroughputBytesCounterDesc.descs) {
				return
			}
			metrics = append(metrics, MustNewMetricWithTypedFactorDesc(p.backendThroughputBytesCounterDesc.descs[i], valueFloat64, labels...))
		default:
			klog.Errorf("Unknown counterType:%s", counterType)
		}
	}
	return
}

func (p *usFsStatCollector) postMountPointStatusMetrics(statusType string, fsClientInfo *fuseInfo, metricsArray []string) (metrics []*Metric) {
	var err error
	for _, value := range metricsArray {
		if statusType == utils.MetricsLastFuseClientExitReason {
			labels := p.getCommonLabels(fsClientInfo, "", value)
			metrics = append(metrics, MustNewMetricWithTypedFactorDesc(*p.lastFuseClientExitReason, 1, labels...))
			continue
		}

		var valueFloat64 float64
		valueFloat64, err = strconv.ParseFloat(value, 64)
		if err != nil {
			klog.Errorf("Convert value %s to float64 is failed, err:%s, counterType:%s, metricsArray:%+v", value, err, statusType, metricsArray)
			continue
		}

		labels := p.getCommonLabels(fsClientInfo, "", "")
		switch statusType {
		case utils.MetricsMountRetryCount:
			metrics = append(metrics, MustNewMetricWithTypedFactorDesc(*p.mountRetryTotalCounter, valueFloat64, labels...))
		case utils.MetricsMountPointStatus:
			metrics = append(metrics, MustNewMetricWithTypedFactorDesc(*p.mountPointStatus, valueFloat64, labels...))
		case utils.MetricsMountPointFailoverCount:
			metrics = append(metrics, MustNewMetricWithTypedFactorDesc(*p.mountPointFailoverTotalCounter, valueFloat64, labels...))
		}
	}
	return
}

func (p *usFsStatCollector) Update(ch chan<- prometheus.Metric) error {
	metrics := p.Get()
	for _, metric := range metrics {
		ch <- prometheus.MustNewConstMetric(metric.Desc, metric.ValueType, metric.Value, convertLabelsToString(metric.Labels)...)
	}
	return nil
}

func (p *usFsStatCollector) Get() (metrics []*Metric) {
	fsClientInfo := new(fuseInfo)
	// foreach fuse client type
	for _, fsClientType := range fsClientTypeArray {
		// exclusive case: podid
		// shared case: sha256(pvname)
		subDirArray, err := getSubDirArray(fsClientPathPrefix, fsClientType)
		if err != nil {
			continue
		}
		//foreach pod uid
		for _, subDir := range subDirArray {
			//stat pod_info, if exists, updateExclusiveMetrics; else updateSharedMetrics
			if utils.IsFileExisting(filepath.Join(fsClientPathPrefix, fsClientType, subDir, utils.PodInfoFile)) {
				// subDir -> podUid
				metrics = append(metrics, p.updateExclusiveMetrics(fsClientType, subDir, fsClientInfo)...)
				continue
			}
			// subDir -> shaVol
			metrics = append(metrics, p.updateSharedMetrics(fsClientType, subDir, fsClientInfo)...)
		}
	}
	return
}

func (p *usFsStatCollector) updateExclusiveMetrics(fsClientType, podUID string, fsClientInfo *fuseInfo) (metrics []*Metric) {
	//get pod info
	podInfoArray, err := readFirstLines(filepath.Join(fsClientPathPrefix, fsClientType, podUID, utils.PodInfoFile))
	if err != nil {
		return
	}
	//namespace pod_name uid top_number
	if len(podInfoArray) < 4 {
		return
	}
	fsClientInfo.Namespace = podInfoArray[0]
	fsClientInfo.PodName = podInfoArray[1]
	fsClientInfo.PodUID = podInfoArray[2]
	// list volume from pod
	volumeArray, err := listDirectory(filepath.Join(fsClientPathPrefix, fsClientType, podUID))
	if err != nil {
		return
	}
	// foreach volume
	for _, volume := range volumeArray {
		volPath := filepath.Join(fsClientPathPrefix, fsClientType, podUID, volume)
		metrics = append(metrics, p.postVolMetrics(volPath, fsClientInfo)...)
	}
	return
}

func (p *usFsStatCollector) updateSharedMetrics(fsClientType, subDir string, fsClientInfo *fuseInfo) []*Metric {
	// /var/run/fsType/sha256(pvname)
	volPath := filepath.Join(fsClientPathPrefix, fsClientType, subDir)
	return p.postVolMetrics(volPath, fsClientInfo)
}

func (p *usFsStatCollector) postVolMetrics(volPath string, fsClientInfo *fuseInfo) (metrics []*Metric) {
	mountPointInfoArray, err := readFirstLines(filepath.Join(volPath, utils.MountPointInfoFile))
	if err != nil {
		return
	}
	//fuse_client storage_type filesystem_id pv_name mount_point
	if len(mountPointInfoArray) < 5 {
		return
	}
	fsClientInfo.ClientName = mountPointInfoArray[0]
	fsClientInfo.BackendStorage = mountPointInfoArray[1]
	fsClientInfo.BucketName = mountPointInfoArray[2]
	fsClientInfo.PvName = mountPointInfoArray[3]
	fsClientInfo.MountPoint = mountPointInfoArray[4]
	// foreach counter metrics
	for _, counterType := range counterTypeArray {
		metricsArray, err := readFirstLines(filepath.Join(volPath, counterType))
		if err != nil {
			continue
		}
		metrics = append(metrics, p.postCounterMetrics(counterType, fsClientInfo, metricsArray)...)
	}
	// foreach hot_top_file metrics
	for _, hotSpotType := range hotSpotArray {
		metricsArray, err := readFirstLines(filepath.Join(volPath, hotSpotType))
		if err != nil {
			continue
		}
		metrics = append(metrics, p.postHotTopFileMetrics(hotSpotType, fsClientInfo, metricsArray)...)
	}
	// foreach backend counter metrics
	for _, backendCounterType := range backendCounterTypeArray {
		metricsArray, err := readFirstLines(filepath.Join(volPath, backendCounterType))
		if err != nil {
			continue
		}
		metrics = append(metrics, p.postBackendCounterMetrics(backendCounterType, fsClientInfo, metricsArray)...)
	}
	// foreach mountpoint status related metrics
	for _, mountPointStatus := range mountPointStatusArray {
		var metricsArray []string
		if mountPointStatus == utils.MetricsLastFuseClientExitReason {
			metrics, err := readAllContent(filepath.Join(volPath, mountPointStatus))
			if err != nil {
				continue
			}
			metricsArray = []string{metrics}
		} else {
			metricsArray, err = readFirstLines(filepath.Join(volPath, mountPointStatus))
			if err != nil {
				continue
			}
		}
		metrics = append(metrics, p.postMountPointStatusMetrics(mountPointStatus, fsClientInfo, metricsArray)...)
	}
	return
}
