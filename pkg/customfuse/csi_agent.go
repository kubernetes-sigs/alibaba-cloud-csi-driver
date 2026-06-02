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
	socketPath string
	ns         *nodeServer
}

func NewCSIAgent(socketPath string) *CSIAgent {
	ns := &nodeServer{
		locks:      utils.NewVolumeLocks(),
		rawMounter: mountutils.NewWithoutSystemd(""),
		socketPath: socketPath,
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
