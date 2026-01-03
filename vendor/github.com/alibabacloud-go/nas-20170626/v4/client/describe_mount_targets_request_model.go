// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMountTargetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDualStackMountTargetDomain(v string) *DescribeMountTargetsRequest
	GetDualStackMountTargetDomain() *string
	SetFileSystemId(v string) *DescribeMountTargetsRequest
	GetFileSystemId() *string
	SetMountTargetDomain(v string) *DescribeMountTargetsRequest
	GetMountTargetDomain() *string
	SetPageNumber(v int32) *DescribeMountTargetsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeMountTargetsRequest
	GetPageSize() *int32
}

type DescribeMountTargetsRequest struct {
	// The dual-stack (IPv4 and IPv6) domain name of the mount target.
	//
	// > Only Extreme NAS file systems that reside in the Chinese mainland support IPv6.
	//
	// example:
	//
	// 1ca404****-x****.dualstack.cn-hangzhou.nas.aliyuncs.com
	DualStackMountTargetDomain *string `json:"DualStackMountTargetDomain,omitempty" xml:"DualStackMountTargetDomain,omitempty"`
	// The ID of the file system.
	//
	// 	- Sample ID of a General-purpose NAS file system: 31a8e4\\*\\*\\*\\*.
	//
	// 	- The IDs of Extreme NAS file systems must start with `extreme-`, for example, extreme-0015\\*\\*\\*\\*.
	//
	// 	- The IDs of CPFS file systems must start with `cpfs-`. Example: cpfs-125487\\*\\*\\*\\*.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1ca404****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The address of the mount target.
	//
	// example:
	//
	// 1ca404****-x****.cn-hangzhou.nas.aliyuncs.com
	MountTargetDomain *string `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty"`
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
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeMountTargetsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMountTargetsRequest) GoString() string {
	return s.String()
}

func (s *DescribeMountTargetsRequest) GetDualStackMountTargetDomain() *string {
	return s.DualStackMountTargetDomain
}

func (s *DescribeMountTargetsRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DescribeMountTargetsRequest) GetMountTargetDomain() *string {
	return s.MountTargetDomain
}

func (s *DescribeMountTargetsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeMountTargetsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeMountTargetsRequest) SetDualStackMountTargetDomain(v string) *DescribeMountTargetsRequest {
	s.DualStackMountTargetDomain = &v
	return s
}

func (s *DescribeMountTargetsRequest) SetFileSystemId(v string) *DescribeMountTargetsRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeMountTargetsRequest) SetMountTargetDomain(v string) *DescribeMountTargetsRequest {
	s.MountTargetDomain = &v
	return s
}

func (s *DescribeMountTargetsRequest) SetPageNumber(v int32) *DescribeMountTargetsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeMountTargetsRequest) SetPageSize(v int32) *DescribeMountTargetsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeMountTargetsRequest) Validate() error {
	return dara.Validate(s)
}
