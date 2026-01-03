// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDedicatedHostAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkAttributes(v *ModifyDedicatedHostAttributeRequestNetworkAttributes) *ModifyDedicatedHostAttributeRequest
	GetNetworkAttributes() *ModifyDedicatedHostAttributeRequestNetworkAttributes
	SetActionOnMaintenance(v string) *ModifyDedicatedHostAttributeRequest
	GetActionOnMaintenance() *string
	SetAutoPlacement(v string) *ModifyDedicatedHostAttributeRequest
	GetAutoPlacement() *string
	SetCpuOverCommitRatio(v float32) *ModifyDedicatedHostAttributeRequest
	GetCpuOverCommitRatio() *float32
	SetDedicatedHostClusterId(v string) *ModifyDedicatedHostAttributeRequest
	GetDedicatedHostClusterId() *string
	SetDedicatedHostId(v string) *ModifyDedicatedHostAttributeRequest
	GetDedicatedHostId() *string
	SetDedicatedHostName(v string) *ModifyDedicatedHostAttributeRequest
	GetDedicatedHostName() *string
	SetDescription(v string) *ModifyDedicatedHostAttributeRequest
	GetDescription() *string
	SetOwnerAccount(v string) *ModifyDedicatedHostAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDedicatedHostAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyDedicatedHostAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyDedicatedHostAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDedicatedHostAttributeRequest
	GetResourceOwnerId() *int64
}

type ModifyDedicatedHostAttributeRequest struct {
	NetworkAttributes *ModifyDedicatedHostAttributeRequestNetworkAttributes `json:"NetworkAttributes,omitempty" xml:"NetworkAttributes,omitempty" type:"Struct"`
	// The policy for migrating the instances deployed on the dedicated host when the dedicated host fails or needs to be repaired online. Valid values:
	//
	// 	- Migrate: The instances are migrated to another physical machine and then restarted.
	//
	// 	- Stop: The instances are stopped. If the dedicated host cannot be repaired, the instances are migrated to another physical machine and then restarted.
	//
	// If the dedicated host has cloud disks attached, the default value is Migrate.
	//
	// If the dedicated host has local disks attached, the default value is Stop.
	//
	// example:
	//
	// Migrate
	ActionOnMaintenance *string `json:"ActionOnMaintenance,omitempty" xml:"ActionOnMaintenance,omitempty"`
	// Specifies whether to add the dedicated host to the resource pool for automatic deployment. If you do not specify **DedicatedHostId*	- when you create an instance on a dedicated host, Alibaba Cloud automatically selects a dedicated host from the resource pool to host the instance. Valid values:
	//
	// 	- on: adds the dedicated host to the resource pool for automatic deployment.
	//
	// 	- off: does not add the dedicated host to the resource pool for automatic deployment.
	//
	// For information about automatic deployment, see [Functions and features](https://help.aliyun.com/document_detail/118938.html).
	//
	// example:
	//
	// on
	AutoPlacement *string `json:"AutoPlacement,omitempty" xml:"AutoPlacement,omitempty"`
	// The CPU overcommit ratio. You can configure CPU overcommit ratios only for the following dedicated host types: g6s, c6s, and r6s. Valid values: 1 to 5.
	//
	// The CPU overcommit ratio affects the number of available vCPUs on a dedicated host. You can use the following formula to calculate the number of available vCPUs on a dedicated host: Number of available vCPUs = Number of physical CPU cores × 2 × CPU overcommit ratio. For example, the number of physical CPU cores on each g6s dedicated host is 52. If you change the CPU overcommit ratio of a g6s dedicated host to 4, the number of available vCPUs on the dedicated host is 416. For scenarios that have minimal requirements for CPU stability or where CPU load is not heavy, such as development and test environments, you can increase the number of available vCPUs on a dedicated host by increasing the CPU overcommit ratio. This allows you to deploy more ECS instances of the same specifications on the dedicated host and reduce the unit deployment cost.
	//
	// example:
	//
	// 1
	CpuOverCommitRatio *float32 `json:"CpuOverCommitRatio,omitempty" xml:"CpuOverCommitRatio,omitempty"`
	// The ID of the dedicated host cluster to which to assign the dedicated host.
	//
	// example:
	//
	// dc-bp165p6xk2tlw61e****
	DedicatedHostClusterId *string `json:"DedicatedHostClusterId,omitempty" xml:"DedicatedHostClusterId,omitempty"`
	// The ID of the dedicated host.
	//
	// This parameter is required.
	//
	// example:
	//
	// dh-bp165p6xk2tlw61e****
	DedicatedHostId *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	// The name of the dedicated host. The name must be 2 to 128 characters in length. The name must start with a letter and cannot start with http:// or https://. The name can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// testDedicatedHostName
	DedicatedHostName *string `json:"DedicatedHostName,omitempty" xml:"DedicatedHostName,omitempty"`
	// The description of the dedicated host. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// testDescription
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the dedicated host resides. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDedicatedHostAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDedicatedHostAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDedicatedHostAttributeRequest) GetNetworkAttributes() *ModifyDedicatedHostAttributeRequestNetworkAttributes {
	return s.NetworkAttributes
}

