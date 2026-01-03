// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteImageComponentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageComponentId(v string) *DeleteImageComponentRequest
	GetImageComponentId() *string
	SetOwnerAccount(v string) *DeleteImageComponentRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteImageComponentRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteImageComponentRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteImageComponentRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteImageComponentRequest
	GetResourceOwnerId() *int64
}

type DeleteImageComponentRequest struct {
	// The ID of the image component.
	//
	// This parameter is required.
	//
	// example:
	//
	// ic-bp67acfmxazb4p****
	ImageComponentId *string `json:"ImageComponentId,omitempty" xml:"ImageComponentId,omitempty"`
	OwnerAccount     *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the image component. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
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

func (s DeleteImageComponentRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteImageComponentRequest) GoString() string {
	return s.String()
}

func (s *DeleteImageComponentRequest) GetImageComponentId() *string {
	return s.ImageComponentId
}

func (s *DeleteImageComponentRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteImageComponentRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteImageComponentRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteImageComponentRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteImageComponentRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteImageComponentRequest) SetImageComponentId(v string) *DeleteImageComponentRequest {
	s.ImageComponentId = &v
	return s
}

func (s *DeleteImageComponentRequest) SetOwnerAccount(v string) *DeleteImageComponentRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteImageComponentRequest) SetOwnerId(v int64) *DeleteImageComponentRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteImageComponentRequest) SetRegionId(v string) *DeleteImageComponentRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteImageComponentRequest) SetResourceOwnerAccount(v string) *DeleteImageComponentRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteImageComponentRequest) SetResourceOwnerId(v int64) *DeleteImageComponentRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteImageComponentRequest) Validate() error {
	return dara.Validate(s)
}
