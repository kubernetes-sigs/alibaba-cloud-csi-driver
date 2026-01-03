// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceDeploymentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAffinity(v string) *ModifyInstanceDeploymentRequest
	GetAffinity() *string
	SetDedicatedHostClusterId(v string) *ModifyInstanceDeploymentRequest
	GetDedicatedHostClusterId() *string
	SetDedicatedHostId(v string) *ModifyInstanceDeploymentRequest
	GetDedicatedHostId() *string
	SetDeploymentSetGroupNo(v int32) *ModifyInstanceDeploymentRequest
	GetDeploymentSetGroupNo() *int32
	SetDeploymentSetId(v string) *ModifyInstanceDeploymentRequest
	GetDeploymentSetId() *string
	SetForce(v bool) *ModifyInstanceDeploymentRequest
	GetForce() *bool
	SetInstanceId(v string) *ModifyInstanceDeploymentRequest
	GetInstanceId() *string
	SetInstanceType(v string) *ModifyInstanceDeploymentRequest
	GetInstanceType() *string
	SetMigrationType(v string) *ModifyInstanceDeploymentRequest
	GetMigrationType() *string
	SetOwnerAccount(v string) *ModifyInstanceDeploymentRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyInstanceDeploymentRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyInstanceDeploymentRequest
	GetRegionId() *string
	SetRemoveFromDeploymentSet(v bool) *ModifyInstanceDeploymentRequest
	GetRemoveFromDeploymentSet() *bool
	SetResourceOwnerAccount(v string) *ModifyInstanceDeploymentRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyInstanceDeploymentRequest
	GetResourceOwnerId() *int64
	SetTenancy(v string) *ModifyInstanceDeploymentRequest
	GetTenancy() *string
}

