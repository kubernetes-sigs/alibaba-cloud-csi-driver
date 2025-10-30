package server

import (
	"context"
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
	"k8s.io/utils/ptr"
)

const (
	MetricsMountRetryCount          = "mount_retry_count"
	MetricsMountPointStatus         = "mount_point_status"
	MetricsMountPointFailoverCount  = "mount_point_failover_count"
	MetricsLastFuseClientExitReason = "last_fuse_client_exit_reason"
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
	Target        string
	Pid           int
	MetricsPath   string
	State         MonitorState
	mu            sync.RWMutex
	stopCh        chan struct{}
	monitorCtx    context.Context
	monitorCancel context.CancelFunc
	raw           mount.Interface
	// Mount retry count (persistent across mount attempts)
	retryCount    int
	failoverCount int
}

// MountMonitorManager manages multiple mount monitors
type MountMonitorManager struct {
	monitors *sync.Map // map[string]*MountMonitor - key is target path
	wg       sync.WaitGroup
}

// NewMountMonitorManager creates a new mount monitor manager
func NewMountMonitorManager() *MountMonitorManager {
	return &MountMonitorManager{
		monitors: new(sync.Map),
	}
}

// GetMountMonitor creates a new monitor or returns existing one
func (manager *MountMonitorManager) GetMountMonitor(target, metricsPath string, raw mount.Interface, createIfNotExist bool) *MountMonitor {
	// Check if monitor already exists
	if existingMonitor, exists := manager.monitors.Load(target); exists {
		monitor := existingMonitor.(*MountMonitor)
		// Not allow to update info like metrics path currently
		klog.V(4).InfoS("Using existing monitor", "target", target)
		return monitor
	}
	if !createIfNotExist {
		return nil
	}

	// Create new monitor
	monitorCtx, monitorCancel := context.WithCancel(context.Background())
	monitor := &MountMonitor{
		Target:        target,
		MetricsPath:   metricsPath,
		State:         MonitorStateInitialized,
		stopCh:        make(chan struct{}),
		monitorCtx:    monitorCtx,
		monitorCancel: monitorCancel,
		raw:           raw,
		retryCount:    0,
		failoverCount: 0,
	}

	manager.monitors.Store(target, monitor)

	// Initialize metrics files to empty/zero state
	monitor.updateMountPointMetrics(&monitor.retryCount, &monitor.failoverCount, nil, nil)

	klog.InfoS("Created new monitor", "target", target)
	return monitor
}

// HandleMountResult handles mount result (success or failure)
func (m *MountMonitor) HandleMountResult(process *exec.Cmd, err error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if err != nil {
		// Mount failed
		if m.retryCount < maxCountRecord {
			m.retryCount++
		}
		m.updateMountPointMetrics(&m.retryCount, nil, ptr.To(true), err)

		m.State = MonitorStateInitialized // Back to initialized for retry

		klog.V(2).InfoS("Mount failed, updated retry count",
			"target", m.Target,
			"retry_count", m.retryCount,
			"error", err)
	} else {
		// Mount successful
		// Update mount_point_status to 0 (healthy)
		m.updateMountPointMetrics(nil, nil, ptr.To(false), nil)

		m.State = MonitorStateMonitoring
		m.Pid = process.Process.Pid
	}
}

// Stop stops the monitoring
func (m *MountMonitor) Stop() {
	close(m.stopCh)
	m.monitorCancel()
}

// WaitForAllMonitoring waits for all monitoring goroutines to finish
func (manager *MountMonitorManager) WaitForAllMonitoring() {
	manager.wg.Wait()
}

