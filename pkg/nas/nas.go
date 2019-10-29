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

package nas

import (
	log "github.com/Sirupsen/logrus"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-csi/drivers/pkg/csi-common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"os"
)

const (
	driverName = "nasplugin.csi.alibabacloud.com"
	// InstanceID is instance id
	InstanceID = "instance-id"
)

var (
	version = "1.0.0"
)

// NAS the NAS object
type NAS struct {
	driver           *csicommon.CSIDriver
	endpoint         string
	idServer         *csicommon.DefaultIdentityServer
	nodeServer       *nodeServer
	controllerServer csi.ControllerServer

	cap   []*csi.VolumeCapability_AccessMode
	cscap []*csi.ControllerServiceCapability
}

//NewDriver create the identity/node/controller server and disk driver
func NewDriver(nodeID, endpoint string) *NAS {
	log.Infof("Driver: %v version: %v", driverName, version)

	d := &NAS{}
	d.endpoint = endpoint
	if nodeID == "" {
		nodeID, _ = utils.GetMetaData(InstanceID)
		log.Infof("Use node id : %s", nodeID)
	}
	csiDriver := csicommon.NewCSIDriver(driverName, version, nodeID)
	csiDriver.AddVolumeCapabilityAccessModes([]csi.VolumeCapability_AccessMode_Mode{csi.VolumeCapability_AccessMode_MULTI_NODE_MULTI_WRITER})
	csiDriver.AddControllerServiceCapabilities([]csi.ControllerServiceCapability_RPC_Type{
		csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME,
		csi.ControllerServiceCapability_RPC_PUBLISH_UNPUBLISH_VOLUME,
	})

	d.driver = csiDriver

	accessKeyID, accessSecret, accessToken := GetDefaultAK()
	c := newNasClient(accessKeyID, accessSecret, accessToken)

	region := os.Getenv("REGION_ID")
	if region == "" {
		region = GetMetaData(RegionTag)
	}
	d.controllerServer = NewControllerServer(d.driver, c, region)

	return d
}

//newNodeServer create the csi node server
func newNodeServer(d *NAS) *nodeServer {
	return &nodeServer{
		DefaultNodeServer: csicommon.NewDefaultNodeServer(d.driver),
	}
}

// Run start a new NodeServer
func (d *NAS) Run() {
	s := csicommon.NewNonBlockingGRPCServer()
	s.Start(d.endpoint,
		csicommon.NewDefaultIdentityServer(d.driver),
		d.controllerServer,
		newNodeServer(d))
	s.Wait()
}
