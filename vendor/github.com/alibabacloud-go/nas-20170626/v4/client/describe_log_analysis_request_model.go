// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogAnalysisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemType(v string) *DescribeLogAnalysisRequest
	GetFileSystemType() *string
	SetPageNumber(v int32) *DescribeLogAnalysisRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeLogAnalysisRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeLogAnalysisRequest
	GetRegionId() *string
}

type DescribeLogAnalysisRequest struct {
	// The type of the file system.
	//
	// Valid values:
	//
	// 	- standard: General-purpose NAS file system
	//
	// 	- extreme: Extreme NAS file system
	//
	// 	- all (default): all types
	//
	// example:
	//
	// all
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeLogAnalysisRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogAnalysisRequest) GoString() string {
	return s.String()
}

func (s *DescribeLogAnalysisRequest) GetFileSystemType() *string {
	return s.FileSystemType
}

func (s *DescribeLogAnalysisRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeLogAnalysisRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLogAnalysisRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLogAnalysisRequest) SetFileSystemType(v string) *DescribeLogAnalysisRequest {
	s.FileSystemType = &v
	return s
}

func (s *DescribeLogAnalysisRequest) SetPageNumber(v int32) *DescribeLogAnalysisRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeLogAnalysisRequest) SetPageSize(v int32) *DescribeLogAnalysisRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLogAnalysisRequest) SetRegionId(v string) *DescribeLogAnalysisRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLogAnalysisRequest) Validate() error {
	return dara.Validate(s)
}
