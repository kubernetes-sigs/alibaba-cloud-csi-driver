// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccessPointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessGroup(v string) *CreateAccessPointRequest
	GetAccessGroup() *string
	SetAccessPointName(v string) *CreateAccessPointRequest
	GetAccessPointName() *string
	SetEnabledRam(v bool) *CreateAccessPointRequest
	GetEnabledRam() *bool
	SetFileSystemId(v string) *CreateAccessPointRequest
	GetFileSystemId() *string
	SetOwnerGroupId(v int32) *CreateAccessPointRequest
	GetOwnerGroupId() *int32
	SetOwnerUserId(v int32) *CreateAccessPointRequest
	GetOwnerUserId() *int32
	SetPermission(v string) *CreateAccessPointRequest
	GetPermission() *string
	SetPosixGroupId(v int32) *CreateAccessPointRequest
	GetPosixGroupId() *int32
	SetPosixSecondaryGroupIds(v string) *CreateAccessPointRequest
	GetPosixSecondaryGroupIds() *string
	SetPosixUserId(v int32) *CreateAccessPointRequest
	GetPosixUserId() *int32
	SetRootDirectory(v string) *CreateAccessPointRequest
	GetRootDirectory() *string
	SetTag(v []*CreateAccessPointRequestTag) *CreateAccessPointRequest
	GetTag() []*CreateAccessPointRequestTag
	SetVpcId(v string) *CreateAccessPointRequest
	GetVpcId() *string
	SetVswId(v string) *CreateAccessPointRequest
	GetVswId() *string
}

