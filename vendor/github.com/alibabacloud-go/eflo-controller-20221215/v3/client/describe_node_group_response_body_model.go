// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNodeGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAz(v string) *DescribeNodeGroupResponseBody
	GetAz() *string
	SetClusterId(v string) *DescribeNodeGroupResponseBody
	GetClusterId() *string
	SetClusterName(v string) *DescribeNodeGroupResponseBody
	GetClusterName() *string
	SetCreateTime(v string) *DescribeNodeGroupResponseBody
	GetCreateTime() *string
	SetFileSystemMountEnabled(v bool) *DescribeNodeGroupResponseBody
	GetFileSystemMountEnabled() *bool
	SetImageId(v string) *DescribeNodeGroupResponseBody
	GetImageId() *string
	SetImageName(v string) *DescribeNodeGroupResponseBody
	GetImageName() *string
	SetKeyPairName(v string) *DescribeNodeGroupResponseBody
	GetKeyPairName() *string
	SetLoginType(v string) *DescribeNodeGroupResponseBody
	GetLoginType() *string
	SetMachineType(v string) *DescribeNodeGroupResponseBody
	GetMachineType() *string
	SetNodeCount(v string) *DescribeNodeGroupResponseBody
	GetNodeCount() *string
	SetNodeGroupDescription(v string) *DescribeNodeGroupResponseBody
	GetNodeGroupDescription() *string
	SetNodeGroupId(v string) *DescribeNodeGroupResponseBody
	GetNodeGroupId() *string
	SetNodeGroupName(v string) *DescribeNodeGroupResponseBody
	GetNodeGroupName() *string
	SetRamRoleName(v string) *DescribeNodeGroupResponseBody
	GetRamRoleName() *string
	SetRequestId(v string) *DescribeNodeGroupResponseBody
	GetRequestId() *string
	SetSystemDisk(v *DescribeNodeGroupResponseBodySystemDisk) *DescribeNodeGroupResponseBody
	GetSystemDisk() *DescribeNodeGroupResponseBodySystemDisk
	SetUpdateTime(v string) *DescribeNodeGroupResponseBody
	GetUpdateTime() *string
	SetUserData(v string) *DescribeNodeGroupResponseBody
	GetUserData() *string
	SetVirtualGpuEnabled(v bool) *DescribeNodeGroupResponseBody
	GetVirtualGpuEnabled() *bool
}

type DescribeNodeGroupResponseBody struct {
	// The zone ID.
	//
	// example:
	//
	// cn-wulanchabu-a
	Az *string `json:"Az,omitempty" xml:"Az,omitempty"`
	// The ID of the cluster.
	//
	// example:
	//
	// i111987311754895199538
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// aliyun-basic-cluster
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2025-07-07T17:19:42.980000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether file storage can be mounted.
	//
	// example:
	//
	// True
	FileSystemMountEnabled *bool `json:"FileSystemMountEnabled,omitempty" xml:"FileSystemMountEnabled,omitempty"`
	// The image ID.
	//
	// example:
	//
	// i190720111752146430748
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The name of the image.
	//
	// example:
	//
	// Alinux3_x86_5.10.134-16.3_NV_RunC_D3_E3C7_570.133.20_V1.0_250428
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The name of the key pair.
	//
	// example:
	//
	// k8s-key
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// Valid values:
	//
	// • password: The node group supports only password-based logon.
	//
	// • keypair: The node group supports only key pair-based logon.
	//
	// • both: The node group supports password-based and key pair-based logon.
	//
	// If this parameter is not returned, no logon method is configured for the node group.
	//
	// example:
	//
	// password
	LoginType *string `json:"LoginType,omitempty" xml:"LoginType,omitempty"`
	// The instance type.
	//
	// example:
	//
	// efg2.NH2cn
	MachineType *string `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	// The number of nodes.
	//
	// example:
	//
	// 1
	NodeCount *string `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	// The description of the node group.
	//
	// example:
	//
	// lingjun alinux node group
	NodeGroupDescription *string `json:"NodeGroupDescription,omitempty" xml:"NodeGroupDescription,omitempty"`
	// The ID of the node group.
	//
	// example:
	//
	// i120982301752461697971
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// The name of the node group.
	//
	// example:
	//
	// test-ack
	NodeGroupName *string `json:"NodeGroupName,omitempty" xml:"NodeGroupName,omitempty"`
	// The name of the RAM role. You can call the ListRoles operation of the RAM API to query the RAM roles that you created.
	//
	// example:
	//
	// xianwen-test-ram-role
	RamRoleName *string `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 83896080-59E3-5775-BDDC-8084691C3D96
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The system disk information.
	SystemDisk *DescribeNodeGroupResponseBodySystemDisk `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	// The time when the node group was last updated.
	//
	// example:
	//
	// 2025-08-20T11:18:11.164000
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The user data.
	//
	// example:
	//
	// #!/bin/bash
	//
	// uptime
	//
	// echo "aaaaaaa" >> /tmp/tttest.sh
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// Indicates whether the node group supports GPU virtualization.
	//
	// example:
	//
	// False
	VirtualGpuEnabled *bool `json:"VirtualGpuEnabled,omitempty" xml:"VirtualGpuEnabled,omitempty"`
}

