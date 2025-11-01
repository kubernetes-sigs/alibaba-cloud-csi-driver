// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterHyperNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHyperNodes(v []*ListClusterHyperNodesResponseBodyHyperNodes) *ListClusterHyperNodesResponseBody
	GetHyperNodes() []*ListClusterHyperNodesResponseBodyHyperNodes
	SetNextToken(v string) *ListClusterHyperNodesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListClusterHyperNodesResponseBody
	GetRequestId() *string
}

type ListClusterHyperNodesResponseBody struct {
	HyperNodes []*ListClusterHyperNodesResponseBodyHyperNodes `json:"HyperNodes,omitempty" xml:"HyperNodes,omitempty" type:"Repeated"`
	// example:
	//
	// 563d42ae0b17572449ec8c97f7f66069
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 4FD06DF0-9167-5C6F-A145-F30CA4A15D54
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListClusterHyperNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListClusterHyperNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterHyperNodesResponseBody) GetHyperNodes() []*ListClusterHyperNodesResponseBodyHyperNodes {
	return s.HyperNodes
}

func (s *ListClusterHyperNodesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListClusterHyperNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListClusterHyperNodesResponseBody) SetHyperNodes(v []*ListClusterHyperNodesResponseBodyHyperNodes) *ListClusterHyperNodesResponseBody {
	s.HyperNodes = v
	return s
}

func (s *ListClusterHyperNodesResponseBody) SetNextToken(v string) *ListClusterHyperNodesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListClusterHyperNodesResponseBody) SetRequestId(v string) *ListClusterHyperNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClusterHyperNodesResponseBody) Validate() error {
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

type ListClusterHyperNodesResponseBodyHyperNodes struct {
	// example:
	//
	// bccluster_eflocomputing_public_cn
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// example:
	//
	// 2025-07-07T17:38:35.391
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2025-04-19T02:32:48Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// False
	FileSystemMountEnabled *bool `json:"FileSystemMountEnabled,omitempty" xml:"FileSystemMountEnabled,omitempty"`
	// example:
	//
	// alywlcb-lingjun-gpu-0025
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// example:
	//
	// B1
	HpnZone *string `json:"HpnZone,omitempty" xml:"HpnZone,omitempty"`
	// example:
	//
	// e01-cn-2r42tmj4z02
	HyperNodeId *string `json:"HyperNodeId,omitempty" xml:"HyperNodeId,omitempty"`
	// example:
	//
	// efg2.NH2cn
	MachineType *string `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	// example:
	//
	// i123229811742436895560
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// example:
	//
	// g1
	NodeGroupName *string `json:"NodeGroupName,omitempty" xml:"NodeGroupName,omitempty"`
	// example:
	//
	// Extending
	Status *string                                            `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags   []*ListClusterHyperNodesResponseBodyHyperNodesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// example:
	//
	// i153907661745288876128
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// vsw-bp1u4ej4ap8c4yiqfi87c
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// vpc-0jl8gs7qmx89739e210dn
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// cn-shanghai-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListClusterHyperNodesResponseBodyHyperNodes) String() string {
	return dara.Prettify(s)
}

func (s ListClusterHyperNodesResponseBodyHyperNodes) GoString() string {
	return s.String()
}

func (s *ListClusterHyperNodesResponseBodyHyperNodes) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *ListClusterHyperNodesResponseBodyHyperNodes) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListClusterHyperNodesResponseBodyHyperNodes) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *ListClusterHyperNodesResponseBodyHyperNodes) GetFileSystemMountEnabled() *bool {
	return s.FileSystemMountEnabled
}

func (s *ListClusterHyperNodesResponseBodyHyperNodes) GetHostname() *string {
	return s.Hostname
}

func (s *ListClusterHyperNodesResponseBodyHyperNodes) GetHpnZone() *string {
	return s.HpnZone
}

func (s *ListClusterHyperNodesResponseBodyHyperNodes) GetHyperNodeId() *string {
	return s.HyperNodeId
}

func (s *ListClusterHyperNodesResponseBodyHyperNodes) GetMachineType() *string {
	return s.MachineType
}

func (s *ListClusterHyperNodesResponseBodyHyperNodes) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *ListClusterHyperNodesResponseBodyHyperNodes) GetNodeGroupName() *string {
	return s.NodeGroupName
}

func (s *ListClusterHyperNodesResponseBodyHyperNodes) GetStatus() *string {
	return s.Status
}

func (s *ListClusterHyperNodesResponseBodyHyperNodes) GetTags() []*ListClusterHyperNodesResponseBodyHyperNodesTags {
	return s.Tags
}

func (s *ListClusterHyperNodesResponseBodyHyperNodes) GetTaskId() *string {
	return s.TaskId
}

