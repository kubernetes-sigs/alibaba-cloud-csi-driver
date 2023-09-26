/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package disk

import (
	"context"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/container-storage-interface/spec/lib/go/csi"
	csicommon "github.com/kubernetes-csi/drivers/pkg/csi-common"
	snapClientset "github.com/kubernetes-csi/external-snapshotter/client/v4/clientset/versioned"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/log"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/options"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	crd "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// PluginFolder defines the location of diskplugin
const (
	driverName              = "diskplugin.csi.alibabacloud.com"
	csiVersion              = "1.0.0"
	TopologyZoneKey         = "topology." + driverName + "/zone"
	TopologyMultiZonePrefix = TopologyZoneKey + "-"
)

// DISK the DISK object
type DISK struct {
	driver           *csicommon.CSIDriver
	endpoint         string
	idServer         csi.IdentityServer
	nodeServer       csi.NodeServer
	controllerServer csi.ControllerServer

	cap   []*csi.VolumeCapability_AccessMode
	cscap []*csi.ControllerServiceCapability
}

// GlobalConfig save global values for plugin
type GlobalConfig struct {
	EcsClient             *ecs.Client
	Region                string
	NodeID                string
	ZoneID                string
	CanAttach             bool
	DiskTagEnable         bool
	AttachMutex           sync.RWMutex
	ADControllerEnable    bool
	DetachDisabled        bool
	MetricEnable          bool
	RunTimeClass          string
	DetachBeforeAttach    bool
	DetachBeforeDelete    bool
	DiskBdfEnable         bool
	ClientSet             *kubernetes.Clientset
	FilesystemLosePercent float64
	ClusterID             string
	DiskPartitionEnable   bool
	ControllerService     bool
	BdfHealthCheck        bool
	DiskMultiTenantEnable bool
	CheckBDFHotPlugin     bool
	SnapClient            *snapClientset.Clientset
	NodeMultiZoneEnable   bool
	WaitBeforeAttach      bool
	AddonVMFatalEvents    []string
	RequestBaseInfo       map[string]string
	SnapshotBeforeDelete  bool
	OmitFilesystemCheck   bool
}

// define global variable
var (
	GlobalConfigVar GlobalConfig
)

// Init checks for the persistent volume file and loads all found volumes
// into a memory structure
func initDriver() {
}

// NewDriver create the identity/node/controller server and disk driver
func NewDriver(nodeID, endpoint string, runAsController bool) *DISK {
	initDriver()
	tmpdisk := &DISK{}
	tmpdisk.endpoint = endpoint

	// Config Global vars
	cfg := GlobalConfigSet(nodeID)

	csiDriver := csicommon.NewCSIDriver(driverName, csiVersion, GlobalConfigVar.NodeID)
	tmpdisk.driver = csiDriver
	tmpdisk.driver.AddControllerServiceCapabilities([]csi.ControllerServiceCapability_RPC_Type{
		csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME,
		csi.ControllerServiceCapability_RPC_PUBLISH_UNPUBLISH_VOLUME,
		csi.ControllerServiceCapability_RPC_CREATE_DELETE_SNAPSHOT,
		csi.ControllerServiceCapability_RPC_LIST_SNAPSHOTS,
		csi.ControllerServiceCapability_RPC_EXPAND_VOLUME,
	})
	tmpdisk.driver.AddVolumeCapabilityAccessModes([]csi.VolumeCapability_AccessMode_Mode{csi.VolumeCapability_AccessMode_MULTI_NODE_MULTI_WRITER})

	// Init ECS Client
	accessControl := utils.GetAccessControl()
	client := newEcsClient(accessControl)
	if accessControl.UseMode == utils.EcsRAMRole || accessControl.UseMode == utils.ManagedToken {
		log.Log.Infof("Starting csi-plugin with sts.")
	} else {
		log.Log.Infof("Starting csi-plugin without sts.")
	}
	GlobalConfigVar.EcsClient = client

	apiExtentionClient, err := crd.NewForConfig(cfg)
	if err != nil {
		log.Log.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	// Create GRPC servers
	tmpdisk.idServer = NewIdentityServer(tmpdisk.driver)
	tmpdisk.controllerServer = NewControllerServer(tmpdisk.driver, apiExtentionClient)

	if !runAsController {
		tmpdisk.nodeServer = NewNodeServer(tmpdisk.driver, client)
	}

	return tmpdisk
}

// Run start a new NodeServer
func (disk *DISK) Run() {
	log.Log.Infof("Starting csi-plugin Driver: %v version: %v", driverName, csiVersion)
	common.RunCSIServer(disk.endpoint, disk.idServer, disk.controllerServer, disk.nodeServer)
}

// GlobalConfigSet set Global Config
func GlobalConfigSet(nodeID string) *restclient.Config {
	configMapName := "csi-plugin"
	isADControllerEnable := false
	isDiskTagEnable := false
	isDiskMetricEnable := true
	isDiskDetachDisable := false
	isDiskDetachBeforeDelete := true
	isDiskBdfEnable := false
	isDiskMultiTenantEnable := false
	isNodeMultiZoneEnable := false

	isWaitBeforeAttach := false
	if waitBeforeAttach := os.Getenv("WAIT_BEFORE_ATTACH"); waitBeforeAttach == "true" {
		isWaitBeforeAttach = true
	}

	// Global Configs Set
	cfg, err := clientcmd.BuildConfigFromFlags(options.MasterURL, options.Kubeconfig)
	if err != nil {
		log.Log.Fatalf("Error building kubeconfig: %s", err.Error())
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
	cfg.AcceptContentTypes = strings.Join([]string{runtime.ContentTypeProtobuf, runtime.ContentTypeJSON}, ",")
	// snapshotClient does not support protobuf
	snapClient, err := snapClientset.NewForConfig(cfg)
	if err != nil {
		log.Log.Fatalf("Error building kubernetes snapclientset: %s", err.Error())
	}

	cfg.ContentType = runtime.ContentTypeProtobuf
	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		log.Log.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	configMap, err := kubeClient.CoreV1().ConfigMaps("kube-system").Get(context.Background(), configMapName, metav1.GetOptions{})
	if err != nil {
		log.Log.Infof("Not found configmap named as csi-plugin under kube-system, with: %v", err)
	} else {
		if value, ok := configMap.Data["disk-adcontroller-enable"]; ok {
			if checkOption(value) {
				log.Log.Infof("AD-Controller is enabled by configMap(%s), CSI Disk Plugin running in AD Controller mode.", value)
				isADControllerEnable = true
			} else if checkOptionFalse(value) {
				log.Log.Infof("AD-Controller is disable by configMap(%s), CSI Disk Plugin running in kubelet mode.", value)
				isADControllerEnable = false
			}
		}
		if value, ok := configMap.Data["disk-tag-enable"]; ok {
			if checkOption(value) {
				log.Log.Infof("Disk Tag is enabled by configMap(%s).", value)
				isDiskTagEnable = true
			}
		}
		if value, ok := configMap.Data["disk-metric-enable"]; ok {
			if checkOption(value) {
				log.Log.Infof("Disk Metric is enabled by configMap(%s).", value)
				isDiskMetricEnable = true
			}
		}
		if value, ok := configMap.Data["disk-detach-disable"]; ok {
			if checkOption(value) {
				log.Log.Infof("Disk Detach is disabled by configMap(%s), this tag only works when adcontroller enabled.", value)
				isDiskDetachDisable = true
			} else if checkOptionFalse(value) {
				log.Log.Infof("Disk Detach is enable by configMap(%s), this tag only works when adcontroller enabled.", value)
				isDiskDetachDisable = false
			}
		}
		if value, ok := configMap.Data["disk-detach-before-delete"]; ok {
			if checkOption(value) {
				log.Log.Infof("Disk Detach before delete is enabled by configMap(%s).", value)
				isDiskDetachBeforeDelete = true
			} else if checkOptionFalse(value) {
				log.Log.Infof("Disk Detach before delete is diskable by configMap(%s).", value)
				isDiskDetachBeforeDelete = false
			}
		}
		if value, ok := configMap.Data["disk-bdf-enable"]; ok {
			if checkOption(value) {
				log.Log.Infof("Disk Bdf is enabled by configMap(%s).", value)
				isDiskBdfEnable = true
			} else if checkOptionFalse(value) {
				log.Log.Infof("Disk Bdf is disable by configMap(%s).", value)
				isDiskBdfEnable = false
			}
		}
	}

	// Env variables
	avmfe := os.Getenv("ADDON_VM_FATAL_EVENTS")
	fatalEvents := []string{}
	if avmfe != "" {
		if strings.Contains(avmfe, ",") {
			fatalEvents = strings.Split(avmfe, ",")
		} else {
			fatalEvents = []string{avmfe}
		}
	}

	adEnable := os.Getenv(DiskAttachByController)
	if adEnable == "true" || adEnable == "yes" {
		log.Log.Infof("AD-Controller is enabled by Env(%s), CSI Disk Plugin running in AD Controller mode.", adEnable)
		isADControllerEnable = true
	} else if adEnable == "false" || adEnable == "no" {
		log.Log.Infof("AD-Controller is disabled by Env(%s), CSI Disk Plugin running in kubelet mode.", adEnable)
		isADControllerEnable = false
	}
	if isADControllerEnable {
		log.Log.Infof("AD-Controller is enabled, CSI Disk Plugin running in AD Controller mode.")
	} else {
		log.Log.Infof("AD-Controller is disabled, CSI Disk Plugin running in kubelet mode.")
	}

	isDetachBeforeAttached := true
	forceDetached := os.Getenv(DiskForceDetached)
	if forceDetached == "false" || forceDetached == "no" {
		isDetachBeforeAttached = false
	}

	tagDiskConf := os.Getenv(DiskTagedByPlugin)
	if tagDiskConf == "true" || tagDiskConf == "yes" {
		isDiskTagEnable = true
	} else if tagDiskConf == "false" || tagDiskConf == "no" {
		isDiskTagEnable = false
	}

	metricDiskConf := os.Getenv(DiskMetricByPlugin)
	if metricDiskConf == "true" || metricDiskConf == "yes" {
		isDiskMetricEnable = true
	} else if metricDiskConf == "false" || metricDiskConf == "no" {
		isDiskMetricEnable = false
	}

	diskDetachDisable := os.Getenv(DiskDetachDisable)
	if diskDetachDisable == "true" || diskDetachDisable == "yes" {
		isDiskDetachDisable = true
	} else if diskDetachDisable == "false" || diskDetachDisable == "no" {
		isDiskDetachDisable = false
	}

	bdfDiskConf := os.Getenv(DiskBdfEnable)
	if bdfDiskConf == "true" || bdfDiskConf == "yes" {
		isDiskBdfEnable = true
	} else if bdfDiskConf == "false" || bdfDiskConf == "no" {
		isDiskBdfEnable = false
	}

	diskDetachBeforeDelete := os.Getenv(DiskDetachBeforeDelete)
	if diskDetachBeforeDelete == "true" || diskDetachBeforeDelete == "yes" {
		isDiskDetachBeforeDelete = true
	} else if diskDetachBeforeDelete == "false" || diskDetachBeforeDelete == "no" {
		isDiskDetachBeforeDelete = false
	}

	// fileSystemLosePercent ...
	fileSystemLosePercent := float64(0.90)
	if fileSystemLoseCapacityPercent := os.Getenv(FileSystemLoseCapacityPercent); fileSystemLoseCapacityPercent != "" {
		percent, err := strconv.ParseFloat(fileSystemLoseCapacityPercent, 64)
		if err == nil {
			fileSystemLosePercent = percent
		}
	}

	nodeName := os.Getenv(kubeNodeName)
	runtimeValue := "runc"
	var regionID, zoneID, vmID string
	nodeInfo, err := kubeClient.CoreV1().Nodes().Get(context.Background(), nodeName, metav1.GetOptions{})
	if err != nil {
		log.Log.Errorf("GlobalConfigSet: get node %s with error: %s", nodeName, err.Error())
		regionID = GetRegionID()
	} else {
		if value, ok := nodeInfo.Labels["alibabacloud.com/container-runtime"]; ok && strings.TrimSpace(value) == "Sandboxed-Container.runv" {
			if value, ok := nodeInfo.Labels["alibabacloud.com/container-runtime-version"]; ok && strings.HasPrefix(strings.TrimSpace(value), "1.") {
				runtimeValue = MixRunTimeMode
			}
		}
		log.Log.Infof("Describe node %s and Set RunTimeClass to %s", nodeName, runtimeValue)

		regionID, zoneID, vmID = getMeta(nodeInfo)
		log.Log.Infof("NewNodeServer: get instance meta info from metadataserver, regionID: %s, zoneID: %s, vmID: %s", regionID, zoneID, vmID)

		if nodeID == "" {
			nodeID = vmID
		}
	}
	if zoneID == "" || !strings.HasPrefix(vmID, "i-") {
		doc, err := retryGetInstanceDoc()
		log.Log.Infof("NewNodeServer: get instance meta info failed from metadataserver, err: %v, doc: %v", err, doc)
		if err == nil {
			zoneID = doc.ZoneID
			nodeID = doc.InstanceID
		}
	}
	runtimeEnv := os.Getenv("RUNTIME")
	if runtimeEnv == MixRunTimeMode {
		runtimeValue = MixRunTimeMode
	}
	clustID := os.Getenv("CLUSTER_ID")

	partition := true
	if partitionEn := os.Getenv("DISK_PARTITION_ENABLE"); partitionEn == "false" {
		partition = false
	}

	controllerServerType := false
	if os.Getenv(utils.ServiceType) == utils.ProvisionerService {
		controllerServerType = true
	}

	bdfCheck := true
	if os.Getenv("BDF_HEALTH_CHECK") == "false" {
		bdfCheck = false
	}
	diskMultiTenantEnable := os.Getenv(DiskMultiTenantEnable)
	if diskMultiTenantEnable == "true" || diskMultiTenantEnable == "yes" {
		log.Log.Infof("Multi tenant is Enabled")
		isDiskMultiTenantEnable = true
	} else if diskMultiTenantEnable == "false" || diskMultiTenantEnable == "no" {
		isDiskMultiTenantEnable = false
	}

	nodeMultiZoneEnable := os.Getenv(NodeMultiZoneEnable)
	if nodeMultiZoneEnable == "true" || nodeMultiZoneEnable == "yes" {
		log.Log.Infof("Multi zone node is Enabled")
		isNodeMultiZoneEnable = true
	}

	delAutoSnap := false
	volumeDelAutoSnap := os.Getenv("VOLUME_DEL_AUTO_SNAP")
	if volumeDelAutoSnap == "true" {
		delAutoSnap = true
	}

	omitFsCheck := false
	disableFsCheck := os.Getenv("DISABLE_FS_CHECK")
	if disableFsCheck == "true" {
		omitFsCheck = true
	}

	log.Log.Infof("Starting with GlobalConfigVar: region(%s), zone(%s), NodeID(%s), ADControllerEnable(%t), DiskTagEnable(%t), DiskBdfEnable(%t), MetricEnable(%t), RunTimeClass(%s), DetachDisabled(%t), DetachBeforeDelete(%t), ClusterID(%s)", regionID, zoneID, nodeID, isADControllerEnable, isDiskTagEnable, isDiskBdfEnable, isDiskMetricEnable, runtimeValue, isDiskDetachDisable, isDiskDetachBeforeDelete, clustID)
	// Global Config Set
	GlobalConfigVar = GlobalConfig{
		Region:                regionID,
		NodeID:                nodeID,
		ZoneID:                zoneID,
		CanAttach:             true,
		ADControllerEnable:    isADControllerEnable,
		DiskTagEnable:         isDiskTagEnable,
		DiskBdfEnable:         isDiskBdfEnable,
		MetricEnable:          isDiskMetricEnable,
		RunTimeClass:          runtimeValue,
		DetachDisabled:        isDiskDetachDisable,
		DetachBeforeDelete:    isDiskDetachBeforeDelete,
		DetachBeforeAttach:    isDetachBeforeAttached,
		ClientSet:             kubeClient,
		SnapClient:            snapClient,
		FilesystemLosePercent: fileSystemLosePercent,
		ClusterID:             clustID,
		DiskPartitionEnable:   partition,
		ControllerService:     controllerServerType,
		BdfHealthCheck:        bdfCheck,
		DiskMultiTenantEnable: isDiskMultiTenantEnable,
		NodeMultiZoneEnable:   isNodeMultiZoneEnable,
		WaitBeforeAttach:      isWaitBeforeAttach,
		AddonVMFatalEvents:    fatalEvents,
		SnapshotBeforeDelete:  delAutoSnap,
		RequestBaseInfo:       map[string]string{"owner": "alibaba-cloud-csi-driver", "nodeName": nodeName},
		OmitFilesystemCheck:   omitFsCheck,
	}
	return cfg
}