func (s DescribeNodeGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNodeGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNodeGroupResponseBody) GetAz() *string {
	return s.Az
}

func (s *DescribeNodeGroupResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeNodeGroupResponseBody) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeNodeGroupResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeNodeGroupResponseBody) GetFileSystemMountEnabled() *bool {
	return s.FileSystemMountEnabled
}

func (s *DescribeNodeGroupResponseBody) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeNodeGroupResponseBody) GetImageName() *string {
	return s.ImageName
}

func (s *DescribeNodeGroupResponseBody) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *DescribeNodeGroupResponseBody) GetLoginType() *string {
	return s.LoginType
}

func (s *DescribeNodeGroupResponseBody) GetMachineType() *string {
	return s.MachineType
}

func (s *DescribeNodeGroupResponseBody) GetNodeCount() *string {
	return s.NodeCount
}

func (s *DescribeNodeGroupResponseBody) GetNodeGroupDescription() *string {
	return s.NodeGroupDescription
}

func (s *DescribeNodeGroupResponseBody) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *DescribeNodeGroupResponseBody) GetNodeGroupName() *string {
	return s.NodeGroupName
}

func (s *DescribeNodeGroupResponseBody) GetRamRoleName() *string {
	return s.RamRoleName
}

func (s *DescribeNodeGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNodeGroupResponseBody) GetSystemDisk() *DescribeNodeGroupResponseBodySystemDisk {
	return s.SystemDisk
}

func (s *DescribeNodeGroupResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeNodeGroupResponseBody) GetUserData() *string {
	return s.UserData
}

func (s *DescribeNodeGroupResponseBody) GetVirtualGpuEnabled() *bool {
	return s.VirtualGpuEnabled
}

func (s *DescribeNodeGroupResponseBody) SetAz(v string) *DescribeNodeGroupResponseBody {
	s.Az = &v
	return s
}

func (s *DescribeNodeGroupResponseBody) SetClusterId(v string) *DescribeNodeGroupResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeNodeGroupResponseBody) SetClusterName(v string) *DescribeNodeGroupResponseBody {
	s.ClusterName = &v
	return s
}

func (s *DescribeNodeGroupResponseBody) SetCreateTime(v string) *DescribeNodeGroupResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeNodeGroupResponseBody) SetFileSystemMountEnabled(v bool) *DescribeNodeGroupResponseBody {
	s.FileSystemMountEnabled = &v
	return s
}

func (s *DescribeNodeGroupResponseBody) SetImageId(v string) *DescribeNodeGroupResponseBody {
	s.ImageId = &v
	return s
}

func (s *DescribeNodeGroupResponseBody) SetImageName(v string) *DescribeNodeGroupResponseBody {
	s.ImageName = &v
	return s
}

func (s *DescribeNodeGroupResponseBody) SetKeyPairName(v string) *DescribeNodeGroupResponseBody {
	s.KeyPairName = &v
	return s
}

func (s *DescribeNodeGroupResponseBody) SetLoginType(v string) *DescribeNodeGroupResponseBody {
	s.LoginType = &v
	return s
}

