package common

import (
	csicommon "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/agent/csi-common"
)

func RunCSIServer(endpoint string, servers csicommon.Servers) {
	servers.Ns = WrapNodeServerWithValidator(servers.Ns)
	servers.Cs = WrapControllerServerWithValidator(servers.Cs)
	if servers.Gcs != nil {
		servers.Gcs = WrapGroupControllerServerWithValidator(servers.Gcs)
	}

	// start grpc server
	s := csicommon.NewNonBlockingGRPCServer()
	s.Start(endpoint, servers)
	s.Wait()
}
