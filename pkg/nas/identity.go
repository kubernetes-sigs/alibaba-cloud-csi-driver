package nas

import (
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
)

type identityServer struct {
	common.GenericIdentityServer
}

func newIdentityServer() *identityServer {
	return &identityServer{
		GenericIdentityServer: common.GenericIdentityServer{
			Name: driverName,
		},
	}
}
