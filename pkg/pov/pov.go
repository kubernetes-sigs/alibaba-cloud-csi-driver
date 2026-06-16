//go:build !windows

package pov

import (
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/options"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"
)

const (
	DriverName = "povplugin.csi.alibabacloud.com"
)

var GlobalConfigVar GlobalConfig

// Pangu Over Virtio
func NewServers(meta *metadata.Metadata, endpoint string, serviceType utils.ServiceType) *common.Servers {
	newGlobalConfig(meta, serviceType)

	var servers common.Servers
	servers.IdentityServer = newIdentityServer()
	if serviceType&utils.Controller != 0 {
		cs := newControllerService()
		servers.ControllerServer = &cs
	}
	if serviceType&utils.Node != 0 {
		ns := newNodeService(meta)
		servers.NodeServer = &ns
	}

	return &servers
}

func newGlobalConfig(meta *metadata.Metadata, serviceType utils.ServiceType) {
	cfg, err := options.GetRestConfig()
	if err != nil {
		klog.Fatalf("newGlobalConfig: build kubeconfig failed: %v", err)
	}
	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	GlobalConfigVar = GlobalConfig{
		client:   kubeClient,
		regionID: metadata.MustGet(meta, metadata.RegionID),
	}
}

type GlobalConfig struct {
	regionID          string
	controllerService bool
	client            kubernetes.Interface
}
