package nas

import (
	"context"
	"errors"

	"github.com/container-storage-interface/spec/lib/go/csi"
	cnfsv1beta1 "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cnfs/v1beta1"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/nas/internal"
)

type CSIAgent struct {
	csi.UnimplementedNodeServer
	ns *nodeServer
}

func NewCSIAgent(socketPath string) *CSIAgent {
	config := &internal.NodeConfig{
		EnablePortCheck:   true,
		EnableVolumeStats: true,
		MountProxySocket:  socketPath,
		CNFSGetter:        unsupportedCNFSGetter{},
	}
	ns := newNodeServer(config)
	return &CSIAgent{
		ns: ns,
	}
}

func (a *CSIAgent) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	return a.ns.NodePublishVolume(ctx, req)
}

func (a *CSIAgent) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	return a.ns.NodeUnpublishVolume(ctx, req)
}

func (a *CSIAgent) NodeGetCapabilities(ctx context.Context, req *csi.NodeGetCapabilitiesRequest) (*csi.NodeGetCapabilitiesResponse, error) {
	return a.ns.NodeGetCapabilities(ctx, req)
}

type unsupportedCNFSGetter struct{}

func (g unsupportedCNFSGetter) GetCNFS(ctx context.Context, name string) (*cnfsv1beta1.ContainerNetworkFileSystem, error) {
	return nil, errors.New("CNFS not supported in csi-agent")
}
