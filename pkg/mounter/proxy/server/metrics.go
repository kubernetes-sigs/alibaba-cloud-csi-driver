package server

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"k8s.io/klog/v2"
	"k8s.io/mount-utils"
	mountutils "k8s.io/mount-utils"
)

var maxCountRecord int = 999

// MonitorState represents the state of a monitor
type MonitorState int

const (
	MonitorStateInitialized MonitorState = iota
	MonitorStateMonitoring
)

// MountMonitor manages monitoring for a single mount point
type MountMonitor struct {
	Target string
	Pid    int
	// Note: MetricsPath is the host path volume mount point,
	// recycled by CSI nodeUnstageVolume
	MetricsPath string
	State       MonitorState
	mu          sync.RWMutex
	stopCh      chan struct{}
	raw         mount.Interface
	// Mount retry count (persistent across mount attempts)
	retryCount    int
	failoverCount int
}

// MountMonitorManager manages multiple mount monitors
type MountMonitorManager struct {
	// map[string]*MountMonitor - key is target path
	// Monitors Lifecycle:
	// - Creation: Monitors are created during the initial mount operation.
	// - Termination: All monitors will be terminated together when the mount-proxy lifecycle ends.
	// TODO: Need to handle the compatibility for cases where the lifecycle of mount-proxy
	// (e.g., efc) and fuse client are not consistent.
	monitors sync.Map
	wg       sync.WaitGroup
}

// NewMountMonitorManager creates a new mount monitor manager
func NewMountMonitorManager() *MountMonitorManager {
	return &MountMonitorManager{}
}

// GetMountMonitor creates a new monitor or returns existing one
func (manager *MountMonitorManager) GetMountMonitor(
	target, metricsPath string,
	raw mount.Interface,
	createIfNotExist bool,
) (monitor *MountMonitor, found bool) {
	// Check if monitor already exists
	if existingMonitor, exists := manager.monitors.Load(target); exists {
		monitor := existingMonitor.(*MountMonitor)
		// Not allow to update info like metrics path currently
		klog.V(4).InfoS("Using existing monitor", "target", target)
		return monitor, true
	}
	if !createIfNotExist {
		return nil, false
	}

	if metricsPath == "" {
		klog.V(4).InfoS("Skipping mount point monitoring for target %s: no metrics path provided", target)
		return nil, false
	}

	// Ensure metricsPath directory exists for rund (agent mode)
	if err := os.MkdirAll(metricsPath, 0755); err != nil {
		klog.ErrorS(err, "Failed to create metrics path directory", "path", metricsPath, "target", target)
		return nil, false
	}

	// Create new monitor
	monitor = &MountMonitor{
		Target:        target,
		MetricsPath:   metricsPath,
		State:         MonitorStateInitialized,
		stopCh:        make(chan struct{}),
		raw:           raw,
		retryCount:    0,
		failoverCount: 0,
	}

	manager.monitors.Store(target, monitor)

	klog.InfoS("Created new monitor", "target", target)
	return monitor, false
}

func (m *MountMonitor) IncreaseMountRetryCount() {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.retryCount < maxCountRecord {
		m.retryCount++
	}
}

// HandleMountFailureOrExit handles the case when mount operation fails
// or fuse client exits.
func (m *MountMonitor) HandleMountFailureOrExit(err error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.State = MonitorStateInitialized // Back to initialized state for retry
	if err == nil {
		// should not set the status to healthy
		klog.Warning("Process exit without error", "pid", m.Pid, "target", m.Target)
		return
	}
	// Update metrics for mount failure
	m.updateMountPointMetrics(&m.retryCount, nil, err)

	klog.ErrorS(err, "Mount failed, updated retry count",
		"target", m.Target,
		"retry_count", m.retryCount)
}

// HandleMountSuccess handles the case when mount operation succeeds
func (m *MountMonitor) HandleMountSuccess(process *exec.Cmd) {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Update metrics for mount success
	m.updateMountPointMetrics(&m.retryCount, nil, nil)

	m.State = MonitorStateMonitoring
	m.Pid = process.Process.Pid

	klog.InfoS("Mount succeeded", "target", m.Target, "pid", m.Pid)
}

