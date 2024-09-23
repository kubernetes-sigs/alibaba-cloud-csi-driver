package pov

import (
	"os"
	"strconv"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/options"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
)

const (
	driverType = "pov"
	DriverName = "povplugin.csi.alibabacloud.com"
)

var GlobalConfigVar GlobalConfig

// Pangu Over Virtio
type PoV struct {
	endpoint string
	servers  common.Servers
}

func NewDriver(meta *metadata.Metadata, endpoint string, serviceType utils.ServiceType) *PoV {
	poV := &PoV{}
	poV.endpoint = endpoint
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
	poV.servers = servers

	return poV
}

// Run start pov driver service
func (p *PoV) Run() {
	klog.Infof("Starting pov driver service, endpoint: %s", p.endpoint)
	common.RunCSIServer(driverType, p.endpoint, p.servers)
}

func newGlobalConfig(meta *metadata.Metadata, serviceType utils.ServiceType) {
	cfg, err := clientcmd.BuildConfigFromFlags(options.MasterURL, options.Kubeconfig)
	if err != nil {
		klog.Fatalf("newGlobalConfig: build kubeconfig failed: %v", err)
	}

	if qps := os.Getenv("KUBE_CLI_API_QPS"); qps != "" {
		if qpsi, err := strconv.Atoi(qps); err == nil {
			cfg.QPS = float32(qpsi)
		}
	}
	if burst := os.Getenv("KUBE_CLI_API_BURST"); burst != "" {
		if qpsi, err := strconv.Atoi(burst); err == nil {
			cfg.Burst = qpsi
		}
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