type CreateAccessPointRequest struct {
	// The name of the permission group.
	//
	// This parameter is required for a General-purpose File Storage NAS (NAS) file system.
	//
	// The default permission group for virtual private clouds (VPCs) is named DEFAULT_VPC_GROUP_NAME.
	//
	// This parameter is required.
	//
	// example:
	//
	// DEFAULT_VPC_GROUP_NAME
	AccessGroup *string `json:"AccessGroup,omitempty" xml:"AccessGroup,omitempty"`
	// The name of the access point.
	//
	// example:
	//
	// test
	AccessPointName *string `json:"AccessPointName,omitempty" xml:"AccessPointName,omitempty"`
	// Specifies whether to enable the RAM policy. Valid values:
	//
	// 	- true: The RAM policy is enabled.
	//
	// 	- false (default): The RAM policy is disabled.
	//
	// >  After the RAM policy is enabled for access points, no RAM user is allowed to use access points to mount and access data by default. To use access points to mount and access data as a RAM user, you must grant the related access permissions to the RAM user. If the RAM policy is disabled, access points can be anonymously mounted. For more information about how to configure permissions on access points, see [Configure a policy for the access point](https://help.aliyun.com/document_detail/2545998.html).
	//
	// example:
	//
	// false
	EnabledRam *bool `json:"EnabledRam,omitempty" xml:"EnabledRam,omitempty"`
	// The ID of the file system.
	//
	// This parameter is required.
	//
	// example:
	//
	// 31a8e4****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The ID of the owner group.
	//
	// This parameter is required if the RootDirectory directory does not exist.
	//
	// example:
	//
	// 1
	OwnerGroupId *int32 `json:"OwnerGroupId,omitempty" xml:"OwnerGroupId,omitempty"`
	// The owner ID.
	//
	// This parameter is required if the RootDirectory directory does not exist.
	//
	// example:
	//
	// 1
	OwnerUserId *int32 `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	// The Portable Operating System Interface for UNIX (POSIX) permission. Default value: 0777.
	//
	// This field takes effect only if you specify the OwnerUserId and OwnerGroupId parameters.
	//
	// example:
	//
	// 0777
	Permission *string `json:"Permission,omitempty" xml:"Permission,omitempty"`
	// The ID of the POSIX user group.
	//
	// example:
	//
	// 123
	PosixGroupId *int32 `json:"PosixGroupId,omitempty" xml:"PosixGroupId,omitempty"`
	// The secondary user group. Separate multiple user group IDs with commas (,).
	//
	// example:
	//
	// 123,345
	PosixSecondaryGroupIds *string `json:"PosixSecondaryGroupIds,omitempty" xml:"PosixSecondaryGroupIds,omitempty"`
	// The ID of the POSIX user.
	//
	// example:
	//
	// 123
	PosixUserId *int32 `json:"PosixUserId,omitempty" xml:"PosixUserId,omitempty"`
	// The root directory of the access point. The default value is /. If the directory does not exist, you must also specify the OwnerUserId and OwnerGroupId parameters.
	//
	// example:
	//
	// /
	RootDirectory *string                        `json:"RootDirectory,omitempty" xml:"RootDirectory,omitempty"`
	Tag           []*CreateAccessPointRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The VPC ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-2zesj9afh3y518k9o****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The vSwitch ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-2zevmwkwyztjuoffg****
	VswId *string `json:"VswId,omitempty" xml:"VswId,omitempty"`
}

func (s CreateAccessPointRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAccessPointRequest) GoString() string {
	return s.String()
}

func (s *CreateAccessPointRequest) GetAccessGroup() *string {
	return s.AccessGroup
}

func (s *CreateAccessPointRequest) GetAccessPointName() *string {
	return s.AccessPointName
}

func (s *CreateAccessPointRequest) GetEnabledRam() *bool {
	return s.EnabledRam
}

func (s *CreateAccessPointRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *CreateAccessPointRequest) GetOwnerGroupId() *int32 {
	return s.OwnerGroupId
}

func (s *CreateAccessPointRequest) GetOwnerUserId() *int32 {
	return s.OwnerUserId
}

func (s *CreateAccessPointRequest) GetPermission() *string {
	return s.Permission
}

func (s *CreateAccessPointRequest) GetPosixGroupId() *int32 {
	return s.PosixGroupId
}

func (s *CreateAccessPointRequest) GetPosixSecondaryGroupIds() *string {
	return s.PosixSecondaryGroupIds
}

func (s *CreateAccessPointRequest) GetPosixUserId() *int32 {
	return s.PosixUserId
}

func (s *CreateAccessPointRequest) GetRootDirectory() *string {
	return s.RootDirectory
}

func (s *CreateAccessPointRequest) GetTag() []*CreateAccessPointRequestTag {
	return s.Tag
}

func (s *CreateAccessPointRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateAccessPointRequest) GetVswId() *string {
	return s.VswId
}

func (s *CreateAccessPointRequest) SetAccessGroup(v string) *CreateAccessPointRequest {
	s.AccessGroup = &v
	return s
}

func (s *CreateAccessPointRequest) SetAccessPointName(v string) *CreateAccessPointRequest {
	s.AccessPointName = &v
	return s
}

func (s *CreateAccessPointRequest) SetEnabledRam(v bool) *CreateAccessPointRequest {
	s.EnabledRam = &v
	return s
}

func (s *CreateAccessPointRequest) SetFileSystemId(v string) *CreateAccessPointRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateAccessPointRequest) SetOwnerGroupId(v int32) *CreateAccessPointRequest {
	s.OwnerGroupId = &v
	return s
}

func (s *CreateAccessPointRequest) SetOwnerUserId(v int32) *CreateAccessPointRequest {
	s.OwnerUserId = &v
	return s
}

func (s *CreateAccessPointRequest) SetPermission(v string) *CreateAccessPointRequest {
	s.Permission = &v
	return s
}

func (s *CreateAccessPointRequest) SetPosixGroupId(v int32) *CreateAccessPointRequest {
	s.PosixGroupId = &v
	return s
}

func (s *CreateAccessPointRequest) SetPosixSecondaryGroupIds(v string) *CreateAccessPointRequest {
	s.PosixSecondaryGroupIds = &v
	return s
}

func (s *CreateAccessPointRequest) SetPosixUserId(v int32) *CreateAccessPointRequest {
	s.PosixUserId = &v
	return s
}

func (s *CreateAccessPointRequest) SetRootDirectory(v string) *CreateAccessPointRequest {
	s.RootDirectory = &v
	return s
}

func (s *CreateAccessPointRequest) SetTag(v []*CreateAccessPointRequestTag) *CreateAccessPointRequest {
	s.Tag = v
	return s
}

func (s *CreateAccessPointRequest) SetVpcId(v string) *CreateAccessPointRequest {
	s.VpcId = &v
	return s
}

func (s *CreateAccessPointRequest) SetVswId(v string) *CreateAccessPointRequest {
	s.VswId = &v
	return s
}

func (s *CreateAccessPointRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateAccessPointRequestTag struct {
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateAccessPointRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateAccessPointRequestTag) GoString() string {
	return s.String()
}

func (s *CreateAccessPointRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateAccessPointRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateAccessPointRequestTag) SetKey(v string) *CreateAccessPointRequestTag {
	s.Key = &v
	return s
}

func (s *CreateAccessPointRequestTag) SetValue(v string) *CreateAccessPointRequestTag {
	s.Value = &v
	return s
}

func (s *CreateAccessPointRequestTag) Validate() error {
	return dara.Validate(s)
}