// Stop stops the monitoring and cleans up metrics files
func (m *MountMonitor) Stop() {
	// Close stopCh first to signal monitoring goroutine to stop
	close(m.stopCh)

	// MetricsPath is a hostPath mount point in the pod, we should only remove files inside,
	// not the directory itself, CSI nodeUnstageVolume will clean up the directory.
	if m.MetricsPath != "" {
		// Remove only the metrics files we created, not the entire directory
		for _, filename := range utils.MountpointMetricsArray {
			filePath := filepath.Join(m.MetricsPath, filename)
			if err := os.Remove(filePath); err != nil && !os.IsNotExist(err) {
				klog.ErrorS(err, "Failed to remove metrics file", "target", m.Target, "file", filePath)
			}
		}

		klog.V(4).InfoS("Cleaned up metrics files", "target", m.Target, "path", m.MetricsPath)
	}
}

// WaitForAllMonitoring waits for all monitoring goroutines to finish with timeout
func (manager *MountMonitorManager) WaitForAllMonitoring() {
	const timeout = 2 * time.Second
	done := make(chan struct{})

	go func() {
		manager.wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		klog.V(4).InfoS("All monitoring goroutines finished")
	case <-time.After(timeout):
		klog.Warningf("WaitForAllMonitoring timed out after %v, some goroutines may still be running", timeout)
	}
}

// checkMountPointStatus checks the current status of a mount point
// Note:
// When [strict] is set to true, any abnormality of the mount point will result in an error.
// This may cause continuous polling failures if the mount point is manually unmounted but the Pod is not deleted.
// If there is no specific requirement to manually unmount the mount point, [strict] should be set to false
// to avoid unnecessary errors during monitoring.
func (m *MountMonitor) checkMountPointStatus(target string, strict bool) (err error) {
	// Check if it's a mount point
	notMnt, err := m.raw.IsLikelyNotMountPoint(target)
	if err != nil {
		// case 1: mount point does not exist
		if os.IsNotExist(err) {
			if strict {
				return fmt.Errorf("Mount point does not exist")
			}
			return nil
		}
		// case 2: mount point is corrupted
		if mountutils.IsCorruptedMnt(err) {
			return fmt.Errorf("Mountpoint is corrupted")
		}
		// case 3: failed to check mount point
		return fmt.Errorf("Error checking mount point, error: %v", err)
	}

	// case 4: mount point is not mounted
	if notMnt && strict {
		return fmt.Errorf("Mountpoint is not mounted")
	}
	return nil
}

// checkAndUpdateMountStatus checks mount status and updates metrics
func (m *MountMonitor) checkAndUpdateMountStatus() {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Only check if we're in monitoring state
	if m.State != MonitorStateMonitoring {
		return
	}
	old := m.readMetricsFiles(utils.MetricsMountPointStatus)
	// metrics file stores notHealthy as "1", healthy as "0"
	oldUnhealthy := old == "1"
	// Check mount point status
	err := m.checkMountPointStatus(m.Target, false)
	if err != nil {
		klog.ErrorS(err, "Check mount point status with error", "target", m.Target)
	} else {
		klog.V(5).InfoS("Check mount point status successfully", "target", m.Target)
	}
	newUnhealthy := err != nil
	if newUnhealthy != oldUnhealthy {
		klog.Warningf("Mount point status changed. but the detailed reason was missing")
		// Update status
		m.updateMountPointMetrics(nil, nil, err)
	}
}

func boolToBinaryString(b bool) string {
	if b {
		return "1"
	}
	return "0"
}

