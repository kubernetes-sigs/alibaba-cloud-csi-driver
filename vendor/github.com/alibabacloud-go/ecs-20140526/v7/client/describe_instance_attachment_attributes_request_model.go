// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceAttachmentAttributesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v string) *DescribeInstanceAttachmentAttributesRequest
	GetInstanceIds() *string
	SetOwnerAccount(v string) *DescribeInstanceAttachmentAttributesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeInstanceAttachmentAttributesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeInstanceAttachmentAttributesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeInstanceAttachmentAttributesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeInstanceAttachmentAttributesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeInstanceAttachmentAttributesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeInstanceAttachmentAttributesRequest
	GetResourceOwnerId() *int64
}

type DescribeInstanceAttachmentAttributesRequest struct {
	// The IDs of the instances. The value can be a JSON array that consists of up to 100 instance IDs. Separate the IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// ["i-bp67acfmxazb4****", "i-bp67acfmxazb5****", "i-bp67acfmxazb6****"]
	InstanceIds  *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number.
	//
	// Page starts from page 1.
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
	// The region ID of the elasticity assurance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
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

func (s DescribeInstanceAttachmentAttributesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAttachmentAttributesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttachmentAttributesRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *DescribeInstanceAttachmentAttributesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeInstanceAttachmentAttributesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeInstanceAttachmentAttributesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeInstanceAttachmentAttributesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInstanceAttachmentAttributesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstanceAttachmentAttributesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeInstanceAttachmentAttributesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeInstanceAttachmentAttributesRequest) SetInstanceIds(v string) *DescribeInstanceAttachmentAttributesRequest {
	s.InstanceIds = &v
	return s
}

func (s *DescribeInstanceAttachmentAttributesRequest) SetOwnerAccount(v string) *DescribeInstanceAttachmentAttributesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeInstanceAttachmentAttributesRequest) SetOwnerId(v int64) *DescribeInstanceAttachmentAttributesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeInstanceAttachmentAttributesRequest) SetPageNumber(v int32) *DescribeInstanceAttachmentAttributesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstanceAttachmentAttributesRequest) SetPageSize(v int32) *DescribeInstanceAttachmentAttributesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstanceAttachmentAttributesRequest) SetRegionId(v string) *DescribeInstanceAttachmentAttributesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceAttachmentAttributesRequest) SetResourceOwnerAccount(v string) *DescribeInstanceAttachmentAttributesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeInstanceAttachmentAttributesRequest) SetResourceOwnerId(v int64) *DescribeInstanceAttachmentAttributesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeInstanceAttachmentAttributesRequest) Validate() error {
	return dara.Validate(s)
}
