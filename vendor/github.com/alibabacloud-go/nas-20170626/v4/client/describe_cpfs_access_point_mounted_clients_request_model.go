// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCpfsAccessPointMountedClientsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessPointId(v string) *DescribeCpfsAccessPointMountedClientsRequest
	GetAccessPointId() *string
	SetFileSystemId(v string) *DescribeCpfsAccessPointMountedClientsRequest
	GetFileSystemId() *string
	SetPageNumber(v int32) *DescribeCpfsAccessPointMountedClientsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCpfsAccessPointMountedClientsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeCpfsAccessPointMountedClientsRequest
	GetRegionId() *string
}

type DescribeCpfsAccessPointMountedClientsRequest struct {
	// The access point ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ap-ie15yd****
	AccessPointId *string `json:"AccessPointId,omitempty" xml:"AccessPointId,omitempty"`
	// The file system ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bmcpfs-370lx1ev9ss27o0****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The page number of the list.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of results on each page. Valid values: 1 to 100.
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
}

func (s DescribeCpfsAccessPointMountedClientsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCpfsAccessPointMountedClientsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCpfsAccessPointMountedClientsRequest) GetAccessPointId() *string {
	return s.AccessPointId
}

func (s *DescribeCpfsAccessPointMountedClientsRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DescribeCpfsAccessPointMountedClientsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCpfsAccessPointMountedClientsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCpfsAccessPointMountedClientsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCpfsAccessPointMountedClientsRequest) SetAccessPointId(v string) *DescribeCpfsAccessPointMountedClientsRequest {
	s.AccessPointId = &v
	return s
}

func (s *DescribeCpfsAccessPointMountedClientsRequest) SetFileSystemId(v string) *DescribeCpfsAccessPointMountedClientsRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeCpfsAccessPointMountedClientsRequest) SetPageNumber(v int32) *DescribeCpfsAccessPointMountedClientsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCpfsAccessPointMountedClientsRequest) SetPageSize(v int32) *DescribeCpfsAccessPointMountedClientsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCpfsAccessPointMountedClientsRequest) SetRegionId(v string) *DescribeCpfsAccessPointMountedClientsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCpfsAccessPointMountedClientsRequest) Validate() error {
	return dara.Validate(s)
}
