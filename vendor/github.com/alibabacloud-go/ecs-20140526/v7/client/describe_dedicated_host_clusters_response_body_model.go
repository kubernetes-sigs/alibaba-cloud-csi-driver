// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDedicatedHostClustersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDedicatedHostClusters(v *DescribeDedicatedHostClustersResponseBodyDedicatedHostClusters) *DescribeDedicatedHostClustersResponseBody
	GetDedicatedHostClusters() *DescribeDedicatedHostClustersResponseBodyDedicatedHostClusters
	SetPageNumber(v int32) *DescribeDedicatedHostClustersResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDedicatedHostClustersResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeDedicatedHostClustersResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeDedicatedHostClustersResponseBody
	GetTotalCount() *int32
}

type DescribeDedicatedHostClustersResponseBody struct {
	DedicatedHostClusters *DescribeDedicatedHostClustersResponseBodyDedicatedHostClusters `json:"DedicatedHostClusters,omitempty" xml:"DedicatedHostClusters,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 5
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 214A2187-B06F-4E49-A081-4D053466A8C7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of dedicated host clusters.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDedicatedHostClustersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostClustersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostClustersResponseBody) GetDedicatedHostClusters() *DescribeDedicatedHostClustersResponseBodyDedicatedHostClusters {
	return s.DedicatedHostClusters
}

func (s *DescribeDedicatedHostClustersResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDedicatedHostClustersResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDedicatedHostClustersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDedicatedHostClustersResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDedicatedHostClustersResponseBody) SetDedicatedHostClusters(v *DescribeDedicatedHostClustersResponseBodyDedicatedHostClusters) *DescribeDedicatedHostClustersResponseBody {
	s.DedicatedHostClusters = v
	return s
}

func (s *DescribeDedicatedHostClustersResponseBody) SetPageNumber(v int32) *DescribeDedicatedHostClustersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDedicatedHostClustersResponseBody) SetPageSize(v int32) *DescribeDedicatedHostClustersResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDedicatedHostClustersResponseBody) SetRequestId(v string) *DescribeDedicatedHostClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDedicatedHostClustersResponseBody) SetTotalCount(v int32) *DescribeDedicatedHostClustersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDedicatedHostClustersResponseBody) Validate() error {
	if s.DedicatedHostClusters != nil {
		if err := s.DedicatedHostClusters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDedicatedHostClustersResponseBodyDedicatedHostClusters struct {
	DedicatedHostCluster []*DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostCluster `json:"DedicatedHostCluster,omitempty" xml:"DedicatedHostCluster,omitempty" type:"Repeated"`
}

func (s DescribeDedicatedHostClustersResponseBodyDedicatedHostClusters) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostClustersResponseBodyDedicatedHostClusters) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClusters) GetDedicatedHostCluster() []*DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostCluster {
	return s.DedicatedHostCluster
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClusters) SetDedicatedHostCluster(v []*DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostCluster) *DescribeDedicatedHostClustersResponseBodyDedicatedHostClusters {
	s.DedicatedHostCluster = v
	return s
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClusters) Validate() error {
	if s.DedicatedHostCluster != nil {
		for _, item := range s.DedicatedHostCluster {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostCluster struct {
	DedicatedHostClusterCapacity *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacity `json:"DedicatedHostClusterCapacity,omitempty" xml:"DedicatedHostClusterCapacity,omitempty" type:"Struct"`
	DedicatedHostClusterId       *string                                                                                                         `json:"DedicatedHostClusterId,omitempty" xml:"DedicatedHostClusterId,omitempty"`
	DedicatedHostClusterName     *string                                                                                                         `json:"DedicatedHostClusterName,omitempty" xml:"DedicatedHostClusterName,omitempty"`
	DedicatedHostIds             *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostIds             `json:"DedicatedHostIds,omitempty" xml:"DedicatedHostIds,omitempty" type:"Struct"`
	Description                  *string                                                                                                         `json:"Description,omitempty" xml:"Description,omitempty"`
	RegionId                     *string                                                                                                         `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId              *string                                                                                                         `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tags                         *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterTags                         `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	ZoneId                       *string                                                                                                         `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostCluster) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostCluster) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostCluster) GetDedicatedHostClusterCapacity() *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacity {
	return s.DedicatedHostClusterCapacity
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostCluster) GetDedicatedHostClusterId() *string {
	return s.DedicatedHostClusterId
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostCluster) GetDedicatedHostClusterName() *string {
	return s.DedicatedHostClusterName
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostCluster) GetDedicatedHostIds() *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostIds {
	return s.DedicatedHostIds
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostCluster) GetDescription() *string {
	return s.Description
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostCluster) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostCluster) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostCluster) GetTags() *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterTags {
	return s.Tags
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostCluster) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostCluster) SetDedicatedHostClusterCapacity(v *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacity) *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostCluster {
	s.DedicatedHostClusterCapacity = v
	return s
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostCluster) SetDedicatedHostClusterId(v string) *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostCluster {
	s.DedicatedHostClusterId = &v
	return s
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostCluster) SetDedicatedHostClusterName(v string) *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostCluster {
	s.DedicatedHostClusterName = &v
	return s
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostCluster) SetDedicatedHostIds(v *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostIds) *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostCluster {
	s.DedicatedHostIds = v
	return s
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostCluster) SetDescription(v string) *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostCluster {
	s.Description = &v
	return s
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostCluster) SetRegionId(v string) *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostCluster {
	s.RegionId = &v
	return s
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostCluster) SetResourceGroupId(v string) *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostCluster {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostCluster) SetTags(v *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterTags) *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostCluster {
	s.Tags = v
	return s
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostCluster) SetZoneId(v string) *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostCluster {
	s.ZoneId = &v
	return s
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostCluster) Validate() error {
	if s.DedicatedHostClusterCapacity != nil {
		if err := s.DedicatedHostClusterCapacity.Validate(); err != nil {
			return err
		}
	}
	if s.DedicatedHostIds != nil {
		if err := s.DedicatedHostIds.Validate(); err != nil {
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

type DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacity struct {
	AvailableInstanceTypes *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacityAvailableInstanceTypes `json:"AvailableInstanceTypes,omitempty" xml:"AvailableInstanceTypes,omitempty" type:"Struct"`
	AvailableMemory        *int32                                                                                                                                `json:"AvailableMemory,omitempty" xml:"AvailableMemory,omitempty"`
	AvailableVcpus         *int32                                                                                                                                `json:"AvailableVcpus,omitempty" xml:"AvailableVcpus,omitempty"`
	LocalStorageCapacities *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacityLocalStorageCapacities `json:"LocalStorageCapacities,omitempty" xml:"LocalStorageCapacities,omitempty" type:"Struct"`
	TotalMemory            *int32                                                                                                                                `json:"TotalMemory,omitempty" xml:"TotalMemory,omitempty"`
	TotalVcpus             *int32                                                                                                                                `json:"TotalVcpus,omitempty" xml:"TotalVcpus,omitempty"`
}

func (s DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacity) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacity) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacity) GetAvailableInstanceTypes() *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacityAvailableInstanceTypes {
	return s.AvailableInstanceTypes
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacity) GetAvailableMemory() *int32 {
	return s.AvailableMemory
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacity) GetAvailableVcpus() *int32 {
	return s.AvailableVcpus
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacity) GetLocalStorageCapacities() *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacityLocalStorageCapacities {
	return s.LocalStorageCapacities
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacity) GetTotalMemory() *int32 {
	return s.TotalMemory
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacity) GetTotalVcpus() *int32 {
	return s.TotalVcpus
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacity) SetAvailableInstanceTypes(v *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacityAvailableInstanceTypes) *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacity {
	s.AvailableInstanceTypes = v
	return s
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacity) SetAvailableMemory(v int32) *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacity {
	s.AvailableMemory = &v
	return s
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacity) SetAvailableVcpus(v int32) *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacity {
	s.AvailableVcpus = &v
	return s
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacity) SetLocalStorageCapacities(v *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacityLocalStorageCapacities) *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacity {
	s.LocalStorageCapacities = v
	return s
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacity) SetTotalMemory(v int32) *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacity {
	s.TotalMemory = &v
	return s
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacity) SetTotalVcpus(v int32) *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacity {
	s.TotalVcpus = &v
	return s
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacity) Validate() error {
	if s.AvailableInstanceTypes != nil {
		if err := s.AvailableInstanceTypes.Validate(); err != nil {
			return err
		}
	}
	if s.LocalStorageCapacities != nil {
		if err := s.LocalStorageCapacities.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacityAvailableInstanceTypes struct {
	AvailableInstanceType []*DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacityAvailableInstanceTypesAvailableInstanceType `json:"AvailableInstanceType,omitempty" xml:"AvailableInstanceType,omitempty" type:"Repeated"`
}

func (s DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacityAvailableInstanceTypes) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacityAvailableInstanceTypes) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacityAvailableInstanceTypes) GetAvailableInstanceType() []*DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacityAvailableInstanceTypesAvailableInstanceType {
	return s.AvailableInstanceType
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacityAvailableInstanceTypes) SetAvailableInstanceType(v []*DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacityAvailableInstanceTypesAvailableInstanceType) *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacityAvailableInstanceTypes {
	s.AvailableInstanceType = v
	return s
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacityAvailableInstanceTypes) Validate() error {
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

type DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacityAvailableInstanceTypesAvailableInstanceType struct {
	AvailableInstanceCapacity *int32  `json:"AvailableInstanceCapacity,omitempty" xml:"AvailableInstanceCapacity,omitempty"`
	InstanceType              *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacityAvailableInstanceTypesAvailableInstanceType) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacityAvailableInstanceTypesAvailableInstanceType) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacityAvailableInstanceTypesAvailableInstanceType) GetAvailableInstanceCapacity() *int32 {
	return s.AvailableInstanceCapacity
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacityAvailableInstanceTypesAvailableInstanceType) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacityAvailableInstanceTypesAvailableInstanceType) SetAvailableInstanceCapacity(v int32) *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacityAvailableInstanceTypesAvailableInstanceType {
	s.AvailableInstanceCapacity = &v
	return s
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacityAvailableInstanceTypesAvailableInstanceType) SetInstanceType(v string) *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacityAvailableInstanceTypesAvailableInstanceType {
	s.InstanceType = &v
	return s
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacityAvailableInstanceTypesAvailableInstanceType) Validate() error {
	return dara.Validate(s)
}

type DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacityLocalStorageCapacities struct {
	LocalStorageCapacity []*DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacityLocalStorageCapacitiesLocalStorageCapacity `json:"LocalStorageCapacity,omitempty" xml:"LocalStorageCapacity,omitempty" type:"Repeated"`
}

func (s DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacityLocalStorageCapacities) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacityLocalStorageCapacities) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacityLocalStorageCapacities) GetLocalStorageCapacity() []*DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacityLocalStorageCapacitiesLocalStorageCapacity {
	return s.LocalStorageCapacity
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacityLocalStorageCapacities) SetLocalStorageCapacity(v []*DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacityLocalStorageCapacitiesLocalStorageCapacity) *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacityLocalStorageCapacities {
	s.LocalStorageCapacity = v
	return s
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacityLocalStorageCapacities) Validate() error {
	if s.LocalStorageCapacity != nil {
		for _, item := range s.LocalStorageCapacity {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacityLocalStorageCapacitiesLocalStorageCapacity struct {
	AvailableDisk    *int32  `json:"AvailableDisk,omitempty" xml:"AvailableDisk,omitempty"`
	DataDiskCategory *string `json:"DataDiskCategory,omitempty" xml:"DataDiskCategory,omitempty"`
	TotalDisk        *int32  `json:"TotalDisk,omitempty" xml:"TotalDisk,omitempty"`
}

func (s DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacityLocalStorageCapacitiesLocalStorageCapacity) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacityLocalStorageCapacitiesLocalStorageCapacity) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacityLocalStorageCapacitiesLocalStorageCapacity) GetAvailableDisk() *int32 {
	return s.AvailableDisk
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacityLocalStorageCapacitiesLocalStorageCapacity) GetDataDiskCategory() *string {
	return s.DataDiskCategory
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacityLocalStorageCapacitiesLocalStorageCapacity) GetTotalDisk() *int32 {
	return s.TotalDisk
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacityLocalStorageCapacitiesLocalStorageCapacity) SetAvailableDisk(v int32) *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacityLocalStorageCapacitiesLocalStorageCapacity {
	s.AvailableDisk = &v
	return s
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacityLocalStorageCapacitiesLocalStorageCapacity) SetDataDiskCategory(v string) *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacityLocalStorageCapacitiesLocalStorageCapacity {
	s.DataDiskCategory = &v
	return s
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacityLocalStorageCapacitiesLocalStorageCapacity) SetTotalDisk(v int32) *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacityLocalStorageCapacitiesLocalStorageCapacity {
	s.TotalDisk = &v
	return s
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostClusterCapacityLocalStorageCapacitiesLocalStorageCapacity) Validate() error {
	return dara.Validate(s)
}

type DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostIds struct {
	DedicatedHostId []*string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty" type:"Repeated"`
}

func (s DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostIds) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostIds) GetDedicatedHostId() []*string {
	return s.DedicatedHostId
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostIds) SetDedicatedHostId(v []*string) *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostIds {
	s.DedicatedHostId = v
	return s
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterDedicatedHostIds) Validate() error {
	return dara.Validate(s)
}

type DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterTags struct {
	Tag []*DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterTags) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterTags) GetTag() []*DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterTagsTag {
	return s.Tag
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterTags) SetTag(v []*DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterTagsTag) *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterTags {
	s.Tag = v
	return s
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterTags) Validate() error {
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

type DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterTagsTag struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterTagsTag) SetTagKey(v string) *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterTagsTag {
	s.TagKey = &v
	return s
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterTagsTag) SetTagValue(v string) *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterTagsTag {
	s.TagValue = &v
	return s
}

func (s *DescribeDedicatedHostClustersResponseBodyDedicatedHostClustersDedicatedHostClusterTagsTag) Validate() error {
	return dara.Validate(s)
}
