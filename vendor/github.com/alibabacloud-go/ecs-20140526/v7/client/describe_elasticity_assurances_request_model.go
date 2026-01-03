// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeElasticityAssurancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPrivatePoolOptions(v *DescribeElasticityAssurancesRequestPrivatePoolOptions) *DescribeElasticityAssurancesRequest
	GetPrivatePoolOptions() *DescribeElasticityAssurancesRequestPrivatePoolOptions
	SetInstanceChargeType(v string) *DescribeElasticityAssurancesRequest
	GetInstanceChargeType() *string
	SetInstanceType(v string) *DescribeElasticityAssurancesRequest
	GetInstanceType() *string
	SetMaxResults(v int32) *DescribeElasticityAssurancesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeElasticityAssurancesRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *DescribeElasticityAssurancesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeElasticityAssurancesRequest
	GetOwnerId() *int64
	SetPackageType(v string) *DescribeElasticityAssurancesRequest
	GetPackageType() *string
	SetPlatform(v string) *DescribeElasticityAssurancesRequest
	GetPlatform() *string
	SetRegionId(v string) *DescribeElasticityAssurancesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeElasticityAssurancesRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeElasticityAssurancesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeElasticityAssurancesRequest
	GetResourceOwnerId() *int64
	SetStatus(v string) *DescribeElasticityAssurancesRequest
	GetStatus() *string
	SetTag(v []*DescribeElasticityAssurancesRequestTag) *DescribeElasticityAssurancesRequest
	GetTag() []*DescribeElasticityAssurancesRequestTag
	SetZoneId(v string) *DescribeElasticityAssurancesRequest
	GetZoneId() *string
}

type DescribeElasticityAssurancesRequest struct {
	PrivatePoolOptions *DescribeElasticityAssurancesRequestPrivatePoolOptions `json:"PrivatePoolOptions,omitempty" xml:"PrivatePoolOptions,omitempty" type:"Struct"`
	// The billing method of the instance. Set the value to PostPaid. Only pay-as-you-go instances can be created by using elasticity assurances.
	//
	// Default value: PostPaid.
	//
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The instance type.
	//
	// example:
	//
	// ecs.c6.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The maximum number of entries per page.
	//
	// Maximum value: 100
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the request to retrieve a new page of results. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The type of the elasticity assurance. Valid values:
	//
	// 	- ElasticityAssurance: the general-purpose elasticity assurance. RecurrenceRules is not specified for a general-purpose elasticity assurance.
	//
	// 	- TimeDivisionElasticityAssurance: the time-segmented elasticity assurance. RecurrenceRules is specified for a time-segmented assurance.
	//
	// >  Time-segmented elasticity assurances are available only in specific regions and to specific users. To use time-segmented elasticity assurances, [submit a ticket](https://smartservice.console.aliyun.com/service/create-ticket-intl).
	//
	// example:
	//
	// ElasticityAssurance
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// > This parameter is deprecated.
	//
	// example:
	//
	// null
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The region ID of the elasticity assurances. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group. If you configure this parameter to query resources, up to 1,000 resources that belong to the specified resource group can be displayed in the response.
	//
	// > Resources in the default resource group are displayed in the response regardless of whether you configure this parameter.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The status of the elasticity assurance. Valid values:
	//
	// 	- All: All states.
	//
	// 	- Deactivated: The elasticity assurance is pending activation. This state is in invitational preview.
	//
	// 	- Preparing: The elasticity assurance is being prepared.
	//
	// 	- Prepared: The elasticity assurance is to take effect.
	//
	// 	- Active: The elasticity assurance is in effect.
	//
	// 	- Released: The elasticity assurance is released.
	//
	// If you do not specify this parameter, elasticity assurances in states other than Pending and Released are queried.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags.
	Tag []*DescribeElasticityAssurancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The zone ID of the elasticity assurances.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeElasticityAssurancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticityAssurancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeElasticityAssurancesRequest) GetPrivatePoolOptions() *DescribeElasticityAssurancesRequestPrivatePoolOptions {
	return s.PrivatePoolOptions
}

func (s *DescribeElasticityAssurancesRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *DescribeElasticityAssurancesRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeElasticityAssurancesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeElasticityAssurancesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeElasticityAssurancesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeElasticityAssurancesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeElasticityAssurancesRequest) GetPackageType() *string {
	return s.PackageType
}

func (s *DescribeElasticityAssurancesRequest) GetPlatform() *string {
	return s.Platform
}

