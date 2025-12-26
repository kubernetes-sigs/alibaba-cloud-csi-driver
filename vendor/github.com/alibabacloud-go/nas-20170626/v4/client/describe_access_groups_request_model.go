// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccessGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessGroupName(v string) *DescribeAccessGroupsRequest
	GetAccessGroupName() *string
	SetFileSystemType(v string) *DescribeAccessGroupsRequest
	GetFileSystemType() *string
	SetPageNumber(v int32) *DescribeAccessGroupsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAccessGroupsRequest
	GetPageSize() *int32
	SetUseUTCDateTime(v bool) *DescribeAccessGroupsRequest
	GetUseUTCDateTime() *bool
}

type DescribeAccessGroupsRequest struct {
	// The name of the permission group.
	//
	// Limits:
	//
	// 	- The name must be 3 to 64 characters in length.
	//
	// 	- The name must start with a letter and can contain letters, digits, underscores (_), and hyphens (-).
	//
	// example:
	//
	// DEFAULT_VPC_GROUP_NAME
	AccessGroupName *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	// The type of the file system.
	//
	// Valid values:
	//
	// 	- standard (default): General-purpose NAS file system.
	//
	// 	- extreme: Extreme NAS file system.
	//
	// 	- cpfs: Cloud Parallel File Storage (CPFS) file system.
	//
	// > CPFS file systems are available only on the China site (aliyun.com).
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
	// 2
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Specifies whether to display the creation time of the permission group in UTC.
	//
	// Valid values:
	//
	// 	- true (default): The time is displayed in UTC.
	//
	// 	- false: The time is not displayed in UTC.
	//
	// example:
	//
	// true
	UseUTCDateTime *bool `json:"UseUTCDateTime,omitempty" xml:"UseUTCDateTime,omitempty"`
}

func (s DescribeAccessGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccessGroupsRequest) GetAccessGroupName() *string {
	return s.AccessGroupName
}

func (s *DescribeAccessGroupsRequest) GetFileSystemType() *string {
	return s.FileSystemType
}

func (s *DescribeAccessGroupsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAccessGroupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAccessGroupsRequest) GetUseUTCDateTime() *bool {
	return s.UseUTCDateTime
}

func (s *DescribeAccessGroupsRequest) SetAccessGroupName(v string) *DescribeAccessGroupsRequest {
	s.AccessGroupName = &v
	return s
}

func (s *DescribeAccessGroupsRequest) SetFileSystemType(v string) *DescribeAccessGroupsRequest {
	s.FileSystemType = &v
	return s
}

func (s *DescribeAccessGroupsRequest) SetPageNumber(v int32) *DescribeAccessGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAccessGroupsRequest) SetPageSize(v int32) *DescribeAccessGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAccessGroupsRequest) SetUseUTCDateTime(v bool) *DescribeAccessGroupsRequest {
	s.UseUTCDateTime = &v
	return s
}

func (s *DescribeAccessGroupsRequest) Validate() error {
	return dara.Validate(s)
}
