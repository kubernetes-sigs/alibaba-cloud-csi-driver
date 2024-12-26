/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http:// www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package dfs

import (
	"context"
	"errors"
	"testing"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/util/wait"
)

func newControllerServerWithMock(t *testing.T) (*controllerServer, *MockDFSClient) {
	t.Helper()
	mockCtrl := gomock.NewController(t)
	mockDFSClient := NewMockDFSClient(mockCtrl)
	return &controllerServer{
		attachBackoff: wait.Backoff{
			Steps: 2,
		},
		newDFSClient: func() (DFSClient, error) {
			return mockDFSClient, nil
		},
	}, mockDFSClient
}

func TestControllerPublishVolume(t *testing.T) {
	t.Run("Succeed to attach", func(t *testing.T) {
		c, mockClient := newControllerServerWithMock(t)
		req := &csi.ControllerPublishVolumeRequest{
			VolumeId: "xxx-xx-xx-xx-xxx/xxx-xxx.dfs.aliyuncs.com",
			NodeId:   "node0",
		}
		fileSystemId, mountPointId, _ := parseVolumeId(req.VolumeId)
		mockClient.EXPECT().Attach(gomock.Eq(fileSystemId), gomock.Eq(mountPointId), gomock.Eq(req.NodeId)).Return(nil)
		mockClient.EXPECT().DescribeAttachment(gomock.Eq(fileSystemId), gomock.Eq(mountPointId), gomock.Eq(req.NodeId)).Return(&VSC{
			VscId:     "vsc1",
			VscType:   VSCType_Primary,
			VscStatus: VSC_CREATING,
		}, nil)

		mockClient.EXPECT().DescribeAttachment(gomock.Eq(fileSystemId), gomock.Eq(mountPointId), gomock.Eq(req.NodeId)).Return(&VSC{
			VscId:     "vsc1",
			VscType:   VSCType_Primary,
			VscStatus: VSC_NORMARL,
		}, nil)

		resp, err := c.ControllerPublishVolume(context.Background(), req)
		assert.NoError(t, err)
		assert.Equal(t, resp.PublishContext["vscId"], "vsc1")
	})

	t.Run("VSC turns to INVALID", func(t *testing.T) {
		c, mockClient := newControllerServerWithMock(t)
		req := &csi.ControllerPublishVolumeRequest{
			VolumeId: "xxx-xx-xx-xx-xxx/xxx-xxx.dfs.aliyuncs.com",
			NodeId:   "node0",
		}
		fileSystemId, mountPointId, _ := parseVolumeId(req.VolumeId)
		mockClient.EXPECT().Attach(gomock.Eq(fileSystemId), gomock.Eq(mountPointId), gomock.Eq(req.NodeId)).Return(nil)
		mockClient.EXPECT().DescribeAttachment(gomock.Eq(fileSystemId), gomock.Eq(mountPointId), gomock.Eq(req.NodeId)).Return(&VSC{
			VscId:     "vsc1",
			VscType:   VSCType_Primary,
			VscStatus: VSC_INVALID,
		}, nil)

		_, err := c.ControllerPublishVolume(context.Background(), req)
		assert.ErrorContains(t, err, "VSC becomes INVALID")
	})

	t.Run("Timed out", func(t *testing.T) {
		c, mockClient := newControllerServerWithMock(t)
		req := &csi.ControllerPublishVolumeRequest{
			VolumeId: "xxx-xx-xx-xx-xxx/xxx-xxx.dfs.aliyuncs.com",
			NodeId:   "node0",
		}
		fileSystemId, mountPointId, _ := parseVolumeId(req.VolumeId)
		mockClient.EXPECT().Attach(gomock.Eq(fileSystemId), gomock.Eq(mountPointId), gomock.Eq(req.NodeId)).Return(nil)
		mockClient.EXPECT().DescribeAttachment(gomock.Eq(fileSystemId), gomock.Eq(mountPointId), gomock.Eq(req.NodeId)).Return(&VSC{
			VscId:     "vsc1",
			VscType:   VSCType_Primary,
			VscStatus: VSC_CREATING,
		}, nil).Times(c.attachBackoff.Steps)

		_, err := c.ControllerPublishVolume(context.Background(), req)
		assert.ErrorContains(t, err, "Failed to wait for the VSC to become NORMAL")
	})
}

func TestControllerUnpublishVolume(t *testing.T) {
	t.Run("Succeed to detach", func(t *testing.T) {
		c, mockClient := newControllerServerWithMock(t)
		req := &csi.ControllerUnpublishVolumeRequest{
			VolumeId: "xxx-xx-xx-xx-xxx/xxx-xxx.dfs.aliyuncs.com",
			NodeId:   "node0",
		}
		fileSystemId, mountPointId, _ := parseVolumeId(req.VolumeId)
		mockClient.EXPECT().Detach(gomock.Eq(fileSystemId), gomock.Eq(mountPointId), gomock.Eq(req.NodeId)).Return(nil)

		_, err := c.ControllerUnpublishVolume(context.Background(), req)
		assert.NoError(t, err)
	})

	t.Run("Fail to detach", func(t *testing.T) {
		c, mockClient := newControllerServerWithMock(t)
		req := &csi.ControllerUnpublishVolumeRequest{
			VolumeId: "xxx-xx-xx-xx-xxx/xxx-xxx.dfs.aliyuncs.com",
			NodeId:   "node0",
		}
		fileSystemId, mountPointId, _ := parseVolumeId(req.VolumeId)
		mockClient.EXPECT().Detach(gomock.Eq(fileSystemId), gomock.Eq(mountPointId), gomock.Eq(req.NodeId)).Return(errors.New("what ever error"))
		_, err := c.ControllerUnpublishVolume(context.Background(), req)
		assert.Error(t, err)
	})
}
func Test_parseVolumeId(t *testing.T) {
	type args struct {
		volumeId string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   string
		wantErr bool
	}{
		{
			"valid",
			args{"xxx-xx-xx-xx-xxx/xxx-xxx.dfs.aliyuncs.com"},
			"xxx-xx-xx-xx-xxx",
			"xxx-xxx.dfs.aliyuncs.com",
			false,
		},
		{
			"missing mountpoint id",
			args{"xxx-xx-xx-xx-xxx"},
			"",
			"",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := parseVolumeId(tt.args.volumeId)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseVolumeId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseVolumeId() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("parseVolumeId() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
