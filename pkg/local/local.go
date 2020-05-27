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

package local

import (
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-csi/drivers/pkg/csi-common"
	log "github.com/sirupsen/logrus"
	"os"
)

// Local the Local struct
type Local struct {
	driver           *csicommon.CSIDriver
	endpoint         string
	idServer         *identityServer
	nodeServer       csi.NodeServer
	controllerServer *controllerServer

	cap   []*csi.VolumeCapability_AccessMode
	cscap []*csi.ControllerServiceCapability
}

// GlobalConfig var
type GlobalConfig struct {
	Region    string
	NodeID    string
	Scheduler string
}

var (
	// GlobalConfigVar var
	GlobalConfigVar GlobalConfig
)

const (
	defaultDriverName = "yodaplugin.csi.alibabacloud.com"
	localDriverName   = "localplugin.csi.alibabacloud.com"
	yodaDriverName    = "yodaplugin.csi.alibabacloud.com"
	csiVersion        = "1.0.0"
)

// Init checks for the persistent volume file and loads all found volumes
// into a memory structure
func initDriver() {
}

// NewDriver create the identity/node/controller server and disk driver
func NewDriver(nodeID, endpoint string) *Local {
	initDriver()
	tmplvm := &Local{}
	tmplvm.endpoint = endpoint

	if nodeID == "" {
		nodeID = GetMetaData(InstanceID)
		log.Infof("Use node id : %s", nodeID)
	}

	// set driver name;
	driverName := defaultDriverName
	tmpValue := os.Getenv("DRIVER_VENDOR")
	if tmpValue != "" {
		driverName = tmpValue
	}

	// GlobalConfig Set
	GlobalConfigSet("", nodeID, driverName)

	csiDriver := csicommon.NewCSIDriver(driverName, csiVersion, nodeID)
	tmplvm.driver = csiDriver
	tmplvm.driver.AddControllerServiceCapabilities([]csi.ControllerServiceCapability_RPC_Type{
		csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME,
		csi.ControllerServiceCapability_RPC_PUBLISH_UNPUBLISH_VOLUME,
		csi.ControllerServiceCapability_RPC_EXPAND_VOLUME,
	})
	tmplvm.driver.AddVolumeCapabilityAccessModes([]csi.VolumeCapability_AccessMode_Mode{csi.VolumeCapability_AccessMode_MULTI_NODE_MULTI_WRITER})

	// Create GRPC servers
	tmplvm.idServer = newIdentityServer(tmplvm.driver)
	tmplvm.nodeServer = NewNodeServer(tmplvm.driver, driverName, nodeID)
	tmplvm.controllerServer = newControllerServer(tmplvm.driver)

	return tmplvm
}

// Run start a new server
func (lvm *Local) Run() {
	server := csicommon.NewNonBlockingGRPCServer()
	server.Start(lvm.endpoint, lvm.idServer, lvm.controllerServer, lvm.nodeServer)
	server.Wait()
}

// GlobalConfigSet set Global Config
func GlobalConfigSet(region, nodeID, driverName string) {
	// Global Config Set
	GlobalConfigVar = GlobalConfig{
		Region:    region,
		NodeID:    nodeID,
		Scheduler: driverName,
	}
}
