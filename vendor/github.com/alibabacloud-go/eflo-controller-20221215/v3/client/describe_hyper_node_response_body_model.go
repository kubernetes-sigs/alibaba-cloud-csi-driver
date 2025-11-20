// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHyperNodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeHyperNodeResponseBody
	GetClusterId() *string
	SetClusterName(v string) *DescribeHyperNodeResponseBody
	GetClusterName() *string
	SetCreateTime(v string) *DescribeHyperNodeResponseBody
	GetCreateTime() *string
	SetExpireTime(v string) *DescribeHyperNodeResponseBody
	GetExpireTime() *string
	SetFileSystemMountEnabled(v bool) *DescribeHyperNodeResponseBody
	GetFileSystemMountEnabled() *bool
	SetHostname(v string) *DescribeHyperNodeResponseBody
	GetHostname() *string
	SetHpnZone(v string) *DescribeHyperNodeResponseBody
	GetHpnZone() *string
	SetHyperNodeId(v string) *DescribeHyperNodeResponseBody
	GetHyperNodeId() *string
	SetMachineType(v string) *DescribeHyperNodeResponseBody
	GetMachineType() *string
	SetNodeGroupId(v string) *DescribeHyperNodeResponseBody
	GetNodeGroupId() *string
	SetNodeGroupName(v string) *DescribeHyperNodeResponseBody
	GetNodeGroupName() *string
	SetNodes(v []*DescribeHyperNodeResponseBodyNodes) *DescribeHyperNodeResponseBody
	GetNodes() []*DescribeHyperNodeResponseBodyNodes
	SetOperatingState(v string) *DescribeHyperNodeResponseBody
	GetOperatingState() *string
	SetRequestId(v string) *DescribeHyperNodeResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *DescribeHyperNodeResponseBody
	GetResourceGroupId() *string
	SetStatus(v string) *DescribeHyperNodeResponseBody
	GetStatus() *string
	SetZoneId(v string) *DescribeHyperNodeResponseBody
	GetZoneId() *string
}

