// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemType(v string) *DescribeRegionsRequest
	GetFileSystemType() *string
	SetPageNumber(v int32) *DescribeRegionsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeRegionsRequest
	GetPageSize() *int32
}

type DescribeRegionsRequest struct {
	// The type of the file system.
	//
	// Valid values:
	//
	// 	- all: all types of file systems
	//
	// 	- standard (default): General-purpose NAS file system
	//
	// 	- extreme: Extreme NAS file system
	//
	// 	- cpfs: Cloud Parallel File Storage (CPFS) file system
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
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) GetFileSystemType() *string {
	return s.FileSystemType
}

func (s *DescribeRegionsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeRegionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRegionsRequest) SetFileSystemType(v string) *DescribeRegionsRequest {
	s.FileSystemType = &v
	return s
}

func (s *DescribeRegionsRequest) SetPageNumber(v int32) *DescribeRegionsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeRegionsRequest) SetPageSize(v int32) *DescribeRegionsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRegionsRequest) Validate() error {
	return dara.Validate(s)
}
