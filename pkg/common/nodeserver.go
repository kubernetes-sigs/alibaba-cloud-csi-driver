package common

import (
	"context"
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
	if req.StagingTargetPath == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Staging path not provided")
	}
	ok, err := filepathContains(utils.KubeletRootDir, req.StagingTargetPath)
	if err != nil || !ok {
		return nil, status.Errorf(codes.InvalidArgument, "Staging path %q is not a subpath of %s", req.StagingTargetPath, utils.KubeletRootDir)
	}
	return s.NodeServer.NodeStageVolume(ctx, req)
}

func (s *NodeServerWithValidator) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	if req.TargetPath == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Target path not provided")
	}
	ok, err := filepathContains(utils.KubeletRootDir, req.TargetPath)
	if err != nil || !ok {
		return nil, status.Errorf(codes.InvalidArgument, "Target path %q is not a subpath of %s", req.TargetPath, utils.KubeletRootDir)
	}
	return s.NodeServer.NodePublishVolume(ctx, req)
}

func filepathContains(basePath, path string) (bool, error) {
	relPath, err := filepath.Rel(basePath, path)
	if err != nil {
		return false, err
	}
	return !strings.HasPrefix(relPath, ".."+string(os.PathSeparator)), nil
}
