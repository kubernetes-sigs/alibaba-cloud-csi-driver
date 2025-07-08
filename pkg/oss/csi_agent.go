package oss

import (
	"context"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/oss"
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

	ossfs := oss.NewFuseOssfs(nil, m)
	ossfs2 := oss.NewFuseOssfs2(nil, m)
	fusePodManagers := map[string]*oss.OSSFusePodManager{
		OssFsType:  oss.NewOSSFusePodManager(ossfs, nil),
		OssFs2Type: oss.NewOSSFusePodManager(ossfs2, nil),
	}
	ns := &nodeServer{
		metadata:        m,
		locks:           utils.NewVolumeLocks(),
		rawMounter:      mountutils.NewWithoutSystemd(""),
		skipAttach:      true,
		fusePodManagers: fusePodManagers,
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
