// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDedicatedHostAutoRenewAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenew(v bool) *ModifyDedicatedHostAutoRenewAttributeRequest
	GetAutoRenew() *bool
	SetAutoRenewWithEcs(v string) *ModifyDedicatedHostAutoRenewAttributeRequest
	GetAutoRenewWithEcs() *string
	SetDedicatedHostIds(v string) *ModifyDedicatedHostAutoRenewAttributeRequest
	GetDedicatedHostIds() *string
	SetDuration(v int32) *ModifyDedicatedHostAutoRenewAttributeRequest
	GetDuration() *int32
	SetOwnerAccount(v string) *ModifyDedicatedHostAutoRenewAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDedicatedHostAutoRenewAttributeRequest
	GetOwnerId() *int64
	SetPeriodUnit(v string) *ModifyDedicatedHostAutoRenewAttributeRequest
	GetPeriodUnit() *string
	SetRegionId(v string) *ModifyDedicatedHostAutoRenewAttributeRequest
	GetRegionId() *string
	SetRenewalStatus(v string) *ModifyDedicatedHostAutoRenewAttributeRequest
	GetRenewalStatus() *string
	SetResourceOwnerAccount(v string) *ModifyDedicatedHostAutoRenewAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDedicatedHostAutoRenewAttributeRequest
	GetResourceOwnerId() *int64
}

