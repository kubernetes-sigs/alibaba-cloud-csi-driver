package utils

import (
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"golang.org/x/sys/unix"
	"k8s.io/klog/v2"
)

// ----------------DISK METRICS-----------------

// GetMetrics get path metric
func GetMetrics(path string) (*csi.NodeGetVolumeStatsResponse, error) {
	if path == "" {
		return nil, fmt.Errorf("getMetrics No path given")
	}
	statfs := &unix.Statfs_t{}
	err := unix.Statfs(path, statfs)
	if err != nil {
		return nil, err
	}

	// Available is blocks available * fragment size
	available := int64(statfs.Bavail) * int64(statfs.Bsize)

	// Capacity is total block count * fragment size
	capacity := int64(statfs.Blocks) * int64(statfs.Bsize)

	// Usage is block being used * fragment size (aka block size).
	usage := (int64(statfs.Blocks) - int64(statfs.Bfree)) * int64(statfs.Bsize)

	inodes := int64(statfs.Files)
	inodesFree := int64(statfs.Ffree)
	inodesUsed := inodes - inodesFree

	return &csi.NodeGetVolumeStatsResponse{
		Usage: []*csi.VolumeUsage{
			{
				Available: available,
				Total:     capacity,
				Used:      usage,
				Unit:      csi.VolumeUsage_BYTES,
			},
			{
				Available: inodesFree,
				Total:     inodes,
				Used:      inodesUsed,
				Unit:      csi.VolumeUsage_INODES,
			},
		},
	}, nil
}

// ----------------FUSE METRICS-----------------

const (
	PodInfoFile        = "pod_info"
	MountPointInfoFile = "mount_point_info"
)

const (
	MetricsMountRetryCount          = "mount_retry_count"
	MetricsMountPointStatus         = "mount_point_status"
	MetricsMountPointFailoverCount  = "mount_point_failover_count"
	MetricsLastFuseClientExitReason = "last_fuse_client_exit_reason"

	MetricsHotSpotReadFileTop  = "hot_spot_read_file_top"
	MetricsHotSpotWriteFileTop = "hot_spot_write_file_top"
	MetricsHotSpotHeadFileTop  = "hot_spot_head_file_top"

	MetricsCapacityCounter   = "capacity_counter"
	MetricsInodesCounter     = "inodes_counter"
	MetricsThroughputCounter = "throughput_counter"
	MetricsIopsCounter       = "iops_counter"
	MetricsLatencyCounter    = "latency_counter"
	MetricsPosixCounter      = "posix_counter"
	MetricsOssObjectCounter  = "oss_object_counter"
)

var MountpointMetricsArray = []string{
	MetricsMountRetryCount,
	MetricsMountPointStatus,
	MetricsMountPointFailoverCount,
	MetricsLastFuseClientExitReason,
}

var CounterTypeMetricsArray = []string{
	MetricsCapacityCounter,
	MetricsInodesCounter,
	MetricsThroughputCounter,
	MetricsIopsCounter,
	MetricsLatencyCounter,
	MetricsPosixCounter,
	MetricsOssObjectCounter,
}

var HotSpotMetricsArray = []string{
	MetricsHotSpotReadFileTop,
	MetricsHotSpotWriteFileTop,
	MetricsHotSpotHeadFileTop,
}

func GetFuseMetricsMountDir(metricsPathPrefix, volumeId string) string {
	return filepath.Join(metricsPathPrefix, fmt.Sprintf("%x", sha256.Sum256([]byte(volumeId))))
}

func RemoveMetrics(
	metricsPathPrefix string,
	req *csi.NodeUnstageVolumeRequest,
) {
	mountPointPath := GetFuseMetricsMountDir(metricsPathPrefix, req.VolumeId)

	// Arrays of metrics files to remove
	metricsArrays := [][]string{
		MountpointMetricsArray,
		CounterTypeMetricsArray,
		HotSpotMetricsArray,
		{MountPointInfoFile},
	}

	// Remove all metrics files from the three arrays
	for _, metricsArray := range metricsArrays {
		for _, filename := range metricsArray {
			filePath := filepath.Join(mountPointPath, filename)
			if err := os.Remove(filePath); err != nil {
				if os.IsNotExist(err) {
					// File doesn't exist, skip silently
					continue
				}
				// File exists but cannot be removed, log error
				klog.ErrorS(err, "Failed to remove metrics file", "file", filePath, "volumeId", req.VolumeId)
			}
		}
	}

	// Remove the mountPointPath directory itself (only if empty)
	// Use os.Remove instead of os.RemoveAll to prevent deletion if there are other files
	if err := os.Remove(mountPointPath); err != nil {
		if os.IsNotExist(err) {
			// Directory doesn't exist, skip silently
			return
		}
		// Directory is not empty or cannot be removed, skip silently
		// This is expected behavior when there are other files in the directory
		klog.ErrorS(err, "Metrics directory not removed", "path", mountPointPath, "volumeId", req.VolumeId)
	} else {
		klog.V(4).InfoS("Removed metrics directory", "path", mountPointPath, "volumeId", req.VolumeId)
	}
}

func WriteSharedMetricsInfo(metricsPathPrefix string,
	req *csi.NodePublishVolumeRequest,
	clientName string, storageBackendName string, fsName string, sharedPath string) (
	mountPointPath string,
) {
	mountPointPath = GetFuseMetricsMountDir(metricsPathPrefix, req.GetVolumeId())
	mountPointName := MountPointInfoFile
	if !IsFileExisting(mountPointPath) {
		_ = os.MkdirAll(mountPointPath, os.FileMode(0755))
	}
	if !IsFileExisting(filepath.Join(mountPointPath, mountPointName)) {
		info := clientName + " " +
			storageBackendName + " " +
			fsName + " " +
			req.GetVolumeId() + " " +
			sharedPath
		_ = WriteAndSyncFile(filepath.Join(mountPointPath, mountPointName), []byte(info), os.FileMode(0644))
	}
	return strings.TrimPrefix(mountPointPath, hostPrefix)
}

func WriteMetricsInfo(metricsPathPrefix string,
	req *csi.NodePublishVolumeRequest,
	metricsTop string, clientName string, storageBackendName string, fsName string) (
	mountPointPath string,
) {
	podUIDPath := metricsPathPrefix + req.VolumeContext["csi.storage.k8s.io/pod.uid"] + "/"
	mountPointPath = podUIDPath + req.GetVolumeId() + "/"
	if !IsFileExisting(mountPointPath) {
		_ = os.MkdirAll(mountPointPath, os.FileMode(0755))
	}
	if !IsFileExisting(podUIDPath + PodInfoFile) {
		info := req.VolumeContext["csi.storage.k8s.io/pod.namespace"] + " " +
			req.VolumeContext["csi.storage.k8s.io/pod.name"] + " " +
			req.VolumeContext["csi.storage.k8s.io/pod.uid"] + " " +
			metricsTop
		_ = WriteAndSyncFile(podUIDPath+PodInfoFile, []byte(info), os.FileMode(0644))
	}

	if !IsFileExisting(mountPointPath + MountPointInfoFile) {
		info := clientName + " " +
			storageBackendName + " " +
			fsName + " " +
			req.GetVolumeId() + " " +
			req.TargetPath
		_ = WriteAndSyncFile(mountPointPath+MountPointInfoFile, []byte(info), os.FileMode(0644))
	}
	return strings.TrimPrefix(mountPointPath, hostPrefix)
}
