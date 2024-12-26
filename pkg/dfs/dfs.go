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

package dfs

import (
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/options"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"
)

const (
	driverType = "dfs"
	driverName = "dfsplugin.csi.alibabacloud.com"
)

type DFSDriver struct {
	endpoint string
	common.Servers
}

func NewDriver(m metadata.MetadataProvider, endpoint string, serviceType utils.ServiceType) *DFSDriver {
	d := new(DFSDriver)
	d.endpoint = endpoint
	d.IdentityServer = newIdentityServer()

	if serviceType&utils.Controller != 0 {
		region := metadata.MustGet(m, metadata.RegionID)
		initDFSEndpoint(region)
		d.ControllerServer = newControllerServer(region)
	}

	if serviceType&utils.Node != 0 {
		kubeClient := kubernetes.NewForConfigOrDie(options.MustGetRestConfig())
		nodeId, err := m.Get(metadata.VscAgentID)
		if err != nil {
			klog.ErrorS(err, "Failed to get vsc agent id")
			nodeId = metadata.MustGet(m, metadata.InstanceID)
		}
		klog.Infof("DFS nodeId: %s", nodeId)
		d.NodeServer = newNodeServer(nodeId, kubeClient)
	}

	return d
}

func (d *DFSDriver) Run() {
	common.RunCSIServer(driverType, d.endpoint, d.Servers)
}
