package pov

import (
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
)

const (
	DriverName = "povplugin.csi.alibabacloud.com"
)

type PoV struct{}

func NewDriver(meta *metadata.Metadata, endpoint string, serviceType utils.ServiceType) *PoV {
	panic("POV driver is not supported on Windows")
}

func (p *PoV) Run() {
	panic("POV driver is not supported on Windows")
}
