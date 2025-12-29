// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyReservedInstanceAutoRenewAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *ModifyReservedInstanceAutoRenewAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyReservedInstanceAutoRenewAttributeRequest
	GetOwnerId() *int64
	SetPeriod(v int32) *ModifyReservedInstanceAutoRenewAttributeRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *ModifyReservedInstanceAutoRenewAttributeRequest
	GetPeriodUnit() *string
	SetRegionId(v string) *ModifyReservedInstanceAutoRenewAttributeRequest
	GetRegionId() *string
	SetRenewalStatus(v string) *ModifyReservedInstanceAutoRenewAttributeRequest
	GetRenewalStatus() *string
	SetReservedInstanceId(v []*string) *ModifyReservedInstanceAutoRenewAttributeRequest
	GetReservedInstanceId() []*string
	SetResourceOwnerAccount(v string) *ModifyReservedInstanceAutoRenewAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyReservedInstanceAutoRenewAttributeRequest
	GetResourceOwnerId() *int64
}

type ModifyReservedInstanceAutoRenewAttributeRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The auto-renewal cycle.
	//
	// Valid values: 1 and 3.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the auto-renewal duration.
	//
	// Valid values: Year and Month.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The region ID of the reserved instances. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to automatically renew the reserved instances. Valid values:
	//
	// 	- AutoRenewal: automatically renews the reserved instances.
	//
	// 	- Normal: manually renews the reserved instances.
	//
	// example:
	//
	// AutoRenewal
	RenewalStatus *string `json:"RenewalStatus,omitempty" xml:"RenewalStatus,omitempty"`
	// The IDs of the reserved instances.
	ReservedInstanceId   []*string `json:"ReservedInstanceId,omitempty" xml:"ReservedInstanceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyReservedInstanceAutoRenewAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyReservedInstanceAutoRenewAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyReservedInstanceAutoRenewAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyReservedInstanceAutoRenewAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyReservedInstanceAutoRenewAttributeRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *ModifyReservedInstanceAutoRenewAttributeRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *ModifyReservedInstanceAutoRenewAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyReservedInstanceAutoRenewAttributeRequest) GetRenewalStatus() *string {
	return s.RenewalStatus
}

func (s *ModifyReservedInstanceAutoRenewAttributeRequest) GetReservedInstanceId() []*string {
	return s.ReservedInstanceId
}

func (s *ModifyReservedInstanceAutoRenewAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyReservedInstanceAutoRenewAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyReservedInstanceAutoRenewAttributeRequest) SetOwnerAccount(v string) *ModifyReservedInstanceAutoRenewAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyReservedInstanceAutoRenewAttributeRequest) SetOwnerId(v int64) *ModifyReservedInstanceAutoRenewAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyReservedInstanceAutoRenewAttributeRequest) SetPeriod(v int32) *ModifyReservedInstanceAutoRenewAttributeRequest {
	s.Period = &v
	return s
}

func (s *ModifyReservedInstanceAutoRenewAttributeRequest) SetPeriodUnit(v string) *ModifyReservedInstanceAutoRenewAttributeRequest {
	s.PeriodUnit = &v
	return s
}

func (s *ModifyReservedInstanceAutoRenewAttributeRequest) SetRegionId(v string) *ModifyReservedInstanceAutoRenewAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyReservedInstanceAutoRenewAttributeRequest) SetRenewalStatus(v string) *ModifyReservedInstanceAutoRenewAttributeRequest {
	s.RenewalStatus = &v
	return s
}

func (s *ModifyReservedInstanceAutoRenewAttributeRequest) SetReservedInstanceId(v []*string) *ModifyReservedInstanceAutoRenewAttributeRequest {
	s.ReservedInstanceId = v
	return s
}

func (s *ModifyReservedInstanceAutoRenewAttributeRequest) SetResourceOwnerAccount(v string) *ModifyReservedInstanceAutoRenewAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyReservedInstanceAutoRenewAttributeRequest) SetResourceOwnerId(v int64) *ModifyReservedInstanceAutoRenewAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyReservedInstanceAutoRenewAttributeRequest) Validate() error {
	return dara.Validate(s)
}
