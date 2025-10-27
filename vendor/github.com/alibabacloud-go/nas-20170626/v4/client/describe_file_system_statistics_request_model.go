// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFileSystemStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeFileSystemStatisticsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeFileSystemStatisticsRequest
	GetPageSize() *int32
}

type DescribeFileSystemStatisticsRequest struct {
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

func (s DescribeFileSystemStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileSystemStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemStatisticsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeFileSystemStatisticsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeFileSystemStatisticsRequest) SetPageNumber(v int32) *DescribeFileSystemStatisticsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeFileSystemStatisticsRequest) SetPageSize(v int32) *DescribeFileSystemStatisticsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeFileSystemStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
