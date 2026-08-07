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
	SetAgenticSpaceId(v string) *CreateAccessPointRequest
	GetAgenticSpaceId() *string
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
	// This parameter is required if the file system is a General-purpose NAS file system.
	//
	// Default permission group: DEFAULT_VPC_GROUP_NAME (the default permission group for VPCs).
	//
	// >Not supported for Agentic file systems.
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
	// The AgenticSpace ID.
	//
	// >This parameter is required for Agentic file systems.
	//
	// example:
	//
	// agentic-229oypxjgpau2****
	AgenticSpaceId *string `json:"AgenticSpaceId,omitempty" xml:"AgenticSpaceId,omitempty"`
	// Specifies whether to enable access point policy.
	//
	// Valid values:
	//
	// - true: enabled.
	//
	// - false (default): not enabled.
	//
	// > After you enable access point policy for the access point, all Resource Access Management (RAM) users are denied access to mount and access data through the access point by default. You must grant the corresponding access permissions through authorization and then mount and access the file system through the access point. After you disable access point policy, the access point allows anonymity mounting. For more information about how to configure access point permissions, see [Configure access point policies](https://help.aliyun.com/document_detail/2545998.html).
	//
	// >For Agentic file systems, this parameter must be set to true.
	//
	// example:
	//
	// false
	EnabledRam *bool `json:"EnabledRam,omitempty" xml:"EnabledRam,omitempty"`
	// The file system ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 31a8e4****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The owner group ID.
	//
	// This parameter is required if the RootDirectory directory does not exist.
	//
	// >Not supported for Agentic file systems.
	//
	// example:
	//
	// 1
	OwnerGroupId *int32 `json:"OwnerGroupId,omitempty" xml:"OwnerGroupId,omitempty"`
	// The owner user ID.
	//
	// This parameter is required if the RootDirectory directory does not exist.
	//
	// >Not supported for Agentic file systems.
	//
	// example:
	//
	// 1
	OwnerUserId *int32 `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	// The POSIX permission. Default value: "0755". The value must be a four-digit octal number that starts with 0.
	//
	// This parameter takes effect only after you specify the OwnerUserId and OwnerGroupId parameters.
	//
	// >Not supported for Agentic file systems.
	//
	// example:
	//
	// 0755
	Permission *string `json:"Permission,omitempty" xml:"Permission,omitempty"`
	// The POSIX group ID.
	//
	// >Not supported for Agentic file systems.
	//
	// example:
	//
	// 123
	PosixGroupId *int32 `json:"PosixGroupId,omitempty" xml:"PosixGroupId,omitempty"`
	// The secondary group IDs. Separate multiple group IDs with commas (,).
	//
	// >Not supported for Agentic file systems.
	//
	// example:
	//
	// 123,345
	PosixSecondaryGroupIds *string `json:"PosixSecondaryGroupIds,omitempty" xml:"PosixSecondaryGroupIds,omitempty"`
	// The POSIX user ID.
	//
	// >Not supported for Agentic file systems.
	//
	// example:
	//
	// 123
	PosixUserId *int32 `json:"PosixUserId,omitempty" xml:"PosixUserId,omitempty"`
	// The root directory of the access point.
	//
	// Default value: "/". If the access point directory does not exist, you must also specify the OwnerUserId and OwnerGroupId parameters.
	//
	// >Supported only for Agentic file systems.
	//
	// example:
	//
	// /
	RootDirectory *string `json:"RootDirectory,omitempty" xml:"RootDirectory,omitempty"`
	// The list of access point tags.
	Tag []*CreateAccessPointRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The virtual private cloud (VPC) ID.
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

func (s *CreateAccessPointRequest) GetAgenticSpaceId() *string {
	return s.AgenticSpaceId
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

func (s *CreateAccessPointRequest) SetAgenticSpaceId(v string) *CreateAccessPointRequest {
	s.AgenticSpaceId = &v
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
	// The tag key.
	//
	// Limits:
	//
	// - Cannot be empty or an empty string.
	//
	// - Can be up to 128 characters in length.
	//
	// - Cannot start with aliyun or acs:.
	//
	// - Cannot contain http:// or https://.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// Limits:
	//
	// - Cannot be empty or an empty string.
	//
	// - Can be up to 128 characters in length.
	//
	// - Cannot contain http:// or https://.
	//
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
