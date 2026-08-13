//go:build !windows

package oss

import (
	"context"

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
	ns *nodeServer
}

func NewCSIAgent(m metadata.MetadataProvider, socketPath string) *CSIAgent {
	ns := &nodeServer{
		metadata:        m,
		locks:           utils.NewVolumeLocks(),
		rawMounter:      mountutils.NewWithoutSystemd(""),
		skipGlobalMount: utils.GetSkipGlobalMount(true),
		fusePodManagers: ossfpm.GetAllOSSFusePodManagers(utils.Config{}, m, nil, nil),
		ossfsPaths:      ossfpm.GetAllFuseMounterPaths(),
		// Socket injection is handled inside nodeServer.NodePublishVolume,
		// which injects mountProxySock into req.PublishContext when non-empty.
		mountProxySock: socketPath,
	}
	return &CSIAgent{
		ns: ns,
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

// NodePublishVolume delegates to nodeServer, which handles socket injection
// via ns.mountProxySock (set from the csi-agent --mount-proxy-sock flag).
// The injection was originally done here; it was moved into nodeServer.NodePublishVolume
// to unify the socket override logic for both csi-agent and csi-plugin binaries.
func (a *CSIAgent) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	return a.ns.NodePublishVolume(ctx, req)
}

func (a *CSIAgent) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	return a.ns.NodeUnpublishVolume(ctx, req)
}
