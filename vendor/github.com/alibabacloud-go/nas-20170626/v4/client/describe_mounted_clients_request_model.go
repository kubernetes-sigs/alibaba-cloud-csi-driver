// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMountedClientsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientIP(v string) *DescribeMountedClientsRequest
	GetClientIP() *string
	SetFileSystemId(v string) *DescribeMountedClientsRequest
	GetFileSystemId() *string
	SetMountTargetDomain(v string) *DescribeMountedClientsRequest
	GetMountTargetDomain() *string
	SetPageNumber(v int32) *DescribeMountedClientsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeMountedClientsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeMountedClientsRequest
	GetRegionId() *string
}

type DescribeMountedClientsRequest struct {
	// The IP address of the client.
	//
	// 	- If you specify an IP address, the operation checks whether the client list includes this IP address. If the client list includes the IP address, the operation returns the IP address. If the client list does not include the IP address, the operation returns an empty list.
	//
	// 	- If you do not specify an IP address, the operation returns the IP addresses of all clients that have accessed the specified NAS file system within the last minute.
	//
	// example:
	//
	// 10.10.10.1
	ClientIP *string `json:"ClientIP,omitempty" xml:"ClientIP,omitempty"`
	// The ID of the file system.
	//
	// This parameter is required.
	//
	// example:
	//
	// 109c****66
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The domain name of the mount target.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1111222****95.cn-hangzhou.nas.aliyuncs.com
	MountTargetDomain *string `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty"`
	// The page number.
	//
	// Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of IP addresses to return on each page.
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
}

func (s DescribeMountedClientsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMountedClientsRequest) GoString() string {
	return s.String()
}

func (s *DescribeMountedClientsRequest) GetClientIP() *string {
	return s.ClientIP
}

func (s *DescribeMountedClientsRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DescribeMountedClientsRequest) GetMountTargetDomain() *string {
	return s.MountTargetDomain
}

func (s *DescribeMountedClientsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeMountedClientsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeMountedClientsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeMountedClientsRequest) SetClientIP(v string) *DescribeMountedClientsRequest {
	s.ClientIP = &v
	return s
}

func (s *DescribeMountedClientsRequest) SetFileSystemId(v string) *DescribeMountedClientsRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeMountedClientsRequest) SetMountTargetDomain(v string) *DescribeMountedClientsRequest {
	s.MountTargetDomain = &v
	return s
}

func (s *DescribeMountedClientsRequest) SetPageNumber(v int32) *DescribeMountedClientsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeMountedClientsRequest) SetPageSize(v int32) *DescribeMountedClientsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeMountedClientsRequest) SetRegionId(v string) *DescribeMountedClientsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMountedClientsRequest) Validate() error {
	return dara.Validate(s)
}
