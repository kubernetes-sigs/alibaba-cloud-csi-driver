// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewReservedInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenew(v bool) *RenewReservedInstancesRequest
	GetAutoRenew() *bool
	SetAutoRenewPeriod(v int32) *RenewReservedInstancesRequest
	GetAutoRenewPeriod() *int32
	SetClientToken(v string) *RenewReservedInstancesRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *RenewReservedInstancesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RenewReservedInstancesRequest
	GetOwnerId() *int64
	SetPeriod(v int32) *RenewReservedInstancesRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *RenewReservedInstancesRequest
	GetPeriodUnit() *string
	SetRegionId(v string) *RenewReservedInstancesRequest
	GetRegionId() *string
	SetReservedInstanceId(v []*string) *RenewReservedInstancesRequest
	GetReservedInstanceId() []*string
	SetResourceOwnerAccount(v string) *RenewReservedInstancesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RenewReservedInstancesRequest
	GetResourceOwnerId() *int64
}

type RenewReservedInstancesRequest struct {
	// Specifies whether to enable auto-renewal for the reserved instance. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The auto-renewal duration. Unit: months. This parameter takes effect only when AutoRenew is set to true.
	//
	// Valid values: 12 and 36. Default value: 12.
	//
	// example:
	//
	// 1
	AutoRenewPeriod *int32 `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **token*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The validity period of the reserved instance.
	//
	// Valid values: 1 and 3.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the validity period of the reserved instance.
	//
	// Set the value to Year.
	//
	// example:
	//
	// Year
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The region ID of the reserved instance.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of the reserved instances.
	ReservedInstanceId   []*string `json:"ReservedInstanceId,omitempty" xml:"ReservedInstanceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s RenewReservedInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s RenewReservedInstancesRequest) GoString() string {
	return s.String()
}

func (s *RenewReservedInstancesRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *RenewReservedInstancesRequest) GetAutoRenewPeriod() *int32 {
	return s.AutoRenewPeriod
}

func (s *RenewReservedInstancesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RenewReservedInstancesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RenewReservedInstancesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RenewReservedInstancesRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *RenewReservedInstancesRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *RenewReservedInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RenewReservedInstancesRequest) GetReservedInstanceId() []*string {
	return s.ReservedInstanceId
}

func (s *RenewReservedInstancesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RenewReservedInstancesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RenewReservedInstancesRequest) SetAutoRenew(v bool) *RenewReservedInstancesRequest {
	s.AutoRenew = &v
	return s
}

func (s *RenewReservedInstancesRequest) SetAutoRenewPeriod(v int32) *RenewReservedInstancesRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *RenewReservedInstancesRequest) SetClientToken(v string) *RenewReservedInstancesRequest {
	s.ClientToken = &v
	return s
}

func (s *RenewReservedInstancesRequest) SetOwnerAccount(v string) *RenewReservedInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RenewReservedInstancesRequest) SetOwnerId(v int64) *RenewReservedInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *RenewReservedInstancesRequest) SetPeriod(v int32) *RenewReservedInstancesRequest {
	s.Period = &v
	return s
}

func (s *RenewReservedInstancesRequest) SetPeriodUnit(v string) *RenewReservedInstancesRequest {
	s.PeriodUnit = &v
	return s
}

func (s *RenewReservedInstancesRequest) SetRegionId(v string) *RenewReservedInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *RenewReservedInstancesRequest) SetReservedInstanceId(v []*string) *RenewReservedInstancesRequest {
	s.ReservedInstanceId = v
	return s
}

func (s *RenewReservedInstancesRequest) SetResourceOwnerAccount(v string) *RenewReservedInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RenewReservedInstancesRequest) SetResourceOwnerId(v int64) *RenewReservedInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RenewReservedInstancesRequest) Validate() error {
	return dara.Validate(s)
}
