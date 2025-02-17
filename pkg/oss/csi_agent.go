package oss

import (
	"context"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
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
		metadata:   m,
		locks:      utils.NewVolumeLocks(),
		rawMounter: mountutils.NewWithoutSystemd(""),
		skipAttach: true,
	}
	return &CSIAgent{
		ns:         ns,
		socketPath: socketPath,
	}
}

func (a *CSIAgent) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	if req.PublishContext == nil {
		req.PublishContext = map[string]string{}
	}
	req.PublishContext[mountProxySocket] = a.socketPath
	return a.ns.NodePublishVolume(ctx, req)
}
