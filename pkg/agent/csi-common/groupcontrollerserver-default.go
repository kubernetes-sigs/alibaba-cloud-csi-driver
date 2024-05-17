package csicommon

import (
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/golang/glog"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DefaultGroupControllerServer struct {
	Driver *CSIDriver
}

// GroupControllerGetCapabilities implements the default GRPC callout.
// Default supports all capabilities
func (gcs *DefaultGroupControllerServer) GroupControllerGetCapabilities(ctx context.Context, req *csi.GroupControllerGetCapabilitiesRequest) (*csi.GroupControllerGetCapabilitiesResponse, error) {
	glog.V(5).Infof("Using default GroupControllerGetCapabilities")

	return &csi.GroupControllerGetCapabilitiesResponse{
		Capabilities: gcs.Driver.gCap,
	}, nil
}

func (gcs *DefaultGroupControllerServer) CreateVolumeGroupSnapshot(ctx context.Context, req *csi.CreateVolumeGroupSnapshotRequest) (*csi.CreateVolumeGroupSnapshotResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (gcs *DefaultGroupControllerServer) DeleteVolumeGroupSnapshot(ctx context.Context, req *csi.DeleteVolumeGroupSnapshotRequest) (*csi.DeleteVolumeGroupSnapshotResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (gcs *DefaultGroupControllerServer) GetVolumeGroupSnapshot(ctx context.Context, req *csi.GetVolumeGroupSnapshotRequest) (*csi.GetVolumeGroupSnapshotResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}
