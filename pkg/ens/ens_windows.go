package ens

import (
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
)

func NewServers(nodeID, endpoint string, serviceType utils.ServiceType) *common.Servers {
	panic("ENS driver is not supported on Windows")
}