// checkMountPointStatus checks the current status of a mount point
func (m *MountMonitor) checkMountPointStatus(target string) (err error) {
	// Check if it's a mount point
	notMnt, err := m.raw.IsLikelyNotMountPoint(target)
	if err != nil {
		// case 1: mount point does not exist
		if os.IsNotExist(err) {
			return fmt.Errorf("Mountpoint does not exist")
		}
		// case 2: mount point is corrupted
		if mountutils.IsCorruptedMnt(err) {
			return fmt.Errorf("Mountpoint is corrupted")
		}
		// case 3: failed to check mount point
		return fmt.Errorf("Error checking mount point, error: %v", err)
	}

	// case 4: mount point is not mounted
	if notMnt {
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
	old := m.readMetricsFiles(MetricsMountPointStatus)
	// metrics file stores notHealthy as "1", healthy as "0"
	oldUnhealthy := old == "1"
	// Check mount point status
	err := m.checkMountPointStatus(m.Target)
	if err != nil {
		klog.ErrorS(err, "Check mount point status with error", "target", m.Target)
	} else {
		klog.V(2).InfoS("Check mount point status successfully", "target", m.Target)
	}
	newUnhealthy := err != nil
	if newUnhealthy != oldUnhealthy {
		klog.Warningf("Mount point status changed. but the detailed reason was missing")
		// Update status
		m.updateMountPointMetrics(nil, nil, ptr.To(newUnhealthy), err)
	}
}

func boolToBinaryString(b bool) string {
	boolMap := map[bool]string{true: "1", false: "0"}
	return boolMap[b]
}

// initializeMetrics initializes metrics files to empty/zero state
func (m *MountMonitor) updateMountPointMetrics(
	retryCont *int,
	failoverCont *int,
	notHealthy *bool,
	lastExistError error) {
	if retryCont != nil {
		// Update mount_retry_count
		retryFile := filepath.Join(m.MetricsPath, MetricsMountRetryCount)
		if err := utils.WriteAndSyncFile(retryFile, []byte(strconv.Itoa(*retryCont)), 0644); err != nil {
			klog.Errorf("Failed to update %s: %v", MetricsMountRetryCount, err)
		}
		klog.V(4).Infof("Update %s: %d", MetricsMountRetryCount, *retryCont)
	}
	if failoverCont != nil {
		// Update mount_point_failover_count
		failoverFile := filepath.Join(m.MetricsPath, MetricsMountPointFailoverCount)
		if err := utils.WriteAndSyncFile(failoverFile, []byte(strconv.Itoa(*failoverCont)), 0644); err != nil {
			klog.Errorf("Failed to update %s: %v", MetricsMountPointFailoverCount, err)
		}
		klog.V(4).Infof("Update %s: %d", MetricsMountPointFailoverCount, *failoverCont)
	}
	if notHealthy != nil {
		// Update mount_point_status
		statusFile := filepath.Join(m.MetricsPath, MetricsMountPointStatus)
		if err := utils.WriteAndSyncFile(statusFile, []byte(boolToBinaryString(*notHealthy)), 0644); err != nil {
			klog.Errorf("Failed to update %s: %v", MetricsMountPointStatus, err)
		}
		klog.V(4).Infof("Update %s: %t", MetricsMountPointStatus, *notHealthy)
	}
	if lastExistError != nil {
		// Update last_fuse_client_exit_reason
		statusFile := filepath.Join(m.MetricsPath, MetricsLastFuseClientExitReason)
		currentTime := time.Now().Format(time.RFC3339)
		errorMessage := fmt.Sprintf("%s:: %s", currentTime, lastExistError.Error())
		if err := utils.WriteAndSyncFile(statusFile, []byte(errorMessage), 0644); err != nil {
			klog.Errorf("Failed to update %s: %v", MetricsLastFuseClientExitReason, err)
		}
		klog.V(4).Infof("Update %s: %s", MetricsLastFuseClientExitReason, errorMessage)
	}

}

func (m *MountMonitor) readMetricsFiles(metrics string) string {
	switch metrics {
	case MetricsMountPointFailoverCount, MetricsLastFuseClientExitReason, MetricsMountPointStatus, MetricsMountRetryCount:
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
	manager.wg.Add(1)
	monitor := manager.GetMountMonitor(target, "", nil, false)
	if monitor == nil {
		// No monitor found; nothing to start
		manager.wg.Done()
		klog.Errorf("StartMonitoring called but monitor not found: target=%s", target)
		return
	}
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
			case <-monitor.monitorCtx.Done():
				klog.InfoS("Mount monitoring context cancelled", "target", monitor.Target)
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
