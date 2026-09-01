// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccessRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessGroupName(v string) *ModifyAccessRuleRequest
	GetAccessGroupName() *string
	SetAccessRuleId(v string) *ModifyAccessRuleRequest
	GetAccessRuleId() *string
	SetFileSystemType(v string) *ModifyAccessRuleRequest
	GetFileSystemType() *string
	SetIpv6SourceCidrIp(v string) *ModifyAccessRuleRequest
	GetIpv6SourceCidrIp() *string
	SetPriority(v int32) *ModifyAccessRuleRequest
	GetPriority() *int32
	SetRWAccessType(v string) *ModifyAccessRuleRequest
	GetRWAccessType() *string
	SetSourceCidrIp(v string) *ModifyAccessRuleRequest
	GetSourceCidrIp() *string
	SetUserAccessType(v string) *ModifyAccessRuleRequest
	GetUserAccessType() *string
}

type ModifyAccessRuleRequest struct {
	// The name of the permission group.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-test
	AccessGroupName *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	// The ID of the permission rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	AccessRuleId *string `json:"AccessRuleId,omitempty" xml:"AccessRuleId,omitempty"`
	// The type of the file system.
	//
	// Valid values:
	//
	// - standard (default): General-purpose NAS.
	//
	// - extreme: Extreme NAS.
	//
	// example:
	//
	// standard
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	// The source IPv6 CIDR block.
	//
	// IPv6 addresses and CIDR blocks are supported.
	//
	// > - Only Extreme NAS file systems in the China (Hohhot) region support IPv6 CIDR blocks.
	//
	// > - Only VPCs are supported.
	//
	// > - IPv4 and IPv6 are mutually exclusive. You cannot convert between the two types.
	//
	// > - You must specify either SourceCidrIp or Ipv6SourceCidrIp. You cannot leave both parameters empty, and you cannot specify both parameters at the same time.
	//
	// example:
	//
	// fe80::3d4a:80fd:f05d:****
	Ipv6SourceCidrIp *string `json:"Ipv6SourceCidrIp,omitempty" xml:"Ipv6SourceCidrIp,omitempty"`
	// The priority of the permission rule.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 1 (highest priority).
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The read and write permission that the authorized object has on the file system.
	//
	// Valid values:
	//
	// - RDWR (default): read and write.
	//
	// - RDONLY: read-only.
	//
	// example:
	//
	// RDWR
	RWAccessType *string `json:"RWAccessType,omitempty" xml:"RWAccessType,omitempty"`
	// The IP address or CIDR block.
	//
	// The value must be a single IP address or a CIDR block.
	//
	// > You must specify either SourceCidrIp or Ipv6SourceCidrIp. You cannot leave both parameters empty, and you cannot specify both parameters at the same time.
	//
	// example:
	//
	// ``192.0.**.**``
	SourceCidrIp *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
	// The access permissions that the system user of the authorization object has on the file system.
	//
	// Valid values:
	//
	// - no_squash: allows access to the file system as the root user.
	//
	// - root_squash: maps the root user to the nobody user when the root user accesses the file system.
	//
	// - all_squash: maps all users to the nobody user regardless of the user identity used to access the file system.
	//
	// The nobody user is a default user in Linux. This user can access only public content on the server and has low privileges and high security.
	//
	// example:
	//
	// all_squash
	UserAccessType *string `json:"UserAccessType,omitempty" xml:"UserAccessType,omitempty"`
}

func (s ModifyAccessRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccessRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccessRuleRequest) GetAccessGroupName() *string {
	return s.AccessGroupName
}

func (s *ModifyAccessRuleRequest) GetAccessRuleId() *string {
	return s.AccessRuleId
}

func (s *ModifyAccessRuleRequest) GetFileSystemType() *string {
	return s.FileSystemType
}

func (s *ModifyAccessRuleRequest) GetIpv6SourceCidrIp() *string {
	return s.Ipv6SourceCidrIp
}

func (s *ModifyAccessRuleRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *ModifyAccessRuleRequest) GetRWAccessType() *string {
	return s.RWAccessType
}

func (s *ModifyAccessRuleRequest) GetSourceCidrIp() *string {
	return s.SourceCidrIp
}

func (s *ModifyAccessRuleRequest) GetUserAccessType() *string {
	return s.UserAccessType
}

func (s *ModifyAccessRuleRequest) SetAccessGroupName(v string) *ModifyAccessRuleRequest {
	s.AccessGroupName = &v
	return s
}

func (s *ModifyAccessRuleRequest) SetAccessRuleId(v string) *ModifyAccessRuleRequest {
	s.AccessRuleId = &v
	return s
}

func (s *ModifyAccessRuleRequest) SetFileSystemType(v string) *ModifyAccessRuleRequest {
	s.FileSystemType = &v
	return s
}

func (s *ModifyAccessRuleRequest) SetIpv6SourceCidrIp(v string) *ModifyAccessRuleRequest {
	s.Ipv6SourceCidrIp = &v
	return s
}

func (s *ModifyAccessRuleRequest) SetPriority(v int32) *ModifyAccessRuleRequest {
	s.Priority = &v
	return s
}

func (s *ModifyAccessRuleRequest) SetRWAccessType(v string) *ModifyAccessRuleRequest {
	s.RWAccessType = &v
	return s
}

func (s *ModifyAccessRuleRequest) SetSourceCidrIp(v string) *ModifyAccessRuleRequest {
	s.SourceCidrIp = &v
	return s
}

func (s *ModifyAccessRuleRequest) SetUserAccessType(v string) *ModifyAccessRuleRequest {
	s.UserAccessType = &v
	return s
}

func (s *ModifyAccessRuleRequest) Validate() error {
	return dara.Validate(s)
}
