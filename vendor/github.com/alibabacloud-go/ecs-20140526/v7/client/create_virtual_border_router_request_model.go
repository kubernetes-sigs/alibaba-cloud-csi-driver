// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVirtualBorderRouterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCircuitCode(v string) *CreateVirtualBorderRouterRequest
	GetCircuitCode() *string
	SetClientToken(v string) *CreateVirtualBorderRouterRequest
	GetClientToken() *string
	SetDescription(v string) *CreateVirtualBorderRouterRequest
	GetDescription() *string
	SetLocalGatewayIp(v string) *CreateVirtualBorderRouterRequest
	GetLocalGatewayIp() *string
	SetName(v string) *CreateVirtualBorderRouterRequest
	GetName() *string
	SetOwnerAccount(v string) *CreateVirtualBorderRouterRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateVirtualBorderRouterRequest
	GetOwnerId() *int64
	SetPeerGatewayIp(v string) *CreateVirtualBorderRouterRequest
	GetPeerGatewayIp() *string
	SetPeeringSubnetMask(v string) *CreateVirtualBorderRouterRequest
	GetPeeringSubnetMask() *string
	SetPhysicalConnectionId(v string) *CreateVirtualBorderRouterRequest
	GetPhysicalConnectionId() *string
	SetRegionId(v string) *CreateVirtualBorderRouterRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateVirtualBorderRouterRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateVirtualBorderRouterRequest
	GetResourceOwnerId() *int64
	SetUserCidr(v string) *CreateVirtualBorderRouterRequest
	GetUserCidr() *string
	SetVbrOwnerId(v int64) *CreateVirtualBorderRouterRequest
	GetVbrOwnerId() *int64
	SetVlanId(v int32) *CreateVirtualBorderRouterRequest
	GetVlanId() *int32
}

type CreateVirtualBorderRouterRequest struct {
	CircuitCode       *string `json:"CircuitCode,omitempty" xml:"CircuitCode,omitempty"`
	ClientToken       *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	LocalGatewayIp    *string `json:"LocalGatewayIp,omitempty" xml:"LocalGatewayIp,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount      *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PeerGatewayIp     *string `json:"PeerGatewayIp,omitempty" xml:"PeerGatewayIp,omitempty"`
	PeeringSubnetMask *string `json:"PeeringSubnetMask,omitempty" xml:"PeeringSubnetMask,omitempty"`
	// This parameter is required.
	PhysicalConnectionId *string `json:"PhysicalConnectionId,omitempty" xml:"PhysicalConnectionId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	UserCidr             *string `json:"UserCidr,omitempty" xml:"UserCidr,omitempty"`
	VbrOwnerId           *int64  `json:"VbrOwnerId,omitempty" xml:"VbrOwnerId,omitempty"`
	// This parameter is required.
	VlanId *int32 `json:"VlanId,omitempty" xml:"VlanId,omitempty"`
}

func (s CreateVirtualBorderRouterRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVirtualBorderRouterRequest) GoString() string {
	return s.String()
}

func (s *CreateVirtualBorderRouterRequest) GetCircuitCode() *string {
	return s.CircuitCode
}

func (s *CreateVirtualBorderRouterRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateVirtualBorderRouterRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateVirtualBorderRouterRequest) GetLocalGatewayIp() *string {
	return s.LocalGatewayIp
}

func (s *CreateVirtualBorderRouterRequest) GetName() *string {
	return s.Name
}

func (s *CreateVirtualBorderRouterRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateVirtualBorderRouterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateVirtualBorderRouterRequest) GetPeerGatewayIp() *string {
	return s.PeerGatewayIp
}

func (s *CreateVirtualBorderRouterRequest) GetPeeringSubnetMask() *string {
	return s.PeeringSubnetMask
}

func (s *CreateVirtualBorderRouterRequest) GetPhysicalConnectionId() *string {
	return s.PhysicalConnectionId
}

func (s *CreateVirtualBorderRouterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateVirtualBorderRouterRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateVirtualBorderRouterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateVirtualBorderRouterRequest) GetUserCidr() *string {
	return s.UserCidr
}

func (s *CreateVirtualBorderRouterRequest) GetVbrOwnerId() *int64 {
	return s.VbrOwnerId
}

func (s *CreateVirtualBorderRouterRequest) GetVlanId() *int32 {
	return s.VlanId
}

func (s *CreateVirtualBorderRouterRequest) SetCircuitCode(v string) *CreateVirtualBorderRouterRequest {
	s.CircuitCode = &v
	return s
}

func (s *CreateVirtualBorderRouterRequest) SetClientToken(v string) *CreateVirtualBorderRouterRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateVirtualBorderRouterRequest) SetDescription(v string) *CreateVirtualBorderRouterRequest {
	s.Description = &v
	return s
}

func (s *CreateVirtualBorderRouterRequest) SetLocalGatewayIp(v string) *CreateVirtualBorderRouterRequest {
	s.LocalGatewayIp = &v
	return s
}

func (s *CreateVirtualBorderRouterRequest) SetName(v string) *CreateVirtualBorderRouterRequest {
	s.Name = &v
	return s
}

func (s *CreateVirtualBorderRouterRequest) SetOwnerAccount(v string) *CreateVirtualBorderRouterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateVirtualBorderRouterRequest) SetOwnerId(v int64) *CreateVirtualBorderRouterRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateVirtualBorderRouterRequest) SetPeerGatewayIp(v string) *CreateVirtualBorderRouterRequest {
	s.PeerGatewayIp = &v
	return s
}

func (s *CreateVirtualBorderRouterRequest) SetPeeringSubnetMask(v string) *CreateVirtualBorderRouterRequest {
	s.PeeringSubnetMask = &v
	return s
}

func (s *CreateVirtualBorderRouterRequest) SetPhysicalConnectionId(v string) *CreateVirtualBorderRouterRequest {
	s.PhysicalConnectionId = &v
	return s
}

func (s *CreateVirtualBorderRouterRequest) SetRegionId(v string) *CreateVirtualBorderRouterRequest {
	s.RegionId = &v
	return s
}

func (s *CreateVirtualBorderRouterRequest) SetResourceOwnerAccount(v string) *CreateVirtualBorderRouterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateVirtualBorderRouterRequest) SetResourceOwnerId(v int64) *CreateVirtualBorderRouterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateVirtualBorderRouterRequest) SetUserCidr(v string) *CreateVirtualBorderRouterRequest {
	s.UserCidr = &v
	return s
}

func (s *CreateVirtualBorderRouterRequest) SetVbrOwnerId(v int64) *CreateVirtualBorderRouterRequest {
	s.VbrOwnerId = &v
	return s
}

func (s *CreateVirtualBorderRouterRequest) SetVlanId(v int32) *CreateVirtualBorderRouterRequest {
	s.VlanId = &v
	return s
}

func (s *CreateVirtualBorderRouterRequest) Validate() error {
	return dara.Validate(s)
}
