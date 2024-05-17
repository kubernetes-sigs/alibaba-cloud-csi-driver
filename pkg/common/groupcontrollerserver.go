package common

import (
	"context"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func WrapGroupControllerServerWithValidator(server csi.GroupControllerServer) csi.GroupControllerServer {
	return &GroupControllerServerWithValidator{GroupControllerServer: server}
}

type GroupControllerServerWithValidator struct {
	csi.GroupControllerServer
}

func (gcs *GroupControllerServerWithValidator) CreateVolumeGroupSnapshot(context context.Context, req *csi.CreateVolumeGroupSnapshotRequest) (*csi.CreateVolumeGroupSnapshotResponse, error) {
	if len(req.Name) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Name is required")
	}
	if len(req.SourceVolumeIds) == 0 {
		return nil, status.Error(codes.InvalidArgument, "SourceVolumeIds is required")
	}
	return gcs.GroupControllerServer.CreateVolumeGroupSnapshot(context, req)
}

func (gcs *GroupControllerServerWithValidator) DeleteVolumeGroupSnapshot(context context.Context, req *csi.DeleteVolumeGroupSnapshotRequest) (*csi.DeleteVolumeGroupSnapshotResponse, error) {
	if len(req.GroupSnapshotId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "GroupSnapshotId is required")
	}
	if len(req.SnapshotIds) == 0 {
		return nil, status.Error(codes.InvalidArgument, "SnapshotIds is required")
	}
	return gcs.GroupControllerServer.DeleteVolumeGroupSnapshot(context, req)
}

func (gcs *GroupControllerServerWithValidator) GetVolumeGroupSnapshot(context context.Context, req *csi.GetVolumeGroupSnapshotRequest) (*csi.GetVolumeGroupSnapshotResponse, error) {
	if len(req.GroupSnapshotId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "GroupSnapshotId is required")
	}
	if len(req.SnapshotIds) == 0 {
		return nil, status.Error(codes.InvalidArgument, "SnapshotIds is required")
	}
	return gcs.GroupControllerServer.GetVolumeGroupSnapshot(context, req)
}
