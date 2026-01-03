// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActivateRouterInterfaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *ActivateRouterInterfaceRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ActivateRouterInterfaceRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ActivateRouterInterfaceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ActivateRouterInterfaceRequest
	GetResourceOwnerId() *int64
	SetRouterInterfaceId(v string) *ActivateRouterInterfaceRequest
	GetRouterInterfaceId() *string
}

type ActivateRouterInterfaceRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	RouterInterfaceId *string `json:"RouterInterfaceId,omitempty" xml:"RouterInterfaceId,omitempty"`
}

func (s ActivateRouterInterfaceRequest) String() string {
	return dara.Prettify(s)
}

func (s ActivateRouterInterfaceRequest) GoString() string {
	return s.String()
}

func (s *ActivateRouterInterfaceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ActivateRouterInterfaceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ActivateRouterInterfaceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ActivateRouterInterfaceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ActivateRouterInterfaceRequest) GetRouterInterfaceId() *string {
	return s.RouterInterfaceId
}

func (s *ActivateRouterInterfaceRequest) SetOwnerId(v int64) *ActivateRouterInterfaceRequest {
	s.OwnerId = &v
	return s
}

func (s *ActivateRouterInterfaceRequest) SetRegionId(v string) *ActivateRouterInterfaceRequest {
	s.RegionId = &v
	return s
}

func (s *ActivateRouterInterfaceRequest) SetResourceOwnerAccount(v string) *ActivateRouterInterfaceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ActivateRouterInterfaceRequest) SetResourceOwnerId(v int64) *ActivateRouterInterfaceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ActivateRouterInterfaceRequest) SetRouterInterfaceId(v string) *ActivateRouterInterfaceRequest {
	s.RouterInterfaceId = &v
	return s
}

func (s *ActivateRouterInterfaceRequest) Validate() error {
	return dara.Validate(s)
}
