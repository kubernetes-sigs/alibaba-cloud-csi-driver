// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFreeHyperNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHpnZone(v string) *ListFreeHyperNodesRequest
	GetHpnZone() *string
	SetMachineType(v string) *ListFreeHyperNodesRequest
	GetMachineType() *string
	SetMaxResults(v int32) *ListFreeHyperNodesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListFreeHyperNodesRequest
	GetNextToken() *string
	SetResourceGroupId(v string) *ListFreeHyperNodesRequest
	GetResourceGroupId() *string
	SetStatus(v []*string) *ListFreeHyperNodesRequest
	GetStatus() []*string
	SetTags(v []*ListFreeHyperNodesRequestTags) *ListFreeHyperNodesRequest
	GetTags() []*ListFreeHyperNodesRequestTags
}

type ListFreeHyperNodesRequest struct {
	// example:
	//
	// A1
	HpnZone *string `json:"HpnZone,omitempty" xml:"HpnZone,omitempty"`
	// example:
	//
	// efg1.nvga1
	MachineType *string `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// a3f2224a5ec7224116c4f5246120****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// rg-aekzyqdwnfabx6q
	ResourceGroupId *string                          `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Status          []*string                        `json:"Status,omitempty" xml:"Status,omitempty" type:"Repeated"`
	Tags            []*ListFreeHyperNodesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListFreeHyperNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFreeHyperNodesRequest) GoString() string {
	return s.String()
}

func (s *ListFreeHyperNodesRequest) GetHpnZone() *string {
	return s.HpnZone
}

func (s *ListFreeHyperNodesRequest) GetMachineType() *string {
	return s.MachineType
}

func (s *ListFreeHyperNodesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListFreeHyperNodesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListFreeHyperNodesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListFreeHyperNodesRequest) GetStatus() []*string {
	return s.Status
}

func (s *ListFreeHyperNodesRequest) GetTags() []*ListFreeHyperNodesRequestTags {
	return s.Tags
}

func (s *ListFreeHyperNodesRequest) SetHpnZone(v string) *ListFreeHyperNodesRequest {
	s.HpnZone = &v
	return s
}

func (s *ListFreeHyperNodesRequest) SetMachineType(v string) *ListFreeHyperNodesRequest {
	s.MachineType = &v
	return s
}

func (s *ListFreeHyperNodesRequest) SetMaxResults(v int32) *ListFreeHyperNodesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListFreeHyperNodesRequest) SetNextToken(v string) *ListFreeHyperNodesRequest {
	s.NextToken = &v
	return s
}

func (s *ListFreeHyperNodesRequest) SetResourceGroupId(v string) *ListFreeHyperNodesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListFreeHyperNodesRequest) SetStatus(v []*string) *ListFreeHyperNodesRequest {
	s.Status = v
	return s
}

func (s *ListFreeHyperNodesRequest) SetTags(v []*ListFreeHyperNodesRequestTags) *ListFreeHyperNodesRequest {
	s.Tags = v
	return s
}

func (s *ListFreeHyperNodesRequest) Validate() error {
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

type ListFreeHyperNodesRequestTags struct {
	// example:
	//
	// my_key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// 129
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListFreeHyperNodesRequestTags) String() string {
	return dara.Prettify(s)
}

func (s ListFreeHyperNodesRequestTags) GoString() string {
	return s.String()
}

func (s *ListFreeHyperNodesRequestTags) GetKey() *string {
	return s.Key
}

func (s *ListFreeHyperNodesRequestTags) GetValue() *string {
	return s.Value
}

func (s *ListFreeHyperNodesRequestTags) SetKey(v string) *ListFreeHyperNodesRequestTags {
	s.Key = &v
	return s
}

func (s *ListFreeHyperNodesRequestTags) SetValue(v string) *ListFreeHyperNodesRequestTags {
	s.Value = &v
	return s
}

func (s *ListFreeHyperNodesRequestTags) Validate() error {
	return dara.Validate(s)
}
