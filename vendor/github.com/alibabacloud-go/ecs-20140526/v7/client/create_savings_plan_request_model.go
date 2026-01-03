// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSavingsPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChargeType(v string) *CreateSavingsPlanRequest
	GetChargeType() *string
	SetCommittedAmount(v string) *CreateSavingsPlanRequest
	GetCommittedAmount() *string
	SetInstanceTypeFamily(v string) *CreateSavingsPlanRequest
	GetInstanceTypeFamily() *string
	SetInstanceTypeFamilyGroup(v string) *CreateSavingsPlanRequest
	GetInstanceTypeFamilyGroup() *string
	SetOfferingType(v string) *CreateSavingsPlanRequest
	GetOfferingType() *string
	SetPeriod(v string) *CreateSavingsPlanRequest
	GetPeriod() *string
	SetPeriodUnit(v string) *CreateSavingsPlanRequest
	GetPeriodUnit() *string
	SetPlanType(v string) *CreateSavingsPlanRequest
	GetPlanType() *string
	SetRegionId(v string) *CreateSavingsPlanRequest
	GetRegionId() *string
	SetResourceId(v []*string) *CreateSavingsPlanRequest
	GetResourceId() []*string
}

type CreateSavingsPlanRequest struct {
	ChargeType              *string   `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	CommittedAmount         *string   `json:"CommittedAmount,omitempty" xml:"CommittedAmount,omitempty"`
	InstanceTypeFamily      *string   `json:"InstanceTypeFamily,omitempty" xml:"InstanceTypeFamily,omitempty"`
	InstanceTypeFamilyGroup *string   `json:"InstanceTypeFamilyGroup,omitempty" xml:"InstanceTypeFamilyGroup,omitempty"`
	OfferingType            *string   `json:"OfferingType,omitempty" xml:"OfferingType,omitempty"`
	Period                  *string   `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit              *string   `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	PlanType                *string   `json:"PlanType,omitempty" xml:"PlanType,omitempty"`
	RegionId                *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId              []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
}

func (s CreateSavingsPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSavingsPlanRequest) GoString() string {
	return s.String()
}

func (s *CreateSavingsPlanRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateSavingsPlanRequest) GetCommittedAmount() *string {
	return s.CommittedAmount
}

func (s *CreateSavingsPlanRequest) GetInstanceTypeFamily() *string {
	return s.InstanceTypeFamily
}

func (s *CreateSavingsPlanRequest) GetInstanceTypeFamilyGroup() *string {
	return s.InstanceTypeFamilyGroup
}

func (s *CreateSavingsPlanRequest) GetOfferingType() *string {
	return s.OfferingType
}

func (s *CreateSavingsPlanRequest) GetPeriod() *string {
	return s.Period
}

func (s *CreateSavingsPlanRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *CreateSavingsPlanRequest) GetPlanType() *string {
	return s.PlanType
}

func (s *CreateSavingsPlanRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateSavingsPlanRequest) GetResourceId() []*string {
	return s.ResourceId
}

func (s *CreateSavingsPlanRequest) SetChargeType(v string) *CreateSavingsPlanRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateSavingsPlanRequest) SetCommittedAmount(v string) *CreateSavingsPlanRequest {
	s.CommittedAmount = &v
	return s
}

func (s *CreateSavingsPlanRequest) SetInstanceTypeFamily(v string) *CreateSavingsPlanRequest {
	s.InstanceTypeFamily = &v
	return s
}

func (s *CreateSavingsPlanRequest) SetInstanceTypeFamilyGroup(v string) *CreateSavingsPlanRequest {
	s.InstanceTypeFamilyGroup = &v
	return s
}

func (s *CreateSavingsPlanRequest) SetOfferingType(v string) *CreateSavingsPlanRequest {
	s.OfferingType = &v
	return s
}

func (s *CreateSavingsPlanRequest) SetPeriod(v string) *CreateSavingsPlanRequest {
	s.Period = &v
	return s
}

func (s *CreateSavingsPlanRequest) SetPeriodUnit(v string) *CreateSavingsPlanRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateSavingsPlanRequest) SetPlanType(v string) *CreateSavingsPlanRequest {
	s.PlanType = &v
	return s
}

func (s *CreateSavingsPlanRequest) SetRegionId(v string) *CreateSavingsPlanRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSavingsPlanRequest) SetResourceId(v []*string) *CreateSavingsPlanRequest {
	s.ResourceId = v
	return s
}

func (s *CreateSavingsPlanRequest) Validate() error {
	return dara.Validate(s)
}
