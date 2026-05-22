// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoProvisioningGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAutoProvisioningGroups(v *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroups) *DescribeAutoProvisioningGroupsResponseBody
	GetAutoProvisioningGroups() *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroups
	SetPageNumber(v int32) *DescribeAutoProvisioningGroupsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAutoProvisioningGroupsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeAutoProvisioningGroupsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeAutoProvisioningGroupsResponseBody
	GetTotalCount() *int32
}

type DescribeAutoProvisioningGroupsResponseBody struct {
	AutoProvisioningGroups *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroups `json:"AutoProvisioningGroups,omitempty" xml:"AutoProvisioningGroups,omitempty" type:"Struct"`
	// The number of the page returned.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 745CEC9F-0DD7-4451-9FE7-8B752F39****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of queried auto provisioning groups.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAutoProvisioningGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoProvisioningGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAutoProvisioningGroupsResponseBody) GetAutoProvisioningGroups() *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroups {
	return s.AutoProvisioningGroups
}

func (s *DescribeAutoProvisioningGroupsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAutoProvisioningGroupsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAutoProvisioningGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAutoProvisioningGroupsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeAutoProvisioningGroupsResponseBody) SetAutoProvisioningGroups(v *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroups) *DescribeAutoProvisioningGroupsResponseBody {
	s.AutoProvisioningGroups = v
	return s
}

