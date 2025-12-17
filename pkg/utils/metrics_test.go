package utils

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetFuseMetricsMountDir(t *testing.T) {
	metricsPrefix := "/test/metrics"
	tests := []struct {
		name     string
		volumeId string
		want     string
	}{
		{
			"empty id",
			"",
			"/test/metrics/e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		},
		{
			"normal volumeid",
			"pv-oss",
			"/test/metrics/8f3e75e1af90a7dcc66882ec1544cb5c7c32c82c2b56b25a821ac77cea60a928",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := GetFuseMetricsMountDir(metricsPrefix, tt.volumeId)
			assert.Equal(t, tt.want, actual)
		})
	}
}

func TestWriteSharedMetricsInfo(t *testing.T) {
	// Create temporary directory
	tmpDir, err := os.MkdirTemp("", "metrics-test-*")
	require.NoError(t, err)
	defer os.RemoveAll(tmpDir)

	// Create metricsPathPrefix in tmpDir with hostPrefix prefix to test TrimPrefix logic
	metricsPathPrefix := filepath.Join(tmpDir, hostPrefix, "metrics")
	err = os.MkdirAll(metricsPathPrefix, 0755)
	require.NoError(t, err)

	// Create test request
	req := &csi.NodePublishVolumeRequest{
		VolumeId:   "test-volume-123",
		TargetPath: "/mnt/test",
	}

	clientName := "ossfs"
	storageBackendName := "oss"
	fsName := "test-bucket"
	sharedPath := "/shared/path"

	// Test: First call should create directory and file
	mountPointPath := WriteSharedMetricsInfo(metricsPathPrefix, req, clientName, storageBackendName, fsName, sharedPath)

	// Verify directory was created
	expectedDir := GetFuseMetricsMountDir(metricsPathPrefix, req.GetVolumeId())
	assert.True(t, IsFileExisting(expectedDir))

	// Verify file was created with correct content
	mountPointInfoPath := filepath.Join(expectedDir, MountPointInfoFile)
	assert.True(t, IsFileExisting(mountPointInfoPath))

	content, err := os.ReadFile(mountPointInfoPath)
	require.NoError(t, err)
	expectedContent := strings.Join([]string{clientName, storageBackendName, fsName, req.GetVolumeId(), sharedPath}, " ")
	assert.Equal(t, expectedContent, string(content))

	// Verify returned path is correct (should trim hostPrefix if present)
	expectedPath := strings.TrimPrefix(expectedDir, hostPrefix)
	assert.Equal(t, expectedPath, mountPointPath)

	// Test: Second call should not overwrite existing file
	WriteSharedMetricsInfo(metricsPathPrefix, req, "new-client", "new-backend", "new-fs", "/new/path")
	content2, err := os.ReadFile(mountPointInfoPath)
	require.NoError(t, err)
	// Content should remain the same
	assert.Equal(t, expectedContent, string(content2))
}

