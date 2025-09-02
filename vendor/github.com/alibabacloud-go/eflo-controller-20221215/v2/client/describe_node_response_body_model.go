// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeNodeResponseBody
	GetClusterId() *string
	SetClusterName(v string) *DescribeNodeResponseBody
	GetClusterName() *string
	SetCreateTime(v string) *DescribeNodeResponseBody
	GetCreateTime() *string
	SetDisks(v []*DescribeNodeResponseBodyDisks) *DescribeNodeResponseBody
	GetDisks() []*DescribeNodeResponseBodyDisks
	SetExpiredTime(v string) *DescribeNodeResponseBody
	GetExpiredTime() *string
	SetFileSystemMountEnabled(v bool) *DescribeNodeResponseBody
	GetFileSystemMountEnabled() *bool
	SetHostname(v string) *DescribeNodeResponseBody
	GetHostname() *string
	SetHpnZone(v string) *DescribeNodeResponseBody
	GetHpnZone() *string
	SetImageId(v string) *DescribeNodeResponseBody
	GetImageId() *string
	SetImageName(v string) *DescribeNodeResponseBody
	GetImageName() *string
	SetMachineType(v string) *DescribeNodeResponseBody
	GetMachineType() *string
	SetNetworks(v []*DescribeNodeResponseBodyNetworks) *DescribeNodeResponseBody
	GetNetworks() []*DescribeNodeResponseBodyNetworks
	SetNodeGroupId(v string) *DescribeNodeResponseBody
	GetNodeGroupId() *string
	SetNodeGroupName(v string) *DescribeNodeResponseBody
	GetNodeGroupName() *string
	SetNodeId(v string) *DescribeNodeResponseBody
	GetNodeId() *string
	SetNodeType(v string) *DescribeNodeResponseBody
	GetNodeType() *string
	SetOperatingState(v string) *DescribeNodeResponseBody
	GetOperatingState() *string
	SetRequestId(v string) *DescribeNodeResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *DescribeNodeResponseBody
	GetResourceGroupId() *string
	SetSn(v string) *DescribeNodeResponseBody
	GetSn() *string
	SetUserData(v string) *DescribeNodeResponseBody
	GetUserData() *string
	SetZoneId(v string) *DescribeNodeResponseBody
	GetZoneId() *string
}

type DescribeNodeResponseBody struct {
	// The cluster ID.
	//
	// example:
	//
	// i116913051662373010974
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// Standard_Cluster
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2022-09-30T03:35:53Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The disks.
	Disks []*DescribeNodeResponseBodyDisks `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Repeated"`
	// The expiration time.
	//
	// example:
	//
	// 2022-06-23T16:00:00Z
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// Indicates whether file storage mounting is supported.
	//
	// example:
	//
	// False
	FileSystemMountEnabled *bool `json:"FileSystemMountEnabled,omitempty" xml:"FileSystemMountEnabled,omitempty"`
	// The hostname.
	//
	// example:
	//
	// 31d38530-241e-11ed-bc63-acde48001122
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// The cluster number.
	//
	// example:
	//
	// A1
	HpnZone *string `json:"HpnZone,omitempty" xml:"HpnZone,omitempty"`
	// The image ID.
	//
	// example:
	//
	// i190297201634099844192
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The image name.
	//
	// example:
	//
	// Centos7.9_all_0811
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The instance type.
	//
	// example:
	//
	// efg1.nvga1
	MachineType *string `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	// The network information.
	Networks []*DescribeNodeResponseBodyNetworks `json:"Networks,omitempty" xml:"Networks,omitempty" type:"Repeated"`
	// The node group ID.
	//
	// example:
	//
	// ng-ec3c96ff0aa4c60d
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// The node group name.
	//
	// example:
	//
	// emr-default
	NodeGroupName *string `json:"NodeGroupName,omitempty" xml:"NodeGroupName,omitempty"`
	// The node ID.
	//
	// example:
	//
	// e01-cn-zvp2tgykr08
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// example:
	//
	// standard
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The node status.
	//
	// Valid values:
	//
	// 	- Extending
	//
	// 	- UnusedNodeStopped
	//
	// 	- UnusedNodeStopping
	//
	// 	- Unused
	//
	// 	- Using
	//
	// 	- ReleaseLocking
	//
	// 	- Operating
	//
	// 	- Cutting
	//
	// 	- ClusterNodeStopped
	//
	// 	- UnusedNodeRecovering
	//
	// 	- ClusterNodeStopping
	//
	// 	- ClusterNodeRecovering
	//
	// 	- Replacing
	//
	// example:
	//
	// Using
	OperatingState *string `json:"OperatingState,omitempty" xml:"OperatingState,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AC4F0004-7BCE-52E0-891B-CAC7D64E3368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmywpvugkh7kq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The serial number of the node.
	//
	// example:
	//
	// sag42ckf4jx
	Sn *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
	// The custom script.
	//
	// example:
	//
	// #!/bin/bash
	//
	// uptime
	//
	// echo "aaaaaaa" >> /tmp/ttttt20250110141010.sh
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeNodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNodeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNodeResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeNodeResponseBody) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeNodeResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeNodeResponseBody) GetDisks() []*DescribeNodeResponseBodyDisks {
	return s.Disks
}

