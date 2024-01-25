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
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/log"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/options"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/version"
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

	cap   []*csi.VolumeCapability_AccessMode
	cscap []*csi.ControllerServiceCapability
}

// GlobalConfig save global values for plugin
type GlobalConfig struct {
	EcsClient             *ecs.Client
	Region                string
	NodeID                string
	ZoneID                string
	DiskTagEnable         bool
	AttachMutex           sync.Mutex
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
		tmpdisk.nodeServer = NewNodeServer(tmpdisk.driver, m)
	}

	return tmpdisk
}

// Run start a new NodeServer
func (disk *DISK) Run() {
	log.Log.Infof("Starting csi-plugin Driver: %v version: %v", driverName, version.VERSION)
	common.RunCSIServer(disk.endpoint, disk.idServer, disk.controllerServer, disk.nodeServer)
}

// GlobalConfigSet set Global Config
func GlobalConfigSet(m metadata.MetadataProvider) *restclient.Config {
	configMapName := "csi-plugin"

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

	csiCfg := utils.Config{}
	configMap, err := kubeClient.CoreV1().ConfigMaps("kube-system").Get(context.Background(), configMapName, metav1.GetOptions{})
	if err != nil {
		log.Log.Infof("Not found configmap named as csi-plugin under kube-system, with: %v", err)
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
	nodeInfo, err := kubeClient.CoreV1().Nodes().Get(context.Background(), nodeName, metav1.GetOptions{})
	if err != nil {
		log.Log.Errorf("GlobalConfigSet: get node %s with error: %s", nodeName, err.Error())
	} else {
		if value, ok := nodeInfo.Labels["alibabacloud.com/container-runtime"]; ok && strings.TrimSpace(value) == "Sandboxed-Container.runv" {
			if value, ok := nodeInfo.Labels["alibabacloud.com/container-runtime-version"]; ok && strings.HasPrefix(strings.TrimSpace(value), "1.") {
				runtimeValue = MixRunTimeMode
			}
		}
		log.Log.Infof("Describe node %s and Set RunTimeClass to %s", nodeName, runtimeValue)
	}
	runtimeEnv := os.Getenv("RUNTIME")
	if runtimeEnv == MixRunTimeMode {
		runtimeValue = MixRunTimeMode
	}
	clustID := os.Getenv("CLUSTER_ID")

	controllerServerType := false
	if os.Getenv(utils.ServiceType) == utils.ProvisionerService {
		controllerServerType = true
	}

	// Global Config Set
	GlobalConfigVar = GlobalConfig{
		Region:                metadata.MustGet(m, metadata.RegionID),
		NodeID:                metadata.MustGet(m, metadata.InstanceID),
		ZoneID:                metadata.MustGet(m, metadata.ZoneID),
		ADControllerEnable:    csiCfg.GetBool("disk-adcontroller-enable", "DISK_AD_CONTROLLER", false),
		DiskTagEnable:         csiCfg.GetBool("disk-tag-by-plugin", "DISK_TAGED_BY_PLUGIN", false),
		DiskBdfEnable:         csiCfg.GetBool("disk-bdf-enable", "DISK_BDF_ENABLE", false),
		MetricEnable:          csiCfg.GetBool("disk-metric-by-plugin", "DISK_METRIC_BY_PLUGIN", true),
		RunTimeClass:          runtimeValue,
		DetachDisabled:        csiCfg.GetBool("disk-detach-disable", "DISK_DETACH_DISABLE", false),
		DetachBeforeDelete:    csiCfg.GetBool("disk-detach-before-delete", "DISK_DETACH_BEFORE_DELETE", true),
		DetachBeforeAttach:    csiCfg.GetBool("disk-detach-before-attach", "DISK_FORCE_DETACHED", true),
		ClientSet:             kubeClient,
		SnapClient:            snapClient,
		FilesystemLosePercent: fileSystemLosePercent,
		ClusterID:             clustID,
		DiskPartitionEnable:   csiCfg.GetBool("disk-partition-enable", "DISK_PARTITION_ENABLE", true),
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
		log.Log.Infof("AD-Controller is enabled, CSI Disk Plugin running in AD Controller mode.")
	} else {
		log.Log.Infof("AD-Controller is disabled, CSI Disk Plugin running in kubelet mode.")
	}
	log.Log.Infof("Starting with GlobalConfigVar: ADControllerEnable(%t), DiskTagEnable(%t), DiskBdfEnable(%t), MetricEnable(%t), RunTimeClass(%s), DetachDisabled(%t), DetachBeforeDelete(%t), ClusterID(%s)",
		GlobalConfigVar.ADControllerEnable,
		GlobalConfigVar.DiskTagEnable,
		GlobalConfigVar.DiskBdfEnable,
		GlobalConfigVar.MetricEnable,
		GlobalConfigVar.RunTimeClass,
		GlobalConfigVar.DetachDisabled,
		GlobalConfigVar.DetachBeforeDelete,
		GlobalConfigVar.ClusterID,
	)

	return cfg
}
