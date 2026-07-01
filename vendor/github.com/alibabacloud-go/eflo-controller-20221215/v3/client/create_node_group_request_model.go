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
	// The ID of the cluster to which the node group belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// i118191731740041623425
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The configurations of the node group.
	//
	// This parameter is required.
	NodeGroup *CreateNodeGroupRequestNodeGroup `json:"NodeGroup,omitempty" xml:"NodeGroup,omitempty" type:"Struct"`
	// The configuration of the node unit.
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
	if s.NodeGroup != nil {
		if err := s.NodeGroup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateNodeGroupRequestNodeGroup struct {
	// The availability zone of the node group.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu-b
	Az *string `json:"Az,omitempty" xml:"Az,omitempty"`
	// Specifies whether to enable file system mounting.
	//
	// example:
	//
	// false
	FileSystemMountEnabled *bool `json:"FileSystemMountEnabled,omitempty" xml:"FileSystemMountEnabled,omitempty"`
	// The image ID for the nodes.
	//
	// This parameter is required.
	//
	// example:
	//
	// i191887641687336652616
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The name of the key pair for SSH login.
	//
	// example:
	//
	// test-keypair
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The password to log in to the nodes.
	//
	// example:
	//
	// test-LoginPassword
	LoginPassword *string `json:"LoginPassword,omitempty" xml:"LoginPassword,omitempty"`
	// The machine type for the nodes.
	//
	// This parameter is required.
	//
	// example:
	//
	// mock-machine-type3
	MachineType *string `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	// The description of the node group.
	//
	// example:
	//
	// Node group description
	NodeGroupDescription *string `json:"NodeGroupDescription,omitempty" xml:"NodeGroupDescription,omitempty"`
	// The name of the node group.
	//
	// This parameter is required.
	//
	// example:
	//
	// PAI-LINGJUN
	NodeGroupName *string `json:"NodeGroupName,omitempty" xml:"NodeGroupName,omitempty"`
	// The name of the RAM role to attach to the nodes. You can call the RAM API `ListRoles` operation to query the RAM roles that you have created. The trust entity of the specified role must be Intelligent Computing Lingjun.<br>**Note:*	- You cannot detach an existing role by clearing this parameter.<br>
	//
	// example:
	//
	// xianwen-test-ram-role
	RamRoleName *string `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	// The system disk configuration for the nodes.
	SystemDisk *CreateNodeGroupRequestNodeGroupSystemDisk `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	// The user data passed to the nodes at launch.
	//
	// example:
	//
	// #!/bin/bash
	//
	// uptime
	//
	// echo "aaaaaaa" >> /tmp/ttttt20250110141010.sh
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// Specifies whether to enable GPU virtualization.
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

func (s *CreateNodeGroupRequestNodeGroup) GetRamRoleName() *string {
	return s.RamRoleName
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

func (s *CreateNodeGroupRequestNodeGroup) SetRamRoleName(v string) *CreateNodeGroupRequestNodeGroup {
	s.RamRoleName = &v
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
	if s.SystemDisk != nil {
		if err := s.SystemDisk.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateNodeGroupRequestNodeGroupSystemDisk struct {
	// The type of the system disk. Valid values:
	//
	// - `cloud_essd`: ESSD.
	//
	// example:
	//
	// clou_essd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The performance level of the ESSD system disk. Valid values:
	//
	// - `PL0`: A single disk delivers up to 10,000 random read/write IOPS.
	//
	// - `PL1`: A single disk delivers up to 50,000 random read/write IOPS.
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The size of the system disk, in GB.
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
