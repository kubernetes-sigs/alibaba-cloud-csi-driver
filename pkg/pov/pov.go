package pov

import (
	"fmt"
	"net"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/options"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
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
	srv      *grpc.Server

	controllerService
	nodeService
}

func initDriver() {}

func NewDriver(nodeID, endpoint string, runAsController bool) *PoV {
	initDriver()
	poV := &PoV{}
	poV.endpoint = endpoint
	newGlobalConfig(runAsController)

	if runAsController {
		poV.controllerService = newControllerService()
	} else {
		poV.nodeService = newNodeService()
	}

	return poV
}

// Run start pov driver service
func (p *PoV) Run() error {
	log.Infof("Starting pov driver service, endpoint: %s", p.endpoint)
	scheme, addr, err := parseEndpoint(p.endpoint)
	if err != nil {
		log.Fatalf("failed to parse endpoint: %v", err)
	}

	listener, err := net.Listen(scheme, addr)
	if err != nil {
		log.Fatalf("failed to listen endpoint: %s,  %v", p.endpoint, err)
	}
	p.srv = grpc.NewServer()

	log.Info("Starting pov register identity service")
	csi.RegisterIdentityServer(p.srv, p)

	if GlobalConfigVar.controllerService {
		log.Info("Starting pov register in controller mode, register controller service")
		csi.RegisterControllerServer(p.srv, p)
	} else {
		log.Info("Starting pov register in node mode, register node service")
		csi.RegisterNodeServer(p.srv, p)
	}
	return p.srv.Serve(listener)
}

func parseEndpoint(endpoint string) (string, string, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return "", "", fmt.Errorf("could not parse endpoint: %w", err)
	}

	addr := filepath.Join(u.Host, filepath.FromSlash(u.Path))

	scheme := strings.ToLower(u.Scheme)
	switch scheme {
	case "tcp":
	case "unix":
		addr = filepath.Join("/", addr)
		if err := os.Remove(addr); err != nil && !os.IsNotExist(err) {
			return "", "", fmt.Errorf("could not remove unix domain socket %q: %w", addr, err)
		}
	default:
		return "", "", fmt.Errorf("unsupported protocol: %s", scheme)
	}

	return scheme, addr, nil
}

func newGlobalConfig(runAsController bool) {
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
		controllerService: runAsController,
		client:            kubeClient,
		regionID:          doc.RegionID,
		instanceID:        doc.InstanceID,
		zoneID:            doc.ZoneID,
	}
}

type GlobalConfig struct {
	regionID          string
	instanceID        string
	zoneID            string
	controllerService bool
	client            kubernetes.Interface
}
