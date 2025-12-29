// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyElasticityAssuranceAutoRenewAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPrivatePoolOptions(v *ModifyElasticityAssuranceAutoRenewAttributeRequestPrivatePoolOptions) *ModifyElasticityAssuranceAutoRenewAttributeRequest
	GetPrivatePoolOptions() *ModifyElasticityAssuranceAutoRenewAttributeRequestPrivatePoolOptions
	SetOwnerAccount(v string) *ModifyElasticityAssuranceAutoRenewAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyElasticityAssuranceAutoRenewAttributeRequest
	GetOwnerId() *int64
	SetPeriod(v int32) *ModifyElasticityAssuranceAutoRenewAttributeRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *ModifyElasticityAssuranceAutoRenewAttributeRequest
	GetPeriodUnit() *string
	SetRegionId(v string) *ModifyElasticityAssuranceAutoRenewAttributeRequest
	GetRegionId() *string
	SetRenewalStatus(v string) *ModifyElasticityAssuranceAutoRenewAttributeRequest
	GetRenewalStatus() *string
	SetResourceOwnerAccount(v string) *ModifyElasticityAssuranceAutoRenewAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyElasticityAssuranceAutoRenewAttributeRequest
	GetResourceOwnerId() *int64
}

type ModifyElasticityAssuranceAutoRenewAttributeRequest struct {
	PrivatePoolOptions *ModifyElasticityAssuranceAutoRenewAttributeRequestPrivatePoolOptions `json:"PrivatePoolOptions,omitempty" xml:"PrivatePoolOptions,omitempty" type:"Struct"`
	OwnerAccount       *string                                                               `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId            *int64                                                                `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The auto-renewal period for the elasticity assurance.
	//
	// 	- Valid values when `PeriodUnit` is set to `Year`: 1, 3, and 5.
	//
	// 	- Valid values when `PeriodUnit` is set to `Month`: 1.
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
	// Default value: Month.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The ID of the region to which the elasticity assurance belongs. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The auto-renewal status of the elasticity assurance. Valid values:
	//
	// 	- AutoRenewal: Auto-renewal is enabled for the elasticity assurance.
	//
	// 	- Normal: Auto-renewal is disabled for the elasticity assurance.
	//
	// 	- NotRenewal: The elasticity assurance is not renewed. The system no longer sends an expiration notification but sends only a renewal notification three days before the elasticity assurance expires. You can change the value of this parameter from NotRenewal to `Normal` for an elasticity assurance, and then manually renew the elasticity assurance. Alternatively, you can set the RenewalStatus parameter to AutoRenewal.
	//
	// example:
	//
	// Normal
	RenewalStatus        *string `json:"RenewalStatus,omitempty" xml:"RenewalStatus,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyElasticityAssuranceAutoRenewAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyElasticityAssuranceAutoRenewAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyElasticityAssuranceAutoRenewAttributeRequest) GetPrivatePoolOptions() *ModifyElasticityAssuranceAutoRenewAttributeRequestPrivatePoolOptions {
	return s.PrivatePoolOptions
}

func (s *ModifyElasticityAssuranceAutoRenewAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyElasticityAssuranceAutoRenewAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyElasticityAssuranceAutoRenewAttributeRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *ModifyElasticityAssuranceAutoRenewAttributeRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *ModifyElasticityAssuranceAutoRenewAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyElasticityAssuranceAutoRenewAttributeRequest) GetRenewalStatus() *string {
	return s.RenewalStatus
}

func (s *ModifyElasticityAssuranceAutoRenewAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyElasticityAssuranceAutoRenewAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyElasticityAssuranceAutoRenewAttributeRequest) SetPrivatePoolOptions(v *ModifyElasticityAssuranceAutoRenewAttributeRequestPrivatePoolOptions) *ModifyElasticityAssuranceAutoRenewAttributeRequest {
	s.PrivatePoolOptions = v
	return s
}

func (s *ModifyElasticityAssuranceAutoRenewAttributeRequest) SetOwnerAccount(v string) *ModifyElasticityAssuranceAutoRenewAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyElasticityAssuranceAutoRenewAttributeRequest) SetOwnerId(v int64) *ModifyElasticityAssuranceAutoRenewAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyElasticityAssuranceAutoRenewAttributeRequest) SetPeriod(v int32) *ModifyElasticityAssuranceAutoRenewAttributeRequest {
	s.Period = &v
	return s
}

func (s *ModifyElasticityAssuranceAutoRenewAttributeRequest) SetPeriodUnit(v string) *ModifyElasticityAssuranceAutoRenewAttributeRequest {
	s.PeriodUnit = &v
	return s
}

func (s *ModifyElasticityAssuranceAutoRenewAttributeRequest) SetRegionId(v string) *ModifyElasticityAssuranceAutoRenewAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyElasticityAssuranceAutoRenewAttributeRequest) SetRenewalStatus(v string) *ModifyElasticityAssuranceAutoRenewAttributeRequest {
	s.RenewalStatus = &v
	return s
}

func (s *ModifyElasticityAssuranceAutoRenewAttributeRequest) SetResourceOwnerAccount(v string) *ModifyElasticityAssuranceAutoRenewAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyElasticityAssuranceAutoRenewAttributeRequest) SetResourceOwnerId(v int64) *ModifyElasticityAssuranceAutoRenewAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyElasticityAssuranceAutoRenewAttributeRequest) Validate() error {
	if s.PrivatePoolOptions != nil {
		if err := s.PrivatePoolOptions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyElasticityAssuranceAutoRenewAttributeRequestPrivatePoolOptions struct {
	// The IDs of elasticity assurances.
	//
	// >  You can renew up to 50 elasticity assurances at a time.
	Id []*string `json:"Id,omitempty" xml:"Id,omitempty" type:"Repeated"`
}

func (s ModifyElasticityAssuranceAutoRenewAttributeRequestPrivatePoolOptions) String() string {
	return dara.Prettify(s)
}

func (s ModifyElasticityAssuranceAutoRenewAttributeRequestPrivatePoolOptions) GoString() string {
	return s.String()
}

func (s *ModifyElasticityAssuranceAutoRenewAttributeRequestPrivatePoolOptions) GetId() []*string {
	return s.Id
}

func (s *ModifyElasticityAssuranceAutoRenewAttributeRequestPrivatePoolOptions) SetId(v []*string) *ModifyElasticityAssuranceAutoRenewAttributeRequestPrivatePoolOptions {
	s.Id = v
	return s
}

func (s *ModifyElasticityAssuranceAutoRenewAttributeRequestPrivatePoolOptions) Validate() error {
	return dara.Validate(s)
}
