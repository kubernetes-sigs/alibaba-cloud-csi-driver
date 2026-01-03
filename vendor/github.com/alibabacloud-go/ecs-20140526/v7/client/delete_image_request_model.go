// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v bool) *DeleteImageRequest
	GetDryRun() *bool
	SetForce(v bool) *DeleteImageRequest
	GetForce() *bool
	SetImageId(v string) *DeleteImageRequest
	GetImageId() *string
	SetOwnerAccount(v string) *DeleteImageRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteImageRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteImageRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteImageRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteImageRequest
	GetResourceOwnerId() *int64
}

type DeleteImageRequest struct {
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to forcefully delete the custom image. Valid values:
	//
	// 	- true: forcefully deletes the custom image, regardless of whether the image is being used by instances.
	//
	// 	- false: verifies that the custom image is not being used by instances and then deletes the image.
	//
	// Default value: false
	//
	// example:
	//
	// false
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// The ID of the image. If the specified custom image does not exist, the request is ignored.
	//
	// This parameter is required.
	//
	// example:
	//
	// m-bp67acfmxazb4p****
	ImageId      *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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

func (s DeleteImageRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteImageRequest) GoString() string {
	return s.String()
}

func (s *DeleteImageRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteImageRequest) GetForce() *bool {
	return s.Force
}

func (s *DeleteImageRequest) GetImageId() *string {
	return s.ImageId
}

func (s *DeleteImageRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteImageRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteImageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteImageRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteImageRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteImageRequest) SetDryRun(v bool) *DeleteImageRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteImageRequest) SetForce(v bool) *DeleteImageRequest {
	s.Force = &v
	return s
}

func (s *DeleteImageRequest) SetImageId(v string) *DeleteImageRequest {
	s.ImageId = &v
	return s
}

func (s *DeleteImageRequest) SetOwnerAccount(v string) *DeleteImageRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteImageRequest) SetOwnerId(v int64) *DeleteImageRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteImageRequest) SetRegionId(v string) *DeleteImageRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteImageRequest) SetResourceOwnerAccount(v string) *DeleteImageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteImageRequest) SetResourceOwnerId(v int64) *DeleteImageRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteImageRequest) Validate() error {
	return dara.Validate(s)
}
