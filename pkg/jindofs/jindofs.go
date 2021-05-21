/*
Copyright 2021 The Kubernetes Authors.

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

package jindofs

import (
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-csi/drivers/pkg/csi-common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"time"

	k8smount "k8s.io/utils/mount"
)

const (
	driverName = "jindofsplugin.csi.alibabacloud.com"
)

var (
	version = "1.0.0"
)

// OSS the OSS object
type JindoFS struct {
	driver     *csicommon.CSIDriver
	endpoint   string
	idServer   *csicommon.DefaultIdentityServer
	nodeServer *nodeServer

	cap   []*csi.VolumeCapability_AccessMode
	cscap []*csi.ControllerServiceCapability
}

// NewDriver init jindofs type of csi driver
func NewDriver(nodeID, endpoint string) *JindoFS {
	log.Infof("Driver: %v version: %v", driverName, version)

	d := &JindoFS{}
	d.endpoint = endpoint

	if nodeID == "" {
		nodeID = GetMetaData(InstanceID)
		log.Infof("Use node id : %s", nodeID)
	}
	csiDriver := csicommon.NewCSIDriver(driverName, version, nodeID)
	csiDriver.AddVolumeCapabilityAccessModes([]csi.VolumeCapability_AccessMode_Mode{csi.VolumeCapability_AccessMode_MULTI_NODE_MULTI_WRITER})
	csiDriver.AddControllerServiceCapabilities([]csi.ControllerServiceCapability_RPC_Type{csi.ControllerServiceCapability_RPC_UNKNOWN})

	d.driver = csiDriver

	return d
}

// newNodeServer init oss type of csi nodeServer
func newNodeServer(d *JindoFS) *nodeServer {
	return &nodeServer{
		k8smounter:        k8smount.New(""),
		DefaultNodeServer: csicommon.NewDefaultNodeServer(d.driver),
	}
}

const (
	JindofsCredentialPathInPod  = "/oss-secret/sts-token"
	JindofsCredentialPathOnHost = "/host/etc/jindofs-credentials"
)

// syncCredential syncs the credential file which is mounted from k8s secret into the host path
// on which the node server pod is running periodically to make sure the jindofs-fuse can always
// get the latest credential file from the host path.
func syncCredential() {
	log.Infof("Start to sync jindofs credentials from %v to %v", JindofsCredentialPathInPod, JindofsCredentialPathOnHost)
	for {
		select {
		case _ = <-time.NewTimer(time.Second * 5).C:
			if !utils.IsFileExisting(JindofsCredentialPathInPod) {
				continue
			}
			credentials, err := ioutil.ReadFile(JindofsCredentialPathInPod)
			if err != nil {
				log.Errorf("Failed to read credentials file %v: %v", JindofsCredentialPathInPod, err)
				continue
			}
			if err := ioutil.WriteFile(JindofsCredentialPathOnHost, credentials, os.ModePerm); err != nil {
				log.Errorf("Failed to write credentials file %v: %v", JindofsCredentialPathOnHost, err)
			}
		}
	}
}

// Run start a newNodeServer
func (d *JindoFS) Run() {
	go syncCredential()

	s := csicommon.NewNonBlockingGRPCServer()
	s.Start(d.endpoint,
		csicommon.NewDefaultIdentityServer(d.driver),
		nil,
		//csicommon.NewDefaultControllerServer(d.driver),
		newNodeServer(d))
	s.Wait()
}
