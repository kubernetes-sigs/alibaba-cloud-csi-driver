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
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"k8s.io/klog/v2"
)

type nodeServer struct {
	common.GenericNodeServer
}

const LingjunConfigFile = "/host/etc/eflo_config/lingjun_config"

const (
	CommonNodeIDPrefix  = "common:"
	LingjunNodeIDPrefix = "lingjun:"
)

func newNodeServer() (*nodeServer, error) {
	// TODO: Move the implementation for obtaining Lingjun node id to metadata providers.
	var nodeID string
	data, err := os.ReadFile(LingjunConfigFile)
	if err != nil {
		if os.IsNotExist(err) {
			// if not lingjun instance
			nodeID = CommonNodeIDPrefix + os.Getenv(metadata.KUBE_NODE_NAME_ENV)
		} else {
			return nil, fmt.Errorf("read lingjun_config file: %w", err)
		}
	} else {
		var lingjunConfig struct {
			NodeId string `json:"NodeId"`
		}
		if err := json.Unmarshal(data, &lingjunConfig); err != nil {
			return nil, fmt.Errorf("parse lingjun_config: %w", err)
		}
		if lingjunConfig.NodeId == "" {
			return nil, errors.New("unexpected lingjun_config: NodeId is empty")
		}
		nodeID = LingjunNodeIDPrefix + lingjunConfig.NodeId
	}
	klog.Infof("bmcpfsplugin nodeId: %s", nodeID)

	return &nodeServer{
		GenericNodeServer: common.GenericNodeServer{NodeID: nodeID},
	}, nil
}
