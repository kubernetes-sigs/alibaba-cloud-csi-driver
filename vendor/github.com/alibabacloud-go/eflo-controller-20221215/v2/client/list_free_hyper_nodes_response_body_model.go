// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFreeHyperNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHyperNodes(v []*ListFreeHyperNodesResponseBodyHyperNodes) *ListFreeHyperNodesResponseBody
	GetHyperNodes() []*ListFreeHyperNodesResponseBodyHyperNodes
	SetMaxResults(v int32) *ListFreeHyperNodesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListFreeHyperNodesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListFreeHyperNodesResponseBody
	GetRequestId() *string
}

type ListFreeHyperNodesResponseBody struct {
	HyperNodes []*ListFreeHyperNodesResponseBodyHyperNodes `json:"HyperNodes,omitempty" xml:"HyperNodes,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 563d42ae0b17572449ec8c97f7f66069
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// A511C02A-0127-51AA-A9F9-966382C9A1B5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFreeHyperNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFreeHyperNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ListFreeHyperNodesResponseBody) GetHyperNodes() []*ListFreeHyperNodesResponseBodyHyperNodes {
	return s.HyperNodes
}

func (s *ListFreeHyperNodesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListFreeHyperNodesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListFreeHyperNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFreeHyperNodesResponseBody) SetHyperNodes(v []*ListFreeHyperNodesResponseBodyHyperNodes) *ListFreeHyperNodesResponseBody {
	s.HyperNodes = v
	return s
}

func (s *ListFreeHyperNodesResponseBody) SetMaxResults(v int32) *ListFreeHyperNodesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListFreeHyperNodesResponseBody) SetNextToken(v string) *ListFreeHyperNodesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListFreeHyperNodesResponseBody) SetRequestId(v string) *ListFreeHyperNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFreeHyperNodesResponseBody) Validate() error {
	if s.HyperNodes != nil {
		for _, item := range s.HyperNodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListFreeHyperNodesResponseBodyHyperNodes struct {
	// example:
	//
	// bccluster_eflocomputing_public_cn
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// example:
	//
	// 2025-08-06T10:11:41.569
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2025-01-22T23:59:59Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// alywlcb-lingjun-gpu-0025
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// example:
	//
	// A1
	HpnZone *string `json:"HpnZone,omitempty" xml:"HpnZone,omitempty"`
	// example:
	//
	// e01-cn-7pp2x193801
	HyperNodeId *string `json:"HyperNodeId,omitempty" xml:"HyperNodeId,omitempty"`
	// example:
	//
	// efg2.ks01L20Z2
	MachineType *string `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	// example:
	//
	// rg-acfmwaateahzoii
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// Using
	Status *string                                         `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags   []*ListFreeHyperNodesResponseBodyHyperNodesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// example:
	//
	// cn-wulanchabu-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListFreeHyperNodesResponseBodyHyperNodes) String() string {
	return dara.Prettify(s)
}

func (s ListFreeHyperNodesResponseBodyHyperNodes) GoString() string {
	return s.String()
}

func (s *ListFreeHyperNodesResponseBodyHyperNodes) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *ListFreeHyperNodesResponseBodyHyperNodes) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListFreeHyperNodesResponseBodyHyperNodes) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *ListFreeHyperNodesResponseBodyHyperNodes) GetHostname() *string {
	return s.Hostname
}

func (s *ListFreeHyperNodesResponseBodyHyperNodes) GetHpnZone() *string {
	return s.HpnZone
}

func (s *ListFreeHyperNodesResponseBodyHyperNodes) GetHyperNodeId() *string {
	return s.HyperNodeId
}

func (s *ListFreeHyperNodesResponseBodyHyperNodes) GetMachineType() *string {
	return s.MachineType
}

func (s *ListFreeHyperNodesResponseBodyHyperNodes) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListFreeHyperNodesResponseBodyHyperNodes) GetStatus() *string {
	return s.Status
}

func (s *ListFreeHyperNodesResponseBodyHyperNodes) GetTags() []*ListFreeHyperNodesResponseBodyHyperNodesTags {
	return s.Tags
}

func (s *ListFreeHyperNodesResponseBodyHyperNodes) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListFreeHyperNodesResponseBodyHyperNodes) SetCommodityCode(v string) *ListFreeHyperNodesResponseBodyHyperNodes {
	s.CommodityCode = &v
	return s
}

func (s *ListFreeHyperNodesResponseBodyHyperNodes) SetCreateTime(v string) *ListFreeHyperNodesResponseBodyHyperNodes {
	s.CreateTime = &v
	return s
}

func (s *ListFreeHyperNodesResponseBodyHyperNodes) SetExpireTime(v string) *ListFreeHyperNodesResponseBodyHyperNodes {
	s.ExpireTime = &v
	return s
}

func (s *ListFreeHyperNodesResponseBodyHyperNodes) SetHostname(v string) *ListFreeHyperNodesResponseBodyHyperNodes {
	s.Hostname = &v
	return s
}

func (s *ListFreeHyperNodesResponseBodyHyperNodes) SetHpnZone(v string) *ListFreeHyperNodesResponseBodyHyperNodes {
	s.HpnZone = &v
	return s
}

func (s *ListFreeHyperNodesResponseBodyHyperNodes) SetHyperNodeId(v string) *ListFreeHyperNodesResponseBodyHyperNodes {
	s.HyperNodeId = &v
	return s
}

func (s *ListFreeHyperNodesResponseBodyHyperNodes) SetMachineType(v string) *ListFreeHyperNodesResponseBodyHyperNodes {
	s.MachineType = &v
	return s
}

func (s *ListFreeHyperNodesResponseBodyHyperNodes) SetResourceGroupId(v string) *ListFreeHyperNodesResponseBodyHyperNodes {
	s.ResourceGroupId = &v
	return s
}

func (s *ListFreeHyperNodesResponseBodyHyperNodes) SetStatus(v string) *ListFreeHyperNodesResponseBodyHyperNodes {
	s.Status = &v
	return s
}

func (s *ListFreeHyperNodesResponseBodyHyperNodes) SetTags(v []*ListFreeHyperNodesResponseBodyHyperNodesTags) *ListFreeHyperNodesResponseBodyHyperNodes {
	s.Tags = v
	return s
}

func (s *ListFreeHyperNodesResponseBodyHyperNodes) SetZoneId(v string) *ListFreeHyperNodesResponseBodyHyperNodes {
	s.ZoneId = &v
	return s
}

func (s *ListFreeHyperNodesResponseBodyHyperNodes) Validate() error {
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

type ListFreeHyperNodesResponseBodyHyperNodesTags struct {
	// example:
	//
	// Cpu_Usage
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// on
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListFreeHyperNodesResponseBodyHyperNodesTags) String() string {
	return dara.Prettify(s)
}

func (s ListFreeHyperNodesResponseBodyHyperNodesTags) GoString() string {
	return s.String()
}

func (s *ListFreeHyperNodesResponseBodyHyperNodesTags) GetKey() *string {
	return s.Key
}

func (s *ListFreeHyperNodesResponseBodyHyperNodesTags) GetValue() *string {
	return s.Value
}

func (s *ListFreeHyperNodesResponseBodyHyperNodesTags) SetKey(v string) *ListFreeHyperNodesResponseBodyHyperNodesTags {
	s.Key = &v
	return s
}

func (s *ListFreeHyperNodesResponseBodyHyperNodesTags) SetValue(v string) *ListFreeHyperNodesResponseBodyHyperNodesTags {
	s.Value = &v
	return s
}

func (s *ListFreeHyperNodesResponseBodyHyperNodesTags) Validate() error {
	return dara.Validate(s)
}
