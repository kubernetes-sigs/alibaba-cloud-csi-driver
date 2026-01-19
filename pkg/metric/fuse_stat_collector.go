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

func fuseMetricDesc(name, help string) *MetaDesc {
	return NewMetaDesc(nodeNamespace, volumeSubsystem, name, help, usFsStatLabelNames, nil)
}

var (
	capacityBytesUsedCounterDesc          = fuseMetricDesc("capacity_bytes_used_counter", ".")
	capacityBytesAvailableCounterDesc     = fuseMetricDesc("capacity_bytes_available_counter", ".")
	capacityBytesTotalCounterDesc         = fuseMetricDesc("capacity_bytes_total_counter", ".")
	inodeBytesUsedCounterDesc             = fuseMetricDesc("inode_bytes_used_counter", ".")
	inodeBytesAvailableCounterDesc        = fuseMetricDesc("inode_bytes_available_counter", ".")
	inodeBytesTotalCounterDesc            = fuseMetricDesc("inode_bytes_total_counter", ".")
	readBytesTotalCounterDesc             = fuseMetricDesc("read_bytes_total_counter", ".")
	writeBytesTotalCounterDesc            = fuseMetricDesc("write_bytes_total_counter", ".")
	readCompletedTotalCounterDesc         = fuseMetricDesc("read_completed_total_counter", ".")
	writeCompletedTotalCounterDesc        = fuseMetricDesc("write_completed_total_counter", ".")
	readTimeMillisecondsTotalCounterDesc  = fuseMetricDesc("read_time_milliseconds_total_counter", ".")
	writeTimeMillisecondsTotalCounterDesc = fuseMetricDesc("write_time_milliseconds_total_counter", ".")
	posixMkdirTotalCounterDesc            = fuseMetricDesc("posix_mkdir_total_counter", ".")
	posixRmdirTotalCounterDesc            = fuseMetricDesc("posix_rmdir_total_counter", ".")
	posixOpendirTotalCounterDesc          = fuseMetricDesc("posix_opendir_total_counter", ".")
	posixReaddirTotalCounterDesc          = fuseMetricDesc("posix_readdir_total_counter", ".")
	posixWriteTotalCounterDesc            = fuseMetricDesc("posix_write_total_counter", ".")
	posixFlushTotalCounterDesc            = fuseMetricDesc("posix_flush_total_counter", ".")
	posixFsyncTotalCounterDesc            = fuseMetricDesc("posix_fsync_total_counter", ".")
	posixReleaseTotalCounterDesc          = fuseMetricDesc("posix_release_total_counter", ".")
	posixReadTotalCounterDesc             = fuseMetricDesc("posix_read_total_counter", ".")
	posixCreateTotalCounterDesc           = fuseMetricDesc("posix_create_total_counter", ".")
	posixOpenTotalCounterDesc             = fuseMetricDesc("posix_open_total_counter", ".")
	posixAccessTotalCounterDesc           = fuseMetricDesc("posix_access_total_counter", ".")
	posixRenameTotalCounterDesc           = fuseMetricDesc("posix_rename_total_counter", ".")
	posixChownTotalCounterDesc            = fuseMetricDesc("posix_chown_total_counter", ".")
	posixChmodTotalCounterDesc            = fuseMetricDesc("posix_chmod_total_counter", ".")
	posixTruncateTotalCounterDesc         = fuseMetricDesc("posix_truncate_total_counter", ".")
	ossPutObjectTotalCounterDesc          = fuseMetricDesc("oss_put_object_total_counter", ".")
	ossGetObjectTotalCounterDesc          = fuseMetricDesc("oss_get_object_total_counter", ".")
	ossHeadObjectTotalCounterDesc         = fuseMetricDesc("oss_head_object_total_counter", ".")
	ossDeleteObjectTotalCounterDesc       = fuseMetricDesc("oss_delete_object_total_counter", ".")
	ossPostObjectTotalCounterDesc         = fuseMetricDesc("oss_post_object_total_counter", ".")
	hotSpotReadFileTopDesc                = fuseMetricDesc("hot_spot_read_file_top", ".")
	hotSpotWriteFileTopDesc               = fuseMetricDesc("hot_spot_write_file_top", ".")
	hotSpotHeadFileTopDesc                = fuseMetricDesc("hot_spot_head_file_top", ".")
)

