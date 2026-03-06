package oss

import (
	"context"
	"os"
	"strconv"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	ossfpm "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/fuse_pod_manager/oss"
	_ "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/fuse_pod_manager/oss/ossfs"
	_ "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/fuse_pod_manager/oss/ossfs2"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	mountutils "k8s.io/mount-utils"
)

type CSIAgent struct {
	csi.UnimplementedNodeServer
	// mount-proxy socket path
	socketPath string
	ns         *nodeServer
}

func NewCSIAgent(m metadata.MetadataProvider, socketPath string) *CSIAgent {
	ns := &nodeServer{
		metadata:        m,
		locks:           utils.NewVolumeLocks(),
		rawMounter:      mountutils.NewWithoutSystemd(""),
		skipGlobalMount: true, // label for csi-agent environment
		fusePodManagers: ossfpm.GetAllOSSFusePodManagers(utils.Config{}, m, nil),
		ossfsPaths:      ossfpm.GetAllFuseMounterPaths(),
	}
	if value := os.Getenv("OSS_SKIP_GLOBAL_MOUNT"); value != "" {
		ns.skipGlobalMount, _ = strconv.ParseBool(value)
	}
	return &CSIAgent{
		ns:         ns,
		socketPath: socketPath,
	}
}

func (a *CSIAgent) NodeGetCapabilities(ctx context.Context, req *csi.NodeGetCapabilitiesRequest) (*csi.NodeGetCapabilitiesResponse, error) {
	return a.ns.NodeGetCapabilities(ctx, req)
}

func (a *CSIAgent) NodeStageVolume(ctx context.Context, req *csi.NodeStageVolumeRequest) (*csi.NodeStageVolumeResponse, error) {
	return a.ns.NodeStageVolume(ctx, req)
}

func (a *CSIAgent) NodeUnstageVolume(ctx context.Context, req *csi.NodeUnstageVolumeRequest) (*csi.NodeUnstageVolumeResponse, error) {
	return a.ns.NodeUnstageVolume(ctx, req)
}

func (a *CSIAgent) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	if req.PublishContext == nil {
		req.PublishContext = map[string]string{}
	}
	req.PublishContext[mountProxySocket] = a.socketPath
	return a.ns.NodePublishVolume(ctx, req)
}

func (a *CSIAgent) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	return a.ns.NodeUnpublishVolume(ctx, req)
}
