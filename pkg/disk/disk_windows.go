package disk

import (
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
)

func NewServers(m metadata.MetadataProvider, ecsV2 cloud.ECSv2Interface, endpoint string, serviceType utils.ServiceType, csiCfg utils.Config, useLabeler bool) *common.Servers {
	panic("Disk driver is not supported on Windows yet")
}

type CSIAgent struct {
	csi.UnimplementedNodeServer
}

func NewCSIAgent() *CSIAgent {
	panic("Disk CSI agent is not supported on Windows yet")
}
