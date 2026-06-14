package disk

import "github.com/container-storage-interface/spec/lib/go/csi"

type CSIAgent struct {
	csi.UnimplementedNodeServer
}

func NewCSIAgent() *CSIAgent {
	panic("Disk CSI agent is not supported on Windows yet")
}