// initializeMetrics initializes metrics files to empty/zero state
func (m *MountMonitor) updateMountPointMetrics(
	retryCont *int,
	failoverCont *int,
	lastExistError error) {
	if retryCont != nil {
		// Update mount_retry_count
		retryFile := filepath.Join(m.MetricsPath, utils.MetricsMountRetryCount)
		if err := os.WriteFile(retryFile, []byte(strconv.Itoa(*retryCont)), 0644); err != nil {
			klog.ErrorS(err, "Failed to update metrics", "key", utils.MetricsMountRetryCount, "val", *retryCont)
		}
		klog.V(5).Infof("Update %s: %d", utils.MetricsMountRetryCount, *retryCont)
	}
	if failoverCont != nil {
		// Update mount_point_failover_count
		failoverFile := filepath.Join(m.MetricsPath, utils.MetricsMountPointFailoverCount)
		if err := os.WriteFile(failoverFile, []byte(strconv.Itoa(*failoverCont)), 0644); err != nil {
			klog.ErrorS(err, "Failed to update metrics", "key", utils.MetricsMountPointFailoverCount, "val", *failoverCont)
		}
		klog.V(5).Infof("Update %s: %d", utils.MetricsMountPointFailoverCount, *failoverCont)
	}
	if lastExistError != nil {
		// Update last_fuse_client_exit_reason
		statusFile := filepath.Join(m.MetricsPath, utils.MetricsLastFuseClientExitReason)
		currentTime := time.Now().Format(time.RFC3339)
		errorMessage := fmt.Sprintf("%s:: %s", currentTime, lastExistError.Error())
		if err := os.WriteFile(statusFile, []byte(errorMessage), 0644); err != nil {
			klog.ErrorS(err, "Failed to update metrics", "key", utils.MetricsLastFuseClientExitReason, "val", errorMessage)
		}
		klog.V(5).Infof("Update %s: %s", utils.MetricsLastFuseClientExitReason, errorMessage)
	}
	{
		notHealthy := lastExistError != nil
		// Update mount_point_status
		statusFile := filepath.Join(m.MetricsPath, utils.MetricsMountPointStatus)
		if err := os.WriteFile(statusFile, []byte(boolToBinaryString(notHealthy)), 0644); err != nil {
			klog.ErrorS(err, "Failed to update metrics", "key", utils.MetricsMountPointStatus, "val", boolToBinaryString(notHealthy))
		}
		klog.V(5).Infof("Update %s: %t", utils.MetricsMountPointStatus, notHealthy)
	}
}

func (m *MountMonitor) readMetricsFiles(metrics string) string {
	switch metrics {
	case utils.MetricsMountPointFailoverCount, utils.MetricsLastFuseClientExitReason, utils.MetricsMountPointStatus, utils.MetricsMountRetryCount:
		file := filepath.Join(m.MetricsPath, metrics)
		content, err := os.ReadFile(file)
		if err != nil && !os.IsNotExist(err) {
			klog.ErrorS(err, "Failed to read metrics file", "file", file)
			return ""
		}
		return string(content)
	default:
		return ""
	}
}

// StartMonitoring starts the monitoring goroutine for this monitor
func (manager *MountMonitorManager) StartMonitoring(target string) {
	monitor, _ := manager.GetMountMonitor(target, "", nil, false)
	if monitor == nil {
		// No monitor found; nothing to start
		klog.ErrorS(nil, "StartMonitoring failed, monitor not found", "target", target)
		return
	}
	manager.wg.Add(1)
	klog.InfoS("Mount successful, started monitoring", "target", monitor.Target, "pid", monitor.Pid)
	go func() {
		defer manager.wg.Done()

		ticker := time.NewTicker(5 * time.Second) // Check every 5 seconds
		defer ticker.Stop()

		for {
			select {
			case <-monitor.stopCh:
				klog.InfoS("Stopping mount monitoring", "target", monitor.Target)
				return
			case <-ticker.C:
				monitor.checkAndUpdateMountStatus()
			}
		}
	}()
}

// StopAllMonitoring stops all mount monitoring
func (manager *MountMonitorManager) StopAllMonitoring() {
	manager.monitors.Range(func(key, value any) bool {
		target := key.(string)
		monitor := value.(*MountMonitor)
		klog.InfoS("Stopping mount monitoring", "target", target)
		monitor.Stop()
		return true
	})
}
