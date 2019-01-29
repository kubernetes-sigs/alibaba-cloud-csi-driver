/*
Copyright 2018 The Kubernetes Authors.

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
	"os"
	"path"

	log "github.com/Sirupsen/logrus"
	"github.com/container-storage-interface/spec/lib/go/csi/v0"
	"github.com/denverdino/aliyungo/common"
	"github.com/denverdino/aliyungo/ecs"
	"github.com/denverdino/aliyungo/metadata"
	"github.com/kubernetes-csi/drivers/pkg/csi-common"
)

// PluginFolder defines the location of diskplugin
const (
	PluginFolder = "/var/lib/kubelet/plugins/csi-diskplugin"
	driverName   = "csi-diskplugin"
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

var (
	//diskDriver *disk
	version = "0.3.0"
)

type diskVolume struct {
	Type     string `json:"type"`
	RegionId string `json:"regionId"`
	ZoneId   string `json:"zoneId"`
	FsType   string `json:"fsType"`
	ReadOnly bool   `json:"readOnly"`
}

var diskVolumes map[string]diskVolume

// Init checks for the persistent volume file and loads all found volumes
// into a memory structure
func initDriver() {
	if _, err := os.Stat(path.Join(PluginFolder, "controller")); os.IsNotExist(err) {
		log.Infof("disk: folder %s not found. Creating... \n", path.Join(PluginFolder, "controller"))
		if err := os.Mkdir(path.Join(PluginFolder, "controller"), 0755); err != nil {
			log.Fatalf("Failed to create a controller's volumes folder with error: %v\n", err)
		}
		return
	}
}

func NewDriver(nodeID, endpoint string) *disk {
	initDriver()
	tmpdisk := &disk{}
	tmpdisk.endpoint = endpoint

	csiDriver := csicommon.NewCSIDriver(driverName, version, nodeID)
	tmpdisk.driver = csiDriver

	tmpdisk.driver.AddControllerServiceCapabilities([]csi.ControllerServiceCapability_RPC_Type{
		csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME,
		csi.ControllerServiceCapability_RPC_PUBLISH_UNPUBLISH_VOLUME,
	})

	tmpdisk.driver.AddVolumeCapabilityAccessModes([]csi.VolumeCapability_AccessMode_Mode{csi.VolumeCapability_AccessMode_SINGLE_NODE_WRITER})

	region := getRegion()
	c := newEcsClient(region)
	// Create GRPC servers
	tmpdisk.idServer = NewIdentityServer(tmpdisk.driver)
	tmpdisk.nodeServer = NewNodeServer(tmpdisk.driver, c, region)
	tmpdisk.controllerServer = NewControllerServer(tmpdisk.driver, c, region)

	return tmpdisk
}

func getRegion() common.Region {
	r := common.Region(DEFAULT_REGION)
	m := metadata.NewMetaData(nil)
	region, err := m.Region()
	if err == nil {
		r = common.Region(region)
	}

	return r
}

func newEcsClient(region common.Region) *ecs.Client {
	accessKeyID, accessSecret, accessToken := GetDefaultAK()

	client := ecs.NewClient(accessKeyID, accessSecret)
	if accessToken != "" {
		client.SetSecurityToken(accessToken)
	}

	client.SetRegionID(region)
	client.SetUserAgent(KUBERNETES_ALICLOUD_IDENTITY)

	return client
}

func (disk *disk) Run() {
	log.Infof("Driver: %v version: %v", driverName, version)

	s := csicommon.NewNonBlockingGRPCServer()
	s.Start(disk.endpoint, disk.idServer, disk.controllerServer, disk.nodeServer)
	s.Wait()
}
