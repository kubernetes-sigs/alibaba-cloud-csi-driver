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

package mem

import (
	"os"
	"strings"

	"github.com/container-storage-interface/spec/lib/go/csi"
	csicommon "github.com/kubernetes-csi/drivers/pkg/csi-common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/options"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// MEM the LVM struct
type MEM struct {
	driver           *csicommon.CSIDriver
	endpoint         string
	idServer         *identityServer
	nodeServer       csi.NodeServer
	controllerServer *controllerServer

	cap   []*csi.VolumeCapability_AccessMode
	cscap []*csi.ControllerServiceCapability
}

const (
	driverName = "memplugin.csi.alibabacloud.com"
	csiVersion = "1.0.0"
)

// Init checks for the persistent volume file and loads all found volumes
// into a memory structure
func initDriver() {
}

// NewDriver create the identity/node/controller server and disk driver
func NewDriver(nodeID, endpoint string) *MEM {
	initDriver()
	tmpmem := &MEM{}
	tmpmem.endpoint = endpoint

	if nodeID == "" {
		nodeID = utils.RetryGetMetaData(InstanceID)
		log.Infof("Use node id : %s", nodeID)
	}
	GlobalConfigSet("", nodeID, driverName)
	csiDriver := csicommon.NewCSIDriver(driverName, csiVersion, nodeID)
	tmpmem.driver = csiDriver
	tmpmem.driver.AddControllerServiceCapabilities([]csi.ControllerServiceCapability_RPC_Type{
		csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME,
		csi.ControllerServiceCapability_RPC_PUBLISH_UNPUBLISH_VOLUME,
		csi.ControllerServiceCapability_RPC_EXPAND_VOLUME,
	})
	tmpmem.driver.AddVolumeCapabilityAccessModes([]csi.VolumeCapability_AccessMode_Mode{csi.VolumeCapability_AccessMode_MULTI_NODE_MULTI_WRITER})

	// Create GRPC servers
	tmpmem.idServer = newIdentityServer(tmpmem.driver)
	tmpmem.nodeServer = NewNodeServer(tmpmem.driver, nodeID)
	tmpmem.controllerServer = newControllerServer(tmpmem.driver)

	return tmpmem
}

// Run start a new server
func (mem *MEM) Run() {
	log.Infof("Driver: %v ", driverName)
	common.RunCSIServer(mem.endpoint, mem.idServer, mem.controllerServer, mem.nodeServer)
}

// GlobalConfigSet set Global Config
func GlobalConfigSet(region, nodeID, driverName string) {
	// Global Config set
	cfg, err := clientcmd.BuildConfigFromFlags(options.MasterURL, options.Kubeconfig)
	if err != nil {
		log.Fatalf("Error building kubeconfig: %s", err.Error())
	}
	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		log.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	nodeName := os.Getenv("KUBE_NODE_NAME")
	kmemEnable := false
	nodeInfo, err := kubeClient.CoreV1().Nodes().Get(context.Background(), nodeName, metav1.GetOptions{})
	if err != nil {
		log.Fatalf("Describe node %s with error: %s", nodeName, err.Error())
	} else {
		if value, ok := nodeInfo.Labels[KmemNodeLabel]; ok {
			nodePmemType := strings.TrimSpace(value)
			if nodePmemType == KmemValue {
				kmemEnable = true
			} else {
				log.Errorf("GlobalConfigSet: unknown pemeType: %s", nodePmemType)
			}
		}
	}

	remoteProvision := true
	remoteConfig := os.Getenv("LOCAL_CONTROLLER_PROVISION")
	if strings.ToLower(remoteConfig) == "false" {
		remoteProvision = false
	}

	// Global Config Set
	GlobalConfigVar = GlobalConfig{
		Region:              region,
		NodeID:              nodeID,
		Scheduler:           driverName,
		KmemEnable:          kmemEnable,
		ControllerProvision: remoteProvision,
		KubeClient:          kubeClient,
	}
	log.Infof("Local Plugin Global Config is: %v", GlobalConfigVar)

}
