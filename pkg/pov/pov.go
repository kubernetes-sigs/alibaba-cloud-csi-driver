package pov

import (
	"os"
	"strconv"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/options"
	log "github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

const (
	DriverName = "povplugin.csi.alibabacloud.com"
)

var GlobalConfigVar GlobalConfig

// Pangu Over Virtio
type PoV struct {
	endpoint string

	controllerService
	nodeService
}

func NewDriver(meta *metadata.Metadata, endpoint string, runAsController bool) *PoV {
	poV := &PoV{}
	poV.endpoint = endpoint
	newGlobalConfig(meta, runAsController)

	if runAsController {
		poV.controllerService = newControllerService()
	} else {
		poV.nodeService = newNodeService(meta)
	}

	return poV
}

// Run start pov driver service
func (p *PoV) Run() {
	log.Infof("Starting pov driver service, endpoint: %s", p.endpoint)
	common.RunCSIServer(p.endpoint, newIdentityServer(), &p.controllerService, &p.nodeService)
}

func newGlobalConfig(meta *metadata.Metadata, runAsController bool) {
	cfg, err := clientcmd.BuildConfigFromFlags(options.MasterURL, options.Kubeconfig)
	if err != nil {
		log.Fatalf("newGlobalConfig: build kubeconfig failed: %v", err)
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
		log.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	GlobalConfigVar = GlobalConfig{
		controllerService: runAsController,
		client:            kubeClient,
		regionID:          metadata.MustGet(meta, metadata.RegionID),
	}
}

type GlobalConfig struct {
	regionID          string
	controllerService bool
	client            kubernetes.Interface
}
