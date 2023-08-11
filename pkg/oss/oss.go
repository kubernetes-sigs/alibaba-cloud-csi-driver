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

package oss

import (
	"os"

	"github.com/container-storage-interface/spec/lib/go/csi"
	csicommon "github.com/kubernetes-csi/drivers/pkg/csi-common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/options"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

const (
	driverName = "ossplugin.csi.alibabacloud.com"
)

var (
	version = "1.0.0"
)

// OSS the OSS object
type OSS struct {
	driver   *csicommon.CSIDriver
	endpoint string
}

// NewDriver init oss type of csi driver
func NewDriver(nodeID, endpoint string) *OSS {
	log.Infof("Driver: %v version: %v", driverName, version)

	d := &OSS{}
	d.endpoint = endpoint

	if nodeID == "" {
		nodeID = utils.RetryGetMetaData(InstanceID)
		log.Infof("Use node id : %s", nodeID)
	}
	csiDriver := csicommon.NewCSIDriver(driverName, version, nodeID)
	csiDriver.AddVolumeCapabilityAccessModes([]csi.VolumeCapability_AccessMode_Mode{
		csi.VolumeCapability_AccessMode_MULTI_NODE_MULTI_WRITER,
	})
	csiDriver.AddControllerServiceCapabilities([]csi.ControllerServiceCapability_RPC_Type{
		csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME,
		csi.ControllerServiceCapability_RPC_UNKNOWN,
	})

	d.driver = csiDriver
	return d
}

func newControllerServer(driver *csicommon.CSIDriver) csi.ControllerServer {
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatalf("Create create kube config is failed, err: %v", err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Create client set is failed, err: %v", err)
	}

	crdClient, err := dynamic.NewForConfig(config)
	if err != nil {
		log.Fatalf("Create dynamic client is failed, err: %v", err)
	}

	c := &controllerServer{
		client:                  clientset,
		DefaultControllerServer: csicommon.NewDefaultControllerServer(driver),
		crdClient:               crdClient,
	}
	return c
}

// newNodeServer init oss type of csi nodeServer
func newNodeServer(driver *csicommon.CSIDriver) *nodeServer {
	nodeName := os.Getenv("KUBE_NODE_NAME")
	if nodeName == "" {
		log.Fatal("env KUBE_NODE_NAME is empty")
	}

	cfg, err := clientcmd.BuildConfigFromFlags(options.MasterURL, options.Kubeconfig)
	if err != nil {
		log.Fatalf("Build kubeconfig is failed, err: %s", err.Error())
	}

	clientset, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		log.Fatalf("Create client set is failed, err: %v", err)
	}
	configmap, err := clientset.CoreV1().ConfigMaps("kube-system").Get(context.Background(), "csi-plugin", metav1.GetOptions{})
	if err != nil {
		log.Fatalf("failed to get configmap kube-system/csi-plugin: %v", err)
	}

	crdClient, err := dynamic.NewForConfig(cfg)
	if err != nil {
		log.Fatalf("Create crd client is failed, err: %v", err)
	}

	return &nodeServer{
		DefaultNodeServer: csicommon.NewDefaultNodeServer(driver),
		dynamicClient:     crdClient,
		sharedPathLock:    utils.NewVolumeLocks(),
		ossfsMounterFac:   mounter.NewContainerizedFuseMounterFactory(mounter.FuseOssfs, configmap, clientset, nodeName),
	}
}

// Run start a newNodeServer
func (d *OSS) Run() {
	common.RunCSIServer(d.endpoint, csicommon.NewDefaultIdentityServer(d.driver), newControllerServer(d.driver), newNodeServer(d.driver))
}
