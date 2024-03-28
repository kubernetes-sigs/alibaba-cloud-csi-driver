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
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	snapClientset "github.com/kubernetes-csi/external-snapshotter/client/v8/clientset/versioned"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/disk/desc"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/disk/waitstatus"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/features"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/options"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/version"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
	"k8s.io/utils/clock"
)

// PluginFolder defines the location of diskplugin
const (
	driverType              = "disk"
	driverName              = "diskplugin.csi.alibabacloud.com"
	TopologyZoneKey         = "topology." + driverName + "/zone"
	TopologyMultiZonePrefix = TopologyZoneKey + "-"
)

// DISK the DISK object
type DISK struct {
	endpoint string
	servers  common.Servers
}

// GlobalConfig save global values for plugin
type GlobalConfig struct {
	EcsClient            *ecs.Client
	Region               string
	NodeID               string
	DiskTagEnable        bool
	ADControllerEnable   bool
	DetachDisabled       bool
	MetricEnable         bool
	DetachBeforeAttach   bool
	DetachBeforeDelete   bool
	DiskBdfEnable        bool
	ClientSet            *kubernetes.Clientset
	ClusterID            string
	ControllerService    bool
	BdfHealthCheck       bool
	CheckBDFHotPlugin    bool
	SnapClient           *snapClientset.Clientset
	NodeMultiZoneEnable  bool
	WaitBeforeAttach     bool
	AddonVMFatalEvents   []string
	RequestBaseInfo      map[string]string
	SnapshotBeforeDelete bool
	OmitFilesystemCheck  bool
	DiskAllowAllType     bool
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
func NewDriver(m metadata.MetadataProvider, endpoint string, serviceType utils.ServiceType) *DISK {
	initDriver()
	tmpdisk := &DISK{}
	tmpdisk.endpoint = endpoint

	// Config Global vars
	csiCfg := GlobalConfigSet(m)

	if serviceType&utils.Node != 0 {
		GlobalConfigVar.NodeID = metadata.MustGet(m, metadata.InstanceID)
		DefaultDeviceManager.DisableSerial = IsVFNode()
	} else {
		GlobalConfigVar.NodeID = "not-retrieved" // make csi-common happy
	}

	// Init ECS Client
	accessControl := utils.GetAccessControl()
	client := newEcsClient(metadata.MustGet(m, metadata.RegionID), accessControl)
	if accessControl.UseMode == utils.EcsRAMRole || accessControl.UseMode == utils.ManagedToken {
		klog.Infof("Starting csi-plugin with sts.")
	} else {
		klog.Infof("Starting csi-plugin without sts.")
	}
	GlobalConfigVar.EcsClient = client

	// Create GRPC servers
	var servers common.Servers
	servers.IdentityServer = NewIdentityServer()
	if serviceType&utils.Controller != 0 {
		servers.ControllerServer = NewControllerServer(csiCfg)
	}
	if serviceType&utils.Node != 0 {
		servers.NodeServer = NewNodeServer(m)
	}
	if features.FunctionalMutableFeatureGate.Enabled(features.EnableVolumeGroupSnapshots) {
		servers.GroupControllerServer = NewGroupControllerServer()
	}
	tmpdisk.servers = servers

	return tmpdisk
}

// Run start a new NodeServer
func (disk *DISK) Run() {
	klog.Infof("Starting csi-plugin Driver: %v version: %v", driverName, version.VERSION)
	common.RunCSIServer(driverType, disk.endpoint, disk.servers)
}

// GlobalConfigSet set Global Config
func GlobalConfigSet(m metadata.MetadataProvider) utils.Config {
	configMapName := "csi-plugin"

	// Global Configs Set
	cfg, err := clientcmd.BuildConfigFromFlags(options.MasterURL, options.Kubeconfig)
	if err != nil {
		klog.Fatalf("Error building kubeconfig: %s", err.Error())
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
		klog.Fatalf("Error building kubernetes snapclientset: %s", err.Error())
	}

	cfg.ContentType = runtime.ContentTypeProtobuf
	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	csiCfg := utils.Config{}
	configMap, err := kubeClient.CoreV1().ConfigMaps("kube-system").Get(context.Background(), configMapName, metav1.GetOptions{})
	if err != nil {
		klog.Infof("Not found configmap named as csi-plugin under kube-system, with: %v", err)
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

	// Global Config Set
	GlobalConfigVar = GlobalConfig{
		Region: metadata.MustGet(m, metadata.RegionID),
		ADControllerEnable: features.FunctionalMutableFeatureGate.Enabled(features.DiskADController) ||
			csiCfg.GetBool("disk-adcontroller-enable", "DISK_AD_CONTROLLER", false),
		DiskTagEnable:        csiCfg.GetBool("disk-tag-by-plugin", "DISK_TAGED_BY_PLUGIN", false),
		DiskBdfEnable:        csiCfg.GetBool("disk-bdf-enable", "DISK_BDF_ENABLE", false),
		MetricEnable:         csiCfg.GetBool("disk-metric-by-plugin", "DISK_METRIC_BY_PLUGIN", true),
		DetachDisabled:       csiCfg.GetBool("disk-detach-disable", "DISK_DETACH_DISABLE", false),
		DetachBeforeDelete:   csiCfg.GetBool("disk-detach-before-delete", "DISK_DETACH_BEFORE_DELETE", true),
		DetachBeforeAttach:   csiCfg.GetBool("disk-detach-before-attach", "DISK_FORCE_DETACHED", true),
		ClientSet:            kubeClient,
		SnapClient:           snapClient,
		ClusterID:            os.Getenv("CLUSTER_ID"),
		BdfHealthCheck:       csiCfg.GetBool("bdf-health-check", "BDF_HEALTH_CHECK", true),
		NodeMultiZoneEnable:  csiCfg.GetBool("node-multi-zone-enable", "NODE_MULTI_ZONE_ENABLE", false),
		WaitBeforeAttach:     csiCfg.GetBool("wait-before-attach", "WAIT_BEFORE_ATTACH", false),
		AddonVMFatalEvents:   fatalEvents,
		SnapshotBeforeDelete: csiCfg.GetBool("volume-del-auto-snap", "VOLUME_DEL_AUTO_SNAP", false),
		RequestBaseInfo:      map[string]string{"owner": "alibaba-cloud-csi-driver", "nodeName": nodeName},
		OmitFilesystemCheck:  csiCfg.GetBool("disable-fs-check", "DISABLE_FS_CHECK", false),
		DiskAllowAllType:     csiCfg.GetBool("disk-allow-all-type", "DISK_ALLOW_ALL_TYPE", false),
	}
	if csiCfg.GetBool("disk-multi-tenant-enable", "DISK_MULTI_TENANT_ENABLE", false) {
		panic("Disk multi tenant support has been removed. Please remove the related config")
	}
	if GlobalConfigVar.ADControllerEnable {
		klog.Infof("AD-Controller is enabled, CSI Disk Plugin running in AD Controller mode.")
	} else {
		klog.Infof("AD-Controller is disabled, CSI Disk Plugin running in kubelet mode.")
	}
	DefaultDeviceManager.EnableDiskPartition = csiCfg.GetBool("disk-partition-enable", "DISK_PARTITION_ENABLE", true)
	klog.Infof("Starting with GlobalConfigVar: ADControllerEnable(%t), DiskTagEnable(%t), DiskBdfEnable(%t), MetricEnable(%t), DetachDisabled(%t), DetachBeforeDelete(%t), ClusterID(%s)",
		GlobalConfigVar.ADControllerEnable,
		GlobalConfigVar.DiskTagEnable,
		GlobalConfigVar.DiskBdfEnable,
		GlobalConfigVar.MetricEnable,
		GlobalConfigVar.DetachDisabled,
		GlobalConfigVar.DetachBeforeDelete,
		GlobalConfigVar.ClusterID,
	)
	return csiCfg
}

func newDiskStatusWaiter(fromNode bool) waitstatus.StatusWaiter[ecs.Disk] {
	client := desc.Disk{Client: GlobalConfigVar.EcsClient}
	interval := 1 * time.Second
	if fromNode {
		interval = 2 * time.Second // We have many nodes, use longer interval to avoid throttling
	}
	waiter := waitstatus.NewBatched(client, clock.RealClock{}, interval, 3*time.Second)
	go waiter.Run(context.Background())
	return waiter
}
