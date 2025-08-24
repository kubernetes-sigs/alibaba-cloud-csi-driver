package common

import (
	"context"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/klog/v2"
)

func WrapControllerServer(server csi.ControllerServer) csi.ControllerServer {
	return ControllerServerWithLog{ControllerServerWithValidator{ControllerServer: server}}
}

type ControllerServerWithLog struct {
	ControllerServerWithValidator
}

func (cs ControllerServerWithLog) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	logger := klog.FromContext(ctx)
	ctx = klog.NewContext(ctx, logger.WithValues("method", "CreateVolume", "name", req.Name))
	return logGRPC(cs.ControllerServerWithValidator.CreateVolume, ctx, req)
}

func (cs ControllerServerWithLog) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	logger := klog.FromContext(ctx)
	ctx = klog.NewContext(ctx, logger.WithValues("method", "DeleteVolume", "volumeID", req.VolumeId))
	return logGRPC(cs.ControllerServerWithValidator.DeleteVolume, ctx, req)
}

func (cs ControllerServerWithLog) ControllerPublishVolume(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	logger := klog.FromContext(ctx)
	ctx = klog.NewContext(ctx, logger.WithValues("method", "ControllerPublishVolume", "volumeID", req.VolumeId, "nodeID", req.NodeId))
	return logGRPC(cs.ControllerServerWithValidator.ControllerPublishVolume, ctx, req)
}

func (cs ControllerServerWithLog) ControllerUnpublishVolume(ctx context.Context, req *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
	logger := klog.FromContext(ctx)
	ctx = klog.NewContext(ctx, logger.WithValues("method", "ControllerUnpublishVolume", "volumeID", req.VolumeId, "nodeID", req.NodeId))
	return logGRPC(cs.ControllerServerWithValidator.ControllerUnpublishVolume, ctx, req)
}

func (cs ControllerServerWithLog) ValidateVolumeCapabilities(ctx context.Context, req *csi.ValidateVolumeCapabilitiesRequest) (*csi.ValidateVolumeCapabilitiesResponse, error) {
	logger := klog.FromContext(ctx)
	ctx = klog.NewContext(ctx, logger.WithValues("method", "ValidateVolumeCapabilities", "volumeID", req.VolumeId))
	return logGRPC(cs.ControllerServerWithValidator.ValidateVolumeCapabilities, ctx, req)
}

func (cs ControllerServerWithLog) CreateSnapshot(ctx context.Context, req *csi.CreateSnapshotRequest) (*csi.CreateSnapshotResponse, error) {
	logger := klog.FromContext(ctx)
	ctx = klog.NewContext(ctx, logger.WithValues("method", "CreateSnapshot", "name", req.Name, "source", req.SourceVolumeId))
	return logGRPC(cs.ControllerServerWithValidator.CreateSnapshot, ctx, req)
}

func (cs ControllerServerWithLog) DeleteSnapshot(ctx context.Context, req *csi.DeleteSnapshotRequest) (*csi.DeleteSnapshotResponse, error) {
	logger := klog.FromContext(ctx)
	ctx = klog.NewContext(ctx, logger.WithValues("method", "DeleteSnapshot", "SnapshotID", req.SnapshotId))
	return logGRPC(cs.ControllerServerWithValidator.DeleteSnapshot, ctx, req)
}

func (cs ControllerServerWithLog) ControllerExpandVolume(ctx context.Context, req *csi.ControllerExpandVolumeRequest) (*csi.ControllerExpandVolumeResponse, error) {
	logger := klog.FromContext(ctx)
	ctx = klog.NewContext(ctx, logger.WithValues("method", "ControllerExpandVolume", "volumeID", req.VolumeId))
	return logGRPC(cs.ControllerServerWithValidator.ControllerExpandVolume, ctx, req)
}

func (cs ControllerServerWithLog) ControllerGetVolume(ctx context.Context, req *csi.ControllerGetVolumeRequest) (*csi.ControllerGetVolumeResponse, error) {
	logger := klog.FromContext(ctx)
	ctx = klog.NewContext(ctx, logger.WithValues("method", "ControllerGetVolume", "volumeID", req.VolumeId))
	return logGRPC(cs.ControllerServerWithValidator.ControllerGetVolume, ctx, req)
}

