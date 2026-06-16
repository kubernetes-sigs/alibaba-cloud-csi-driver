//go:build !windows

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

package bmcpfs

import (
	"strings"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"k8s.io/klog/v2"
)

const (
	driverName = "bmcpfsplugin.csi.alibabacloud.com"

	// keys in volume context or publish context
	_vpcMountTarget = "vpcMountTarget"
	_vscMountTarget = "vscMountTarget"
	_vscID          = "vscId"
	_networkType    = "networkType"
	_path           = "path"
	_mpAutoSwitch   = "mountpointAutoSwitch"

	// CommonNodeIDPrefix is the prefix for common node IDs
	CommonNodeIDPrefix = "common:"
	// VSCNodeIDPrefix is the prefix for ecs node IDs which enabled vsc
	VSCNodeIDPrefix = "vsc:"
	// LingjunNodeIDPrefix is the prefix for lingjun node IDs
	LingjunNodeIDPrefix = "lingjun:"

	// network types of CPFS mount targets
	networkTypeVPC = "vpc"
	networkTypeVSC = "vsc"

	volumeHandleDelimiter = "+"

	FILESET_DESCRIBE_TIMEOUT = 5 * time.Minute
)

// nodeKind identifies the source/topology of a CSI node ID.
type nodeKind int

const (
	// nodeKindUnknown is returned when the node ID has no recognized prefix.
	nodeKindUnknown nodeKind = iota
	// nodeKindCommon represents an ECS node without VSC support; uses VPC mount.
	nodeKindCommon
	// nodeKindVSC represents an ECS node with VSC enabled.
	nodeKindVSC
	// nodeKindLingjun represents a Lingjun (eflo) node.
	nodeKindLingjun
)

// supportsVSC reports whether VSC-based attach should be performed for this node kind.
func (k nodeKind) supportsVSC() bool {
	return k == nodeKindLingjun || k == nodeKindVSC
}

// parseNodeID extracts the underlying instance/node ID from a CSI node ID and
// classifies the node kind based on its prefix. Returns empty instanceID and
// nodeKindUnknown when the prefix is missing or the body is empty.
func parseNodeID(nodeID string) (instanceID string, kind nodeKind) {
	switch {
	case strings.HasPrefix(nodeID, LingjunNodeIDPrefix):
		instanceID = strings.TrimPrefix(nodeID, LingjunNodeIDPrefix)
		kind = nodeKindLingjun
	case strings.HasPrefix(nodeID, VSCNodeIDPrefix):
		instanceID = strings.TrimPrefix(nodeID, VSCNodeIDPrefix)
		kind = nodeKindVSC
	case strings.HasPrefix(nodeID, CommonNodeIDPrefix):
		instanceID = strings.TrimPrefix(nodeID, CommonNodeIDPrefix)
		kind = nodeKindCommon
	default:
		return "", nodeKindUnknown
	}
	if instanceID == "" {
		return "", nodeKindUnknown
	}
	return instanceID, kind
}

func NewServers(meta *metadata.Metadata, endpoint string, serviceType utils.ServiceType) *common.Servers {
	var servers common.Servers
	servers.IdentityServer = newIdentityServer()

	if serviceType&utils.Controller != 0 {
		cs, err := newControllerServer(meta)
		if err != nil {
			klog.Fatalf("Init controller server: %v", err)
		}
		servers.ControllerServer = cs
	}
	if serviceType&utils.Node != 0 {
		ns, err := newNodeServer(meta)
		if err != nil {
			klog.Fatalf("Init node server: %v", err)
		}
		servers.NodeServer = ns
	}

	return &servers
}
