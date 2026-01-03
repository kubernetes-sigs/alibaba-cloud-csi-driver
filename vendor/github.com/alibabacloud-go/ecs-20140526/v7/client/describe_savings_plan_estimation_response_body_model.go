// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSavingsPlanEstimationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCommittedAmount(v string) *DescribeSavingsPlanEstimationResponseBody
	GetCommittedAmount() *string
	SetCurrency(v string) *DescribeSavingsPlanEstimationResponseBody
	GetCurrency() *string
	SetInstanceTypeFamily(v string) *DescribeSavingsPlanEstimationResponseBody
	GetInstanceTypeFamily() *string
	SetInstanceTypeFamilyGroup(v string) *DescribeSavingsPlanEstimationResponseBody
	GetInstanceTypeFamilyGroup() *string
	SetOfferingType(v string) *DescribeSavingsPlanEstimationResponseBody
	GetOfferingType() *string
	SetPeriod(v int32) *DescribeSavingsPlanEstimationResponseBody
	GetPeriod() *int32
	SetPeriodUnit(v string) *DescribeSavingsPlanEstimationResponseBody
	GetPeriodUnit() *string
	SetPlanType(v string) *DescribeSavingsPlanEstimationResponseBody
	GetPlanType() *string
	SetRequestId(v string) *DescribeSavingsPlanEstimationResponseBody
	GetRequestId() *string
	SetResourceId(v string) *DescribeSavingsPlanEstimationResponseBody
	GetResourceId() *string
}

type DescribeSavingsPlanEstimationResponseBody struct {
	CommittedAmount         *string `json:"CommittedAmount,omitempty" xml:"CommittedAmount,omitempty"`
	Currency                *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	InstanceTypeFamily      *string `json:"InstanceTypeFamily,omitempty" xml:"InstanceTypeFamily,omitempty"`
	InstanceTypeFamilyGroup *string `json:"InstanceTypeFamilyGroup,omitempty" xml:"InstanceTypeFamilyGroup,omitempty"`
	OfferingType            *string `json:"OfferingType,omitempty" xml:"OfferingType,omitempty"`
	Period                  *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit              *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	PlanType                *string `json:"PlanType,omitempty" xml:"PlanType,omitempty"`
	RequestId               *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceId              *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s DescribeSavingsPlanEstimationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSavingsPlanEstimationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSavingsPlanEstimationResponseBody) GetCommittedAmount() *string {
	return s.CommittedAmount
}

func (s *DescribeSavingsPlanEstimationResponseBody) GetCurrency() *string {
	return s.Currency
}

func (s *DescribeSavingsPlanEstimationResponseBody) GetInstanceTypeFamily() *string {
	return s.InstanceTypeFamily
}

func (s *DescribeSavingsPlanEstimationResponseBody) GetInstanceTypeFamilyGroup() *string {
	return s.InstanceTypeFamilyGroup
}

func (s *DescribeSavingsPlanEstimationResponseBody) GetOfferingType() *string {
	return s.OfferingType
}

func (s *DescribeSavingsPlanEstimationResponseBody) GetPeriod() *int32 {
	return s.Period
}

func (s *DescribeSavingsPlanEstimationResponseBody) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *DescribeSavingsPlanEstimationResponseBody) GetPlanType() *string {
	return s.PlanType
}

func (s *DescribeSavingsPlanEstimationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSavingsPlanEstimationResponseBody) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeSavingsPlanEstimationResponseBody) SetCommittedAmount(v string) *DescribeSavingsPlanEstimationResponseBody {
	s.CommittedAmount = &v
	return s
}

func (s *DescribeSavingsPlanEstimationResponseBody) SetCurrency(v string) *DescribeSavingsPlanEstimationResponseBody {
	s.Currency = &v
	return s
}

func (s *DescribeSavingsPlanEstimationResponseBody) SetInstanceTypeFamily(v string) *DescribeSavingsPlanEstimationResponseBody {
	s.InstanceTypeFamily = &v
	return s
}

func (s *DescribeSavingsPlanEstimationResponseBody) SetInstanceTypeFamilyGroup(v string) *DescribeSavingsPlanEstimationResponseBody {
	s.InstanceTypeFamilyGroup = &v
	return s
}

func (s *DescribeSavingsPlanEstimationResponseBody) SetOfferingType(v string) *DescribeSavingsPlanEstimationResponseBody {
	s.OfferingType = &v
	return s
}

func (s *DescribeSavingsPlanEstimationResponseBody) SetPeriod(v int32) *DescribeSavingsPlanEstimationResponseBody {
	s.Period = &v
	return s
}

func (s *DescribeSavingsPlanEstimationResponseBody) SetPeriodUnit(v string) *DescribeSavingsPlanEstimationResponseBody {
	s.PeriodUnit = &v
	return s
}

func (s *DescribeSavingsPlanEstimationResponseBody) SetPlanType(v string) *DescribeSavingsPlanEstimationResponseBody {
	s.PlanType = &v
	return s
}

func (s *DescribeSavingsPlanEstimationResponseBody) SetRequestId(v string) *DescribeSavingsPlanEstimationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSavingsPlanEstimationResponseBody) SetResourceId(v string) *DescribeSavingsPlanEstimationResponseBody {
	s.ResourceId = &v
	return s
}

func (s *DescribeSavingsPlanEstimationResponseBody) Validate() error {
	return dara.Validate(s)
}
