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
	"strings"
	"time"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/klog/v2"
)

type controllerServer struct {
	common.GenericControllerServer
	attachBackoff wait.Backoff
	newDFSClient  func() (DFSClient, error)
}

func newControllerServer(region string) *controllerServer {
	return &controllerServer{
		attachBackoff: wait.Backoff{
			Duration: time.Millisecond * 400,
			Factor:   2,
			Steps:    6,
		},
		newDFSClient: func() (DFSClient, error) {
			client, err := newDFSClient(region)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "Failed to init dfs client: %v", err)
			}
			return client, nil
		},
	}
}

func (c *controllerServer) ControllerPublishVolume(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	fileSystemId, mountPointId, err := parseVolumeId(req.VolumeId)
	if err != nil {
		return nil, err
	}
	dfsClient, err := c.newDFSClient()
	if err != nil {
		return nil, err
	}

	// Attach
	err = dfsClient.Attach(fileSystemId, mountPointId, req.NodeId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Wait for the VSC to become NORMAL
	var vsc *VSC
	err = wait.ExponentialBackoffWithContext(ctx, c.attachBackoff, func(_ context.Context) (done bool, err error) {
		vsc, err = dfsClient.DescribeAttachment(fileSystemId, mountPointId, req.NodeId)
		if err != nil {
			return false, status.Error(codes.Internal, err.Error())
		}
		if vsc == nil {
			return false, status.Error(codes.Internal, "VSC not found after AttachVscMountPoint")
		}
		klog.InfoS("Waiting for the VSC status", "nodeId", req.NodeId, "volumeId", req.VolumeId, "status", vsc.VscStatus)
		switch vsc.VscStatus {
		case VSC_NORMARL:
			return true, nil
		case VSC_INVALID:
			return false, status.Error(codes.Internal, "VSC becomes INVALID")
		default:
			return false, nil
		}
	})
	if err != nil {
		if wait.Interrupted(err) {
			return nil, status.Error(codes.DeadlineExceeded, "Failed to wait for the VSC to become NORMAL")
		}
		return nil, err
	}

	klog.InfoS("ControllerPublishVolume succeeded", "nodeId", req.NodeId, "volumeId", req.VolumeId, "vscId", vsc.VscId)
	return &csi.ControllerPublishVolumeResponse{
		PublishContext: map[string]string{
			"vscId": vsc.VscId,
		},
	}, nil
}

func (c *controllerServer) ControllerUnpublishVolume(_ context.Context, req *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
	fileSystemId, mountPointId, err := parseVolumeId(req.VolumeId)
	if err != nil {
		return nil, err
	}
	dfsClient, err := c.newDFSClient()
	if err != nil {
		return nil, err
	}

	// Detach
	err = dfsClient.Detach(fileSystemId, mountPointId, req.NodeId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	klog.InfoS("ControllerUnpublishVolume succeeded", "nodeId", req.NodeId, "volumeId", req.VolumeId)
	return &csi.ControllerUnpublishVolumeResponse{}, nil
}

func (c *controllerServer) ValidateVolumeCapabilities(_ context.Context, req *csi.ValidateVolumeCapabilitiesRequest) (*csi.ValidateVolumeCapabilitiesResponse, error) {
	return &csi.ValidateVolumeCapabilitiesResponse{
		Confirmed: &csi.ValidateVolumeCapabilitiesResponse_Confirmed{
			VolumeCapabilities: req.VolumeCapabilities,
		},
	}, nil
}

func (c *controllerServer) ControllerGetCapabilities(_ context.Context, _ *csi.ControllerGetCapabilitiesRequest) (*csi.ControllerGetCapabilitiesResponse, error) {
	return &csi.ControllerGetCapabilitiesResponse{
		Capabilities: common.ControllerRPCCapabilities(
			csi.ControllerServiceCapability_RPC_PUBLISH_UNPUBLISH_VOLUME,
		),
	}, nil
}

func parseVolumeId(volumeId string) (string, string, error) {
	fileSystemId, mountPointId, found := strings.Cut(volumeId, "/")
	if !found {
		return "", "", status.Error(codes.InvalidArgument, "Invalid volume id")
	}
	if fileSystemId == "" || mountPointId == "" {
		return "", "", status.Error(codes.InvalidArgument, "Invalid volume id")
	}
	return fileSystemId, mountPointId, nil
}
