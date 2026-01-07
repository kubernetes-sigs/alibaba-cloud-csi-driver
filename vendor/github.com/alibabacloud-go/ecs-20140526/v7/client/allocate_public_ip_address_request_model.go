// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocatePublicIpAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *AllocatePublicIpAddressRequest
	GetInstanceId() *string
	SetIpAddress(v string) *AllocatePublicIpAddressRequest
	GetIpAddress() *string
	SetOwnerAccount(v string) *AllocatePublicIpAddressRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AllocatePublicIpAddressRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *AllocatePublicIpAddressRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AllocatePublicIpAddressRequest
	GetResourceOwnerId() *int64
	SetVlanId(v string) *AllocatePublicIpAddressRequest
	GetVlanId() *string
}

type AllocatePublicIpAddressRequest struct {
	// The ID of the instance to which you want to assign a public IP address.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp1gtjxuuvwj17zr****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The static public IP address that you want to assign to the instance. This parameter is empty by default, which indicates that a static public IP address is randomly assigned by the system.
	//
	// >  Only users in the whitelist can specify this parameter.
	//
	// example:
	//
	// ``112.124.**.**``
	IpAddress            *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The virtual LAN (VLAN) ID of the instance.
	//
	// > This parameter will be removed in the future. To ensure future compatibility, we recommend that you use other parameters.
	//
	// example:
	//
	// 720
	VlanId *string `json:"VlanId,omitempty" xml:"VlanId,omitempty"`
}

func (s AllocatePublicIpAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s AllocatePublicIpAddressRequest) GoString() string {
	return s.String()
}

func (s *AllocatePublicIpAddressRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AllocatePublicIpAddressRequest) GetIpAddress() *string {
	return s.IpAddress
}

func (s *AllocatePublicIpAddressRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AllocatePublicIpAddressRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AllocatePublicIpAddressRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AllocatePublicIpAddressRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AllocatePublicIpAddressRequest) GetVlanId() *string {
	return s.VlanId
}

func (s *AllocatePublicIpAddressRequest) SetInstanceId(v string) *AllocatePublicIpAddressRequest {
	s.InstanceId = &v
	return s
}

func (s *AllocatePublicIpAddressRequest) SetIpAddress(v string) *AllocatePublicIpAddressRequest {
	s.IpAddress = &v
	return s
}

func (s *AllocatePublicIpAddressRequest) SetOwnerAccount(v string) *AllocatePublicIpAddressRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AllocatePublicIpAddressRequest) SetOwnerId(v int64) *AllocatePublicIpAddressRequest {
	s.OwnerId = &v
	return s
}

func (s *AllocatePublicIpAddressRequest) SetResourceOwnerAccount(v string) *AllocatePublicIpAddressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AllocatePublicIpAddressRequest) SetResourceOwnerId(v int64) *AllocatePublicIpAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AllocatePublicIpAddressRequest) SetVlanId(v string) *AllocatePublicIpAddressRequest {
	s.VlanId = &v
	return s
}

func (s *AllocatePublicIpAddressRequest) Validate() error {
	return dara.Validate(s)
}
