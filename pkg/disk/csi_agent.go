package disk

import (
	"context"

	"github.com/container-storage-interface/spec/lib/go/csi"
)

type CSIAgent struct {
	csi.UnimplementedNodeServer
}

func NewCSIAgent() *CSIAgent {
	return &CSIAgent{}
}

func (a *CSIAgent) NodeExpandVolume(ctx context.Context, req *csi.NodeExpandVolumeRequest) (*csi.NodeExpandVolumeResponse, error) {
	return localExpandVolume(ctx, req)
}
