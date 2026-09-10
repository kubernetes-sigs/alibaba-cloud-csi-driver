//go:build !windows

package customfuse

import (
	"context"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	mountutils "k8s.io/mount-utils"
)

// CSIAgent serves customfuse volumes in csi-agent mode (sandbox scenario only).
// ACS GPU scenario is NOT supported — csi-agent cannot distinguish sandbox from ACS
// since both resolve to the same runtime type.
type CSIAgent struct {
	csi.UnimplementedNodeServer
	ns *nodeServer
}

func NewCSIAgent(socketPath string) *CSIAgent {
	ns := &nodeServer{
		locks:           utils.NewVolumeLocks(),
		rawMounter:      mountutils.NewWithoutSystemd(""),
		skipGlobalMount: utils.GetSkipGlobalMount(true),
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
// via ns.mountProxySock. The injection was originally done here; it was moved
// into nodeServer.NodePublishVolume to unify the logic for csi-agent and csi-plugin.
func (a *CSIAgent) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	return a.ns.NodePublishVolume(ctx, req)
}

func (a *CSIAgent) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	return a.ns.NodeUnpublishVolume(ctx, req)
}
