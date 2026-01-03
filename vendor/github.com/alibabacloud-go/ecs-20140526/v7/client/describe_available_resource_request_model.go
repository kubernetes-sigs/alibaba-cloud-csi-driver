// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAvailableResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCores(v int32) *DescribeAvailableResourceRequest
	GetCores() *int32
	SetDataDiskCategory(v string) *DescribeAvailableResourceRequest
	GetDataDiskCategory() *string
	SetDedicatedHostId(v string) *DescribeAvailableResourceRequest
	GetDedicatedHostId() *string
	SetDestinationResource(v string) *DescribeAvailableResourceRequest
	GetDestinationResource() *string
	SetInstanceChargeType(v string) *DescribeAvailableResourceRequest
	GetInstanceChargeType() *string
	SetInstanceType(v string) *DescribeAvailableResourceRequest
	GetInstanceType() *string
	SetIoOptimized(v string) *DescribeAvailableResourceRequest
	GetIoOptimized() *string
	SetMemory(v float32) *DescribeAvailableResourceRequest
	GetMemory() *float32
	SetNetworkCategory(v string) *DescribeAvailableResourceRequest
	GetNetworkCategory() *string
	SetOwnerAccount(v string) *DescribeAvailableResourceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeAvailableResourceRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeAvailableResourceRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeAvailableResourceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeAvailableResourceRequest
	GetResourceOwnerId() *int64
	SetResourceType(v string) *DescribeAvailableResourceRequest
	GetResourceType() *string
	SetScope(v string) *DescribeAvailableResourceRequest
	GetScope() *string
	SetSpotDuration(v int32) *DescribeAvailableResourceRequest
	GetSpotDuration() *int32
	SetSpotStrategy(v string) *DescribeAvailableResourceRequest
	GetSpotStrategy() *string
	SetSystemDiskCategory(v string) *DescribeAvailableResourceRequest
	GetSystemDiskCategory() *string
	SetZoneId(v string) *DescribeAvailableResourceRequest
	GetZoneId() *string
}

type DescribeAvailableResourceRequest struct {
	// example:
	//
	// 2
	Cores *int32 `json:"Cores,omitempty" xml:"Cores,omitempty"`
	// example:
	//
	// cloud_ssd
	DataDiskCategory *string `json:"DataDiskCategory,omitempty" xml:"DataDiskCategory,omitempty"`
	// example:
	//
	// dh-bp165p6xk2tlw61e****
	DedicatedHostId *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// InstanceType
	DestinationResource *string `json:"DestinationResource,omitempty" xml:"DestinationResource,omitempty"`
	// example:
	//
	// PrePaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// example:
	//
	// ecs.g5.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// example:
	//
	// optimized
	IoOptimized *string `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	// example:
	//
	// 8.0
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// example:
	//
	// vpc
	NetworkCategory *string `json:"NetworkCategory,omitempty" xml:"NetworkCategory,omitempty"`
	OwnerAccount    *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// Region
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// example:
	//
	// 1
	SpotDuration *int32 `json:"SpotDuration,omitempty" xml:"SpotDuration,omitempty"`
	// example:
	//
	// NoSpot
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// example:
	//
	// cloud_ssd
	SystemDiskCategory *string `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty"`
	// example:
	//
	// cn-hangzhou-e
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeAvailableResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceRequest) GetCores() *int32 {
	return s.Cores
}

func (s *DescribeAvailableResourceRequest) GetDataDiskCategory() *string {
	return s.DataDiskCategory
}

func (s *DescribeAvailableResourceRequest) GetDedicatedHostId() *string {
	return s.DedicatedHostId
}

func (s *DescribeAvailableResourceRequest) GetDestinationResource() *string {
	return s.DestinationResource
}

func (s *DescribeAvailableResourceRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *DescribeAvailableResourceRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeAvailableResourceRequest) GetIoOptimized() *string {
	return s.IoOptimized
}

func (s *DescribeAvailableResourceRequest) GetMemory() *float32 {
	return s.Memory
}

func (s *DescribeAvailableResourceRequest) GetNetworkCategory() *string {
	return s.NetworkCategory
}

func (s *DescribeAvailableResourceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeAvailableResourceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeAvailableResourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAvailableResourceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeAvailableResourceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAvailableResourceRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeAvailableResourceRequest) GetScope() *string {
	return s.Scope
}

func (s *DescribeAvailableResourceRequest) GetSpotDuration() *int32 {
	return s.SpotDuration
}

func (s *DescribeAvailableResourceRequest) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *DescribeAvailableResourceRequest) GetSystemDiskCategory() *string {
	return s.SystemDiskCategory
}

func (s *DescribeAvailableResourceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeAvailableResourceRequest) SetCores(v int32) *DescribeAvailableResourceRequest {
	s.Cores = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetDataDiskCategory(v string) *DescribeAvailableResourceRequest {
	s.DataDiskCategory = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetDedicatedHostId(v string) *DescribeAvailableResourceRequest {
	s.DedicatedHostId = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetDestinationResource(v string) *DescribeAvailableResourceRequest {
	s.DestinationResource = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetInstanceChargeType(v string) *DescribeAvailableResourceRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetInstanceType(v string) *DescribeAvailableResourceRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetIoOptimized(v string) *DescribeAvailableResourceRequest {
	s.IoOptimized = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetMemory(v float32) *DescribeAvailableResourceRequest {
	s.Memory = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetNetworkCategory(v string) *DescribeAvailableResourceRequest {
	s.NetworkCategory = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetOwnerAccount(v string) *DescribeAvailableResourceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetOwnerId(v int64) *DescribeAvailableResourceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetRegionId(v string) *DescribeAvailableResourceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetResourceOwnerAccount(v string) *DescribeAvailableResourceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetResourceOwnerId(v int64) *DescribeAvailableResourceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetResourceType(v string) *DescribeAvailableResourceRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetScope(v string) *DescribeAvailableResourceRequest {
	s.Scope = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetSpotDuration(v int32) *DescribeAvailableResourceRequest {
	s.SpotDuration = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetSpotStrategy(v string) *DescribeAvailableResourceRequest {
	s.SpotStrategy = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetSystemDiskCategory(v string) *DescribeAvailableResourceRequest {
	s.SystemDiskCategory = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetZoneId(v string) *DescribeAvailableResourceRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeAvailableResourceRequest) Validate() error {
	return dara.Validate(s)
}
