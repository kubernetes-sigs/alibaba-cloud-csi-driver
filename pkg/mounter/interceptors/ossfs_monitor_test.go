package interceptors

import (
	"context"
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
