// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeZonesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemType(v string) *DescribeZonesRequest
	GetFileSystemType() *string
	SetRegionId(v string) *DescribeZonesRequest
	GetRegionId() *string
}

type DescribeZonesRequest struct {
	// The type of the file system.
	//
	// Valid value:
	//
	// 	- standard: General-purpose Apsara File Storage NAS (NAS) file system
	//
	// 	- extreme: Extreme NAS file system.
	//
	// 	- cpfs: CPFS file system.
	//
	// example:
	//
	// standard
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	// The ID of the region where you want to query zones.
	//
	// You can call the DescribeRegions operation to query the latest region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeZonesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesRequest) GoString() string {
	return s.String()
}

func (s *DescribeZonesRequest) GetFileSystemType() *string {
	return s.FileSystemType
}

func (s *DescribeZonesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeZonesRequest) SetFileSystemType(v string) *DescribeZonesRequest {
	s.FileSystemType = &v
	return s
}

func (s *DescribeZonesRequest) SetRegionId(v string) *DescribeZonesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeZonesRequest) Validate() error {
	return dara.Validate(s)
}
