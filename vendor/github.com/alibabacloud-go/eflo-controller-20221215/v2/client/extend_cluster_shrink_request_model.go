// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExtendClusterShrinkRequest interface {
  dara.Model
  String() string
  GoString() string
  SetClusterId(v string) *ExtendClusterShrinkRequest
  GetClusterId() *string 
  SetIgnoreFailedNodeTasks(v bool) *ExtendClusterShrinkRequest
  GetIgnoreFailedNodeTasks() *bool 
  SetIpAllocationPolicyShrink(v string) *ExtendClusterShrinkRequest
  GetIpAllocationPolicyShrink() *string 
  SetNodeGroupsShrink(v string) *ExtendClusterShrinkRequest
  GetNodeGroupsShrink() *string 
  SetVSwitchZoneId(v string) *ExtendClusterShrinkRequest
  GetVSwitchZoneId() *string 
  SetVpdSubnetsShrink(v string) *ExtendClusterShrinkRequest
  GetVpdSubnetsShrink() *string 
}

type ExtendClusterShrinkRequest struct {
  // Cluster ID
  // 
  // example:
  // 
  // i15b480fbd2fcdbc2869cd80
  ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
  // Whether to allow skipping failed node tasks, default value is False
  // 
  // example:
  // 
  // False
  IgnoreFailedNodeTasks *bool `json:"IgnoreFailedNodeTasks,omitempty" xml:"IgnoreFailedNodeTasks,omitempty"`
  // IP allocation combination policy: Each policy can only choose one type, and multiple policies can be combined
  IpAllocationPolicyShrink *string `json:"IpAllocationPolicy,omitempty" xml:"IpAllocationPolicy,omitempty"`
  // Node Groups
  NodeGroupsShrink *string `json:"NodeGroups,omitempty" xml:"NodeGroups,omitempty"`
  // VSwitch availability zone ID
  // 
  // example:
  // 
  // cn-shanghai-b
  VSwitchZoneId *string `json:"VSwitchZoneId,omitempty" xml:"VSwitchZoneId,omitempty"`
  // List of cluster subnets
  VpdSubnetsShrink *string `json:"VpdSubnets,omitempty" xml:"VpdSubnets,omitempty"`
}

func (s ExtendClusterShrinkRequest) String() string {
  return dara.Prettify(s)
}

func (s ExtendClusterShrinkRequest) GoString() string {
  return s.String()
}

func (s *ExtendClusterShrinkRequest) GetClusterId() *string  {
  return s.ClusterId
}

func (s *ExtendClusterShrinkRequest) GetIgnoreFailedNodeTasks() *bool  {
  return s.IgnoreFailedNodeTasks
}

func (s *ExtendClusterShrinkRequest) GetIpAllocationPolicyShrink() *string  {
  return s.IpAllocationPolicyShrink
}

func (s *ExtendClusterShrinkRequest) GetNodeGroupsShrink() *string  {
  return s.NodeGroupsShrink
}

func (s *ExtendClusterShrinkRequest) GetVSwitchZoneId() *string  {
  return s.VSwitchZoneId
}

func (s *ExtendClusterShrinkRequest) GetVpdSubnetsShrink() *string  {
  return s.VpdSubnetsShrink
}

func (s *ExtendClusterShrinkRequest) SetClusterId(v string) *ExtendClusterShrinkRequest {
  s.ClusterId = &v
  return s
}

func (s *ExtendClusterShrinkRequest) SetIgnoreFailedNodeTasks(v bool) *ExtendClusterShrinkRequest {
  s.IgnoreFailedNodeTasks = &v
  return s
}

func (s *ExtendClusterShrinkRequest) SetIpAllocationPolicyShrink(v string) *ExtendClusterShrinkRequest {
  s.IpAllocationPolicyShrink = &v
  return s
}

func (s *ExtendClusterShrinkRequest) SetNodeGroupsShrink(v string) *ExtendClusterShrinkRequest {
  s.NodeGroupsShrink = &v
  return s
}

func (s *ExtendClusterShrinkRequest) SetVSwitchZoneId(v string) *ExtendClusterShrinkRequest {
  s.VSwitchZoneId = &v
  return s
}

func (s *ExtendClusterShrinkRequest) SetVpdSubnetsShrink(v string) *ExtendClusterShrinkRequest {
  s.VpdSubnetsShrink = &v
  return s
}

func (s *ExtendClusterShrinkRequest) Validate() error {
  return dara.Validate(s)
}

