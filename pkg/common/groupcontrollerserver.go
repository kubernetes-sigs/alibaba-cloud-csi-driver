package common

import (
	"context"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/klog/v2"
)

func WrapGroupControllerServer(server csi.GroupControllerServer) csi.GroupControllerServer {
	return GroupControllerServerWithLog{GroupControllerServerWithValidator{GroupControllerServer: server}}
}

type GroupControllerServerWithLog struct {
	GroupControllerServerWithValidator
}

func (gcs GroupControllerServerWithLog) CreateVolumeGroupSnapshot(ctx context.Context, req *csi.CreateVolumeGroupSnapshotRequest) (*csi.CreateVolumeGroupSnapshotResponse, error) {
	logger := klog.FromContext(ctx)
	ctx = klog.NewContext(ctx, logger.WithValues("method", "CreateVolumeGroupSnapshot", "name", req.Name))
	return logGRPC(gcs.GroupControllerServerWithValidator.CreateVolumeGroupSnapshot, ctx, req)
}

func (gcs GroupControllerServerWithLog) DeleteVolumeGroupSnapshot(ctx context.Context, req *csi.DeleteVolumeGroupSnapshotRequest) (*csi.DeleteVolumeGroupSnapshotResponse, error) {
	logger := klog.FromContext(ctx)
	ctx = klog.NewContext(ctx, logger.WithValues("method", "DeleteVolumeGroupSnapshot", "name", req.GroupSnapshotId))
	return logGRPC(gcs.GroupControllerServerWithValidator.DeleteVolumeGroupSnapshot, ctx, req)
}

func (gcs GroupControllerServerWithLog) GetVolumeGroupSnapshot(ctx context.Context, req *csi.GetVolumeGroupSnapshotRequest) (*csi.GetVolumeGroupSnapshotResponse, error) {
	logger := klog.FromContext(ctx)
	ctx = klog.NewContext(ctx, logger.WithValues("method", "GetVolumeGroupSnapshot", "name", req.GroupSnapshotId))
	return logGRPC(gcs.GroupControllerServerWithValidator.GetVolumeGroupSnapshot, ctx, req)
}

type GroupControllerServerWithValidator struct {
	csi.GroupControllerServer
}

func (gcs GroupControllerServerWithValidator) CreateVolumeGroupSnapshot(context context.Context, req *csi.CreateVolumeGroupSnapshotRequest) (*csi.CreateVolumeGroupSnapshotResponse, error) {
	if len(req.Name) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Name is required")
	}
	if len(req.SourceVolumeIds) == 0 {
		return nil, status.Error(codes.InvalidArgument, "SourceVolumeIds is required")
	}
	return gcs.GroupControllerServer.CreateVolumeGroupSnapshot(context, req)
}

func (gcs GroupControllerServerWithValidator) DeleteVolumeGroupSnapshot(context context.Context, req *csi.DeleteVolumeGroupSnapshotRequest) (*csi.DeleteVolumeGroupSnapshotResponse, error) {
	if len(req.GroupSnapshotId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "GroupSnapshotId is required")
	}
	return gcs.GroupControllerServer.DeleteVolumeGroupSnapshot(context, req)
}

func (gcs GroupControllerServerWithValidator) GetVolumeGroupSnapshot(context context.Context, req *csi.GetVolumeGroupSnapshotRequest) (*csi.GetVolumeGroupSnapshotResponse, error) {
	if len(req.GroupSnapshotId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "GroupSnapshotId is required")
	}
	return gcs.GroupControllerServer.GetVolumeGroupSnapshot(context, req)
}

type GenericGroupControllerServer struct {
	csi.UnimplementedGroupControllerServer
}

func GroupControllerRPCCapabilities(capsRPC ...csi.GroupControllerServiceCapability_RPC_Type) []*csi.GroupControllerServiceCapability {
	caps := make([]*csi.GroupControllerServiceCapability, 0, len(capsRPC))
	for _, c := range capsRPC {
		caps = append(caps, &csi.GroupControllerServiceCapability{
			Type: &csi.GroupControllerServiceCapability_Rpc{
				Rpc: &csi.GroupControllerServiceCapability_RPC{
					Type: c,
				},
			},
		})
	}
	return caps
}
