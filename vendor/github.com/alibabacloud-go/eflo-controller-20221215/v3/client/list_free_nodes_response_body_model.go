// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFreeNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListFreeNodesResponseBody
	GetNextToken() *string
	SetNodes(v []*ListFreeNodesResponseBodyNodes) *ListFreeNodesResponseBody
	GetNodes() []*ListFreeNodesResponseBodyNodes
	SetRequestId(v string) *ListFreeNodesResponseBody
	GetRequestId() *string
}

type ListFreeNodesResponseBody struct {
	// The returned pagination token which can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// a3f2224a5ec7224116c4f5246120abe4
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The nodes.
	Nodes []*ListFreeNodesResponseBodyNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// AA14CB86-70C4-5CB7-9E7B-6CCA77F3512B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFreeNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFreeNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ListFreeNodesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListFreeNodesResponseBody) GetNodes() []*ListFreeNodesResponseBodyNodes {
	return s.Nodes
}

func (s *ListFreeNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFreeNodesResponseBody) SetNextToken(v string) *ListFreeNodesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListFreeNodesResponseBody) SetNodes(v []*ListFreeNodesResponseBodyNodes) *ListFreeNodesResponseBody {
	s.Nodes = v
	return s
}

func (s *ListFreeNodesResponseBody) SetRequestId(v string) *ListFreeNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFreeNodesResponseBody) Validate() error {
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

type ListFreeNodesResponseBodyNodes struct {
	// The commodity code.
	//
	// example:
	//
	// bccluster_eflocomputing_public_cn
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 1652321554
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the node expires.
	//
	// example:
	//
	// 1673107200
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The cluster number.
	//
	// example:
	//
	// A1
	HpnZone *string `json:"HpnZone,omitempty" xml:"HpnZone,omitempty"`
	// example:
	//
	// e01-cn-zvp2tgykr08
	HyperNodeId *string `json:"HyperNodeId,omitempty" xml:"HyperNodeId,omitempty"`
	// The instance type.
	//
	// example:
	//
	// efg1.nvga1
	MachineType *string `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	// The node ID.
	//
	// example:
	//
	// e01-cn-7pp2x193801
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The node status.
	//
	// example:
	//
	// Unused
	OperatingState *string `json:"OperatingState,omitempty" xml:"OperatingState,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aekzkkbrpl4owgy
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The serial number of the node.
	//
	// example:
	//
	// sn_pozkHBgicd
	Sn *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
	// The tags.
	Tags []*ListFreeNodesResponseBodyNodesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-j
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListFreeNodesResponseBodyNodes) String() string {
	return dara.Prettify(s)
}

func (s ListFreeNodesResponseBodyNodes) GoString() string {
	return s.String()
}

func (s *ListFreeNodesResponseBodyNodes) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *ListFreeNodesResponseBodyNodes) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListFreeNodesResponseBodyNodes) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *ListFreeNodesResponseBodyNodes) GetHpnZone() *string {
	return s.HpnZone
}

func (s *ListFreeNodesResponseBodyNodes) GetHyperNodeId() *string {
	return s.HyperNodeId
}

func (s *ListFreeNodesResponseBodyNodes) GetMachineType() *string {
	return s.MachineType
}

func (s *ListFreeNodesResponseBodyNodes) GetNodeId() *string {
	return s.NodeId
}

func (s *ListFreeNodesResponseBodyNodes) GetOperatingState() *string {
	return s.OperatingState
}

func (s *ListFreeNodesResponseBodyNodes) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListFreeNodesResponseBodyNodes) GetSn() *string {
	return s.Sn
}

func (s *ListFreeNodesResponseBodyNodes) GetTags() []*ListFreeNodesResponseBodyNodesTags {
	return s.Tags
}

func (s *ListFreeNodesResponseBodyNodes) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListFreeNodesResponseBodyNodes) SetCommodityCode(v string) *ListFreeNodesResponseBodyNodes {
	s.CommodityCode = &v
	return s
}

func (s *ListFreeNodesResponseBodyNodes) SetCreateTime(v string) *ListFreeNodesResponseBodyNodes {
	s.CreateTime = &v
	return s
}

func (s *ListFreeNodesResponseBodyNodes) SetExpiredTime(v string) *ListFreeNodesResponseBodyNodes {
	s.ExpiredTime = &v
	return s
}

func (s *ListFreeNodesResponseBodyNodes) SetHpnZone(v string) *ListFreeNodesResponseBodyNodes {
	s.HpnZone = &v
	return s
}

func (s *ListFreeNodesResponseBodyNodes) SetHyperNodeId(v string) *ListFreeNodesResponseBodyNodes {
	s.HyperNodeId = &v
	return s
}

func (s *ListFreeNodesResponseBodyNodes) SetMachineType(v string) *ListFreeNodesResponseBodyNodes {
	s.MachineType = &v
	return s
}

func (s *ListFreeNodesResponseBodyNodes) SetNodeId(v string) *ListFreeNodesResponseBodyNodes {
	s.NodeId = &v
	return s
}

func (s *ListFreeNodesResponseBodyNodes) SetOperatingState(v string) *ListFreeNodesResponseBodyNodes {
	s.OperatingState = &v
	return s
}

func (s *ListFreeNodesResponseBodyNodes) SetResourceGroupId(v string) *ListFreeNodesResponseBodyNodes {
	s.ResourceGroupId = &v
	return s
}

func (s *ListFreeNodesResponseBodyNodes) SetSn(v string) *ListFreeNodesResponseBodyNodes {
	s.Sn = &v
	return s
}

func (s *ListFreeNodesResponseBodyNodes) SetTags(v []*ListFreeNodesResponseBodyNodesTags) *ListFreeNodesResponseBodyNodes {
	s.Tags = v
	return s
}

func (s *ListFreeNodesResponseBodyNodes) SetZoneId(v string) *ListFreeNodesResponseBodyNodes {
	s.ZoneId = &v
	return s
}

func (s *ListFreeNodesResponseBodyNodes) Validate() error {
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

type ListFreeNodesResponseBodyNodesTags struct {
	// The tag key.
	//
	// example:
	//
	// aa_key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// aa_vakye
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListFreeNodesResponseBodyNodesTags) String() string {
	return dara.Prettify(s)
}

func (s ListFreeNodesResponseBodyNodesTags) GoString() string {
	return s.String()
}

func (s *ListFreeNodesResponseBodyNodesTags) GetKey() *string {
	return s.Key
}

func (s *ListFreeNodesResponseBodyNodesTags) GetValue() *string {
	return s.Value
}

func (s *ListFreeNodesResponseBodyNodesTags) SetKey(v string) *ListFreeNodesResponseBodyNodesTags {
	s.Key = &v
	return s
}

func (s *ListFreeNodesResponseBodyNodesTags) SetValue(v string) *ListFreeNodesResponseBodyNodesTags {
	s.Value = &v
	return s
}

func (s *ListFreeNodesResponseBodyNodesTags) Validate() error {
	return dara.Validate(s)
}
