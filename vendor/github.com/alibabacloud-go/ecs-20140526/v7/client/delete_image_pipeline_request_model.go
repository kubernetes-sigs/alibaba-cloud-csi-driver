// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteImagePipelineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImagePipelineId(v string) *DeleteImagePipelineRequest
	GetImagePipelineId() *string
	SetOwnerAccount(v string) *DeleteImagePipelineRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteImagePipelineRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteImagePipelineRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteImagePipelineRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteImagePipelineRequest
	GetResourceOwnerId() *int64
}

type DeleteImagePipelineRequest struct {
	// The ID of the image template.
	//
	// This parameter is required.
	//
	// example:
	//
	// ip-2ze5tsl5bp6nf2b3****
	ImagePipelineId *string `json:"ImagePipelineId,omitempty" xml:"ImagePipelineId,omitempty"`
	OwnerAccount    *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the image template. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
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

func (s DeleteImagePipelineRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteImagePipelineRequest) GoString() string {
	return s.String()
}

func (s *DeleteImagePipelineRequest) GetImagePipelineId() *string {
	return s.ImagePipelineId
}

func (s *DeleteImagePipelineRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteImagePipelineRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteImagePipelineRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteImagePipelineRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteImagePipelineRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteImagePipelineRequest) SetImagePipelineId(v string) *DeleteImagePipelineRequest {
	s.ImagePipelineId = &v
	return s
}

func (s *DeleteImagePipelineRequest) SetOwnerAccount(v string) *DeleteImagePipelineRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteImagePipelineRequest) SetOwnerId(v int64) *DeleteImagePipelineRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteImagePipelineRequest) SetRegionId(v string) *DeleteImagePipelineRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteImagePipelineRequest) SetResourceOwnerAccount(v string) *DeleteImagePipelineRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteImagePipelineRequest) SetResourceOwnerId(v int64) *DeleteImagePipelineRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteImagePipelineRequest) Validate() error {
	return dara.Validate(s)
}
