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

package nas

import (
	"context"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/nas/internal"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"
)

const defaultVolumeAs = "subpath"

func init() {
	internal.RegisterControllerMode(newSubpathController)
	internal.RegisterControllerMode(newSharepathController)
	internal.RegisterControllerMode(newFilesystemController)
	internal.RegisterControllerMode(newAccesspointController)
}

type controllerServer struct {
	*internal.ControllerFactory
	kubeClient kubernetes.Interface
	locks      *utils.VolumeLocks
	common.GenericControllerServer
}

func newControllerServer(config *internal.ControllerConfig) (*controllerServer, error) {
	fac, err := internal.NewControllerFactory(config, defaultVolumeAs)
	if err != nil {
		return nil, err
	}
	c := &controllerServer{
		ControllerFactory: fac,
		kubeClient:        config.KubeClient,
		locks:             utils.NewVolumeLocks(),
	}
	return c, nil
}

func (cs *controllerServer) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	klog.V(2).InfoS("CreateVolume: starting", "request", req)
	if !cs.locks.TryAcquire(req.Name) {
		return nil, status.Errorf(codes.Aborted, "There is already an operation for volume %s", req.Name)
	}
	defer cs.locks.Release(req.Name)

	controller, err := cs.VolumeAs(req.Parameters["volumeAs"])
	if err != nil {
		return nil, err
	}
	resp, err := controller.CreateVolume(ctx, req)
	if err != nil {
		return nil, err
	}

	parameters := req.Parameters
	resp.Volume.VolumeContext["volumeAs"] = controller.VolumeAs()
	if mode := parameters["mode"]; mode != "" {
		resp.Volume.VolumeContext["mode"] = mode
		modeType := parameters["modeType"]
		if modeType == "" {
			modeType = "non-recursive"
		}
		resp.Volume.VolumeContext["modeType"] = modeType
	}
	if options := parameters["options"]; options != "" {
		resp.Volume.VolumeContext["options"] = options
	}

	klog.V(2).InfoS("CreateVolume: succeeded", "response", resp)
	return resp, err
}

func (cs *controllerServer) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	klog.Infof("DeleteVolume: Starting deleting volume %s", req.GetVolumeId())
	if !cs.locks.TryAcquire(req.VolumeId) {
		return nil, status.Errorf(codes.Aborted, "There is already an operation for volume %s", req.VolumeId)
	}
	defer cs.locks.Release(req.VolumeId)

	pv, err := cs.kubeClient.CoreV1().PersistentVolumes().Get(ctx, req.VolumeId, metav1.GetOptions{})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	controller, err := cs.VolumeAs(pv.Spec.CSI.VolumeAttributes["volumeAs"])
	if err != nil {
		return nil, err
	}
	resp, err := controller.DeleteVolume(ctx, req, pv)
	if err == nil {
		klog.V(2).InfoS("DeleteVolume: succeeded", "volumeId", req.VolumeId)
	}
	return resp, err
}

func (cs *controllerServer) ControllerExpandVolume(ctx context.Context, req *csi.ControllerExpandVolumeRequest) (*csi.ControllerExpandVolumeResponse, error) {
	klog.Infof("ControllerExpandVolume: starting to expand nas volume with %v", req)
	if !cs.locks.TryAcquire(req.VolumeId) {
		return nil, status.Errorf(codes.Aborted, "There is already an operation for volume %s", req.VolumeId)
	}
	defer cs.locks.Release(req.VolumeId)

	pv, err := cs.kubeClient.CoreV1().PersistentVolumes().Get(ctx, req.VolumeId, metav1.GetOptions{})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	controller, err := cs.VolumeAs(pv.Spec.CSI.VolumeAttributes["volumeAs"])
	if err != nil {
		return nil, err
	}
	resp, err := controller.ControllerExpandVolume(ctx, req, pv)
	if err == nil {
		klog.V(2).InfoS("ControllerExpandVolume: succeeded", "response", resp)
	}
	return resp, err
}

func (cs *controllerServer) ValidateVolumeCapabilities(ctx context.Context, req *csi.ValidateVolumeCapabilitiesRequest) (*csi.ValidateVolumeCapabilitiesResponse, error) {
	for _, cap := range req.VolumeCapabilities {
		if cap.GetAccessMode().GetMode() != csi.VolumeCapability_AccessMode_MULTI_NODE_MULTI_WRITER {
			return &csi.ValidateVolumeCapabilitiesResponse{Message: ""}, nil
		}
	}
	return &csi.ValidateVolumeCapabilitiesResponse{
		Confirmed: &csi.ValidateVolumeCapabilitiesResponse_Confirmed{
			VolumeCapabilities: req.VolumeCapabilities,
		},
	}, nil
}

func (cs *controllerServer) ControllerGetCapabilities(ctx context.Context, req *csi.ControllerGetCapabilitiesRequest) (*csi.ControllerGetCapabilitiesResponse, error) {
	return &csi.ControllerGetCapabilitiesResponse{
		Capabilities: common.ControllerRPCCapabilities(
			csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME,
			csi.ControllerServiceCapability_RPC_EXPAND_VOLUME,
		)}, nil
}

func (cs *controllerServer) ControllerPublishVolume(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (cs *controllerServer) ControllerUnpublishVolume(ctx context.Context, req *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (cs *controllerServer) ListVolumes(ctx context.Context, req *csi.ListVolumesRequest) (*csi.ListVolumesResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (cs *controllerServer) GetCapacity(ctx context.Context, req *csi.GetCapacityRequest) (*csi.GetCapacityResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}
func (cs *controllerServer) CreateSnapshot(ctx context.Context, req *csi.CreateSnapshotRequest) (*csi.CreateSnapshotResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (cs *controllerServer) DeleteSnapshot(ctx context.Context, req *csi.DeleteSnapshotRequest) (*csi.DeleteSnapshotResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (cs *controllerServer) ListSnapshots(ctx context.Context, req *csi.ListSnapshotsRequest) (*csi.ListSnapshotsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (cs *controllerServer) ControllerGetVolume(ctx context.Context, req *csi.ControllerGetVolumeRequest,
) (*csi.ControllerGetVolumeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (cs *controllerServer) ControllerModifyVolume(ctx context.Context, req *csi.ControllerModifyVolumeRequest,
) (*csi.ControllerModifyVolumeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}
