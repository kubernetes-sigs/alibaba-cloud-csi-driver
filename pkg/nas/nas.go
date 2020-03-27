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
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-csi/drivers/pkg/csi-common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"strings"
)

const (
	driverName = "nasplugin.csi.alibabacloud.com"
	// InstanceID is instance id
	InstanceID = "instance-id"
)

var (
	version    = "1.0.0"
	masterURL  string
	kubeconfig string
	// GlobalConfigVar Global Config
	GlobalConfigVar GlobalConfig
)

// GlobalConfig save global values for plugin
type GlobalConfig struct {
	Region             string
	NasTagEnable       bool
	ADControllerEnable bool
	MetricEnable       bool
	RunTimeClass       string
}

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

	// Global Configs Set
	GlobalConfigSet()

	d.driver = csiDriver
	accessKeyID, accessSecret, accessToken := utils.GetDefaultAK()
	c := newNasClient(accessKeyID, accessSecret, accessToken)
	region := os.Getenv("REGION_ID")
	if region == "" {
		region = GetMetaData(RegionTag)
	}
	d.controllerServer = NewControllerServer(d.driver, c, region)

	return d
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

// GlobalConfigSet set global config
func GlobalConfigSet() {
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
	isNasMetricEnable := false

	configMap, err := kubeClient.CoreV1().ConfigMaps("kube-system").Get(configMapName, metav1.GetOptions{})
	if err != nil {
		log.Infof("Not found configmap named as csi-plugin under kube-system, with error: %v", err)
	} else {
		if value, ok := configMap.Data["nas-metric-enable"]; ok {
			if value == "enable" || value == "yes" || value == "true" {
				log.Infof("Nas Metric is enabled by configMap(%s).", value)
				isNasMetricEnable = true
			}
		}
	}

	metricNasConf := os.Getenv(NasMetricByPlugin)
	if metricNasConf == "true" || metricNasConf == "yes" {
		isNasMetricEnable = true
	} else if metricNasConf == "false" || metricNasConf == "no" {
		isNasMetricEnable = false
	}

	nodeName := os.Getenv("KUBE_NODE_NAME")
	runtimeValue := "runc"
	nodeInfo, err := kubeClient.CoreV1().Nodes().Get(nodeName, metav1.GetOptions{})
	if err != nil {
		log.Errorf("Describe node %s with error: %s", nodeName, err.Error())
	} else {
		if value, ok := nodeInfo.Labels["alibabacloud.com/container-runtime"]; ok && strings.TrimSpace(value) == "Sandboxed-Container.runv" {
			runtimeValue = MixRunTimeMode
		}
		log.Infof("Describe node %s and set RunTimeClass to %s", nodeName, runtimeValue)
	}

	GlobalConfigVar.MetricEnable = isNasMetricEnable
	GlobalConfigVar.RunTimeClass = runtimeValue

}