func (s *ListClusterHyperNodesResponseBodyHyperNodes) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ListClusterHyperNodesResponseBodyHyperNodes) GetVpcId() *string {
	return s.VpcId
}

func (s *ListClusterHyperNodesResponseBodyHyperNodes) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListClusterHyperNodesResponseBodyHyperNodes) SetCommodityCode(v string) *ListClusterHyperNodesResponseBodyHyperNodes {
	s.CommodityCode = &v
	return s
}

func (s *ListClusterHyperNodesResponseBodyHyperNodes) SetCreateTime(v string) *ListClusterHyperNodesResponseBodyHyperNodes {
	s.CreateTime = &v
	return s
}

func (s *ListClusterHyperNodesResponseBodyHyperNodes) SetExpireTime(v string) *ListClusterHyperNodesResponseBodyHyperNodes {
	s.ExpireTime = &v
	return s
}

func (s *ListClusterHyperNodesResponseBodyHyperNodes) SetFileSystemMountEnabled(v bool) *ListClusterHyperNodesResponseBodyHyperNodes {
	s.FileSystemMountEnabled = &v
	return s
}

func (s *ListClusterHyperNodesResponseBodyHyperNodes) SetHostname(v string) *ListClusterHyperNodesResponseBodyHyperNodes {
	s.Hostname = &v
	return s
}

func (s *ListClusterHyperNodesResponseBodyHyperNodes) SetHpnZone(v string) *ListClusterHyperNodesResponseBodyHyperNodes {
	s.HpnZone = &v
	return s
}

func (s *ListClusterHyperNodesResponseBodyHyperNodes) SetHyperNodeId(v string) *ListClusterHyperNodesResponseBodyHyperNodes {
	s.HyperNodeId = &v
	return s
}

func (s *ListClusterHyperNodesResponseBodyHyperNodes) SetMachineType(v string) *ListClusterHyperNodesResponseBodyHyperNodes {
	s.MachineType = &v
	return s
}

func (s *ListClusterHyperNodesResponseBodyHyperNodes) SetNodeGroupId(v string) *ListClusterHyperNodesResponseBodyHyperNodes {
	s.NodeGroupId = &v
	return s
}

func (s *ListClusterHyperNodesResponseBodyHyperNodes) SetNodeGroupName(v string) *ListClusterHyperNodesResponseBodyHyperNodes {
	s.NodeGroupName = &v
	return s
}

func (s *ListClusterHyperNodesResponseBodyHyperNodes) SetStatus(v string) *ListClusterHyperNodesResponseBodyHyperNodes {
	s.Status = &v
	return s
}

func (s *ListClusterHyperNodesResponseBodyHyperNodes) SetTags(v []*ListClusterHyperNodesResponseBodyHyperNodesTags) *ListClusterHyperNodesResponseBodyHyperNodes {
	s.Tags = v
	return s
}

func (s *ListClusterHyperNodesResponseBodyHyperNodes) SetTaskId(v string) *ListClusterHyperNodesResponseBodyHyperNodes {
	s.TaskId = &v
	return s
}

func (s *ListClusterHyperNodesResponseBodyHyperNodes) SetVSwitchId(v string) *ListClusterHyperNodesResponseBodyHyperNodes {
	s.VSwitchId = &v
	return s
}

func (s *ListClusterHyperNodesResponseBodyHyperNodes) SetVpcId(v string) *ListClusterHyperNodesResponseBodyHyperNodes {
	s.VpcId = &v
	return s
}

func (s *ListClusterHyperNodesResponseBodyHyperNodes) SetZoneId(v string) *ListClusterHyperNodesResponseBodyHyperNodes {
	s.ZoneId = &v
	return s
}

func (s *ListClusterHyperNodesResponseBodyHyperNodes) Validate() error {
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

type ListClusterHyperNodesResponseBodyHyperNodesTags struct {
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// ali-zeekr-ota-doris-prod-hz-selectdb
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListClusterHyperNodesResponseBodyHyperNodesTags) String() string {
	return dara.Prettify(s)
}

func (s ListClusterHyperNodesResponseBodyHyperNodesTags) GoString() string {
	return s.String()
}

func (s *ListClusterHyperNodesResponseBodyHyperNodesTags) GetKey() *string {
	return s.Key
}

func (s *ListClusterHyperNodesResponseBodyHyperNodesTags) GetValue() *string {
	return s.Value
}

func (s *ListClusterHyperNodesResponseBodyHyperNodesTags) SetKey(v string) *ListClusterHyperNodesResponseBodyHyperNodesTags {
	s.Key = &v
	return s
}

func (s *ListClusterHyperNodesResponseBodyHyperNodesTags) SetValue(v string) *ListClusterHyperNodesResponseBodyHyperNodesTags {
	s.Value = &v
	return s
}

func (s *ListClusterHyperNodesResponseBodyHyperNodesTags) Validate() error {
	return dara.Validate(s)
}
