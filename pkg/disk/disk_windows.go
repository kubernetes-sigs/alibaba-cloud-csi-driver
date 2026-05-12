package disk

import (
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
)

type DISK struct{}

func NewDriver(m metadata.MetadataProvider, endpoint string, serviceType utils.ServiceType, csiCfg utils.Config) *DISK {
	panic("Disk driver is not supported on Windows yet")
}

func (d *DISK) Run() {
	panic("Disk driver is not supported on Windows yet")
}

type CSIAgent struct {
	csi.UnimplementedNodeServer
}

func NewCSIAgent() *CSIAgent {
	panic("Disk CSI agent is not supported on Windows yet")
}
