package server

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
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
				MetricsMountPointStatus: "0",
			},
		},
		{
			name:      "error case - with error",
			handleErr: fmt.Errorf("test error"),
			wantFiles: map[string]string{
				MetricsMountPointStatus: "1",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Clean up any existing files
			for filename := range tt.wantFiles {
				file := filepath.Join(tempDir, filename)
				os.Remove(file)
			}

			// Call the function
			updateMountPointMetrics(tempDir, tt.handleErr)

			// Check mount_point_status file
			statusFile := filepath.Join(tempDir, MetricsMountPointStatus)
			content, err := os.ReadFile(statusFile)
			require.NoError(t, err)
			assert.Equal(t, tt.wantFiles[MetricsMountPointStatus], string(content))

			// Check last_fuse_client_exit_reason file (only when there's an error)
			if tt.handleErr != nil {
				reasonFile := filepath.Join(tempDir, MetricsLastFuseClientExitReason)
				content, err := os.ReadFile(reasonFile)
				require.NoError(t, err)
				reasonContent := string(content)
				assert.Contains(t, reasonContent, tt.handleErr.Error())
			}

			// Check mount_retry_count file (only when there's an error)
			if tt.handleErr != nil {
				retryFile := filepath.Join(tempDir, MetricsMountRetryCount)
				content, err := os.ReadFile(retryFile)
				require.NoError(t, err)
				count, err := strconv.Atoi(string(content))
				require.NoError(t, err)
				// When file doesn't exist initially, updateCounter sets counter to 0
				// When file exists, it increments the counter
				// For first error, we expect 1 since file doesn't exist initially
				assert.Equal(t, 1, count)
			}
		})
	}
}

func TestUpdateCounter(t *testing.T) {
	// Create temporary directory for testing
	tempDir, err := os.MkdirTemp("", "counter_test")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir) // Clean up

	tests := []struct {
		name          string
		initialValue  string
		expectedValue int
		expectError   bool
		setupFile     bool
	}{
		{
			name:          "new file - should start at 1",
			setupFile:     false,
			expectedValue: 1,
			expectError:   false,
		},
		{
			name:          "existing file with value 5 - should increment to 6",
			setupFile:     true,
			initialValue:  "5",
			expectedValue: 6,
			expectError:   false,
		},
		{
			name:          "existing file with value 0 - should increment to 1",
			setupFile:     true,
			initialValue:  "0",
			expectedValue: 1,
			expectError:   false,
		},
		{
			name:          "existing file with invalid value - should return error",
			setupFile:     true,
			initialValue:  "invalid",
			expectedValue: 0,
			expectError:   true,
		},
		{
			name:          "existing empty file - should start at 1",
			setupFile:     true,
			initialValue:  "",
			expectedValue: 1,
			expectError:   false,
		},
		{
			name:          "reach max record value",
			setupFile:     true,
			initialValue:  strconv.Itoa(maxCountRecord),
			expectedValue: maxCountRecord,
			expectError:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			counterFile := filepath.Join(tempDir, "counter_test")

			// Setup initial file if needed
			if tt.setupFile {
				err := utils.WriteAndSyncFile(counterFile, []byte(tt.initialValue), 0644)
				require.NoError(t, err)
			}

			// Call updateCounter
			err = updateCounter(counterFile)

			// Check error expectation
			require.Equal(t, tt.expectError, err != nil)

			// Check counter value if no error expected
			if !tt.expectError {
				content, err := os.ReadFile(counterFile)
				require.NoError(t, err)
				actualValue, err := strconv.Atoi(string(content))
				require.NoError(t, err)
				assert.Equal(t, tt.expectedValue, actualValue)
			}

			// Clean up for next test
			os.Remove(counterFile)
		})
	}
}
