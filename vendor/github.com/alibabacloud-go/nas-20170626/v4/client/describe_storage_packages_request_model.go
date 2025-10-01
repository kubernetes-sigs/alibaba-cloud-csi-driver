// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStoragePackagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeStoragePackagesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeStoragePackagesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeStoragePackagesRequest
	GetRegionId() *string
	SetUseUTCDateTime(v bool) *DescribeStoragePackagesRequest
	GetUseUTCDateTime() *bool
}

type DescribeStoragePackagesRequest struct {
	// The number of the page to return.
	//
	// Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of storage plans to return on each page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether the time to return is in UTC.
	//
	// Valid values:
	//
	// 	- true (default): returns UTC time.
	//
	// 	- false: returns UNIX timestamp.
	//
	// example:
	//
	// true
	UseUTCDateTime *bool `json:"UseUTCDateTime,omitempty" xml:"UseUTCDateTime,omitempty"`
}

func (s DescribeStoragePackagesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeStoragePackagesRequest) GoString() string {
	return s.String()
}

func (s *DescribeStoragePackagesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeStoragePackagesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeStoragePackagesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeStoragePackagesRequest) GetUseUTCDateTime() *bool {
	return s.UseUTCDateTime
}

func (s *DescribeStoragePackagesRequest) SetPageNumber(v int32) *DescribeStoragePackagesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeStoragePackagesRequest) SetPageSize(v int32) *DescribeStoragePackagesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeStoragePackagesRequest) SetRegionId(v string) *DescribeStoragePackagesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeStoragePackagesRequest) SetUseUTCDateTime(v bool) *DescribeStoragePackagesRequest {
	s.UseUTCDateTime = &v
	return s
}

func (s *DescribeStoragePackagesRequest) Validate() error {
	return dara.Validate(s)
}