func (cs ControllerServerWithLog) ControllerModifyVolume(ctx context.Context, req *csi.ControllerModifyVolumeRequest) (*csi.ControllerModifyVolumeResponse, error) {
	logger := klog.FromContext(ctx)
	ctx = klog.NewContext(ctx, logger.WithValues("method", "ControllerModifyVolume", "volumeID", req.VolumeId))
	return logGRPC(cs.ControllerServerWithValidator.ControllerModifyVolume, ctx, req)
}

type ControllerServerWithValidator struct {
	csi.ControllerServer
}

func (cs ControllerServerWithValidator) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	if len(req.Name) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Name is required")
	}
	if len(req.VolumeCapabilities) == 0 {
		return nil, status.Error(codes.InvalidArgument, "VolumeCapabilities is required")
	}
	return cs.ControllerServer.CreateVolume(ctx, req)
}

func (cs ControllerServerWithValidator) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	if len(req.VolumeId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "VolumeId is required")
	}
	return cs.ControllerServer.DeleteVolume(ctx, req)
}

func (cs ControllerServerWithValidator) ControllerPublishVolume(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	if len(req.VolumeId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "VolumeId is required")
	}
	if len(req.NodeId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "NodeId is required")
	}
	if req.VolumeCapability == nil {
		return nil, status.Error(codes.InvalidArgument, "VolumeCapability is required")
	}
	return cs.ControllerServer.ControllerPublishVolume(ctx, req)
}

func (cs ControllerServerWithValidator) ControllerUnpublishVolume(ctx context.Context, req *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
	if len(req.VolumeId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "VolumeId is required")
	}
	if len(req.NodeId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "NodeId is required")
	}
	return cs.ControllerServer.ControllerUnpublishVolume(ctx, req)
}

func (cs ControllerServerWithValidator) ValidateVolumeCapabilities(ctx context.Context, req *csi.ValidateVolumeCapabilitiesRequest) (*csi.ValidateVolumeCapabilitiesResponse, error) {
	if len(req.VolumeId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "VolumeId is required")
	}
	if len(req.VolumeCapabilities) == 0 {
		return nil, status.Error(codes.InvalidArgument, "VolumeCapabilities is required")
	}
	return cs.ControllerServer.ValidateVolumeCapabilities(ctx, req)
}

func (cs ControllerServerWithValidator) CreateSnapshot(ctx context.Context, req *csi.CreateSnapshotRequest) (*csi.CreateSnapshotResponse, error) {
	if len(req.SourceVolumeId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "SourceVolumeId is required")
	}
	if len(req.Name) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Name is required")
	}
	return cs.ControllerServer.CreateSnapshot(ctx, req)
}

func (cs ControllerServerWithValidator) DeleteSnapshot(ctx context.Context, req *csi.DeleteSnapshotRequest) (*csi.DeleteSnapshotResponse, error) {
	if len(req.SnapshotId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "SnapshotId is required")
	}
	return cs.ControllerServer.DeleteSnapshot(ctx, req)
}

func (cs ControllerServerWithValidator) ControllerExpandVolume(ctx context.Context, req *csi.ControllerExpandVolumeRequest) (*csi.ControllerExpandVolumeResponse, error) {
	if len(req.VolumeId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "VolumeId is required")
	}
	if req.CapacityRange == nil {
		return nil, status.Error(codes.InvalidArgument, "CapacityRange is required")
	}
	return cs.ControllerServer.ControllerExpandVolume(ctx, req)
}

func (cs ControllerServerWithValidator) ControllerGetVolume(ctx context.Context, req *csi.ControllerGetVolumeRequest) (*csi.ControllerGetVolumeResponse, error) {
	if len(req.VolumeId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "VolumeId is required")
	}
	return cs.ControllerServer.ControllerGetVolume(ctx, req)
}

func (cs ControllerServerWithValidator) ControllerModifyVolume(ctx context.Context, req *csi.ControllerModifyVolumeRequest) (*csi.ControllerModifyVolumeResponse, error) {
	if len(req.VolumeId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "VolumeId is required")
	}
	if req.MutableParameters == nil {
		return nil, status.Error(codes.InvalidArgument, "MutableParameters is required")
	}
	return cs.ControllerServer.ControllerModifyVolume(ctx, req)
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
