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

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/container-storage-interface/spec/lib/go/csi"
	csicommon "github.com/kubernetes-csi/drivers/pkg/csi-common"
	snapClientset "github.com/kubernetes-csi/external-snapshotter/client/v4/clientset/versioned"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/features"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/options"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/version"
	log "github.com/sirupsen/logrus"
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
}

// GlobalConfig save global values for plugin
type GlobalConfig struct {
	EcsClient             *ecs.Client
	Region                string
	NodeID                string
	DiskTagEnable         bool
	AttachDetachSlots     AttachDetachSlots
	ADControllerEnable    bool
	DetachDisabled        bool
	MetricEnable          bool
	DetachBeforeAttach    bool
	DetachBeforeDelete    bool
	DiskBdfEnable         bool
	ClientSet             *kubernetes.Clientset
	ClusterID             string
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
	DiskAllowAllType      bool
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
func NewDriver(m metadata.MetadataProvider, endpoint string, runAsController bool) *DISK {
	initDriver()
	tmpdisk := &DISK{}
	tmpdisk.endpoint = endpoint

	// Config Global vars
	cfg := GlobalConfigSet(m)

	csiDriver := csicommon.NewCSIDriver(driverName, version.VERSION, GlobalConfigVar.NodeID)
	tmpdisk.driver = csiDriver

	// Init ECS Client
	accessControl := utils.GetAccessControl()
	client := newEcsClient(accessControl)
	if accessControl.UseMode == utils.EcsRAMRole || accessControl.UseMode == utils.ManagedToken {
		log.Infof("Starting csi-plugin with sts.")
	} else {
		log.Infof("Starting csi-plugin without sts.")
	}
	GlobalConfigVar.EcsClient = client

	apiExtentionClient, err := crd.NewForConfig(cfg)
	if err != nil {
		log.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	// Create GRPC servers
	tmpdisk.idServer = NewIdentityServer(tmpdisk.driver)
	tmpdisk.controllerServer = NewControllerServer(apiExtentionClient)

	if !runAsController {
		tmpdisk.nodeServer = NewNodeServer(m)
	}

	return tmpdisk
}

// Run start a new NodeServer
func (disk *DISK) Run() {
	log.Infof("Starting csi-plugin Driver: %v version: %v", driverName, version.VERSION)
	common.RunCSIServer(disk.endpoint, disk.idServer, disk.controllerServer, disk.nodeServer)
}

// GlobalConfigSet set Global Config
func GlobalConfigSet(m metadata.MetadataProvider) *restclient.Config {
	configMapName := "csi-plugin"

	// Global Configs Set
	cfg, err := clientcmd.BuildConfigFromFlags(options.MasterURL, options.Kubeconfig)
	if err != nil {
		log.Fatalf("Error building kubeconfig: %s", err.Error())
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
		log.Fatalf("Error building kubernetes snapclientset: %s", err.Error())
	}

	cfg.ContentType = runtime.ContentTypeProtobuf
	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		log.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	csiCfg := utils.Config{}
	configMap, err := kubeClient.CoreV1().ConfigMaps("kube-system").Get(context.Background(), configMapName, metav1.GetOptions{})
	if err != nil {
		log.Infof("Not found configmap named as csi-plugin under kube-system, with: %v", err)
	} else {
		csiCfg.ConfigMap = configMap.Data
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

	nodeName := os.Getenv(kubeNodeName)

	controllerServerType := false
	nodeID := ""
	if os.Getenv(utils.ServiceType) == utils.ProvisionerService {
		controllerServerType = true
		nodeID = "controller" // make csi-common happy
	} else {
		nodeID = metadata.MustGet(m, metadata.InstanceID)
	}

	// Global Config Set
	GlobalConfigVar = GlobalConfig{
		Region: metadata.MustGet(m, metadata.RegionID),
		NodeID: nodeID,
		ADControllerEnable: features.FunctionalMutableFeatureGate.Enabled(features.DiskADController) ||
			csiCfg.GetBool("disk-adcontroller-enable", "DISK_AD_CONTROLLER", false),
		DiskTagEnable:         csiCfg.GetBool("disk-tag-by-plugin", "DISK_TAGED_BY_PLUGIN", false),
		DiskBdfEnable:         csiCfg.GetBool("disk-bdf-enable", "DISK_BDF_ENABLE", false),
		MetricEnable:          csiCfg.GetBool("disk-metric-by-plugin", "DISK_METRIC_BY_PLUGIN", true),
		DetachDisabled:        csiCfg.GetBool("disk-detach-disable", "DISK_DETACH_DISABLE", false),
		DetachBeforeDelete:    csiCfg.GetBool("disk-detach-before-delete", "DISK_DETACH_BEFORE_DELETE", true),
		DetachBeforeAttach:    csiCfg.GetBool("disk-detach-before-attach", "DISK_FORCE_DETACHED", true),
		ClientSet:             kubeClient,
		SnapClient:            snapClient,
		ClusterID:             os.Getenv("CLUSTER_ID"),
		ControllerService:     controllerServerType,
		BdfHealthCheck:        csiCfg.GetBool("bdf-health-check", "BDF_HEALTH_CHECK", true),
		DiskMultiTenantEnable: csiCfg.GetBool("disk-multi-tenant-enable", "DISK_MULTI_TENANT_ENABLE", false),
		NodeMultiZoneEnable:   csiCfg.GetBool("node-multi-zone-enable", "NODE_MULTI_ZONE_ENABLE", false),
		WaitBeforeAttach:      csiCfg.GetBool("wait-before-attach", "WAIT_BEFORE_ATTACH", false),
		AddonVMFatalEvents:    fatalEvents,
		SnapshotBeforeDelete:  csiCfg.GetBool("volume-del-auto-snap", "VOLUME_DEL_AUTO_SNAP", false),
		RequestBaseInfo:       map[string]string{"owner": "alibaba-cloud-csi-driver", "nodeName": nodeName},
		OmitFilesystemCheck:   csiCfg.GetBool("disable-fs-check", "DISABLE_FS_CHECK", false),
		DiskAllowAllType:      csiCfg.GetBool("disk-allow-all-type", "DISK_ALLOW_ALL_TYPE", false),
	}
	if GlobalConfigVar.ADControllerEnable {
		log.Infof("AD-Controller is enabled, CSI Disk Plugin running in AD Controller mode.")
	} else {
		log.Infof("AD-Controller is disabled, CSI Disk Plugin running in kubelet mode.")
	}
	DefaultDeviceManager.EnableDiskPartition = csiCfg.GetBool("disk-partition-enable", "DISK_PARTITION_ENABLE", true)
	DefaultDeviceManager.DisableSerial = !controllerServerType && IsVFNode()
	log.Infof("Starting with GlobalConfigVar: ADControllerEnable(%t), DiskTagEnable(%t), DiskBdfEnable(%t), MetricEnable(%t), DetachDisabled(%t), DetachBeforeDelete(%t), ClusterID(%s)",
		GlobalConfigVar.ADControllerEnable,
		GlobalConfigVar.DiskTagEnable,
		GlobalConfigVar.DiskBdfEnable,
		GlobalConfigVar.MetricEnable,
		GlobalConfigVar.DetachDisabled,
		GlobalConfigVar.DetachBeforeDelete,
		GlobalConfigVar.ClusterID,
	)

	if controllerServerType {
		pDetach := features.FunctionalMutableFeatureGate.Enabled(features.DiskParallelDetach)
		pAttach := features.FunctionalMutableFeatureGate.Enabled(features.DiskParallelAttach)
		GlobalConfigVar.AttachDetachSlots = NewSlots(!pDetach, !pAttach)
		if pAttach {
			log.Infof("Disk parallel attach enabled")
		}
		if pDetach {
			log.Infof("Disk parallel detach enabled")
		}
	} else {
		GlobalConfigVar.AttachDetachSlots = NewSlots(true, true)
	}

	return cfg
}
