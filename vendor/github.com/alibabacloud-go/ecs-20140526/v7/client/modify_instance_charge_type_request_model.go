// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceChargeTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *ModifyInstanceChargeTypeRequest
	GetAutoPay() *bool
	SetClientToken(v string) *ModifyInstanceChargeTypeRequest
	GetClientToken() *string
	SetDryRun(v bool) *ModifyInstanceChargeTypeRequest
	GetDryRun() *bool
	SetIncludeDataDisks(v bool) *ModifyInstanceChargeTypeRequest
	GetIncludeDataDisks() *bool
	SetInstanceChargeType(v string) *ModifyInstanceChargeTypeRequest
	GetInstanceChargeType() *string
	SetInstanceIds(v string) *ModifyInstanceChargeTypeRequest
	GetInstanceIds() *string
	SetIsDetailFee(v bool) *ModifyInstanceChargeTypeRequest
	GetIsDetailFee() *bool
	SetOwnerAccount(v string) *ModifyInstanceChargeTypeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyInstanceChargeTypeRequest
	GetOwnerId() *int64
	SetPeriod(v int32) *ModifyInstanceChargeTypeRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *ModifyInstanceChargeTypeRequest
	GetPeriodUnit() *string
	SetRegionId(v string) *ModifyInstanceChargeTypeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyInstanceChargeTypeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyInstanceChargeTypeRequest
	GetResourceOwnerId() *int64
}

type ModifyInstanceChargeTypeRequest struct {
	// Specifies whether to automatically complete the payment. Valid values:
	//
	// 	- true: enables automatic payment. Maintain a sufficient account balance. Otherwise, your order becomes invalid and is canceled.
	//
	// 	- false: disables automatic payment. An order is generated but no payment is made.
	//
	// Default value: true.
	//
	// >  If your account balance is insufficient, you can set AutoPay to false to generate an unpaid order. Then, you can log on to the ECS console to pay for the order.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **token*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- true: performs only a dry run. The system checks the request for potential issues, including invalid AccessKey pairs, unauthorized Resource Access Management (RAM) users, and missing parameter values. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- false: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to change the billing method of all data disks on the instance from pay-as-you-go to subscription. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// example:
	//
	// false
	IncludeDataDisks *bool `json:"IncludeDataDisks,omitempty" xml:"IncludeDataDisks,omitempty"`
	// The new billing method of the instance. Valid values:
	//
	// 	- PrePaid: subscription
	//
	// 	- PostPaid: pay-as-you-go
	//
	// Default value: PrePaid.
	//
	// example:
	//
	// PrePaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The instance IDs. The value can be a JSON array that consists of up to 20 instance IDs. Separate the instance IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// ["i-bp67acfmxazb4p****","i-bp67acfmxazb4d****"]
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// Specifies whether to return cost details of the order after the billing method is changed from subscription to pay-as-you-go. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// example:
	//
	// false
	IsDetailFee  *bool   `json:"IsDetailFee,omitempty" xml:"IsDetailFee,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The renewal duration of the subscription instance. If the instance is hosted on a dedicated host, the renewal duration of the instance cannot exceed the subscription duration of the dedicated host. Valid values:
	//
	// Valid values when `PeriodUnit` is set to Month: `1, 2, 3, 4, 5, 6, 7, 8, 9, and 12`.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the renewal duration specified by `Period`. Valid values:
	//
	// Month
	//
	// Default value: Month.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The region ID of the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyInstanceChargeTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceChargeTypeRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceChargeTypeRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *ModifyInstanceChargeTypeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyInstanceChargeTypeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyInstanceChargeTypeRequest) GetIncludeDataDisks() *bool {
	return s.IncludeDataDisks
}

func (s *ModifyInstanceChargeTypeRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *ModifyInstanceChargeTypeRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *ModifyInstanceChargeTypeRequest) GetIsDetailFee() *bool {
	return s.IsDetailFee
}

func (s *ModifyInstanceChargeTypeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyInstanceChargeTypeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyInstanceChargeTypeRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *ModifyInstanceChargeTypeRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *ModifyInstanceChargeTypeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyInstanceChargeTypeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyInstanceChargeTypeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyInstanceChargeTypeRequest) SetAutoPay(v bool) *ModifyInstanceChargeTypeRequest {
	s.AutoPay = &v
	return s
}

func (s *ModifyInstanceChargeTypeRequest) SetClientToken(v string) *ModifyInstanceChargeTypeRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyInstanceChargeTypeRequest) SetDryRun(v bool) *ModifyInstanceChargeTypeRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyInstanceChargeTypeRequest) SetIncludeDataDisks(v bool) *ModifyInstanceChargeTypeRequest {
	s.IncludeDataDisks = &v
	return s
}

func (s *ModifyInstanceChargeTypeRequest) SetInstanceChargeType(v string) *ModifyInstanceChargeTypeRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *ModifyInstanceChargeTypeRequest) SetInstanceIds(v string) *ModifyInstanceChargeTypeRequest {
	s.InstanceIds = &v
	return s
}

func (s *ModifyInstanceChargeTypeRequest) SetIsDetailFee(v bool) *ModifyInstanceChargeTypeRequest {
	s.IsDetailFee = &v
	return s
}

func (s *ModifyInstanceChargeTypeRequest) SetOwnerAccount(v string) *ModifyInstanceChargeTypeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyInstanceChargeTypeRequest) SetOwnerId(v int64) *ModifyInstanceChargeTypeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyInstanceChargeTypeRequest) SetPeriod(v int32) *ModifyInstanceChargeTypeRequest {
	s.Period = &v
	return s
}

func (s *ModifyInstanceChargeTypeRequest) SetPeriodUnit(v string) *ModifyInstanceChargeTypeRequest {
	s.PeriodUnit = &v
	return s
}

func (s *ModifyInstanceChargeTypeRequest) SetRegionId(v string) *ModifyInstanceChargeTypeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyInstanceChargeTypeRequest) SetResourceOwnerAccount(v string) *ModifyInstanceChargeTypeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyInstanceChargeTypeRequest) SetResourceOwnerId(v int64) *ModifyInstanceChargeTypeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyInstanceChargeTypeRequest) Validate() error {
	return dara.Validate(s)
}
