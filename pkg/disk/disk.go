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
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-csi/drivers/pkg/csi-common"
	log "github.com/sirupsen/logrus"
	"os"
)

// PluginFolder defines the location of diskplugin
const (
	driverName      = "diskplugin.csi.alibabacloud.com"
	csiVersion      = "1.0.0"
	TopologyZoneKey = "topology." + driverName + "/zone"
)

type disk struct {
	driver           *csicommon.CSIDriver
	endpoint         string
	idServer         csi.IdentityServer
	nodeServer       csi.NodeServer
	controllerServer csi.ControllerServer

	cap   []*csi.VolumeCapability_AccessMode
	cscap []*csi.ControllerServiceCapability
}

// Init checks for the persistent volume file and loads all found volumes
// into a memory structure
func initDriver() {

}

//NewDriver create the identity/node/controller server and disk driver
func NewDriver(nodeID, endpoint string, runAsController bool) *disk {
	initDriver()
	tmpdisk := &disk{}
	tmpdisk.endpoint = endpoint

	if nodeID == "" {
		nodeID = GetMetaData(INSTANCE_ID)
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
		region = GetMetaData(REGIONID_TAG)
	}
	// Create GRPC servers
	tmpdisk.idServer = NewIdentityServer(tmpdisk.driver)
	tmpdisk.controllerServer = NewControllerServer(tmpdisk.driver, c, region)

	if !runAsController {
		tmpdisk.nodeServer = NewNodeServer(tmpdisk.driver, c)
	}

	return tmpdisk
}

func (disk *disk) Run() {
	log.Infof("Starting csi-plugin Driver: %v version: %v", driverName, csiVersion)
	s := csicommon.NewNonBlockingGRPCServer()
	s.Start(disk.endpoint, disk.idServer, disk.controllerServer, disk.nodeServer)
	s.Wait()
}
