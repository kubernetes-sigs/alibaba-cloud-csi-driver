// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRenewalPriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExpectedRenewDay(v int32) *DescribeRenewalPriceRequest
	GetExpectedRenewDay() *int32
	SetOwnerAccount(v string) *DescribeRenewalPriceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeRenewalPriceRequest
	GetOwnerId() *int64
	SetPeriod(v int32) *DescribeRenewalPriceRequest
	GetPeriod() *int32
	SetPriceUnit(v string) *DescribeRenewalPriceRequest
	GetPriceUnit() *string
	SetRegionId(v string) *DescribeRenewalPriceRequest
	GetRegionId() *string
	SetResourceId(v string) *DescribeRenewalPriceRequest
	GetResourceId() *string
	SetResourceOwnerAccount(v string) *DescribeRenewalPriceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeRenewalPriceRequest
	GetResourceOwnerId() *int64
	SetResourceType(v string) *DescribeRenewalPriceRequest
	GetResourceType() *string
}

type DescribeRenewalPriceRequest struct {
	// The synchronized expiration date. If you specify this parameter, the price for renewing a specified instance to the specified synchronized expiration date is queried. Valid values: 1 to 28.
	//
	// For information about how to synchronize the expiration dates of instances, see [Synchronize the expiration dates of instances](https://help.aliyun.com/document_detail/108486.html).
	//
	// > The renewal period-related parameter pair (`Period` and `PeriodUnit`) and the `ExpectedRenewDay` parameter are mutually exclusive.
	//
	// example:
	//
	// 5
	ExpectedRenewDay *int32  `json:"ExpectedRenewDay,omitempty" xml:"ExpectedRenewDay,omitempty"`
	OwnerAccount     *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The renewal period. Valid values:
	//
	// 	- Valid values when the `PriceUnit` parameter is set to `Month`: 1, 2, 3, 4, 5, 6, 7, 8, and 9.
	//
	// 	- Valid values when the `PriceUnit` parameter is set to `Year`: 1, 2, 3.
	//
	// Default value: 1.
	//
	// > The renewal period-related parameter pair (`Period` and `PeriodUnit`) and the `ExpectedRenewDay` parameter are mutually exclusive.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the renewal period. Valid values:
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
	PriceUnit *string `json:"PriceUnit,omitempty" xml:"PriceUnit,omitempty"`
	// The region ID of the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent list of regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource. If the `ResourceType` parameter is set to`  instance `, the value of the `ResourceId` parameter is the ID of the specified instance.``
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp1f2o4ldh8l29zv****
	ResourceId           *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of the resource. Set the value to instance.
	//
	// Default value: instance.
	//
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribeRenewalPriceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenewalPriceRequest) GoString() string {
	return s.String()
}

func (s *DescribeRenewalPriceRequest) GetExpectedRenewDay() *int32 {
	return s.ExpectedRenewDay
}

func (s *DescribeRenewalPriceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeRenewalPriceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeRenewalPriceRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *DescribeRenewalPriceRequest) GetPriceUnit() *string {
	return s.PriceUnit
}

func (s *DescribeRenewalPriceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRenewalPriceRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeRenewalPriceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeRenewalPriceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeRenewalPriceRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeRenewalPriceRequest) SetExpectedRenewDay(v int32) *DescribeRenewalPriceRequest {
	s.ExpectedRenewDay = &v
	return s
}

func (s *DescribeRenewalPriceRequest) SetOwnerAccount(v string) *DescribeRenewalPriceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeRenewalPriceRequest) SetOwnerId(v int64) *DescribeRenewalPriceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRenewalPriceRequest) SetPeriod(v int32) *DescribeRenewalPriceRequest {
	s.Period = &v
	return s
}

func (s *DescribeRenewalPriceRequest) SetPriceUnit(v string) *DescribeRenewalPriceRequest {
	s.PriceUnit = &v
	return s
}

func (s *DescribeRenewalPriceRequest) SetRegionId(v string) *DescribeRenewalPriceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRenewalPriceRequest) SetResourceId(v string) *DescribeRenewalPriceRequest {
	s.ResourceId = &v
	return s
}

func (s *DescribeRenewalPriceRequest) SetResourceOwnerAccount(v string) *DescribeRenewalPriceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRenewalPriceRequest) SetResourceOwnerId(v int64) *DescribeRenewalPriceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRenewalPriceRequest) SetResourceType(v string) *DescribeRenewalPriceRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeRenewalPriceRequest) Validate() error {
	return dara.Validate(s)
}
