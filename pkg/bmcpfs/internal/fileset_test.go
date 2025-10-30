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

package internal

import (
	"context"
	"errors"
	"testing"

	nasclient "github.com/alibabacloud-go/nas-20170626/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewCPFSFileSetManager(t *testing.T) {
	mockClient := NewMockNasClient()
	manager := NewCPFSFileSetManager(mockClient)

	assert.NotNil(t, manager)
	assert.IsType(t, &cpfsFileSetManager{}, manager)
}

func TestCPFSFileSetManager_CreateFileSet_Success(t *testing.T) {
	mockClient := NewMockNasClient()
	manager := NewCPFSFileSetManager(mockClient)

	ctx := context.Background()
	fsID := "fs-123"
	pvName := "test-pv"
	fileSystemPath := "/test/path"
	countLimit := int64(1000)
	sizeBytes := int64(1024 * 1024) // 1MB
	deleteProtection := true
	expectedFileSetID := "fset-456"

	expectedRequest := &nasclient.CreateFilesetRequest{
		FileSystemId:       tea.String(fsID),
		ClientToken:        tea.String(pvName),
		DeletionProtection: tea.Bool(deleteProtection),
		FileSystemPath:     tea.String(fileSystemPath),
		Quota: &nasclient.CreateFilesetRequestQuota{
			SizeLimit:      tea.Int64(sizeBytes),
			FileCountLimit: tea.Int64(countLimit),
		},
	}

	mockResponse := &nasclient.CreateFilesetResponse{
		Body: &nasclient.CreateFilesetResponseBody{
			FsetId: tea.String(expectedFileSetID),
		},
	}

	mockClient.CreateFilesetFunc = func(request *nasclient.CreateFilesetRequest) (*nasclient.CreateFilesetResponse, error) {
		assert.Equal(t, expectedRequest.FileSystemId, request.FileSystemId)
		assert.Equal(t, expectedRequest.ClientToken, request.ClientToken)
		assert.Equal(t, expectedRequest.DeletionProtection, request.DeletionProtection)
		assert.Equal(t, expectedRequest.FileSystemPath, request.FileSystemPath)
		assert.Equal(t, expectedRequest.Quota.SizeLimit, request.Quota.SizeLimit)
		assert.Equal(t, expectedRequest.Quota.FileCountLimit, request.Quota.FileCountLimit)
		return mockResponse, nil
	}

	fileSetID, err := manager.CreateFileSet(ctx, fsID, pvName, fileSystemPath, countLimit, sizeBytes, deleteProtection)

	require.NoError(t, err)
	assert.Equal(t, expectedFileSetID, fileSetID)
}

func TestCPFSFileSetManager_CreateFileSet_WithZeroCountLimit(t *testing.T) {
	mockClient := NewMockNasClient()
	manager := NewCPFSFileSetManager(mockClient)

	ctx := context.Background()
	fsID := "fs-123"
	pvName := "test-pv"
	fileSystemPath := "/test/path"
	countLimit := int64(0) // Zero count limit
	sizeBytes := int64(1024 * 1024)
	deleteProtection := false
	expectedFileSetID := "fset-789"

	mockResponse := &nasclient.CreateFilesetResponse{
		Body: &nasclient.CreateFilesetResponseBody{
			FsetId: tea.String(expectedFileSetID),
		},
	}

	mockClient.CreateFilesetFunc = func(request *nasclient.CreateFilesetRequest) (*nasclient.CreateFilesetResponse, error) {
		// When countLimit is 0, FileCountLimit should not be set in the request
		assert.Nil(t, request.Quota.FileCountLimit)
		return mockResponse, nil
	}

	fileSetID, err := manager.CreateFileSet(ctx, fsID, pvName, fileSystemPath, countLimit, sizeBytes, deleteProtection)

	require.NoError(t, err)
	assert.Equal(t, expectedFileSetID, fileSetID)
}

func TestCPFSFileSetManager_CreateFileSet_Error(t *testing.T) {
	mockClient := NewMockNasClient()
	manager := NewCPFSFileSetManager(mockClient)

	ctx := context.Background()
	fsID := "fs-123"
	pvName := "test-pv"
	fileSystemPath := "/test/path"
	countLimit := int64(1000)
	sizeBytes := int64(1024 * 1024)
	deleteProtection := true

	clientError := errors.New("client error")
	mockClient.CreateFilesetFunc = func(request *nasclient.CreateFilesetRequest) (*nasclient.CreateFilesetResponse, error) {
		return nil, clientError
	}

	fileSetID, err := manager.CreateFileSet(ctx, fsID, pvName, fileSystemPath, countLimit, sizeBytes, deleteProtection)

	require.Error(t, err)
	assert.Empty(t, fileSetID)
	assert.Contains(t, err.Error(), "create fileset fs-123/test-pv failed")
	assert.Contains(t, err.Error(), "client error")
}

