// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlanMaintenanceWindowsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnable(v bool) *DescribePlanMaintenanceWindowsShrinkRequest
	GetEnable() *bool
	SetMaxResults(v int32) *DescribePlanMaintenanceWindowsShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribePlanMaintenanceWindowsShrinkRequest
	GetNextToken() *string
	SetPlanWindowId(v string) *DescribePlanMaintenanceWindowsShrinkRequest
	GetPlanWindowId() *string
	SetPlanWindowName(v string) *DescribePlanMaintenanceWindowsShrinkRequest
	GetPlanWindowName() *string
	SetRegionId(v string) *DescribePlanMaintenanceWindowsShrinkRequest
	GetRegionId() *string
	SetTargetResourceGroupId(v string) *DescribePlanMaintenanceWindowsShrinkRequest
	GetTargetResourceGroupId() *string
	SetTargetResourceTagsShrink(v string) *DescribePlanMaintenanceWindowsShrinkRequest
	GetTargetResourceTagsShrink() *string
}

type DescribePlanMaintenanceWindowsShrinkRequest struct {
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAdDWBF2
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// pw-bp1jarob1aup7yvlrdt6
	PlanWindowId   *string `json:"PlanWindowId,omitempty" xml:"PlanWindowId,omitempty"`
	PlanWindowName *string `json:"PlanWindowName,omitempty" xml:"PlanWindowName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-d85g5yocioezmdrll
	TargetResourceGroupId    *string `json:"TargetResourceGroupId,omitempty" xml:"TargetResourceGroupId,omitempty"`
	TargetResourceTagsShrink *string `json:"TargetResourceTags,omitempty" xml:"TargetResourceTags,omitempty"`
}

func (s DescribePlanMaintenanceWindowsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePlanMaintenanceWindowsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribePlanMaintenanceWindowsShrinkRequest) GetEnable() *bool {
	return s.Enable
}

func (s *DescribePlanMaintenanceWindowsShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribePlanMaintenanceWindowsShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribePlanMaintenanceWindowsShrinkRequest) GetPlanWindowId() *string {
	return s.PlanWindowId
}

func (s *DescribePlanMaintenanceWindowsShrinkRequest) GetPlanWindowName() *string {
	return s.PlanWindowName
}

func (s *DescribePlanMaintenanceWindowsShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePlanMaintenanceWindowsShrinkRequest) GetTargetResourceGroupId() *string {
	return s.TargetResourceGroupId
}

func (s *DescribePlanMaintenanceWindowsShrinkRequest) GetTargetResourceTagsShrink() *string {
	return s.TargetResourceTagsShrink
}

func (s *DescribePlanMaintenanceWindowsShrinkRequest) SetEnable(v bool) *DescribePlanMaintenanceWindowsShrinkRequest {
	s.Enable = &v
	return s
}

func (s *DescribePlanMaintenanceWindowsShrinkRequest) SetMaxResults(v int32) *DescribePlanMaintenanceWindowsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribePlanMaintenanceWindowsShrinkRequest) SetNextToken(v string) *DescribePlanMaintenanceWindowsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *DescribePlanMaintenanceWindowsShrinkRequest) SetPlanWindowId(v string) *DescribePlanMaintenanceWindowsShrinkRequest {
	s.PlanWindowId = &v
	return s
}

func (s *DescribePlanMaintenanceWindowsShrinkRequest) SetPlanWindowName(v string) *DescribePlanMaintenanceWindowsShrinkRequest {
	s.PlanWindowName = &v
	return s
}

func (s *DescribePlanMaintenanceWindowsShrinkRequest) SetRegionId(v string) *DescribePlanMaintenanceWindowsShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePlanMaintenanceWindowsShrinkRequest) SetTargetResourceGroupId(v string) *DescribePlanMaintenanceWindowsShrinkRequest {
	s.TargetResourceGroupId = &v
	return s
}

func (s *DescribePlanMaintenanceWindowsShrinkRequest) SetTargetResourceTagsShrink(v string) *DescribePlanMaintenanceWindowsShrinkRequest {
	s.TargetResourceTagsShrink = &v
	return s
}

func (s *DescribePlanMaintenanceWindowsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
