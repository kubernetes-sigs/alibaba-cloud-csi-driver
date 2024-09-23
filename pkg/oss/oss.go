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
	"context"
	"os"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	cnfsv1beta1 "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cnfs/v1beta1"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/options"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/version"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"
	mountutils "k8s.io/mount-utils"
)

const (
	driverName = "ossplugin.csi.alibabacloud.com"
)

// OSS the OSS object
type OSS struct {
	endpoint string

	controllerServer csi.ControllerServer
	nodeServer       csi.NodeServer
}

// NewDriver init oss type of csi driver
func NewDriver(endpoint string, m metadata.MetadataProvider, runAsController bool) *OSS {
	klog.Infof("Driver: %v version: %v", driverName, version.VERSION)

	switch os.Getenv(utils.ServiceType) {
	case utils.ProvisionerService:
		runAsController = true
	case utils.PluginService:
		runAsController = false
	}

	d := &OSS{}
	d.endpoint = endpoint

	nodeName := os.Getenv("KUBE_NODE_NAME")
	if !runAsController {
		if nodeName == "" {
			klog.Fatal("env KUBE_NODE_NAME is empty")
		}
	} else {
		nodeName = "controller" // any non-empty value to avoid csi-common panic
	}

	cfg := options.MustGetRestConfig()
	clientset := kubernetes.NewForConfigOrDie(cfg)
	cnfsGetter := cnfsv1beta1.NewCNFSGetter(dynamic.NewForConfigOrDie(cfg))

	configmap, err := clientset.CoreV1().ConfigMaps("kube-system").Get(context.Background(), "csi-plugin", metav1.GetOptions{})
	if err != nil && !errors.IsNotFound(err) {
		klog.Fatalf("failed to get configmap kube-system/csi-plugin: %v", err)
	}

	ossfs := mounter.NewFuseOssfs(configmap, m)
	fusePodManager := mounter.NewFusePodManager(ossfs, clientset)

	if runAsController {
		d.controllerServer = &controllerServer{
			client:         clientset,
			cnfsGetter:     cnfsGetter,
			metadata:       m,
			fusePodManager: fusePodManager,
		}
	} else {
		d.nodeServer = &nodeServer{
			metadata:   m,
			locks:      utils.NewVolumeLocks(),
			nodeName:   nodeName,
			clientset:  clientset,
			cnfsGetter: cnfsGetter,
			rawMounter: mountutils.New(""),
			ossfs:      ossfs,
			GenericNodeServer: common.GenericNodeServer{
				NodeID: nodeName,
			},
		}
	}

	return d
}

type identityServer struct {
	common.GenericIdentityServer
}

func newIdentityServer() *identityServer {
	return &identityServer{
		GenericIdentityServer: common.GenericIdentityServer{
			Name: driverName,
		},
	}
}

// Run start a newNodeServer
func (d *OSS) Run() {
	common.RunCSIServer(d.endpoint, newIdentityServer(), d.controllerServer, d.nodeServer)
}
