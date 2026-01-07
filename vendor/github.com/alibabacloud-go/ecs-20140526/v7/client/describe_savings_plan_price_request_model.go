// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSavingsPlanPriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommittedAmount(v string) *DescribeSavingsPlanPriceRequest
	GetCommittedAmount() *string
	SetInstanceTypeFamily(v string) *DescribeSavingsPlanPriceRequest
	GetInstanceTypeFamily() *string
	SetInstanceTypeFamilyGroup(v string) *DescribeSavingsPlanPriceRequest
	GetInstanceTypeFamilyGroup() *string
	SetOfferingType(v string) *DescribeSavingsPlanPriceRequest
	GetOfferingType() *string
	SetPeriod(v int32) *DescribeSavingsPlanPriceRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *DescribeSavingsPlanPriceRequest
	GetPeriodUnit() *string
	SetPlanType(v string) *DescribeSavingsPlanPriceRequest
	GetPlanType() *string
	SetRegionId(v string) *DescribeSavingsPlanPriceRequest
	GetRegionId() *string
	SetResourceId(v []*string) *DescribeSavingsPlanPriceRequest
	GetResourceId() []*string
}

type DescribeSavingsPlanPriceRequest struct {
	CommittedAmount         *string   `json:"CommittedAmount,omitempty" xml:"CommittedAmount,omitempty"`
	InstanceTypeFamily      *string   `json:"InstanceTypeFamily,omitempty" xml:"InstanceTypeFamily,omitempty"`
	InstanceTypeFamilyGroup *string   `json:"InstanceTypeFamilyGroup,omitempty" xml:"InstanceTypeFamilyGroup,omitempty"`
	OfferingType            *string   `json:"OfferingType,omitempty" xml:"OfferingType,omitempty"`
	Period                  *int32    `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit              *string   `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	PlanType                *string   `json:"PlanType,omitempty" xml:"PlanType,omitempty"`
	RegionId                *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId              []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
}

func (s DescribeSavingsPlanPriceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSavingsPlanPriceRequest) GoString() string {
	return s.String()
}

func (s *DescribeSavingsPlanPriceRequest) GetCommittedAmount() *string {
	return s.CommittedAmount
}

func (s *DescribeSavingsPlanPriceRequest) GetInstanceTypeFamily() *string {
	return s.InstanceTypeFamily
}

func (s *DescribeSavingsPlanPriceRequest) GetInstanceTypeFamilyGroup() *string {
	return s.InstanceTypeFamilyGroup
}

func (s *DescribeSavingsPlanPriceRequest) GetOfferingType() *string {
	return s.OfferingType
}

func (s *DescribeSavingsPlanPriceRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *DescribeSavingsPlanPriceRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *DescribeSavingsPlanPriceRequest) GetPlanType() *string {
	return s.PlanType
}

func (s *DescribeSavingsPlanPriceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSavingsPlanPriceRequest) GetResourceId() []*string {
	return s.ResourceId
}

func (s *DescribeSavingsPlanPriceRequest) SetCommittedAmount(v string) *DescribeSavingsPlanPriceRequest {
	s.CommittedAmount = &v
	return s
}

func (s *DescribeSavingsPlanPriceRequest) SetInstanceTypeFamily(v string) *DescribeSavingsPlanPriceRequest {
	s.InstanceTypeFamily = &v
	return s
}

func (s *DescribeSavingsPlanPriceRequest) SetInstanceTypeFamilyGroup(v string) *DescribeSavingsPlanPriceRequest {
	s.InstanceTypeFamilyGroup = &v
	return s
}

func (s *DescribeSavingsPlanPriceRequest) SetOfferingType(v string) *DescribeSavingsPlanPriceRequest {
	s.OfferingType = &v
	return s
}

func (s *DescribeSavingsPlanPriceRequest) SetPeriod(v int32) *DescribeSavingsPlanPriceRequest {
	s.Period = &v
	return s
}

func (s *DescribeSavingsPlanPriceRequest) SetPeriodUnit(v string) *DescribeSavingsPlanPriceRequest {
	s.PeriodUnit = &v
	return s
}

func (s *DescribeSavingsPlanPriceRequest) SetPlanType(v string) *DescribeSavingsPlanPriceRequest {
	s.PlanType = &v
	return s
}

func (s *DescribeSavingsPlanPriceRequest) SetRegionId(v string) *DescribeSavingsPlanPriceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSavingsPlanPriceRequest) SetResourceId(v []*string) *DescribeSavingsPlanPriceRequest {
	s.ResourceId = v
	return s
}

func (s *DescribeSavingsPlanPriceRequest) Validate() error {
	return dara.Validate(s)
}
