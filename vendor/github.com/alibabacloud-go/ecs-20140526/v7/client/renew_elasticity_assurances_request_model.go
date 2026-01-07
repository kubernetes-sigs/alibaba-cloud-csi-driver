// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewElasticityAssurancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPrivatePoolOptions(v *RenewElasticityAssurancesRequestPrivatePoolOptions) *RenewElasticityAssurancesRequest
	GetPrivatePoolOptions() *RenewElasticityAssurancesRequestPrivatePoolOptions
	SetAutoPay(v bool) *RenewElasticityAssurancesRequest
	GetAutoPay() *bool
	SetAutoRenew(v bool) *RenewElasticityAssurancesRequest
	GetAutoRenew() *bool
	SetAutoRenewPeriod(v int32) *RenewElasticityAssurancesRequest
	GetAutoRenewPeriod() *int32
	SetClientToken(v string) *RenewElasticityAssurancesRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *RenewElasticityAssurancesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RenewElasticityAssurancesRequest
	GetOwnerId() *int64
	SetPeriod(v int32) *RenewElasticityAssurancesRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *RenewElasticityAssurancesRequest
	GetPeriodUnit() *string
	SetRegionId(v string) *RenewElasticityAssurancesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *RenewElasticityAssurancesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RenewElasticityAssurancesRequest
	GetResourceOwnerId() *int64
}

type RenewElasticityAssurancesRequest struct {
	PrivatePoolOptions *RenewElasticityAssurancesRequestPrivatePoolOptions `json:"PrivatePoolOptions,omitempty" xml:"PrivatePoolOptions,omitempty" type:"Struct"`
	// Specifies whether to enable automatic payment. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: true.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal for the elasticity assurance. Valid values:
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
	// The auto-renewal period. Unit: month. Valid values: 1, 2, 3, 6, 12, 24, and 36.
	//
	// 	- If you set `PeriodUnit` to Month, the default value is 1.
	//
	// 	- If you set `PeriodUnit` to Year, the default value is 12.
	//
	// >  This parameter is required if you set `AutoRenew` to `true`.
	//
	// example:
	//
	// 1
	AutoRenewPeriod *int32 `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests.
	//
	// The `token` can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The renewal duration. The unit of the renewal duration is determined by the `PeriodUnit` value. Valid values:
	//
	// 	- Valid values if you set `PeriodUnit` to `Month`: 1, 2, 3, 4, 5, 6, 7, 8, and 9.
	//
	// 	- Valid values if you set `PeriodUnit` to `Year`: 1, 2, and 3.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the renewal duration. Valid values:
	//
	// 	- Month
	//
	// 	- Year
	//
	// Default value: Year.
	//
	// example:
	//
	// Year
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The region ID of the elasticity assurance.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/2680071.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s RenewElasticityAssurancesRequest) String() string {
	return dara.Prettify(s)
}

func (s RenewElasticityAssurancesRequest) GoString() string {
	return s.String()
}

func (s *RenewElasticityAssurancesRequest) GetPrivatePoolOptions() *RenewElasticityAssurancesRequestPrivatePoolOptions {
	return s.PrivatePoolOptions
}

func (s *RenewElasticityAssurancesRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *RenewElasticityAssurancesRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *RenewElasticityAssurancesRequest) GetAutoRenewPeriod() *int32 {
	return s.AutoRenewPeriod
}

func (s *RenewElasticityAssurancesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RenewElasticityAssurancesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RenewElasticityAssurancesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RenewElasticityAssurancesRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *RenewElasticityAssurancesRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *RenewElasticityAssurancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RenewElasticityAssurancesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RenewElasticityAssurancesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RenewElasticityAssurancesRequest) SetPrivatePoolOptions(v *RenewElasticityAssurancesRequestPrivatePoolOptions) *RenewElasticityAssurancesRequest {
	s.PrivatePoolOptions = v
	return s
}

func (s *RenewElasticityAssurancesRequest) SetAutoPay(v bool) *RenewElasticityAssurancesRequest {
	s.AutoPay = &v
	return s
}

func (s *RenewElasticityAssurancesRequest) SetAutoRenew(v bool) *RenewElasticityAssurancesRequest {
	s.AutoRenew = &v
	return s
}

func (s *RenewElasticityAssurancesRequest) SetAutoRenewPeriod(v int32) *RenewElasticityAssurancesRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *RenewElasticityAssurancesRequest) SetClientToken(v string) *RenewElasticityAssurancesRequest {
	s.ClientToken = &v
	return s
}

func (s *RenewElasticityAssurancesRequest) SetOwnerAccount(v string) *RenewElasticityAssurancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RenewElasticityAssurancesRequest) SetOwnerId(v int64) *RenewElasticityAssurancesRequest {
	s.OwnerId = &v
	return s
}

func (s *RenewElasticityAssurancesRequest) SetPeriod(v int32) *RenewElasticityAssurancesRequest {
	s.Period = &v
	return s
}

func (s *RenewElasticityAssurancesRequest) SetPeriodUnit(v string) *RenewElasticityAssurancesRequest {
	s.PeriodUnit = &v
	return s
}

func (s *RenewElasticityAssurancesRequest) SetRegionId(v string) *RenewElasticityAssurancesRequest {
	s.RegionId = &v
	return s
}

func (s *RenewElasticityAssurancesRequest) SetResourceOwnerAccount(v string) *RenewElasticityAssurancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RenewElasticityAssurancesRequest) SetResourceOwnerId(v int64) *RenewElasticityAssurancesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RenewElasticityAssurancesRequest) Validate() error {
	if s.PrivatePoolOptions != nil {
		if err := s.PrivatePoolOptions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RenewElasticityAssurancesRequestPrivatePoolOptions struct {
	// The IDs of elasticity assurances.
	//
	// **Limits**: You can renew up to 20 elasticity assurances at a time.
	//
	// You can call the [DescribeElasticityAssurances](https://help.aliyun.com/document_detail/2679748.html) operation to query the elasticity assurances that you purchased.
	Id []*string `json:"Id,omitempty" xml:"Id,omitempty" type:"Repeated"`
}

func (s RenewElasticityAssurancesRequestPrivatePoolOptions) String() string {
	return dara.Prettify(s)
}

func (s RenewElasticityAssurancesRequestPrivatePoolOptions) GoString() string {
	return s.String()
}

func (s *RenewElasticityAssurancesRequestPrivatePoolOptions) GetId() []*string {
	return s.Id
}

func (s *RenewElasticityAssurancesRequestPrivatePoolOptions) SetId(v []*string) *RenewElasticityAssurancesRequestPrivatePoolOptions {
	s.Id = v
	return s
}

func (s *RenewElasticityAssurancesRequestPrivatePoolOptions) Validate() error {
	return dara.Validate(s)
}
