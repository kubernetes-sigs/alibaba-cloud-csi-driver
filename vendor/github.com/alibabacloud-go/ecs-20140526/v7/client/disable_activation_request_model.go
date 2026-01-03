// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableActivationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActivationId(v string) *DisableActivationRequest
	GetActivationId() *string
	SetOwnerAccount(v string) *DisableActivationRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DisableActivationRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DisableActivationRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DisableActivationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DisableActivationRequest
	GetResourceOwnerId() *int64
}

type DisableActivationRequest struct {
	// The request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4ECEEE12-56F1-4FBC-9AB1-890F1234****
	ActivationId *string `json:"ActivationId,omitempty" xml:"ActivationId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. Supported regions: China (Qingdao), China (Beijing), China (Zhangjiakou), China (Hohhot), China (Ulanqab), China (Hangzhou), China (Shanghai), China (Shenzhen), China (Heyuan), China (Guangzhou), China (Chengdu), China (Hong Kong), Singapore, Japan (Tokyo), US (Silicon Valley), and US (Virginia).
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

func (s DisableActivationRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableActivationRequest) GoString() string {
	return s.String()
}

func (s *DisableActivationRequest) GetActivationId() *string {
	return s.ActivationId
}

func (s *DisableActivationRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DisableActivationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DisableActivationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DisableActivationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DisableActivationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DisableActivationRequest) SetActivationId(v string) *DisableActivationRequest {
	s.ActivationId = &v
	return s
}

func (s *DisableActivationRequest) SetOwnerAccount(v string) *DisableActivationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DisableActivationRequest) SetOwnerId(v int64) *DisableActivationRequest {
	s.OwnerId = &v
	return s
}

func (s *DisableActivationRequest) SetRegionId(v string) *DisableActivationRequest {
	s.RegionId = &v
	return s
}

func (s *DisableActivationRequest) SetResourceOwnerAccount(v string) *DisableActivationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DisableActivationRequest) SetResourceOwnerId(v int64) *DisableActivationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DisableActivationRequest) Validate() error {
	return dara.Validate(s)
}
