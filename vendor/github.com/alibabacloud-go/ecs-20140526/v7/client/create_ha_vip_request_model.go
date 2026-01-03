// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHaVipRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateHaVipRequest
	GetClientToken() *string
	SetDescription(v string) *CreateHaVipRequest
	GetDescription() *string
	SetIpAddress(v string) *CreateHaVipRequest
	GetIpAddress() *string
	SetOwnerAccount(v string) *CreateHaVipRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateHaVipRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateHaVipRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateHaVipRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateHaVipRequest
	GetResourceOwnerId() *int64
	SetVSwitchId(v string) *CreateHaVipRequest
	GetVSwitchId() *string
}

type CreateHaVipRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	IpAddress    *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s CreateHaVipRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHaVipRequest) GoString() string {
	return s.String()
}

func (s *CreateHaVipRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateHaVipRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateHaVipRequest) GetIpAddress() *string {
	return s.IpAddress
}

func (s *CreateHaVipRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateHaVipRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateHaVipRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateHaVipRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateHaVipRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateHaVipRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateHaVipRequest) SetClientToken(v string) *CreateHaVipRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateHaVipRequest) SetDescription(v string) *CreateHaVipRequest {
	s.Description = &v
	return s
}

func (s *CreateHaVipRequest) SetIpAddress(v string) *CreateHaVipRequest {
	s.IpAddress = &v
	return s
}

func (s *CreateHaVipRequest) SetOwnerAccount(v string) *CreateHaVipRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateHaVipRequest) SetOwnerId(v int64) *CreateHaVipRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateHaVipRequest) SetRegionId(v string) *CreateHaVipRequest {
	s.RegionId = &v
	return s
}

func (s *CreateHaVipRequest) SetResourceOwnerAccount(v string) *CreateHaVipRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateHaVipRequest) SetResourceOwnerId(v int64) *CreateHaVipRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateHaVipRequest) SetVSwitchId(v string) *CreateHaVipRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateHaVipRequest) Validate() error {
	return dara.Validate(s)
}
