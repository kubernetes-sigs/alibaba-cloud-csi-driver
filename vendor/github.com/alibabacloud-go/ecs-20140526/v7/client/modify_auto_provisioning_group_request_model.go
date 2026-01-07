// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAutoProvisioningGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoProvisioningGroupId(v string) *ModifyAutoProvisioningGroupRequest
	GetAutoProvisioningGroupId() *string
	SetAutoProvisioningGroupName(v string) *ModifyAutoProvisioningGroupRequest
	GetAutoProvisioningGroupName() *string
	SetDefaultTargetCapacityType(v string) *ModifyAutoProvisioningGroupRequest
	GetDefaultTargetCapacityType() *string
	SetExcessCapacityTerminationPolicy(v string) *ModifyAutoProvisioningGroupRequest
	GetExcessCapacityTerminationPolicy() *string
	SetLaunchTemplateConfig(v []*ModifyAutoProvisioningGroupRequestLaunchTemplateConfig) *ModifyAutoProvisioningGroupRequest
	GetLaunchTemplateConfig() []*ModifyAutoProvisioningGroupRequestLaunchTemplateConfig
	SetMaxSpotPrice(v float32) *ModifyAutoProvisioningGroupRequest
	GetMaxSpotPrice() *float32
	SetOwnerAccount(v string) *ModifyAutoProvisioningGroupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyAutoProvisioningGroupRequest
	GetOwnerId() *int64
	SetPayAsYouGoTargetCapacity(v string) *ModifyAutoProvisioningGroupRequest
	GetPayAsYouGoTargetCapacity() *string
	SetRegionId(v string) *ModifyAutoProvisioningGroupRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyAutoProvisioningGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyAutoProvisioningGroupRequest
	GetResourceOwnerId() *int64
	SetSpotTargetCapacity(v string) *ModifyAutoProvisioningGroupRequest
	GetSpotTargetCapacity() *string
	SetTerminateInstancesWithExpiration(v bool) *ModifyAutoProvisioningGroupRequest
	GetTerminateInstancesWithExpiration() *bool
	SetTotalTargetCapacity(v string) *ModifyAutoProvisioningGroupRequest
	GetTotalTargetCapacity() *string
}

type ModifyAutoProvisioningGroupRequest struct {
	// The auto-provisioning group ID.
	//
	// example:
	//
	// apg-bp67acfmxazb4ph****
	AutoProvisioningGroupId *string `json:"AutoProvisioningGroupId,omitempty" xml:"AutoProvisioningGroupId,omitempty"`
	// The name of the auto-provisioning group. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with http:// or https://.[ It can contain letters, digits, colons (:), underscores (_), and hyphens (-).](http://https://。、（:）、（_）（-）。)
	//
	// example:
	//
	// apg-test
	AutoProvisioningGroupName *string `json:"AutoProvisioningGroupName,omitempty" xml:"AutoProvisioningGroupName,omitempty"`
	// The type of supplemental instances. When the sum of the PayAsYouGoTargetCapacity and SpotTargetCapacity values is smaller than the TotalTargetCapacity value, the auto-provisioning group creates instances of the specified type to meet the target capacity. Valid values:
	//
	// 	- PayAsYouGo: pay-as-you-go instances
	//
	// 	- Spot: spot instances
	//
	// example:
	//
	// Spot
	DefaultTargetCapacityType *string `json:"DefaultTargetCapacityType,omitempty" xml:"DefaultTargetCapacityType,omitempty"`
	// Specifies whether to release the removed instances when the real-time capacity of the auto-provisioning group exceeds the target capacity and a scale-in event is triggered. Valid values:
	//
	// 	- termination: releases the removed instances.
	//
	// 	- no-termination: removes the instances from the auto-provisioning group but does not release them.
	//
	// example:
	//
	// no-termination
	ExcessCapacityTerminationPolicy *string `json:"ExcessCapacityTerminationPolicy,omitempty" xml:"ExcessCapacityTerminationPolicy,omitempty"`
	// The extended configurations of the launch template.
	LaunchTemplateConfig []*ModifyAutoProvisioningGroupRequestLaunchTemplateConfig `json:"LaunchTemplateConfig,omitempty" xml:"LaunchTemplateConfig,omitempty" type:"Repeated"`
	// The maximum price of spot instances in the auto-provisioning group.
	//
	// > When both the MaxSpotPrice and LaunchTemplateConfig.N.MaxPrice parameters are specified, the smaller one of the two parameter values is used. The LaunchTemplateConfig.N.MaxPrice parameter is specified when the auto-provisioning group is created, and cannot be modified.
	//
	// example:
	//
	// 0.5
	MaxSpotPrice *float32 `json:"MaxSpotPrice,omitempty" xml:"MaxSpotPrice,omitempty"`
	OwnerAccount *string  `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64   `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The target capacity of pay-as-you-go instances in the auto-provisioning group. Valid values: Set this parameter to a value smaller than the TotalTargetCapacity value.
	//
	// example:
	//
	// 30
	PayAsYouGoTargetCapacity *string `json:"PayAsYouGoTargetCapacity,omitempty" xml:"PayAsYouGoTargetCapacity,omitempty"`
	// The region ID of the auto-provisioning group. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The target capacity of spot instances in the auto-provisioning group. Valid values: Set this parameter to a value smaller than the TotalTargetCapacity value.
	//
	// example:
	//
	// 30
	SpotTargetCapacity *string `json:"SpotTargetCapacity,omitempty" xml:"SpotTargetCapacity,omitempty"`
	// Specifies whether to release instances that are located in the auto-provisioning group after the group expires. Valid values:
	//
	// 	- true: releases instances that are located in the auto-provisioning group.
	//
	// 	- false: removes instances from the auto-provisioning group but does not release them.
	//
	// example:
	//
	// false
	TerminateInstancesWithExpiration *bool `json:"TerminateInstancesWithExpiration,omitempty" xml:"TerminateInstancesWithExpiration,omitempty"`
	// The total target capacity of the auto-provisioning group. The value must be a positive integer.
	//
	// The total target capacity of the auto-provisioning group must be greater than or equal to the sum of the target capacity of pay-as-you-go instances specified by the PayAsYouGoTargetCapacity parameter as well as the target capacity of spot instances specified by the SpotTargetCapacity parameter.
	//
	// example:
	//
	// 70
	TotalTargetCapacity *string `json:"TotalTargetCapacity,omitempty" xml:"TotalTargetCapacity,omitempty"`
}

func (s ModifyAutoProvisioningGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAutoProvisioningGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyAutoProvisioningGroupRequest) GetAutoProvisioningGroupId() *string {
	return s.AutoProvisioningGroupId
}

func (s *ModifyAutoProvisioningGroupRequest) GetAutoProvisioningGroupName() *string {
	return s.AutoProvisioningGroupName
}

func (s *ModifyAutoProvisioningGroupRequest) GetDefaultTargetCapacityType() *string {
	return s.DefaultTargetCapacityType
}

func (s *ModifyAutoProvisioningGroupRequest) GetExcessCapacityTerminationPolicy() *string {
	return s.ExcessCapacityTerminationPolicy
}

func (s *ModifyAutoProvisioningGroupRequest) GetLaunchTemplateConfig() []*ModifyAutoProvisioningGroupRequestLaunchTemplateConfig {
	return s.LaunchTemplateConfig
}

func (s *ModifyAutoProvisioningGroupRequest) GetMaxSpotPrice() *float32 {
	return s.MaxSpotPrice
}

func (s *ModifyAutoProvisioningGroupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyAutoProvisioningGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyAutoProvisioningGroupRequest) GetPayAsYouGoTargetCapacity() *string {
	return s.PayAsYouGoTargetCapacity
}

func (s *ModifyAutoProvisioningGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyAutoProvisioningGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyAutoProvisioningGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyAutoProvisioningGroupRequest) GetSpotTargetCapacity() *string {
	return s.SpotTargetCapacity
}

func (s *ModifyAutoProvisioningGroupRequest) GetTerminateInstancesWithExpiration() *bool {
	return s.TerminateInstancesWithExpiration
}

func (s *ModifyAutoProvisioningGroupRequest) GetTotalTargetCapacity() *string {
	return s.TotalTargetCapacity
}

func (s *ModifyAutoProvisioningGroupRequest) SetAutoProvisioningGroupId(v string) *ModifyAutoProvisioningGroupRequest {
	s.AutoProvisioningGroupId = &v
	return s
}

func (s *ModifyAutoProvisioningGroupRequest) SetAutoProvisioningGroupName(v string) *ModifyAutoProvisioningGroupRequest {
	s.AutoProvisioningGroupName = &v
	return s
}

func (s *ModifyAutoProvisioningGroupRequest) SetDefaultTargetCapacityType(v string) *ModifyAutoProvisioningGroupRequest {
	s.DefaultTargetCapacityType = &v
	return s
}

func (s *ModifyAutoProvisioningGroupRequest) SetExcessCapacityTerminationPolicy(v string) *ModifyAutoProvisioningGroupRequest {
	s.ExcessCapacityTerminationPolicy = &v
	return s
}

func (s *ModifyAutoProvisioningGroupRequest) SetLaunchTemplateConfig(v []*ModifyAutoProvisioningGroupRequestLaunchTemplateConfig) *ModifyAutoProvisioningGroupRequest {
	s.LaunchTemplateConfig = v
	return s
}

func (s *ModifyAutoProvisioningGroupRequest) SetMaxSpotPrice(v float32) *ModifyAutoProvisioningGroupRequest {
	s.MaxSpotPrice = &v
	return s
}

func (s *ModifyAutoProvisioningGroupRequest) SetOwnerAccount(v string) *ModifyAutoProvisioningGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyAutoProvisioningGroupRequest) SetOwnerId(v int64) *ModifyAutoProvisioningGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyAutoProvisioningGroupRequest) SetPayAsYouGoTargetCapacity(v string) *ModifyAutoProvisioningGroupRequest {
	s.PayAsYouGoTargetCapacity = &v
	return s
}

