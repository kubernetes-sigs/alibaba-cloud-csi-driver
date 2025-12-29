// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTerminateVirtualBorderRouterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *TerminateVirtualBorderRouterRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *TerminateVirtualBorderRouterRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *TerminateVirtualBorderRouterRequest
	GetOwnerId() *int64
	SetRegionId(v string) *TerminateVirtualBorderRouterRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *TerminateVirtualBorderRouterRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *TerminateVirtualBorderRouterRequest
	GetResourceOwnerId() *int64
	SetUserCidr(v string) *TerminateVirtualBorderRouterRequest
	GetUserCidr() *string
	SetVbrId(v string) *TerminateVirtualBorderRouterRequest
	GetVbrId() *string
}

type TerminateVirtualBorderRouterRequest struct {
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

func (s TerminateVirtualBorderRouterRequest) String() string {
	return dara.Prettify(s)
}

func (s TerminateVirtualBorderRouterRequest) GoString() string {
	return s.String()
}

func (s *TerminateVirtualBorderRouterRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *TerminateVirtualBorderRouterRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *TerminateVirtualBorderRouterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *TerminateVirtualBorderRouterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *TerminateVirtualBorderRouterRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *TerminateVirtualBorderRouterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *TerminateVirtualBorderRouterRequest) GetUserCidr() *string {
	return s.UserCidr
}

func (s *TerminateVirtualBorderRouterRequest) GetVbrId() *string {
	return s.VbrId
}

func (s *TerminateVirtualBorderRouterRequest) SetClientToken(v string) *TerminateVirtualBorderRouterRequest {
	s.ClientToken = &v
	return s
}

func (s *TerminateVirtualBorderRouterRequest) SetOwnerAccount(v string) *TerminateVirtualBorderRouterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *TerminateVirtualBorderRouterRequest) SetOwnerId(v int64) *TerminateVirtualBorderRouterRequest {
	s.OwnerId = &v
	return s
}

func (s *TerminateVirtualBorderRouterRequest) SetRegionId(v string) *TerminateVirtualBorderRouterRequest {
	s.RegionId = &v
	return s
}

func (s *TerminateVirtualBorderRouterRequest) SetResourceOwnerAccount(v string) *TerminateVirtualBorderRouterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TerminateVirtualBorderRouterRequest) SetResourceOwnerId(v int64) *TerminateVirtualBorderRouterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TerminateVirtualBorderRouterRequest) SetUserCidr(v string) *TerminateVirtualBorderRouterRequest {
	s.UserCidr = &v
	return s
}

func (s *TerminateVirtualBorderRouterRequest) SetVbrId(v string) *TerminateVirtualBorderRouterRequest {
	s.VbrId = &v
	return s
}

func (s *TerminateVirtualBorderRouterRequest) Validate() error {
	return dara.Validate(s)
}
