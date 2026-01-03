// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNatGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNatGatewayId(v string) *DeleteNatGatewayRequest
	GetNatGatewayId() *string
	SetOwnerAccount(v string) *DeleteNatGatewayRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteNatGatewayRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteNatGatewayRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteNatGatewayRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteNatGatewayRequest
	GetResourceOwnerId() *int64
}

type DeleteNatGatewayRequest struct {
	// This parameter is required.
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteNatGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNatGatewayRequest) GoString() string {
	return s.String()
}

func (s *DeleteNatGatewayRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DeleteNatGatewayRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteNatGatewayRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteNatGatewayRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteNatGatewayRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteNatGatewayRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteNatGatewayRequest) SetNatGatewayId(v string) *DeleteNatGatewayRequest {
	s.NatGatewayId = &v
	return s
}

func (s *DeleteNatGatewayRequest) SetOwnerAccount(v string) *DeleteNatGatewayRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteNatGatewayRequest) SetOwnerId(v int64) *DeleteNatGatewayRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteNatGatewayRequest) SetRegionId(v string) *DeleteNatGatewayRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteNatGatewayRequest) SetResourceOwnerAccount(v string) *DeleteNatGatewayRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteNatGatewayRequest) SetResourceOwnerId(v int64) *DeleteNatGatewayRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteNatGatewayRequest) Validate() error {
	return dara.Validate(s)
}
