// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourcesModificationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConditions(v []*string) *DescribeResourcesModificationRequest
	GetConditions() []*string
	SetCores(v int32) *DescribeResourcesModificationRequest
	GetCores() *int32
	SetDestinationResource(v string) *DescribeResourcesModificationRequest
	GetDestinationResource() *string
	SetInstanceType(v string) *DescribeResourcesModificationRequest
	GetInstanceType() *string
	SetMemory(v float32) *DescribeResourcesModificationRequest
	GetMemory() *float32
	SetMigrateAcrossZone(v bool) *DescribeResourcesModificationRequest
	GetMigrateAcrossZone() *bool
	SetOperationType(v string) *DescribeResourcesModificationRequest
	GetOperationType() *string
	SetOwnerAccount(v string) *DescribeResourcesModificationRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeResourcesModificationRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeResourcesModificationRequest
	GetRegionId() *string
	SetResourceId(v string) *DescribeResourcesModificationRequest
	GetResourceId() *string
	SetResourceOwnerAccount(v string) *DescribeResourcesModificationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeResourcesModificationRequest
	GetResourceOwnerId() *int64
	SetZoneId(v string) *DescribeResourcesModificationRequest
	GetZoneId() *string
}

type DescribeResourcesModificationRequest struct {
	// The conditions.
	Conditions []*string `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	// The number of vCPUs of the instance type. For information about the valid values, see [Overview of instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// This parameter is valid only when the DestinationResource parameter is set to InstanceType.
	//
	// example:
	//
	// 2
	Cores *int32 `json:"Cores,omitempty" xml:"Cores,omitempty"`
	// The resource type that you want to change. Valid values:
	//
	// 	- InstanceType
	//
	// 	- SystemDisk
	//
	//     If you set this parameter to SystemDisk, you must specify the InstanceType parameter. In this case, this operation queries the system disk categories supported by the specified instance type.
	//
	// This parameter is required.
	//
	// example:
	//
	// InstanceType
	DestinationResource *string `json:"DestinationResource,omitempty" xml:"DestinationResource,omitempty"`
	// The instance type to which you want to change the instance type of the instance. For more information, see [Overview of instance families](https://help.aliyun.com/document_detail/25378.html). You can also call the [DescribeInstanceTypes](https://help.aliyun.com/document_detail/25620.html) operation to query the most recent instance type list.
	//
	// If you set the DestinationResource parameter to SystemDisk, you must specify the InstanceType parameter. In this case, this operation queries the system disk categories supported by the specified instance type.
	//
	// example:
	//
	// ecs.g5.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The memory size of the instance type. Unit: GiB. For information about the valid values, see [Overview of instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// This parameter is valid only when the DestinationResource parameter is set to InstanceType.
	//
	// example:
	//
	// 8.0
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// Specifies whether cross-cluster instance type upgrades are supported. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// When MigrateAcrossZone is set to true and you upgrade the instance type of an instance based on the returned information, take note of the following items:
	//
	// 	- Instance that resides in the classic network:
	//
	//     	- For [retired instance types](https://help.aliyun.com/document_detail/55263.html), when a non-I/O optimized instance is upgraded to an I/O optimized instance, the private IP address, disk device names, and software authorization codes of the instance change. For a Linux instance, basic disks (cloud) are identified as xvd\\	- such as xvda and xvdb, and ultra disks (cloud_efficiency) and standard SSDs (cloud_ssd) are identified as vd\\	- such as vda and vdb.
	//
	//     	- For [instance families available for purchase](https://help.aliyun.com/document_detail/25378.html), when the instance type of an instance is changed, the private IP address of the instance changes.
	//
	// 	- Instance that resides in a virtual private cloud (VPC): For [retired instance types](https://help.aliyun.com/document_detail/55263.html), when a non-I/O optimized instance is upgraded to an I/O optimized instance, the disk device names and software authorization codes of the instance change. For a Linux instance, basic disks (cloud) are identified as xvd\\	- such as xvda and xvdb, and ultra disks (cloud_efficiency) and standard SSDs (cloud_ssd) are identified as vd\\	- such as vda and vdb.
	//
	// example:
	//
	// true
	MigrateAcrossZone *bool `json:"MigrateAcrossZone,omitempty" xml:"MigrateAcrossZone,omitempty"`
	// The operation of changing resource configurations.
	//
	// 	- Valid values for subscription resources:
	//
	//     	- Upgrade: upgrades resources.
	//
	//     	- Downgrade: downgrades resources.
	//
	//     	- RenewDowngrade: renews and downgrades resources.
	//
	//     	- RenewModify: renews an expired instance and changes its configurations.
	//
	// 	- Set the value to Upgrade for pay-as-you-go resources.
	//
	// Default value: Upgrade.
	//
	// example:
	//
	// Upgrade
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance for which you want to change the instance type or system disk category. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the instance for which you want to change the instance type or system disk category.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp67acfmxazb4p****
	ResourceId           *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the destination zone to which you want to migrate the instance.
	//
	// If you want to change the instance type across zones, you must specify this parameter.
	//
	// example:
	//
	// cn-hangzhou-e
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeResourcesModificationRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourcesModificationRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourcesModificationRequest) GetConditions() []*string {
	return s.Conditions
}

func (s *DescribeResourcesModificationRequest) GetCores() *int32 {
	return s.Cores
}

func (s *DescribeResourcesModificationRequest) GetDestinationResource() *string {
	return s.DestinationResource
}

func (s *DescribeResourcesModificationRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeResourcesModificationRequest) GetMemory() *float32 {
	return s.Memory
}

func (s *DescribeResourcesModificationRequest) GetMigrateAcrossZone() *bool {
	return s.MigrateAcrossZone
}

func (s *DescribeResourcesModificationRequest) GetOperationType() *string {
	return s.OperationType
}

func (s *DescribeResourcesModificationRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeResourcesModificationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeResourcesModificationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeResourcesModificationRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeResourcesModificationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeResourcesModificationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeResourcesModificationRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeResourcesModificationRequest) SetConditions(v []*string) *DescribeResourcesModificationRequest {
	s.Conditions = v
	return s
}

func (s *DescribeResourcesModificationRequest) SetCores(v int32) *DescribeResourcesModificationRequest {
	s.Cores = &v
	return s
}

func (s *DescribeResourcesModificationRequest) SetDestinationResource(v string) *DescribeResourcesModificationRequest {
	s.DestinationResource = &v
	return s
}

func (s *DescribeResourcesModificationRequest) SetInstanceType(v string) *DescribeResourcesModificationRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeResourcesModificationRequest) SetMemory(v float32) *DescribeResourcesModificationRequest {
	s.Memory = &v
	return s
}

func (s *DescribeResourcesModificationRequest) SetMigrateAcrossZone(v bool) *DescribeResourcesModificationRequest {
	s.MigrateAcrossZone = &v
	return s
}

func (s *DescribeResourcesModificationRequest) SetOperationType(v string) *DescribeResourcesModificationRequest {
	s.OperationType = &v
	return s
}

func (s *DescribeResourcesModificationRequest) SetOwnerAccount(v string) *DescribeResourcesModificationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeResourcesModificationRequest) SetOwnerId(v int64) *DescribeResourcesModificationRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeResourcesModificationRequest) SetRegionId(v string) *DescribeResourcesModificationRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeResourcesModificationRequest) SetResourceId(v string) *DescribeResourcesModificationRequest {
	s.ResourceId = &v
	return s
}

func (s *DescribeResourcesModificationRequest) SetResourceOwnerAccount(v string) *DescribeResourcesModificationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeResourcesModificationRequest) SetResourceOwnerId(v int64) *DescribeResourcesModificationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeResourcesModificationRequest) SetZoneId(v string) *DescribeResourcesModificationRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeResourcesModificationRequest) Validate() error {
	return dara.Validate(s)
}