func (s *DescribeNodeResponseBody) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeNodeResponseBody) GetFileSystemMountEnabled() *bool {
	return s.FileSystemMountEnabled
}

func (s *DescribeNodeResponseBody) GetHostname() *string {
	return s.Hostname
}

func (s *DescribeNodeResponseBody) GetHpnZone() *string {
	return s.HpnZone
}

func (s *DescribeNodeResponseBody) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeNodeResponseBody) GetImageName() *string {
	return s.ImageName
}

func (s *DescribeNodeResponseBody) GetMachineType() *string {
	return s.MachineType
}

func (s *DescribeNodeResponseBody) GetNetworks() []*DescribeNodeResponseBodyNetworks {
	return s.Networks
}

func (s *DescribeNodeResponseBody) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *DescribeNodeResponseBody) GetNodeGroupName() *string {
	return s.NodeGroupName
}

func (s *DescribeNodeResponseBody) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeNodeResponseBody) GetNodeType() *string {
	return s.NodeType
}

func (s *DescribeNodeResponseBody) GetOperatingState() *string {
	return s.OperatingState
}

func (s *DescribeNodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNodeResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeNodeResponseBody) GetSn() *string {
	return s.Sn
}

func (s *DescribeNodeResponseBody) GetUserData() *string {
	return s.UserData
}

func (s *DescribeNodeResponseBody) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeNodeResponseBody) SetClusterId(v string) *DescribeNodeResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeNodeResponseBody) SetClusterName(v string) *DescribeNodeResponseBody {
	s.ClusterName = &v
	return s
}

func (s *DescribeNodeResponseBody) SetCreateTime(v string) *DescribeNodeResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeNodeResponseBody) SetDisks(v []*DescribeNodeResponseBodyDisks) *DescribeNodeResponseBody {
	s.Disks = v
	return s
}

func (s *DescribeNodeResponseBody) SetExpiredTime(v string) *DescribeNodeResponseBody {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeNodeResponseBody) SetFileSystemMountEnabled(v bool) *DescribeNodeResponseBody {
	s.FileSystemMountEnabled = &v
	return s
}

func (s *DescribeNodeResponseBody) SetHostname(v string) *DescribeNodeResponseBody {
	s.Hostname = &v
	return s
}

func (s *DescribeNodeResponseBody) SetHpnZone(v string) *DescribeNodeResponseBody {
	s.HpnZone = &v
	return s
}

func (s *DescribeNodeResponseBody) SetImageId(v string) *DescribeNodeResponseBody {
	s.ImageId = &v
	return s
}

func (s *DescribeNodeResponseBody) SetImageName(v string) *DescribeNodeResponseBody {
	s.ImageName = &v
	return s
}

func (s *DescribeNodeResponseBody) SetMachineType(v string) *DescribeNodeResponseBody {
	s.MachineType = &v
	return s
}

func (s *DescribeNodeResponseBody) SetNetworks(v []*DescribeNodeResponseBodyNetworks) *DescribeNodeResponseBody {
	s.Networks = v
	return s
}

func (s *DescribeNodeResponseBody) SetNodeGroupId(v string) *DescribeNodeResponseBody {
	s.NodeGroupId = &v
	return s
}

func (s *DescribeNodeResponseBody) SetNodeGroupName(v string) *DescribeNodeResponseBody {
	s.NodeGroupName = &v
	return s
}

func (s *DescribeNodeResponseBody) SetNodeId(v string) *DescribeNodeResponseBody {
	s.NodeId = &v
	return s
}

func (s *DescribeNodeResponseBody) SetNodeType(v string) *DescribeNodeResponseBody {
	s.NodeType = &v
	return s
}

func (s *DescribeNodeResponseBody) SetOperatingState(v string) *DescribeNodeResponseBody {
	s.OperatingState = &v
	return s
}

func (s *DescribeNodeResponseBody) SetRequestId(v string) *DescribeNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNodeResponseBody) SetResourceGroupId(v string) *DescribeNodeResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeNodeResponseBody) SetSn(v string) *DescribeNodeResponseBody {
	s.Sn = &v
	return s
}

