// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPurchaseElasticityAssuranceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPrivatePoolOptions(v *PurchaseElasticityAssuranceRequestPrivatePoolOptions) *PurchaseElasticityAssuranceRequest
	GetPrivatePoolOptions() *PurchaseElasticityAssuranceRequestPrivatePoolOptions
	SetClientToken(v string) *PurchaseElasticityAssuranceRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *PurchaseElasticityAssuranceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *PurchaseElasticityAssuranceRequest
	GetOwnerId() *int64
	SetPeriod(v int32) *PurchaseElasticityAssuranceRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *PurchaseElasticityAssuranceRequest
	GetPeriodUnit() *string
	SetRegionId(v string) *PurchaseElasticityAssuranceRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *PurchaseElasticityAssuranceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *PurchaseElasticityAssuranceRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *PurchaseElasticityAssuranceRequest
	GetStartTime() *string
}

type PurchaseElasticityAssuranceRequest struct {
	PrivatePoolOptions *PurchaseElasticityAssuranceRequestPrivatePoolOptions `json:"PrivatePoolOptions,omitempty" xml:"PrivatePoolOptions,omitempty" type:"Struct"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The `token` can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The validity period of the elasticity assurance. The unit of the validity period is determined by the PeriodUnit value. Valid values:
	//
	// 	- When PeriodUnit is set to Month, valid values are 1, 2, 3, 4, 5, 6, 7, 8, and 9.
	//
	// 	- When PeriodUnit is set to Year, valid values are 1, 2, 3, 4, and 5.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the validity period of the elasticity assurance. Valid values:
	//
	// 	- Month
	//
	// 	- Year
	//
	// Default value: Year.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The ID of the region in which to purchase the elasticity assurance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/2679950.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The time when the elasticity assurance takes effect. The default value is the time when the elasticity assurance is created. Specify the time in the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC. For more information, see [ISO 8601](https://help.aliyun.com/document_detail/25696.html).
	//
	// example:
	//
	// 2024-06-18T00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s PurchaseElasticityAssuranceRequest) String() string {
	return dara.Prettify(s)
}

func (s PurchaseElasticityAssuranceRequest) GoString() string {
	return s.String()
}

func (s *PurchaseElasticityAssuranceRequest) GetPrivatePoolOptions() *PurchaseElasticityAssuranceRequestPrivatePoolOptions {
	return s.PrivatePoolOptions
}

func (s *PurchaseElasticityAssuranceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *PurchaseElasticityAssuranceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *PurchaseElasticityAssuranceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *PurchaseElasticityAssuranceRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *PurchaseElasticityAssuranceRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *PurchaseElasticityAssuranceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *PurchaseElasticityAssuranceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *PurchaseElasticityAssuranceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *PurchaseElasticityAssuranceRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *PurchaseElasticityAssuranceRequest) SetPrivatePoolOptions(v *PurchaseElasticityAssuranceRequestPrivatePoolOptions) *PurchaseElasticityAssuranceRequest {
	s.PrivatePoolOptions = v
	return s
}

func (s *PurchaseElasticityAssuranceRequest) SetClientToken(v string) *PurchaseElasticityAssuranceRequest {
	s.ClientToken = &v
	return s
}

func (s *PurchaseElasticityAssuranceRequest) SetOwnerAccount(v string) *PurchaseElasticityAssuranceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *PurchaseElasticityAssuranceRequest) SetOwnerId(v int64) *PurchaseElasticityAssuranceRequest {
	s.OwnerId = &v
	return s
}

func (s *PurchaseElasticityAssuranceRequest) SetPeriod(v int32) *PurchaseElasticityAssuranceRequest {
	s.Period = &v
	return s
}

func (s *PurchaseElasticityAssuranceRequest) SetPeriodUnit(v string) *PurchaseElasticityAssuranceRequest {
	s.PeriodUnit = &v
	return s
}

func (s *PurchaseElasticityAssuranceRequest) SetRegionId(v string) *PurchaseElasticityAssuranceRequest {
	s.RegionId = &v
	return s
}

func (s *PurchaseElasticityAssuranceRequest) SetResourceOwnerAccount(v string) *PurchaseElasticityAssuranceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *PurchaseElasticityAssuranceRequest) SetResourceOwnerId(v int64) *PurchaseElasticityAssuranceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *PurchaseElasticityAssuranceRequest) SetStartTime(v string) *PurchaseElasticityAssuranceRequest {
	s.StartTime = &v
	return s
}

func (s *PurchaseElasticityAssuranceRequest) Validate() error {
	if s.PrivatePoolOptions != nil {
		if err := s.PrivatePoolOptions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PurchaseElasticityAssuranceRequestPrivatePoolOptions struct {
	// The ID of the elasticity assurance.
	//
	// This parameter is required.
	//
	// example:
	//
	// eap-bp67acfmxazb4****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The type of the private pool with which you want to associate the elasticity assurance. Valid values:
	//
	// 	- Open: open private pool. If you use the elasticity assurance to create Elastic Compute Service (ECS) instances, the open private pool that is associated with the elasticity assurance is automatically matched. If no capacity is available in the open private pool, resources in the public pool are automatically used to create the ECS instances.
	//
	// 	- Target: targeted private pool. If you use the elasticity assurance to create ECS instances, the specified private pool that is associated with the elasticity assurance is automatically matched. If no capacity is available in the private pool, the ECS instances fail to be created.
	//
	// Default value: Open.
	//
	// example:
	//
	// Open
	MatchCriteria *string `json:"MatchCriteria,omitempty" xml:"MatchCriteria,omitempty"`
}

func (s PurchaseElasticityAssuranceRequestPrivatePoolOptions) String() string {
	return dara.Prettify(s)
}

func (s PurchaseElasticityAssuranceRequestPrivatePoolOptions) GoString() string {
	return s.String()
}

func (s *PurchaseElasticityAssuranceRequestPrivatePoolOptions) GetId() *string {
	return s.Id
}

func (s *PurchaseElasticityAssuranceRequestPrivatePoolOptions) GetMatchCriteria() *string {
	return s.MatchCriteria
}

func (s *PurchaseElasticityAssuranceRequestPrivatePoolOptions) SetId(v string) *PurchaseElasticityAssuranceRequestPrivatePoolOptions {
	s.Id = &v
	return s
}

func (s *PurchaseElasticityAssuranceRequestPrivatePoolOptions) SetMatchCriteria(v string) *PurchaseElasticityAssuranceRequestPrivatePoolOptions {
	s.MatchCriteria = &v
	return s
}

func (s *PurchaseElasticityAssuranceRequestPrivatePoolOptions) Validate() error {
	return dara.Validate(s)
}
