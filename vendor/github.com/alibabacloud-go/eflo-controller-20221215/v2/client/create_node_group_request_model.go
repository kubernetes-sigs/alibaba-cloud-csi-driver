// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNodeGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CreateNodeGroupRequest
	GetClusterId() *string
	SetNodeGroup(v *CreateNodeGroupRequestNodeGroup) *CreateNodeGroupRequest
	GetNodeGroup() *CreateNodeGroupRequestNodeGroup
	SetNodeUnit(v map[string]interface{}) *CreateNodeGroupRequest
	GetNodeUnit() map[string]interface{}
}

type CreateNodeGroupRequest struct {
	// Cluster ID
	//
	// This parameter is required.
	//
	// example:
	//
	// i118191731740041623425
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Node ID.
	//
	// This parameter is required.
	NodeGroup *CreateNodeGroupRequestNodeGroup `json:"NodeGroup,omitempty" xml:"NodeGroup,omitempty" type:"Struct"`
	// Node information
	//
	// example:
	//
	// {\\"NodeUnitId\\":\\"3c2999a8-2b95-4409-93c5-ad3985fc5c9f\\",\\"ResourceGroupId\\":\\"\\",\\"MaxNodes\\":0,\\"NodeUnitName\\":\\"asi_cn-serverless-sale_e01-lingjun-psale\\"}
	NodeUnit map[string]interface{} `json:"NodeUnit,omitempty" xml:"NodeUnit,omitempty"`
}

func (s CreateNodeGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNodeGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateNodeGroupRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateNodeGroupRequest) GetNodeGroup() *CreateNodeGroupRequestNodeGroup {
	return s.NodeGroup
}

func (s *CreateNodeGroupRequest) GetNodeUnit() map[string]interface{} {
	return s.NodeUnit
}

func (s *CreateNodeGroupRequest) SetClusterId(v string) *CreateNodeGroupRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateNodeGroupRequest) SetNodeGroup(v *CreateNodeGroupRequestNodeGroup) *CreateNodeGroupRequest {
	s.NodeGroup = v
	return s
}

func (s *CreateNodeGroupRequest) SetNodeUnit(v map[string]interface{}) *CreateNodeGroupRequest {
	s.NodeUnit = v
	return s
}

func (s *CreateNodeGroupRequest) Validate() error {
	return dara.Validate(s)
}

type CreateNodeGroupRequestNodeGroup struct {
	// Availability Zone
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu-b
	Az *string `json:"Az,omitempty" xml:"Az,omitempty"`
	// Whether file storage mounting is supported
	//
	// example:
	//
	// true
	FileSystemMountEnabled *bool `json:"FileSystemMountEnabled,omitempty" xml:"FileSystemMountEnabled,omitempty"`
	// Image ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// i191887641687336652616
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// Key pair name.
	//
	// example:
	//
	// test-keypair
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// Password
	//
	// example:
	//
	// test-LoginPassword
	LoginPassword *string `json:"LoginPassword,omitempty" xml:"LoginPassword,omitempty"`
	// Machine type
	//
	// This parameter is required.
	//
	// example:
	//
	// mock-machine-type3
	MachineType *string `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	// Node group description
	//
	// example:
	//
	// describe for node group
	NodeGroupDescription *string `json:"NodeGroupDescription,omitempty" xml:"NodeGroupDescription,omitempty"`
	// Node group name
	//
	// This parameter is required.
	//
	// example:
	//
	// PAI-LINGJUN
	NodeGroupName *string `json:"NodeGroupName,omitempty" xml:"NodeGroupName,omitempty"`
	// Details of the node system disk configuration.
	SystemDisk *CreateNodeGroupRequestNodeGroupSystemDisk `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	// User-defined data
	//
	// example:
	//
	// #!/bin/bash
	//
	// uptime
	//
	// echo "aaaaaaa" >> /tmp/ttttt20250110141010.sh
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// Whether to enable gpu virtualization or not
	//
	// example:
	//
	// false
	VirtualGpuEnabled *bool `json:"VirtualGpuEnabled,omitempty" xml:"VirtualGpuEnabled,omitempty"`
}

func (s CreateNodeGroupRequestNodeGroup) String() string {
	return dara.Prettify(s)
}

func (s CreateNodeGroupRequestNodeGroup) GoString() string {
	return s.String()
}

func (s *CreateNodeGroupRequestNodeGroup) GetAz() *string {
	return s.Az
}

func (s *CreateNodeGroupRequestNodeGroup) GetFileSystemMountEnabled() *bool {
	return s.FileSystemMountEnabled
}

func (s *CreateNodeGroupRequestNodeGroup) GetImageId() *string {
	return s.ImageId
}

func (s *CreateNodeGroupRequestNodeGroup) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *CreateNodeGroupRequestNodeGroup) GetLoginPassword() *string {
	return s.LoginPassword
}

