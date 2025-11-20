// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHyperNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterName(v string) *ListHyperNodesRequest
	GetClusterName() *string
	SetHpnZone(v string) *ListHyperNodesRequest
	GetHpnZone() *string
	SetHyperNodeId(v string) *ListHyperNodesRequest
	GetHyperNodeId() *string
	SetMachineType(v string) *ListHyperNodesRequest
	GetMachineType() *string
	SetMaxResults(v int32) *ListHyperNodesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListHyperNodesRequest
	GetNextToken() *string
	SetNodeGroupName(v string) *ListHyperNodesRequest
	GetNodeGroupName() *string
	SetResourceGroupId(v string) *ListHyperNodesRequest
	GetResourceGroupId() *string
	SetTags(v []*ListHyperNodesRequestTags) *ListHyperNodesRequest
	GetTags() []*ListHyperNodesRequestTags
	SetZoneId(v string) *ListHyperNodesRequest
	GetZoneId() *string
}

type ListHyperNodesRequest struct {
	// example:
	//
	// Standard_Cluster
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// example:
	//
	// A1
	HpnZone *string `json:"HpnZone,omitempty" xml:"HpnZone,omitempty"`
	// example:
	//
	// e01-cn-zvp2tgykr08
	HyperNodeId *string `json:"HyperNodeId,omitempty" xml:"HyperNodeId,omitempty"`
	// example:
	//
	// efg1.nvga1
	MachineType *string `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 3a6b93229825ac667104463b56790c91
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// emr-default
	NodeGroupName *string `json:"NodeGroupName,omitempty" xml:"NodeGroupName,omitempty"`
	// example:
	//
	// rg-acfmwfm33rlt6zi
	ResourceGroupId *string                      `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tags            []*ListHyperNodesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListHyperNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHyperNodesRequest) GoString() string {
	return s.String()
}

func (s *ListHyperNodesRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *ListHyperNodesRequest) GetHpnZone() *string {
	return s.HpnZone
}

func (s *ListHyperNodesRequest) GetHyperNodeId() *string {
	return s.HyperNodeId
}

func (s *ListHyperNodesRequest) GetMachineType() *string {
	return s.MachineType
}

func (s *ListHyperNodesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListHyperNodesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListHyperNodesRequest) GetNodeGroupName() *string {
	return s.NodeGroupName
}

func (s *ListHyperNodesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListHyperNodesRequest) GetTags() []*ListHyperNodesRequestTags {
	return s.Tags
}

func (s *ListHyperNodesRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListHyperNodesRequest) SetClusterName(v string) *ListHyperNodesRequest {
	s.ClusterName = &v
	return s
}

func (s *ListHyperNodesRequest) SetHpnZone(v string) *ListHyperNodesRequest {
	s.HpnZone = &v
	return s
}

func (s *ListHyperNodesRequest) SetHyperNodeId(v string) *ListHyperNodesRequest {
	s.HyperNodeId = &v
	return s
}

func (s *ListHyperNodesRequest) SetMachineType(v string) *ListHyperNodesRequest {
	s.MachineType = &v
	return s
}

func (s *ListHyperNodesRequest) SetMaxResults(v int32) *ListHyperNodesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListHyperNodesRequest) SetNextToken(v string) *ListHyperNodesRequest {
	s.NextToken = &v
	return s
}

func (s *ListHyperNodesRequest) SetNodeGroupName(v string) *ListHyperNodesRequest {
	s.NodeGroupName = &v
	return s
}

func (s *ListHyperNodesRequest) SetResourceGroupId(v string) *ListHyperNodesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListHyperNodesRequest) SetTags(v []*ListHyperNodesRequestTags) *ListHyperNodesRequest {
	s.Tags = v
	return s
}

func (s *ListHyperNodesRequest) SetZoneId(v string) *ListHyperNodesRequest {
	s.ZoneId = &v
	return s
}

func (s *ListHyperNodesRequest) Validate() error {
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

type ListHyperNodesRequestTags struct {
	// example:
	//
	// alarm_xdc
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// 129
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListHyperNodesRequestTags) String() string {
	return dara.Prettify(s)
}

func (s ListHyperNodesRequestTags) GoString() string {
	return s.String()
}

func (s *ListHyperNodesRequestTags) GetKey() *string {
	return s.Key
}

func (s *ListHyperNodesRequestTags) GetValue() *string {
	return s.Value
}

func (s *ListHyperNodesRequestTags) SetKey(v string) *ListHyperNodesRequestTags {
	s.Key = &v
	return s
}

func (s *ListHyperNodesRequestTags) SetValue(v string) *ListHyperNodesRequestTags {
	s.Value = &v
	return s
}

func (s *ListHyperNodesRequestTags) Validate() error {
	return dara.Validate(s)
}
