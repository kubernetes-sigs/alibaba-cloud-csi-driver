package pov

import (
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
)

const (
	DriverName = "povplugin.csi.alibabacloud.com"
)

func NewServers(meta *metadata.Metadata, endpoint string, serviceType utils.ServiceType) *common.Servers {
	panic("POV driver is not supported on Windows")
}