func (s *DescribeNodeResponseBody) SetUserData(v string) *DescribeNodeResponseBody {
	s.UserData = &v
	return s
}

func (s *DescribeNodeResponseBody) SetZoneId(v string) *DescribeNodeResponseBody {
	s.ZoneId = &v
	return s
}

func (s *DescribeNodeResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeNodeResponseBodyDisks struct {
	// The disk type. Valid values:
	//
	// 	- cloud_essd
	//
	// example:
	//
	// cloud_essd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The disk ID.
	//
	// example:
	//
	// d-bp1fi88ryk4yah8a6yos
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The performance level of the ESSD that is used as the system disk. Valid values:
	//
	// 	- PL0: A single ESSD can deliver up to 10,000 random read/write IOPS.
	//
	// 	- PL1: A single ESSD can deliver up to 50,000 random read/write IOPS.
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The disk size. Unit: GiB.
	//
	// example:
	//
	// 100
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The disk type. Valid values:
	//
	// 	- system: system disk
	//
	// 	- data: data disk
	//
	// example:
	//
	// system
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeNodeResponseBodyDisks) String() string {
	return dara.Prettify(s)
}

func (s DescribeNodeResponseBodyDisks) GoString() string {
	return s.String()
}

func (s *DescribeNodeResponseBodyDisks) GetCategory() *string {
	return s.Category
}

func (s *DescribeNodeResponseBodyDisks) GetDiskId() *string {
	return s.DiskId
}

func (s *DescribeNodeResponseBodyDisks) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *DescribeNodeResponseBodyDisks) GetSize() *int32 {
	return s.Size
}

func (s *DescribeNodeResponseBodyDisks) GetType() *string {
	return s.Type
}

func (s *DescribeNodeResponseBodyDisks) SetCategory(v string) *DescribeNodeResponseBodyDisks {
	s.Category = &v
	return s
}

func (s *DescribeNodeResponseBodyDisks) SetDiskId(v string) *DescribeNodeResponseBodyDisks {
	s.DiskId = &v
	return s
}

func (s *DescribeNodeResponseBodyDisks) SetPerformanceLevel(v string) *DescribeNodeResponseBodyDisks {
	s.PerformanceLevel = &v
	return s
}

func (s *DescribeNodeResponseBodyDisks) SetSize(v int32) *DescribeNodeResponseBodyDisks {
	s.Size = &v
	return s
}

func (s *DescribeNodeResponseBodyDisks) SetType(v string) *DescribeNodeResponseBodyDisks {
	s.Type = &v
	return s
}

func (s *DescribeNodeResponseBodyDisks) Validate() error {
	return dara.Validate(s)
}

type DescribeNodeResponseBodyNetworks struct {
	// The port information of the elastic network interface (ENI).
	//
	// example:
	//
	// Bond0
	BondName *string `json:"BondName,omitempty" xml:"BondName,omitempty"`
	// The IP address of the node.
	//
	// example:
	//
	// 47.254.235.44
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The subnet ID.
	//
	// example:
	//
	// vsw-uf68v51fldm5egmui5a6k
	SubnetId *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
	// The ID of the cluster network.
	//
	// example:
	//
	// vpd-xcuhjyrj
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
}

func (s DescribeNodeResponseBodyNetworks) String() string {
	return dara.Prettify(s)
}

func (s DescribeNodeResponseBodyNetworks) GoString() string {
	return s.String()
}

func (s *DescribeNodeResponseBodyNetworks) GetBondName() *string {
	return s.BondName
}

func (s *DescribeNodeResponseBodyNetworks) GetIp() *string {
	return s.Ip
}

func (s *DescribeNodeResponseBodyNetworks) GetSubnetId() *string {
	return s.SubnetId
}

func (s *DescribeNodeResponseBodyNetworks) GetVpdId() *string {
	return s.VpdId
}

func (s *DescribeNodeResponseBodyNetworks) SetBondName(v string) *DescribeNodeResponseBodyNetworks {
	s.BondName = &v
	return s
}

func (s *DescribeNodeResponseBodyNetworks) SetIp(v string) *DescribeNodeResponseBodyNetworks {
	s.Ip = &v
	return s
}

func (s *DescribeNodeResponseBodyNetworks) SetSubnetId(v string) *DescribeNodeResponseBodyNetworks {
	s.SubnetId = &v
	return s
}

func (s *DescribeNodeResponseBodyNetworks) SetVpdId(v string) *DescribeNodeResponseBodyNetworks {
	s.VpdId = &v
	return s
}

func (s *DescribeNodeResponseBodyNetworks) Validate() error {
	return dara.Validate(s)
}
