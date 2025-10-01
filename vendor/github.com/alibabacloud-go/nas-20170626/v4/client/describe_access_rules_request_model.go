// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccessRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessGroupName(v string) *DescribeAccessRulesRequest
	GetAccessGroupName() *string
	SetAccessRuleId(v string) *DescribeAccessRulesRequest
	GetAccessRuleId() *string
	SetFileSystemType(v string) *DescribeAccessRulesRequest
	GetFileSystemType() *string
	SetPageNumber(v int32) *DescribeAccessRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAccessRulesRequest
	GetPageSize() *int32
}

type DescribeAccessRulesRequest struct {
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
	// example:
	//
	// 1
	AccessRuleId *string `json:"AccessRuleId,omitempty" xml:"AccessRuleId,omitempty"`
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
	// The page number.
	//
	// Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeAccessRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccessRulesRequest) GetAccessGroupName() *string {
	return s.AccessGroupName
}

func (s *DescribeAccessRulesRequest) GetAccessRuleId() *string {
	return s.AccessRuleId
}

func (s *DescribeAccessRulesRequest) GetFileSystemType() *string {
	return s.FileSystemType
}

func (s *DescribeAccessRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAccessRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAccessRulesRequest) SetAccessGroupName(v string) *DescribeAccessRulesRequest {
	s.AccessGroupName = &v
	return s
}

func (s *DescribeAccessRulesRequest) SetAccessRuleId(v string) *DescribeAccessRulesRequest {
	s.AccessRuleId = &v
	return s
}

func (s *DescribeAccessRulesRequest) SetFileSystemType(v string) *DescribeAccessRulesRequest {
	s.FileSystemType = &v
	return s
}

func (s *DescribeAccessRulesRequest) SetPageNumber(v int32) *DescribeAccessRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAccessRulesRequest) SetPageSize(v int32) *DescribeAccessRulesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAccessRulesRequest) Validate() error {
	return dara.Validate(s)
}
