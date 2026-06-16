package oss

import (
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	k8sver "k8s.io/apimachinery/pkg/util/version"
)

func NewServers(endpoint string, meta *metadata.Metadata, serviceType utils.ServiceType, csiCfg utils.Config, k8sVersion *k8sver.Version) *common.Servers {
	panic("OSS driver is not supported on Windows")
}

type CSIAgent struct {
	csi.UnimplementedNodeServer
}

func NewCSIAgent(m metadata.MetadataProvider, socketPath string) *CSIAgent {
	panic("OSS CSI agent is not supported on Windows")
}
