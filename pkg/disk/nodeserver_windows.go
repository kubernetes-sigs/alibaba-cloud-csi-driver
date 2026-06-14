package disk

import (
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
)

func NewNodeServer(csiCfg utils.Config, ecs cloud.ECSInterface, m metadata.MetadataProvider) csi.NodeServer {
	panic("disk driver is not supported on Windows")
}