func (s *ModifyAutoProvisioningGroupRequest) SetRegionId(v string) *ModifyAutoProvisioningGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyAutoProvisioningGroupRequest) SetResourceOwnerAccount(v string) *ModifyAutoProvisioningGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyAutoProvisioningGroupRequest) SetResourceOwnerId(v int64) *ModifyAutoProvisioningGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyAutoProvisioningGroupRequest) SetSpotTargetCapacity(v string) *ModifyAutoProvisioningGroupRequest {
	s.SpotTargetCapacity = &v
	return s
}

func (s *ModifyAutoProvisioningGroupRequest) SetTerminateInstancesWithExpiration(v bool) *ModifyAutoProvisioningGroupRequest {
	s.TerminateInstancesWithExpiration = &v
	return s
}

func (s *ModifyAutoProvisioningGroupRequest) SetTotalTargetCapacity(v string) *ModifyAutoProvisioningGroupRequest {
	s.TotalTargetCapacity = &v
	return s
}

func (s *ModifyAutoProvisioningGroupRequest) Validate() error {
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

type ModifyAutoProvisioningGroupRequestLaunchTemplateConfig struct {
	// The instance type in extended configuration N. Valid values of N: 1 to 20. For more information about the valid values of this parameter, see [Instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// example:
	//
	// ecs.g5.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The maximum price of spot instances in extended configuration N.
	//
	// example:
	//
	// 3
	MaxPrice *float64 `json:"MaxPrice,omitempty" xml:"MaxPrice,omitempty"`
	// The priority of extended configuration N. A value of 0 indicates the highest priority. The value must be greater than 0.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ID of the vSwitch in extended configuration N. The zone of the instances created from the extended configuration is determined by the vSwitch.
	//
	// example:
	//
	// vsw-sn5bsitu4lfzgc5o7****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The weight of the instance type specified in the extended configuration. A greater weight indicates that a single instance has more computing power and fewer instances are required. The value must be greater than 0.
	//
	// The weight is calculated based on the computing power of the instance type and the minimum computing power of a single instance in the cluster that can created by the auto-provisioning group. For example, assume that the minimum computing power of a single instance is 8 vCPUs and 60 GiB of memory.
	//
	// 	- For an instance type with 8 vCPUs and 60 GiB of memory, you can set the weight to 1.
	//
	// 	- For an instance type with 16 vCPUs and 120 GiB of memory, you can set the weight to 2.
	//
	// example:
	//
	// 2
	WeightedCapacity *float64 `json:"WeightedCapacity,omitempty" xml:"WeightedCapacity,omitempty"`
}

func (s ModifyAutoProvisioningGroupRequestLaunchTemplateConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyAutoProvisioningGroupRequestLaunchTemplateConfig) GoString() string {
	return s.String()
}

func (s *ModifyAutoProvisioningGroupRequestLaunchTemplateConfig) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ModifyAutoProvisioningGroupRequestLaunchTemplateConfig) GetMaxPrice() *float64 {
	return s.MaxPrice
}

func (s *ModifyAutoProvisioningGroupRequestLaunchTemplateConfig) GetPriority() *int32 {
	return s.Priority
}

func (s *ModifyAutoProvisioningGroupRequestLaunchTemplateConfig) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ModifyAutoProvisioningGroupRequestLaunchTemplateConfig) GetWeightedCapacity() *float64 {
	return s.WeightedCapacity
}

func (s *ModifyAutoProvisioningGroupRequestLaunchTemplateConfig) SetInstanceType(v string) *ModifyAutoProvisioningGroupRequestLaunchTemplateConfig {
	s.InstanceType = &v
	return s
}

func (s *ModifyAutoProvisioningGroupRequestLaunchTemplateConfig) SetMaxPrice(v float64) *ModifyAutoProvisioningGroupRequestLaunchTemplateConfig {
	s.MaxPrice = &v
	return s
}

func (s *ModifyAutoProvisioningGroupRequestLaunchTemplateConfig) SetPriority(v int32) *ModifyAutoProvisioningGroupRequestLaunchTemplateConfig {
	s.Priority = &v
	return s
}

func (s *ModifyAutoProvisioningGroupRequestLaunchTemplateConfig) SetVSwitchId(v string) *ModifyAutoProvisioningGroupRequestLaunchTemplateConfig {
	s.VSwitchId = &v
	return s
}

func (s *ModifyAutoProvisioningGroupRequestLaunchTemplateConfig) SetWeightedCapacity(v float64) *ModifyAutoProvisioningGroupRequestLaunchTemplateConfig {
	s.WeightedCapacity = &v
	return s
}

func (s *ModifyAutoProvisioningGroupRequestLaunchTemplateConfig) Validate() error {
	return dara.Validate(s)
}