type ModifyInstanceDeploymentRequest struct {
	// Specifies whether to associate the instance with a dedicated host. Valid values:
	//
	// 	- host: associates the instance with a dedicated host. When you start a stopped instance in economical mode, the instance remains on its original dedicated host.
	//
	// 	- default: does not associate the instance with a dedicated host. When you start a stopped instance in economical mode, the instance can be automatically deployed to another dedicated host in the automatic deployment resource pool if the resources of the original dedicated host are insufficient.
	//
	// If you want to migrate the instance from a shared host to a dedicated host, use the default value. Default value: default.
	//
	// example:
	//
	// host
	Affinity *string `json:"Affinity,omitempty" xml:"Affinity,omitempty"`
	// The ID of the dedicated host cluster.
	//
	// example:
	//
	// dc-bp67acfmxazb4ph****
	DedicatedHostClusterId *string `json:"DedicatedHostClusterId,omitempty" xml:"DedicatedHostClusterId,omitempty"`
	// The ID of the destination dedicated host. You can call the [DescribeDedicatedHosts](https://help.aliyun.com/document_detail/134242.html) operation to query the most recent list of dedicated hosts.
	//
	// When you migrate an instance from a shared host to a dedicated host or between dedicated hosts, take note of the following items:
	//
	// 	- To migrate the instance to a specific dedicated host, specify this parameter.
	//
	// 	- To migrate the instance to a system-selected dedicated host, leave this parameter empty and set `Tenancy` to host.
	//
	// For information about the automatic deployment feature, see [Functions and features](https://help.aliyun.com/document_detail/118938.html).
	//
	// example:
	//
	// dh-bp67acfmxazb4ph****
	DedicatedHostId *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	// The number of the deployment set group in which to deploy the instance in the destination deployment set. This parameter is valid only when the destination deployment set uses the high availability group strategy (AvailabilityGroup). Valid values: 1 to 7.
	//
	// > If you call this operation to deploy an instance to a deployment set that uses the high availability group strategy (`AvailablilityGroup`) and leave this parameter empty, the system evenly distributes instances among the deployment set groups in the deployment set. If you call this operation to change the deployment set of an instance and specify the current deployment set of the instance as the destination deployment set, the system evenly distributes instances again among the deployment set groups in the deployment set.
	//
	// example:
	//
	// 3
	DeploymentSetGroupNo *int32 `json:"DeploymentSetGroupNo,omitempty" xml:"DeploymentSetGroupNo,omitempty"`
	// The ID of the destination deployment set.
	//
	// This parameter is required when you add an instance to a deployment set or change the deployment set of an instance.
	//
	// > You cannot change the deployment set when you modify dedicated host configurations, including the `Tenancy`, `Affinity`, and `DedicatedHostId` parameters.
	//
	// example:
	//
	// ds-bp67acfmxazb4ph****
	DeploymentSetId *string `json:"DeploymentSetId,omitempty" xml:"DeploymentSetId,omitempty"`
	// Specifies whether to forcefully change the host of the instance when the deployment set of the instance is changed. Valid values:
	//
	// 	- true: forcefully changes the host of the instance when the deployment set of the instance is changed. Hosts can be forcefully changed only for instances in the Running (Running) or Stopped (Stopped) state. The instances that are in the Stopped (Stopped) state do not include pay-as-you-go instances that are stopped in economical mode.
	//
	//     **
	//
	//     **Note*	- If the specified instance has local disks attached, the local disks are forcefully changed when the host of the instance is forcefully changed. This may cause data loss in the local disks. Proceed with caution.
	//
	// 	- false: does not forcefully change the host of the instance when the deployment set of the instance is changed. You can add the instance to a deployment set only when the instance remains on the current host. When the Force parameter is set to false, the deployment set may fail to be changed.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp67acfmxazb4ph***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance type to which the instance is changed. You can call the [DescribeInstanceTypes](https://help.aliyun.com/document_detail/25620.html) operation to query the most recent list of instance types.
	//
	// You can change the instance type of an instance when you migrate the instance to a dedicated host. The new instance type must match the type of the specified dedicated host. For more information, see [Dedicated host types](https://help.aliyun.com/document_detail/68564.html).
	//
	// 	- If you specify this parameter, you must also specify `DedicatedHostId`.
	//
	// 	- You cannot change the instance type of an instance if you use the automatic deployment feature to migrate the instance.
	//
	// example:
	//
	// ecs.c6.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// Specifies whether to stop the instance before it is migrated to the destination dedicated host. Valid values:
	//
	// 	- reboot: stops the instance before it is migrated.
	//
	// 	- live: migrates the instance without stopping it. If you set MigrationType to live, you must specify DedicatedHostId. In this case, you cannot change the instance type of the instance when the instance is migrated.
	//
	// Default value: reboot.
	//
	// example:
	//
	// live
	MigrationType *string `json:"MigrationType,omitempty" xml:"MigrationType,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to remove the specified instance from the specified deployment set. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// > If you set this parameter to true, you must specify InstanceId and DeploymentSetId and make sure that the specified instance belongs to the specified deployment set.
	//
	// example:
	//
	// false
	RemoveFromDeploymentSet *bool   `json:"RemoveFromDeploymentSet,omitempty" xml:"RemoveFromDeploymentSet,omitempty"`
	ResourceOwnerAccount    *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId         *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to deploy the instance on a dedicated host. Set the value to host, which indicates that the instance is deployed on a dedicated host.
	//
	// example:
	//
	// host
	Tenancy *string `json:"Tenancy,omitempty" xml:"Tenancy,omitempty"`
}

func (s ModifyInstanceDeploymentRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceDeploymentRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceDeploymentRequest) GetAffinity() *string {
	return s.Affinity
}

func (s *ModifyInstanceDeploymentRequest) GetDedicatedHostClusterId() *string {
	return s.DedicatedHostClusterId
}

func (s *ModifyInstanceDeploymentRequest) GetDedicatedHostId() *string {
	return s.DedicatedHostId
}

func (s *ModifyInstanceDeploymentRequest) GetDeploymentSetGroupNo() *int32 {
	return s.DeploymentSetGroupNo
}

func (s *ModifyInstanceDeploymentRequest) GetDeploymentSetId() *string {
	return s.DeploymentSetId
}

func (s *ModifyInstanceDeploymentRequest) GetForce() *bool {
	return s.Force
}

func (s *ModifyInstanceDeploymentRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceDeploymentRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ModifyInstanceDeploymentRequest) GetMigrationType() *string {
	return s.MigrationType
}

func (s *ModifyInstanceDeploymentRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyInstanceDeploymentRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyInstanceDeploymentRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyInstanceDeploymentRequest) GetRemoveFromDeploymentSet() *bool {
	return s.RemoveFromDeploymentSet
}

func (s *ModifyInstanceDeploymentRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyInstanceDeploymentRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyInstanceDeploymentRequest) GetTenancy() *string {
	return s.Tenancy
}

func (s *ModifyInstanceDeploymentRequest) SetAffinity(v string) *ModifyInstanceDeploymentRequest {
	s.Affinity = &v
	return s
}

func (s *ModifyInstanceDeploymentRequest) SetDedicatedHostClusterId(v string) *ModifyInstanceDeploymentRequest {
	s.DedicatedHostClusterId = &v
	return s
}

func (s *ModifyInstanceDeploymentRequest) SetDedicatedHostId(v string) *ModifyInstanceDeploymentRequest {
	s.DedicatedHostId = &v
	return s
}

func (s *ModifyInstanceDeploymentRequest) SetDeploymentSetGroupNo(v int32) *ModifyInstanceDeploymentRequest {
	s.DeploymentSetGroupNo = &v
	return s
}

func (s *ModifyInstanceDeploymentRequest) SetDeploymentSetId(v string) *ModifyInstanceDeploymentRequest {
	s.DeploymentSetId = &v
	return s
}

func (s *ModifyInstanceDeploymentRequest) SetForce(v bool) *ModifyInstanceDeploymentRequest {
	s.Force = &v
	return s
}

func (s *ModifyInstanceDeploymentRequest) SetInstanceId(v string) *ModifyInstanceDeploymentRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceDeploymentRequest) SetInstanceType(v string) *ModifyInstanceDeploymentRequest {
	s.InstanceType = &v
	return s
}

func (s *ModifyInstanceDeploymentRequest) SetMigrationType(v string) *ModifyInstanceDeploymentRequest {
	s.MigrationType = &v
	return s
}

func (s *ModifyInstanceDeploymentRequest) SetOwnerAccount(v string) *ModifyInstanceDeploymentRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyInstanceDeploymentRequest) SetOwnerId(v int64) *ModifyInstanceDeploymentRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyInstanceDeploymentRequest) SetRegionId(v string) *ModifyInstanceDeploymentRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyInstanceDeploymentRequest) SetRemoveFromDeploymentSet(v bool) *ModifyInstanceDeploymentRequest {
	s.RemoveFromDeploymentSet = &v
	return s
}

func (s *ModifyInstanceDeploymentRequest) SetResourceOwnerAccount(v string) *ModifyInstanceDeploymentRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyInstanceDeploymentRequest) SetResourceOwnerId(v int64) *ModifyInstanceDeploymentRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyInstanceDeploymentRequest) SetTenancy(v string) *ModifyInstanceDeploymentRequest {
	s.Tenancy = &v
	return s
}

func (s *ModifyInstanceDeploymentRequest) Validate() error {
	return dara.Validate(s)
}
