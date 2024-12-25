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
		MountProxySocket: socketPath,
		CNFSGetter:       unsupportedCNFSGetter{},
	}
	ns := newNodeServer(config)
	return &CSIAgent{
		ns: ns,
	}
}

func (a *CSIAgent) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	return a.ns.NodePublishVolume(ctx, req)
}

type unsupportedCNFSGetter struct{}

func (g unsupportedCNFSGetter) GetCNFS(ctx context.Context, name string) (*cnfsv1beta1.ContainerNetworkFileSystem, error) {
	return nil, errors.New("CNFS not supported in csi-agent")
}