func TestCPFSFileSetManager_CreateFileSet_NilResponse(t *testing.T) {
	mockClient := NewMockNasClient()
	manager := NewCPFSFileSetManager(mockClient)

	ctx := context.Background()
	fsID := "fs-123"
	pvName := "test-pv"
	fileSystemPath := "/test/path"
	countLimit := int64(1000)
	sizeBytes := int64(1024 * 1024)
	deleteProtection := true

	mockClient.CreateFilesetFunc = func(request *nasclient.CreateFilesetRequest) (*nasclient.CreateFilesetResponse, error) {
		return nil, nil
	}

	fileSetID, err := manager.CreateFileSet(ctx, fsID, pvName, fileSystemPath, countLimit, sizeBytes, deleteProtection)

	require.Error(t, err)
	assert.Empty(t, fileSetID)
	assert.Contains(t, err.Error(), "create fileset fs-123//test/path failed")
	assert.Contains(t, err.Error(), "empty response")
}

func TestCPFSFileSetManager_CreateFileSet_NilResponseBody(t *testing.T) {
	mockClient := NewMockNasClient()
	manager := NewCPFSFileSetManager(mockClient)

	ctx := context.Background()
	fsID := "fs-123"
	pvName := "test-pv"
	fileSystemPath := "/test/path"
	countLimit := int64(1000)
	sizeBytes := int64(1024 * 1024)
	deleteProtection := true

	mockResponse := &nasclient.CreateFilesetResponse{
		Body: nil,
	}

	mockClient.CreateFilesetFunc = func(request *nasclient.CreateFilesetRequest) (*nasclient.CreateFilesetResponse, error) {
		return mockResponse, nil
	}

	fileSetID, err := manager.CreateFileSet(ctx, fsID, pvName, fileSystemPath, countLimit, sizeBytes, deleteProtection)

	require.Error(t, err)
	assert.Empty(t, fileSetID)
	assert.Contains(t, err.Error(), "create fileset fs-123//test/path failed")
	assert.Contains(t, err.Error(), "empty response")
}

func TestCPFSFileSetManager_CreateFileSet_NilFsetId(t *testing.T) {
	mockClient := NewMockNasClient()
	manager := NewCPFSFileSetManager(mockClient)

	ctx := context.Background()
	fsID := "fs-123"
	pvName := "test-pv"
	fileSystemPath := "/test/path"
	countLimit := int64(1000)
	sizeBytes := int64(1024 * 1024)
	deleteProtection := true

	mockResponse := &nasclient.CreateFilesetResponse{
		Body: &nasclient.CreateFilesetResponseBody{
			FsetId: nil,
		},
	}

	mockClient.CreateFilesetFunc = func(request *nasclient.CreateFilesetRequest) (*nasclient.CreateFilesetResponse, error) {
		return mockResponse, nil
	}

	fileSetID, err := manager.CreateFileSet(ctx, fsID, pvName, fileSystemPath, countLimit, sizeBytes, deleteProtection)

	require.Error(t, err)
	assert.Empty(t, fileSetID)
	assert.Contains(t, err.Error(), "create fileset fs-123//test/path failed")
	assert.Contains(t, err.Error(), "empty response")
}

func TestCPFSFileSetManager_DeleteFileSet_Success(t *testing.T) {
	mockClient := NewMockNasClient()
	manager := NewCPFSFileSetManager(mockClient)

	ctx := context.Background()
	fsID := "fs-123"
	fileSetID := "fset-456"

	expectedRequest := &nasclient.DeleteFilesetRequest{
		FileSystemId: tea.String(fsID),
		FsetId:       tea.String(fileSetID),
	}

	mockResponse := &nasclient.DeleteFilesetResponse{
		Body: &nasclient.DeleteFilesetResponseBody{},
	}

	mockClient.DeleteFilesetFunc = func(request *nasclient.DeleteFilesetRequest) (*nasclient.DeleteFilesetResponse, error) {
		assert.Equal(t, expectedRequest.FileSystemId, request.FileSystemId)
		assert.Equal(t, expectedRequest.FsetId, request.FsetId)
		return mockResponse, nil
	}

	err := manager.DeleteFileSet(ctx, fsID, fileSetID)

	require.NoError(t, err)
}

