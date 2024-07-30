package common

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func WrapNodeServerWithValidator(server csi.NodeServer) csi.NodeServer {
	return &NodeServerWithValidator{NodeServer: server}
}

type NodeServerWithValidator struct {
	csi.NodeServer
}

func (s *NodeServerWithValidator) NodeStageVolume(ctx context.Context, req *csi.NodeStageVolumeRequest) (*csi.NodeStageVolumeResponse, error) {
	if len(req.VolumeId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "VolumeId is required")
	}
	if req.VolumeCapability == nil {
		return nil, status.Error(codes.InvalidArgument, "VolumeCapability is required")
	}
	if len(req.StagingTargetPath) == 0 {
		return nil, status.Error(codes.InvalidArgument, "StagingTargetPath is required")
	}
	ok, err := filepathContains(utils.KubeletRootDir, req.StagingTargetPath)
	if err != nil || !ok {
		return nil, status.Errorf(codes.InvalidArgument, "Staging path %q is not a subpath of %s", req.StagingTargetPath, utils.KubeletRootDir)
	}
	return s.NodeServer.NodeStageVolume(ctx, req)
}

func (s *NodeServerWithValidator) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	if len(req.VolumeId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "VolumeId is required")
	}
	if req.VolumeCapability == nil {
		return nil, status.Error(codes.InvalidArgument, "VolumeCapability is required")
	}
	if len(req.TargetPath) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "TargetPath is required")
	}
	ok, err := filepathContains(utils.KubeletRootDir, req.TargetPath)
	if err != nil || !ok {
		return nil, status.Errorf(codes.InvalidArgument, "Target path %q is not a subpath of %s", req.TargetPath, utils.KubeletRootDir)
	}
	return s.NodeServer.NodePublishVolume(ctx, req)
}

func (s *NodeServerWithValidator) NodeUnstageVolume(context context.Context, req *csi.NodeUnstageVolumeRequest) (*csi.NodeUnstageVolumeResponse, error) {
	if len(req.VolumeId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "VolumeId is required")
	}
	if len(req.StagingTargetPath) == 0 {
		return nil, status.Error(codes.InvalidArgument, "StagingTargetPath is required")
	}
	return s.NodeServer.NodeUnstageVolume(context, req)
}

func (s *NodeServerWithValidator) NodeUnpublishVolume(context context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	if len(req.VolumeId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "VolumeId is required")
	}
	if len(req.TargetPath) == 0 {
		return nil, status.Error(codes.InvalidArgument, "TargetPath is required")
	}
	return s.NodeServer.NodeUnpublishVolume(context, req)
}

func (s *NodeServerWithValidator) NodeGetVolumeStats(context context.Context, req *csi.NodeGetVolumeStatsRequest) (*csi.NodeGetVolumeStatsResponse, error) {
	if len(req.VolumeId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "VolumeId is required")
	}
	if len(req.VolumePath) == 0 {
		return nil, status.Error(codes.InvalidArgument, "VolumePath is required")
	}
	return s.NodeServer.NodeGetVolumeStats(context, req)
}

func (s *NodeServerWithValidator) NodeExpandVolume(context context.Context, req *csi.NodeExpandVolumeRequest) (*csi.NodeExpandVolumeResponse, error) {
	if len(req.VolumeId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "VolumeId is required")
	}
	if len(req.VolumePath) == 0 {
		return nil, status.Error(codes.InvalidArgument, "VolumePath is required")
	}
	return s.NodeServer.NodeExpandVolume(context, req)
}

func filepathContains(basePath, path string) (bool, error) {
	relPath, err := filepath.Rel(basePath, path)
	if err != nil {
		return false, err
	}
	return !strings.HasPrefix(relPath, ".."+string(os.PathSeparator)), nil
}

type GenericNodeServer struct {
	NodeID string
}

func (ns *GenericNodeServer) NodeGetInfo(ctx context.Context, req *csi.NodeGetInfoRequest) (*csi.NodeGetInfoResponse, error) {
	return &csi.NodeGetInfoResponse{
		NodeId: ns.NodeID,
	}, nil
}

func (ns *GenericNodeServer) NodeGetCapabilities(context context.Context, req *csi.NodeGetCapabilitiesRequest) (*csi.NodeGetCapabilitiesResponse, error) {
	return &csi.NodeGetCapabilitiesResponse{}, nil
}

func (*GenericNodeServer) NodeGetVolumeStats(ctx context.Context, req *csi.NodeGetVolumeStatsRequest) (*csi.NodeGetVolumeStatsResponse, error) {
	resp, err := utils.GetMetrics(req.VolumePath)
	if errors.Is(err, os.ErrNotExist) {
		return nil, status.Errorf(codes.NotFound, "VolumePath %s not found: %v", req.VolumePath, err)
	}
	return resp, err
}

func (*GenericNodeServer) NodeStageVolume(ctx context.Context, req *csi.NodeStageVolumeRequest) (*csi.NodeStageVolumeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (*GenericNodeServer) NodeUnstageVolume(ctx context.Context, req *csi.NodeUnstageVolumeRequest) (*csi.NodeUnstageVolumeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (*GenericNodeServer) NodeExpandVolume(ctx context.Context, req *csi.NodeExpandVolumeRequest) (*csi.NodeExpandVolumeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}
