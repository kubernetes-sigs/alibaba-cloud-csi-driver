// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRouterInterfaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessPointId(v string) *CreateRouterInterfaceRequest
	GetAccessPointId() *string
	SetAutoPay(v bool) *CreateRouterInterfaceRequest
	GetAutoPay() *bool
	SetClientToken(v string) *CreateRouterInterfaceRequest
	GetClientToken() *string
	SetDescription(v string) *CreateRouterInterfaceRequest
	GetDescription() *string
	SetHealthCheckSourceIp(v string) *CreateRouterInterfaceRequest
	GetHealthCheckSourceIp() *string
	SetHealthCheckTargetIp(v string) *CreateRouterInterfaceRequest
	GetHealthCheckTargetIp() *string
	SetInstanceChargeType(v string) *CreateRouterInterfaceRequest
	GetInstanceChargeType() *string
	SetName(v string) *CreateRouterInterfaceRequest
	GetName() *string
	SetOppositeAccessPointId(v string) *CreateRouterInterfaceRequest
	GetOppositeAccessPointId() *string
	SetOppositeInterfaceId(v string) *CreateRouterInterfaceRequest
	GetOppositeInterfaceId() *string
	SetOppositeInterfaceOwnerId(v string) *CreateRouterInterfaceRequest
	GetOppositeInterfaceOwnerId() *string
	SetOppositeRegionId(v string) *CreateRouterInterfaceRequest
	GetOppositeRegionId() *string
	SetOppositeRouterId(v string) *CreateRouterInterfaceRequest
	GetOppositeRouterId() *string
	SetOppositeRouterType(v string) *CreateRouterInterfaceRequest
	GetOppositeRouterType() *string
	SetOwnerAccount(v string) *CreateRouterInterfaceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateRouterInterfaceRequest
	GetOwnerId() *int64
	SetPeriod(v int32) *CreateRouterInterfaceRequest
	GetPeriod() *int32
	SetPricingCycle(v string) *CreateRouterInterfaceRequest
	GetPricingCycle() *string
	SetRegionId(v string) *CreateRouterInterfaceRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateRouterInterfaceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateRouterInterfaceRequest
	GetResourceOwnerId() *int64
	SetRole(v string) *CreateRouterInterfaceRequest
	GetRole() *string
	SetRouterId(v string) *CreateRouterInterfaceRequest
	GetRouterId() *string
	SetRouterType(v string) *CreateRouterInterfaceRequest
	GetRouterType() *string
	SetSpec(v string) *CreateRouterInterfaceRequest
	GetSpec() *string
	SetUserCidr(v string) *CreateRouterInterfaceRequest
	GetUserCidr() *string
}

type CreateRouterInterfaceRequest struct {
	AccessPointId            *string `json:"AccessPointId,omitempty" xml:"AccessPointId,omitempty"`
	AutoPay                  *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	ClientToken              *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description              *string `json:"Description,omitempty" xml:"Description,omitempty"`
	HealthCheckSourceIp      *string `json:"HealthCheckSourceIp,omitempty" xml:"HealthCheckSourceIp,omitempty"`
	HealthCheckTargetIp      *string `json:"HealthCheckTargetIp,omitempty" xml:"HealthCheckTargetIp,omitempty"`
	InstanceChargeType       *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	Name                     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OppositeAccessPointId    *string `json:"OppositeAccessPointId,omitempty" xml:"OppositeAccessPointId,omitempty"`
	OppositeInterfaceId      *string `json:"OppositeInterfaceId,omitempty" xml:"OppositeInterfaceId,omitempty"`
	OppositeInterfaceOwnerId *string `json:"OppositeInterfaceOwnerId,omitempty" xml:"OppositeInterfaceOwnerId,omitempty"`
	// This parameter is required.
	OppositeRegionId   *string `json:"OppositeRegionId,omitempty" xml:"OppositeRegionId,omitempty"`
	OppositeRouterId   *string `json:"OppositeRouterId,omitempty" xml:"OppositeRouterId,omitempty"`
	OppositeRouterType *string `json:"OppositeRouterType,omitempty" xml:"OppositeRouterType,omitempty"`
	OwnerAccount       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Period             *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	PricingCycle       *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// This parameter is required.
	RouterId *string `json:"RouterId,omitempty" xml:"RouterId,omitempty"`
	// This parameter is required.
	RouterType *string `json:"RouterType,omitempty" xml:"RouterType,omitempty"`
	// This parameter is required.
	Spec     *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	UserCidr *string `json:"UserCidr,omitempty" xml:"UserCidr,omitempty"`
}

func (s CreateRouterInterfaceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRouterInterfaceRequest) GoString() string {
	return s.String()
}

func (s *CreateRouterInterfaceRequest) GetAccessPointId() *string {
	return s.AccessPointId
}

func (s *CreateRouterInterfaceRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateRouterInterfaceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateRouterInterfaceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateRouterInterfaceRequest) GetHealthCheckSourceIp() *string {
	return s.HealthCheckSourceIp
}

func (s *CreateRouterInterfaceRequest) GetHealthCheckTargetIp() *string {
	return s.HealthCheckTargetIp
}

func (s *CreateRouterInterfaceRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *CreateRouterInterfaceRequest) GetName() *string {
	return s.Name
}

func (s *CreateRouterInterfaceRequest) GetOppositeAccessPointId() *string {
	return s.OppositeAccessPointId
}