func TestWriteMetricsInfo(t *testing.T) {
	// Create temporary directory
	tmpDir, err := os.MkdirTemp("", "metrics-test-*")
	require.NoError(t, err)
	defer os.RemoveAll(tmpDir)

	// Create metricsPathPrefix in tmpDir with hostPrefix prefix to test TrimPrefix logic
	metricsPathPrefix := filepath.Join(tmpDir, hostPrefix, "metrics") + "/"
	err = os.MkdirAll(filepath.Dir(metricsPathPrefix), 0755)
	require.NoError(t, err)

	podUID := "pod-uid-12345"
	podNamespace := "test-namespace"
	podName := "test-pod"

	// Create test request
	req := &csi.NodePublishVolumeRequest{
		VolumeId:   "test-volume-456",
		TargetPath: "/mnt/test-target",
		VolumeContext: map[string]string{
			"csi.storage.k8s.io/pod.uid":       podUID,
			"csi.storage.k8s.io/pod.namespace": podNamespace,
			"csi.storage.k8s.io/pod.name":      podName,
		},
	}

	metricsTop := "/metrics/top"
	clientName := "ossfs2"
	storageBackendName := "oss"
	fsName := "test-bucket-2"

	// Test: First call should create directories and files
	mountPointPath := WriteMetricsInfo(metricsPathPrefix, req, metricsTop, clientName, storageBackendName, fsName)

	// Verify pod UID directory was created
	podUIDPath := metricsPathPrefix + podUID + "/"
	assert.True(t, IsFileExisting(podUIDPath))

	// Verify pod_info file was created with correct content
	podInfoPath := podUIDPath + PodInfoFile
	assert.True(t, IsFileExisting(podInfoPath))

	podInfoContent, err := os.ReadFile(podInfoPath)
	require.NoError(t, err)
	expectedPodInfo := strings.Join([]string{podNamespace, podName, podUID, metricsTop}, " ")
	assert.Equal(t, expectedPodInfo, string(podInfoContent))

	// Verify mount point directory was created
	expectedMountDir := podUIDPath + req.GetVolumeId() + "/"
	assert.True(t, IsFileExisting(expectedMountDir))

	// Verify mount_point_info file was created with correct content
	mountPointInfoPath := expectedMountDir + MountPointInfoFile
	assert.True(t, IsFileExisting(mountPointInfoPath))

	mountPointContent, err := os.ReadFile(mountPointInfoPath)
	require.NoError(t, err)
	expectedMountPointInfo := strings.Join([]string{clientName, storageBackendName, fsName, req.GetVolumeId(), req.TargetPath}, " ")
	assert.Equal(t, expectedMountPointInfo, string(mountPointContent))

	// Verify returned path is correct (should trim hostPrefix)
	expectedPath := strings.TrimPrefix(expectedMountDir, hostPrefix)
	assert.Equal(t, expectedPath, mountPointPath)

	// Test: Second call should not overwrite existing files
	WriteMetricsInfo(metricsPathPrefix, req, "/new/metrics/top", "new-client", "new-backend", "new-fs")

	// Pod info should remain the same
	podInfoContent2, err := os.ReadFile(podInfoPath)
	require.NoError(t, err)
	assert.Equal(t, expectedPodInfo, string(podInfoContent2))

	// Mount point info should remain the same
	mountPointContent2, err := os.ReadFile(mountPointInfoPath)
	require.NoError(t, err)
	assert.Equal(t, expectedMountPointInfo, string(mountPointContent2))
}

