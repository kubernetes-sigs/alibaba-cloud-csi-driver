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

package bmcpfs

import (
	"testing"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseVolumeID_Success(t *testing.T) {
	fsID, fileSetID, err := parseVolumeID("fs-123+fset-456")

	require.NoError(t, err)
	assert.Equal(t, "fs-123", fsID)
	assert.Equal(t, "fset-456", fileSetID)
}

func TestParseVolumeID_InvalidFormat_NoSeparator(t *testing.T) {
	fsID, fileSetID, err := parseVolumeID("fs-123fset-456")

	require.Error(t, err)
	assert.Empty(t, fsID)
	assert.Empty(t, fileSetID)
	assert.Contains(t, err.Error(), "invalid volume ID format")
}

func TestParseVolumeID_InvalidFormat_MultipleSeparators(t *testing.T) {
	fsID, fileSetID, err := parseVolumeID("fs-123+fset-456+extra")

	require.Error(t, err)
	assert.Empty(t, fsID)
	assert.Empty(t, fileSetID)
	assert.Contains(t, err.Error(), "invalid volume ID format")
}

func TestParseVolumeID_EmptyFilesystemID(t *testing.T) {
	fsID, fileSetID, err := parseVolumeID("+fset-456")

	require.Error(t, err)
	assert.Empty(t, fsID)
	assert.Empty(t, fileSetID)
	assert.Contains(t, err.Error(), "filesystem ID or fileset ID is empty")
}

func TestParseVolumeID_EmptyFilesetID(t *testing.T) {
	fsID, fileSetID, err := parseVolumeID("fs-123+")

	require.Error(t, err)
	assert.Empty(t, fsID)
	assert.Empty(t, fileSetID)
	assert.Contains(t, err.Error(), "filesystem ID or fileset ID is empty")
}

func TestParseVolumeID_EmptyVolumeID(t *testing.T) {
	fsID, fileSetID, err := parseVolumeID("")

	require.Error(t, err)
	assert.Empty(t, fsID)
	assert.Empty(t, fileSetID)
	assert.Contains(t, err.Error(), "invalid volume ID format")
}

func TestParseVolumeID_WithSpecialCharacters(t *testing.T) {
	fsID, fileSetID, err := parseVolumeID("fs-abc-123+fset-def-456")

	require.NoError(t, err)
	assert.Equal(t, "fs-abc-123", fsID)
	assert.Equal(t, "fset-def-456", fileSetID)
}

func TestValidateFileSetParameters_MissingBmcpfsId(t *testing.T) {
	req := &csi.CreateVolumeRequest{
		Parameters: map[string]string{},
	}

	err := validateFileSetParameters(req)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "bmcpfsId parameter is required")
}

func TestValidateFileSetParameters_InvalidFsType(t *testing.T) {
	req := &csi.CreateVolumeRequest{
		Parameters: map[string]string{
			"bmcpfsId": "fs-123",
		},
		VolumeCapabilities: []*csi.VolumeCapability{
			{
				AccessType: &csi.VolumeCapability_Mount{
					Mount: &csi.VolumeCapability_MountVolume{
						FsType: "ext4", // Invalid fsType
					},
				},
			},
		},
	}

	err := validateFileSetParameters(req)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "fsType ext4 is not supported, only 'cpfs' is supported")
}

func TestValidateFileSetParameters_ValidParams(t *testing.T) {
	req := &csi.CreateVolumeRequest{
		Parameters: map[string]string{
			"bmcpfsId": "fs-123",
		},
		VolumeCapabilities: []*csi.VolumeCapability{
			{
				AccessType: &csi.VolumeCapability_Mount{
					Mount: &csi.VolumeCapability_MountVolume{
						FsType: "bmcpfs", // Valid fsType
					},
				},
			},
		},
	}

	err := validateFileSetParameters(req)
	require.NoError(t, err)
}

func TestCreateVolumeResponse(t *testing.T) {
	fileSetID := "fset-123"
	bmcpfsID := "fs-456"
	volSizeBytes := int64(1024 * 1024 * 1024) // 1GB
	volumeContext := map[string]string{
		"test-key": "test-value",
	}

	result := createVolumeResponse(fileSetID, bmcpfsID, volSizeBytes, volumeContext)

	assert.Equal(t, volSizeBytes, result.CapacityBytes)
	assert.Equal(t, "fs-456", result.VolumeId)
	assert.Equal(t, "bmcpfs", result.VolumeContext[VolumeType])
	assert.Equal(t, "fset-123", result.VolumeContext[FileSetID])
	assert.Equal(t, "test-value", result.VolumeContext["test-key"])
}

func TestUpdateVolumeContext(t *testing.T) {
	inputContext := map[string]string{
		"test-key":                                 "test-value",  // should keep
		"csi.alibabacloud.com/lastApply":           "some-value",  // should remove
		common.PVNameKey:                           "pv-name",     // should remove
		common.PVCNameKey:                          "pvc-name",    // should remove
		common.PVCNamespaceKey:                     "namespace",   // should remove
		"csi.alibabacloud.com/storage-provisioner": "provisioner", // should remove
	}

	result := updateVolumeContext(inputContext)

	// Check that values we want to keep are still there
	assert.Equal(t, "test-value", result["test-key"])

	// Check that values we want to remove are gone
	_, ok := result["csi.alibabacloud.com/lastApply"]
	assert.False(t, ok)
	_, ok = result["csi.alibabacloud.com/storage-provisioner"]
	assert.False(t, ok)
	_, ok = result[common.PVNameKey]
	assert.False(t, ok)
	_, ok = result[common.PVCNameKey]
	assert.False(t, ok)
	_, ok = result[common.PVCNamespaceKey]
	assert.False(t, ok)
}
