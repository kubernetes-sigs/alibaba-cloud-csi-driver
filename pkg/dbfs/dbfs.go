package dbfs

import (
	"context"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dbfs"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-csi/drivers/pkg/csi-common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
)

const (
	driverName = "dbfsplugin.csi.alibabacloud.com"
	// InstanceID is instance id
	InstanceID = "instance-id"
)

var (
	version    = "1.0.0"
	masterURL  string
	kubeconfig string
	// GlobalConfigVar Global Config
	GlobalConfigVar GlobalConfig
)

// GlobalConfig save global values for plugin
type GlobalConfig struct {
	Region       string
	DbfsClient   *dbfs.Client
	MetricEnable bool
	NodeName     string
	EcsInstanceID string
}

type DBFS struct {
	driver           *csicommon.CSIDriver
	endpoint         string
	idServer         *csicommon.DefaultIdentityServer
	nodeServer       *nodeServer
	controllerServer csi.ControllerServer

	cap   []*csi.VolumeCapability_AccessMode
	cscap []*csi.ControllerServiceCapability
}

//NewDriver create the identity/node/controller server and disk driver
func NewDriver(nodeID, endpoint string) *DBFS {
	log.Infof("Driver: %v version: %v", driverName, version)

	d := &DBFS{}
	d.endpoint = endpoint
	if nodeID == "" {
		nodeID, _ = utils.GetMetaData(InstanceID)
		log.Infof("Use node id : %s", nodeID)
	}
	csiDriver := csicommon.NewCSIDriver(driverName, version, nodeID)
	csiDriver.AddVolumeCapabilityAccessModes([]csi.VolumeCapability_AccessMode_Mode{csi.VolumeCapability_AccessMode_MULTI_NODE_MULTI_WRITER})
	csiDriver.AddControllerServiceCapabilities([]csi.ControllerServiceCapability_RPC_Type{
		csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME,
		csi.ControllerServiceCapability_RPC_PUBLISH_UNPUBLISH_VOLUME,
	})

	// Global Configs Set
	GlobalConfigSet()

	d.driver = csiDriver
	accessKeyID, accessSecret, accessToken := utils.GetDefaultAK()
	c := newDbfsClient(accessKeyID, accessSecret, accessToken, "")
	region := os.Getenv("REGION_ID")
	if region == "" {
		region = GetMetaData(RegionTag)
	}
	d.controllerServer = NewControllerServer(d.driver, c, region)
	GlobalConfigVar.DbfsClient = c

	return d
}

// Run start a new NodeServer
func (d *DBFS) Run() {
	s := csicommon.NewNonBlockingGRPCServer()
	s.Start(d.endpoint,
		csicommon.NewDefaultIdentityServer(d.driver),
		d.controllerServer,
		newNodeServer(d))
	s.Wait()
}

// GlobalConfigSet set global config
func GlobalConfigSet() {
	// Global Configs Set
	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		log.Fatalf("Error building kubeconfig: %s", err.Error())
	}
	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		log.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	configMapName := "csi-plugin"
	isMetricEnable := false

	configMap, err := kubeClient.CoreV1().ConfigMaps("kube-system").Get(context.Background(), configMapName, metav1.GetOptions{})
	if err != nil {
		log.Infof("Not found configmap named as csi-plugin under kube-system, with error: %v", err)
	} else {
		if value, ok := configMap.Data["dbfs-metric-enable"]; ok {
			if value == "enable" || value == "yes" || value == "true" {
				log.Infof("Nas Metric is enabled by configMap(%s).", value)
				isMetricEnable = true
			}
		}
	}

	metricDbfsConf := os.Getenv(DbfsMetricByPlugin)
	if metricDbfsConf == "true" || metricDbfsConf == "yes" {
		isMetricEnable = true
	} else if metricDbfsConf == "false" || metricDbfsConf == "no" {
		isMetricEnable = false
	}

	nodeName := os.Getenv("KUBE_NODE_NAME")
	GlobalConfigVar.MetricEnable = isMetricEnable
	GlobalConfigVar.NodeName = nodeName
	GlobalConfigVar.EcsInstanceID, _ = utils.GetMetaData(InstanceID)
}
