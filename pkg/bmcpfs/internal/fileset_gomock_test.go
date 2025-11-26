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
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCPFSFileSetManager_CreateFileSet_WithGoMock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockNas := NewGoMockNasClient(ctrl)
	manager := NewCPFSFileSetManager(mockNas)

	fsID := "fs-123"
	pvName := "pv-demo"
	fileSystemPath := "/demo/path"
	countLimit := int64(50)
	sizeBytes := int64(2048)
	deleteProtection := true
	expectedFileSetID := "fset-abc"

	mockNas.EXPECT().
		CreateFileset(gomock.Any()).
		DoAndReturn(func(req *nasclient.CreateFilesetRequest) (*nasclient.CreateFilesetResponse, error) {
			require.Equal(t, fsID, tea.StringValue(req.FileSystemId))
			require.Equal(t, pvName, tea.StringValue(req.ClientToken))
			require.Equal(t, fileSystemPath, tea.StringValue(req.FileSystemPath))
			require.Equal(t, countLimit, tea.Int64Value(req.Quota.FileCountLimit))
			require.Equal(t, sizeBytes, tea.Int64Value(req.Quota.SizeLimit))
			require.Equal(t, deleteProtection, tea.BoolValue(req.DeletionProtection))
			return &nasclient.CreateFilesetResponse{
				Body: &nasclient.CreateFilesetResponseBody{
					FsetId: tea.String(expectedFileSetID),
				},
			}, nil
		})

	fileSetID, err := manager.CreateFileSet(context.Background(), fsID, pvName, fileSystemPath, countLimit, sizeBytes, deleteProtection)
	require.NoError(t, err)
	assert.Equal(t, expectedFileSetID, fileSetID)
}

func TestCPFSFileSetManager_DeleteFileSet_WithGoMock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockNas := NewGoMockNasClient(ctrl)
	manager := NewCPFSFileSetManager(mockNas)

	fsID := "fs-123"
	fileSetID := "fset-delete"

	mockNas.EXPECT().
		DeleteFileset(gomock.Any()).
		DoAndReturn(func(req *nasclient.DeleteFilesetRequest) (*nasclient.DeleteFilesetResponse, error) {
			require.Equal(t, fsID, tea.StringValue(req.FileSystemId))
			require.Equal(t, fileSetID, tea.StringValue(req.FsetId))
			return &nasclient.DeleteFilesetResponse{
				Body: &nasclient.DeleteFilesetResponseBody{},
			}, nil
		})

	err := manager.DeleteFileSet(context.Background(), fsID, fileSetID)
	require.NoError(t, err)
}

func TestCPFSFileSetManager_DescribeFilesetByFsetID_ErrorWithGoMock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockNas := NewGoMockNasClient(ctrl)
	manager := NewCPFSFileSetManager(mockNas)

	fsID := "fs-123"
	fileSetID := "fset-789"
	expectedErr := errors.New("describe failed")

	mockNas.EXPECT().
		DescribeFilesets(gomock.Any()).
		DoAndReturn(func(req *nasclient.DescribeFilesetsRequest) (*nasclient.DescribeFilesetsResponse, error) {
			require.Equal(t, fsID, tea.StringValue(req.FileSystemId))
			require.Len(t, req.Filters, 1)
			require.Equal(t, fileSetID, tea.StringValue(req.Filters[0].Value))
			return nil, expectedErr
		})

	fileset, err := manager.DescribeFilesetByFsetID(context.Background(), fsID, fileSetID)
	require.Error(t, err)
	assert.Nil(t, fileset)
	assert.Contains(t, err.Error(), expectedErr.Error())
}
