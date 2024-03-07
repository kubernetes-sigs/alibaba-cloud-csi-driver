package common

import (
	"github.com/container-storage-interface/spec/lib/go/csi"
	csicommon "github.com/kubernetes-csi/drivers/pkg/csi-common"
)

func RunCSIServer(endpoint string, ids csi.IdentityServer, cs csi.ControllerServer, ns csi.NodeServer) {
	ns = WrapNodeServerWithValidator(ns)
	cs = WrapControllerServerWithValidator(cs)
	// TODO: WrapGroupControllerServerWithValidator

	// start grpc server
	s := csicommon.NewNonBlockingGRPCServer()
	s.Start(endpoint, ids, cs, ns)
	s.Wait()
}
