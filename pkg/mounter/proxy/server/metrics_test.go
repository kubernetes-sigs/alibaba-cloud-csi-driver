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
	tempDir, err := os.MkdirTemp("", "metrics_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir) // Clean up

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