func (s *CreateNodeGroupRequestNodeGroup) GetMachineType() *string {
	return s.MachineType
}

func (s *CreateNodeGroupRequestNodeGroup) GetNodeGroupDescription() *string {
	return s.NodeGroupDescription
}

func (s *CreateNodeGroupRequestNodeGroup) GetNodeGroupName() *string {
	return s.NodeGroupName
}

func (s *CreateNodeGroupRequestNodeGroup) GetSystemDisk() *CreateNodeGroupRequestNodeGroupSystemDisk {
	return s.SystemDisk
}

func (s *CreateNodeGroupRequestNodeGroup) GetUserData() *string {
	return s.UserData
}

func (s *CreateNodeGroupRequestNodeGroup) GetVirtualGpuEnabled() *bool {
	return s.VirtualGpuEnabled
}

func (s *CreateNodeGroupRequestNodeGroup) SetAz(v string) *CreateNodeGroupRequestNodeGroup {
	s.Az = &v
	return s
}

func (s *CreateNodeGroupRequestNodeGroup) SetFileSystemMountEnabled(v bool) *CreateNodeGroupRequestNodeGroup {
	s.FileSystemMountEnabled = &v
	return s
}

func (s *CreateNodeGroupRequestNodeGroup) SetImageId(v string) *CreateNodeGroupRequestNodeGroup {
	s.ImageId = &v
	return s
}

func (s *CreateNodeGroupRequestNodeGroup) SetKeyPairName(v string) *CreateNodeGroupRequestNodeGroup {
	s.KeyPairName = &v
	return s
}

func (s *CreateNodeGroupRequestNodeGroup) SetLoginPassword(v string) *CreateNodeGroupRequestNodeGroup {
	s.LoginPassword = &v
	return s
}

func (s *CreateNodeGroupRequestNodeGroup) SetMachineType(v string) *CreateNodeGroupRequestNodeGroup {
	s.MachineType = &v
	return s
}

func (s *CreateNodeGroupRequestNodeGroup) SetNodeGroupDescription(v string) *CreateNodeGroupRequestNodeGroup {
	s.NodeGroupDescription = &v
	return s
}

func (s *CreateNodeGroupRequestNodeGroup) SetNodeGroupName(v string) *CreateNodeGroupRequestNodeGroup {
	s.NodeGroupName = &v
	return s
}

func (s *CreateNodeGroupRequestNodeGroup) SetSystemDisk(v *CreateNodeGroupRequestNodeGroupSystemDisk) *CreateNodeGroupRequestNodeGroup {
	s.SystemDisk = v
	return s
}

func (s *CreateNodeGroupRequestNodeGroup) SetUserData(v string) *CreateNodeGroupRequestNodeGroup {
	s.UserData = &v
	return s
}

func (s *CreateNodeGroupRequestNodeGroup) SetVirtualGpuEnabled(v bool) *CreateNodeGroupRequestNodeGroup {
	s.VirtualGpuEnabled = &v
	return s
}

func (s *CreateNodeGroupRequestNodeGroup) Validate() error {
	return dara.Validate(s)
}

type CreateNodeGroupRequestNodeGroupSystemDisk struct {
	// Disk type. Value range:
	//
	//  - cloud_essd: ESSD cloud disk.
	//
	// example:
	//
	// clou_essd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// When creating an ESSD cloud disk as a system disk, set the performance level of the cloud disk. Value range:
	//
	// - PL0: Maximum random read/write IOPS per disk 10,000.
	//
	// - PL1: Maximum random read/write IOPS per disk 50,000.
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// Unit: GB.
	//
	// example:
	//
	// 1000
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s CreateNodeGroupRequestNodeGroupSystemDisk) String() string {
	return dara.Prettify(s)
}

func (s CreateNodeGroupRequestNodeGroupSystemDisk) GoString() string {
	return s.String()
}

func (s *CreateNodeGroupRequestNodeGroupSystemDisk) GetCategory() *string {
	return s.Category
}

func (s *CreateNodeGroupRequestNodeGroupSystemDisk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *CreateNodeGroupRequestNodeGroupSystemDisk) GetSize() *int32 {
	return s.Size
}

func (s *CreateNodeGroupRequestNodeGroupSystemDisk) SetCategory(v string) *CreateNodeGroupRequestNodeGroupSystemDisk {
	s.Category = &v
	return s
}

func (s *CreateNodeGroupRequestNodeGroupSystemDisk) SetPerformanceLevel(v string) *CreateNodeGroupRequestNodeGroupSystemDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *CreateNodeGroupRequestNodeGroupSystemDisk) SetSize(v int32) *CreateNodeGroupRequestNodeGroupSystemDisk {
	s.Size = &v
	return s
}

func (s *CreateNodeGroupRequestNodeGroupSystemDisk) Validate() error {
	return dara.Validate(s)
}
