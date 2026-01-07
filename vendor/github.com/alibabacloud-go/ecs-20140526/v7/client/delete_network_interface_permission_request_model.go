// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNetworkInterfacePermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForce(v bool) *DeleteNetworkInterfacePermissionRequest
	GetForce() *bool
	SetNetworkInterfacePermissionId(v string) *DeleteNetworkInterfacePermissionRequest
	GetNetworkInterfacePermissionId() *string
	SetOwnerAccount(v string) *DeleteNetworkInterfacePermissionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteNetworkInterfacePermissionRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteNetworkInterfacePermissionRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteNetworkInterfacePermissionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteNetworkInterfacePermissionRequest
	GetResourceOwnerId() *int64
}

type DeleteNetworkInterfacePermissionRequest struct {
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// This parameter is required.
	NetworkInterfacePermissionId *string `json:"NetworkInterfacePermissionId,omitempty" xml:"NetworkInterfacePermissionId,omitempty"`
	OwnerAccount                 *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteNetworkInterfacePermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNetworkInterfacePermissionRequest) GoString() string {
	return s.String()
}

func (s *DeleteNetworkInterfacePermissionRequest) GetForce() *bool {
	return s.Force
}

func (s *DeleteNetworkInterfacePermissionRequest) GetNetworkInterfacePermissionId() *string {
	return s.NetworkInterfacePermissionId
}

func (s *DeleteNetworkInterfacePermissionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteNetworkInterfacePermissionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteNetworkInterfacePermissionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteNetworkInterfacePermissionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteNetworkInterfacePermissionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteNetworkInterfacePermissionRequest) SetForce(v bool) *DeleteNetworkInterfacePermissionRequest {
	s.Force = &v
	return s
}

func (s *DeleteNetworkInterfacePermissionRequest) SetNetworkInterfacePermissionId(v string) *DeleteNetworkInterfacePermissionRequest {
	s.NetworkInterfacePermissionId = &v
	return s
}

func (s *DeleteNetworkInterfacePermissionRequest) SetOwnerAccount(v string) *DeleteNetworkInterfacePermissionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteNetworkInterfacePermissionRequest) SetOwnerId(v int64) *DeleteNetworkInterfacePermissionRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteNetworkInterfacePermissionRequest) SetRegionId(v string) *DeleteNetworkInterfacePermissionRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteNetworkInterfacePermissionRequest) SetResourceOwnerAccount(v string) *DeleteNetworkInterfacePermissionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteNetworkInterfacePermissionRequest) SetResourceOwnerId(v int64) *DeleteNetworkInterfacePermissionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteNetworkInterfacePermissionRequest) Validate() error {
	return dara.Validate(s)
}
