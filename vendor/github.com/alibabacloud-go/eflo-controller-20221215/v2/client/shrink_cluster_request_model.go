// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iShrinkClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ShrinkClusterRequest
	GetClusterId() *string
	SetIgnoreFailedNodeTasks(v bool) *ShrinkClusterRequest
	GetIgnoreFailedNodeTasks() *bool
	SetNodeGroups(v []*ShrinkClusterRequestNodeGroups) *ShrinkClusterRequest
	GetNodeGroups() []*ShrinkClusterRequestNodeGroups
}

type ShrinkClusterRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// i15dfa12e8f27c44f4a006c2c8bb
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Specifies whether to allow skipping failed nodes. Default value: False.
	//
	// example:
	//
	// False
	IgnoreFailedNodeTasks *bool `json:"IgnoreFailedNodeTasks,omitempty" xml:"IgnoreFailedNodeTasks,omitempty"`
	// The node groups.
	NodeGroups []*ShrinkClusterRequestNodeGroups `json:"NodeGroups,omitempty" xml:"NodeGroups,omitempty" type:"Repeated"`
}

func (s ShrinkClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s ShrinkClusterRequest) GoString() string {
	return s.String()
}

func (s *ShrinkClusterRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ShrinkClusterRequest) GetIgnoreFailedNodeTasks() *bool {
	return s.IgnoreFailedNodeTasks
}

func (s *ShrinkClusterRequest) GetNodeGroups() []*ShrinkClusterRequestNodeGroups {
	return s.NodeGroups
}

func (s *ShrinkClusterRequest) SetClusterId(v string) *ShrinkClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *ShrinkClusterRequest) SetIgnoreFailedNodeTasks(v bool) *ShrinkClusterRequest {
	s.IgnoreFailedNodeTasks = &v
	return s
}

func (s *ShrinkClusterRequest) SetNodeGroups(v []*ShrinkClusterRequestNodeGroups) *ShrinkClusterRequest {
	s.NodeGroups = v
	return s
}

func (s *ShrinkClusterRequest) Validate() error {
	if s.NodeGroups != nil {
		for _, item := range s.NodeGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ShrinkClusterRequestNodeGroups struct {
	HyperNodes []*ShrinkClusterRequestNodeGroupsHyperNodes `json:"HyperNodes,omitempty" xml:"HyperNodes,omitempty" type:"Repeated"`
	// The node group ID.
	//
	// example:
	//
	// ng-3b6fbd24b1b845a0
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// The nodes.
	Nodes []*ShrinkClusterRequestNodeGroupsNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
}

func (s ShrinkClusterRequestNodeGroups) String() string {
	return dara.Prettify(s)
}

func (s ShrinkClusterRequestNodeGroups) GoString() string {
	return s.String()
}

func (s *ShrinkClusterRequestNodeGroups) GetHyperNodes() []*ShrinkClusterRequestNodeGroupsHyperNodes {
	return s.HyperNodes
}

func (s *ShrinkClusterRequestNodeGroups) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *ShrinkClusterRequestNodeGroups) GetNodes() []*ShrinkClusterRequestNodeGroupsNodes {
	return s.Nodes
}

func (s *ShrinkClusterRequestNodeGroups) SetHyperNodes(v []*ShrinkClusterRequestNodeGroupsHyperNodes) *ShrinkClusterRequestNodeGroups {
	s.HyperNodes = v
	return s
}

func (s *ShrinkClusterRequestNodeGroups) SetNodeGroupId(v string) *ShrinkClusterRequestNodeGroups {
	s.NodeGroupId = &v
	return s
}

func (s *ShrinkClusterRequestNodeGroups) SetNodes(v []*ShrinkClusterRequestNodeGroupsNodes) *ShrinkClusterRequestNodeGroups {
	s.Nodes = v
	return s
}

func (s *ShrinkClusterRequestNodeGroups) Validate() error {
	if s.HyperNodes != nil {
		for _, item := range s.HyperNodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Nodes != nil {
		for _, item := range s.Nodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ShrinkClusterRequestNodeGroupsHyperNodes struct {
	HyperNodeId *string `json:"HyperNodeId,omitempty" xml:"HyperNodeId,omitempty"`
}

func (s ShrinkClusterRequestNodeGroupsHyperNodes) String() string {
	return dara.Prettify(s)
}

func (s ShrinkClusterRequestNodeGroupsHyperNodes) GoString() string {
	return s.String()
}

func (s *ShrinkClusterRequestNodeGroupsHyperNodes) GetHyperNodeId() *string {
	return s.HyperNodeId
}

func (s *ShrinkClusterRequestNodeGroupsHyperNodes) SetHyperNodeId(v string) *ShrinkClusterRequestNodeGroupsHyperNodes {
	s.HyperNodeId = &v
	return s
}

func (s *ShrinkClusterRequestNodeGroupsHyperNodes) Validate() error {
	return dara.Validate(s)
}

type ShrinkClusterRequestNodeGroupsNodes struct {
	// The node ID.
	//
	// example:
	//
	// e01poc-cn-zmb2ypjdc01
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s ShrinkClusterRequestNodeGroupsNodes) String() string {
	return dara.Prettify(s)
}

func (s ShrinkClusterRequestNodeGroupsNodes) GoString() string {
	return s.String()
}

func (s *ShrinkClusterRequestNodeGroupsNodes) GetNodeId() *string {
	return s.NodeId
}

func (s *ShrinkClusterRequestNodeGroupsNodes) SetNodeId(v string) *ShrinkClusterRequestNodeGroupsNodes {
	s.NodeId = &v
	return s
}

func (s *ShrinkClusterRequestNodeGroupsNodes) Validate() error {
	return dara.Validate(s)
}
