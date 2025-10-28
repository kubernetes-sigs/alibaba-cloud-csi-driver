package server

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"k8s.io/klog/v2"
)

const (
	MetricsMountRetryCount          = "mount_retry_count"
	MetricsMountPointStatus         = "mount_point_status"
	MetricsMountPointFailoverCount  = "mount_point_failover_count"
	MetricsLastFuseClientExitReason = "last_fuse_client_exit_reason"
)

var maxCountRecord int = 999

func updateMountPointMetrics(metricsPath string, handleErr error) {
	{
		// update mount_retry_count
		file := filepath.Join(metricsPath, MetricsMountRetryCount)
		if handleErr != nil {
			if err := updateCounter(file); err != nil {
				klog.Errorf("Failed to update %s: %v", MetricsMountRetryCount, err)
			}
		}
	}
	{
		// update mount_point_status
		file := filepath.Join(metricsPath, MetricsMountPointStatus)
		var err error
		if handleErr != nil {
			err = utils.WriteAndSyncFile(file, []byte("1"), 0644)
		} else {
			err = utils.WriteAndSyncFile(file, []byte("0"), 0644)
		}
		if err != nil {
			klog.Errorf("Failed to update %s: %v", MetricsMountPointStatus, err)
		}
	}
	{
		// update last_fuse_client_exit_reason
		file := filepath.Join(metricsPath, MetricsLastFuseClientExitReason)
		currentTime := time.Now().Format(time.RFC3339)
		errorMessage := fmt.Sprintf("%s:: %v", currentTime, handleErr)
		if handleErr != nil {
			if err := utils.WriteAndSyncFile(file, []byte(errorMessage), 0644); err != nil {
				klog.Errorf("Failed to update %s: %v", MetricsLastFuseClientExitReason, err)
			}
		}
	}
}

func updateCounter(file string) (err error) {
	data, err := os.ReadFile(file)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("Failed to read file %s: %v", file, err)
	}
	var counter int
	if err != nil {
		// not exists
		counter = 0
	} else {
		if len(data) > 0 {
			counter, err = strconv.Atoi(string(data))
			if err != nil {
				return fmt.Errorf("Failed to parse counter from file %s: %v", file, err)
			}
		}
	}
	if counter < maxCountRecord {
		counter += 1
	}
	err = utils.WriteAndSyncFile(file, []byte(fmt.Sprintf("%d", counter)), 0644)
	if err != nil {
		return fmt.Errorf("Failed to write counter to file %s: %v", file, err)
	}
	return nil
}
