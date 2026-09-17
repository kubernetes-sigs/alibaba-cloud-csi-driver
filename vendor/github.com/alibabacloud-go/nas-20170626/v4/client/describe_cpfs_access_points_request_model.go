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
	SetTag(v []*DescribeCpfsAccessPointsRequestTag) *DescribeCpfsAccessPointsRequest
	GetTag() []*DescribeCpfsAccessPointsRequestTag
}

type DescribeCpfsAccessPointsRequest struct {
	// The access point ID.
	//
	// example:
	//
	// ap-ie15yd****
	AccessPointId *string `json:"AccessPointId,omitempty" xml:"AccessPointId,omitempty"`
	// The file system ID.
	//
	// - CPFS: The ID must start with `cpfs-`, such as cpfs-099394bd928c****.
	//
	// - CPFS for Lingjun: The ID must start with `bmcpfs-`, such as bmcpfs-290w65p03ok64ya****.
	//
	// This parameter is required.
	//
	// example:
	//
	// bmcpfs-290rg9crq96m362ups2
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The page number of the list.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of results per query.
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
	// The list of CPFS access point tags.
	Tag []*DescribeCpfsAccessPointsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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

func (s *DescribeCpfsAccessPointsRequest) GetTag() []*DescribeCpfsAccessPointsRequestTag {
	return s.Tag
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

func (s *DescribeCpfsAccessPointsRequest) SetTag(v []*DescribeCpfsAccessPointsRequestTag) *DescribeCpfsAccessPointsRequest {
	s.Tag = v
	return s
}

func (s *DescribeCpfsAccessPointsRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCpfsAccessPointsRequestTag struct {
	// The key of the CPFS access point tag.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the CPFS access point tag.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeCpfsAccessPointsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeCpfsAccessPointsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeCpfsAccessPointsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeCpfsAccessPointsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeCpfsAccessPointsRequestTag) SetKey(v string) *DescribeCpfsAccessPointsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeCpfsAccessPointsRequestTag) SetValue(v string) *DescribeCpfsAccessPointsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeCpfsAccessPointsRequestTag) Validate() error {
	return dara.Validate(s)
}
