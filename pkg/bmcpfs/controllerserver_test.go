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
	"fmt"
	"testing"

	"context"
	"errors"

	nasclient "github.com/alibabacloud-go/nas-20170626/v4/client"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/golang/mock/gomock"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/bmcpfs/internal"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/stretchr/testify/assert"
)

// TestParseVolumeHandle tests the parseVolumeHandle function with various inputs
func TestParseVolumeHandle(t *testing.T) {
	tests := []struct {
		name           string
		volumeHandle   string
		expectedFsID   string
		expectedFsetID string
	}{
		{
			name:           "single part",
			volumeHandle:   "fs-12345",
			expectedFsID:   "fs-12345",
			expectedFsetID: "",
		},
		{
			name:           "two parts",
			volumeHandle:   "fs-12345+fset-67890",
			expectedFsID:   "fs-12345",
			expectedFsetID: "fset-67890",
		},
		{
			name:           "multiple delimiters",
			volumeHandle:   "fs-12345+fset-67890+extra",
			expectedFsID:   "fs-12345",
			expectedFsetID: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fsID, fsetID := parseVolumeHandle(tt.volumeHandle)
			assert.Equal(t, tt.expectedFsID, fsID)
			assert.Equal(t, tt.expectedFsetID, fsetID)
		})
	}
}

func TestControllerServer_CreateVolume(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockNas := internal.NewGoMockNasClient(ctrl)

	// Mock the CreateFileset method
	mockNas.EXPECT().
		CreateFileset(gomock.Any()).
		DoAndReturn(func(req *nasclient.CreateFilesetRequest) (*nasclient.CreateFilesetResponse, error) {
			if *req.FileSystemId == "" {
				return nil, errors.New("file system ID is required")
			}
			if *req.FileSystemPath == "" {
				return nil, errors.New("file system path is required")
			}
			if *req.FileSystemId == "valid_fs" {
				return &nasclient.CreateFilesetResponse{
					Body: &nasclient.CreateFilesetResponseBody{
						FsetId: req.ClientToken,
					},
				}, nil
			}
			if *req.FileSystemId == "create_failed" {
				return nil, errors.New("create failed")
			}
			return nil, nil
		}).AnyTimes()

	// Mock the DescribeFilesets method for checking if fileset exists
	mockNas.EXPECT().
		DescribeFilesets(gomock.Any()).
		DoAndReturn(func(req *nasclient.DescribeFilesetsRequest) (*nasclient.DescribeFilesetsResponse, error) {
			// Check if we're filtering for a specific fileset
			isSpecificFileset := false
			for _, filter := range req.Filters {
				if *filter.Key == "FsetIds" {
					isSpecificFileset = true
					break
				}
			}

			// Check if we're filtering by filesystem path
			isPathFilter := false
			for _, filter := range req.Filters {
				if *filter.Key == "FileSystemPath" {
					isPathFilter = true
					break
				}
			}

			if *req.FileSystemId == "valid_fs" && isSpecificFileset {
				// Return the specific fileset as CREATED
				status := "CREATED"
				return &nasclient.DescribeFilesetsResponse{
					Body: &nasclient.DescribeFilesetsResponseBody{
						Entries: &nasclient.DescribeFilesetsResponseBodyEntries{
							Entrie: []*nasclient.DescribeFilesetsResponseBodyEntriesEntrie{
								{
									FsetId: req.Filters[0].Value,
									Status: &status,
								},
							},
						},
					},
				}, nil
			} else if *req.FileSystemId == "valid_fs" && isPathFilter {
				// Simulate fileset not found by path (for CreateVolume flow)
				// Return empty entries to simulate fileset doesn't exist yet
				return &nasclient.DescribeFilesetsResponse{
					Body: &nasclient.DescribeFilesetsResponseBody{
						Entries: &nasclient.DescribeFilesetsResponseBodyEntries{
							Entrie: []*nasclient.DescribeFilesetsResponseBodyEntriesEntrie{},
						},
					},
				}, nil
			} else if *req.FileSystemId == "valid_fs" {
				// Return empty entries to simulate fileset doesn't exist yet
				return &nasclient.DescribeFilesetsResponse{
					Body: &nasclient.DescribeFilesetsResponseBody{
						Entries: &nasclient.DescribeFilesetsResponseBodyEntries{
							Entrie: []*nasclient.DescribeFilesetsResponseBodyEntriesEntrie{},
						},
					},
				}, nil
			}
			return nil, errors.New("unexpected call")
		}).AnyTimes()

	// Create fileset manager
	fsm := internal.NewCPFSFileSetManager(mockNas)
	ctx := context.Background()

	// Test cases for fileset manager
	tests := []struct {
		name           string
		fsID           string
		pvName         string
		fileSystemPath string
		expectError    string
	}{
		{
			name:           "filesystem empty",
			fsID:           "",
			pvName:         "test-pv",
			fileSystemPath: "/test-path",
			expectError:    fmt.Sprintf("create fileset %s/%s failed: %s", "", "test-pv", "file system ID is required"),
		},
		{
			name:           "normal create",
			fsID:           "valid_fs",
			pvName:         "test-pv",
			fileSystemPath: "/test-path",
			expectError:    "",
		},
		{
			name:           "create failed",
			fsID:           "create_failed",
			pvName:         "test-pv",
			fileSystemPath: "/test-path",
			expectError:    fmt.Sprintf("create fileset %s/%s failed: %s", "create_failed", "test-pv", "create failed"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := fsm.CreateFileSet(ctx, tt.fsID, tt.pvName, tt.fileSystemPath, 1000000, 1024*1024*1024, false)
			if err != nil {
				assert.Equal(t, tt.expectError, err.Error())
			}
		})
	}

	// Test cases for controller server CreateVolume
	controllerTests := []struct {
		name        string
		request     *csi.CreateVolumeRequest
		expectError string
	}{
		{
			name: "missing bmcpfsId parameter",
			request: &csi.CreateVolumeRequest{
				Name: "test-volume",
				Parameters: map[string]string{
					"path": "/test",
				},
				VolumeCapabilities: []*csi.VolumeCapability{
					{
						AccessType: &csi.VolumeCapability_Mount{
							Mount: &csi.VolumeCapability_MountVolume{
								FsType: "bmcpfs",
							},
						},
					},
				},
			},
			expectError: "Invalid parameters from input: test-volume, with error: bmcpfsId parameter is required",
		},
		{
			name: "valid request",
			request: &csi.CreateVolumeRequest{
				Name: "test-volume",
				Parameters: map[string]string{
					"bmcpfsId": "valid_fs",
					"path":     "/test",
				},
				VolumeCapabilities: []*csi.VolumeCapability{
					{
						AccessType: &csi.VolumeCapability_Mount{
							Mount: &csi.VolumeCapability_MountVolume{
								FsType: "bmcpfs",
							},
						},
					},
				},
				CapacityRange: &csi.CapacityRange{
					RequiredBytes: 1024 * 1024 * 1024,
				},
			},
			expectError: "",
		},
	}

	// Create a mock controller server
	cs := &controllerServer{
		filesetManager: fsm,
		locks:          utils.NewVolumeLocks(),
	}

	for _, tt := range controllerTests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := cs.CreateVolume(ctx, tt.request)
			if tt.expectError != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectError)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, tt.request.Parameters["bmcpfsId"], resp.Volume.VolumeId)
			}
		})
	}
}

