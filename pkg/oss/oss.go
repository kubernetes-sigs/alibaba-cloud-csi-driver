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

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	cnfsv1beta1 "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cnfs/v1beta1"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/oss"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/options"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/version"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"
	mountutils "k8s.io/mount-utils"
)

const (
	driverType = "oss"
	driverName = "ossplugin.csi.alibabacloud.com"
)

// OSS the OSS object
type OSS struct {
	endpoint string
	servers  common.Servers
}

// NewDriver init oss type of csi driver
func NewDriver(endpoint string, m metadata.MetadataProvider, serviceType utils.ServiceType) *OSS {
	klog.Infof("Driver: %v version: %v", driverName, version.VERSION)

	d := &OSS{}
	d.endpoint = endpoint

	nodeName := os.Getenv("KUBE_NODE_NAME")
	if serviceType&utils.Node > 0 {
		if nodeName == "" {
			klog.Fatal("env KUBE_NODE_NAME is empty")
		}
	} else {
		nodeName = "controller" // any non-empty value to avoid csi-common panic
	}

	cfg := options.MustGetRestConfig()
	clientset := kubernetes.NewForConfigOrDie(cfg)
	crdCfg := options.GetRestConfigForCRD(*cfg)
	cnfsGetter := cnfsv1beta1.NewCNFSGetter(dynamic.NewForConfigOrDie(crdCfg))

	configmap := utils.DefaultConfig()

	ossfs := oss.NewFuseOssfs(&configmap, m)
	ossfs2 := oss.NewFuseOssfs2(&configmap, m)
	fusePodManagers := map[string]*oss.OSSFusePodManager{
		OssFsType:  oss.NewOSSFusePodManager(ossfs, clientset),
		OssFs2Type: oss.NewOSSFusePodManager(ossfs2, clientset),
	}

	var servers common.Servers
	servers.IdentityServer = newIdentityServer()

	if serviceType&utils.Controller != 0 {
		servers.ControllerServer = &controllerServer{
			client:          clientset,
			cnfsGetter:      cnfsGetter,
			metadata:        m,
			fusePodManagers: fusePodManagers,
		}
	}
	if serviceType&utils.Node != 0 {
		servers.NodeServer = &nodeServer{
			metadata:        m,
			locks:           utils.NewVolumeLocks(),
			nodeName:        nodeName,
			clientset:       clientset,
			cnfsGetter:      cnfsGetter,
			rawMounter:      mountutils.NewWithoutSystemd(""),
			fusePodManagers: fusePodManagers,
			GenericNodeServer: common.GenericNodeServer{
				NodeID: nodeName,
			},
		}
	}
	d.servers = servers

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
	common.RunCSIServer(driverType, d.endpoint, d.servers)
}
