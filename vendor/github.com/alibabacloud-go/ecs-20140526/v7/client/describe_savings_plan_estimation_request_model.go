// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSavingsPlanEstimationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEstimationResource(v string) *DescribeSavingsPlanEstimationRequest
	GetEstimationResource() *string
	SetInstanceTypeScope(v string) *DescribeSavingsPlanEstimationRequest
	GetInstanceTypeScope() *string
	SetOfferingType(v string) *DescribeSavingsPlanEstimationRequest
	GetOfferingType() *string
	SetPeriod(v string) *DescribeSavingsPlanEstimationRequest
	GetPeriod() *string
	SetPeriodUnit(v string) *DescribeSavingsPlanEstimationRequest
	GetPeriodUnit() *string
	SetPlanType(v string) *DescribeSavingsPlanEstimationRequest
	GetPlanType() *string
	SetRegionId(v string) *DescribeSavingsPlanEstimationRequest
	GetRegionId() *string
	SetResourceId(v string) *DescribeSavingsPlanEstimationRequest
	GetResourceId() *string
}

type DescribeSavingsPlanEstimationRequest struct {
	EstimationResource *string `json:"EstimationResource,omitempty" xml:"EstimationResource,omitempty"`
	InstanceTypeScope  *string `json:"InstanceTypeScope,omitempty" xml:"InstanceTypeScope,omitempty"`
	OfferingType       *string `json:"OfferingType,omitempty" xml:"OfferingType,omitempty"`
	Period             *string `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit         *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	PlanType           *string `json:"PlanType,omitempty" xml:"PlanType,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId         *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s DescribeSavingsPlanEstimationRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSavingsPlanEstimationRequest) GoString() string {
	return s.String()
}

func (s *DescribeSavingsPlanEstimationRequest) GetEstimationResource() *string {
	return s.EstimationResource
}

func (s *DescribeSavingsPlanEstimationRequest) GetInstanceTypeScope() *string {
	return s.InstanceTypeScope
}

func (s *DescribeSavingsPlanEstimationRequest) GetOfferingType() *string {
	return s.OfferingType
}

func (s *DescribeSavingsPlanEstimationRequest) GetPeriod() *string {
	return s.Period
}

func (s *DescribeSavingsPlanEstimationRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *DescribeSavingsPlanEstimationRequest) GetPlanType() *string {
	return s.PlanType
}

func (s *DescribeSavingsPlanEstimationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSavingsPlanEstimationRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeSavingsPlanEstimationRequest) SetEstimationResource(v string) *DescribeSavingsPlanEstimationRequest {
	s.EstimationResource = &v
	return s
}

func (s *DescribeSavingsPlanEstimationRequest) SetInstanceTypeScope(v string) *DescribeSavingsPlanEstimationRequest {
	s.InstanceTypeScope = &v
	return s
}

func (s *DescribeSavingsPlanEstimationRequest) SetOfferingType(v string) *DescribeSavingsPlanEstimationRequest {
	s.OfferingType = &v
	return s
}

func (s *DescribeSavingsPlanEstimationRequest) SetPeriod(v string) *DescribeSavingsPlanEstimationRequest {
	s.Period = &v
	return s
}

func (s *DescribeSavingsPlanEstimationRequest) SetPeriodUnit(v string) *DescribeSavingsPlanEstimationRequest {
	s.PeriodUnit = &v
	return s
}

func (s *DescribeSavingsPlanEstimationRequest) SetPlanType(v string) *DescribeSavingsPlanEstimationRequest {
	s.PlanType = &v
	return s
}

func (s *DescribeSavingsPlanEstimationRequest) SetRegionId(v string) *DescribeSavingsPlanEstimationRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSavingsPlanEstimationRequest) SetResourceId(v string) *DescribeSavingsPlanEstimationRequest {
	s.ResourceId = &v
	return s
}

func (s *DescribeSavingsPlanEstimationRequest) Validate() error {
	return dara.Validate(s)
}
