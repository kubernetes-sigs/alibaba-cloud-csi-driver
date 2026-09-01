// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFreeNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHpnZone(v string) *ListFreeNodesRequest
	GetHpnZone() *string
	SetMachineType(v string) *ListFreeNodesRequest
	GetMachineType() *string
	SetMaxResults(v int64) *ListFreeNodesRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListFreeNodesRequest
	GetNextToken() *string
	SetOperatingStates(v []*string) *ListFreeNodesRequest
	GetOperatingStates() []*string
	SetResourceGroupId(v string) *ListFreeNodesRequest
	GetResourceGroupId() *string
	SetTags(v []*ListFreeNodesRequestTags) *ListFreeNodesRequest
	GetTags() []*ListFreeNodesRequestTags
}

type ListFreeNodesRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// A1
	HpnZone *string `json:"HpnZone,omitempty" xml:"HpnZone,omitempty"`
	// The machine type.
	//
	// example:
	//
	// mock-machine-type2
	MachineType *string `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	// The number of entries to return on each page for a paged query. The default value is 20.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The query token. Set it to the NextToken value returned from the previous API call.
	//
	// example:
	//
	// a3f2224a5ec7224116c4f5246120abe4
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The type of unused nodes to be returned.
	OperatingStates []*string `json:"OperatingStates,omitempty" xml:"OperatingStates,omitempty" type:"Repeated"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmxno4vh5muoq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tag information.
	Tags []*ListFreeNodesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListFreeNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFreeNodesRequest) GoString() string {
	return s.String()
}

func (s *ListFreeNodesRequest) GetHpnZone() *string {
	return s.HpnZone
}

func (s *ListFreeNodesRequest) GetMachineType() *string {
	return s.MachineType
}

func (s *ListFreeNodesRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListFreeNodesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListFreeNodesRequest) GetOperatingStates() []*string {
	return s.OperatingStates
}

func (s *ListFreeNodesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListFreeNodesRequest) GetTags() []*ListFreeNodesRequestTags {
	return s.Tags
}

func (s *ListFreeNodesRequest) SetHpnZone(v string) *ListFreeNodesRequest {
	s.HpnZone = &v
	return s
}

func (s *ListFreeNodesRequest) SetMachineType(v string) *ListFreeNodesRequest {
	s.MachineType = &v
	return s
}

func (s *ListFreeNodesRequest) SetMaxResults(v int64) *ListFreeNodesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListFreeNodesRequest) SetNextToken(v string) *ListFreeNodesRequest {
	s.NextToken = &v
	return s
}

func (s *ListFreeNodesRequest) SetOperatingStates(v []*string) *ListFreeNodesRequest {
	s.OperatingStates = v
	return s
}

func (s *ListFreeNodesRequest) SetResourceGroupId(v string) *ListFreeNodesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListFreeNodesRequest) SetTags(v []*ListFreeNodesRequestTags) *ListFreeNodesRequest {
	s.Tags = v
	return s
}

func (s *ListFreeNodesRequest) Validate() error {
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

type ListFreeNodesRequestTags struct {
	// The tag key.
	//
	// example:
	//
	// my_key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// 129
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListFreeNodesRequestTags) String() string {
	return dara.Prettify(s)
}

func (s ListFreeNodesRequestTags) GoString() string {
	return s.String()
}

func (s *ListFreeNodesRequestTags) GetKey() *string {
	return s.Key
}

func (s *ListFreeNodesRequestTags) GetValue() *string {
	return s.Value
}

func (s *ListFreeNodesRequestTags) SetKey(v string) *ListFreeNodesRequestTags {
	s.Key = &v
	return s
}

func (s *ListFreeNodesRequestTags) SetValue(v string) *ListFreeNodesRequestTags {
	s.Value = &v
	return s
}

func (s *ListFreeNodesRequestTags) Validate() error {
	return dara.Validate(s)
}
