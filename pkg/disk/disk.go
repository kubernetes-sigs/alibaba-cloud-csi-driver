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
	"os"
	"strings"
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
	TagDisk        string
	ADControllerEn bool
}

// GlobalConfigVar define global variable
var GlobalConfigVar GlobalConfig

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

	accessKeyID, accessSecret, accessToken := GetDefaultAK()
	c := newEcsClient(accessKeyID, accessSecret, accessToken)
	if accessToken == "" {
		log.Infof("Starting csi-plugin without sts")
	} else {
		log.Infof("Starting csi-plugin with sts")
	}

	region := os.Getenv("REGION_ID")
	if region == "" {
		region = GetMetaData(RegionIDTag)
	}

	// Global Configs Set
	tagDiskConf := os.Getenv(DiskTagedByPlugin)
	aDControllerEn := IsDeviceMappedDiskIDEnabled()
	if aDControllerEn {
		log.Infof("Device/DiskId mapping is enabled, CSI Disk Plugin running in AD Controller mode.")
	} else {
		log.Infof("Device/DiskId mapping isn't enabled, CSI Disk Plugin running in kubelet mode.")
	}
	adEnable := os.Getenv(DiskAttachByController)
	if adEnable == "true" || adEnable == "yes" {
		aDControllerEn = true
	} else if adEnable == "false" || adEnable == "no" {
		aDControllerEn = false
	}

	GlobalConfigVar = GlobalConfig{
		EcsClient:      c,
		Region:         region,
		CanAttach:      true,
		ADControllerEn: aDControllerEn,
		TagDisk:        strings.ToLower(tagDiskConf),
	}

	// Create GRPC servers
	tmpdisk.idServer = NewIdentityServer(tmpdisk.driver)
	tmpdisk.controllerServer = NewControllerServer(tmpdisk.driver, c, region)

	if !runAsController {
		tmpdisk.nodeServer = NewNodeServer(tmpdisk.driver, c)
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
