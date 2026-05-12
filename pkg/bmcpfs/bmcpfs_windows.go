package bmcpfs

import (
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
)

type Driver struct{}

func NewDriver(meta *metadata.Metadata, endpoint string, serviceType utils.ServiceType) *Driver {
	panic("BMCPFS driver is not supported on Windows")
}

func (d *Driver) Run() {
	panic("BMCPFS driver is not supported on Windows")
}
