package oss

import (
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
)

type OSS struct{}

func NewDriver(endpoint string, meta *metadata.Metadata, serviceType utils.ServiceType, csiCfg utils.Config) *OSS {
	panic("OSS driver is not supported on Windows")
}

func (d *OSS) Run() {
	panic("OSS driver is not supported on Windows")
}

type CSIAgent struct {
	csi.UnimplementedNodeServer
}

func NewCSIAgent(m metadata.MetadataProvider, socketPath string) *CSIAgent {
	panic("OSS CSI agent is not supported on Windows")
}