type DescribeHyperNodeResponseBody struct {
	// example:
	//
	// i119982311660892626523
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// Standard_Cluster
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// example:
	//
	// 2022-11-30T02:00:00.852Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2125-06-24T16:52:44.318000
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// True
	FileSystemMountEnabled *bool `json:"FileSystemMountEnabled,omitempty" xml:"FileSystemMountEnabled,omitempty"`
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
	// e01-cn-zvp2tgykr08
	HyperNodeId *string `json:"HyperNodeId,omitempty" xml:"HyperNodeId,omitempty"`
	// example:
	//
	// efg1.nvga1
	MachineType *string `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	// example:
	//
	// ng-ec3c96ff0aa4c60d
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// example:
	//
	// lisan-nodegroup
	NodeGroupName *string                               `json:"NodeGroupName,omitempty" xml:"NodeGroupName,omitempty"`
	Nodes         []*DescribeHyperNodeResponseBodyNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// example:
	//
	// Using
	OperatingState *string `json:"OperatingState,omitempty" xml:"OperatingState,omitempty"`
	// example:
	//
	// D6058705-1C45-35C9-9461-02504897D4D0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// rg-acfmxno4vh5****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Deprecated
	//
	// example:
	//
	// Operating
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// cn-wulanchabu-c
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeHyperNodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHyperNodeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHyperNodeResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeHyperNodeResponseBody) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeHyperNodeResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeHyperNodeResponseBody) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeHyperNodeResponseBody) GetFileSystemMountEnabled() *bool {
	return s.FileSystemMountEnabled
}

func (s *DescribeHyperNodeResponseBody) GetHostname() *string {
	return s.Hostname
}

func (s *DescribeHyperNodeResponseBody) GetHpnZone() *string {
	return s.HpnZone
}

func (s *DescribeHyperNodeResponseBody) GetHyperNodeId() *string {
	return s.HyperNodeId
}

func (s *DescribeHyperNodeResponseBody) GetMachineType() *string {
	return s.MachineType
}

func (s *DescribeHyperNodeResponseBody) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *DescribeHyperNodeResponseBody) GetNodeGroupName() *string {
	return s.NodeGroupName
}

func (s *DescribeHyperNodeResponseBody) GetNodes() []*DescribeHyperNodeResponseBodyNodes {
	return s.Nodes
}

func (s *DescribeHyperNodeResponseBody) GetOperatingState() *string {
	return s.OperatingState
}

func (s *DescribeHyperNodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHyperNodeResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeHyperNodeResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeHyperNodeResponseBody) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeHyperNodeResponseBody) SetClusterId(v string) *DescribeHyperNodeResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeHyperNodeResponseBody) SetClusterName(v string) *DescribeHyperNodeResponseBody {
	s.ClusterName = &v
	return s
}

func (s *DescribeHyperNodeResponseBody) SetCreateTime(v string) *DescribeHyperNodeResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeHyperNodeResponseBody) SetExpireTime(v string) *DescribeHyperNodeResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *DescribeHyperNodeResponseBody) SetFileSystemMountEnabled(v bool) *DescribeHyperNodeResponseBody {
	s.FileSystemMountEnabled = &v
	return s
}

func (s *DescribeHyperNodeResponseBody) SetHostname(v string) *DescribeHyperNodeResponseBody {
	s.Hostname = &v
	return s
}

func (s *DescribeHyperNodeResponseBody) SetHpnZone(v string) *DescribeHyperNodeResponseBody {
	s.HpnZone = &v
	return s
}

func (s *DescribeHyperNodeResponseBody) SetHyperNodeId(v string) *DescribeHyperNodeResponseBody {
	s.HyperNodeId = &v
	return s
}

func (s *DescribeHyperNodeResponseBody) SetMachineType(v string) *DescribeHyperNodeResponseBody {
	s.MachineType = &v
	return s
}

func (s *DescribeHyperNodeResponseBody) SetNodeGroupId(v string) *DescribeHyperNodeResponseBody {
	s.NodeGroupId = &v
	return s
}

func (s *DescribeHyperNodeResponseBody) SetNodeGroupName(v string) *DescribeHyperNodeResponseBody {
	s.NodeGroupName = &v
	return s
}

func (s *DescribeHyperNodeResponseBody) SetNodes(v []*DescribeHyperNodeResponseBodyNodes) *DescribeHyperNodeResponseBody {
	s.Nodes = v
	return s
}

func (s *DescribeHyperNodeResponseBody) SetOperatingState(v string) *DescribeHyperNodeResponseBody {
	s.OperatingState = &v
	return s
}

func (s *DescribeHyperNodeResponseBody) SetRequestId(v string) *DescribeHyperNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHyperNodeResponseBody) SetResourceGroupId(v string) *DescribeHyperNodeResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeHyperNodeResponseBody) SetStatus(v string) *DescribeHyperNodeResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeHyperNodeResponseBody) SetZoneId(v string) *DescribeHyperNodeResponseBody {
	s.ZoneId = &v
	return s
}

func (s *DescribeHyperNodeResponseBody) Validate() error {
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

type DescribeHyperNodeResponseBodyNodes struct {
	Disks []*DescribeHyperNodeResponseBodyNodesDisks `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Repeated"`
	// example:
	//
	// 457db5ca-241d-11ed-9fd7-acde48001122
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// example:
	//
	// i190297201669099844192
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// example:
	//
	// Alinux3_x86_gu8xf_P_Host_D3_C7E3_550.127_Legacy_N_241230
	ImageName *string                                     `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	Networks  *DescribeHyperNodeResponseBodyNodesNetworks `json:"Networks,omitempty" xml:"Networks,omitempty" type:"Struct"`
	// example:
	//
	// e01-cn-zvp2tgykr0b
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// example:
	//
	// Using
	OperatingState *string `json:"OperatingState,omitempty" xml:"OperatingState,omitempty"`
	// Deprecated
	//
	// example:
	//
	// InProgress
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// #!/bin/bash
	//
	// uptime
	//
	// echo "aaaaaaa" >> /tmp/ttttttest.sh
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s DescribeHyperNodeResponseBodyNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeHyperNodeResponseBodyNodes) GoString() string {
	return s.String()
}

func (s *DescribeHyperNodeResponseBodyNodes) GetDisks() []*DescribeHyperNodeResponseBodyNodesDisks {
	return s.Disks
}

func (s *DescribeHyperNodeResponseBodyNodes) GetHostname() *string {
	return s.Hostname
}

func (s *DescribeHyperNodeResponseBodyNodes) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeHyperNodeResponseBodyNodes) GetImageName() *string {
	return s.ImageName
}

func (s *DescribeHyperNodeResponseBodyNodes) GetNetworks() *DescribeHyperNodeResponseBodyNodesNetworks {
	return s.Networks
}

func (s *DescribeHyperNodeResponseBodyNodes) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeHyperNodeResponseBodyNodes) GetOperatingState() *string {
	return s.OperatingState
}

func (s *DescribeHyperNodeResponseBodyNodes) GetStatus() *string {
	return s.Status
}

func (s *DescribeHyperNodeResponseBodyNodes) GetUserData() *string {
	return s.UserData
}

func (s *DescribeHyperNodeResponseBodyNodes) SetDisks(v []*DescribeHyperNodeResponseBodyNodesDisks) *DescribeHyperNodeResponseBodyNodes {
	s.Disks = v
	return s
}

func (s *DescribeHyperNodeResponseBodyNodes) SetHostname(v string) *DescribeHyperNodeResponseBodyNodes {
	s.Hostname = &v
	return s
}

func (s *DescribeHyperNodeResponseBodyNodes) SetImageId(v string) *DescribeHyperNodeResponseBodyNodes {
	s.ImageId = &v
	return s
}

func (s *DescribeHyperNodeResponseBodyNodes) SetImageName(v string) *DescribeHyperNodeResponseBodyNodes {
	s.ImageName = &v
	return s
}

func (s *DescribeHyperNodeResponseBodyNodes) SetNetworks(v *DescribeHyperNodeResponseBodyNodesNetworks) *DescribeHyperNodeResponseBodyNodes {
	s.Networks = v
	return s
}

func (s *DescribeHyperNodeResponseBodyNodes) SetNodeId(v string) *DescribeHyperNodeResponseBodyNodes {
	s.NodeId = &v
	return s
}

func (s *DescribeHyperNodeResponseBodyNodes) SetOperatingState(v string) *DescribeHyperNodeResponseBodyNodes {
	s.OperatingState = &v
	return s
}

func (s *DescribeHyperNodeResponseBodyNodes) SetStatus(v string) *DescribeHyperNodeResponseBodyNodes {
	s.Status = &v
	return s
}

func (s *DescribeHyperNodeResponseBodyNodes) SetUserData(v string) *DescribeHyperNodeResponseBodyNodes {
	s.UserData = &v
	return s
}

func (s *DescribeHyperNodeResponseBodyNodes) Validate() error {
	if s.Disks != nil {
		for _, item := range s.Disks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Networks != nil {
		if err := s.Networks.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeHyperNodeResponseBodyNodesDisks struct {
	// example:
	//
	// DOWNLINK_PACKET
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// d-bp1564bcc2306uui4zpk
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// example:
	//
	// PL0
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// example:
	//
	// 5
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// system
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeHyperNodeResponseBodyNodesDisks) String() string {
	return dara.Prettify(s)
}

func (s DescribeHyperNodeResponseBodyNodesDisks) GoString() string {
	return s.String()
}

func (s *DescribeHyperNodeResponseBodyNodesDisks) GetCategory() *string {
	return s.Category
}

func (s *DescribeHyperNodeResponseBodyNodesDisks) GetDiskId() *string {
	return s.DiskId
}

func (s *DescribeHyperNodeResponseBodyNodesDisks) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *DescribeHyperNodeResponseBodyNodesDisks) GetSize() *int32 {
	return s.Size
}

func (s *DescribeHyperNodeResponseBodyNodesDisks) GetType() *string {
	return s.Type
}

func (s *DescribeHyperNodeResponseBodyNodesDisks) SetCategory(v string) *DescribeHyperNodeResponseBodyNodesDisks {
	s.Category = &v
	return s
}

func (s *DescribeHyperNodeResponseBodyNodesDisks) SetDiskId(v string) *DescribeHyperNodeResponseBodyNodesDisks {
	s.DiskId = &v
	return s
}

func (s *DescribeHyperNodeResponseBodyNodesDisks) SetPerformanceLevel(v string) *DescribeHyperNodeResponseBodyNodesDisks {
	s.PerformanceLevel = &v
	return s
}

func (s *DescribeHyperNodeResponseBodyNodesDisks) SetSize(v int32) *DescribeHyperNodeResponseBodyNodesDisks {
	s.Size = &v
	return s
}

func (s *DescribeHyperNodeResponseBodyNodesDisks) SetType(v string) *DescribeHyperNodeResponseBodyNodesDisks {
	s.Type = &v
	return s
}

func (s *DescribeHyperNodeResponseBodyNodesDisks) Validate() error {
	return dara.Validate(s)
}

type DescribeHyperNodeResponseBodyNodesNetworks struct {
	// example:
	//
	// bond0
	BondName *string `json:"BondName,omitempty" xml:"BondName,omitempty"`
	// example:
	//
	// 192.168.22.2
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s DescribeHyperNodeResponseBodyNodesNetworks) String() string {
	return dara.Prettify(s)
}

func (s DescribeHyperNodeResponseBodyNodesNetworks) GoString() string {
	return s.String()
}

func (s *DescribeHyperNodeResponseBodyNodesNetworks) GetBondName() *string {
	return s.BondName
}

func (s *DescribeHyperNodeResponseBodyNodesNetworks) GetIp() *string {
	return s.Ip
}

func (s *DescribeHyperNodeResponseBodyNodesNetworks) SetBondName(v string) *DescribeHyperNodeResponseBodyNodesNetworks {
	s.BondName = &v
	return s
}

func (s *DescribeHyperNodeResponseBodyNodesNetworks) SetIp(v string) *DescribeHyperNodeResponseBodyNodesNetworks {
	s.Ip = &v
	return s
}

func (s *DescribeHyperNodeResponseBodyNodesNetworks) Validate() error {
	return dara.Validate(s)
}
