package server

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpdateMountPointMetrics(t *testing.T) {
	// Create temporary directory for testing
	tempDir := t.TempDir()

	tests := []struct {
		name      string
		handleErr error
		wantFiles map[string]string
	}{
		{
			name:      "success case - no error",
			handleErr: nil,
			wantFiles: map[string]string{
				utils.MetricsMountPointStatus: "0",
			},
		},
		{
			name:      "error case - with error",
			handleErr: fmt.Errorf("test error"),
			wantFiles: map[string]string{
				utils.MetricsMountPointStatus: "1",
			},
		},
	}

	mm := NewMountMonitorManager()
	m, _ := mm.GetMountMonitor("test-target", tempDir, nil, true)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Clean up any existing files
			for filename := range tt.wantFiles {
				file := filepath.Join(tempDir, filename)
				os.Remove(file)
			}

			// Call the function
			m.updateMountPointMetrics(nil, nil, tt.handleErr)

			// Check mount_point_status file
			statusFile := filepath.Join(tempDir, utils.MetricsMountPointStatus)
			content, err := os.ReadFile(statusFile)
			require.NoError(t, err)
			assert.Equal(t, tt.wantFiles[utils.MetricsMountPointStatus], string(content))

			// Check last_fuse_client_exit_reason file (only when there's an error)
			if tt.handleErr != nil {
				reasonFile := filepath.Join(tempDir, utils.MetricsLastFuseClientExitReason)
				content, err := os.ReadFile(reasonFile)
				require.NoError(t, err)
				reasonContent := string(content)
				assert.Contains(t, reasonContent, tt.handleErr.Error())
			}
		})
	}
}

func TestGetMountMonitor_CreatesMetricsPath(t *testing.T) {
	// Create temporary directory for testing
	tempDir := t.TempDir()

	// Test case 1: metricsPath doesn't exist, should be created
	nonExistentPath := filepath.Join(tempDir, "non-existent", "path")
	mm := NewMountMonitorManager()
	monitor, found := mm.GetMountMonitor("test-target-1", nonExistentPath, nil, true)
	require.NotNil(t, monitor, "Monitor should be created")
	assert.False(t, found, "Monitor should be newly created")

	// Verify directory was created
	assert.DirExists(t, nonExistentPath)

	// Test case 2: metricsPath == "", should not create monitor
	monitor2, found2 := mm.GetMountMonitor("test-target-2", "", nil, true)
	assert.Nil(t, monitor2, "Monitor should not be created when metricsPath is empty")
	assert.False(t, found2, "Monitor should not be found")

	// Test case 3: metricsPath exists, should work normally
	existingPath := filepath.Join(tempDir, "existing")
	err := os.MkdirAll(existingPath, 0755)
	require.NoError(t, err, "Failed to create existing directory")

	monitor3, found3 := mm.GetMountMonitor("test-target-3", existingPath, nil, true)
	require.NotNil(t, monitor3, "Monitor should be created")
	assert.False(t, found3, "Monitor should be newly created")
}
