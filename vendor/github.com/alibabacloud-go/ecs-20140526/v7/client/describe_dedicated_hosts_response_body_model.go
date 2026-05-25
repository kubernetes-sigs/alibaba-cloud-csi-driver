// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDedicatedHostsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDedicatedHosts(v *DescribeDedicatedHostsResponseBodyDedicatedHosts) *DescribeDedicatedHostsResponseBody
	GetDedicatedHosts() *DescribeDedicatedHostsResponseBodyDedicatedHosts
	SetNextToken(v string) *DescribeDedicatedHostsResponseBody
	GetNextToken() *string
	SetPageNumber(v int32) *DescribeDedicatedHostsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDedicatedHostsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeDedicatedHostsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeDedicatedHostsResponseBody
	GetTotalCount() *int32
}

type DescribeDedicatedHostsResponseBody struct {
	DedicatedHosts *DescribeDedicatedHostsResponseBodyDedicatedHosts `json:"DedicatedHosts,omitempty" xml:"DedicatedHosts,omitempty" type:"Struct"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists. If the return value of this parameter is empty when you specify MaxResults and NextToken for a paged query, no more results are to be returned.
	//
	// example:
	//
	// e71d8a535bd9cc11
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number.
	//
	// example:
	//
	// 5
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 7654525A-9964-4ABB-8BCD-98F8835E809A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of dedicated hosts.
	//
	// example:
	//
	// 3
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDedicatedHostsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostsResponseBody) GetDedicatedHosts() *DescribeDedicatedHostsResponseBodyDedicatedHosts {
	return s.DedicatedHosts
}

func (s *DescribeDedicatedHostsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDedicatedHostsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDedicatedHostsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDedicatedHostsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDedicatedHostsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDedicatedHostsResponseBody) SetDedicatedHosts(v *DescribeDedicatedHostsResponseBodyDedicatedHosts) *DescribeDedicatedHostsResponseBody {
	s.DedicatedHosts = v
	return s
}

func (s *DescribeDedicatedHostsResponseBody) SetNextToken(v string) *DescribeDedicatedHostsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBody) SetPageNumber(v int32) *DescribeDedicatedHostsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBody) SetPageSize(v int32) *DescribeDedicatedHostsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBody) SetRequestId(v string) *DescribeDedicatedHostsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBody) SetTotalCount(v int32) *DescribeDedicatedHostsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBody) Validate() error {
	if s.DedicatedHosts != nil {
		if err := s.DedicatedHosts.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDedicatedHostsResponseBodyDedicatedHosts struct {
	DedicatedHost []*DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost `json:"DedicatedHost,omitempty" xml:"DedicatedHost,omitempty" type:"Repeated"`
}

func (s DescribeDedicatedHostsResponseBodyDedicatedHosts) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostsResponseBodyDedicatedHosts) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHosts) GetDedicatedHost() []*DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost {
	return s.DedicatedHost
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHosts) SetDedicatedHost(v []*DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) *DescribeDedicatedHostsResponseBodyDedicatedHosts {
	s.DedicatedHost = v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHosts) Validate() error {
	if s.DedicatedHost != nil {
		for _, item := range s.DedicatedHost {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost struct {
	SchedulerOptions                    *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostSchedulerOptions                    `json:"SchedulerOptions,omitempty" xml:"SchedulerOptions,omitempty" type:"Struct"`
	ActionOnMaintenance                 *string                                                                                           `json:"ActionOnMaintenance,omitempty" xml:"ActionOnMaintenance,omitempty"`
	AutoPlacement                       *string                                                                                           `json:"AutoPlacement,omitempty" xml:"AutoPlacement,omitempty"`
	AutoReleaseTime                     *string                                                                                           `json:"AutoReleaseTime,omitempty" xml:"AutoReleaseTime,omitempty"`
	Capacity                            *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacity                            `json:"Capacity,omitempty" xml:"Capacity,omitempty" type:"Struct"`
	ChargeType                          *string                                                                                           `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	Cores                               *int32                                                                                            `json:"Cores,omitempty" xml:"Cores,omitempty"`
	CpuOverCommitRatio                  *float32                                                                                          `json:"CpuOverCommitRatio,omitempty" xml:"CpuOverCommitRatio,omitempty"`
	CreationTime                        *string                                                                                           `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	DedicatedHostClusterId              *string                                                                                           `json:"DedicatedHostClusterId,omitempty" xml:"DedicatedHostClusterId,omitempty"`
	DedicatedHostId                     *string                                                                                           `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	DedicatedHostName                   *string                                                                                           `json:"DedicatedHostName,omitempty" xml:"DedicatedHostName,omitempty"`
	DedicatedHostOwnerId                *int64                                                                                            `json:"DedicatedHostOwnerId,omitempty" xml:"DedicatedHostOwnerId,omitempty"`
	DedicatedHostType                   *string                                                                                           `json:"DedicatedHostType,omitempty" xml:"DedicatedHostType,omitempty"`
	Description                         *string                                                                                           `json:"Description,omitempty" xml:"Description,omitempty"`
	ExpiredTime                         *string                                                                                           `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	GPUSpec                             *string                                                                                           `json:"GPUSpec,omitempty" xml:"GPUSpec,omitempty"`
	HostDetailInfo                      *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostHostDetailInfo                      `json:"HostDetailInfo,omitempty" xml:"HostDetailInfo,omitempty" type:"Struct"`
	Instances                           *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostInstances                           `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Struct"`
	MachineId                           *string                                                                                           `json:"MachineId,omitempty" xml:"MachineId,omitempty"`
	NetworkAttributes                   *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostNetworkAttributes                   `json:"NetworkAttributes,omitempty" xml:"NetworkAttributes,omitempty" type:"Struct"`
	OperationLocks                      *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostOperationLocks                      `json:"OperationLocks,omitempty" xml:"OperationLocks,omitempty" type:"Struct"`
	PhysicalGpus                        *int32                                                                                            `json:"PhysicalGpus,omitempty" xml:"PhysicalGpus,omitempty"`
	RegionId                            *string                                                                                           `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId                     *string                                                                                           `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SaleCycle                           *string                                                                                           `json:"SaleCycle,omitempty" xml:"SaleCycle,omitempty"`
	Sockets                             *int32                                                                                            `json:"Sockets,omitempty" xml:"Sockets,omitempty"`
	Status                              *string                                                                                           `json:"Status,omitempty" xml:"Status,omitempty"`
	SupportedCustomInstanceTypeFamilies *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostSupportedCustomInstanceTypeFamilies `json:"SupportedCustomInstanceTypeFamilies,omitempty" xml:"SupportedCustomInstanceTypeFamilies,omitempty" type:"Struct"`
	SupportedInstanceTypeFamilies       *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostSupportedInstanceTypeFamilies       `json:"SupportedInstanceTypeFamilies,omitempty" xml:"SupportedInstanceTypeFamilies,omitempty" type:"Struct"`
	SupportedInstanceTypesList          *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostSupportedInstanceTypesList          `json:"SupportedInstanceTypesList,omitempty" xml:"SupportedInstanceTypesList,omitempty" type:"Struct"`
	Tags                                *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostTags                                `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	ZoneId                              *string                                                                                           `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) GetSchedulerOptions() *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostSchedulerOptions {
	return s.SchedulerOptions
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) GetActionOnMaintenance() *string {
	return s.ActionOnMaintenance
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) GetAutoPlacement() *string {
	return s.AutoPlacement
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) GetAutoReleaseTime() *string {
	return s.AutoReleaseTime
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) GetCapacity() *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacity {
	return s.Capacity
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) GetCores() *int32 {
	return s.Cores
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) GetCpuOverCommitRatio() *float32 {
	return s.CpuOverCommitRatio
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) GetDedicatedHostClusterId() *string {
	return s.DedicatedHostClusterId
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) GetDedicatedHostId() *string {
	return s.DedicatedHostId
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) GetDedicatedHostName() *string {
	return s.DedicatedHostName
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) GetDedicatedHostOwnerId() *int64 {
	return s.DedicatedHostOwnerId
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) GetDedicatedHostType() *string {
	return s.DedicatedHostType
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) GetDescription() *string {
	return s.Description
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) GetGPUSpec() *string {
	return s.GPUSpec
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) GetHostDetailInfo() *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostHostDetailInfo {
	return s.HostDetailInfo
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) GetInstances() *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostInstances {
	return s.Instances
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) GetMachineId() *string {
	return s.MachineId
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) GetNetworkAttributes() *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostNetworkAttributes {
	return s.NetworkAttributes
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) GetOperationLocks() *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostOperationLocks {
	return s.OperationLocks
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) GetPhysicalGpus() *int32 {
	return s.PhysicalGpus
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) GetSaleCycle() *string {
	return s.SaleCycle
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) GetSockets() *int32 {
	return s.Sockets
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) GetStatus() *string {
	return s.Status
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) GetSupportedCustomInstanceTypeFamilies() *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostSupportedCustomInstanceTypeFamilies {
	return s.SupportedCustomInstanceTypeFamilies
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) GetSupportedInstanceTypeFamilies() *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostSupportedInstanceTypeFamilies {
	return s.SupportedInstanceTypeFamilies
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) GetSupportedInstanceTypesList() *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostSupportedInstanceTypesList {
	return s.SupportedInstanceTypesList
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) GetTags() *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostTags {
	return s.Tags
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) SetSchedulerOptions(v *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostSchedulerOptions) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost {
	s.SchedulerOptions = v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) SetActionOnMaintenance(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost {
	s.ActionOnMaintenance = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) SetAutoPlacement(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost {
	s.AutoPlacement = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) SetAutoReleaseTime(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost {
	s.AutoReleaseTime = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) SetCapacity(v *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacity) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost {
	s.Capacity = v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) SetChargeType(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost {
	s.ChargeType = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) SetCores(v int32) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost {
	s.Cores = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) SetCpuOverCommitRatio(v float32) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost {
	s.CpuOverCommitRatio = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) SetCreationTime(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost {
	s.CreationTime = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) SetDedicatedHostClusterId(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost {
	s.DedicatedHostClusterId = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) SetDedicatedHostId(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost {
	s.DedicatedHostId = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) SetDedicatedHostName(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost {
	s.DedicatedHostName = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) SetDedicatedHostOwnerId(v int64) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost {
	s.DedicatedHostOwnerId = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) SetDedicatedHostType(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost {
	s.DedicatedHostType = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) SetDescription(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost {
	s.Description = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) SetExpiredTime(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) SetGPUSpec(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost {
	s.GPUSpec = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) SetHostDetailInfo(v *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostHostDetailInfo) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost {
	s.HostDetailInfo = v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) SetInstances(v *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostInstances) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost {
	s.Instances = v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) SetMachineId(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost {
	s.MachineId = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) SetNetworkAttributes(v *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostNetworkAttributes) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost {
	s.NetworkAttributes = v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) SetOperationLocks(v *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostOperationLocks) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost {
	s.OperationLocks = v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) SetPhysicalGpus(v int32) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost {
	s.PhysicalGpus = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) SetRegionId(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost {
	s.RegionId = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) SetResourceGroupId(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) SetSaleCycle(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost {
	s.SaleCycle = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) SetSockets(v int32) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost {
	s.Sockets = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) SetStatus(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost {
	s.Status = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) SetSupportedCustomInstanceTypeFamilies(v *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostSupportedCustomInstanceTypeFamilies) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost {
	s.SupportedCustomInstanceTypeFamilies = v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) SetSupportedInstanceTypeFamilies(v *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostSupportedInstanceTypeFamilies) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost {
	s.SupportedInstanceTypeFamilies = v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) SetSupportedInstanceTypesList(v *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostSupportedInstanceTypesList) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost {
	s.SupportedInstanceTypesList = v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) SetTags(v *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostTags) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost {
	s.Tags = v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) SetZoneId(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost {
	s.ZoneId = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHost) Validate() error {
	if s.SchedulerOptions != nil {
		if err := s.SchedulerOptions.Validate(); err != nil {
			return err
		}
	}
	if s.Capacity != nil {
		if err := s.Capacity.Validate(); err != nil {
			return err
		}
	}
	if s.HostDetailInfo != nil {
		if err := s.HostDetailInfo.Validate(); err != nil {
			return err
		}
	}
	if s.Instances != nil {
		if err := s.Instances.Validate(); err != nil {
			return err
		}
	}
	if s.NetworkAttributes != nil {
		if err := s.NetworkAttributes.Validate(); err != nil {
			return err
		}
	}
	if s.OperationLocks != nil {
		if err := s.OperationLocks.Validate(); err != nil {
			return err
		}
	}
	if s.SupportedCustomInstanceTypeFamilies != nil {
		if err := s.SupportedCustomInstanceTypeFamilies.Validate(); err != nil {
			return err
		}
	}
	if s.SupportedInstanceTypeFamilies != nil {
		if err := s.SupportedInstanceTypeFamilies.Validate(); err != nil {
			return err
		}
	}
	if s.SupportedInstanceTypesList != nil {
		if err := s.SupportedInstanceTypesList.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostSchedulerOptions struct {
	ManagedPrivateSpaceId *string `json:"ManagedPrivateSpaceId,omitempty" xml:"ManagedPrivateSpaceId,omitempty"`
}

func (s DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostSchedulerOptions) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostSchedulerOptions) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostSchedulerOptions) GetManagedPrivateSpaceId() *string {
	return s.ManagedPrivateSpaceId
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostSchedulerOptions) SetManagedPrivateSpaceId(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostSchedulerOptions {
	s.ManagedPrivateSpaceId = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostSchedulerOptions) Validate() error {
	return dara.Validate(s)
}

type DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacity struct {
	AvailableInstanceTypes *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacityAvailableInstanceTypes `json:"AvailableInstanceTypes,omitempty" xml:"AvailableInstanceTypes,omitempty" type:"Struct"`
	AvailableLocalStorage  *int32                                                                                       `json:"AvailableLocalStorage,omitempty" xml:"AvailableLocalStorage,omitempty"`
	AvailableMemory        *float32                                                                                     `json:"AvailableMemory,omitempty" xml:"AvailableMemory,omitempty"`
	AvailableVcpus         *int32                                                                                       `json:"AvailableVcpus,omitempty" xml:"AvailableVcpus,omitempty"`
	AvailableVgpus         *int32                                                                                       `json:"AvailableVgpus,omitempty" xml:"AvailableVgpus,omitempty"`
	LocalStorageCategory   *string                                                                                      `json:"LocalStorageCategory,omitempty" xml:"LocalStorageCategory,omitempty"`
	SocketCapacities       *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacitySocketCapacities       `json:"SocketCapacities,omitempty" xml:"SocketCapacities,omitempty" type:"Struct"`
	TotalLocalStorage      *int32                                                                                       `json:"TotalLocalStorage,omitempty" xml:"TotalLocalStorage,omitempty"`
	TotalMemory            *float32                                                                                     `json:"TotalMemory,omitempty" xml:"TotalMemory,omitempty"`
	TotalVcpus             *int32                                                                                       `json:"TotalVcpus,omitempty" xml:"TotalVcpus,omitempty"`
	TotalVgpus             *int32                                                                                       `json:"TotalVgpus,omitempty" xml:"TotalVgpus,omitempty"`
}

func (s DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacity) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacity) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacity) GetAvailableInstanceTypes() *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacityAvailableInstanceTypes {
	return s.AvailableInstanceTypes
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacity) GetAvailableLocalStorage() *int32 {
	return s.AvailableLocalStorage
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacity) GetAvailableMemory() *float32 {
	return s.AvailableMemory
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacity) GetAvailableVcpus() *int32 {
	return s.AvailableVcpus
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacity) GetAvailableVgpus() *int32 {
	return s.AvailableVgpus
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacity) GetLocalStorageCategory() *string {
	return s.LocalStorageCategory
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacity) GetSocketCapacities() *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacitySocketCapacities {
	return s.SocketCapacities
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacity) GetTotalLocalStorage() *int32 {
	return s.TotalLocalStorage
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacity) GetTotalMemory() *float32 {
	return s.TotalMemory
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacity) GetTotalVcpus() *int32 {
	return s.TotalVcpus
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacity) GetTotalVgpus() *int32 {
	return s.TotalVgpus
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacity) SetAvailableInstanceTypes(v *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacityAvailableInstanceTypes) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacity {
	s.AvailableInstanceTypes = v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacity) SetAvailableLocalStorage(v int32) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacity {
	s.AvailableLocalStorage = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacity) SetAvailableMemory(v float32) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacity {
	s.AvailableMemory = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacity) SetAvailableVcpus(v int32) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacity {
	s.AvailableVcpus = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacity) SetAvailableVgpus(v int32) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacity {
	s.AvailableVgpus = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacity) SetLocalStorageCategory(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacity {
	s.LocalStorageCategory = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacity) SetSocketCapacities(v *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacitySocketCapacities) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacity {
	s.SocketCapacities = v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacity) SetTotalLocalStorage(v int32) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacity {
	s.TotalLocalStorage = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacity) SetTotalMemory(v float32) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacity {
	s.TotalMemory = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacity) SetTotalVcpus(v int32) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacity {
	s.TotalVcpus = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacity) SetTotalVgpus(v int32) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacity {
	s.TotalVgpus = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacity) Validate() error {
	if s.AvailableInstanceTypes != nil {
		if err := s.AvailableInstanceTypes.Validate(); err != nil {
			return err
		}
	}
	if s.SocketCapacities != nil {
		if err := s.SocketCapacities.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacityAvailableInstanceTypes struct {
	AvailableInstanceType []*DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacityAvailableInstanceTypesAvailableInstanceType `json:"AvailableInstanceType,omitempty" xml:"AvailableInstanceType,omitempty" type:"Repeated"`
}

func (s DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacityAvailableInstanceTypes) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacityAvailableInstanceTypes) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacityAvailableInstanceTypes) GetAvailableInstanceType() []*DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacityAvailableInstanceTypesAvailableInstanceType {
	return s.AvailableInstanceType
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacityAvailableInstanceTypes) SetAvailableInstanceType(v []*DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacityAvailableInstanceTypesAvailableInstanceType) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacityAvailableInstanceTypes {
	s.AvailableInstanceType = v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacityAvailableInstanceTypes) Validate() error {
	if s.AvailableInstanceType != nil {
		for _, item := range s.AvailableInstanceType {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacityAvailableInstanceTypesAvailableInstanceType struct {
	AvailableInstanceCapacity *int32  `json:"AvailableInstanceCapacity,omitempty" xml:"AvailableInstanceCapacity,omitempty"`
	InstanceType              *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacityAvailableInstanceTypesAvailableInstanceType) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacityAvailableInstanceTypesAvailableInstanceType) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacityAvailableInstanceTypesAvailableInstanceType) GetAvailableInstanceCapacity() *int32 {
	return s.AvailableInstanceCapacity
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacityAvailableInstanceTypesAvailableInstanceType) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacityAvailableInstanceTypesAvailableInstanceType) SetAvailableInstanceCapacity(v int32) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacityAvailableInstanceTypesAvailableInstanceType {
	s.AvailableInstanceCapacity = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacityAvailableInstanceTypesAvailableInstanceType) SetInstanceType(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacityAvailableInstanceTypesAvailableInstanceType {
	s.InstanceType = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacityAvailableInstanceTypesAvailableInstanceType) Validate() error {
	return dara.Validate(s)
}

type DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacitySocketCapacities struct {
	SocketCapacity []*DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacitySocketCapacitiesSocketCapacity `json:"SocketCapacity,omitempty" xml:"SocketCapacity,omitempty" type:"Repeated"`
}

func (s DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacitySocketCapacities) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacitySocketCapacities) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacitySocketCapacities) GetSocketCapacity() []*DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacitySocketCapacitiesSocketCapacity {
	return s.SocketCapacity
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacitySocketCapacities) SetSocketCapacity(v []*DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacitySocketCapacitiesSocketCapacity) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacitySocketCapacities {
	s.SocketCapacity = v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacitySocketCapacities) Validate() error {
	if s.SocketCapacity != nil {
		for _, item := range s.SocketCapacity {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacitySocketCapacitiesSocketCapacity struct {
	AvailableMemory *float32 `json:"AvailableMemory,omitempty" xml:"AvailableMemory,omitempty"`
	AvailableVcpu   *int32   `json:"AvailableVcpu,omitempty" xml:"AvailableVcpu,omitempty"`
	SocketId        *int32   `json:"SocketId,omitempty" xml:"SocketId,omitempty"`
	TotalMemory     *float32 `json:"TotalMemory,omitempty" xml:"TotalMemory,omitempty"`
	TotalVcpu       *int32   `json:"TotalVcpu,omitempty" xml:"TotalVcpu,omitempty"`
}

func (s DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacitySocketCapacitiesSocketCapacity) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacitySocketCapacitiesSocketCapacity) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacitySocketCapacitiesSocketCapacity) GetAvailableMemory() *float32 {
	return s.AvailableMemory
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacitySocketCapacitiesSocketCapacity) GetAvailableVcpu() *int32 {
	return s.AvailableVcpu
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacitySocketCapacitiesSocketCapacity) GetSocketId() *int32 {
	return s.SocketId
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacitySocketCapacitiesSocketCapacity) GetTotalMemory() *float32 {
	return s.TotalMemory
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacitySocketCapacitiesSocketCapacity) GetTotalVcpu() *int32 {
	return s.TotalVcpu
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacitySocketCapacitiesSocketCapacity) SetAvailableMemory(v float32) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacitySocketCapacitiesSocketCapacity {
	s.AvailableMemory = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacitySocketCapacitiesSocketCapacity) SetAvailableVcpu(v int32) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacitySocketCapacitiesSocketCapacity {
	s.AvailableVcpu = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacitySocketCapacitiesSocketCapacity) SetSocketId(v int32) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacitySocketCapacitiesSocketCapacity {
	s.SocketId = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacitySocketCapacitiesSocketCapacity) SetTotalMemory(v float32) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacitySocketCapacitiesSocketCapacity {
	s.TotalMemory = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacitySocketCapacitiesSocketCapacity) SetTotalVcpu(v int32) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacitySocketCapacitiesSocketCapacity {
	s.TotalVcpu = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostCapacitySocketCapacitiesSocketCapacity) Validate() error {
	return dara.Validate(s)
}

type DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostHostDetailInfo struct {
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
}

func (s DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostHostDetailInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostHostDetailInfo) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostHostDetailInfo) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostHostDetailInfo) SetSerialNumber(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostHostDetailInfo {
	s.SerialNumber = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostHostDetailInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostInstances struct {
	Instance []*DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostInstancesInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
}

func (s DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostInstances) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostInstances) GetInstance() []*DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostInstancesInstance {
	return s.Instance
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostInstances) SetInstance(v []*DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostInstancesInstance) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostInstances {
	s.Instance = v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostInstances) Validate() error {
	if s.Instance != nil {
		for _, item := range s.Instance {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostInstancesInstance struct {
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceOwnerId *int64  `json:"InstanceOwnerId,omitempty" xml:"InstanceOwnerId,omitempty"`
	InstanceType    *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	SocketId        *string `json:"SocketId,omitempty" xml:"SocketId,omitempty"`
}

func (s DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostInstancesInstance) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostInstancesInstance) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostInstancesInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostInstancesInstance) GetInstanceOwnerId() *int64 {
	return s.InstanceOwnerId
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostInstancesInstance) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostInstancesInstance) GetSocketId() *string {
	return s.SocketId
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostInstancesInstance) SetInstanceId(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostInstancesInstance {
	s.InstanceId = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostInstancesInstance) SetInstanceOwnerId(v int64) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostInstancesInstance {
	s.InstanceOwnerId = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostInstancesInstance) SetInstanceType(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostInstancesInstance {
	s.InstanceType = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostInstancesInstance) SetSocketId(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostInstancesInstance {
	s.SocketId = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostInstancesInstance) Validate() error {
	return dara.Validate(s)
}

type DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostNetworkAttributes struct {
	SlbUdpTimeout *int32 `json:"SlbUdpTimeout,omitempty" xml:"SlbUdpTimeout,omitempty"`
	UdpTimeout    *int32 `json:"UdpTimeout,omitempty" xml:"UdpTimeout,omitempty"`
}

func (s DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostNetworkAttributes) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostNetworkAttributes) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostNetworkAttributes) GetSlbUdpTimeout() *int32 {
	return s.SlbUdpTimeout
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostNetworkAttributes) GetUdpTimeout() *int32 {
	return s.UdpTimeout
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostNetworkAttributes) SetSlbUdpTimeout(v int32) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostNetworkAttributes {
	s.SlbUdpTimeout = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostNetworkAttributes) SetUdpTimeout(v int32) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostNetworkAttributes {
	s.UdpTimeout = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostNetworkAttributes) Validate() error {
	return dara.Validate(s)
}

type DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostOperationLocks struct {
	OperationLock []*DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostOperationLocksOperationLock `json:"OperationLock,omitempty" xml:"OperationLock,omitempty" type:"Repeated"`
}

func (s DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostOperationLocks) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostOperationLocks) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostOperationLocks) GetOperationLock() []*DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostOperationLocksOperationLock {
	return s.OperationLock
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostOperationLocks) SetOperationLock(v []*DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostOperationLocksOperationLock) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostOperationLocks {
	s.OperationLock = v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostOperationLocks) Validate() error {
	if s.OperationLock != nil {
		for _, item := range s.OperationLock {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostOperationLocksOperationLock struct {
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
}

func (s DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostOperationLocksOperationLock) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostOperationLocksOperationLock) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostOperationLocksOperationLock) GetLockReason() *string {
	return s.LockReason
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostOperationLocksOperationLock) SetLockReason(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostOperationLocksOperationLock {
	s.LockReason = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostOperationLocksOperationLock) Validate() error {
	return dara.Validate(s)
}

type DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostSupportedCustomInstanceTypeFamilies struct {
	SupportedCustomInstanceTypeFamily []*string `json:"SupportedCustomInstanceTypeFamily,omitempty" xml:"SupportedCustomInstanceTypeFamily,omitempty" type:"Repeated"`
}

func (s DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostSupportedCustomInstanceTypeFamilies) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostSupportedCustomInstanceTypeFamilies) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostSupportedCustomInstanceTypeFamilies) GetSupportedCustomInstanceTypeFamily() []*string {
	return s.SupportedCustomInstanceTypeFamily
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostSupportedCustomInstanceTypeFamilies) SetSupportedCustomInstanceTypeFamily(v []*string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostSupportedCustomInstanceTypeFamilies {
	s.SupportedCustomInstanceTypeFamily = v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostSupportedCustomInstanceTypeFamilies) Validate() error {
	return dara.Validate(s)
}

type DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostSupportedInstanceTypeFamilies struct {
	SupportedInstanceTypeFamily []*string `json:"SupportedInstanceTypeFamily,omitempty" xml:"SupportedInstanceTypeFamily,omitempty" type:"Repeated"`
}

func (s DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostSupportedInstanceTypeFamilies) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostSupportedInstanceTypeFamilies) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostSupportedInstanceTypeFamilies) GetSupportedInstanceTypeFamily() []*string {
	return s.SupportedInstanceTypeFamily
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostSupportedInstanceTypeFamilies) SetSupportedInstanceTypeFamily(v []*string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostSupportedInstanceTypeFamilies {
	s.SupportedInstanceTypeFamily = v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostSupportedInstanceTypeFamilies) Validate() error {
	return dara.Validate(s)
}

type DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostSupportedInstanceTypesList struct {
	SupportedInstanceTypesList []*string `json:"SupportedInstanceTypesList,omitempty" xml:"SupportedInstanceTypesList,omitempty" type:"Repeated"`
}

func (s DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostSupportedInstanceTypesList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostSupportedInstanceTypesList) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostSupportedInstanceTypesList) GetSupportedInstanceTypesList() []*string {
	return s.SupportedInstanceTypesList
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostSupportedInstanceTypesList) SetSupportedInstanceTypesList(v []*string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostSupportedInstanceTypesList {
	s.SupportedInstanceTypesList = v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostSupportedInstanceTypesList) Validate() error {
	return dara.Validate(s)
}

type DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostTags struct {
	Tag []*DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostTags) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostTags) GetTag() []*DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostTagsTag {
	return s.Tag
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostTags) SetTag(v []*DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostTagsTag) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostTags {
	s.Tag = v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostTags) Validate() error {
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

type DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostTagsTag struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostTagsTag) SetTagKey(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostTagsTag {
	s.TagKey = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostTagsTag) SetTagValue(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostTagsTag {
	s.TagValue = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHostTagsTag) Validate() error {
	return dara.Validate(s)
}