func TestCPFSFileSetManager_DeleteFileSet_Error(t *testing.T) {
	mockClient := NewMockNasClient()
	manager := NewCPFSFileSetManager(mockClient)

	ctx := context.Background()
	fsID := "fs-123"
	fileSetID := "fset-456"
	expectedError := errors.New("delete failed")

	mockClient.DeleteFilesetFunc = func(request *nasclient.DeleteFilesetRequest) (*nasclient.DeleteFilesetResponse, error) {
		return nil, expectedError
	}

	err := manager.DeleteFileSet(ctx, fsID, fileSetID)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "delete fileset")
	assert.Contains(t, err.Error(), expectedError.Error())
}

func TestCPFSFileSetManager_DescribeFileSets_Success(t *testing.T) {
	mockClient := NewMockNasClient()
	manager := NewCPFSFileSetManager(mockClient)

	ctx := context.Background()
	fsID := "fs-123"
	expectedFileSet1 := &nasclient.DescribeFilesetsResponseBodyEntriesEntrie{
		FsetId: tea.String("fset-1"),
	}
	expectedFileSet2 := &nasclient.DescribeFilesetsResponseBodyEntriesEntrie{
		FsetId: tea.String("fset-2"),
	}

	// First page
	mockResponse1 := &nasclient.DescribeFilesetsResponse{
		Body: &nasclient.DescribeFilesetsResponseBody{
			Entries: &nasclient.DescribeFilesetsResponseBodyEntries{
				Entrie: []*nasclient.DescribeFilesetsResponseBodyEntriesEntrie{
					expectedFileSet1,
				},
			},
			NextToken: tea.String("next-token"),
		},
	}

	// Second page
	mockResponse2 := &nasclient.DescribeFilesetsResponse{
		Body: &nasclient.DescribeFilesetsResponseBody{
			Entries: &nasclient.DescribeFilesetsResponseBodyEntries{
				Entrie: []*nasclient.DescribeFilesetsResponseBodyEntriesEntrie{
					expectedFileSet2,
				},
			},
			NextToken: tea.String(""),
		},
	}

	callCount := 0
	mockClient.DescribeFilesetsFunc = func(request *nasclient.DescribeFilesetsRequest) (*nasclient.DescribeFilesetsResponse, error) {
		assert.Equal(t, fsID, *request.FileSystemId)
		callCount++
		if callCount == 1 {
			return mockResponse1, nil
		}
		return mockResponse2, nil
	}

	fileSets, err := manager.DescribeFileSets(ctx, fsID)

	require.NoError(t, err)
	assert.Len(t, fileSets, 2)
	assert.Equal(t, expectedFileSet1, fileSets[0])
	assert.Equal(t, expectedFileSet2, fileSets[1])
}

func TestCPFSFileSetManager_DescribeFileSets_Error(t *testing.T) {
	mockClient := NewMockNasClient()
	manager := NewCPFSFileSetManager(mockClient)

	ctx := context.Background()
	fsID := "fs-123"
	expectedError := errors.New("describe failed")

	mockClient.DescribeFilesetsFunc = func(request *nasclient.DescribeFilesetsRequest) (*nasclient.DescribeFilesetsResponse, error) {
		return nil, expectedError
	}

	fileSets, err := manager.DescribeFileSets(ctx, fsID)

	require.Error(t, err)
	assert.Nil(t, fileSets)
	assert.Contains(t, err.Error(), "describe filesets")
	assert.Contains(t, err.Error(), expectedError.Error())
}

func TestCPFSFileSetManager_DescribeFileSets_EmptyResponse(t *testing.T) {
	mockClient := NewMockNasClient()
	manager := NewCPFSFileSetManager(mockClient)

	ctx := context.Background()
	fsID := "fs-123"

	mockResponse := &nasclient.DescribeFilesetsResponse{
		Body: &nasclient.DescribeFilesetsResponseBody{},
	}

	mockClient.DescribeFilesetsFunc = func(request *nasclient.DescribeFilesetsRequest) (*nasclient.DescribeFilesetsResponse, error) {
		return mockResponse, nil
	}

	fileSets, err := manager.DescribeFileSets(ctx, fsID)

	require.NoError(t, err)
	assert.Empty(t, fileSets)
}

