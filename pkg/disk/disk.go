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
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-csi/drivers/pkg/csi-common"
	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"sync"
)

// PluginFolder defines the location of diskplugin
const (
	driverName      = "diskplugin.csi.alibabacloud.com"
	csiVersion      = "1.0.0"
	TopologyZoneKey = "topology." + driverName + "/zone"
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
	EcsClient      *ecs.Client
	Region         string
	AttachMutex    sync.RWMutex
	CanAttach      bool
	TagDisk        bool
	ADControllerEn bool
}

// define global variable
var (
	masterURL       string
	kubeconfig      string
	GlobalConfigVar GlobalConfig
)

// Init checks for the persistent volume file and loads all found volumes
// into a memory structure
func initDriver() {
}

//NewDriver create the identity/node/controller server and disk driver
func NewDriver(nodeID, endpoint string, runAsController bool) *DISK {
	initDriver()
	tmpdisk := &DISK{}
	tmpdisk.endpoint = endpoint

	if nodeID == "" {
		nodeID = GetMetaData(InstanceID)
		log.Infof("Use node id : %s", nodeID)
	}
	csiDriver := csicommon.NewCSIDriver(driverName, csiVersion, nodeID)
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
	accessKeyID, accessSecret, accessToken := GetDefaultAK()
	client := newEcsClient(accessKeyID, accessSecret, accessToken)
	if accessToken == "" {
		log.Infof("Starting csi-plugin without sts")
	} else {
		log.Infof("Starting csi-plugin with sts")
	}

	// Set Region ID
	region := os.Getenv("REGION_ID")
	if region == "" {
		region = GetMetaData(RegionIDTag)
	}

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
	tagDiskConf := ""
	aDControllerEnable := false
	configMap, err := kubeClient.CoreV1().ConfigMaps("kube-system").Get(configMapName, metav1.GetOptions{})
	if err != nil {
		log.Infof("Not found configmap named as csi-plugin under kube-system")
	} else {
		if value, ok := configMap.Data["disk-adcontroller"]; ok {
			if value == "enable" || value == "yes" || value == "true" {
				log.Infof("AD-Controller is enabled by configMap(%s), CSI Disk Plugin running in AD Controller mode.", value)
				aDControllerEnable = true
			} else if value == "disable" || value == "no" || value == "false" {
				log.Infof("AD-Controller is disable by configMap(%s), CSI Disk Plugin running in kubelet mode.", value)
				aDControllerEnable = false
			}
		}
		if value, ok := configMap.Data["patch-disk-tag"]; ok {
			log.Infof("Patch Disk Tag is enabled by configMap(%s).", value)
			tagDiskConf = value
		}
	}

	// Env variables
	adEnable := os.Getenv(DiskAttachByController)
	if adEnable == "true" || adEnable == "yes" {
		log.Infof("AD-Controller is enabled by Env(%s), CSI Disk Plugin running in AD Controller mode.", adEnable)
		aDControllerEnable = true
	} else if adEnable == "false" || adEnable == "no" {
		log.Infof("AD-Controller is disabled by Env(%s), CSI Disk Plugin running in kubelet mode.", adEnable)
		aDControllerEnable = false
	}
	if aDControllerEnable {
		log.Infof("AD-Controller is enabled, CSI Disk Plugin running in AD Controller mode.")
	} else {
		log.Infof("AD-Controller is disabled, CSI Disk Plugin running in kubelet mode.")
	}
	tagDiskConf = os.Getenv(DiskTagedByPlugin)
	tagDiskConfEn := false
	if tagDiskConf == "true" || tagDiskConf == "yes" {
		tagDiskConfEn = true
	} else if tagDiskConf == "false" || tagDiskConf == "no" {
		tagDiskConfEn = false
	}

	// Global Config Set
	GlobalConfigVar = GlobalConfig{
		EcsClient:      client,
		Region:         region,
		CanAttach:      true,
		ADControllerEn: aDControllerEnable,
		TagDisk:        tagDiskConfEn,
	}

	// Create GRPC servers
	tmpdisk.idServer = NewIdentityServer(tmpdisk.driver)
	tmpdisk.controllerServer = NewControllerServer(tmpdisk.driver, client, region)

	if !runAsController {
		tmpdisk.nodeServer = NewNodeServer(tmpdisk.driver, client)
	}

	return tmpdisk
}

// Run start a new NodeServer
func (disk *DISK) Run() {
	log.Infof("Starting csi-plugin Driver: %v version: %v", driverName, csiVersion)
	s := csicommon.NewNonBlockingGRPCServer()
	s.Start(disk.endpoint, disk.idServer, disk.controllerServer, disk.nodeServer)
	s.Wait()
}