func (s *DescribeElasticityAssurancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeElasticityAssurancesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeElasticityAssurancesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeElasticityAssurancesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeElasticityAssurancesRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeElasticityAssurancesRequest) GetTag() []*DescribeElasticityAssurancesRequestTag {
	return s.Tag
}

func (s *DescribeElasticityAssurancesRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeElasticityAssurancesRequest) SetPrivatePoolOptions(v *DescribeElasticityAssurancesRequestPrivatePoolOptions) *DescribeElasticityAssurancesRequest {
	s.PrivatePoolOptions = v
	return s
}

func (s *DescribeElasticityAssurancesRequest) SetInstanceChargeType(v string) *DescribeElasticityAssurancesRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeElasticityAssurancesRequest) SetInstanceType(v string) *DescribeElasticityAssurancesRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeElasticityAssurancesRequest) SetMaxResults(v int32) *DescribeElasticityAssurancesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeElasticityAssurancesRequest) SetNextToken(v string) *DescribeElasticityAssurancesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeElasticityAssurancesRequest) SetOwnerAccount(v string) *DescribeElasticityAssurancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeElasticityAssurancesRequest) SetOwnerId(v int64) *DescribeElasticityAssurancesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeElasticityAssurancesRequest) SetPackageType(v string) *DescribeElasticityAssurancesRequest {
	s.PackageType = &v
	return s
}

func (s *DescribeElasticityAssurancesRequest) SetPlatform(v string) *DescribeElasticityAssurancesRequest {
	s.Platform = &v
	return s
}

func (s *DescribeElasticityAssurancesRequest) SetRegionId(v string) *DescribeElasticityAssurancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeElasticityAssurancesRequest) SetResourceGroupId(v string) *DescribeElasticityAssurancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeElasticityAssurancesRequest) SetResourceOwnerAccount(v string) *DescribeElasticityAssurancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeElasticityAssurancesRequest) SetResourceOwnerId(v int64) *DescribeElasticityAssurancesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeElasticityAssurancesRequest) SetStatus(v string) *DescribeElasticityAssurancesRequest {
	s.Status = &v
	return s
}

func (s *DescribeElasticityAssurancesRequest) SetTag(v []*DescribeElasticityAssurancesRequestTag) *DescribeElasticityAssurancesRequest {
	s.Tag = v
	return s
}

func (s *DescribeElasticityAssurancesRequest) SetZoneId(v string) *DescribeElasticityAssurancesRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeElasticityAssurancesRequest) Validate() error {
	if s.PrivatePoolOptions != nil {
		if err := s.PrivatePoolOptions.Validate(); err != nil {
			return err
		}
	}
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeElasticityAssurancesRequestPrivatePoolOptions struct {
	// The IDs of the elasticity assurances. The value can be a JSON array that consists of up to 100 elasticity assurance IDs. Separate the IDs with commas (,).
	//
	// example:
	//
	// ["eap-bp67acfmxazb4****", "eap-bp67acfmxazb5****"]
	Ids *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
}

func (s DescribeElasticityAssurancesRequestPrivatePoolOptions) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticityAssurancesRequestPrivatePoolOptions) GoString() string {
	return s.String()
}

func (s *DescribeElasticityAssurancesRequestPrivatePoolOptions) GetIds() *string {
	return s.Ids
}

func (s *DescribeElasticityAssurancesRequestPrivatePoolOptions) SetIds(v string) *DescribeElasticityAssurancesRequestPrivatePoolOptions {
	s.Ids = &v
	return s
}

func (s *DescribeElasticityAssurancesRequestPrivatePoolOptions) Validate() error {
	return dara.Validate(s)
}

type DescribeElasticityAssurancesRequestTag struct {
	// The key of tag N. Valid values of N: 1 to 20.
	//
	// If you specify a tag to query resources, up to 1,000 resources with this tag are returned in the response. If you specify multiple tags to query resources, up to 1,000 resources with all these tags are returned in the response. To query more than 1,000 resources that have specified tags added, call the [ListTagResources](https://help.aliyun.com/document_detail/110425.html) operation.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N. Valid values of N: 1 to 20.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeElasticityAssurancesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticityAssurancesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeElasticityAssurancesRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeElasticityAssurancesRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeElasticityAssurancesRequestTag) SetKey(v string) *DescribeElasticityAssurancesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeElasticityAssurancesRequestTag) SetValue(v string) *DescribeElasticityAssurancesRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeElasticityAssurancesRequestTag) Validate() error {
	return dara.Validate(s)
}
