// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCpfsAccessPointsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessPointId(v string) *DescribeCpfsAccessPointsRequest
	GetAccessPointId() *string
	SetFileSystemId(v string) *DescribeCpfsAccessPointsRequest
	GetFileSystemId() *string
	SetPageNumber(v int32) *DescribeCpfsAccessPointsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCpfsAccessPointsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeCpfsAccessPointsRequest
	GetRegionId() *string
}

type DescribeCpfsAccessPointsRequest struct {
	// example:
	//
	// ap-ie15yd****
	AccessPointId *string `json:"AccessPointId,omitempty" xml:"AccessPointId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// bmcpfs-290rg9crq96m362ups2
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCpfsAccessPointsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCpfsAccessPointsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCpfsAccessPointsRequest) GetAccessPointId() *string {
	return s.AccessPointId
}

func (s *DescribeCpfsAccessPointsRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DescribeCpfsAccessPointsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCpfsAccessPointsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCpfsAccessPointsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCpfsAccessPointsRequest) SetAccessPointId(v string) *DescribeCpfsAccessPointsRequest {
	s.AccessPointId = &v
	return s
}

func (s *DescribeCpfsAccessPointsRequest) SetFileSystemId(v string) *DescribeCpfsAccessPointsRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeCpfsAccessPointsRequest) SetPageNumber(v int32) *DescribeCpfsAccessPointsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCpfsAccessPointsRequest) SetPageSize(v int32) *DescribeCpfsAccessPointsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCpfsAccessPointsRequest) SetRegionId(v string) *DescribeCpfsAccessPointsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCpfsAccessPointsRequest) Validate() error {
	return dara.Validate(s)
}
