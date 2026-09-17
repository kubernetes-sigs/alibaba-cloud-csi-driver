// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNodeGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemMountEnabled(v bool) *UpdateNodeGroupRequest
	GetFileSystemMountEnabled() *bool
	SetImageId(v string) *UpdateNodeGroupRequest
	GetImageId() *string
	SetKeyPairName(v string) *UpdateNodeGroupRequest
	GetKeyPairName() *string
	SetLoginPassword(v string) *UpdateNodeGroupRequest
	GetLoginPassword() *string
	SetNewNodeGroupName(v string) *UpdateNodeGroupRequest
	GetNewNodeGroupName() *string
	SetNodeGroupId(v string) *UpdateNodeGroupRequest
	GetNodeGroupId() *string
	SetRamRoleName(v string) *UpdateNodeGroupRequest
	GetRamRoleName() *string
	SetSystemDisk(v *UpdateNodeGroupRequestSystemDisk) *UpdateNodeGroupRequest
	GetSystemDisk() *UpdateNodeGroupRequestSystemDisk
	SetUserData(v string) *UpdateNodeGroupRequest
	GetUserData() *string
}

type UpdateNodeGroupRequest struct {
	// Specifies whether file storage mounting is supported.
	//
	// example:
	//
	// True
	FileSystemMountEnabled *bool `json:"FileSystemMountEnabled,omitempty" xml:"FileSystemMountEnabled,omitempty"`
	// The default image ID of the node group. If this parameter is not specified, the image remains unchanged.
	//
	// example:
	//
	// i1232142432432
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The key pair name.
	//
	// example:
	//
	// test
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The logon password of the machines in the node group.
	//
	// example:
	//
	// 密码
	LoginPassword *string `json:"LoginPassword,omitempty" xml:"LoginPassword,omitempty"`
	// The node group name.
	//
	// example:
	//
	// test-update
	NewNodeGroupName *string `json:"NewNodeGroupName,omitempty" xml:"NewNodeGroupName,omitempty"`
	// The node group ID.
	//
	// example:
	//
	// i120021051733814190732
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// The RAM role name of the node. You can call the RAM API ListRoles to query the node RAM roles that you have created. The trusted entity of the role must be set to Lingjun AI Computing Service.
	//
	// Note: Clearing an existing role is not supported.
	//
	// example:
	//
	// xianwen-test-ram-role
	RamRoleName *string                           `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	SystemDisk  *UpdateNodeGroupRequestSystemDisk `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	// The custom executable shell script. The script must be Base64-encoded. The maximum size of the raw data is 16 KB.
	//
	// example:
	//
	// ZWNobyBoZWxsbyBlY3Mh
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s UpdateNodeGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateNodeGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateNodeGroupRequest) GetFileSystemMountEnabled() *bool {
	return s.FileSystemMountEnabled
}

func (s *UpdateNodeGroupRequest) GetImageId() *string {
	return s.ImageId
}

func (s *UpdateNodeGroupRequest) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *UpdateNodeGroupRequest) GetLoginPassword() *string {
	return s.LoginPassword
}

func (s *UpdateNodeGroupRequest) GetNewNodeGroupName() *string {
	return s.NewNodeGroupName
}

func (s *UpdateNodeGroupRequest) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *UpdateNodeGroupRequest) GetRamRoleName() *string {
	return s.RamRoleName
}

func (s *UpdateNodeGroupRequest) GetSystemDisk() *UpdateNodeGroupRequestSystemDisk {
	return s.SystemDisk
}

func (s *UpdateNodeGroupRequest) GetUserData() *string {
	return s.UserData
}

func (s *UpdateNodeGroupRequest) SetFileSystemMountEnabled(v bool) *UpdateNodeGroupRequest {
	s.FileSystemMountEnabled = &v
	return s
}

func (s *UpdateNodeGroupRequest) SetImageId(v string) *UpdateNodeGroupRequest {
	s.ImageId = &v
	return s
}

func (s *UpdateNodeGroupRequest) SetKeyPairName(v string) *UpdateNodeGroupRequest {
	s.KeyPairName = &v
	return s
}

func (s *UpdateNodeGroupRequest) SetLoginPassword(v string) *UpdateNodeGroupRequest {
	s.LoginPassword = &v
	return s
}

func (s *UpdateNodeGroupRequest) SetNewNodeGroupName(v string) *UpdateNodeGroupRequest {
	s.NewNodeGroupName = &v
	return s
}

func (s *UpdateNodeGroupRequest) SetNodeGroupId(v string) *UpdateNodeGroupRequest {
	s.NodeGroupId = &v
	return s
}

func (s *UpdateNodeGroupRequest) SetRamRoleName(v string) *UpdateNodeGroupRequest {
	s.RamRoleName = &v
	return s
}

func (s *UpdateNodeGroupRequest) SetSystemDisk(v *UpdateNodeGroupRequestSystemDisk) *UpdateNodeGroupRequest {
	s.SystemDisk = v
	return s
}

func (s *UpdateNodeGroupRequest) SetUserData(v string) *UpdateNodeGroupRequest {
	s.UserData = &v
	return s
}

func (s *UpdateNodeGroupRequest) Validate() error {
	if s.SystemDisk != nil {
		if err := s.SystemDisk.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateNodeGroupRequestSystemDisk struct {
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
}

func (s UpdateNodeGroupRequestSystemDisk) String() string {
	return dara.Prettify(s)
}

func (s UpdateNodeGroupRequestSystemDisk) GoString() string {
	return s.String()
}

func (s *UpdateNodeGroupRequestSystemDisk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *UpdateNodeGroupRequestSystemDisk) SetPerformanceLevel(v string) *UpdateNodeGroupRequestSystemDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *UpdateNodeGroupRequestSystemDisk) Validate() error {
	return dara.Validate(s)
}
