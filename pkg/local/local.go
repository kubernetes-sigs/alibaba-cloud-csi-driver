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
	"context"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-csi/drivers/pkg/csi-common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local/types"
	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"strings"
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

const (
	defaultDriverName = "localplugin.csi.alibabacloud.com"
	localDriverName   = "localplugin.csi.alibabacloud.com"
	yodaDriverName    = "yodaplugin.csi.alibabacloud.com"
	csiVersion        = "1.0.0"
)

// PmemSupportType ...
var PmemSupportType = []string{types.PmemLVMType, types.PmemDirectType, types.PmemKmemType, types.PmemQuotaPathType}

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
	// Global Configs Set
	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		log.Fatalf("Error building kubeconfig: %s", err.Error())
	}
	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		log.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	nodeName := os.Getenv("KUBE_NODE_NAME")
	pmemEnable := false
	pmeType := ""
	nodeInfo, err := kubeClient.CoreV1().Nodes().Get(context.Background(), nodeName, metav1.GetOptions{})
	if err != nil {
		log.Errorf("Describe node %s with error: %s", nodeName, err.Error())
	} else {
		if value, ok := nodeInfo.Labels[types.PmemNodeLable]; ok {
			nodePmemType := strings.TrimSpace(value)
			pmemEnable = true
			for _, supportPmemType := range PmemSupportType {
				if nodePmemType == supportPmemType {
					pmeType = supportPmemType
				}
			}
			if pmeType == "" {
				log.Fatalf("GlobalConfigSet: unknown pemeType: %s", nodePmemType)
			}
		}
		log.Infof("GlobalConfigSet: Describe node %s and Set PMEM to %v, %s", nodeName, pmemEnable, pmeType)
	}

	grpcProvision := true
	remoteConfig := os.Getenv("LOCAL_GRPC_PROVISION")
	if strings.ToLower(remoteConfig) == "false" {
		grpcProvision = false
	}

	hostNameAsTop := false
	hostNameEnv := os.Getenv("LOCAL_HOSTNAME_AS_TOPO")
	if strings.ToLower(hostNameEnv) == "true" {
		hostNameAsTop = true
	}

	topoKeyDefine := TopologyNodeKey
	topoKeyStr := os.Getenv("LOCAL_TOPO_KEY_DEFINED")
	if topoKeyStr != "" {
		topoKeyDefine = topoKeyStr
	}

	// Global Config Set
	types.GlobalConfigVar = types.GlobalConfig{
		Region:         region,
		NodeID:         nodeID,
		Scheduler:      driverName,
		PmemEnable:     pmemEnable,
		PmemType:       pmeType,
		GrpcProvision:  grpcProvision,
		KubeClient:     kubeClient,
		HostNameAsTopo: hostNameAsTop,
		TopoKeyDefine:  topoKeyDefine,
	}
	log.Infof("Local Plugin Global Config is: %v", types.GlobalConfigVar)
}