func (s *DescribeAutoProvisioningGroupsResponseBody) SetPageNumber(v int32) *DescribeAutoProvisioningGroupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAutoProvisioningGroupsResponseBody) SetPageSize(v int32) *DescribeAutoProvisioningGroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAutoProvisioningGroupsResponseBody) SetRequestId(v string) *DescribeAutoProvisioningGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAutoProvisioningGroupsResponseBody) SetTotalCount(v int32) *DescribeAutoProvisioningGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAutoProvisioningGroupsResponseBody) Validate() error {
	if s.AutoProvisioningGroups != nil {
		if err := s.AutoProvisioningGroups.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroups struct {
	AutoProvisioningGroup []*DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup `json:"AutoProvisioningGroup,omitempty" xml:"AutoProvisioningGroup,omitempty" type:"Repeated"`
}

func (s DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroups) GoString() string {
	return s.String()
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroups) GetAutoProvisioningGroup() []*DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup {
	return s.AutoProvisioningGroup
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroups) SetAutoProvisioningGroup(v []*DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup) *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroups {
	s.AutoProvisioningGroup = v
	return s
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroups) Validate() error {
	if s.AutoProvisioningGroup != nil {
		for _, item := range s.AutoProvisioningGroup {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup struct {
	AutoProvisioningGroupId          *string                                                                                                           `json:"AutoProvisioningGroupId,omitempty" xml:"AutoProvisioningGroupId,omitempty"`
	AutoProvisioningGroupName        *string                                                                                                           `json:"AutoProvisioningGroupName,omitempty" xml:"AutoProvisioningGroupName,omitempty"`
	AutoProvisioningGroupType        *string                                                                                                           `json:"AutoProvisioningGroupType,omitempty" xml:"AutoProvisioningGroupType,omitempty"`
	CandidateOptions                 *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupCandidateOptions            `json:"CandidateOptions,omitempty" xml:"CandidateOptions,omitempty" type:"Struct"`
	CapacitySpecification            *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupCapacitySpecification       `json:"CapacitySpecification,omitempty" xml:"CapacitySpecification,omitempty" type:"Struct"`
	CreationTime                     *string                                                                                                           `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	ExcessCapacityTerminationPolicy  *string                                                                                                           `json:"ExcessCapacityTerminationPolicy,omitempty" xml:"ExcessCapacityTerminationPolicy,omitempty"`
	LaunchTemplateConfigs            *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupLaunchTemplateConfigs       `json:"LaunchTemplateConfigs,omitempty" xml:"LaunchTemplateConfigs,omitempty" type:"Struct"`
	LaunchTemplateId                 *string                                                                                                           `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty"`
	LaunchTemplateVersion            *string                                                                                                           `json:"LaunchTemplateVersion,omitempty" xml:"LaunchTemplateVersion,omitempty"`
	MaxSpotPrice                     *float32                                                                                                          `json:"MaxSpotPrice,omitempty" xml:"MaxSpotPrice,omitempty"`
	PayAsYouGoOptions                *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupPayAsYouGoOptions           `json:"PayAsYouGoOptions,omitempty" xml:"PayAsYouGoOptions,omitempty" type:"Struct"`
	RegionId                         *string                                                                                                           `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId                  *string                                                                                                           `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SpotOptions                      *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupSpotOptions                 `json:"SpotOptions,omitempty" xml:"SpotOptions,omitempty" type:"Struct"`
	State                            *string                                                                                                           `json:"State,omitempty" xml:"State,omitempty"`
	Status                           *string                                                                                                           `json:"Status,omitempty" xml:"Status,omitempty"`
	SuspendedProcesses               *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupSuspendedProcesses          `json:"SuspendedProcesses,omitempty" xml:"SuspendedProcesses,omitempty" type:"Struct"`
	Tags                             *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupTags                        `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	TargetCapacitySpecification      *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupTargetCapacitySpecification `json:"TargetCapacitySpecification,omitempty" xml:"TargetCapacitySpecification,omitempty" type:"Struct"`
	TerminateInstances               *bool                                                                                                             `json:"TerminateInstances,omitempty" xml:"TerminateInstances,omitempty"`
	TerminateInstancesWithExpiration *bool                                                                                                             `json:"TerminateInstancesWithExpiration,omitempty" xml:"TerminateInstancesWithExpiration,omitempty"`
	ValidFrom                        *string                                                                                                           `json:"ValidFrom,omitempty" xml:"ValidFrom,omitempty"`
	ValidUntil                       *string                                                                                                           `json:"ValidUntil,omitempty" xml:"ValidUntil,omitempty"`
}

func (s DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup) GoString() string {
	return s.String()
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup) GetAutoProvisioningGroupId() *string {
	return s.AutoProvisioningGroupId
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup) GetAutoProvisioningGroupName() *string {
	return s.AutoProvisioningGroupName
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup) GetAutoProvisioningGroupType() *string {
	return s.AutoProvisioningGroupType
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup) GetCandidateOptions() *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupCandidateOptions {
	return s.CandidateOptions
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup) GetCapacitySpecification() *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupCapacitySpecification {
	return s.CapacitySpecification
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup) GetExcessCapacityTerminationPolicy() *string {
	return s.ExcessCapacityTerminationPolicy
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup) GetLaunchTemplateConfigs() *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupLaunchTemplateConfigs {
	return s.LaunchTemplateConfigs
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup) GetLaunchTemplateId() *string {
	return s.LaunchTemplateId
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup) GetLaunchTemplateVersion() *string {
	return s.LaunchTemplateVersion
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup) GetMaxSpotPrice() *float32 {
	return s.MaxSpotPrice
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup) GetPayAsYouGoOptions() *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupPayAsYouGoOptions {
	return s.PayAsYouGoOptions
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup) GetSpotOptions() *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupSpotOptions {
	return s.SpotOptions
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup) GetState() *string {
	return s.State
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup) GetStatus() *string {
	return s.Status
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup) GetSuspendedProcesses() *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupSuspendedProcesses {
	return s.SuspendedProcesses
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup) GetTags() *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupTags {
	return s.Tags
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup) GetTargetCapacitySpecification() *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupTargetCapacitySpecification {
	return s.TargetCapacitySpecification
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup) GetTerminateInstances() *bool {
	return s.TerminateInstances
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup) GetTerminateInstancesWithExpiration() *bool {
	return s.TerminateInstancesWithExpiration
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup) GetValidFrom() *string {
	return s.ValidFrom
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup) GetValidUntil() *string {
	return s.ValidUntil
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup) SetAutoProvisioningGroupId(v string) *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup {
	s.AutoProvisioningGroupId = &v
	return s
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup) SetAutoProvisioningGroupName(v string) *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup {
	s.AutoProvisioningGroupName = &v
	return s
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup) SetAutoProvisioningGroupType(v string) *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup {
	s.AutoProvisioningGroupType = &v
	return s
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup) SetCandidateOptions(v *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupCandidateOptions) *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup {
	s.CandidateOptions = v
	return s
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup) SetCapacitySpecification(v *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupCapacitySpecification) *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup {
	s.CapacitySpecification = v
	return s
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup) SetCreationTime(v string) *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup {
	s.CreationTime = &v
	return s
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup) SetExcessCapacityTerminationPolicy(v string) *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup {
	s.ExcessCapacityTerminationPolicy = &v
	return s
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup) SetLaunchTemplateConfigs(v *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupLaunchTemplateConfigs) *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup {
	s.LaunchTemplateConfigs = v
	return s
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup) SetLaunchTemplateId(v string) *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup {
	s.LaunchTemplateId = &v
	return s
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup) SetLaunchTemplateVersion(v string) *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup {
	s.LaunchTemplateVersion = &v
	return s
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup) SetMaxSpotPrice(v float32) *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup {
	s.MaxSpotPrice = &v
	return s
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup) SetPayAsYouGoOptions(v *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupPayAsYouGoOptions) *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup {
	s.PayAsYouGoOptions = v
	return s
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup) SetRegionId(v string) *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup {
	s.RegionId = &v
	return s
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup) SetResourceGroupId(v string) *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup) SetSpotOptions(v *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupSpotOptions) *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup {
	s.SpotOptions = v
	return s
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup) SetState(v string) *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup {
	s.State = &v
	return s
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup) SetStatus(v string) *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup {
	s.Status = &v
	return s
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup) SetSuspendedProcesses(v *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupSuspendedProcesses) *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup {
	s.SuspendedProcesses = v
	return s
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup) SetTags(v *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupTags) *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup {
	s.Tags = v
	return s
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup) SetTargetCapacitySpecification(v *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupTargetCapacitySpecification) *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup {
	s.TargetCapacitySpecification = v
	return s
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup) SetTerminateInstances(v bool) *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup {
	s.TerminateInstances = &v
	return s
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup) SetTerminateInstancesWithExpiration(v bool) *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup {
	s.TerminateInstancesWithExpiration = &v
	return s
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup) SetValidFrom(v string) *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup {
	s.ValidFrom = &v
	return s
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup) SetValidUntil(v string) *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup {
	s.ValidUntil = &v
	return s
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroup) Validate() error {
	if s.CandidateOptions != nil {
		if err := s.CandidateOptions.Validate(); err != nil {
			return err
		}
	}
	if s.CapacitySpecification != nil {
		if err := s.CapacitySpecification.Validate(); err != nil {
			return err
		}
	}
	if s.LaunchTemplateConfigs != nil {
		if err := s.LaunchTemplateConfigs.Validate(); err != nil {
			return err
		}
	}
	if s.PayAsYouGoOptions != nil {
		if err := s.PayAsYouGoOptions.Validate(); err != nil {
			return err
		}
	}
	if s.SpotOptions != nil {
		if err := s.SpotOptions.Validate(); err != nil {
			return err
		}
	}
	if s.SuspendedProcesses != nil {
		if err := s.SuspendedProcesses.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	if s.TargetCapacitySpecification != nil {
		if err := s.TargetCapacitySpecification.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupCandidateOptions struct {
	// example:
	//
	// 60
	TimeoutMinutes *int32 `json:"TimeoutMinutes,omitempty" xml:"TimeoutMinutes,omitempty"`
}

func (s DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupCandidateOptions) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupCandidateOptions) GoString() string {
	return s.String()
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupCandidateOptions) GetTimeoutMinutes() *int32 {
	return s.TimeoutMinutes
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupCandidateOptions) SetTimeoutMinutes(v int32) *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupCandidateOptions {
	s.TimeoutMinutes = &v
	return s
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupCandidateOptions) Validate() error {
	return dara.Validate(s)
}

type DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupCapacitySpecification struct {
	// example:
	//
	// 2
	PayAsYouGoCapacity *float32 `json:"PayAsYouGoCapacity,omitempty" xml:"PayAsYouGoCapacity,omitempty"`
	// example:
	//
	// 0
	PrePaidCapacity *float32 `json:"PrePaidCapacity,omitempty" xml:"PrePaidCapacity,omitempty"`
	// example:
	//
	// 3
	SpotCapacity *float32 `json:"SpotCapacity,omitempty" xml:"SpotCapacity,omitempty"`
	// example:
	//
	// 5
	TotalCapacity *float32 `json:"TotalCapacity,omitempty" xml:"TotalCapacity,omitempty"`
}

func (s DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupCapacitySpecification) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupCapacitySpecification) GoString() string {
	return s.String()
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupCapacitySpecification) GetPayAsYouGoCapacity() *float32 {
	return s.PayAsYouGoCapacity
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupCapacitySpecification) GetPrePaidCapacity() *float32 {
	return s.PrePaidCapacity
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupCapacitySpecification) GetSpotCapacity() *float32 {
	return s.SpotCapacity
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupCapacitySpecification) GetTotalCapacity() *float32 {
	return s.TotalCapacity
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupCapacitySpecification) SetPayAsYouGoCapacity(v float32) *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupCapacitySpecification {
	s.PayAsYouGoCapacity = &v
	return s
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupCapacitySpecification) SetPrePaidCapacity(v float32) *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupCapacitySpecification {
	s.PrePaidCapacity = &v
	return s
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupCapacitySpecification) SetSpotCapacity(v float32) *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupCapacitySpecification {
	s.SpotCapacity = &v
	return s
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupCapacitySpecification) SetTotalCapacity(v float32) *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupCapacitySpecification {
	s.TotalCapacity = &v
	return s
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupCapacitySpecification) Validate() error {
	return dara.Validate(s)
}

type DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupLaunchTemplateConfigs struct {
	LaunchTemplateConfig []*DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupLaunchTemplateConfigsLaunchTemplateConfig `json:"LaunchTemplateConfig,omitempty" xml:"LaunchTemplateConfig,omitempty" type:"Repeated"`
}

func (s DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupLaunchTemplateConfigs) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupLaunchTemplateConfigs) GoString() string {
	return s.String()
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupLaunchTemplateConfigs) GetLaunchTemplateConfig() []*DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupLaunchTemplateConfigsLaunchTemplateConfig {
	return s.LaunchTemplateConfig
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupLaunchTemplateConfigs) SetLaunchTemplateConfig(v []*DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupLaunchTemplateConfigsLaunchTemplateConfig) *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupLaunchTemplateConfigs {
	s.LaunchTemplateConfig = v
	return s
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupLaunchTemplateConfigs) Validate() error {
	if s.LaunchTemplateConfig != nil {
		for _, item := range s.LaunchTemplateConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupLaunchTemplateConfigsLaunchTemplateConfig struct {
	InstanceType     *string  `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	MaxPrice         *float32 `json:"MaxPrice,omitempty" xml:"MaxPrice,omitempty"`
	Priority         *float32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	VSwitchId        *string  `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	WeightedCapacity *float32 `json:"WeightedCapacity,omitempty" xml:"WeightedCapacity,omitempty"`
}

func (s DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupLaunchTemplateConfigsLaunchTemplateConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupLaunchTemplateConfigsLaunchTemplateConfig) GoString() string {
	return s.String()
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupLaunchTemplateConfigsLaunchTemplateConfig) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupLaunchTemplateConfigsLaunchTemplateConfig) GetMaxPrice() *float32 {
	return s.MaxPrice
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupLaunchTemplateConfigsLaunchTemplateConfig) GetPriority() *float32 {
	return s.Priority
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupLaunchTemplateConfigsLaunchTemplateConfig) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupLaunchTemplateConfigsLaunchTemplateConfig) GetWeightedCapacity() *float32 {
	return s.WeightedCapacity
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupLaunchTemplateConfigsLaunchTemplateConfig) SetInstanceType(v string) *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupLaunchTemplateConfigsLaunchTemplateConfig {
	s.InstanceType = &v
	return s
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupLaunchTemplateConfigsLaunchTemplateConfig) SetMaxPrice(v float32) *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupLaunchTemplateConfigsLaunchTemplateConfig {
	s.MaxPrice = &v
	return s
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupLaunchTemplateConfigsLaunchTemplateConfig) SetPriority(v float32) *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupLaunchTemplateConfigsLaunchTemplateConfig {
	s.Priority = &v
	return s
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupLaunchTemplateConfigsLaunchTemplateConfig) SetVSwitchId(v string) *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupLaunchTemplateConfigsLaunchTemplateConfig {
	s.VSwitchId = &v
	return s
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupLaunchTemplateConfigsLaunchTemplateConfig) SetWeightedCapacity(v float32) *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupLaunchTemplateConfigsLaunchTemplateConfig {
	s.WeightedCapacity = &v
	return s
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupLaunchTemplateConfigsLaunchTemplateConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupPayAsYouGoOptions struct {
	AllocationStrategy *string `json:"AllocationStrategy,omitempty" xml:"AllocationStrategy,omitempty"`
}

func (s DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupPayAsYouGoOptions) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupPayAsYouGoOptions) GoString() string {
	return s.String()
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupPayAsYouGoOptions) GetAllocationStrategy() *string {
	return s.AllocationStrategy
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupPayAsYouGoOptions) SetAllocationStrategy(v string) *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupPayAsYouGoOptions {
	s.AllocationStrategy = &v
	return s
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupPayAsYouGoOptions) Validate() error {
	return dara.Validate(s)
}

type DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupSpotOptions struct {
	AllocationStrategy           *string `json:"AllocationStrategy,omitempty" xml:"AllocationStrategy,omitempty"`
	InstanceInterruptionBehavior *string `json:"InstanceInterruptionBehavior,omitempty" xml:"InstanceInterruptionBehavior,omitempty"`
	InstancePoolsToUseCount      *int32  `json:"InstancePoolsToUseCount,omitempty" xml:"InstancePoolsToUseCount,omitempty"`
}

func (s DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupSpotOptions) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupSpotOptions) GoString() string {
	return s.String()
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupSpotOptions) GetAllocationStrategy() *string {
	return s.AllocationStrategy
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupSpotOptions) GetInstanceInterruptionBehavior() *string {
	return s.InstanceInterruptionBehavior
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupSpotOptions) GetInstancePoolsToUseCount() *int32 {
	return s.InstancePoolsToUseCount
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupSpotOptions) SetAllocationStrategy(v string) *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupSpotOptions {
	s.AllocationStrategy = &v
	return s
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupSpotOptions) SetInstanceInterruptionBehavior(v string) *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupSpotOptions {
	s.InstanceInterruptionBehavior = &v
	return s
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupSpotOptions) SetInstancePoolsToUseCount(v int32) *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupSpotOptions {
	s.InstancePoolsToUseCount = &v
	return s
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupSpotOptions) Validate() error {
	return dara.Validate(s)
}

type DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupSuspendedProcesses struct {
	SuspendedProcess []*string `json:"SuspendedProcess,omitempty" xml:"SuspendedProcess,omitempty" type:"Repeated"`
}

func (s DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupSuspendedProcesses) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupSuspendedProcesses) GoString() string {
	return s.String()
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupSuspendedProcesses) GetSuspendedProcess() []*string {
	return s.SuspendedProcess
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupSuspendedProcesses) SetSuspendedProcess(v []*string) *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupSuspendedProcesses {
	s.SuspendedProcess = v
	return s
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupSuspendedProcesses) Validate() error {
	return dara.Validate(s)
}

type DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupTags struct {
	Tag []*DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupTags) GoString() string {
	return s.String()
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupTags) GetTag() []*DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupTagsTag {
	return s.Tag
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupTags) SetTag(v []*DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupTagsTag) *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupTags {
	s.Tag = v
	return s
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupTags) Validate() error {
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

type DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupTagsTag struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupTagsTag) SetTagKey(v string) *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupTagsTag {
	s.TagKey = &v
	return s
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupTagsTag) SetTagValue(v string) *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupTagsTag {
	s.TagValue = &v
	return s
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupTagsTag) Validate() error {
	return dara.Validate(s)
}

type DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupTargetCapacitySpecification struct {
	DefaultTargetCapacityType *string  `json:"DefaultTargetCapacityType,omitempty" xml:"DefaultTargetCapacityType,omitempty"`
	PayAsYouGoTargetCapacity  *float32 `json:"PayAsYouGoTargetCapacity,omitempty" xml:"PayAsYouGoTargetCapacity,omitempty"`
	SpotTargetCapacity        *float32 `json:"SpotTargetCapacity,omitempty" xml:"SpotTargetCapacity,omitempty"`
	TotalTargetCapacity       *float32 `json:"TotalTargetCapacity,omitempty" xml:"TotalTargetCapacity,omitempty"`
}

func (s DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupTargetCapacitySpecification) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupTargetCapacitySpecification) GoString() string {
	return s.String()
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupTargetCapacitySpecification) GetDefaultTargetCapacityType() *string {
	return s.DefaultTargetCapacityType
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupTargetCapacitySpecification) GetPayAsYouGoTargetCapacity() *float32 {
	return s.PayAsYouGoTargetCapacity
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupTargetCapacitySpecification) GetSpotTargetCapacity() *float32 {
	return s.SpotTargetCapacity
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupTargetCapacitySpecification) GetTotalTargetCapacity() *float32 {
	return s.TotalTargetCapacity
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupTargetCapacitySpecification) SetDefaultTargetCapacityType(v string) *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupTargetCapacitySpecification {
	s.DefaultTargetCapacityType = &v
	return s
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupTargetCapacitySpecification) SetPayAsYouGoTargetCapacity(v float32) *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupTargetCapacitySpecification {
	s.PayAsYouGoTargetCapacity = &v
	return s
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupTargetCapacitySpecification) SetSpotTargetCapacity(v float32) *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupTargetCapacitySpecification {
	s.SpotTargetCapacity = &v
	return s
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupTargetCapacitySpecification) SetTotalTargetCapacity(v float32) *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupTargetCapacitySpecification {
	s.TotalTargetCapacity = &v
	return s
}

func (s *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupTargetCapacitySpecification) Validate() error {
	return dara.Validate(s)
}