type ModifyDedicatedHostAutoRenewAttributeRequest struct {
	// Specifies whether to automatically renew the subscription. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// Specifies whether to automatically renew the subscription dedicated hosts along with the subscription ECS instances hosted on the dedicated hosts.
	//
	// If auto-renewal is enabled for the subscription ECS instances hosted on the subscription dedicated hosts, you can specify this parameter to automatically renew the dedicated hosts along with the subscription ECS instances. When the subscription ECS instances hosted on your dedicated hosts are automatically renewed, the subscription dedicated hosts are also automatically renewed if the expiration time of the dedicated hosts is earlier than the expiration time of the renewed instances. Take note of the following items:
	//
	// When the subscription dedicated hosts are configured to be automatically renewed along with the subscription ECS instances hosted on the dedicated hosts, the system checks the expiration time of the renewed instances and selects a minimum renewal duration for the dedicated hosts so that the dedicated hosts are renewed by a duration that ends later than the expiration time of the renewed instances. For more information about supported renewal durations, see the descriptions of the `PeriodUnit` and `Duration` parameters.
	//
	// For example, assume that a dedicated host expires on January 15 of the current year. Subscription ECS instances hosted on the dedicated host are configured to be automatically renewed to November 15 of the same year. The expiration time of the dedicated host is earlier than the expiration time of the ECS instances by 10 months. In this case, the system selects a renewal duration of 12 months (a minimum duration calculated based on a `Duration` value of 12 and a `PeriodUnit` value of Month) for the dedicated host. This ensures that the dedicated host expires later than the ECS instances.
	//
	// Valid values:
	//
	// 	- AutoRenewWithEcs: automatically renews the subscription dedicated hosts along with the subscription ECS instances hosted on the dedicated hosts.
	//
	// 	- StopRenewWithEcs: does not automatically renew the subscription dedicated hosts along with the subscription ECS instances hosted on the dedicated hosts.
	//
	// 	- NoOperation: does not change the current settings for the dedicated hosts.
	//
	// > If you set this parameter to AutoRenewWithEcs, make sure that `AutoRenew` is set to true to enable auto-renewal for the dedicated hosts. Otherwise, the subscription dedicated hosts are not automatically renewed along with the subscription ECS instances hosted on the dedicated hosts.
	//
	// Default value: NoOperation.
	//
	// example:
	//
	// StopRenewWithEcs
	AutoRenewWithEcs *string `json:"AutoRenewWithEcs,omitempty" xml:"AutoRenewWithEcs,omitempty"`
	// The IDs of dedicated hosts. You can specify up to 100 subscription dedicated host IDs. Separate the IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// dh-bp165p6xk2tlw61e****
	DedicatedHostIds *string `json:"DedicatedHostIds,omitempty" xml:"DedicatedHostIds,omitempty"`
	// The renewal duration.
	//
	// 	- Valid values when PeriodUnit is set to Month: 1 and 12
	//
	// 	- Valid values when PeriodUnit is set to Year: 1 and 12
	//
	// example:
	//
	// 1
	Duration     *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The unit of the renewal period. Valid values:
	//
	// 	- Month
	//
	// 	- Year
	//
	// Default value: Month
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The region ID of the dedicated host.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to automatically renew the subscription dedicated host. The `RenewalStatus` parameter takes precedence over the `AutoRenew` parameter. Valid values:
	//
	// 	- AutoRenewal: The dedicated hosts are automatically renewed.
	//
	// 	- Normal: The dedicated hosts are not automatically renewed, and renewal notifications are sent.
	//
	// 	- NotRenewal: The dedicated hosts are not automatically renewed, and no expiration notification is sent. A notification of no renewal is automatically sent three days before the end of the current subscription cycle. You can change the value of this parameter from NotRenewal to Normal and manually renew the dedicated hosts by calling the [RenewDedicatedHosts](https://help.aliyun.com/document_detail/134250.html) operation. Alternatively, you can renew the dedicated hosts by setting this parameter to AutoRenewal.
	//
	// example:
	//
	// Normal
	RenewalStatus        *string `json:"RenewalStatus,omitempty" xml:"RenewalStatus,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDedicatedHostAutoRenewAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDedicatedHostAutoRenewAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDedicatedHostAutoRenewAttributeRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *ModifyDedicatedHostAutoRenewAttributeRequest) GetAutoRenewWithEcs() *string {
	return s.AutoRenewWithEcs
}

func (s *ModifyDedicatedHostAutoRenewAttributeRequest) GetDedicatedHostIds() *string {
	return s.DedicatedHostIds
}

func (s *ModifyDedicatedHostAutoRenewAttributeRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *ModifyDedicatedHostAutoRenewAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDedicatedHostAutoRenewAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDedicatedHostAutoRenewAttributeRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *ModifyDedicatedHostAutoRenewAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDedicatedHostAutoRenewAttributeRequest) GetRenewalStatus() *string {
	return s.RenewalStatus
}

func (s *ModifyDedicatedHostAutoRenewAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDedicatedHostAutoRenewAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDedicatedHostAutoRenewAttributeRequest) SetAutoRenew(v bool) *ModifyDedicatedHostAutoRenewAttributeRequest {
	s.AutoRenew = &v
	return s
}

func (s *ModifyDedicatedHostAutoRenewAttributeRequest) SetAutoRenewWithEcs(v string) *ModifyDedicatedHostAutoRenewAttributeRequest {
	s.AutoRenewWithEcs = &v
	return s
}

func (s *ModifyDedicatedHostAutoRenewAttributeRequest) SetDedicatedHostIds(v string) *ModifyDedicatedHostAutoRenewAttributeRequest {
	s.DedicatedHostIds = &v
	return s
}

func (s *ModifyDedicatedHostAutoRenewAttributeRequest) SetDuration(v int32) *ModifyDedicatedHostAutoRenewAttributeRequest {
	s.Duration = &v
	return s
}

func (s *ModifyDedicatedHostAutoRenewAttributeRequest) SetOwnerAccount(v string) *ModifyDedicatedHostAutoRenewAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDedicatedHostAutoRenewAttributeRequest) SetOwnerId(v int64) *ModifyDedicatedHostAutoRenewAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDedicatedHostAutoRenewAttributeRequest) SetPeriodUnit(v string) *ModifyDedicatedHostAutoRenewAttributeRequest {
	s.PeriodUnit = &v
	return s
}

func (s *ModifyDedicatedHostAutoRenewAttributeRequest) SetRegionId(v string) *ModifyDedicatedHostAutoRenewAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDedicatedHostAutoRenewAttributeRequest) SetRenewalStatus(v string) *ModifyDedicatedHostAutoRenewAttributeRequest {
	s.RenewalStatus = &v
	return s
}

func (s *ModifyDedicatedHostAutoRenewAttributeRequest) SetResourceOwnerAccount(v string) *ModifyDedicatedHostAutoRenewAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDedicatedHostAutoRenewAttributeRequest) SetResourceOwnerId(v int64) *ModifyDedicatedHostAutoRenewAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDedicatedHostAutoRenewAttributeRequest) Validate() error {
	return dara.Validate(s)
}