func (s *ModifyDedicatedHostAttributeRequest) GetActionOnMaintenance() *string {
	return s.ActionOnMaintenance
}

func (s *ModifyDedicatedHostAttributeRequest) GetAutoPlacement() *string {
	return s.AutoPlacement
}

func (s *ModifyDedicatedHostAttributeRequest) GetCpuOverCommitRatio() *float32 {
	return s.CpuOverCommitRatio
}

func (s *ModifyDedicatedHostAttributeRequest) GetDedicatedHostClusterId() *string {
	return s.DedicatedHostClusterId
}

func (s *ModifyDedicatedHostAttributeRequest) GetDedicatedHostId() *string {
	return s.DedicatedHostId
}

func (s *ModifyDedicatedHostAttributeRequest) GetDedicatedHostName() *string {
	return s.DedicatedHostName
}

func (s *ModifyDedicatedHostAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyDedicatedHostAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDedicatedHostAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDedicatedHostAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDedicatedHostAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDedicatedHostAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDedicatedHostAttributeRequest) SetNetworkAttributes(v *ModifyDedicatedHostAttributeRequestNetworkAttributes) *ModifyDedicatedHostAttributeRequest {
	s.NetworkAttributes = v
	return s
}

func (s *ModifyDedicatedHostAttributeRequest) SetActionOnMaintenance(v string) *ModifyDedicatedHostAttributeRequest {
	s.ActionOnMaintenance = &v
	return s
}

func (s *ModifyDedicatedHostAttributeRequest) SetAutoPlacement(v string) *ModifyDedicatedHostAttributeRequest {
	s.AutoPlacement = &v
	return s
}

func (s *ModifyDedicatedHostAttributeRequest) SetCpuOverCommitRatio(v float32) *ModifyDedicatedHostAttributeRequest {
	s.CpuOverCommitRatio = &v
	return s
}

func (s *ModifyDedicatedHostAttributeRequest) SetDedicatedHostClusterId(v string) *ModifyDedicatedHostAttributeRequest {
	s.DedicatedHostClusterId = &v
	return s
}

func (s *ModifyDedicatedHostAttributeRequest) SetDedicatedHostId(v string) *ModifyDedicatedHostAttributeRequest {
	s.DedicatedHostId = &v
	return s
}

func (s *ModifyDedicatedHostAttributeRequest) SetDedicatedHostName(v string) *ModifyDedicatedHostAttributeRequest {
	s.DedicatedHostName = &v
	return s
}

func (s *ModifyDedicatedHostAttributeRequest) SetDescription(v string) *ModifyDedicatedHostAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyDedicatedHostAttributeRequest) SetOwnerAccount(v string) *ModifyDedicatedHostAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDedicatedHostAttributeRequest) SetOwnerId(v int64) *ModifyDedicatedHostAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDedicatedHostAttributeRequest) SetRegionId(v string) *ModifyDedicatedHostAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDedicatedHostAttributeRequest) SetResourceOwnerAccount(v string) *ModifyDedicatedHostAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDedicatedHostAttributeRequest) SetResourceOwnerId(v int64) *ModifyDedicatedHostAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDedicatedHostAttributeRequest) Validate() error {
	if s.NetworkAttributes != nil {
		if err := s.NetworkAttributes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyDedicatedHostAttributeRequestNetworkAttributes struct {
	// The timeout period for a UDP session between a Server Load Balancer (SLB) instance and the dedicated host. Unit: seconds. Valid values: 15 to 310.
	//
	// example:
	//
	// 60
	SlbUdpTimeout *int32 `json:"SlbUdpTimeout,omitempty" xml:"SlbUdpTimeout,omitempty"`
	// The timeout period for a UDP session between a user and an Alibaba Cloud service on the dedicated host. Unit: seconds. Valid values: 15 to 310.
	//
	// example:
	//
	// 60
	UdpTimeout *int32 `json:"UdpTimeout,omitempty" xml:"UdpTimeout,omitempty"`
}

func (s ModifyDedicatedHostAttributeRequestNetworkAttributes) String() string {
	return dara.Prettify(s)
}

func (s ModifyDedicatedHostAttributeRequestNetworkAttributes) GoString() string {
	return s.String()
}

func (s *ModifyDedicatedHostAttributeRequestNetworkAttributes) GetSlbUdpTimeout() *int32 {
	return s.SlbUdpTimeout
}

func (s *ModifyDedicatedHostAttributeRequestNetworkAttributes) GetUdpTimeout() *int32 {
	return s.UdpTimeout
}

func (s *ModifyDedicatedHostAttributeRequestNetworkAttributes) SetSlbUdpTimeout(v int32) *ModifyDedicatedHostAttributeRequestNetworkAttributes {
	s.SlbUdpTimeout = &v
	return s
}

func (s *ModifyDedicatedHostAttributeRequestNetworkAttributes) SetUdpTimeout(v int32) *ModifyDedicatedHostAttributeRequestNetworkAttributes {
	s.UdpTimeout = &v
	return s
}

func (s *ModifyDedicatedHostAttributeRequestNetworkAttributes) Validate() error {
	return dara.Validate(s)
}
