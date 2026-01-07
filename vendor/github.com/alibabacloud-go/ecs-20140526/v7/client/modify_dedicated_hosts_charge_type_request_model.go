// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDedicatedHostsChargeTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *ModifyDedicatedHostsChargeTypeRequest
	GetAutoPay() *bool
	SetClientToken(v string) *ModifyDedicatedHostsChargeTypeRequest
	GetClientToken() *string
	SetDedicatedHostChargeType(v string) *ModifyDedicatedHostsChargeTypeRequest
	GetDedicatedHostChargeType() *string
	SetDedicatedHostIds(v string) *ModifyDedicatedHostsChargeTypeRequest
	GetDedicatedHostIds() *string
	SetDetailFee(v bool) *ModifyDedicatedHostsChargeTypeRequest
	GetDetailFee() *bool
	SetDryRun(v bool) *ModifyDedicatedHostsChargeTypeRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *ModifyDedicatedHostsChargeTypeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDedicatedHostsChargeTypeRequest
	GetOwnerId() *int64
	SetPeriod(v int32) *ModifyDedicatedHostsChargeTypeRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *ModifyDedicatedHostsChargeTypeRequest
	GetPeriodUnit() *string
	SetRegionId(v string) *ModifyDedicatedHostsChargeTypeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyDedicatedHostsChargeTypeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDedicatedHostsChargeTypeRequest
	GetResourceOwnerId() *int64
}

type ModifyDedicatedHostsChargeTypeRequest struct {
	// Specifies whether to automatically complete the payment. Valid value:
	//
	// 	- true: The payment is automatically completed. Ensure that your account balance is sufficient. Otherwise, your order becomes invalid and must be canceled.
	//
	// 	- false: An order is generated but no payment is made.
	//
	// Default value: true.
	//
	// >  If you do not have sufficient balance in your account, you can set `AutoPay` to `false` to generate an unpaid order. Then, you can pay for the order.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The `token` can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The new billing method for the dedicated host. Valid value:
	//
	// 	- PrePaid: changes the billing method from pay-as-you-go to subscription.
	//
	// 	- PostPaid: changes the billing method from subscription to pay-as-you-go.
	//
	// Default value: PrePaid.
	//
	// example:
	//
	// PrePaid
	DedicatedHostChargeType *string `json:"DedicatedHostChargeType,omitempty" xml:"DedicatedHostChargeType,omitempty"`
	// The IDs of the dedicated hosts. The value can be a JSON array that consists of up to 20 dedicated host IDs. Separate the IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// ["dh-bp181e5064b5sotr****","dh-bp18064b5sotrr9c****"]
	DedicatedHostIds *string `json:"DedicatedHostIds,omitempty" xml:"DedicatedHostIds,omitempty"`
	// Specifies whether to return the billing details of the order when the billing method is changed from subscription to pay-as-you-go.
	//
	// Indicates whether the nfs volume is set to the read-only mode. Default value: false.
	//
	// example:
	//
	// false
	DetailFee *bool `json:"DetailFee,omitempty" xml:"DetailFee,omitempty"`
	// Specifies whether to perform only a dry run. Valid value:
	//
	// 	- true: performs only a dry run. The system checks the request for potential issues, including invalid AccessKey pairs, unauthorized Resource Access Management (RAM) users, and missing parameter values. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- false: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// Indicates whether the nfs volume is set to the read-only mode. Default value: false.
	//
	// example:
	//
	// false
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The renewal duration of the subscription dedicated hosts. Valid values:
	//
	// 	- If you set `PeriodUnit` to Week, valid values of `Period` are 1, 2, 3, and 4.
	//
	// 	- If you set `PeriodUnit` to Month, valid values of `Period` are 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 48, and 60.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the renewal duration specified by `Period`. Valid values:
	//
	// 	- Week
	//
	// 	- Month
	//
	// Default value: Month.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The region ID of the dedicated hosts. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
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

func (s ModifyDedicatedHostsChargeTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDedicatedHostsChargeTypeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDedicatedHostsChargeTypeRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *ModifyDedicatedHostsChargeTypeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyDedicatedHostsChargeTypeRequest) GetDedicatedHostChargeType() *string {
	return s.DedicatedHostChargeType
}

func (s *ModifyDedicatedHostsChargeTypeRequest) GetDedicatedHostIds() *string {
	return s.DedicatedHostIds
}

func (s *ModifyDedicatedHostsChargeTypeRequest) GetDetailFee() *bool {
	return s.DetailFee
}

func (s *ModifyDedicatedHostsChargeTypeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyDedicatedHostsChargeTypeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDedicatedHostsChargeTypeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDedicatedHostsChargeTypeRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *ModifyDedicatedHostsChargeTypeRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *ModifyDedicatedHostsChargeTypeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDedicatedHostsChargeTypeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDedicatedHostsChargeTypeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDedicatedHostsChargeTypeRequest) SetAutoPay(v bool) *ModifyDedicatedHostsChargeTypeRequest {
	s.AutoPay = &v
	return s
}

func (s *ModifyDedicatedHostsChargeTypeRequest) SetClientToken(v string) *ModifyDedicatedHostsChargeTypeRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDedicatedHostsChargeTypeRequest) SetDedicatedHostChargeType(v string) *ModifyDedicatedHostsChargeTypeRequest {
	s.DedicatedHostChargeType = &v
	return s
}

func (s *ModifyDedicatedHostsChargeTypeRequest) SetDedicatedHostIds(v string) *ModifyDedicatedHostsChargeTypeRequest {
	s.DedicatedHostIds = &v
	return s
}

func (s *ModifyDedicatedHostsChargeTypeRequest) SetDetailFee(v bool) *ModifyDedicatedHostsChargeTypeRequest {
	s.DetailFee = &v
	return s
}

func (s *ModifyDedicatedHostsChargeTypeRequest) SetDryRun(v bool) *ModifyDedicatedHostsChargeTypeRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyDedicatedHostsChargeTypeRequest) SetOwnerAccount(v string) *ModifyDedicatedHostsChargeTypeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDedicatedHostsChargeTypeRequest) SetOwnerId(v int64) *ModifyDedicatedHostsChargeTypeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDedicatedHostsChargeTypeRequest) SetPeriod(v int32) *ModifyDedicatedHostsChargeTypeRequest {
	s.Period = &v
	return s
}

func (s *ModifyDedicatedHostsChargeTypeRequest) SetPeriodUnit(v string) *ModifyDedicatedHostsChargeTypeRequest {
	s.PeriodUnit = &v
	return s
}

func (s *ModifyDedicatedHostsChargeTypeRequest) SetRegionId(v string) *ModifyDedicatedHostsChargeTypeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDedicatedHostsChargeTypeRequest) SetResourceOwnerAccount(v string) *ModifyDedicatedHostsChargeTypeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDedicatedHostsChargeTypeRequest) SetResourceOwnerId(v int64) *ModifyDedicatedHostsChargeTypeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDedicatedHostsChargeTypeRequest) Validate() error {
	return dara.Validate(s)
}
