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
	// Details about the auto provisioning groups.
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
	// The ID of the auto provisioning group.
	//
	// example:
	//
	// apg-sn54avj8htgvtyh8****
	AutoProvisioningGroupId *string `json:"AutoProvisioningGroupId,omitempty" xml:"AutoProvisioningGroupId,omitempty"`
	// The name of the auto provisioning group.
	//
	// example:
	//
	// EcsDocTest
	AutoProvisioningGroupName *string `json:"AutoProvisioningGroupName,omitempty" xml:"AutoProvisioningGroupName,omitempty"`
	// The delivery type of the auto provisioning group. Valid values:
	//
	// 	- request: one-time delivery. When the auto provisioning group is started, it delivers instances only once. If the instances fail to be delivered, the auto provisioning group does not retry the delivery.
	//
	// 	- maintain: continuous delivery. When the auto provisioning group is started, it attempts to deliver instances that meet the target capacity and monitors the real-time capacity. If the target capacity of the auto provisioning group is not reached, the auto provisioning group continues to create instances until the target capacity is reached.
	//
	// example:
	//
	// maintain
	AutoProvisioningGroupType *string `json:"AutoProvisioningGroupType,omitempty" xml:"AutoProvisioningGroupType,omitempty"`
	// The time when the auto provisioning group was created.
	//
	// example:
	//
	// 2019-04-01T15:10:20Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// Indicates whether to release the scaled-in instances when the real-time capacity of the auto provisioning group exceeds the target capacity and the group is triggered to scale in. Valid values:
	//
	// 	- termination: releases the scaled-in instances.
	//
	// 	- no-termination: only removes the scaled-in instances from the auto provisioning group but does not release the instances.
	//
	// example:
	//
	// termination
	ExcessCapacityTerminationPolicy *string `json:"ExcessCapacityTerminationPolicy,omitempty" xml:"ExcessCapacityTerminationPolicy,omitempty"`
	// Details about the extended configurations.
	LaunchTemplateConfigs *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupLaunchTemplateConfigs `json:"LaunchTemplateConfigs,omitempty" xml:"LaunchTemplateConfigs,omitempty" type:"Struct"`
	// The ID of the launch template associated with the auto provisioning group.
	//
	// example:
	//
	// lt-bp1fgzds4bdogu03****
	LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty"`
	// The version of the launch template associated with the auto provisioning group.
	//
	// example:
	//
	// 1
	LaunchTemplateVersion *string `json:"LaunchTemplateVersion,omitempty" xml:"LaunchTemplateVersion,omitempty"`
	// The maximum price of spot  instances in the auto provisioning group.
	//
	// >  When both the MaxSpotPrice and LaunchTemplateConfig.N.MaxPrice parameters are specified, the smaller one of the two parameter values is used.
	//
	// The LaunchTemplateConfig.N.Priority parameter is set when the auto provisioning group is created, and cannot be modified.
	//
	// example:
	//
	// 2
	MaxSpotPrice *float32 `json:"MaxSpotPrice,omitempty" xml:"MaxSpotPrice,omitempty"`
	// The policies related to pay-as-you-go instances.
	PayAsYouGoOptions *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupPayAsYouGoOptions `json:"PayAsYouGoOptions,omitempty" xml:"PayAsYouGoOptions,omitempty" type:"Struct"`
	// The region ID of the auto provisioning group.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the auto provisioning group belongs.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The policy related to spot instances.
	SpotOptions *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupSpotOptions `json:"SpotOptions,omitempty" xml:"SpotOptions,omitempty" type:"Struct"`
	// The overall status of instance scheduling in the auto provisioning group. Valid values:
	//
	// 	- fulfilled: Scheduling was complete and the instances were delivered.
	//
	// 	- pending-fulfillment: The instances were being created.
	//
	// 	- pending-termination: The instances were being removed.
	//
	// 	- error: An exception occurred during scheduling and the instances were not delivered.
	//
	// example:
	//
	// fulfilled
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The status of the auto provisioning group. Valid values:
	//
	// 	- submitted: The auto provisioning group was created but did not execute scheduling tasks.
	//
	// 	- active: The auto provisioning group was executing scheduling tasks.
	//
	// 	- deleted: The auto provisioning group was deleted.
	//
	// 	- delete-running: The auto provisioning group was being deleted.
	//
	// 	- modifying: The auto provisioning group was being modified.
	//
	// example:
	//
	// submitted
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags that are added to the auto provisioning group.
	Tags *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The settings of the target capacity of the auto provisioning group.
	TargetCapacitySpecification *DescribeAutoProvisioningGroupsResponseBodyAutoProvisioningGroupsAutoProvisioningGroupTargetCapacitySpecification `json:"TargetCapacitySpecification,omitempty" xml:"TargetCapacitySpecification,omitempty" type:"Struct"`
	// Indicates whether to release instances in the auto provisioning group when the auto provisioning group is deleted. Valid values:
	//
	// 	- true: releases the instances.
	//
	// 	- false: only removes the instances from the auto provisioning group but does not release the instances.
	//
	// example:
	//
	// false
	TerminateInstances *bool `json:"TerminateInstances,omitempty" xml:"TerminateInstances,omitempty"`
	// Indicates whether to release instances in the auto provisioning group when the group expires. Valid values:
	//
	// 	- true: releases the instances.
	//
	// 	- false: only removes the instances from the auto provisioning group but does not release the instances.
	//
	// example:
	//
	// true
	TerminateInstancesWithExpiration *bool `json:"TerminateInstancesWithExpiration,omitempty" xml:"TerminateInstancesWithExpiration,omitempty"`
	// The time at which the auto provisioning group is started. The provisioning group is effective until the point in time specified by `ValidUntil`.
	//
	// example:
	//
	// 2019-04-01T15:10:20Z
	ValidFrom *string `json:"ValidFrom,omitempty" xml:"ValidFrom,omitempty"`
	// The time at which the auto provisioning group expires. The period of time between this point in time and the point in time specified by the `ValidFrom` parameter is the validity period of the auto provisioning group.
	//
	// example:
	//
	// 2019-06-01T15:10:20Z
	ValidUntil *string `json:"ValidUntil,omitempty" xml:"ValidUntil,omitempty"`
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
	// The instance type that is specified in the extended configuration.
	//
	// example:
	//
	// ecs.g5.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The maximum price of the instance type specified in the extended configuration.
	//
	// example:
	//
	// 3
	MaxPrice *float32 `json:"MaxPrice,omitempty" xml:"MaxPrice,omitempty"`
	// The priority of the instance type specified in the extended configuration. A value of 0 indicates the highest priority.
	//
	// example:
	//
	// 1
	Priority *float32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ID of the vSwitch specified in the extended configuration.
	//
	// example:
	//
	// vsw-sn5bsitu4lfzgc5o7****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The weight of the instance type specified in the extended configuration.
	//
	// example:
	//
	// 2
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
	// The policy for creating pay-as-you-go instances. Valid values:
	//
	// 	- lowest-price: cost optimization policy. This policy indicates that lowest-cost instance types are used to create instances.
	//
	// 	- prioritized: priority-based policy. This policy indicates that instances are created based on the priority specified by the LaunchTemplateConfig.N.Priority parameter.
	//
	// >  The LaunchTemplateConfig.N.Priority parameter is set when the auto provisioning group is created, and cannot be modified.
	//
	// example:
	//
	// prioritized
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
	// The policy for creating spot instances. Valid values:
	//
	// 	- lowest-price: cost optimization policy. This policy indicates that the lowest-priced instance type is used to create instances.
	//
	// 	- diversified: balanced distribution policy. This policy indicates that instances are created evenly across multiple zones specified in the extended configuration.
	//
	// example:
	//
	// diversified
	AllocationStrategy *string `json:"AllocationStrategy,omitempty" xml:"AllocationStrategy,omitempty"`
	// The action to be performed after the excess spot instances are stopped. Valid values:
	//
	// 	- stop: retains the excess spot instances in the stopped state.
	//
	// 	- terminate: releases the excess spot instances.
	//
	// example:
	//
	// stop
	InstanceInterruptionBehavior *string `json:"InstanceInterruptionBehavior,omitempty" xml:"InstanceInterruptionBehavior,omitempty"`
	// The number of instances that the auto provisioning group creates by selecting the instance type of the lowest price.
	//
	// >  This parameter is set when the auto provisioning group is created, and cannot be modified.
	//
	// example:
	//
	// 2
	InstancePoolsToUseCount *int32 `json:"InstancePoolsToUseCount,omitempty" xml:"InstancePoolsToUseCount,omitempty"`
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
	// The key of tag N that is added to the auto provisioning group.
	//
	// Valid values of N: 1 to 20. The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot contain http:// or https://. The tag key cannot start with acs: or aliyun.
	//
	// example:
	//
	// TestKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of tag N that is added to the auto provisioning group.
	//
	// Valid values of N: 1 to 20. The tag value can be an empty string. The tag value can be up to 128 characters in length and cannot contain http:// or https://.
	//
	// example:
	//
	// TestValue
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
	// The type of supplemental instances. When the sum of the `PayAsYouGoTargetCapacity` and `SpotTargetCapacity` values is less than the `TotalTargetCapacity` value, the auto provisioning group creates instances of the specified billing method to meet the target capacity. Valid values:
	//
	// 	- PayAsYouGo: pay-as-you-go instances.
	//
	// 	- Spot: spot instances.
	//
	// example:
	//
	// Spot
	DefaultTargetCapacityType *string `json:"DefaultTargetCapacityType,omitempty" xml:"DefaultTargetCapacityType,omitempty"`
	// The target capacity of pay-as-you-go instances that the auto provisioning group provisions.
	//
	// example:
	//
	// 30
	PayAsYouGoTargetCapacity *float32 `json:"PayAsYouGoTargetCapacity,omitempty" xml:"PayAsYouGoTargetCapacity,omitempty"`
	// The target capacity of spot instances that the auto provisioning group provisions.
	//
	// example:
	//
	// 20
	SpotTargetCapacity *float32 `json:"SpotTargetCapacity,omitempty" xml:"SpotTargetCapacity,omitempty"`
	// The target capacity of the auto provisioning group. The capacity consists of the following parts:
	//
	// 	- PayAsYouGoTargetCapacity
	//
	// 	- SpotTargetCapacity
	//
	// 	- The supplemental capacity besides instance capacities specified by PayAsYouGoTargetCapacity and SpotTargetCapacity.
	//
	// example:
	//
	// 60
	TotalTargetCapacity *float32 `json:"TotalTargetCapacity,omitempty" xml:"TotalTargetCapacity,omitempty"`
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
