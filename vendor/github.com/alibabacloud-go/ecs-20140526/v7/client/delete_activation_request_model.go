// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteActivationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActivationId(v string) *DeleteActivationRequest
	GetActivationId() *string
	SetOwnerAccount(v string) *DeleteActivationRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteActivationRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteActivationRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteActivationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteActivationRequest
	GetResourceOwnerId() *int64
}

type DeleteActivationRequest struct {
	// The ID of the unused activation code.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4ECEEE12-56F1-4FBC-9AB1-890F1234****
	ActivationId *string `json:"ActivationId,omitempty" xml:"ActivationId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the activation code. Supported regions: China (Qingdao), China (Beijing), China (Zhangjiakou), China (Hohhot), China (Ulanqab), China (Hangzhou), China (Shanghai), China (Shenzhen), China (Heyuan), China (Guangzhou), China (Chengdu), China (Hong Kong), Singapore, Japan (Tokyo), US (Silicon Valley), and US (Virginia).
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
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

func (s DeleteActivationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteActivationRequest) GoString() string {
	return s.String()
}

func (s *DeleteActivationRequest) GetActivationId() *string {
	return s.ActivationId
}

func (s *DeleteActivationRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteActivationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteActivationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteActivationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteActivationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteActivationRequest) SetActivationId(v string) *DeleteActivationRequest {
	s.ActivationId = &v
	return s
}

func (s *DeleteActivationRequest) SetOwnerAccount(v string) *DeleteActivationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteActivationRequest) SetOwnerId(v int64) *DeleteActivationRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteActivationRequest) SetRegionId(v string) *DeleteActivationRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteActivationRequest) SetResourceOwnerAccount(v string) *DeleteActivationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteActivationRequest) SetResourceOwnerId(v int64) *DeleteActivationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteActivationRequest) Validate() error {
	return dara.Validate(s)
}
