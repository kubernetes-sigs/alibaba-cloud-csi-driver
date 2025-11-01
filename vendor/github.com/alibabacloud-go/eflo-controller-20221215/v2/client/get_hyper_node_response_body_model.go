// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHyperNodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetHyperNodeResponseBody
	GetClusterId() *string
	SetClusterName(v string) *GetHyperNodeResponseBody
	GetClusterName() *string
	SetCreateTime(v string) *GetHyperNodeResponseBody
	GetCreateTime() *string
	SetExpireTime(v string) *GetHyperNodeResponseBody
	GetExpireTime() *string
	SetFileSystemMountEnabled(v bool) *GetHyperNodeResponseBody
	GetFileSystemMountEnabled() *bool
	SetHostname(v string) *GetHyperNodeResponseBody
	GetHostname() *string
	SetHpnZone(v string) *GetHyperNodeResponseBody
	GetHpnZone() *string
	SetHyperNodeId(v string) *GetHyperNodeResponseBody
	GetHyperNodeId() *string
	SetMachineType(v string) *GetHyperNodeResponseBody
	GetMachineType() *string
	SetNodeGroupId(v string) *GetHyperNodeResponseBody
	GetNodeGroupId() *string
	SetNodeGroupName(v string) *GetHyperNodeResponseBody
	GetNodeGroupName() *string
	SetNodes(v []*GetHyperNodeResponseBodyNodes) *GetHyperNodeResponseBody
	GetNodes() []*GetHyperNodeResponseBodyNodes
	SetRequestId(v string) *GetHyperNodeResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *GetHyperNodeResponseBody
	GetResourceGroupId() *string
	SetStatus(v string) *GetHyperNodeResponseBody
	GetStatus() *string
	SetZoneId(v string) *GetHyperNodeResponseBody
	GetZoneId() *string
}

