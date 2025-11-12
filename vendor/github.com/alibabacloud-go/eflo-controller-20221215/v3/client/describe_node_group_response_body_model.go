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
	Az                     *string `json:"Az,omitempty" xml:"Az,omitempty"`
	ClusterId              *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterName            *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	CreateTime             *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	FileSystemMountEnabled *bool   `json:"FileSystemMountEnabled,omitempty" xml:"FileSystemMountEnabled,omitempty"`
	ImageId                *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageName              *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	KeyPairName            *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	LoginType              *string `json:"LoginType,omitempty" xml:"LoginType,omitempty"`
	MachineType            *string `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	NodeCount              *string `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	NodeGroupDescription   *string `json:"NodeGroupDescription,omitempty" xml:"NodeGroupDescription,omitempty"`
	NodeGroupId            *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	NodeGroupName          *string `json:"NodeGroupName,omitempty" xml:"NodeGroupName,omitempty"`
	// Id of the request
	RequestId         *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SystemDisk        *DescribeNodeGroupResponseBodySystemDisk `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	UpdateTime        *string                                  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserData          *string                                  `json:"UserData,omitempty" xml:"UserData,omitempty"`
	VirtualGpuEnabled *bool                                    `json:"VirtualGpuEnabled,omitempty" xml:"VirtualGpuEnabled,omitempty"`
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
	Category         *string `json:"Category,omitempty" xml:"Category,omitempty"`
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	Size             *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
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
