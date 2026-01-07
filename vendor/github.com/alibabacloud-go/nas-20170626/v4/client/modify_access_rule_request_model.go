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
	// The rule ID.
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
	// 	- standard (default): General-purpose NAS file system.
	//
	// 	- extreme: Extreme NAS file system.
	//
	// example:
	//
	// standard
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	// The IPv6 address or CIDR block of the authorized object.
	//
	// You must set this parameter to an IPv6 IP address or CIDR block.
	//
	// > 	- Only Extreme NAS file systems that reside in the China (Hohhot) region support IPv6.
	//
	// >	- Only permission groups that reside in virtual private clouds (VPCs) support IPv6.
	//
	// >	- This parameter is unavailable if you specify the SourceCidrIp parameter.
	//
	// example:
	//
	// fe80::3d4a:80fd:f05d:****
	Ipv6SourceCidrIp *string `json:"Ipv6SourceCidrIp,omitempty" xml:"Ipv6SourceCidrIp,omitempty"`
	// The priority of the rule.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 1, which indicates the highest priority.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The access permissions of the authorized object on the file system.
	//
	// Valid values:
	//
	// 	- RDWR (default): the read and write permissions.
	//
	// 	- RDONLY: the read-only permissions.
	//
	// example:
	//
	// RDWR
	RWAccessType *string `json:"RWAccessType,omitempty" xml:"RWAccessType,omitempty"`
	// The IP address or CIDR block of the authorized object.
	//
	// You must set this parameter to an IP address or CIDR block.
	//
	// example:
	//
	// ``192.0.**.**``
	SourceCidrIp *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
	// The access permissions for different types of users in the authorized object.
	//
	// Valid values:
	//
	// 	- no_squash: allows access from root users to the file system.
	//
	// 	- root_squash: grants root users the least permissions as the nobody user.
	//
	// 	- all_squash: grants all users the least permissions as the nobody user.
	//
	// The nobody user has the least permissions in Linux and can access only the public content of the file system. This ensures the security of the file system.
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
