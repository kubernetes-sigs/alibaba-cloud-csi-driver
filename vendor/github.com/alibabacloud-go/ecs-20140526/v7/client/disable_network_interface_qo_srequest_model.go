// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableNetworkInterfaceQoSRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkInterfaceId(v string) *DisableNetworkInterfaceQoSRequest
	GetNetworkInterfaceId() *string
	SetOwnerAccount(v string) *DisableNetworkInterfaceQoSRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DisableNetworkInterfaceQoSRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DisableNetworkInterfaceQoSRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DisableNetworkInterfaceQoSRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DisableNetworkInterfaceQoSRequest
	GetResourceOwnerId() *int64
}

type DisableNetworkInterfaceQoSRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eni-bp1iqejowblx6h8j****
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	OwnerAccount       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DisableNetworkInterfaceQoSRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableNetworkInterfaceQoSRequest) GoString() string {
	return s.String()
}

func (s *DisableNetworkInterfaceQoSRequest) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *DisableNetworkInterfaceQoSRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DisableNetworkInterfaceQoSRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DisableNetworkInterfaceQoSRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DisableNetworkInterfaceQoSRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DisableNetworkInterfaceQoSRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DisableNetworkInterfaceQoSRequest) SetNetworkInterfaceId(v string) *DisableNetworkInterfaceQoSRequest {
	s.NetworkInterfaceId = &v
	return s
}

func (s *DisableNetworkInterfaceQoSRequest) SetOwnerAccount(v string) *DisableNetworkInterfaceQoSRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DisableNetworkInterfaceQoSRequest) SetOwnerId(v int64) *DisableNetworkInterfaceQoSRequest {
	s.OwnerId = &v
	return s
}

func (s *DisableNetworkInterfaceQoSRequest) SetRegionId(v string) *DisableNetworkInterfaceQoSRequest {
	s.RegionId = &v
	return s
}

func (s *DisableNetworkInterfaceQoSRequest) SetResourceOwnerAccount(v string) *DisableNetworkInterfaceQoSRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DisableNetworkInterfaceQoSRequest) SetResourceOwnerId(v int64) *DisableNetworkInterfaceQoSRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DisableNetworkInterfaceQoSRequest) Validate() error {
	return dara.Validate(s)
}
