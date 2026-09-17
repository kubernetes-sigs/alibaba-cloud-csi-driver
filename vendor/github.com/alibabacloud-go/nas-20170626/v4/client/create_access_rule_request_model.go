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
	// The value supports CIDR format and IPv6 format address range.
	//
	// > - The IPv6 feature is supported only by Extreme NAS file systems in regions in the Chinese mainland, and IPv6 must be enabled for the file system.
	//
	// >- Only VPC networks are supported.
	//
	// >- IPv4 and IPv6 are mutually exclusive.
	//
	// example:
	//
	// 2001:250:6000::***
	Ipv6SourceCidrIp *string `json:"Ipv6SourceCidrIp,omitempty" xml:"Ipv6SourceCidrIp,omitempty"`
	// The priority of the permission rule.
	//
	// If an authorized address matches multiple rules, the rule with the highest priority takes effect.
	//
	// Valid values: 1 to 100. The value 1 indicates the highest priority.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The read and write permissions of the authorized address on the file system.
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
	// The IP address or CIDR block of the authorized address.
	//
	// The value must be a single IP address or a CIDR block.
	//
	// > Permission groups of the classic network type support only IP addresses.
	//
	// example:
	//
	// 192.0.2.0/16
	SourceCidrIp *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
	// The access permissions of the system user of the authorized address on the file system.
	//
	// Valid values:
	//
	// - no_squash (default): allows access to the file system as the root user.
	//
	// - root_squash: maps the root user to the nobody user when the root user accesses the file system.
	//
	// - all_squash: maps all users to the nobody user regardless of the user identity.
	//
	// The nobody user is a default user in Linux. The nobody user can access only public content on the server and has low privileges and high security. Authorization is required for the system user to access the file system.
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
