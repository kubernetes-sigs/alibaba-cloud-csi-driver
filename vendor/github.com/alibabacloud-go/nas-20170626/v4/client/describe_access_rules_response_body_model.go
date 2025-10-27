// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccessRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessRules(v *DescribeAccessRulesResponseBodyAccessRules) *DescribeAccessRulesResponseBody
	GetAccessRules() *DescribeAccessRulesResponseBodyAccessRules
	SetPageNumber(v int32) *DescribeAccessRulesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAccessRulesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeAccessRulesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeAccessRulesResponseBody
	GetTotalCount() *int32
}

type DescribeAccessRulesResponseBody struct {
	// The rules in the permission group.
	AccessRules *DescribeAccessRulesResponseBodyAccessRules `json:"AccessRules,omitempty" xml:"AccessRules,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 86D89E82-4297-4343-8E1E-A2495B35****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of rules.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAccessRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccessRulesResponseBody) GetAccessRules() *DescribeAccessRulesResponseBodyAccessRules {
	return s.AccessRules
}

func (s *DescribeAccessRulesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAccessRulesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAccessRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAccessRulesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeAccessRulesResponseBody) SetAccessRules(v *DescribeAccessRulesResponseBodyAccessRules) *DescribeAccessRulesResponseBody {
	s.AccessRules = v
	return s
}

func (s *DescribeAccessRulesResponseBody) SetPageNumber(v int32) *DescribeAccessRulesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAccessRulesResponseBody) SetPageSize(v int32) *DescribeAccessRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAccessRulesResponseBody) SetRequestId(v string) *DescribeAccessRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccessRulesResponseBody) SetTotalCount(v int32) *DescribeAccessRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAccessRulesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAccessRulesResponseBodyAccessRules struct {
	AccessRule []*DescribeAccessRulesResponseBodyAccessRulesAccessRule `json:"AccessRule,omitempty" xml:"AccessRule,omitempty" type:"Repeated"`
}

func (s DescribeAccessRulesResponseBodyAccessRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessRulesResponseBodyAccessRules) GoString() string {
	return s.String()
}

func (s *DescribeAccessRulesResponseBodyAccessRules) GetAccessRule() []*DescribeAccessRulesResponseBodyAccessRulesAccessRule {
	return s.AccessRule
}

func (s *DescribeAccessRulesResponseBodyAccessRules) SetAccessRule(v []*DescribeAccessRulesResponseBodyAccessRulesAccessRule) *DescribeAccessRulesResponseBodyAccessRules {
	s.AccessRule = v
	return s
}

func (s *DescribeAccessRulesResponseBodyAccessRules) Validate() error {
	return dara.Validate(s)
}

type DescribeAccessRulesResponseBodyAccessRulesAccessRule struct {
	// The name of the permission group.
	//
	// example:
	//
	// test
	AccessGroupName *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	// The ID of the rule.
	//
	// example:
	//
	// 1
	AccessRuleId *string `json:"AccessRuleId,omitempty" xml:"AccessRuleId,omitempty"`
	// The type of the file system.
	//
	// Valid values:
	//
	// 	- standard: General-purpose File Storage NAS (NAS) file system
	//
	// 	- extreme: Extreme NAS file system
	//
	// example:
	//
	// standard
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	// The IPv6 address or CIDR block of the authorized object.
	//
	// example:
	//
	// 2001:250:6000::***
	Ipv6SourceCidrIp *string `json:"Ipv6SourceCidrIp,omitempty" xml:"Ipv6SourceCidrIp,omitempty"`
	// The priority of the rule.
	//
	// If multiple rules are attached to the authorized object, the rule with the highest priority takes effect.
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
	RWAccess *string `json:"RWAccess,omitempty" xml:"RWAccess,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IP address or CIDR block of the authorized object.
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
	// no_squash
	UserAccess *string `json:"UserAccess,omitempty" xml:"UserAccess,omitempty"`
}

func (s DescribeAccessRulesResponseBodyAccessRulesAccessRule) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessRulesResponseBodyAccessRulesAccessRule) GoString() string {
	return s.String()
}

func (s *DescribeAccessRulesResponseBodyAccessRulesAccessRule) GetAccessGroupName() *string {
	return s.AccessGroupName
}

func (s *DescribeAccessRulesResponseBodyAccessRulesAccessRule) GetAccessRuleId() *string {
	return s.AccessRuleId
}

func (s *DescribeAccessRulesResponseBodyAccessRulesAccessRule) GetFileSystemType() *string {
	return s.FileSystemType
}

func (s *DescribeAccessRulesResponseBodyAccessRulesAccessRule) GetIpv6SourceCidrIp() *string {
	return s.Ipv6SourceCidrIp
}

func (s *DescribeAccessRulesResponseBodyAccessRulesAccessRule) GetPriority() *int32 {
	return s.Priority
}

func (s *DescribeAccessRulesResponseBodyAccessRulesAccessRule) GetRWAccess() *string {
	return s.RWAccess
}

func (s *DescribeAccessRulesResponseBodyAccessRulesAccessRule) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAccessRulesResponseBodyAccessRulesAccessRule) GetSourceCidrIp() *string {
	return s.SourceCidrIp
}

func (s *DescribeAccessRulesResponseBodyAccessRulesAccessRule) GetUserAccess() *string {
	return s.UserAccess
}

func (s *DescribeAccessRulesResponseBodyAccessRulesAccessRule) SetAccessGroupName(v string) *DescribeAccessRulesResponseBodyAccessRulesAccessRule {
	s.AccessGroupName = &v
	return s
}

func (s *DescribeAccessRulesResponseBodyAccessRulesAccessRule) SetAccessRuleId(v string) *DescribeAccessRulesResponseBodyAccessRulesAccessRule {
	s.AccessRuleId = &v
	return s
}

func (s *DescribeAccessRulesResponseBodyAccessRulesAccessRule) SetFileSystemType(v string) *DescribeAccessRulesResponseBodyAccessRulesAccessRule {
	s.FileSystemType = &v
	return s
}

func (s *DescribeAccessRulesResponseBodyAccessRulesAccessRule) SetIpv6SourceCidrIp(v string) *DescribeAccessRulesResponseBodyAccessRulesAccessRule {
	s.Ipv6SourceCidrIp = &v
	return s
}

func (s *DescribeAccessRulesResponseBodyAccessRulesAccessRule) SetPriority(v int32) *DescribeAccessRulesResponseBodyAccessRulesAccessRule {
	s.Priority = &v
	return s
}

func (s *DescribeAccessRulesResponseBodyAccessRulesAccessRule) SetRWAccess(v string) *DescribeAccessRulesResponseBodyAccessRulesAccessRule {
	s.RWAccess = &v
	return s
}

func (s *DescribeAccessRulesResponseBodyAccessRulesAccessRule) SetRegionId(v string) *DescribeAccessRulesResponseBodyAccessRulesAccessRule {
	s.RegionId = &v
	return s
}

func (s *DescribeAccessRulesResponseBodyAccessRulesAccessRule) SetSourceCidrIp(v string) *DescribeAccessRulesResponseBodyAccessRulesAccessRule {
	s.SourceCidrIp = &v
	return s
}

func (s *DescribeAccessRulesResponseBodyAccessRulesAccessRule) SetUserAccess(v string) *DescribeAccessRulesResponseBodyAccessRulesAccessRule {
	s.UserAccess = &v
	return s
}

func (s *DescribeAccessRulesResponseBodyAccessRulesAccessRule) Validate() error {
	return dara.Validate(s)
}