var (
	backendReadBytesTotalCounterDesc             = fuseMetricDesc("backend_read_bytes_total_counter", ".")
	backendWriteBytesTotalCounterDesc            = fuseMetricDesc("backend_write_bytes_total_counter", ".")
	backendReadCompletedTotalCounterDesc         = fuseMetricDesc("backend_read_completed_total_counter", ".")
	backendWriteCompletedTotalCounterDesc        = fuseMetricDesc("backend_write_completed_total_counter", ".")
	backendReadTimeMillisecondsTotalCounterDesc  = fuseMetricDesc("backend_read_time_milliseconds_total_counter", ".")
	backendWriteTimeMillisecondsTotalCounterDesc = fuseMetricDesc("backend_write_time_milliseconds_total_counter", ".")
	backendPosixGetAttrTotalCounterDesc          = fuseMetricDesc("backend_posix_getattr_total_counter", ".")
	backendPosixGetModeTotalCounterDesc          = fuseMetricDesc("backend_posix_getmode_total_counter", ".")
	backendPosixAccessTotalCounterDesc           = fuseMetricDesc("backend_posix_access_total_counter", ".")
	backendPosixLookupTotalCounterDesc           = fuseMetricDesc("backend_posix_lookup_total_counter", ".")
	backendPosixMknodTotalCounterDesc            = fuseMetricDesc("backend_posix_mknod_total_counter", ".")
	backendPosixRemoveTotalCounterDesc           = fuseMetricDesc("backend_posix_remove_total_counter", ".")
	backendPosixSetAttrTotalCounterDesc          = fuseMetricDesc("backend_posix_setattr_total_counter", ".")
	backendPosixLinkTotalCounterDesc             = fuseMetricDesc("backend_posix_link_total_counter", ".")
	backendPosixReadLinkTotalCounterDesc         = fuseMetricDesc("backend_posix_readlink_total_counter", ".")
	backendPosixStatfsTotalCounterDesc           = fuseMetricDesc("backend_posix_statfs_total_counter", ".")
	backendPosixRenameTotalCounterDesc           = fuseMetricDesc("backend_posix_rename_total_counter", ".")
	backendPosixReaddirTotalCounterDesc          = fuseMetricDesc("backend_posix_readdir_total_counter", ".")
	mountRetryTotalCounterDesc                   = fuseMetricDesc(utils.MetricsMountRetryCount, ".")
	mountPointStatusDesc                         = fuseMetricDesc(utils.MetricsMountPointStatus, ".")
	mountPointFailoverTotalCountDesc             = fuseMetricDesc(utils.MetricsMountPointFailoverCount, ".")
	lastFuseClientExitReasonDesc                 = fuseMetricDesc(utils.MetricsLastFuseClientExitReason, ".")
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
	hotSpotReadFileTop                    *MetaDesc
	hotSpotWriteFileTop                   *MetaDesc
	hotSpotHeadFileTop                    *MetaDesc
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
				{MetaDesc: capacityBytesUsedCounterDesc, valueType: prometheus.CounterValue},
				{MetaDesc: capacityBytesAvailableCounterDesc, valueType: prometheus.CounterValue},
				{MetaDesc: capacityBytesTotalCounterDesc, valueType: prometheus.CounterValue},
			},
		},
		inodeBytesCounterDesc: inodeBytesCounterDesc{
			descs: []typedFactorDesc{
				{MetaDesc: inodeBytesUsedCounterDesc, valueType: prometheus.CounterValue},
				{MetaDesc: inodeBytesAvailableCounterDesc, valueType: prometheus.CounterValue},
				{MetaDesc: inodeBytesTotalCounterDesc, valueType: prometheus.CounterValue},
			},
		},
		throughputBytesCounterDesc: throughputBytesCounterDesc{
			descs: []typedFactorDesc{
				{MetaDesc: readBytesTotalCounterDesc, valueType: prometheus.CounterValue},
				{MetaDesc: writeBytesTotalCounterDesc, valueType: prometheus.CounterValue},
			},
		},
		iopsCompletedCounterDesc: iopsCompletedCounterDesc{
			descs: []typedFactorDesc{
				{MetaDesc: readCompletedTotalCounterDesc, valueType: prometheus.CounterValue},
				{MetaDesc: writeCompletedTotalCounterDesc, valueType: prometheus.CounterValue},
			},
		},
		latencyMillisecondsCounterDesc: latencyMillisecondsCounterDesc{
			descs: []typedFactorDesc{
				{MetaDesc: readTimeMillisecondsTotalCounterDesc, valueType: prometheus.CounterValue, factor: .001},
				{MetaDesc: writeTimeMillisecondsTotalCounterDesc, valueType: prometheus.CounterValue, factor: .001},
			},
		},
		posixCounterDesc: posixCounterDesc{
			descs: []typedFactorDesc{
				{MetaDesc: posixMkdirTotalCounterDesc, valueType: prometheus.CounterValue},
				{MetaDesc: posixRmdirTotalCounterDesc, valueType: prometheus.CounterValue},
				{MetaDesc: posixOpendirTotalCounterDesc, valueType: prometheus.CounterValue},
				{MetaDesc: posixReaddirTotalCounterDesc, valueType: prometheus.CounterValue},
				{MetaDesc: posixWriteTotalCounterDesc, valueType: prometheus.CounterValue},
				{MetaDesc: posixFlushTotalCounterDesc, valueType: prometheus.CounterValue},
				{MetaDesc: posixFsyncTotalCounterDesc, valueType: prometheus.CounterValue},
				{MetaDesc: posixReleaseTotalCounterDesc, valueType: prometheus.CounterValue},
				{MetaDesc: posixReadTotalCounterDesc, valueType: prometheus.CounterValue},
				{MetaDesc: posixCreateTotalCounterDesc, valueType: prometheus.CounterValue},
				{MetaDesc: posixOpenTotalCounterDesc, valueType: prometheus.CounterValue},
				{MetaDesc: posixAccessTotalCounterDesc, valueType: prometheus.CounterValue},
				{MetaDesc: posixRenameTotalCounterDesc, valueType: prometheus.CounterValue},
				{MetaDesc: posixChownTotalCounterDesc, valueType: prometheus.CounterValue},
				{MetaDesc: posixChmodTotalCounterDesc, valueType: prometheus.CounterValue},
				{MetaDesc: posixTruncateTotalCounterDesc, valueType: prometheus.CounterValue},
			},
		},
		ossObjectCounterDesc: ossObjectCounterDesc{
			descs: []typedFactorDesc{
				{MetaDesc: ossPutObjectTotalCounterDesc, valueType: prometheus.CounterValue},
				{MetaDesc: ossGetObjectTotalCounterDesc, valueType: prometheus.CounterValue},
				{MetaDesc: ossHeadObjectTotalCounterDesc, valueType: prometheus.CounterValue},
				{MetaDesc: ossDeleteObjectTotalCounterDesc, valueType: prometheus.CounterValue},
				{MetaDesc: ossPostObjectTotalCounterDesc, valueType: prometheus.CounterValue},
			},
		},
		backendThroughputBytesCounterDesc: backendThroughputBytesCounterDesc{
			descs: []typedFactorDesc{
				{MetaDesc: backendReadBytesTotalCounterDesc, valueType: prometheus.CounterValue},
				{MetaDesc: backendWriteBytesTotalCounterDesc, valueType: prometheus.CounterValue},
			},
		},
		backendIOPSCompletedCounterDesc: backendIOPSCompletedCounterDesc{
			descs: []typedFactorDesc{
				{MetaDesc: backendReadCompletedTotalCounterDesc, valueType: prometheus.CounterValue},
				{MetaDesc: backendWriteCompletedTotalCounterDesc, valueType: prometheus.CounterValue},
			},
		},
		backendLatencyMillisecondsCounterDesc: backendLatencyMillisecondsCounterDesc{
			descs: []typedFactorDesc{
				{MetaDesc: backendReadTimeMillisecondsTotalCounterDesc, valueType: prometheus.CounterValue, factor: .001},
				{MetaDesc: backendWriteTimeMillisecondsTotalCounterDesc, valueType: prometheus.CounterValue, factor: .001},
			},
		},
		backendPosixCounterDesc: backendPosixCounterDesc{
			descs: []typedFactorDesc{
				{MetaDesc: backendPosixGetAttrTotalCounterDesc, valueType: prometheus.CounterValue},
				{MetaDesc: backendPosixGetModeTotalCounterDesc, valueType: prometheus.CounterValue},
				{MetaDesc: backendPosixAccessTotalCounterDesc, valueType: prometheus.CounterValue},
				{MetaDesc: backendPosixLookupTotalCounterDesc, valueType: prometheus.CounterValue},
				{MetaDesc: backendPosixMknodTotalCounterDesc, valueType: prometheus.CounterValue},
				{MetaDesc: backendPosixRemoveTotalCounterDesc, valueType: prometheus.CounterValue},
				{MetaDesc: backendPosixSetAttrTotalCounterDesc, valueType: prometheus.CounterValue},
				{MetaDesc: backendPosixLinkTotalCounterDesc, valueType: prometheus.CounterValue},
				{MetaDesc: backendPosixReadLinkTotalCounterDesc, valueType: prometheus.CounterValue},
				{MetaDesc: backendPosixStatfsTotalCounterDesc, valueType: prometheus.CounterValue},
				{MetaDesc: backendPosixRenameTotalCounterDesc, valueType: prometheus.CounterValue},
				{MetaDesc: backendPosixReaddirTotalCounterDesc, valueType: prometheus.CounterValue},
			},
		},
		mountRetryTotalCounter:         &typedFactorDesc{MetaDesc: mountRetryTotalCounterDesc, valueType: prometheus.CounterValue},
		mountPointStatus:               &typedFactorDesc{MetaDesc: mountPointStatusDesc, valueType: prometheus.GaugeValue},
		mountPointFailoverTotalCounter: &typedFactorDesc{MetaDesc: mountPointFailoverTotalCountDesc, valueType: prometheus.CounterValue},
		lastFuseClientExitReason:       &typedFactorDesc{MetaDesc: lastFuseClientExitReasonDesc, valueType: prometheus.GaugeValue},
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
			metrics = append(metrics, MustNewMetricWithMetaDesc(p.hotSpotReadFileTop, valueFloat64, prometheus.GaugeValue, labels...))
		case "hot_spot_write_file_top":
			metrics = append(metrics, MustNewMetricWithMetaDesc(p.hotSpotWriteFileTop, valueFloat64, prometheus.GaugeValue, labels...))
		case "hot_spot_head_file_top":
			metrics = append(metrics, MustNewMetricWithMetaDesc(p.hotSpotHeadFileTop, valueFloat64, prometheus.GaugeValue, labels...))
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
		ch <- prometheus.MustNewConstMetric(metric.Desc, metric.ValueType, metric.Value, convertLabelsToString(metric.VariableLabelPairs)...)
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