func (s *CreateRouterInterfaceRequest) GetOppositeInterfaceId() *string {
	return s.OppositeInterfaceId
}

func (s *CreateRouterInterfaceRequest) GetOppositeInterfaceOwnerId() *string {
	return s.OppositeInterfaceOwnerId
}

func (s *CreateRouterInterfaceRequest) GetOppositeRegionId() *string {
	return s.OppositeRegionId
}

func (s *CreateRouterInterfaceRequest) GetOppositeRouterId() *string {
	return s.OppositeRouterId
}

func (s *CreateRouterInterfaceRequest) GetOppositeRouterType() *string {
	return s.OppositeRouterType
}

func (s *CreateRouterInterfaceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateRouterInterfaceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateRouterInterfaceRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateRouterInterfaceRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *CreateRouterInterfaceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateRouterInterfaceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateRouterInterfaceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateRouterInterfaceRequest) GetRole() *string {
	return s.Role
}

func (s *CreateRouterInterfaceRequest) GetRouterId() *string {
	return s.RouterId
}

func (s *CreateRouterInterfaceRequest) GetRouterType() *string {
	return s.RouterType
}

func (s *CreateRouterInterfaceRequest) GetSpec() *string {
	return s.Spec
}

func (s *CreateRouterInterfaceRequest) GetUserCidr() *string {
	return s.UserCidr
}

func (s *CreateRouterInterfaceRequest) SetAccessPointId(v string) *CreateRouterInterfaceRequest {
	s.AccessPointId = &v
	return s
}

func (s *CreateRouterInterfaceRequest) SetAutoPay(v bool) *CreateRouterInterfaceRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateRouterInterfaceRequest) SetClientToken(v string) *CreateRouterInterfaceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateRouterInterfaceRequest) SetDescription(v string) *CreateRouterInterfaceRequest {
	s.Description = &v
	return s
}

func (s *CreateRouterInterfaceRequest) SetHealthCheckSourceIp(v string) *CreateRouterInterfaceRequest {
	s.HealthCheckSourceIp = &v
	return s
}

func (s *CreateRouterInterfaceRequest) SetHealthCheckTargetIp(v string) *CreateRouterInterfaceRequest {
	s.HealthCheckTargetIp = &v
	return s
}

func (s *CreateRouterInterfaceRequest) SetInstanceChargeType(v string) *CreateRouterInterfaceRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *CreateRouterInterfaceRequest) SetName(v string) *CreateRouterInterfaceRequest {
	s.Name = &v
	return s
}

func (s *CreateRouterInterfaceRequest) SetOppositeAccessPointId(v string) *CreateRouterInterfaceRequest {
	s.OppositeAccessPointId = &v
	return s
}

func (s *CreateRouterInterfaceRequest) SetOppositeInterfaceId(v string) *CreateRouterInterfaceRequest {
	s.OppositeInterfaceId = &v
	return s
}

func (s *CreateRouterInterfaceRequest) SetOppositeInterfaceOwnerId(v string) *CreateRouterInterfaceRequest {
	s.OppositeInterfaceOwnerId = &v
	return s
}

func (s *CreateRouterInterfaceRequest) SetOppositeRegionId(v string) *CreateRouterInterfaceRequest {
	s.OppositeRegionId = &v
	return s
}

func (s *CreateRouterInterfaceRequest) SetOppositeRouterId(v string) *CreateRouterInterfaceRequest {
	s.OppositeRouterId = &v
	return s
}

func (s *CreateRouterInterfaceRequest) SetOppositeRouterType(v string) *CreateRouterInterfaceRequest {
	s.OppositeRouterType = &v
	return s
}

func (s *CreateRouterInterfaceRequest) SetOwnerAccount(v string) *CreateRouterInterfaceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateRouterInterfaceRequest) SetOwnerId(v int64) *CreateRouterInterfaceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateRouterInterfaceRequest) SetPeriod(v int32) *CreateRouterInterfaceRequest {
	s.Period = &v
	return s
}

func (s *CreateRouterInterfaceRequest) SetPricingCycle(v string) *CreateRouterInterfaceRequest {
	s.PricingCycle = &v
	return s
}

func (s *CreateRouterInterfaceRequest) SetRegionId(v string) *CreateRouterInterfaceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateRouterInterfaceRequest) SetResourceOwnerAccount(v string) *CreateRouterInterfaceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateRouterInterfaceRequest) SetResourceOwnerId(v int64) *CreateRouterInterfaceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateRouterInterfaceRequest) SetRole(v string) *CreateRouterInterfaceRequest {
	s.Role = &v
	return s
}

func (s *CreateRouterInterfaceRequest) SetRouterId(v string) *CreateRouterInterfaceRequest {
	s.RouterId = &v
	return s
}

func (s *CreateRouterInterfaceRequest) SetRouterType(v string) *CreateRouterInterfaceRequest {
	s.RouterType = &v
	return s
}

func (s *CreateRouterInterfaceRequest) SetSpec(v string) *CreateRouterInterfaceRequest {
	s.Spec = &v
	return s
}

func (s *CreateRouterInterfaceRequest) SetUserCidr(v string) *CreateRouterInterfaceRequest {
	s.UserCidr = &v
	return s
}

func (s *CreateRouterInterfaceRequest) Validate() error {
	return dara.Validate(s)
}
