// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClusterShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterDescription(v string) *CreateClusterShrinkRequest
	GetClusterDescription() *string
	SetClusterName(v string) *CreateClusterShrinkRequest
	GetClusterName() *string
	SetClusterType(v string) *CreateClusterShrinkRequest
	GetClusterType() *string
	SetComponentsShrink(v string) *CreateClusterShrinkRequest
	GetComponentsShrink() *string
	SetHpnZone(v string) *CreateClusterShrinkRequest
	GetHpnZone() *string
	SetIgnoreFailedNodeTasks(v bool) *CreateClusterShrinkRequest
	GetIgnoreFailedNodeTasks() *bool
	SetNetworksShrink(v string) *CreateClusterShrinkRequest
	GetNetworksShrink() *string
	SetNimizVSwitchesShrink(v string) *CreateClusterShrinkRequest
	GetNimizVSwitchesShrink() *string
	SetNodeGroupsShrink(v string) *CreateClusterShrinkRequest
	GetNodeGroupsShrink() *string
	SetOpenEniJumboFrame(v bool) *CreateClusterShrinkRequest
	GetOpenEniJumboFrame() *bool
	SetResourceGroupId(v string) *CreateClusterShrinkRequest
	GetResourceGroupId() *string
	SetTag(v []*CreateClusterShrinkRequestTag) *CreateClusterShrinkRequest
	GetTag() []*CreateClusterShrinkRequestTag
}

type CreateClusterShrinkRequest struct {
	// Cluster description
	//
	// example:
	//
	// Cluster description
	ClusterDescription *string `json:"ClusterDescription,omitempty" xml:"ClusterDescription,omitempty"`
	// Cluster name
	//
	// example:
	//
	// Standard_Cluster
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// Cluster type
	//
	// example:
	//
	// Lite
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// Components (software instances)
	ComponentsShrink *string `json:"Components,omitempty" xml:"Components,omitempty"`
	// Cluster number
	//
	// example:
	//
	// A1
	HpnZone *string `json:"HpnZone,omitempty" xml:"HpnZone,omitempty"`
	// Whether to allow skipping failed nodes, the default value is False
	//
	// example:
	//
	// False
	IgnoreFailedNodeTasks *bool `json:"IgnoreFailedNodeTasks,omitempty" xml:"IgnoreFailedNodeTasks,omitempty"`
	// Network information
	NetworksShrink *string `json:"Networks,omitempty" xml:"Networks,omitempty"`
	// Node VSwitches
	NimizVSwitchesShrink *string `json:"NimizVSwitches,omitempty" xml:"NimizVSwitches,omitempty"`
	// Node group list
	NodeGroupsShrink *string `json:"NodeGroups,omitempty" xml:"NodeGroups,omitempty"`
	// Whether the network interface supports jumbo frames
	//
	// example:
	//
	// false
	OpenEniJumboFrame *bool `json:"OpenEniJumboFrame,omitempty" xml:"OpenEniJumboFrame,omitempty"`
	// Resource group ID
	//
	// example:
	//
	// rg-aek2xdkc6icwfha
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Resource tags
	Tag []*CreateClusterShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateClusterShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterShrinkRequest) GetClusterDescription() *string {
	return s.ClusterDescription
}

func (s *CreateClusterShrinkRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *CreateClusterShrinkRequest) GetClusterType() *string {
	return s.ClusterType
}

func (s *CreateClusterShrinkRequest) GetComponentsShrink() *string {
	return s.ComponentsShrink
}

func (s *CreateClusterShrinkRequest) GetHpnZone() *string {
	return s.HpnZone
}

func (s *CreateClusterShrinkRequest) GetIgnoreFailedNodeTasks() *bool {
	return s.IgnoreFailedNodeTasks
}

func (s *CreateClusterShrinkRequest) GetNetworksShrink() *string {
	return s.NetworksShrink
}

func (s *CreateClusterShrinkRequest) GetNimizVSwitchesShrink() *string {
	return s.NimizVSwitchesShrink
}

func (s *CreateClusterShrinkRequest) GetNodeGroupsShrink() *string {
	return s.NodeGroupsShrink
}

func (s *CreateClusterShrinkRequest) GetOpenEniJumboFrame() *bool {
	return s.OpenEniJumboFrame
}

func (s *CreateClusterShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateClusterShrinkRequest) GetTag() []*CreateClusterShrinkRequestTag {
	return s.Tag
}

func (s *CreateClusterShrinkRequest) SetClusterDescription(v string) *CreateClusterShrinkRequest {
	s.ClusterDescription = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetClusterName(v string) *CreateClusterShrinkRequest {
	s.ClusterName = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetClusterType(v string) *CreateClusterShrinkRequest {
	s.ClusterType = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetComponentsShrink(v string) *CreateClusterShrinkRequest {
	s.ComponentsShrink = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetHpnZone(v string) *CreateClusterShrinkRequest {
	s.HpnZone = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetIgnoreFailedNodeTasks(v bool) *CreateClusterShrinkRequest {
	s.IgnoreFailedNodeTasks = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetNetworksShrink(v string) *CreateClusterShrinkRequest {
	s.NetworksShrink = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetNimizVSwitchesShrink(v string) *CreateClusterShrinkRequest {
	s.NimizVSwitchesShrink = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetNodeGroupsShrink(v string) *CreateClusterShrinkRequest {
	s.NodeGroupsShrink = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetOpenEniJumboFrame(v bool) *CreateClusterShrinkRequest {
	s.OpenEniJumboFrame = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetResourceGroupId(v string) *CreateClusterShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetTag(v []*CreateClusterShrinkRequestTag) *CreateClusterShrinkRequest {
	s.Tag = v
	return s
}

func (s *CreateClusterShrinkRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateClusterShrinkRequestTag struct {
	// Key
	//
	// example:
	//
	// env-name
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Value
	//
	// example:
	//
	// dev
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateClusterShrinkRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *CreateClusterShrinkRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateClusterShrinkRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateClusterShrinkRequestTag) SetKey(v string) *CreateClusterShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *CreateClusterShrinkRequestTag) SetValue(v string) *CreateClusterShrinkRequestTag {
	s.Value = &v
	return s
}

func (s *CreateClusterShrinkRequestTag) Validate() error {
	return dara.Validate(s)
}
