// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSnapshotPackageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeSnapshotPackageRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSnapshotPackageRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeSnapshotPackageRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSnapshotPackageRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeSnapshotPackageRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeSnapshotPackageRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSnapshotPackageRequest
	GetResourceOwnerId() *int64
}

type DescribeSnapshotPackageRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Pages start from page 1.
	//
	// Default value: 1
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: 1 to 100.
	//
	// Default value: 10
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeSnapshotPackageRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnapshotPackageRequest) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotPackageRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSnapshotPackageRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSnapshotPackageRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSnapshotPackageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSnapshotPackageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSnapshotPackageRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSnapshotPackageRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSnapshotPackageRequest) SetOwnerAccount(v string) *DescribeSnapshotPackageRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSnapshotPackageRequest) SetOwnerId(v int64) *DescribeSnapshotPackageRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSnapshotPackageRequest) SetPageNumber(v int32) *DescribeSnapshotPackageRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSnapshotPackageRequest) SetPageSize(v int32) *DescribeSnapshotPackageRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSnapshotPackageRequest) SetRegionId(v string) *DescribeSnapshotPackageRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSnapshotPackageRequest) SetResourceOwnerAccount(v string) *DescribeSnapshotPackageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSnapshotPackageRequest) SetResourceOwnerId(v int64) *DescribeSnapshotPackageRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSnapshotPackageRequest) Validate() error {
	return dara.Validate(s)
}
