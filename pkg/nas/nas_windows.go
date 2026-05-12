package nas

import (
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
)

const driverName = "nasplugin.csi.alibabacloud.com"

type NAS struct{}

func NewDriver(meta *metadata.Metadata, endpoint string, serviceType utils.ServiceType, csiCfg utils.Config) *NAS {
	panic("NAS driver is not supported on Windows")
}

func (d *NAS) Run() {
	panic("NAS driver is not supported on Windows")
}

type CSIAgent struct {
	csi.UnimplementedNodeServer
}

func NewCSIAgent(socketPath string) *CSIAgent {
	panic("NAS CSI agent is not supported on Windows")
}
