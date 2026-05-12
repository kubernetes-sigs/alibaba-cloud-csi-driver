package ens

import "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"

type ENS struct{}

func NewDriver(nodeID, endpoint string, serviceType utils.ServiceType) *ENS {
	panic("ENS driver is not supported on Windows")
}

func (d *ENS) Run() {
	panic("ENS driver is not supported on Windows")
}
