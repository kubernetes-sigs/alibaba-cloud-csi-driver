// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelCopyImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v string) *CancelCopyImageRequest
	GetImageId() *string
	SetOwnerAccount(v string) *CancelCopyImageRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CancelCopyImageRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CancelCopyImageRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CancelCopyImageRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CancelCopyImageRequest
	GetResourceOwnerId() *int64
}

type CancelCopyImageRequest struct {
	// The ID of the image that is being copied.
	//
	// This parameter is required.
	//
	// example:
	//
	// m-bp1caf3yicx5jlfl****
	ImageId      *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the image copy. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
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

func (s CancelCopyImageRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelCopyImageRequest) GoString() string {
	return s.String()
}

func (s *CancelCopyImageRequest) GetImageId() *string {
	return s.ImageId
}

func (s *CancelCopyImageRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CancelCopyImageRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CancelCopyImageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CancelCopyImageRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CancelCopyImageRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CancelCopyImageRequest) SetImageId(v string) *CancelCopyImageRequest {
	s.ImageId = &v
	return s
}

func (s *CancelCopyImageRequest) SetOwnerAccount(v string) *CancelCopyImageRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CancelCopyImageRequest) SetOwnerId(v int64) *CancelCopyImageRequest {
	s.OwnerId = &v
	return s
}

func (s *CancelCopyImageRequest) SetRegionId(v string) *CancelCopyImageRequest {
	s.RegionId = &v
	return s
}

func (s *CancelCopyImageRequest) SetResourceOwnerAccount(v string) *CancelCopyImageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CancelCopyImageRequest) SetResourceOwnerId(v int64) *CancelCopyImageRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CancelCopyImageRequest) Validate() error {
	return dara.Validate(s)
}
