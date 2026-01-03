// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageSharePermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v string) *DescribeImageSharePermissionRequest
	GetImageId() *string
	SetOwnerAccount(v string) *DescribeImageSharePermissionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeImageSharePermissionRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeImageSharePermissionRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeImageSharePermissionRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeImageSharePermissionRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeImageSharePermissionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeImageSharePermissionRequest
	GetResourceOwnerId() *int64
}

type DescribeImageSharePermissionRequest struct {
	// The ID of the custom image.
	//
	// This parameter is required.
	//
	// example:
	//
	// m-bp1caf3yicx5jlfl****
	ImageId      *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number.
	//
	// Pages start from page 1.
	//
	// Default value: 1.
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
	// The region ID of the custom image. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
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

func (s DescribeImageSharePermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageSharePermissionRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageSharePermissionRequest) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeImageSharePermissionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeImageSharePermissionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeImageSharePermissionRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeImageSharePermissionRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeImageSharePermissionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeImageSharePermissionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeImageSharePermissionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeImageSharePermissionRequest) SetImageId(v string) *DescribeImageSharePermissionRequest {
	s.ImageId = &v
	return s
}

func (s *DescribeImageSharePermissionRequest) SetOwnerAccount(v string) *DescribeImageSharePermissionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeImageSharePermissionRequest) SetOwnerId(v int64) *DescribeImageSharePermissionRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeImageSharePermissionRequest) SetPageNumber(v int32) *DescribeImageSharePermissionRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeImageSharePermissionRequest) SetPageSize(v int32) *DescribeImageSharePermissionRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeImageSharePermissionRequest) SetRegionId(v string) *DescribeImageSharePermissionRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeImageSharePermissionRequest) SetResourceOwnerAccount(v string) *DescribeImageSharePermissionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeImageSharePermissionRequest) SetResourceOwnerId(v int64) *DescribeImageSharePermissionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeImageSharePermissionRequest) Validate() error {
	return dara.Validate(s)
}
