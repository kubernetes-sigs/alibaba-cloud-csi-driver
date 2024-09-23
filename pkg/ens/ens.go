package ens

import (
	"context"
	"os"
	"strconv"
	"sync"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/options"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/version"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
)

const (
	driverType            = "ens"
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
	endpoint string
	servers  common.Servers
}

func initDriver() {}

func NewDriver(nodeID, endpoint string, serviceType utils.ServiceType) *ENS {

	initDriver()
	tmpENS := &ENS{}
	tmpENS.endpoint = endpoint

	NewGlobalConfig()

	var servers common.Servers
	servers.IdentityServer = NewIdentityServer()
	if serviceType&utils.Controller != 0 {
		servers.ControllerServer = NewControllerServer()
	}
	if serviceType&utils.Node != 0 {
		servers.NodeServer = NewNodeServer()
	}
	tmpENS.servers = servers
	return tmpENS
}

func (ens *ENS) Run() {
	klog.Infof("Run: start csi-plugin driver: %s, version %s", driverName, version.VERSION)
	common.RunCSIServer(driverType, ens.endpoint, ens.servers)
}

func NewGlobalConfig() {
	ctx := context.Background()
	cfg, err := clientcmd.BuildConfigFromFlags(options.MasterURL, options.Kubeconfig)
	if err != nil {
		klog.Fatalf("NewGlobalConfig: Error building kubeconfig: %s", err.Error())
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

	nodeName := os.Getenv("KUBE_NODE_NAME")
	if nodeName == "" {
		klog.Fatalf("NewGlobalConfig: failed to get NODE_NAME from env, NODE_NAME is necessary")
	}

	node, err := kubeClient.CoreV1().Nodes().Get(ctx, nodeName, metav1.GetOptions{})
	if err != nil {
		klog.Fatalf("NewGlobalConfig: failed to get node info from cluster: %+v", err)
	}
	instanceID, ok := node.Labels[ensInstanceIDLabelKey]
	if !ok {
		klog.Fatalf("NewGlobalConfig: ens instance key not in node labels %+v", node.Labels)
	}
	profile, err := kubeClient.CoreV1().ConfigMaps(kubeSystemNamespace).Get(ctx, clusterProfileKey, metav1.GetOptions{})
	if err != nil {
		klog.Fatalf("NewGlobalConfig: failed to get ack profile configmap %+v", err)
	}
	clusterID, ok := profile.Data[clusterIdKey]
	if !ok {
		klog.Fatalf("NewGlobalConfig: clusterid key not in configmap %+v", profile.Data)
	}
	ensClient := newENSClient()

	instances, err := ensClient.DescribeInstance(instanceID)
	if err != nil {
		klog.Fatalf("NewGlobalConfig: describe instance failed: %+v", err)
	}
	var regionID string
	if len(instances) == 1 {
		regionID = *instances[0].EnsRegionId
	} else {
		klog.Fatalf("NewGlobalConfig: get wrong ens instanceid count from describeInstance, instanceID: %s, instances: %+v", instanceID, instances)
	}
	attachDetachController := os.Getenv("DISK_AD_CONTROLLER")
	if attachDetachController == "" {
		attachDetachController = "false"
	}
	detachBeforeAttach, err := strconv.ParseBool(os.Getenv("DETACH_BEFORE_ATTACH"))
	if err != nil {
		detachBeforeAttach = true
	}

	GlobalConfigVar = GlobalConfig{
		KClient:                      kubeClient,
		InstanceID:                   instanceID,
		ENSCli:                       ensClient,
		ClusterID:                    clusterID,
		RegionID:                     regionID,
		EnableAttachDetachController: attachDetachController,
		DetachBeforeAttach:           detachBeforeAttach,
	}
}

type GlobalConfig struct {
	RegionID           string
	InstanceID         string
	ClusterID          string
	DetachBeforeAttach bool

	EnableDiskPartition          string
	EnableAttachDetachController string
	KClient                      kubernetes.Interface
	ENSCli                       ENSClient
	AttachMutex                  sync.Mutex
}
