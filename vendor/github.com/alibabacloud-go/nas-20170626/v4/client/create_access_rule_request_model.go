// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccessRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessGroupName(v string) *CreateAccessRuleRequest
	GetAccessGroupName() *string
	SetFileSystemType(v string) *CreateAccessRuleRequest
	GetFileSystemType() *string
	SetIpv6SourceCidrIp(v string) *CreateAccessRuleRequest
	GetIpv6SourceCidrIp() *string
	SetPriority(v int32) *CreateAccessRuleRequest
	GetPriority() *int32
	SetRWAccessType(v string) *CreateAccessRuleRequest
	GetRWAccessType() *string
	SetSourceCidrIp(v string) *CreateAccessRuleRequest
	GetSourceCidrIp() *string
	SetUserAccessType(v string) *CreateAccessRuleRequest
	GetUserAccessType() *string
}

type CreateAccessRuleRequest struct {
	// The name of the permission group.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-test
	AccessGroupName *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	// The type of the file system.
	//
	// Valid values:
	//
	// 	- standard (default): General-purpose NAS file system
	//
	// 	- extreme: Extreme NAS file system
	//
	// example:
	//
	// standard
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	// The IPv6 address or CIDR block of the authorized object.
	//
	// You must set this parameter to an IPv6 address or CIDR block.
	//
	// > 	- Only Extreme NAS file systems that reside in the Chinese mainland support IPv6. If you specify this parameter, you must enable IPv6 for the file system.
	//
	// >	- Only permission groups that reside in virtual private clouds (VPCs) support IPv6.
	//
	// >	- You cannot specify an IPv4 address and an IPv6 address at the same time.
	//
	// example:
	//
	// 2001:250:6000::***
	Ipv6SourceCidrIp *string `json:"Ipv6SourceCidrIp,omitempty" xml:"Ipv6SourceCidrIp,omitempty"`
	// The priority of the rule.
	//
	// The rule with the highest priority takes effect if multiple rules are attached to the authorized object.
	//
	// Valid values: 1 to 100. The value 1 indicates the highest priority.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The access permissions of the authorized object on the file system.
	//
	// Valid values:
	//
	// 	- RDWR (default): the read and write permissions
	//
	// 	- RDONLY: the read-only permissions
	//
	// example:
	//
	// RDWR
	RWAccessType *string `json:"RWAccessType,omitempty" xml:"RWAccessType,omitempty"`
	// The IP address or CIDR block of the authorized object.
	//
	// You must set this parameter to an IP address or CIDR block.
	//
	// > If the permission group resides in the classic network, you must set this parameter to an IP address.
	//
	// example:
	//
	// 192.0.2.0/16
	SourceCidrIp *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
	// The access permissions for different types of users in the authorized object.
	//
	// Valid values:
	//
	// 	- no_squash (default): grants root users the permissions to access the file system.
	//
	// 	- root_squash: grants root users the least permissions as the nobody user.
	//
	// 	- all_squash: grants all users the least permissions as the nobody user.
	//
	// The nobody user has the least permissions in Linux and can access only the public content of the file system. This ensures the security of the file system.
	//
	// example:
	//
	// no_squash
	UserAccessType *string `json:"UserAccessType,omitempty" xml:"UserAccessType,omitempty"`
}

func (s CreateAccessRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAccessRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateAccessRuleRequest) GetAccessGroupName() *string {
	return s.AccessGroupName
}

func (s *CreateAccessRuleRequest) GetFileSystemType() *string {
	return s.FileSystemType
}

func (s *CreateAccessRuleRequest) GetIpv6SourceCidrIp() *string {
	return s.Ipv6SourceCidrIp
}

func (s *CreateAccessRuleRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateAccessRuleRequest) GetRWAccessType() *string {
	return s.RWAccessType
}

func (s *CreateAccessRuleRequest) GetSourceCidrIp() *string {
	return s.SourceCidrIp
}

func (s *CreateAccessRuleRequest) GetUserAccessType() *string {
	return s.UserAccessType
}

func (s *CreateAccessRuleRequest) SetAccessGroupName(v string) *CreateAccessRuleRequest {
	s.AccessGroupName = &v
	return s
}

func (s *CreateAccessRuleRequest) SetFileSystemType(v string) *CreateAccessRuleRequest {
	s.FileSystemType = &v
	return s
}

func (s *CreateAccessRuleRequest) SetIpv6SourceCidrIp(v string) *CreateAccessRuleRequest {
	s.Ipv6SourceCidrIp = &v
	return s
}

func (s *CreateAccessRuleRequest) SetPriority(v int32) *CreateAccessRuleRequest {
	s.Priority = &v
	return s
}

func (s *CreateAccessRuleRequest) SetRWAccessType(v string) *CreateAccessRuleRequest {
	s.RWAccessType = &v
	return s
}

func (s *CreateAccessRuleRequest) SetSourceCidrIp(v string) *CreateAccessRuleRequest {
	s.SourceCidrIp = &v
	return s
}

func (s *CreateAccessRuleRequest) SetUserAccessType(v string) *CreateAccessRuleRequest {
	s.UserAccessType = &v
	return s
}

func (s *CreateAccessRuleRequest) Validate() error {
	return dara.Validate(s)
}
