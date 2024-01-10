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

package lvm

import (
	"github.com/container-storage-interface/spec/lib/go/csi"
	csicommon "github.com/kubernetes-csi/drivers/pkg/csi-common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/version"
	log "github.com/sirupsen/logrus"
)

// LVM the LVM struct
type LVM struct {
	driver           *csicommon.CSIDriver
	endpoint         string
	idServer         *identityServer
	nodeServer       csi.NodeServer
	controllerServer *controllerServer

	cap   []*csi.VolumeCapability_AccessMode
	cscap []*csi.ControllerServiceCapability
}

const (
	driverName = "lvmplugin.csi.alibabacloud.com"
)

// Init checks for the persistent volume file and loads all found volumes
// into a memory structure
func initDriver() {
}

// NewDriver create the identity/node/controller server and disk driver
func NewDriver(nodeID, endpoint string) *LVM {
	initDriver()
	tmplvm := &LVM{}
	tmplvm.endpoint = endpoint

	if nodeID == "" {
		nodeID, _ = utils.GetMetaData(InstanceID)
		log.Infof("Use node id : %s", nodeID)
	}
	csiDriver := csicommon.NewCSIDriver(driverName, version.VERSION, nodeID)
	tmplvm.driver = csiDriver
	tmplvm.driver.AddControllerServiceCapabilities([]csi.ControllerServiceCapability_RPC_Type{
		csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME,
		csi.ControllerServiceCapability_RPC_PUBLISH_UNPUBLISH_VOLUME,
		csi.ControllerServiceCapability_RPC_EXPAND_VOLUME,
	})
	tmplvm.driver.AddVolumeCapabilityAccessModes([]csi.VolumeCapability_AccessMode_Mode{csi.VolumeCapability_AccessMode_MULTI_NODE_MULTI_WRITER})

	// Create GRPC servers
	tmplvm.idServer = newIdentityServer(tmplvm.driver)
	tmplvm.nodeServer = NewNodeServer(tmplvm.driver, nodeID)
	tmplvm.controllerServer = newControllerServer(tmplvm.driver)

	return tmplvm
}

// Run start a new server
func (lvm *LVM) Run() {
	log.Infof("Driver: %v ", driverName)
	common.RunCSIServer(lvm.endpoint, lvm.idServer, lvm.controllerServer, lvm.nodeServer)
}
