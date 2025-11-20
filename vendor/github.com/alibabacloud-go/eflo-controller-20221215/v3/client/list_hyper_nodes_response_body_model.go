// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHyperNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHyperNodes(v []*ListHyperNodesResponseBodyHyperNodes) *ListHyperNodesResponseBody
	GetHyperNodes() []*ListHyperNodesResponseBodyHyperNodes
	SetMaxResults(v int32) *ListHyperNodesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListHyperNodesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListHyperNodesResponseBody
	GetRequestId() *string
}

type ListHyperNodesResponseBody struct {
	HyperNodes []*ListHyperNodesResponseBodyHyperNodes `json:"HyperNodes,omitempty" xml:"HyperNodes,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 7ed93fda-5e7f-436a-ae5a-bd8e6b04e36b
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// E64F9128-E2FC-5998-B769-199B0CB18138
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListHyperNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHyperNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ListHyperNodesResponseBody) GetHyperNodes() []*ListHyperNodesResponseBodyHyperNodes {
	return s.HyperNodes
}

func (s *ListHyperNodesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListHyperNodesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListHyperNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHyperNodesResponseBody) SetHyperNodes(v []*ListHyperNodesResponseBodyHyperNodes) *ListHyperNodesResponseBody {
	s.HyperNodes = v
	return s
}

func (s *ListHyperNodesResponseBody) SetMaxResults(v int32) *ListHyperNodesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListHyperNodesResponseBody) SetNextToken(v string) *ListHyperNodesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListHyperNodesResponseBody) SetRequestId(v string) *ListHyperNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHyperNodesResponseBody) Validate() error {
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

type ListHyperNodesResponseBodyHyperNodes struct {
	// example:
	//
	// i115226661755786900341
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// test-ack
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// example:
	//
	// bccluster_eflocomputing_public_cn
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// example:
	//
	// 2025-07-09T10:41:56.577
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2025-09-28T16:00:00Z
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
	// e01-cn-2r42tmj4z02
	HyperNodeId *string `json:"HyperNodeId,omitempty" xml:"HyperNodeId,omitempty"`
	// example:
	//
	// efg2.GN9C.cn8
	MachineType *string `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	// example:
	//
	// i121135081698451727812
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// example:
	//
	// l20c-0801
	NodeGroupName *string `json:"NodeGroupName,omitempty" xml:"NodeGroupName,omitempty"`
	// example:
	//
	// Using
	OperatingState *string `json:"OperatingState,omitempty" xml:"OperatingState,omitempty"`
	// example:
	//
	// rg-acfmwfm33rlt6zi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Deprecated
	//
	// example:
	//
	// Extending
	Status *string                                     `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags   []*ListHyperNodesResponseBodyHyperNodesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// example:
	//
	// i153333771756952392398
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListHyperNodesResponseBodyHyperNodes) String() string {
	return dara.Prettify(s)
}

func (s ListHyperNodesResponseBodyHyperNodes) GoString() string {
	return s.String()
}

func (s *ListHyperNodesResponseBodyHyperNodes) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListHyperNodesResponseBodyHyperNodes) GetClusterName() *string {
	return s.ClusterName
}

func (s *ListHyperNodesResponseBodyHyperNodes) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *ListHyperNodesResponseBodyHyperNodes) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListHyperNodesResponseBodyHyperNodes) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *ListHyperNodesResponseBodyHyperNodes) GetHostname() *string {
	return s.Hostname
}

func (s *ListHyperNodesResponseBodyHyperNodes) GetHpnZone() *string {
	return s.HpnZone
}

func (s *ListHyperNodesResponseBodyHyperNodes) GetHyperNodeId() *string {
	return s.HyperNodeId
}

func (s *ListHyperNodesResponseBodyHyperNodes) GetMachineType() *string {
	return s.MachineType
}

func (s *ListHyperNodesResponseBodyHyperNodes) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *ListHyperNodesResponseBodyHyperNodes) GetNodeGroupName() *string {
	return s.NodeGroupName
}

func (s *ListHyperNodesResponseBodyHyperNodes) GetOperatingState() *string {
	return s.OperatingState
}

func (s *ListHyperNodesResponseBodyHyperNodes) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListHyperNodesResponseBodyHyperNodes) GetStatus() *string {
	return s.Status
}

func (s *ListHyperNodesResponseBodyHyperNodes) GetTags() []*ListHyperNodesResponseBodyHyperNodesTags {
	return s.Tags
}

func (s *ListHyperNodesResponseBodyHyperNodes) GetTaskId() *string {
	return s.TaskId
}

func (s *ListHyperNodesResponseBodyHyperNodes) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListHyperNodesResponseBodyHyperNodes) SetClusterId(v string) *ListHyperNodesResponseBodyHyperNodes {
	s.ClusterId = &v
	return s
}

func (s *ListHyperNodesResponseBodyHyperNodes) SetClusterName(v string) *ListHyperNodesResponseBodyHyperNodes {
	s.ClusterName = &v
	return s
}

func (s *ListHyperNodesResponseBodyHyperNodes) SetCommodityCode(v string) *ListHyperNodesResponseBodyHyperNodes {
	s.CommodityCode = &v
	return s
}

func (s *ListHyperNodesResponseBodyHyperNodes) SetCreateTime(v string) *ListHyperNodesResponseBodyHyperNodes {
	s.CreateTime = &v
	return s
}

func (s *ListHyperNodesResponseBodyHyperNodes) SetExpireTime(v string) *ListHyperNodesResponseBodyHyperNodes {
	s.ExpireTime = &v
	return s
}

func (s *ListHyperNodesResponseBodyHyperNodes) SetHostname(v string) *ListHyperNodesResponseBodyHyperNodes {
	s.Hostname = &v
	return s
}

func (s *ListHyperNodesResponseBodyHyperNodes) SetHpnZone(v string) *ListHyperNodesResponseBodyHyperNodes {
	s.HpnZone = &v
	return s
}

func (s *ListHyperNodesResponseBodyHyperNodes) SetHyperNodeId(v string) *ListHyperNodesResponseBodyHyperNodes {
	s.HyperNodeId = &v
	return s
}

func (s *ListHyperNodesResponseBodyHyperNodes) SetMachineType(v string) *ListHyperNodesResponseBodyHyperNodes {
	s.MachineType = &v
	return s
}

func (s *ListHyperNodesResponseBodyHyperNodes) SetNodeGroupId(v string) *ListHyperNodesResponseBodyHyperNodes {
	s.NodeGroupId = &v
	return s
}

func (s *ListHyperNodesResponseBodyHyperNodes) SetNodeGroupName(v string) *ListHyperNodesResponseBodyHyperNodes {
	s.NodeGroupName = &v
	return s
}

func (s *ListHyperNodesResponseBodyHyperNodes) SetOperatingState(v string) *ListHyperNodesResponseBodyHyperNodes {
	s.OperatingState = &v
	return s
}

func (s *ListHyperNodesResponseBodyHyperNodes) SetResourceGroupId(v string) *ListHyperNodesResponseBodyHyperNodes {
	s.ResourceGroupId = &v
	return s
}

func (s *ListHyperNodesResponseBodyHyperNodes) SetStatus(v string) *ListHyperNodesResponseBodyHyperNodes {
	s.Status = &v
	return s
}

func (s *ListHyperNodesResponseBodyHyperNodes) SetTags(v []*ListHyperNodesResponseBodyHyperNodesTags) *ListHyperNodesResponseBodyHyperNodes {
	s.Tags = v
	return s
}

func (s *ListHyperNodesResponseBodyHyperNodes) SetTaskId(v string) *ListHyperNodesResponseBodyHyperNodes {
	s.TaskId = &v
	return s
}

func (s *ListHyperNodesResponseBodyHyperNodes) SetZoneId(v string) *ListHyperNodesResponseBodyHyperNodes {
	s.ZoneId = &v
	return s
}

func (s *ListHyperNodesResponseBodyHyperNodes) Validate() error {
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

type ListHyperNodesResponseBodyHyperNodesTags struct {
	// example:
	//
	// alarm_xdc
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// 97
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListHyperNodesResponseBodyHyperNodesTags) String() string {
	return dara.Prettify(s)
}

func (s ListHyperNodesResponseBodyHyperNodesTags) GoString() string {
	return s.String()
}

func (s *ListHyperNodesResponseBodyHyperNodesTags) GetKey() *string {
	return s.Key
}

func (s *ListHyperNodesResponseBodyHyperNodesTags) GetValue() *string {
	return s.Value
}

func (s *ListHyperNodesResponseBodyHyperNodesTags) SetKey(v string) *ListHyperNodesResponseBodyHyperNodesTags {
	s.Key = &v
	return s
}

func (s *ListHyperNodesResponseBodyHyperNodesTags) SetValue(v string) *ListHyperNodesResponseBodyHyperNodesTags {
	s.Value = &v
	return s
}

func (s *ListHyperNodesResponseBodyHyperNodesTags) Validate() error {
	return dara.Validate(s)
}
