package bmcpfs

import (
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
)

func NewServers(meta *metadata.Metadata, endpoint string, serviceType utils.ServiceType, mountProxySock string) *common.Servers {
	panic("BMCPFS driver is not supported on Windows")
}
