package common

import (
	"context"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func WrapControllerServerWithValidator(server csi.ControllerServer) csi.ControllerServer {
	return &ControllerServerWithValidator{ControllerServer: server}
}

type ControllerServerWithValidator struct {
	csi.ControllerServer
}

func (cs *ControllerServerWithValidator) CreateVolume(context context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	if len(req.Name) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Name is required")
	}
	if len(req.VolumeCapabilities) == 0 {
		return nil, status.Error(codes.InvalidArgument, "VolumeCapabilities is required")
	}
	return cs.ControllerServer.CreateVolume(context, req)
}

func (cs *ControllerServerWithValidator) DeleteVolume(context context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	if len(req.VolumeId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "VolumeId is required")
	}
	return cs.ControllerServer.DeleteVolume(context, req)
}

func (cs *ControllerServerWithValidator) ControllerPublishVolume(context context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	if len(req.VolumeId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "VolumeId is required")
	}
	if len(req.NodeId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "NodeId is required")
	}
	if req.VolumeCapability == nil {
		return nil, status.Error(codes.InvalidArgument, "VolumeCapability is required")
	}
	return cs.ControllerServer.ControllerPublishVolume(context, req)
}

func (cs *ControllerServerWithValidator) ControllerUnpublishVolume(context context.Context, req *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
	if len(req.VolumeId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "VolumeId is required")
	}
	if len(req.NodeId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "NodeId is required")
	}
	return cs.ControllerServer.ControllerUnpublishVolume(context, req)
}

func (cs *ControllerServerWithValidator) ValidateVolumeCapabilities(context context.Context, req *csi.ValidateVolumeCapabilitiesRequest) (*csi.ValidateVolumeCapabilitiesResponse, error) {
	if len(req.VolumeId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "VolumeId is required")
	}
	if len(req.VolumeCapabilities) == 0 {
		return nil, status.Error(codes.InvalidArgument, "VolumeCapabilities is required")
	}
	return cs.ControllerServer.ValidateVolumeCapabilities(context, req)
}

func (cs *ControllerServerWithValidator) CreateSnapshot(context context.Context, req *csi.CreateSnapshotRequest) (*csi.CreateSnapshotResponse, error) {
	if len(req.SourceVolumeId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "SourceVolumeId is required")
	}
	if len(req.Name) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Name is required")
	}
	return cs.ControllerServer.CreateSnapshot(context, req)
}

func (cs *ControllerServerWithValidator) DeleteSnapshot(context context.Context, req *csi.DeleteSnapshotRequest) (*csi.DeleteSnapshotResponse, error) {
	if len(req.SnapshotId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "SnapshotId is required")
	}
	return cs.ControllerServer.DeleteSnapshot(context, req)
}

func (cs *ControllerServerWithValidator) ControllerExpandVolume(context context.Context, req *csi.ControllerExpandVolumeRequest) (*csi.ControllerExpandVolumeResponse, error) {
	if len(req.VolumeId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "VolumeId is required")
	}
	if req.CapacityRange == nil {
		return nil, status.Error(codes.InvalidArgument, "CapacityRange is required")
	}
	return cs.ControllerServer.ControllerExpandVolume(context, req)
}