func TestControllerServer_DeleteVolume(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockNas := internal.NewGoMockNasClient(ctrl)

	mockNas.EXPECT().
		DeleteFileset(gomock.Any()).AnyTimes().
		DoAndReturn(func(req *nasclient.DeleteFilesetRequest) (*nasclient.DeleteFilesetResponse, error) {
			if *req.FileSystemId == "" {
				return nil, errors.New("file system ID is required")
			}
			if *req.FsetId == "" {
				return nil, errors.New("file set ID is required")
			}
			if *req.FileSystemId == "normal_delete" {
				return &nasclient.DeleteFilesetResponse{
					Body: &nasclient.DeleteFilesetResponseBody{},
				}, nil
			}
			if *req.FileSystemId == "delete_failed" {
				return nil, errors.New("delete failed")
			}
			// Return a valid response for the successful case
			return &nasclient.DeleteFilesetResponse{
				Body: &nasclient.DeleteFilesetResponseBody{},
			}, nil
		})

	fsm := internal.NewCPFSFileSetManager(mockNas)
	ctx := context.Background()

	tests := []struct {
		name        string
		fsID        string
		filesetID   string
		expectError string
	}{
		{
			name:        "filesystem empty",
			fsID:        "",
			filesetID:   "",
			expectError: fmt.Sprintf("delete fileset %s/%s failed: %s", "", "", "file system ID is required"),
		},
		{
			name:        "fileset empty",
			fsID:        "xxx",
			filesetID:   "",
			expectError: fmt.Sprintf("delete fileset %s/%s failed: %s", "xxx", "", "file set ID is required"),
		},
		{
			name:        "normal delete",
			fsID:        "normal_delete",
			filesetID:   "fset-123",
			expectError: "",
		},
		{
			name:        "delete failed",
			fsID:        "delete_failed",
			filesetID:   "fset-123",
			expectError: "delete fileset delete_failed/fset-123 failed: delete failed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := fsm.DeleteFileSet(ctx, tt.fsID, tt.filesetID)
			if tt.expectError != "" {
				assert.Error(t, err)
				assert.Equal(t, tt.expectError, err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}

	// Test controller server DeleteVolume
	controllerTests := []struct {
		name        string
		volumeId    string
		expectError string
	}{
		{
			name:        "valid volume id",
			volumeId:    "fs-123+fset-456",
			expectError: "",
		},
		{
			name:        "invalid volume id format",
			volumeId:    "invalid-volume-id",
			expectError: "rpc error: code = InvalidArgument desc = invalid volume ID format: invalid-volume-id, expected format: filesystemId+filesetId",
		},
		{
			name:        "empty volume id",
			volumeId:    "",
			expectError: "rpc error: code = InvalidArgument desc = invalid volume ID format: , expected format: filesystemId+filesetId",
		},
	}

	cs := &controllerServer{
		filesetManager: fsm,
	}

	for _, tt := range controllerTests {
		t.Run(tt.name, func(t *testing.T) {
			req := &csi.DeleteVolumeRequest{
				VolumeId: tt.volumeId,
			}
			_, err := cs.DeleteVolume(ctx, req)
			if tt.expectError != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestControllerServer_ControllerGetCapabilities(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	cs := &controllerServer{}

	req := &csi.ControllerGetCapabilitiesRequest{}
	resp, err := cs.ControllerGetCapabilities(ctx, req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Capabilities)
}
