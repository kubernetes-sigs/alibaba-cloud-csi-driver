// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageFromFamilyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageFamily(v string) *DescribeImageFromFamilyRequest
	GetImageFamily() *string
	SetOwnerAccount(v string) *DescribeImageFromFamilyRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeImageFromFamilyRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeImageFromFamilyRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeImageFromFamilyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeImageFromFamilyRequest
	GetResourceOwnerId() *int64
}

type DescribeImageFromFamilyRequest struct {
	// The family name of the image that you want to use to create the instances.
	//
	// You can configure image families for custom images, public images, community images, and shared images. For more information, see [Overview of image families](https://help.aliyun.com/document_detail/174241.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// hangzhou-daily-update
	ImageFamily  *string `json:"ImageFamily,omitempty" xml:"ImageFamily,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region in which to create the custom image. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
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

func (s DescribeImageFromFamilyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageFromFamilyRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageFromFamilyRequest) GetImageFamily() *string {
	return s.ImageFamily
}

func (s *DescribeImageFromFamilyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeImageFromFamilyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeImageFromFamilyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeImageFromFamilyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeImageFromFamilyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeImageFromFamilyRequest) SetImageFamily(v string) *DescribeImageFromFamilyRequest {
	s.ImageFamily = &v
	return s
}

func (s *DescribeImageFromFamilyRequest) SetOwnerAccount(v string) *DescribeImageFromFamilyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeImageFromFamilyRequest) SetOwnerId(v int64) *DescribeImageFromFamilyRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeImageFromFamilyRequest) SetRegionId(v string) *DescribeImageFromFamilyRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeImageFromFamilyRequest) SetResourceOwnerAccount(v string) *DescribeImageFromFamilyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeImageFromFamilyRequest) SetResourceOwnerId(v int64) *DescribeImageFromFamilyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeImageFromFamilyRequest) Validate() error {
	return dara.Validate(s)
}
