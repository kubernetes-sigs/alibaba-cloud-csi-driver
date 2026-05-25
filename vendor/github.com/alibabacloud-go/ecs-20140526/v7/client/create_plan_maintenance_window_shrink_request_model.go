// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePlanMaintenanceWindowShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnable(v bool) *CreatePlanMaintenanceWindowShrinkRequest
	GetEnable() *bool
	SetPlanWindowName(v string) *CreatePlanMaintenanceWindowShrinkRequest
	GetPlanWindowName() *string
	SetRegionId(v string) *CreatePlanMaintenanceWindowShrinkRequest
	GetRegionId() *string
	SetSupportMaintenanceAction(v string) *CreatePlanMaintenanceWindowShrinkRequest
	GetSupportMaintenanceAction() *string
	SetTargetResourceShrink(v string) *CreatePlanMaintenanceWindowShrinkRequest
	GetTargetResourceShrink() *string
	SetTimePeriodShrink(v string) *CreatePlanMaintenanceWindowShrinkRequest
	GetTimePeriodShrink() *string
}

type CreatePlanMaintenanceWindowShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// This parameter is required.
	PlanWindowName *string `json:"PlanWindowName,omitempty" xml:"PlanWindowName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Reboot
	SupportMaintenanceAction *string `json:"SupportMaintenanceAction,omitempty" xml:"SupportMaintenanceAction,omitempty"`
	// This parameter is required.
	TargetResourceShrink *string `json:"TargetResource,omitempty" xml:"TargetResource,omitempty"`
	// This parameter is required.
	TimePeriodShrink *string `json:"TimePeriod,omitempty" xml:"TimePeriod,omitempty"`
}

func (s CreatePlanMaintenanceWindowShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePlanMaintenanceWindowShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreatePlanMaintenanceWindowShrinkRequest) GetEnable() *bool {
	return s.Enable
}

func (s *CreatePlanMaintenanceWindowShrinkRequest) GetPlanWindowName() *string {
	return s.PlanWindowName
}

func (s *CreatePlanMaintenanceWindowShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreatePlanMaintenanceWindowShrinkRequest) GetSupportMaintenanceAction() *string {
	return s.SupportMaintenanceAction
}

func (s *CreatePlanMaintenanceWindowShrinkRequest) GetTargetResourceShrink() *string {
	return s.TargetResourceShrink
}

func (s *CreatePlanMaintenanceWindowShrinkRequest) GetTimePeriodShrink() *string {
	return s.TimePeriodShrink
}

func (s *CreatePlanMaintenanceWindowShrinkRequest) SetEnable(v bool) *CreatePlanMaintenanceWindowShrinkRequest {
	s.Enable = &v
	return s
}

func (s *CreatePlanMaintenanceWindowShrinkRequest) SetPlanWindowName(v string) *CreatePlanMaintenanceWindowShrinkRequest {
	s.PlanWindowName = &v
	return s
}

func (s *CreatePlanMaintenanceWindowShrinkRequest) SetRegionId(v string) *CreatePlanMaintenanceWindowShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePlanMaintenanceWindowShrinkRequest) SetSupportMaintenanceAction(v string) *CreatePlanMaintenanceWindowShrinkRequest {
	s.SupportMaintenanceAction = &v
	return s
}

func (s *CreatePlanMaintenanceWindowShrinkRequest) SetTargetResourceShrink(v string) *CreatePlanMaintenanceWindowShrinkRequest {
	s.TargetResourceShrink = &v
	return s
}

func (s *CreatePlanMaintenanceWindowShrinkRequest) SetTimePeriodShrink(v string) *CreatePlanMaintenanceWindowShrinkRequest {
	s.TimePeriodShrink = &v
	return s
}

func (s *CreatePlanMaintenanceWindowShrinkRequest) Validate() error {
	return dara.Validate(s)
}
