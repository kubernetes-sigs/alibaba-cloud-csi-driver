package dbfs

import (
	"context"
	"os"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/dbfs"
	"github.com/container-storage-interface/spec/lib/go/csi"
	csicommon "github.com/kubernetes-csi/drivers/pkg/csi-common"
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
	driverName = "dbfsplugin.csi.alibabacloud.com"
	// InstanceID is instance id
	InstanceID = "instance-id"
	// DBFSAttachByController tag
	DBFSAttachByController = "DBFS_AD_CONTROLLER"
	// DbfsDetachDisable tag
	DbfsDetachDisable = "DBFS_DETACH_DISABLE"
)

var (
	// GlobalConfigVar Global Config
	GlobalConfigVar GlobalConfig
)

// GlobalConfig save global values for plugin
type GlobalConfig struct {
	Region       string
	DbfsClient   *dbfs.Client
	MetricEnable bool
	NodeName     string
	DBFSDomain   string
	// only useful in node
	EcsInstanceID      string
	ADControllerEnable bool
	DBFSDetachDisable  bool
}

// DBFS define driver
type DBFS struct {
	driver           *csicommon.CSIDriver
	endpoint         string
	idServer         *csicommon.DefaultIdentityServer
	nodeServer       *nodeServer
	controllerServer csi.ControllerServer
}

// NewDriver create the identity/node/controller server and dbfs driver
func NewDriver(nodeID, endpoint string) *DBFS {
	log.Infof("Driver: %v version: %v", driverName, version.VERSION)

	d := &DBFS{}
	d.endpoint = endpoint
	if nodeID == "" {
		nodeID = utils.RetryGetMetaData(InstanceID)
		log.Infof("DBFS Use node id : %s", nodeID)
	}
	csiDriver := csicommon.NewCSIDriver(driverName, version.VERSION, nodeID)
	csiDriver.AddVolumeCapabilityAccessModes([]csi.VolumeCapability_AccessMode_Mode{csi.VolumeCapability_AccessMode_MULTI_NODE_MULTI_WRITER})
	csiDriver.AddControllerServiceCapabilities([]csi.ControllerServiceCapability_RPC_Type{
		csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME,
		csi.ControllerServiceCapability_RPC_PUBLISH_UNPUBLISH_VOLUME,
		csi.ControllerServiceCapability_RPC_CREATE_DELETE_SNAPSHOT,
		csi.ControllerServiceCapability_RPC_LIST_SNAPSHOTS,
		csi.ControllerServiceCapability_RPC_EXPAND_VOLUME,
	})

	d.driver = csiDriver
	accessControl := utils.GetAccessControl()
	c := newDbfsClient(accessControl, "")
	region := os.Getenv("REGION_ID")
	if region == "" {
		region, _ = utils.GetMetaData(RegionTag)
	}
	d.controllerServer = NewControllerServer(d.driver, c, region)
	GlobalConfigVar.DbfsClient = c

	// Global Configs Set
	GlobalConfigSet(region)
	return d
}

// Run start a new NodeServer
func (d *DBFS) Run() {
	common.RunCSIServer(d.endpoint, NewIdentityServer(d.driver), d.controllerServer, newNodeServer(d))
}

// GlobalConfigSet set global config
func GlobalConfigSet(region string) {
	// Global Configs Set
	cfg, err := clientcmd.BuildConfigFromFlags(options.MasterURL, options.Kubeconfig)
	if err != nil {
		log.Fatalf("Error building kubeconfig: %s", err.Error())
	}
	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		log.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	isADControllerEnable := true
	configMapName := "csi-plugin"
	isMetricEnable := true

	configMap, err := kubeClient.CoreV1().ConfigMaps("kube-system").Get(context.Background(), configMapName, metav1.GetOptions{})
	if err != nil {
		log.Infof("Not found configmap named as csi-plugin under kube-system, with error: %v", err)
	} else {
		if value, ok := configMap.Data["dbfs-metric-enable"]; ok {
			if value == "enable" || value == "yes" || value == "true" {
				log.Infof("Dbfs Metric is enabled by configMap(%s).", value)
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

	// Env variables
	adEnable := os.Getenv(DBFSAttachByController)
	if adEnable == "true" || adEnable == "yes" {
		log.Infof("AD-Controller is enabled by Env(%s), CSI DBFS Plugin running in AD Controller mode.", adEnable)
		isADControllerEnable = true
	} else if adEnable == "false" || adEnable == "no" {
		log.Infof("AD-Controller is disabled by Env(%s), CSI DBFS Plugin running in kubelet mode.", adEnable)
		isADControllerEnable = false
	}
	if isADControllerEnable {
		log.Infof("AD-Controller is enabled, CSI DBFS Plugin running in AD Controller mode.")
	} else {
		log.Infof("AD-Controller is disabled, CSI DBFS Plugin running in kubelet mode.")
	}

	isDbfsDetachDisable := false
	dbfsDetachDisable := os.Getenv(DbfsDetachDisable)
	if dbfsDetachDisable == "true" || dbfsDetachDisable == "yes" {
		isDbfsDetachDisable = true
	} else if dbfsDetachDisable == "false" || dbfsDetachDisable == "no" {
		isDbfsDetachDisable = false
	}

	nodeName := os.Getenv("KUBE_NODE_NAME")
	GlobalConfigVar.MetricEnable = isMetricEnable
	GlobalConfigVar.NodeName = nodeName
	GlobalConfigVar.Region = region
	GlobalConfigVar.EcsInstanceID, _ = utils.GetMetaData(InstanceID)
	GlobalConfigVar.ADControllerEnable = isADControllerEnable
	GlobalConfigVar.DBFSDomain = "dbfs-vpc." + GlobalConfigVar.Region + ".aliyuncs.com"
	GlobalConfigVar.DBFSDetachDisable = isDbfsDetachDisable
	log.Infof("DBFS Global Config: %v", GlobalConfigVar)
}
