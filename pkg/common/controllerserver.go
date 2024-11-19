package common

import (
	"context"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/klog/v2"
)

func WrapControllerServerWithValidator(server csi.ControllerServer) csi.ControllerServer {
	return &ControllerServerWithValidator{ControllerServer: server}
}

type ControllerServerWithValidator struct {
	csi.ControllerServer
}

func (cs *ControllerServerWithValidator) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	if len(req.Name) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Name is required")
	}
	if len(req.VolumeCapabilities) == 0 {
		return nil, status.Error(codes.InvalidArgument, "VolumeCapabilities is required")
	}
	logger := klog.FromContext(ctx)
	ctx = klog.NewContext(ctx, logger.WithValues("name", req.Name))
	return cs.ControllerServer.CreateVolume(ctx, req)
}

func (cs *ControllerServerWithValidator) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	if len(req.VolumeId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "VolumeId is required")
	}
	logger := klog.FromContext(ctx)
	ctx = klog.NewContext(ctx, logger.WithValues("volumeID", req.VolumeId))
	return cs.ControllerServer.DeleteVolume(ctx, req)
}

func (cs *ControllerServerWithValidator) ControllerPublishVolume(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	if len(req.VolumeId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "VolumeId is required")
	}
	if len(req.NodeId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "NodeId is required")
	}
	if req.VolumeCapability == nil {
		return nil, status.Error(codes.InvalidArgument, "VolumeCapability is required")
	}
	logger := klog.FromContext(ctx)
	ctx = klog.NewContext(ctx, logger.WithValues("volumeID", req.VolumeId, "nodeID", req.NodeId))
	return cs.ControllerServer.ControllerPublishVolume(ctx, req)
}

func (cs *ControllerServerWithValidator) ControllerUnpublishVolume(ctx context.Context, req *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
	if len(req.VolumeId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "VolumeId is required")
	}
	if len(req.NodeId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "NodeId is required")
	}
	logger := klog.FromContext(ctx)
	ctx = klog.NewContext(ctx, logger.WithValues("volumeID", req.VolumeId, "nodeID", req.NodeId))
	return cs.ControllerServer.ControllerUnpublishVolume(ctx, req)
}

func (cs *ControllerServerWithValidator) ValidateVolumeCapabilities(ctx context.Context, req *csi.ValidateVolumeCapabilitiesRequest) (*csi.ValidateVolumeCapabilitiesResponse, error) {
	if len(req.VolumeId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "VolumeId is required")
	}
	if len(req.VolumeCapabilities) == 0 {
		return nil, status.Error(codes.InvalidArgument, "VolumeCapabilities is required")
	}
	logger := klog.FromContext(ctx)
	ctx = klog.NewContext(ctx, logger.WithValues("volumeID", req.VolumeId))
	return cs.ControllerServer.ValidateVolumeCapabilities(ctx, req)
}

func (cs *ControllerServerWithValidator) CreateSnapshot(ctx context.Context, req *csi.CreateSnapshotRequest) (*csi.CreateSnapshotResponse, error) {
	if len(req.SourceVolumeId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "SourceVolumeId is required")
	}
	if len(req.Name) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Name is required")
	}
	logger := klog.FromContext(ctx)
	ctx = klog.NewContext(ctx, logger.WithValues("name", req.Name, "source", req.SourceVolumeId))
	return cs.ControllerServer.CreateSnapshot(ctx, req)
}

func (cs *ControllerServerWithValidator) DeleteSnapshot(ctx context.Context, req *csi.DeleteSnapshotRequest) (*csi.DeleteSnapshotResponse, error) {
	if len(req.SnapshotId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "SnapshotId is required")
	}
	logger := klog.FromContext(ctx)
	ctx = klog.NewContext(ctx, logger.WithValues("SnapshotID", req.SnapshotId))
	return cs.ControllerServer.DeleteSnapshot(ctx, req)
}

func (cs *ControllerServerWithValidator) ControllerExpandVolume(ctx context.Context, req *csi.ControllerExpandVolumeRequest) (*csi.ControllerExpandVolumeResponse, error) {
	if len(req.VolumeId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "VolumeId is required")
	}
	if req.CapacityRange == nil {
		return nil, status.Error(codes.InvalidArgument, "CapacityRange is required")
	}
	logger := klog.FromContext(ctx)
	ctx = klog.NewContext(ctx, logger.WithValues("volumeID", req.VolumeId))
	return cs.ControllerServer.ControllerExpandVolume(ctx, req)
}

type GenericControllerServer struct {
	csi.UnimplementedControllerServer
}

func ControllerRPCCapabilities(capsRPC ...csi.ControllerServiceCapability_RPC_Type) []*csi.ControllerServiceCapability {
	caps := make([]*csi.ControllerServiceCapability, 0, len(capsRPC))
	for _, c := range capsRPC {
		caps = append(caps, &csi.ControllerServiceCapability{
			Type: &csi.ControllerServiceCapability_Rpc{
				Rpc: &csi.ControllerServiceCapability_RPC{
					Type: c,
				},
			},
		})
	}
	return caps
}
