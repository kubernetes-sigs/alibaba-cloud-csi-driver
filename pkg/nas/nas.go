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
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/nas/internal"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/version"
	"k8s.io/klog/v2"
)

const (
	driverType = "nas"
	driverName = "nasplugin.csi.alibabacloud.com"
)

// NAS the NAS object
type NAS struct {
	endpoint string
	servers  common.Servers
}

func NewDriver(meta *metadata.Metadata, endpoint string, serviceType utils.ServiceType, csiCfg utils.Config) *NAS {
	klog.Infof("Driver: %v version: %v", driverName, version.VERSION)

	var d NAS
	d.endpoint = endpoint
	var servers common.Servers
	servers.IdentityServer = newIdentityServer()

	if serviceType&utils.Controller != 0 {
		config, err := internal.GetControllerConfig(meta, csiCfg)
		if err != nil {
			klog.Fatalf("Get nas controller config: %v", err)
		}
		cs, err := newControllerServer(config)
		if err != nil {
			klog.Fatalf("Failed to init nas controller server: %v", err)
		}
		servers.ControllerServer = cs
	}
	if serviceType&utils.Node != 0 {
		config, err := internal.GetNodeConfig(csiCfg)
		if err != nil {
			klog.Fatalf("Get nas node config: %v", err)
		}
		servers.NodeServer = newNodeServer(config)
	}
	d.servers = servers

	return &d
}

func (d *NAS) Run() {
	common.RunCSIServer(driverType, d.endpoint, d.servers)
}
