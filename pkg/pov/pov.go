package pov

import (
	"os"
	"strconv"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/options"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
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

func initDriver() {}

func NewDriver(nodeID, endpoint string, serviceType utils.ServiceType) *PoV {
	initDriver()
	poV := &PoV{}
	poV.endpoint = endpoint
	newGlobalConfig()

	if serviceType&utils.Controller != 0 {
		poV.controllerService = newControllerService()
	}
	if serviceType&utils.Node != 0 {
		poV.nodeService = newNodeService()
	}

	return poV
}

// Run start pov driver service
func (p *PoV) Run() {
	log.Infof("Starting pov driver service, endpoint: %s", p.endpoint)
	common.RunCSIServer(p.endpoint, p, &p.controllerService, &p.nodeService)
}

func newGlobalConfig() {
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

	ms := NewMetadataService()
	doc, err := ms.GetDoc()
	if err != nil {
		log.Fatalf("Failed to get baseinfo from metadataserver: %+v", err)
	}

	GlobalConfigVar = GlobalConfig{
		client:     kubeClient,
		regionID:   doc.RegionID,
		instanceID: doc.InstanceID,
		zoneID:     doc.ZoneID,
	}
}

type GlobalConfig struct {
	regionID   string
	instanceID string
	zoneID     string
	client     kubernetes.Interface
}
