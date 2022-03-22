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
	"context"
	"fmt"
	aliNas "github.com/aliyun/alibaba-cloud-sdk-go/services/nas"
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
	CpfsNfsEnable      bool
	ADControllerEnable bool
	MetricEnable       bool
	NasFakeProvision   bool
	RunTimeClass       string
	NodeID             string
	NodeIP             string
	ClusterID          string
	LosetupEnable      bool
	KubeClient         *kubernetes.Clientset
	NasClient          *aliNas.Client
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
func NewDriver(nodeID, endpoint, serviceType string) *NAS {
	log.Infof("Driver: %v version: %v", driverName, version)

	d := &NAS{}
	d.endpoint = endpoint
	if nodeID == "" {
		nodeID = utils.RetryGetMetaData(InstanceID)
		log.Infof("Use node id : %s", nodeID)
	}
	csiDriver := csicommon.NewCSIDriver(driverName, version, nodeID)
	csiDriver.AddVolumeCapabilityAccessModes([]csi.VolumeCapability_AccessMode_Mode{csi.VolumeCapability_AccessMode_MULTI_NODE_MULTI_WRITER})
	csiDriver.AddControllerServiceCapabilities([]csi.ControllerServiceCapability_RPC_Type{
		csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME,
		csi.ControllerServiceCapability_RPC_PUBLISH_UNPUBLISH_VOLUME,
		csi.ControllerServiceCapability_RPC_EXPAND_VOLUME,
	})

	// Global Configs Set
	GlobalConfigSet(serviceType)

	d.driver = csiDriver

	ac := utils.GetAccessControl()
	c := newNasClient(ac, "")
	region := os.Getenv("REGION_ID")
	limit := os.Getenv("NAS_LIMIT_PERSECOND")
	if limit == "" {
		limit = "2"
	}
	if region == "" {
		region = GetMetaData(RegionTag)
	}
	d.controllerServer = NewControllerServer(d.driver, c, region, limit)

	GlobalConfigVar.NasClient = c
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
func GlobalConfigSet(serviceType string) {
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
	isNasFakeProvisioner := false
	isCpfsNfsEnable := false

	configMap, err := kubeClient.CoreV1().ConfigMaps("kube-system").Get(context.Background(), configMapName, metav1.GetOptions{})
	if err != nil {
		log.Infof("Not found configmap named as csi-plugin under kube-system, with error: %v", err)
	} else {
		if value, ok := configMap.Data["nas-metric-enable"]; ok {
			if value == "enable" || value == "yes" || value == "true" {
				log.Infof("Nas Metric is enabled by configMap(%s).", value)
				isNasMetricEnable = true
			}
		}
		if value, ok := configMap.Data["nas-fake-provision"]; ok {
			if value == "enable" || value == "yes" || value == "true" {
				isNasFakeProvisioner = true
			}
		}
		if value, ok := configMap.Data["cpfs-nas-enable"]; ok {
			if value == "enable" || value == "yes" || value == "true" {
				isCpfsNfsEnable = true
				queryCmd := fmt.Sprintf("%s rpm -qa | grep aliyun-cpfs-utils", NsenterCmd)
				res, _ := utils.Run(queryCmd)
				if len(res) == 0 && serviceType == utils.PluginService {
					cpfsRpm := "aliyun-cpfs-nfs-utils-1.0-2.al.noarch.rpm"
					installCmd := fmt.Sprintf("%s yum localinstall -y /etc/csi-tool/%s", NsenterCmd, cpfsRpm)
					_, err := utils.Run(installCmd)
					if err != nil {
						log.Errorf("Exec cmd %s is failed, err: %v", installCmd, err)
					} else {
						log.Infof("Exec cmd %s is successfully", installCmd)
					}
				}
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
	nodeInfo, err := kubeClient.CoreV1().Nodes().Get(context.Background(), nodeName, metav1.GetOptions{})
	if err != nil {
		log.Errorf("Describe node %s with error: %s", nodeName, err.Error())
	} else {
		if value, ok := nodeInfo.Labels["alibabacloud.com/container-runtime"]; ok && strings.TrimSpace(value) == "Sandboxed-Container.runv" {
			if value, ok := nodeInfo.Labels["alibabacloud.com/container-runtime-version"]; ok && strings.HasPrefix(strings.TrimSpace(value), "1.") {
				runtimeValue = MixRunTimeMode
			}
		}
		log.Infof("Describe node %s and set RunTimeClass to %s", nodeName, runtimeValue)
	}

	if nodeInfo != nil {
		for _, address := range nodeInfo.Status.Addresses {
			if address.Type == "InternalIP" {
				log.Infof("Node InternalIP is: %s", address.Address)
				GlobalConfigVar.NodeIP = address.Address
			}
		}
	}

	GlobalConfigVar.LosetupEnable = false
	losetupEn := os.Getenv("NAS_LOSETUP_ENABLE")
	if losetupEn == "true" || losetupEn == "yes" {
		GlobalConfigVar.LosetupEnable = true
	}

	if GlobalConfigVar.LosetupEnable && GlobalConfigVar.NodeIP == "" {
		log.Fatal("Init GlobalConfigVar with NodeIP Empty, Nas losetup feature may be useless")
	}
	clustID := os.Getenv("CLUSTER_ID")

	GlobalConfigVar.KubeClient = kubeClient
	GlobalConfigVar.MetricEnable = isNasMetricEnable
	GlobalConfigVar.RunTimeClass = runtimeValue
	GlobalConfigVar.NodeID = nodeName
	GlobalConfigVar.ClusterID = clustID
	GlobalConfigVar.NasFakeProvision = isNasFakeProvisioner
	GlobalConfigVar.CpfsNfsEnable = isCpfsNfsEnable
	log.Infof("NAS Global Config: %v", GlobalConfigVar)
}
