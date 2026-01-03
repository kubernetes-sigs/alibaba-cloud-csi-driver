// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRouterInterfaceAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifyRouterInterfaceAttributeRequest
	GetDescription() *string
	SetHealthCheckSourceIp(v string) *ModifyRouterInterfaceAttributeRequest
	GetHealthCheckSourceIp() *string
	SetHealthCheckTargetIp(v string) *ModifyRouterInterfaceAttributeRequest
	GetHealthCheckTargetIp() *string
	SetName(v string) *ModifyRouterInterfaceAttributeRequest
	GetName() *string
	SetOppositeInterfaceId(v string) *ModifyRouterInterfaceAttributeRequest
	GetOppositeInterfaceId() *string
	SetOppositeInterfaceOwnerId(v int64) *ModifyRouterInterfaceAttributeRequest
	GetOppositeInterfaceOwnerId() *int64
	SetOppositeRouterId(v string) *ModifyRouterInterfaceAttributeRequest
	GetOppositeRouterId() *string
	SetOppositeRouterType(v string) *ModifyRouterInterfaceAttributeRequest
	GetOppositeRouterType() *string
	SetOwnerId(v int64) *ModifyRouterInterfaceAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyRouterInterfaceAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyRouterInterfaceAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyRouterInterfaceAttributeRequest
	GetResourceOwnerId() *int64
	SetRouterInterfaceId(v string) *ModifyRouterInterfaceAttributeRequest
	GetRouterInterfaceId() *string
}

type ModifyRouterInterfaceAttributeRequest struct {
	Description              *string `json:"Description,omitempty" xml:"Description,omitempty"`
	HealthCheckSourceIp      *string `json:"HealthCheckSourceIp,omitempty" xml:"HealthCheckSourceIp,omitempty"`
	HealthCheckTargetIp      *string `json:"HealthCheckTargetIp,omitempty" xml:"HealthCheckTargetIp,omitempty"`
	Name                     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OppositeInterfaceId      *string `json:"OppositeInterfaceId,omitempty" xml:"OppositeInterfaceId,omitempty"`
	OppositeInterfaceOwnerId *int64  `json:"OppositeInterfaceOwnerId,omitempty" xml:"OppositeInterfaceOwnerId,omitempty"`
	OppositeRouterId         *string `json:"OppositeRouterId,omitempty" xml:"OppositeRouterId,omitempty"`
	OppositeRouterType       *string `json:"OppositeRouterType,omitempty" xml:"OppositeRouterType,omitempty"`
	OwnerId                  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	RouterInterfaceId *string `json:"RouterInterfaceId,omitempty" xml:"RouterInterfaceId,omitempty"`
}

func (s ModifyRouterInterfaceAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyRouterInterfaceAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyRouterInterfaceAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyRouterInterfaceAttributeRequest) GetHealthCheckSourceIp() *string {
	return s.HealthCheckSourceIp
}

func (s *ModifyRouterInterfaceAttributeRequest) GetHealthCheckTargetIp() *string {
	return s.HealthCheckTargetIp
}

func (s *ModifyRouterInterfaceAttributeRequest) GetName() *string {
	return s.Name
}

func (s *ModifyRouterInterfaceAttributeRequest) GetOppositeInterfaceId() *string {
	return s.OppositeInterfaceId
}

func (s *ModifyRouterInterfaceAttributeRequest) GetOppositeInterfaceOwnerId() *int64 {
	return s.OppositeInterfaceOwnerId
}

func (s *ModifyRouterInterfaceAttributeRequest) GetOppositeRouterId() *string {
	return s.OppositeRouterId
}

func (s *ModifyRouterInterfaceAttributeRequest) GetOppositeRouterType() *string {
	return s.OppositeRouterType
}

func (s *ModifyRouterInterfaceAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyRouterInterfaceAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyRouterInterfaceAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyRouterInterfaceAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyRouterInterfaceAttributeRequest) GetRouterInterfaceId() *string {
	return s.RouterInterfaceId
}

func (s *ModifyRouterInterfaceAttributeRequest) SetDescription(v string) *ModifyRouterInterfaceAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyRouterInterfaceAttributeRequest) SetHealthCheckSourceIp(v string) *ModifyRouterInterfaceAttributeRequest {
	s.HealthCheckSourceIp = &v
	return s
}

func (s *ModifyRouterInterfaceAttributeRequest) SetHealthCheckTargetIp(v string) *ModifyRouterInterfaceAttributeRequest {
	s.HealthCheckTargetIp = &v
	return s
}

func (s *ModifyRouterInterfaceAttributeRequest) SetName(v string) *ModifyRouterInterfaceAttributeRequest {
	s.Name = &v
	return s
}

func (s *ModifyRouterInterfaceAttributeRequest) SetOppositeInterfaceId(v string) *ModifyRouterInterfaceAttributeRequest {
	s.OppositeInterfaceId = &v
	return s
}

func (s *ModifyRouterInterfaceAttributeRequest) SetOppositeInterfaceOwnerId(v int64) *ModifyRouterInterfaceAttributeRequest {
	s.OppositeInterfaceOwnerId = &v
	return s
}

func (s *ModifyRouterInterfaceAttributeRequest) SetOppositeRouterId(v string) *ModifyRouterInterfaceAttributeRequest {
	s.OppositeRouterId = &v
	return s
}

func (s *ModifyRouterInterfaceAttributeRequest) SetOppositeRouterType(v string) *ModifyRouterInterfaceAttributeRequest {
	s.OppositeRouterType = &v
	return s
}

func (s *ModifyRouterInterfaceAttributeRequest) SetOwnerId(v int64) *ModifyRouterInterfaceAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyRouterInterfaceAttributeRequest) SetRegionId(v string) *ModifyRouterInterfaceAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyRouterInterfaceAttributeRequest) SetResourceOwnerAccount(v string) *ModifyRouterInterfaceAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyRouterInterfaceAttributeRequest) SetResourceOwnerId(v int64) *ModifyRouterInterfaceAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyRouterInterfaceAttributeRequest) SetRouterInterfaceId(v string) *ModifyRouterInterfaceAttributeRequest {
	s.RouterInterfaceId = &v
	return s
}

func (s *ModifyRouterInterfaceAttributeRequest) Validate() error {
	return dara.Validate(s)
}
