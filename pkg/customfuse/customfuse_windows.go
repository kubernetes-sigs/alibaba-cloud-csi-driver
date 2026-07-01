package customfuse

import (
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	k8sver "k8s.io/apimachinery/pkg/util/version"
)

func NewServers(m metadata.MetadataProvider, endpoint string, serviceType utils.ServiceType, csiCfg utils.Config, k8sVersion *k8sver.Version, mountProxySock string) *common.Servers {
	panic("CustomFuse driver is not supported on Windows")
}

type CSIAgent struct {
	csi.UnimplementedNodeServer
}

func NewCSIAgent(socketPath string) *CSIAgent {
	panic("CustomFuse CSI agent is not supported on Windows")
}
