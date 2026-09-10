package interceptors

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/server"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestOssfsMonitorInterceptor(t *testing.T) {
	metricsDir := t.TempDir()
	tests := []struct {
		name      string
		handler   mounter.MountHandler
		op        *mounter.MountOperation
		expectErr bool
	}{
		{
			name:    "nil operation",
			handler: successMountHandler,
		},
		{
			name:    "nil metrics path",
			handler: successMountHandler,
			op:      &mounter.MountOperation{},
		},
		{
			name:    "mount error reservation",
			handler: failureMountHandler,
			op: &mounter.MountOperation{
				MetricsPath: metricsDir,
			},
			expectErr: true,
		},
		{
			name:    "nil mount result",
			handler: successMountHandler,
			op: &mounter.MountOperation{
				MetricsPath: metricsDir,
				Target:      "target1",
			},
		},
		{
			name:    "invalid mount result",
			handler: successMountHandler,
			op: &mounter.MountOperation{
				MetricsPath: metricsDir,
				MountResult: "invalid",
				Target:      "target2",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := OssfsMonitorInterceptor(context.Background(), tt.op, tt.handler)
			if tt.expectErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			if tt.op == nil || tt.op.MetricsPath == "" {
				return
			}

			monitor, found := monitorManager.GetMountMonitor(tt.op.Target, tt.op.MetricsPath, raw, false)
			assert.True(t, found)
			assert.NotNil(t, monitor)

		})
	}

	monitorManager.StopAllMonitoring()
	monitorManager.WaitForAllMonitoring()
	monitorManager = server.NewMountMonitorManager()
	defer func() {
		monitorManager.StopAllMonitoring()
		monitorManager.WaitForAllMonitoring()
	}()

	op := &mounter.MountOperation{
		Target:      "volume1",
		MetricsPath: metricsDir,
		MountResult: server.OssfsMountResult{
			PID:      123,
			ExitChan: make(chan error),
		},
	}
	err := OssfsMonitorInterceptor(context.Background(), op, successMountHandler)
	assert.NoError(t, err)
	monitor, found := monitorManager.GetMountMonitor(op.Target, op.MetricsPath, raw, false)
	assert.True(t, found)
	assert.NotNil(t, monitor)
	assertMountMetricValue(t, op.MetricsPath, utils.MetricsMountPointStatus, "0")
	assertMountMetricValue(t, op.MetricsPath, utils.MetricsMountRetryCount, "0")

	err = OssfsMonitorInterceptor(context.Background(), op, failureMountHandler)
	assert.Error(t, err)
	assertMountMetricValue(t, op.MetricsPath, utils.MetricsMountPointStatus, "1")
	assertMountMetricValue(t, op.MetricsPath, utils.MetricsMountRetryCount, "1")
}

func assertMountMetricValue(t *testing.T, metricsDir, metricsFile string, expected string) {
	actual, err := os.ReadFile(filepath.Join(metricsDir, metricsFile))
	assert.NoError(t, err)
	assert.Equal(t, expected, string(actual))
}

func TestRecoveryMetrics(t *testing.T) {
	monitorManager.StopAllMonitoring()
	monitorManager.WaitForAllMonitoring()
	monitorManager = server.NewMountMonitorManager()
	defer func() {
		monitorManager.StopAllMonitoring()
		monitorManager.WaitForAllMonitoring()
	}()

	metricsDir := t.TempDir()
	op := &mounter.MountOperation{
		Target:      "recovery-test-volume",
		MetricsPath: metricsDir,
		MountResult: server.OssfsMountResult{
			PID:      123,
			ExitChan: make(chan error),
		},
	}

	// Initial mount success
	err := OssfsMonitorInterceptor(context.Background(), op, successMountHandler)
	assert.NoError(t, err)

	monitor, found := monitorManager.GetMountMonitor(op.Target, op.MetricsPath, raw, false)
	assert.True(t, found)
	assert.NotNil(t, monitor)

	// Verify initial metrics
	assertMountMetricValue(t, metricsDir, utils.MetricsMountPointStatus, "0")
	assertMountMetricValue(t, metricsDir, utils.MetricsMountRetryCount, "0")

	// Verify callbacks are registered on op
	assert.NotNil(t, op.OnProcessExit, "OnProcessExit callback should be set")
	assert.NotNil(t, op.OnRecoverySuccess, "OnRecoverySuccess callback should be set")
	assert.NotNil(t, op.OnRecoveryFailed, "OnRecoveryFailed callback should be set")

	// Simulate first recovery exit
	exitErr1 := fmt.Errorf("ossfs2 exited with error: connection reset")
	op.OnProcessExit(exitErr1)

	// Verify failover count incremented and exit reason updated
	assertMountMetricValue(t, metricsDir, utils.MetricsMountPointFailoverCount, "1")
	assertMountMetricValue(t, metricsDir, utils.MetricsMountPointStatus, "1") // unhealthy during restart
	exitReason, err := os.ReadFile(filepath.Join(metricsDir, utils.MetricsLastFuseClientExitReason))
	assert.NoError(t, err)
	assert.Contains(t, string(exitReason), "connection reset")

	// Simulate recovery success after 2 attempts
	op.OnRecoverySuccess(456, exitErr1, 2)

	// Verify status back to healthy
	assertMountMetricValue(t, metricsDir, utils.MetricsMountPointStatus, "0")
	// Failover count should remain at 1
	assertMountMetricValue(t, metricsDir, utils.MetricsMountPointFailoverCount, "1")
	// Exit reason should contain original error and recovery context
	exitReason, err = os.ReadFile(filepath.Join(metricsDir, utils.MetricsLastFuseClientExitReason))
	assert.NoError(t, err)
	assert.Contains(t, string(exitReason), "connection reset")
	assert.Contains(t, string(exitReason), "recovered after 2 attempt(s)")

	// Simulate second recovery exit
	exitErr2 := fmt.Errorf("ossfs2 exited with error: timeout")
	op.OnProcessExit(exitErr2)

	// Verify failover count incremented to 2
	assertMountMetricValue(t, metricsDir, utils.MetricsMountPointFailoverCount, "2")
	exitReason, err = os.ReadFile(filepath.Join(metricsDir, utils.MetricsLastFuseClientExitReason))
	assert.NoError(t, err)
	assert.Contains(t, string(exitReason), "timeout")

	// Simulate recovery failure
	recoveryErr := fmt.Errorf("credential expired")
	op.OnRecoveryFailed(exitErr2, recoveryErr, 5)

	// Verify status stays unhealthy, exit reason shows recovery failure context
	assertMountMetricValue(t, metricsDir, utils.MetricsMountPointStatus, "1")
	exitReason, err = os.ReadFile(filepath.Join(metricsDir, utils.MetricsLastFuseClientExitReason))
	assert.NoError(t, err)
	assert.Contains(t, string(exitReason), "timeout")
	assert.Contains(t, string(exitReason), "recovery exhausted after 5 attempt(s)")
	assert.Contains(t, string(exitReason), "credential expired")
}
