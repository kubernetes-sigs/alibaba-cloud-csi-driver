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

	"context"
	"errors"

	nasclient "github.com/alibabacloud-go/nas-20170626/v4/client"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

// MockFileSetManager is a mock implementation of CPFSFileSetManager for testing
type MockFileSetManager struct {
	DescribeFilesetFunc  func(ctx context.Context, fsID, fileSetID string) (*nasclient.DescribeFilesetsResponseBodyEntriesEntrie, error)
	DeleteFileSetFunc    func(ctx context.Context, fsID, fileSetID string) error
	DescribeFileSetsFunc func(ctx context.Context, fsID string) ([]*nasclient.DescribeFilesetsResponseBodyEntriesEntrie, error)
}

func (m *MockFileSetManager) CreateFileSet(ctx context.Context, fsID, pvName, fileSystemPath string, countLimit, sizeBytes int64, deleteProtection bool) (string, error) {
	return "", nil
}

func (m *MockFileSetManager) DeleteFileSet(ctx context.Context, fsID, fileSetID string) error {
	if m.DeleteFileSetFunc != nil {
		return m.DeleteFileSetFunc(ctx, fsID, fileSetID)
	}
	return nil
}

func (m *MockFileSetManager) DescribeFileSets(ctx context.Context, fsID string) ([]*nasclient.DescribeFilesetsResponseBodyEntriesEntrie, error) {
	if m.DescribeFileSetsFunc != nil {
		return m.DescribeFileSetsFunc(ctx, fsID)
	}
	return nil, nil
}

func (m *MockFileSetManager) DescribeFileset(ctx context.Context, fsID, fileSetID string) (*nasclient.DescribeFilesetsResponseBodyEntriesEntrie, error) {
	if m.DescribeFilesetFunc != nil {
		return m.DescribeFilesetFunc(ctx, fsID, fileSetID)
	}
	return nil, nil
}

// Create a test controller server with mock file set manager
func createTestControllerServer() *controllerServer {
	return &controllerServer{
		filesetManager: &MockFileSetManager{},
	}
}

func TestControllerServer_DeleteVolume_Success(t *testing.T) {
	cs := createTestControllerServer()
	mockManager := cs.filesetManager.(*MockFileSetManager)

	ctx := context.Background()
	req := &csi.DeleteVolumeRequest{
		VolumeId: "fs-123+fset-456",
	}

	// Mock DescribeFileset to return existing fileset
	mockManager.DescribeFilesetFunc = func(ctx context.Context, fsID, fileSetID string) (*nasclient.DescribeFilesetsResponseBodyEntriesEntrie, error) {
		assert.Equal(t, "fs-123", fsID)
		assert.Equal(t, "fset-456", fileSetID)
		return &nasclient.DescribeFilesetsResponseBodyEntriesEntrie{
			FsetId: &fileSetID,
		}, nil
	}

	// Mock DeleteFileSet to succeed
	mockManager.DeleteFileSetFunc = func(ctx context.Context, fsID, fileSetID string) error {
		assert.Equal(t, "fs-123", fsID)
		assert.Equal(t, "fset-456", fileSetID)
		return nil
	}

	resp, err := cs.DeleteVolume(ctx, req)

	require.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestControllerServer_DeleteVolume_InvalidVolumeId(t *testing.T) {
	cs := createTestControllerServer()

	ctx := context.Background()
	req := &csi.DeleteVolumeRequest{
		VolumeId: "invalid-format",
	}

	resp, err := cs.DeleteVolume(ctx, req)

	require.Error(t, err)
	assert.Nil(t, resp)

	statusErr, ok := status.FromError(err)
	require.True(t, ok)
	assert.Equal(t, codes.InvalidArgument, statusErr.Code())
	assert.Contains(t, statusErr.Message(), "invalid volume ID format")
}

func TestControllerServer_DeleteVolume_EmptyVolumeId(t *testing.T) {
	cs := createTestControllerServer()

	ctx := context.Background()
	req := &csi.DeleteVolumeRequest{
		VolumeId: "",
	}

	resp, err := cs.DeleteVolume(ctx, req)

	require.Error(t, err)
	assert.Nil(t, resp)

	statusErr, ok := status.FromError(err)
	require.True(t, ok)
	assert.Equal(t, codes.InvalidArgument, statusErr.Code())
}

func TestControllerServer_DeleteVolume_FilesetNotFound(t *testing.T) {
	cs := createTestControllerServer()
	mockManager := cs.filesetManager.(*MockFileSetManager)

	ctx := context.Background()
	req := &csi.DeleteVolumeRequest{
		VolumeId: "fs-123+fset-nonexistent",
	}

	// Mock DescribeFileset to return error (fileset not found)
	mockManager.DescribeFilesetFunc = func(ctx context.Context, fsID, fileSetID string) (*nasclient.DescribeFilesetsResponseBodyEntriesEntrie, error) {
		return nil, errors.New("fileset not found")
	}

	// DeleteFileSet should not be called
	mockManager.DeleteFileSetFunc = func(ctx context.Context, fsID, fileSetID string) error {
		t.Fatal("DeleteFileSet should not be called when fileset doesn't exist")
		return nil
	}

	resp, err := cs.DeleteVolume(ctx, req)

	require.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestControllerServer_DeleteVolume_DeleteFileSetError(t *testing.T) {
	cs := createTestControllerServer()
	mockManager := cs.filesetManager.(*MockFileSetManager)

	ctx := context.Background()
	req := &csi.DeleteVolumeRequest{
		VolumeId: "fs-123+fset-456",
	}

	// Mock DescribeFileset to return existing fileset
	mockManager.DescribeFilesetFunc = func(ctx context.Context, fsID, fileSetID string) (*nasclient.DescribeFilesetsResponseBodyEntriesEntrie, error) {
		return &nasclient.DescribeFilesetsResponseBodyEntriesEntrie{
			FsetId: &fileSetID,
		}, nil
	}

	// Mock DeleteFileSet to return error
	mockManager.DeleteFileSetFunc = func(ctx context.Context, fsID, fileSetID string) error {
		return errors.New("delete failed")
	}

	resp, err := cs.DeleteVolume(ctx, req)

	require.Error(t, err)
	assert.Nil(t, resp)

	statusErr, ok := status.FromError(err)
	require.True(t, ok)
	assert.Equal(t, codes.Internal, statusErr.Code())
	assert.Contains(t, statusErr.Message(), "delete failed")
}