func TestRemoveMetrics(t *testing.T) {
	// Create temporary directory
	tmpDir, err := os.MkdirTemp("", "metrics-remove-test-*")
	require.NoError(t, err)
	defer os.RemoveAll(tmpDir)

	metricsPathPrefix := filepath.Join(tmpDir, "metrics")
	volumeId := "test-volume-remove"
	req := &csi.NodeUnstageVolumeRequest{
		VolumeId: volumeId,
	}

	mountPointPath := GetFuseMetricsMountDir(metricsPathPrefix, volumeId)

	tests := []struct {
		name        string
		setupFiles  func() // Setup function to create files
		verifyFiles func() // Verify function to check files are removed
		verifyDir   bool   // Whether to verify directory is removed
	}{
		{
			name: "remove all metrics files and directory",
			setupFiles: func() {
				// Create directory
				err := os.MkdirAll(mountPointPath, 0755)
				require.NoError(t, err)

				// Create files from all three arrays
				allFiles := []string{}
				allFiles = append(allFiles, MountpointMetricsArray...)
				allFiles = append(allFiles, CounterTypeMetricsArray...)
				allFiles = append(allFiles, HotSpotMetricsArray...)

				for _, filename := range allFiles {
					filePath := filepath.Join(mountPointPath, filename)
					err := os.WriteFile(filePath, []byte("test content"), 0644)
					require.NoError(t, err)
				}
			},
			verifyFiles: func() {
				// Verify all files are removed
				allFiles := []string{}
				allFiles = append(allFiles, MountpointMetricsArray...)
				allFiles = append(allFiles, CounterTypeMetricsArray...)
				allFiles = append(allFiles, HotSpotMetricsArray...)

				for _, filename := range allFiles {
					filePath := filepath.Join(mountPointPath, filename)
					assert.False(t, IsFileExisting(filePath), "File should be removed: %s", filePath)
				}
			},
			verifyDir: true,
		},
		{
			name: "remove when some files don't exist",
			setupFiles: func() {
				// Create directory
				err := os.MkdirAll(mountPointPath, 0755)
				require.NoError(t, err)

				// Create only some files from MountpointMetricsArray
				for i, filename := range MountpointMetricsArray {
					if i < 2 { // Only create first 2 files
						filePath := filepath.Join(mountPointPath, filename)
						err := os.WriteFile(filePath, []byte("test content"), 0644)
						require.NoError(t, err)
					}
				}
			},
			verifyFiles: func() {
				// Verify all files are removed (including non-existent ones should not cause errors)
				allFiles := []string{}
				allFiles = append(allFiles, MountpointMetricsArray...)
				allFiles = append(allFiles, CounterTypeMetricsArray...)
				allFiles = append(allFiles, HotSpotMetricsArray...)

				for _, filename := range allFiles {
					filePath := filepath.Join(mountPointPath, filename)
					assert.False(t, IsFileExisting(filePath), "File should be removed or not exist: %s", filePath)
				}
			},
			verifyDir: true,
		},
		{
			name: "remove when directory doesn't exist",
			setupFiles: func() {
				// Don't create directory
			},
			verifyFiles: func() {
				// Should not cause errors
			},
			verifyDir: false,
		},
		{
			name: "remove when directory is empty",
			setupFiles: func() {
				// Create empty directory
				err := os.MkdirAll(mountPointPath, 0755)
				require.NoError(t, err)
			},
			verifyFiles: func() {
				// Should not cause errors
			},
			verifyDir: true,
		},
		{
			name: "directory not removed when contains other files",
			setupFiles: func() {
				// Create directory
				err := os.MkdirAll(mountPointPath, 0755)
				require.NoError(t, err)

				// Create files from all three arrays
				allFiles := []string{}
				allFiles = append(allFiles, MountpointMetricsArray...)
				allFiles = append(allFiles, CounterTypeMetricsArray...)
				allFiles = append(allFiles, HotSpotMetricsArray...)

				for _, filename := range allFiles {
					filePath := filepath.Join(mountPointPath, filename)
					err := os.WriteFile(filePath, []byte("test content"), 0644)
					require.NoError(t, err)
				}

				// Create an additional file that is not in any of the arrays
				otherFile := filepath.Join(mountPointPath, "other_file.txt")
				err = os.WriteFile(otherFile, []byte("other content"), 0644)
				require.NoError(t, err)
			},
			verifyFiles: func() {
				// Verify all metrics files are removed
				allFiles := []string{}
				allFiles = append(allFiles, MountpointMetricsArray...)
				allFiles = append(allFiles, CounterTypeMetricsArray...)
				allFiles = append(allFiles, HotSpotMetricsArray...)

				for _, filename := range allFiles {
					filePath := filepath.Join(mountPointPath, filename)
					assert.False(t, IsFileExisting(filePath), "File should be removed: %s", filePath)
				}

				// Verify other file still exists
				otherFile := filepath.Join(mountPointPath, "other_file.txt")
				assert.True(t, IsFileExisting(otherFile), "Other file should still exist: %s", otherFile)
			},
			verifyDir: false, // Directory should not be removed because it contains other files
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			tt.setupFiles()

			// Execute
			RemoveMetrics(metricsPathPrefix, req)

			// Verify
			tt.verifyFiles()
			if tt.verifyDir {
				assert.False(t, IsFileExisting(mountPointPath), "Directory should be removed: %s", mountPointPath)
			} else {
				// If verifyDir is false, directory should still exist (may contain other files)
				// Only check if directory was created in setup
				if _, err := os.Stat(mountPointPath); err == nil {
					assert.True(t, IsFileExisting(mountPointPath), "Directory should still exist: %s", mountPointPath)
				}
			}
		})
	}
}
