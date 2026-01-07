// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVirtualBorderRouterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteVirtualBorderRouterRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *DeleteVirtualBorderRouterRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteVirtualBorderRouterRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteVirtualBorderRouterRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteVirtualBorderRouterRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteVirtualBorderRouterRequest
	GetResourceOwnerId() *int64
	SetUserCidr(v string) *DeleteVirtualBorderRouterRequest
	GetUserCidr() *string
	SetVbrId(v string) *DeleteVirtualBorderRouterRequest
	GetVbrId() *string
}

type DeleteVirtualBorderRouterRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	UserCidr             *string `json:"UserCidr,omitempty" xml:"UserCidr,omitempty"`
	// This parameter is required.
	VbrId *string `json:"VbrId,omitempty" xml:"VbrId,omitempty"`
}

func (s DeleteVirtualBorderRouterRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVirtualBorderRouterRequest) GoString() string {
	return s.String()
}

func (s *DeleteVirtualBorderRouterRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteVirtualBorderRouterRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteVirtualBorderRouterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteVirtualBorderRouterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteVirtualBorderRouterRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteVirtualBorderRouterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteVirtualBorderRouterRequest) GetUserCidr() *string {
	return s.UserCidr
}

func (s *DeleteVirtualBorderRouterRequest) GetVbrId() *string {
	return s.VbrId
}

func (s *DeleteVirtualBorderRouterRequest) SetClientToken(v string) *DeleteVirtualBorderRouterRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteVirtualBorderRouterRequest) SetOwnerAccount(v string) *DeleteVirtualBorderRouterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteVirtualBorderRouterRequest) SetOwnerId(v int64) *DeleteVirtualBorderRouterRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteVirtualBorderRouterRequest) SetRegionId(v string) *DeleteVirtualBorderRouterRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteVirtualBorderRouterRequest) SetResourceOwnerAccount(v string) *DeleteVirtualBorderRouterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteVirtualBorderRouterRequest) SetResourceOwnerId(v int64) *DeleteVirtualBorderRouterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteVirtualBorderRouterRequest) SetUserCidr(v string) *DeleteVirtualBorderRouterRequest {
	s.UserCidr = &v
	return s
}

func (s *DeleteVirtualBorderRouterRequest) SetVbrId(v string) *DeleteVirtualBorderRouterRequest {
	s.VbrId = &v
	return s
}

func (s *DeleteVirtualBorderRouterRequest) Validate() error {
	return dara.Validate(s)
}
