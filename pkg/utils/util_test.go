/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package utils

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/klog/v2"
	k8smount "k8s.io/mount-utils"
)

func TestGetServiceType(t *testing.T) {
	defer func() { klog.OsExit = os.Exit }()
	klog.OsExit = func(c int) { panic(c) }
	tests := []struct {
		name                 string
		runAsController      bool
		runControllerService bool
		runNodeService       bool
		serviceTypeEnv       string
		want                 ServiceType
		fatal                bool
	}{
		{
			name:                 "default",
			runControllerService: true,
			runNodeService:       true,
			want:                 Controller | Node,
		},
		{
			name:            "Run as controller",
			runAsController: true,
			want:            Controller,
		},
		{
			name:           "env provisioner",
			serviceTypeEnv: ProvisionerService,
			want:           Controller,
		},
		{
			name:           "env plugin",
			serviceTypeEnv: PluginService,
			want:           Node,
		},
		{
			name:                 "Run controller",
			runControllerService: true,
			runNodeService:       false,
			want:                 Controller,
		},
		{
			name:                 "Run node",
			runControllerService: false,
			runNodeService:       true,
			want:                 Node,
		},
		{
			name:                 "nothing",
			runControllerService: false,
			runNodeService:       false,
			want:                 0,
		},
		{
			name:           "invalid env",
			serviceTypeEnv: "invalid",
			fatal:          true,
		},
		{
			name:            "conflict env and run-as-controller",
			runAsController: true,
			serviceTypeEnv:  PluginService,
			fatal:           true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("SERVICE_TYPE", tt.serviceTypeEnv)
			if tt.fatal {
				assert.Panics(t, func() {
					GetServiceType(tt.runAsController, tt.runControllerService, tt.runNodeService)
				})
				return
			}
			got := GetServiceType(tt.runAsController, tt.runControllerService, tt.runNodeService)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestIsDirTmpfs(t *testing.T) {
	mounter := k8smount.NewFakeMounter([]k8smount.MountPoint{
		{
			Path: "/",
		},
		{
			Path: "/some/tmpfs",
			Type: "tmpfs",
		},
		{
			Path: "/some/other",
			Type: "ext4",
		},
	})
	isTmpfs, err := IsDirTmpfs(mounter, "/some/tmpfs")
	assert.Nil(t, err)
	assert.True(t, isTmpfs)
}

func TestIsDirTmpfsFalse(t *testing.T) {
	mounter := k8smount.NewFakeMounter([]k8smount.MountPoint{
		{
			Path: "/",
		},
		{
			Path: "/some/other",
			Type: "ext4",
		},
	})
	isTmpfs, err := IsDirTmpfs(mounter, "/some/tmpfs")
	assert.Nil(t, err)
	assert.False(t, isTmpfs)
}

func TestGetFuseMetricsMountDir(t *testing.T) {
	metrcisPrefix := "/test/metrics"
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
			actual := GetFuseMetricsMountDir(metrcisPrefix, tt.volumeId)
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
	mountPointInfoPath := filepath.Join(expectedDir, "mount_point_info")
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
	podInfoPath := podUIDPath + "pod_info"
	assert.True(t, IsFileExisting(podInfoPath))

	podInfoContent, err := os.ReadFile(podInfoPath)
	require.NoError(t, err)
	expectedPodInfo := strings.Join([]string{podNamespace, podName, podUID, metricsTop}, " ")
	assert.Equal(t, expectedPodInfo, string(podInfoContent))

	// Verify mount point directory was created
	expectedMountDir := podUIDPath + req.GetVolumeId() + "/"
	assert.True(t, IsFileExisting(expectedMountDir))

	// Verify mount_point_info file was created with correct content
	mountPointInfoPath := expectedMountDir + "mount_point_info"
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
