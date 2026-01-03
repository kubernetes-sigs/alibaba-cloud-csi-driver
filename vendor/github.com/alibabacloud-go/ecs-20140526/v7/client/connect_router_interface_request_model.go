// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConnectRouterInterfaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *ConnectRouterInterfaceRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ConnectRouterInterfaceRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ConnectRouterInterfaceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ConnectRouterInterfaceRequest
	GetResourceOwnerId() *int64
	SetRouterInterfaceId(v string) *ConnectRouterInterfaceRequest
	GetRouterInterfaceId() *string
}

type ConnectRouterInterfaceRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	RouterInterfaceId *string `json:"RouterInterfaceId,omitempty" xml:"RouterInterfaceId,omitempty"`
}

func (s ConnectRouterInterfaceRequest) String() string {
	return dara.Prettify(s)
}

func (s ConnectRouterInterfaceRequest) GoString() string {
	return s.String()
}

func (s *ConnectRouterInterfaceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ConnectRouterInterfaceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ConnectRouterInterfaceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ConnectRouterInterfaceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ConnectRouterInterfaceRequest) GetRouterInterfaceId() *string {
	return s.RouterInterfaceId
}

func (s *ConnectRouterInterfaceRequest) SetOwnerId(v int64) *ConnectRouterInterfaceRequest {
	s.OwnerId = &v
	return s
}

func (s *ConnectRouterInterfaceRequest) SetRegionId(v string) *ConnectRouterInterfaceRequest {
	s.RegionId = &v
	return s
}

func (s *ConnectRouterInterfaceRequest) SetResourceOwnerAccount(v string) *ConnectRouterInterfaceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ConnectRouterInterfaceRequest) SetResourceOwnerId(v int64) *ConnectRouterInterfaceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ConnectRouterInterfaceRequest) SetRouterInterfaceId(v string) *ConnectRouterInterfaceRequest {
	s.RouterInterfaceId = &v
	return s
}

func (s *ConnectRouterInterfaceRequest) Validate() error {
	return dara.Validate(s)
}
