// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RenewInstanceRequest
	GetClientToken() *string
	SetExpectedRenewDay(v int32) *RenewInstanceRequest
	GetExpectedRenewDay() *int32
	SetInstanceId(v string) *RenewInstanceRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *RenewInstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RenewInstanceRequest
	GetOwnerId() *int64
	SetPeriod(v int32) *RenewInstanceRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *RenewInstanceRequest
	GetPeriodUnit() *string
	SetResourceOwnerAccount(v string) *RenewInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RenewInstanceRequest
	GetResourceOwnerId() *int64
}

type RenewInstanceRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **token*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies the [unified expiration date](https://help.aliyun.com/document_detail/63396.html). Valid values: 1 to 28.
	//
	// To use this parameter, you must [specify a unified expiration date for the ECS instance](~~63396#694cb636c0rp6~~). The value of this parameter must be the same as the specified unified expiration date. Otherwise, the call fails.
	//
	// >  You must specify the renewal period-related parameter pair (`Period` and `PeriodUnit`) or `ExpectedRenewDay`, but not both.
	//
	// example:
	//
	// 5
	ExpectedRenewDay *int32 `json:"ExpectedRenewDay,omitempty" xml:"ExpectedRenewDay,omitempty"`
	// The ID of the instance that you want to renew.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp67acfmxazb4p****
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The renewal period of the subscription instance. If `DedicatedHostId` is specified, the value of Period cannot exceed the subscription period of the specified dedicated host.
	//
	// Valid values when PeriodUnit is set to Month: 1, 2, 3, 4, 5, 6, 7, 8, 9, and 12.
	//
	// > The renewal period-related parameter pair (`Period` and `PeriodUnit`) and `ExpectedRenewDay` are mutually exclusive.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the renewal period. Valid values:
	//
	// Month
	//
	// Default value: Month.
	//
	// example:
	//
	// Month
	PeriodUnit           *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s RenewInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RenewInstanceRequest) GoString() string {
	return s.String()
}

func (s *RenewInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RenewInstanceRequest) GetExpectedRenewDay() *int32 {
	return s.ExpectedRenewDay
}

func (s *RenewInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RenewInstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RenewInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RenewInstanceRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *RenewInstanceRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *RenewInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RenewInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RenewInstanceRequest) SetClientToken(v string) *RenewInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *RenewInstanceRequest) SetExpectedRenewDay(v int32) *RenewInstanceRequest {
	s.ExpectedRenewDay = &v
	return s
}

func (s *RenewInstanceRequest) SetInstanceId(v string) *RenewInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *RenewInstanceRequest) SetOwnerAccount(v string) *RenewInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RenewInstanceRequest) SetOwnerId(v int64) *RenewInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *RenewInstanceRequest) SetPeriod(v int32) *RenewInstanceRequest {
	s.Period = &v
	return s
}

func (s *RenewInstanceRequest) SetPeriodUnit(v string) *RenewInstanceRequest {
	s.PeriodUnit = &v
	return s
}

func (s *RenewInstanceRequest) SetResourceOwnerAccount(v string) *RenewInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RenewInstanceRequest) SetResourceOwnerId(v int64) *RenewInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RenewInstanceRequest) Validate() error {
	return dara.Validate(s)
}
