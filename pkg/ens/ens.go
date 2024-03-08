package ens

import (
	"context"
	"os"
	"strconv"
	"sync"

	"github.com/container-storage-interface/spec/lib/go/csi"
	csicommon "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/agent/csi-common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/options"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/version"
	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

const (
	driverName            = "ensplugin.csi.alibabacloud.com"
	ensInstanceIDLabelKey = "alibabacloud.com/ens-instance-id"
	clusterProfileKey     = "ack-cluster-profile"
	clusterIdKey          = "clusterid"
	kubeSystemNamespace   = "kube-system"

	// ens disk type
	CLOUD_EFFICIENCY   = "cloud_efficiency"
	CLOUD_SSD          = "cloud_ssd"
	LOCAL_HDD          = "local_hdd"
	LOCAL_SSD          = "local_ssd"
	ENS_DISK_AVAILABLE = "available"

	// NodeSchedueTag in annotations
	NodeSchedueTag = "volume.kubernetes.io/selected-node"

	// DiskNotAvailable error
	DiskNotAvailable = "InvalidDataDiskCategory.NotSupported"
	// DiskNotAvailableVer2 error
	DiskNotAvailableVer2 = "'DataDisk.n.Category' is not valid in this region."

	TopologyRegionKey         = "topology." + driverName + "/region"
	TopologyMultiRegionPrefix = TopologyRegionKey + "-"
)

var (
	GlobalConfigVar GlobalConfig
	ENSDiskTypeMap  = map[string]string{CLOUD_EFFICIENCY: "4", CLOUD_SSD: "3", LOCAL_HDD: "2", LOCAL_SSD: "1"}
)

type ENS struct {
	driver           *csicommon.CSIDriver
	endpoint         string
	idServer         csi.IdentityServer
	nodeServer       csi.NodeServer
	controllerServer csi.ControllerServer
}

func initDriver() {}

func NewDriver(nodeID, endpoint string) *ENS {

	initDriver()
	tmpENS := &ENS{}
	tmpENS.endpoint = endpoint

	NewGlobalConfig()

	csiDriver := csicommon.NewCSIDriver(driverName, version.VERSION, GlobalConfigVar.InstanceID)
	tmpENS.driver = csiDriver

	tmpENS.driver.AddControllerServiceCapabilities([]csi.ControllerServiceCapability_RPC_Type{
		csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME,
		csi.ControllerServiceCapability_RPC_PUBLISH_UNPUBLISH_VOLUME,
		// volume expansion is not support
		// csi.ControllerServiceCapability_RPC_EXPAND_VOLUME,
	})
	tmpENS.driver.AddVolumeCapabilityAccessModes([]csi.VolumeCapability_AccessMode_Mode{csi.VolumeCapability_AccessMode_MULTI_NODE_MULTI_WRITER})

	tmpENS.idServer = NewIdentityServer(tmpENS.driver)
	if GlobalConfigVar.ControllerService {
		tmpENS.controllerServer = NewControllerServer(tmpENS.driver)
	} else {
		tmpENS.nodeServer = NewNodeServer(tmpENS.driver)
	}
	return tmpENS
}

func (ens *ENS) Run() {
	log.Infof("Run: start csi-plugin driver: %s, version %s", driverName, version.VERSION)
	servers := csicommon.Servers{
		Ids: ens.idServer,
		Cs:  ens.controllerServer,
		Ns:  ens.nodeServer,
	}
	common.RunCSIServer(ens.endpoint, servers)
}

func NewGlobalConfig() {
	ctx := context.Background()
	cfg, err := clientcmd.BuildConfigFromFlags(options.MasterURL, options.Kubeconfig)
	if err != nil {
		log.Fatalf("NewGlobalConfig: Error building kubeconfig: %s", err.Error())
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

	nodeName := os.Getenv("KUBE_NODE_NAME")
	if nodeName == "" {
		log.Fatalf("NewGlobalConfig: failed to get NODE_NAME from env, NODE_NAME is necessary")
	}

	node, err := kubeClient.CoreV1().Nodes().Get(ctx, nodeName, metav1.GetOptions{})
	if err != nil {
		log.Fatalf("NewGlobalConfig: failed to get node info from cluster: %+v", err)
	}
	instanceID, ok := node.Labels[ensInstanceIDLabelKey]
	if !ok {
		log.Fatalf("NewGlobalConfig: ens instance key not in node labels %+v", node.Labels)
	}
	profile, err := kubeClient.CoreV1().ConfigMaps(kubeSystemNamespace).Get(ctx, clusterProfileKey, metav1.GetOptions{})
	if err != nil {
		log.Fatalf("NewGlobalConfig: failed to get ack profile configmap %+v", err)
	}
	clusterID, ok := profile.Data[clusterIdKey]
	if !ok {
		log.Fatalf("NewGlobalConfig: clusterid key not in configmap %+v", profile.Data)
	}
	ensClient := newENSClient()

	instances, err := ensClient.DescribeInstance(instanceID)
	if err != nil {
		log.Fatalf("NewGlobalConfig: describe instance failed: %+v", err)
	}
	var regionID string
	if len(instances) == 1 {
		regionID = *instances[0].EnsRegionId
	} else {
		log.Fatalf("NewGlobalConfig: get wrong ens instanceid count from describeInstance, instanceID: %s, instances: %+v", instanceID, instances)
	}
	attachDetachController := os.Getenv("DISK_AD_CONTROLLER")
	if attachDetachController == "" {
		attachDetachController = "false"
	}
	detachBeforeAttach, err := strconv.ParseBool(os.Getenv("DETACH_BEFORE_ATTACH"))
	if err != nil {
		detachBeforeAttach = true
	}

	controllerServerType := false
	if os.Getenv(utils.ServiceType) == utils.ProvisionerService {
		controllerServerType = true
	}

	GlobalConfigVar = GlobalConfig{
		KClient:                      kubeClient,
		InstanceID:                   instanceID,
		ENSCli:                       ensClient,
		ClusterID:                    clusterID,
		RegionID:                     regionID,
		EnableAttachDetachController: attachDetachController,
		DetachBeforeAttach:           detachBeforeAttach,
		ControllerService:            controllerServerType,
	}
}

type GlobalConfig struct {
	RegionID           string
	InstanceID         string
	ClusterID          string
	DetachBeforeAttach bool
	ControllerService  bool

	EnableDiskPartition          string
	EnableAttachDetachController string
	KClient                      kubernetes.Interface
	ENSCli                       ENSClient
	AttachMutex                  sync.Mutex
}