func (s *DescribeNodeGroupResponseBody) SetMachineType(v string) *DescribeNodeGroupResponseBody {
	s.MachineType = &v
	return s
}

func (s *DescribeNodeGroupResponseBody) SetNodeCount(v string) *DescribeNodeGroupResponseBody {
	s.NodeCount = &v
	return s
}

func (s *DescribeNodeGroupResponseBody) SetNodeGroupDescription(v string) *DescribeNodeGroupResponseBody {
	s.NodeGroupDescription = &v
	return s
}

func (s *DescribeNodeGroupResponseBody) SetNodeGroupId(v string) *DescribeNodeGroupResponseBody {
	s.NodeGroupId = &v
	return s
}

func (s *DescribeNodeGroupResponseBody) SetNodeGroupName(v string) *DescribeNodeGroupResponseBody {
	s.NodeGroupName = &v
	return s
}

func (s *DescribeNodeGroupResponseBody) SetRamRoleName(v string) *DescribeNodeGroupResponseBody {
	s.RamRoleName = &v
	return s
}

func (s *DescribeNodeGroupResponseBody) SetRequestId(v string) *DescribeNodeGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNodeGroupResponseBody) SetSystemDisk(v *DescribeNodeGroupResponseBodySystemDisk) *DescribeNodeGroupResponseBody {
	s.SystemDisk = v
	return s
}

func (s *DescribeNodeGroupResponseBody) SetUpdateTime(v string) *DescribeNodeGroupResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *DescribeNodeGroupResponseBody) SetUserData(v string) *DescribeNodeGroupResponseBody {
	s.UserData = &v
	return s
}

func (s *DescribeNodeGroupResponseBody) SetVirtualGpuEnabled(v bool) *DescribeNodeGroupResponseBody {
	s.VirtualGpuEnabled = &v
	return s
}

func (s *DescribeNodeGroupResponseBody) Validate() error {
	if s.SystemDisk != nil {
		if err := s.SystemDisk.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeNodeGroupResponseBodySystemDisk struct {
	// The category of the disk.
	//
	// example:
	//
	// system
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The performance level of the ESSD. Valid values:
	//
	// - PL0: A single ESSD can deliver up to 10,000 random read/write IOPS.
	//
	// - PL1: A single ESSD can deliver up to 50,000 random read/write IOPS.
	//
	// example:
	//
	// PL0
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The size of the system disk. Unit: GiB. Valid values:
	//
	// - Basic disk: 20–500.
	//
	// - ESSD:
	//
	//   - PL0: 1–2,048.
	//
	//   - PL1: 20–2,048.
	//
	//   - PL2: 461–2,048.
	//
	//   - PL3: 1,261–2,048.
	//
	// - ESSD AutoPL disk: 1–2,048.
	//
	// - Other disk categories: 20–2,048.
	//
	// Default value: the larger value between 20 and the size of the image that is specified by `ImageId`.
	//
	// example:
	//
	// 1024
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DescribeNodeGroupResponseBodySystemDisk) String() string {
	return dara.Prettify(s)
}

func (s DescribeNodeGroupResponseBodySystemDisk) GoString() string {
	return s.String()
}

func (s *DescribeNodeGroupResponseBodySystemDisk) GetCategory() *string {
	return s.Category
}

func (s *DescribeNodeGroupResponseBodySystemDisk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *DescribeNodeGroupResponseBodySystemDisk) GetSize() *int32 {
	return s.Size
}

func (s *DescribeNodeGroupResponseBodySystemDisk) SetCategory(v string) *DescribeNodeGroupResponseBodySystemDisk {
	s.Category = &v
	return s
}

func (s *DescribeNodeGroupResponseBodySystemDisk) SetPerformanceLevel(v string) *DescribeNodeGroupResponseBodySystemDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *DescribeNodeGroupResponseBodySystemDisk) SetSize(v int32) *DescribeNodeGroupResponseBodySystemDisk {
	s.Size = &v
	return s
}

func (s *DescribeNodeGroupResponseBodySystemDisk) Validate() error {
	return dara.Validate(s)
}