type GetHyperNodeResponseBody struct {
	// example:
	//
	// i112138561737531371671
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
	// 2026-05-14T00:00:00
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
	// efg2.C48cNHmcn
	MachineType *string `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	// example:
	//
	// ng-ec3c96ff0aa4c60d
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// example:
	//
	// emr-default
	NodeGroupName *string                          `json:"NodeGroupName,omitempty" xml:"NodeGroupName,omitempty"`
	Nodes         []*GetHyperNodeResponseBodyNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// example:
	//
	// 2FE2B22C-CF9D-59DE-BF63-DC9B9B33A9D1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// rg-acfmwfm33rlt6zi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// Using
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// cn-wulanchabu-c
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetHyperNodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHyperNodeResponseBody) GoString() string {
	return s.String()
}

func (s *GetHyperNodeResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetHyperNodeResponseBody) GetClusterName() *string {
	return s.ClusterName
}

func (s *GetHyperNodeResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetHyperNodeResponseBody) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *GetHyperNodeResponseBody) GetFileSystemMountEnabled() *bool {
	return s.FileSystemMountEnabled
}

func (s *GetHyperNodeResponseBody) GetHostname() *string {
	return s.Hostname
}

func (s *GetHyperNodeResponseBody) GetHpnZone() *string {
	return s.HpnZone
}

func (s *GetHyperNodeResponseBody) GetHyperNodeId() *string {
	return s.HyperNodeId
}

func (s *GetHyperNodeResponseBody) GetMachineType() *string {
	return s.MachineType
}

func (s *GetHyperNodeResponseBody) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *GetHyperNodeResponseBody) GetNodeGroupName() *string {
	return s.NodeGroupName
}

func (s *GetHyperNodeResponseBody) GetNodes() []*GetHyperNodeResponseBodyNodes {
	return s.Nodes
}

func (s *GetHyperNodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHyperNodeResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetHyperNodeResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetHyperNodeResponseBody) GetZoneId() *string {
	return s.ZoneId
}

func (s *GetHyperNodeResponseBody) SetClusterId(v string) *GetHyperNodeResponseBody {
	s.ClusterId = &v
	return s
}

func (s *GetHyperNodeResponseBody) SetClusterName(v string) *GetHyperNodeResponseBody {
	s.ClusterName = &v
	return s
}

func (s *GetHyperNodeResponseBody) SetCreateTime(v string) *GetHyperNodeResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetHyperNodeResponseBody) SetExpireTime(v string) *GetHyperNodeResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *GetHyperNodeResponseBody) SetFileSystemMountEnabled(v bool) *GetHyperNodeResponseBody {
	s.FileSystemMountEnabled = &v
	return s
}

func (s *GetHyperNodeResponseBody) SetHostname(v string) *GetHyperNodeResponseBody {
	s.Hostname = &v
	return s
}

func (s *GetHyperNodeResponseBody) SetHpnZone(v string) *GetHyperNodeResponseBody {
	s.HpnZone = &v
	return s
}

func (s *GetHyperNodeResponseBody) SetHyperNodeId(v string) *GetHyperNodeResponseBody {
	s.HyperNodeId = &v
	return s
}

func (s *GetHyperNodeResponseBody) SetMachineType(v string) *GetHyperNodeResponseBody {
	s.MachineType = &v
	return s
}

func (s *GetHyperNodeResponseBody) SetNodeGroupId(v string) *GetHyperNodeResponseBody {
	s.NodeGroupId = &v
	return s
}

func (s *GetHyperNodeResponseBody) SetNodeGroupName(v string) *GetHyperNodeResponseBody {
	s.NodeGroupName = &v
	return s
}

func (s *GetHyperNodeResponseBody) SetNodes(v []*GetHyperNodeResponseBodyNodes) *GetHyperNodeResponseBody {
	s.Nodes = v
	return s
}

func (s *GetHyperNodeResponseBody) SetRequestId(v string) *GetHyperNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHyperNodeResponseBody) SetResourceGroupId(v string) *GetHyperNodeResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetHyperNodeResponseBody) SetStatus(v string) *GetHyperNodeResponseBody {
	s.Status = &v
	return s
}

func (s *GetHyperNodeResponseBody) SetZoneId(v string) *GetHyperNodeResponseBody {
	s.ZoneId = &v
	return s
}

func (s *GetHyperNodeResponseBody) Validate() error {
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

type GetHyperNodeResponseBodyNodes struct {
	Disks []*GetHyperNodeResponseBodyNodesDisks `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Repeated"`
	// example:
	//
	// alywlcb-lingjun-gpu-0025
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// example:
	//
	// i190297201669099844192
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// example:
	//
	// Alinux3_x86_5.10.134-16.3_NV_RunC_D3_E3C7_570.133.20_V1.0_250428
	ImageName *string                                `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	Networks  *GetHyperNodeResponseBodyNodesNetworks `json:"Networks,omitempty" xml:"Networks,omitempty" type:"Struct"`
	// example:
	//
	// e01-cn-zvp2tgykr0b
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
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
	// echo "aaaaaaa" >> /tmp/ttttt20250110141010.sh
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s GetHyperNodeResponseBodyNodes) String() string {
	return dara.Prettify(s)
}

func (s GetHyperNodeResponseBodyNodes) GoString() string {
	return s.String()
}

func (s *GetHyperNodeResponseBodyNodes) GetDisks() []*GetHyperNodeResponseBodyNodesDisks {
	return s.Disks
}

func (s *GetHyperNodeResponseBodyNodes) GetHostname() *string {
	return s.Hostname
}

func (s *GetHyperNodeResponseBodyNodes) GetImageId() *string {
	return s.ImageId
}

func (s *GetHyperNodeResponseBodyNodes) GetImageName() *string {
	return s.ImageName
}

func (s *GetHyperNodeResponseBodyNodes) GetNetworks() *GetHyperNodeResponseBodyNodesNetworks {
	return s.Networks
}

func (s *GetHyperNodeResponseBodyNodes) GetNodeId() *string {
	return s.NodeId
}

func (s *GetHyperNodeResponseBodyNodes) GetStatus() *string {
	return s.Status
}

func (s *GetHyperNodeResponseBodyNodes) GetUserData() *string {
	return s.UserData
}

func (s *GetHyperNodeResponseBodyNodes) SetDisks(v []*GetHyperNodeResponseBodyNodesDisks) *GetHyperNodeResponseBodyNodes {
	s.Disks = v
	return s
}

func (s *GetHyperNodeResponseBodyNodes) SetHostname(v string) *GetHyperNodeResponseBodyNodes {
	s.Hostname = &v
	return s
}

func (s *GetHyperNodeResponseBodyNodes) SetImageId(v string) *GetHyperNodeResponseBodyNodes {
	s.ImageId = &v
	return s
}

func (s *GetHyperNodeResponseBodyNodes) SetImageName(v string) *GetHyperNodeResponseBodyNodes {
	s.ImageName = &v
	return s
}

func (s *GetHyperNodeResponseBodyNodes) SetNetworks(v *GetHyperNodeResponseBodyNodesNetworks) *GetHyperNodeResponseBodyNodes {
	s.Networks = v
	return s
}

func (s *GetHyperNodeResponseBodyNodes) SetNodeId(v string) *GetHyperNodeResponseBodyNodes {
	s.NodeId = &v
	return s
}

func (s *GetHyperNodeResponseBodyNodes) SetStatus(v string) *GetHyperNodeResponseBodyNodes {
	s.Status = &v
	return s
}

func (s *GetHyperNodeResponseBodyNodes) SetUserData(v string) *GetHyperNodeResponseBodyNodes {
	s.UserData = &v
	return s
}

func (s *GetHyperNodeResponseBodyNodes) Validate() error {
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

type GetHyperNodeResponseBodyNodesDisks struct {
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

func (s GetHyperNodeResponseBodyNodesDisks) String() string {
	return dara.Prettify(s)
}

func (s GetHyperNodeResponseBodyNodesDisks) GoString() string {
	return s.String()
}

func (s *GetHyperNodeResponseBodyNodesDisks) GetCategory() *string {
	return s.Category
}

func (s *GetHyperNodeResponseBodyNodesDisks) GetDiskId() *string {
	return s.DiskId
}

func (s *GetHyperNodeResponseBodyNodesDisks) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *GetHyperNodeResponseBodyNodesDisks) GetSize() *int32 {
	return s.Size
}

func (s *GetHyperNodeResponseBodyNodesDisks) GetType() *string {
	return s.Type
}

func (s *GetHyperNodeResponseBodyNodesDisks) SetCategory(v string) *GetHyperNodeResponseBodyNodesDisks {
	s.Category = &v
	return s
}

func (s *GetHyperNodeResponseBodyNodesDisks) SetDiskId(v string) *GetHyperNodeResponseBodyNodesDisks {
	s.DiskId = &v
	return s
}

func (s *GetHyperNodeResponseBodyNodesDisks) SetPerformanceLevel(v string) *GetHyperNodeResponseBodyNodesDisks {
	s.PerformanceLevel = &v
	return s
}

func (s *GetHyperNodeResponseBodyNodesDisks) SetSize(v int32) *GetHyperNodeResponseBodyNodesDisks {
	s.Size = &v
	return s
}

func (s *GetHyperNodeResponseBodyNodesDisks) SetType(v string) *GetHyperNodeResponseBodyNodesDisks {
	s.Type = &v
	return s
}

func (s *GetHyperNodeResponseBodyNodesDisks) Validate() error {
	return dara.Validate(s)
}

type GetHyperNodeResponseBodyNodesNetworks struct {
	// example:
	//
	// bond0
	BondName *string `json:"BondName,omitempty" xml:"BondName,omitempty"`
	// example:
	//
	// 172.17.231.113
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s GetHyperNodeResponseBodyNodesNetworks) String() string {
	return dara.Prettify(s)
}

func (s GetHyperNodeResponseBodyNodesNetworks) GoString() string {
	return s.String()
}

func (s *GetHyperNodeResponseBodyNodesNetworks) GetBondName() *string {
	return s.BondName
}

func (s *GetHyperNodeResponseBodyNodesNetworks) GetIp() *string {
	return s.Ip
}

func (s *GetHyperNodeResponseBodyNodesNetworks) SetBondName(v string) *GetHyperNodeResponseBodyNodesNetworks {
	s.BondName = &v
	return s
}

func (s *GetHyperNodeResponseBodyNodesNetworks) SetIp(v string) *GetHyperNodeResponseBodyNodesNetworks {
	s.Ip = &v
	return s
}

func (s *GetHyperNodeResponseBodyNodesNetworks) Validate() error {
	return dara.Validate(s)
}
