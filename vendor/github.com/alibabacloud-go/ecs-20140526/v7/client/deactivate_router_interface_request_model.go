// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeactivateRouterInterfaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DeactivateRouterInterfaceRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeactivateRouterInterfaceRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeactivateRouterInterfaceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeactivateRouterInterfaceRequest
	GetResourceOwnerId() *int64
	SetRouterInterfaceId(v string) *DeactivateRouterInterfaceRequest
	GetRouterInterfaceId() *string
}

type DeactivateRouterInterfaceRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	RouterInterfaceId *string `json:"RouterInterfaceId,omitempty" xml:"RouterInterfaceId,omitempty"`
}

func (s DeactivateRouterInterfaceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeactivateRouterInterfaceRequest) GoString() string {
	return s.String()
}

func (s *DeactivateRouterInterfaceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeactivateRouterInterfaceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeactivateRouterInterfaceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeactivateRouterInterfaceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeactivateRouterInterfaceRequest) GetRouterInterfaceId() *string {
	return s.RouterInterfaceId
}

func (s *DeactivateRouterInterfaceRequest) SetOwnerId(v int64) *DeactivateRouterInterfaceRequest {
	s.OwnerId = &v
	return s
}

func (s *DeactivateRouterInterfaceRequest) SetRegionId(v string) *DeactivateRouterInterfaceRequest {
	s.RegionId = &v
	return s
}

func (s *DeactivateRouterInterfaceRequest) SetResourceOwnerAccount(v string) *DeactivateRouterInterfaceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeactivateRouterInterfaceRequest) SetResourceOwnerId(v int64) *DeactivateRouterInterfaceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeactivateRouterInterfaceRequest) SetRouterInterfaceId(v string) *DeactivateRouterInterfaceRequest {
	s.RouterInterfaceId = &v
	return s
}

func (s *DeactivateRouterInterfaceRequest) Validate() error {
	return dara.Validate(s)
}
