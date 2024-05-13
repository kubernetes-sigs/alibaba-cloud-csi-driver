package common

import (
	csicommon "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/agent/csi-common"
)

func RunCSIServer(endpoint string, servers csicommon.Servers) {
	servers.NodeServer = WrapNodeServerWithValidator(servers.NodeServer)
	servers.ControllerServer = WrapControllerServerWithValidator(servers.ControllerServer)
	if servers.GroupControllerServer != nil {
		servers.GroupControllerServer = WrapGroupControllerServerWithValidator(servers.GroupControllerServer)
	}

	// start grpc server
	s := csicommon.NewNonBlockingGRPCServer()
	s.Start(endpoint, servers)
	s.Wait()
}
