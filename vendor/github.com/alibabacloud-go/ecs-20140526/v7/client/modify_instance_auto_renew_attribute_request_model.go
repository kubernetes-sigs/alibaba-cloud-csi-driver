// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceAutoRenewAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenew(v bool) *ModifyInstanceAutoRenewAttributeRequest
	GetAutoRenew() *bool
	SetDuration(v int32) *ModifyInstanceAutoRenewAttributeRequest
	GetDuration() *int32
	SetInstanceId(v string) *ModifyInstanceAutoRenewAttributeRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *ModifyInstanceAutoRenewAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyInstanceAutoRenewAttributeRequest
	GetOwnerId() *int64
	SetPeriodUnit(v string) *ModifyInstanceAutoRenewAttributeRequest
	GetPeriodUnit() *string
	SetRegionId(v string) *ModifyInstanceAutoRenewAttributeRequest
	GetRegionId() *string
	SetRenewalStatus(v string) *ModifyInstanceAutoRenewAttributeRequest
	GetRenewalStatus() *string
	SetResourceOwnerAccount(v string) *ModifyInstanceAutoRenewAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyInstanceAutoRenewAttributeRequest
	GetResourceOwnerId() *int64
}

type ModifyInstanceAutoRenewAttributeRequest struct {
	// Specifies whether to enable auto-renewal.
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
	// The auto-renewal period of the instance.
	//
	// 	- Valid values of `Duration` when `PeriodUnit` is set to `Year`: 1, 2, 3, 4, and 5.
	//
	// 	- Valid values of `Duration` when `PeriodUnit` is set to `Month`: 1, 2, 3, 6, 12, 24, 36, 48, and 60.
	//
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The ID of the instance. You can specify up to 100 subscription instance IDs. Separate the instance IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp67acfmxazb4ph****,i-bp67acfmxazb4pi****
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The unit of the renewal period (`Duration`). Valid values:
	//
	// 	- Month (default)
	//
	// 	- Year
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
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The auto-renewal status of the instance. Valid values:
	//
	// 	- AutoRenewal: Auto-renewal is enabled for the instance.
	//
	// 	- Normal: Auto-renewal is disabled for the instance.
	//
	// 	- NotRenewal: The instance is not renewed. The system no longer sends an expiration notification but sends only a renewal notification three days before the instance expires. You can change the value of this parameter from NotRenewal to `Normal` for an instance, and then manually renew the instance. Alternatively, you can set the RenewalStatus parameter to AutoRenewal.
	//
	// > `RenewalStatus` takes precedence over `AutoRenew`. If you do not specify `RenewalStatus`, `AutoRenew` is used by default.
	//
	// example:
	//
	// AutoRenewal
	RenewalStatus        *string `json:"RenewalStatus,omitempty" xml:"RenewalStatus,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyInstanceAutoRenewAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceAutoRenewAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAutoRenewAttributeRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *ModifyInstanceAutoRenewAttributeRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *ModifyInstanceAutoRenewAttributeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceAutoRenewAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyInstanceAutoRenewAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyInstanceAutoRenewAttributeRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *ModifyInstanceAutoRenewAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyInstanceAutoRenewAttributeRequest) GetRenewalStatus() *string {
	return s.RenewalStatus
}

func (s *ModifyInstanceAutoRenewAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyInstanceAutoRenewAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyInstanceAutoRenewAttributeRequest) SetAutoRenew(v bool) *ModifyInstanceAutoRenewAttributeRequest {
	s.AutoRenew = &v
	return s
}

func (s *ModifyInstanceAutoRenewAttributeRequest) SetDuration(v int32) *ModifyInstanceAutoRenewAttributeRequest {
	s.Duration = &v
	return s
}

func (s *ModifyInstanceAutoRenewAttributeRequest) SetInstanceId(v string) *ModifyInstanceAutoRenewAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceAutoRenewAttributeRequest) SetOwnerAccount(v string) *ModifyInstanceAutoRenewAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyInstanceAutoRenewAttributeRequest) SetOwnerId(v int64) *ModifyInstanceAutoRenewAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyInstanceAutoRenewAttributeRequest) SetPeriodUnit(v string) *ModifyInstanceAutoRenewAttributeRequest {
	s.PeriodUnit = &v
	return s
}

func (s *ModifyInstanceAutoRenewAttributeRequest) SetRegionId(v string) *ModifyInstanceAutoRenewAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyInstanceAutoRenewAttributeRequest) SetRenewalStatus(v string) *ModifyInstanceAutoRenewAttributeRequest {
	s.RenewalStatus = &v
	return s
}

func (s *ModifyInstanceAutoRenewAttributeRequest) SetResourceOwnerAccount(v string) *ModifyInstanceAutoRenewAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyInstanceAutoRenewAttributeRequest) SetResourceOwnerId(v int64) *ModifyInstanceAutoRenewAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyInstanceAutoRenewAttributeRequest) Validate() error {
	return dara.Validate(s)
}
