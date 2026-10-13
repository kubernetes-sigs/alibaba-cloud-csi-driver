// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHyperNodesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterName(v string) *ListHyperNodesShrinkRequest
	GetClusterName() *string
	SetCommodityCode(v string) *ListHyperNodesShrinkRequest
	GetCommodityCode() *string
	SetHpnZone(v string) *ListHyperNodesShrinkRequest
	GetHpnZone() *string
	SetHyperNodeId(v string) *ListHyperNodesShrinkRequest
	GetHyperNodeId() *string
	SetHyperNodeIdsShrink(v string) *ListHyperNodesShrinkRequest
	GetHyperNodeIdsShrink() *string
	SetMachineType(v string) *ListHyperNodesShrinkRequest
	GetMachineType() *string
	SetMaxResults(v int32) *ListHyperNodesShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListHyperNodesShrinkRequest
	GetNextToken() *string
	SetNodeGroupName(v string) *ListHyperNodesShrinkRequest
	GetNodeGroupName() *string
	SetOperatingStatesShrink(v string) *ListHyperNodesShrinkRequest
	GetOperatingStatesShrink() *string
	SetResourceGroupId(v string) *ListHyperNodesShrinkRequest
	GetResourceGroupId() *string
	SetTags(v []*ListHyperNodesShrinkRequestTags) *ListHyperNodesShrinkRequest
	GetTags() []*ListHyperNodesShrinkRequestTags
	SetZoneId(v string) *ListHyperNodesShrinkRequest
	GetZoneId() *string
}

type ListHyperNodesShrinkRequest struct {
	// The name of the cluster.
	//
	// example:
	//
	// Standard_Cluster
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The commodity code.
	//
	// example:
	//
	// rds_machineinstanceba_public_cn
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The cluster number.
	//
	// example:
	//
	// A1
	HpnZone *string `json:"HpnZone,omitempty" xml:"HpnZone,omitempty"`
	// The ID of the node.
	//
	// example:
	//
	// e01-cn-zvp2tgykr08
	HyperNodeId *string `json:"HyperNodeId,omitempty" xml:"HyperNodeId,omitempty"`
	// The list of node IDs.
	HyperNodeIdsShrink *string `json:"HyperNodeIds,omitempty" xml:"HyperNodeIds,omitempty"`
	// The machine type.
	//
	// example:
	//
	// efg1.nvga1
	MachineType *string `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	// The number of entries to return on each page. The maximum value is 100.
	//
	// Default values:
	//
	// • If you do not specify this parameter or you specify a value that is less than 20, the default value is 20.
	//
	// • If you specify a value that is greater than 100, the default value is 100.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Set this parameter to the NextToken value that is returned from a previous call.
	//
	// example:
	//
	// 3a6b93229825ac667104463b56790c91
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The name of the node group.
	//
	// example:
	//
	// emr-default
	NodeGroupName *string `json:"NodeGroupName,omitempty" xml:"NodeGroupName,omitempty"`
	// The state of the node. If you do not specify this parameter, nodes in all states are returned. Valid values: Extending, UnusedNodeStopped, UnusedNodeStopping, Unused, Using, ReleaseLocking, Operating, Cutting, ClusterNodeStopped, UnusedNodeRecovering, ClusterNodeStopping, ClusterNodeRecovering, and Replacing.
	OperatingStatesShrink *string `json:"OperatingStates,omitempty" xml:"OperatingStates,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmwfm33rlt6zi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags.
	Tags []*ListHyperNodesShrinkRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the zone.
	//
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListHyperNodesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHyperNodesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListHyperNodesShrinkRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *ListHyperNodesShrinkRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *ListHyperNodesShrinkRequest) GetHpnZone() *string {
	return s.HpnZone
}

func (s *ListHyperNodesShrinkRequest) GetHyperNodeId() *string {
	return s.HyperNodeId
}

func (s *ListHyperNodesShrinkRequest) GetHyperNodeIdsShrink() *string {
	return s.HyperNodeIdsShrink
}

func (s *ListHyperNodesShrinkRequest) GetMachineType() *string {
	return s.MachineType
}

func (s *ListHyperNodesShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListHyperNodesShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListHyperNodesShrinkRequest) GetNodeGroupName() *string {
	return s.NodeGroupName
}

func (s *ListHyperNodesShrinkRequest) GetOperatingStatesShrink() *string {
	return s.OperatingStatesShrink
}

func (s *ListHyperNodesShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListHyperNodesShrinkRequest) GetTags() []*ListHyperNodesShrinkRequestTags {
	return s.Tags
}

func (s *ListHyperNodesShrinkRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListHyperNodesShrinkRequest) SetClusterName(v string) *ListHyperNodesShrinkRequest {
	s.ClusterName = &v
	return s
}

func (s *ListHyperNodesShrinkRequest) SetCommodityCode(v string) *ListHyperNodesShrinkRequest {
	s.CommodityCode = &v
	return s
}

func (s *ListHyperNodesShrinkRequest) SetHpnZone(v string) *ListHyperNodesShrinkRequest {
	s.HpnZone = &v
	return s
}

func (s *ListHyperNodesShrinkRequest) SetHyperNodeId(v string) *ListHyperNodesShrinkRequest {
	s.HyperNodeId = &v
	return s
}

func (s *ListHyperNodesShrinkRequest) SetHyperNodeIdsShrink(v string) *ListHyperNodesShrinkRequest {
	s.HyperNodeIdsShrink = &v
	return s
}

func (s *ListHyperNodesShrinkRequest) SetMachineType(v string) *ListHyperNodesShrinkRequest {
	s.MachineType = &v
	return s
}

func (s *ListHyperNodesShrinkRequest) SetMaxResults(v int32) *ListHyperNodesShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListHyperNodesShrinkRequest) SetNextToken(v string) *ListHyperNodesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListHyperNodesShrinkRequest) SetNodeGroupName(v string) *ListHyperNodesShrinkRequest {
	s.NodeGroupName = &v
	return s
}

func (s *ListHyperNodesShrinkRequest) SetOperatingStatesShrink(v string) *ListHyperNodesShrinkRequest {
	s.OperatingStatesShrink = &v
	return s
}

func (s *ListHyperNodesShrinkRequest) SetResourceGroupId(v string) *ListHyperNodesShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListHyperNodesShrinkRequest) SetTags(v []*ListHyperNodesShrinkRequestTags) *ListHyperNodesShrinkRequest {
	s.Tags = v
	return s
}

func (s *ListHyperNodesShrinkRequest) SetZoneId(v string) *ListHyperNodesShrinkRequest {
	s.ZoneId = &v
	return s
}

func (s *ListHyperNodesShrinkRequest) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListHyperNodesShrinkRequestTags struct {
	// The tag key.
	//
	// example:
	//
	// alarm_xdc
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// 129
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListHyperNodesShrinkRequestTags) String() string {
	return dara.Prettify(s)
}

func (s ListHyperNodesShrinkRequestTags) GoString() string {
	return s.String()
}

func (s *ListHyperNodesShrinkRequestTags) GetKey() *string {
	return s.Key
}

func (s *ListHyperNodesShrinkRequestTags) GetValue() *string {
	return s.Value
}

func (s *ListHyperNodesShrinkRequestTags) SetKey(v string) *ListHyperNodesShrinkRequestTags {
	s.Key = &v
	return s
}

func (s *ListHyperNodesShrinkRequestTags) SetValue(v string) *ListHyperNodesShrinkRequestTags {
	s.Value = &v
	return s
}

func (s *ListHyperNodesShrinkRequestTags) Validate() error {
	return dara.Validate(s)
}