func TestCPFSFileSetManager_DescribeFileset_Success(t *testing.T) {
	mockClient := NewMockNasClient()
	manager := NewCPFSFileSetManager(mockClient)

	ctx := context.Background()
	fsID := "fs-123"
	fileSetID := "fset-456"
	expectedFileSet := &nasclient.DescribeFilesetsResponseBodyEntriesEntrie{
		FsetId: tea.String(fileSetID),
	}

	mockResponse := &nasclient.DescribeFilesetsResponse{
		Body: &nasclient.DescribeFilesetsResponseBody{
			Entries: &nasclient.DescribeFilesetsResponseBodyEntries{
				Entrie: []*nasclient.DescribeFilesetsResponseBodyEntriesEntrie{
					expectedFileSet,
				},
			},
		},
	}

	mockClient.DescribeFilesetsFunc = func(request *nasclient.DescribeFilesetsRequest) (*nasclient.DescribeFilesetsResponse, error) {
		assert.Equal(t, fsID, *request.FileSystemId)
		assert.Len(t, request.Filters, 1)
		assert.Equal(t, "FsetIds", *request.Filters[0].Key)
		assert.Equal(t, fileSetID, *request.Filters[0].Value)
		return mockResponse, nil
	}

	fileSet, err := manager.DescribeFileset(ctx, fsID, fileSetID)

	require.NoError(t, err)
	assert.Equal(t, expectedFileSet, fileSet)
}

func TestCPFSFileSetManager_DescribeFileset_NotFound(t *testing.T) {
	mockClient := NewMockNasClient()
	manager := NewCPFSFileSetManager(mockClient)

	ctx := context.Background()
	fsID := "fs-123"
	fileSetID := "fset-456"

	mockResponse := &nasclient.DescribeFilesetsResponse{
		Body: &nasclient.DescribeFilesetsResponseBody{
			Entries: &nasclient.DescribeFilesetsResponseBodyEntries{},
		},
	}

	mockClient.DescribeFilesetsFunc = func(request *nasclient.DescribeFilesetsRequest) (*nasclient.DescribeFilesetsResponse, error) {
		return mockResponse, nil
	}

	fileSet, err := manager.DescribeFileset(ctx, fsID, fileSetID)

	require.Error(t, err)
	assert.Nil(t, fileSet)
	assert.Contains(t, err.Error(), "not found")
}

func TestCPFSFileSetManager_DescribeFileset_Error(t *testing.T) {
	mockClient := NewMockNasClient()
	manager := NewCPFSFileSetManager(mockClient)

	ctx := context.Background()
	fsID := "fs-123"
	fileSetID := "fset-456"
	expectedError := errors.New("describe failed")

	mockClient.DescribeFilesetsFunc = func(request *nasclient.DescribeFilesetsRequest) (*nasclient.DescribeFilesetsResponse, error) {
		return nil, expectedError
	}

	fileSet, err := manager.DescribeFileset(ctx, fsID, fileSetID)

	require.Error(t, err)
	assert.Nil(t, fileSet)
	assert.Contains(t, err.Error(), "describe fileset")
	assert.Contains(t, err.Error(), expectedError.Error())
}

func TestCPFSFileSetManager_DescribeFileset_EmptyResponse(t *testing.T) {
	mockClient := NewMockNasClient()
	manager := NewCPFSFileSetManager(mockClient)

	ctx := context.Background()
	fsID := "fs-123"
	fileSetID := "fset-456"

	mockClient.DescribeFilesetsFunc = func(request *nasclient.DescribeFilesetsRequest) (*nasclient.DescribeFilesetsResponse, error) {
		return nil, nil
	}

	fileSet, err := manager.DescribeFileset(ctx, fsID, fileSetID)

	require.Error(t, err)
	assert.Nil(t, fileSet)
	assert.Contains(t, err.Error(), "empty response")
}

// Test to verify that CPFSFileSetManager implements the CPFSFileSetManager interface
func TestCPFSFileSetManager_InterfaceCompliance(t *testing.T) {
	mockClient := NewMockNasClient()
	manager := NewCPFSFileSetManager(mockClient)

	// This test ensures that cpfsFileSetManager implements CPFSFileSetManager interface
	var _ CPFSFileSetManager = manager
}
