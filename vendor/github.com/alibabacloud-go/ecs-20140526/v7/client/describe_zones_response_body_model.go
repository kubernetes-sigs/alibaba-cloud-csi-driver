// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeZonesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeZonesResponseBody
	GetRequestId() *string
	SetZones(v *DescribeZonesResponseBodyZones) *DescribeZonesResponseBody
	GetZones() *DescribeZonesResponseBodyZones
}

type DescribeZonesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Details about the zones and their supported resources.
	Zones *DescribeZonesResponseBodyZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Struct"`
}

func (s DescribeZonesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeZonesResponseBody) GetZones() *DescribeZonesResponseBodyZones {
	return s.Zones
}

func (s *DescribeZonesResponseBody) SetRequestId(v string) *DescribeZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeZonesResponseBody) SetZones(v *DescribeZonesResponseBodyZones) *DescribeZonesResponseBody {
	s.Zones = v
	return s
}

func (s *DescribeZonesResponseBody) Validate() error {
	if s.Zones != nil {
		if err := s.Zones.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeZonesResponseBodyZones struct {
	Zone []*DescribeZonesResponseBodyZonesZone `json:"Zone,omitempty" xml:"Zone,omitempty" type:"Repeated"`
}

func (s DescribeZonesResponseBodyZones) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesResponseBodyZones) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZones) GetZone() []*DescribeZonesResponseBodyZonesZone {
	return s.Zone
}

func (s *DescribeZonesResponseBodyZones) SetZone(v []*DescribeZonesResponseBodyZonesZone) *DescribeZonesResponseBodyZones {
	s.Zone = v
	return s
}

func (s *DescribeZonesResponseBodyZones) Validate() error {
	if s.Zone != nil {
		for _, item := range s.Zone {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeZonesResponseBodyZonesZone struct {
	// The supported dedicated host types.
	AvailableDedicatedHostTypes *DescribeZonesResponseBodyZonesZoneAvailableDedicatedHostTypes `json:"AvailableDedicatedHostTypes,omitempty" xml:"AvailableDedicatedHostTypes,omitempty" type:"Struct"`
	// The categories of cloud disks that can be created. Valid values:
	//
	// 	- cloud: basic disk
	//
	// 	- cloud_ssd: standard SSD
	//
	// 	- cloud_efficiency: ultra disk
	//
	// 	- cloud_essd: ESSD
	AvailableDiskCategories *DescribeZonesResponseBodyZonesZoneAvailableDiskCategories `json:"AvailableDiskCategories,omitempty" xml:"AvailableDiskCategories,omitempty" type:"Struct"`
	// The supported instance types.
	AvailableInstanceTypes *DescribeZonesResponseBodyZonesZoneAvailableInstanceTypes `json:"AvailableInstanceTypes,omitempty" xml:"AvailableInstanceTypes,omitempty" type:"Struct"`
	// The types of resources that can be created. Valid values:
	//
	// 	- VSwitch: vSwitch
	//
	// 	- IoOptimized: I/O optimized instance
	//
	// 	- Instance: instance
	//
	// 	- DedicatedHost: dedicated host
	//
	// 	- disk: cloud disk
	AvailableResourceCreation *DescribeZonesResponseBodyZonesZoneAvailableResourceCreation `json:"AvailableResourceCreation,omitempty" xml:"AvailableResourceCreation,omitempty" type:"Struct"`
	// Details about the resources that can be created in the zone.
	AvailableResources *DescribeZonesResponseBodyZonesZoneAvailableResources `json:"AvailableResources,omitempty" xml:"AvailableResources,omitempty" type:"Struct"`
	// The supported Shared Block Storage device categories.
	AvailableVolumeCategories *DescribeZonesResponseBodyZonesZoneAvailableVolumeCategories `json:"AvailableVolumeCategories,omitempty" xml:"AvailableVolumeCategories,omitempty" type:"Struct"`
	// The supported generations of dedicated hosts.
	DedicatedHostGenerations *DescribeZonesResponseBodyZonesZoneDedicatedHostGenerations `json:"DedicatedHostGenerations,omitempty" xml:"DedicatedHostGenerations,omitempty" type:"Struct"`
	// The name of the zone in the local language.
	//
	// example:
	//
	// Hangzhou Zone G
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The ID of the zone.
	//
	// example:
	//
	// cn-hangzhou-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// The type of the zone. Valid values:
	//
	// 	- AvailabilityZone: zone for the Alibaba Cloud public cloud
	//
	// 	- CloudBoxZone: zone for CloudBox
	//
	// example:
	//
	// AvailabilityZone
	ZoneType *string `json:"ZoneType,omitempty" xml:"ZoneType,omitempty"`
}

func (s DescribeZonesResponseBodyZonesZone) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesResponseBodyZonesZone) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZonesZone) GetAvailableDedicatedHostTypes() *DescribeZonesResponseBodyZonesZoneAvailableDedicatedHostTypes {
	return s.AvailableDedicatedHostTypes
}

func (s *DescribeZonesResponseBodyZonesZone) GetAvailableDiskCategories() *DescribeZonesResponseBodyZonesZoneAvailableDiskCategories {
	return s.AvailableDiskCategories
}

func (s *DescribeZonesResponseBodyZonesZone) GetAvailableInstanceTypes() *DescribeZonesResponseBodyZonesZoneAvailableInstanceTypes {
	return s.AvailableInstanceTypes
}

func (s *DescribeZonesResponseBodyZonesZone) GetAvailableResourceCreation() *DescribeZonesResponseBodyZonesZoneAvailableResourceCreation {
	return s.AvailableResourceCreation
}

func (s *DescribeZonesResponseBodyZonesZone) GetAvailableResources() *DescribeZonesResponseBodyZonesZoneAvailableResources {
	return s.AvailableResources
}

func (s *DescribeZonesResponseBodyZonesZone) GetAvailableVolumeCategories() *DescribeZonesResponseBodyZonesZoneAvailableVolumeCategories {
	return s.AvailableVolumeCategories
}

func (s *DescribeZonesResponseBodyZonesZone) GetDedicatedHostGenerations() *DescribeZonesResponseBodyZonesZoneDedicatedHostGenerations {
	return s.DedicatedHostGenerations
}

func (s *DescribeZonesResponseBodyZonesZone) GetLocalName() *string {
	return s.LocalName
}

func (s *DescribeZonesResponseBodyZonesZone) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeZonesResponseBodyZonesZone) GetZoneType() *string {
	return s.ZoneType
}

func (s *DescribeZonesResponseBodyZonesZone) SetAvailableDedicatedHostTypes(v *DescribeZonesResponseBodyZonesZoneAvailableDedicatedHostTypes) *DescribeZonesResponseBodyZonesZone {
	s.AvailableDedicatedHostTypes = v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetAvailableDiskCategories(v *DescribeZonesResponseBodyZonesZoneAvailableDiskCategories) *DescribeZonesResponseBodyZonesZone {
	s.AvailableDiskCategories = v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetAvailableInstanceTypes(v *DescribeZonesResponseBodyZonesZoneAvailableInstanceTypes) *DescribeZonesResponseBodyZonesZone {
	s.AvailableInstanceTypes = v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetAvailableResourceCreation(v *DescribeZonesResponseBodyZonesZoneAvailableResourceCreation) *DescribeZonesResponseBodyZonesZone {
	s.AvailableResourceCreation = v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetAvailableResources(v *DescribeZonesResponseBodyZonesZoneAvailableResources) *DescribeZonesResponseBodyZonesZone {
	s.AvailableResources = v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetAvailableVolumeCategories(v *DescribeZonesResponseBodyZonesZoneAvailableVolumeCategories) *DescribeZonesResponseBodyZonesZone {
	s.AvailableVolumeCategories = v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetDedicatedHostGenerations(v *DescribeZonesResponseBodyZonesZoneDedicatedHostGenerations) *DescribeZonesResponseBodyZonesZone {
	s.DedicatedHostGenerations = v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetLocalName(v string) *DescribeZonesResponseBodyZonesZone {
	s.LocalName = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetZoneId(v string) *DescribeZonesResponseBodyZonesZone {
	s.ZoneId = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetZoneType(v string) *DescribeZonesResponseBodyZonesZone {
	s.ZoneType = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) Validate() error {
	if s.AvailableDedicatedHostTypes != nil {
		if err := s.AvailableDedicatedHostTypes.Validate(); err != nil {
			return err
		}
	}
	if s.AvailableDiskCategories != nil {
		if err := s.AvailableDiskCategories.Validate(); err != nil {
			return err
		}
	}
	if s.AvailableInstanceTypes != nil {
		if err := s.AvailableInstanceTypes.Validate(); err != nil {
			return err
		}
	}
	if s.AvailableResourceCreation != nil {
		if err := s.AvailableResourceCreation.Validate(); err != nil {
			return err
		}
	}
	if s.AvailableResources != nil {
		if err := s.AvailableResources.Validate(); err != nil {
			return err
		}
	}
	if s.AvailableVolumeCategories != nil {
		if err := s.AvailableVolumeCategories.Validate(); err != nil {
			return err
		}
	}
	if s.DedicatedHostGenerations != nil {
		if err := s.DedicatedHostGenerations.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeZonesResponseBodyZonesZoneAvailableDedicatedHostTypes struct {
	DedicatedHostType []*string `json:"DedicatedHostType,omitempty" xml:"DedicatedHostType,omitempty" type:"Repeated"`
}

func (s DescribeZonesResponseBodyZonesZoneAvailableDedicatedHostTypes) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesResponseBodyZonesZoneAvailableDedicatedHostTypes) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZonesZoneAvailableDedicatedHostTypes) GetDedicatedHostType() []*string {
	return s.DedicatedHostType
}

func (s *DescribeZonesResponseBodyZonesZoneAvailableDedicatedHostTypes) SetDedicatedHostType(v []*string) *DescribeZonesResponseBodyZonesZoneAvailableDedicatedHostTypes {
	s.DedicatedHostType = v
	return s
}

func (s *DescribeZonesResponseBodyZonesZoneAvailableDedicatedHostTypes) Validate() error {
	return dara.Validate(s)
}

type DescribeZonesResponseBodyZonesZoneAvailableDiskCategories struct {
	DiskCategories []*string `json:"DiskCategories,omitempty" xml:"DiskCategories,omitempty" type:"Repeated"`
}

func (s DescribeZonesResponseBodyZonesZoneAvailableDiskCategories) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesResponseBodyZonesZoneAvailableDiskCategories) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZonesZoneAvailableDiskCategories) GetDiskCategories() []*string {
	return s.DiskCategories
}

func (s *DescribeZonesResponseBodyZonesZoneAvailableDiskCategories) SetDiskCategories(v []*string) *DescribeZonesResponseBodyZonesZoneAvailableDiskCategories {
	s.DiskCategories = v
	return s
}

func (s *DescribeZonesResponseBodyZonesZoneAvailableDiskCategories) Validate() error {
	return dara.Validate(s)
}

type DescribeZonesResponseBodyZonesZoneAvailableInstanceTypes struct {
	InstanceTypes []*string `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
}

func (s DescribeZonesResponseBodyZonesZoneAvailableInstanceTypes) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesResponseBodyZonesZoneAvailableInstanceTypes) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZonesZoneAvailableInstanceTypes) GetInstanceTypes() []*string {
	return s.InstanceTypes
}

func (s *DescribeZonesResponseBodyZonesZoneAvailableInstanceTypes) SetInstanceTypes(v []*string) *DescribeZonesResponseBodyZonesZoneAvailableInstanceTypes {
	s.InstanceTypes = v
	return s
}

func (s *DescribeZonesResponseBodyZonesZoneAvailableInstanceTypes) Validate() error {
	return dara.Validate(s)
}

type DescribeZonesResponseBodyZonesZoneAvailableResourceCreation struct {
	ResourceTypes []*string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty" type:"Repeated"`
}

func (s DescribeZonesResponseBodyZonesZoneAvailableResourceCreation) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesResponseBodyZonesZoneAvailableResourceCreation) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZonesZoneAvailableResourceCreation) GetResourceTypes() []*string {
	return s.ResourceTypes
}

func (s *DescribeZonesResponseBodyZonesZoneAvailableResourceCreation) SetResourceTypes(v []*string) *DescribeZonesResponseBodyZonesZoneAvailableResourceCreation {
	s.ResourceTypes = v
	return s
}

func (s *DescribeZonesResponseBodyZonesZoneAvailableResourceCreation) Validate() error {
	return dara.Validate(s)
}

type DescribeZonesResponseBodyZonesZoneAvailableResources struct {
	ResourcesInfo []*DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfo `json:"ResourcesInfo,omitempty" xml:"ResourcesInfo,omitempty" type:"Repeated"`
}

func (s DescribeZonesResponseBodyZonesZoneAvailableResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesResponseBodyZonesZoneAvailableResources) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZonesZoneAvailableResources) GetResourcesInfo() []*DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfo {
	return s.ResourcesInfo
}

func (s *DescribeZonesResponseBodyZonesZoneAvailableResources) SetResourcesInfo(v []*DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfo) *DescribeZonesResponseBodyZonesZoneAvailableResources {
	s.ResourcesInfo = v
	return s
}

func (s *DescribeZonesResponseBodyZonesZoneAvailableResources) Validate() error {
	if s.ResourcesInfo != nil {
		for _, item := range s.ResourcesInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfo struct {
	// The categories of data disks that can be created.
	DataDiskCategories *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoDataDiskCategories `json:"DataDiskCategories,omitempty" xml:"DataDiskCategories,omitempty" type:"Struct"`
	// The supported generations of instance families.
	InstanceGenerations *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoInstanceGenerations `json:"InstanceGenerations,omitempty" xml:"InstanceGenerations,omitempty" type:"Struct"`
	// The supported instance families.
	InstanceTypeFamilies *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoInstanceTypeFamilies `json:"InstanceTypeFamilies,omitempty" xml:"InstanceTypeFamilies,omitempty" type:"Struct"`
	// The supported instance types.
	InstanceTypes *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoInstanceTypes `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Struct"`
	// Indicates whether the instance is I/O optimized.
	//
	// example:
	//
	// true
	IoOptimized *bool `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	// The supported network types.
	NetworkTypes *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoNetworkTypes `json:"NetworkTypes,omitempty" xml:"NetworkTypes,omitempty" type:"Struct"`
	// The categories of system disks that can be created.
	SystemDiskCategories *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoSystemDiskCategories `json:"SystemDiskCategories,omitempty" xml:"SystemDiskCategories,omitempty" type:"Struct"`
}

func (s DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfo) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfo) GetDataDiskCategories() *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoDataDiskCategories {
	return s.DataDiskCategories
}

func (s *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfo) GetInstanceGenerations() *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoInstanceGenerations {
	return s.InstanceGenerations
}

func (s *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfo) GetInstanceTypeFamilies() *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoInstanceTypeFamilies {
	return s.InstanceTypeFamilies
}

func (s *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfo) GetInstanceTypes() *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoInstanceTypes {
	return s.InstanceTypes
}

func (s *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfo) GetIoOptimized() *bool {
	return s.IoOptimized
}

func (s *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfo) GetNetworkTypes() *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoNetworkTypes {
	return s.NetworkTypes
}

func (s *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfo) GetSystemDiskCategories() *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoSystemDiskCategories {
	return s.SystemDiskCategories
}

func (s *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfo) SetDataDiskCategories(v *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoDataDiskCategories) *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfo {
	s.DataDiskCategories = v
	return s
}

func (s *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfo) SetInstanceGenerations(v *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoInstanceGenerations) *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfo {
	s.InstanceGenerations = v
	return s
}

func (s *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfo) SetInstanceTypeFamilies(v *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoInstanceTypeFamilies) *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfo {
	s.InstanceTypeFamilies = v
	return s
}

func (s *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfo) SetInstanceTypes(v *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoInstanceTypes) *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfo {
	s.InstanceTypes = v
	return s
}

func (s *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfo) SetIoOptimized(v bool) *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfo {
	s.IoOptimized = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfo) SetNetworkTypes(v *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoNetworkTypes) *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfo {
	s.NetworkTypes = v
	return s
}

func (s *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfo) SetSystemDiskCategories(v *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoSystemDiskCategories) *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfo {
	s.SystemDiskCategories = v
	return s
}

func (s *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfo) Validate() error {
	if s.DataDiskCategories != nil {
		if err := s.DataDiskCategories.Validate(); err != nil {
			return err
		}
	}
	if s.InstanceGenerations != nil {
		if err := s.InstanceGenerations.Validate(); err != nil {
			return err
		}
	}
	if s.InstanceTypeFamilies != nil {
		if err := s.InstanceTypeFamilies.Validate(); err != nil {
			return err
		}
	}
	if s.InstanceTypes != nil {
		if err := s.InstanceTypes.Validate(); err != nil {
			return err
		}
	}
	if s.NetworkTypes != nil {
		if err := s.NetworkTypes.Validate(); err != nil {
			return err
		}
	}
	if s.SystemDiskCategories != nil {
		if err := s.SystemDiskCategories.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoDataDiskCategories struct {
	SupportedDataDiskCategory []*string `json:"supportedDataDiskCategory,omitempty" xml:"supportedDataDiskCategory,omitempty" type:"Repeated"`
}

func (s DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoDataDiskCategories) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoDataDiskCategories) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoDataDiskCategories) GetSupportedDataDiskCategory() []*string {
	return s.SupportedDataDiskCategory
}

func (s *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoDataDiskCategories) SetSupportedDataDiskCategory(v []*string) *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoDataDiskCategories {
	s.SupportedDataDiskCategory = v
	return s
}

func (s *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoDataDiskCategories) Validate() error {
	return dara.Validate(s)
}

type DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoInstanceGenerations struct {
	SupportedInstanceGeneration []*string `json:"supportedInstanceGeneration,omitempty" xml:"supportedInstanceGeneration,omitempty" type:"Repeated"`
}

func (s DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoInstanceGenerations) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoInstanceGenerations) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoInstanceGenerations) GetSupportedInstanceGeneration() []*string {
	return s.SupportedInstanceGeneration
}

func (s *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoInstanceGenerations) SetSupportedInstanceGeneration(v []*string) *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoInstanceGenerations {
	s.SupportedInstanceGeneration = v
	return s
}

func (s *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoInstanceGenerations) Validate() error {
	return dara.Validate(s)
}

type DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoInstanceTypeFamilies struct {
	SupportedInstanceTypeFamily []*string `json:"supportedInstanceTypeFamily,omitempty" xml:"supportedInstanceTypeFamily,omitempty" type:"Repeated"`
}

func (s DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoInstanceTypeFamilies) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoInstanceTypeFamilies) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoInstanceTypeFamilies) GetSupportedInstanceTypeFamily() []*string {
	return s.SupportedInstanceTypeFamily
}

func (s *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoInstanceTypeFamilies) SetSupportedInstanceTypeFamily(v []*string) *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoInstanceTypeFamilies {
	s.SupportedInstanceTypeFamily = v
	return s
}

func (s *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoInstanceTypeFamilies) Validate() error {
	return dara.Validate(s)
}

type DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoInstanceTypes struct {
	SupportedInstanceType []*string `json:"supportedInstanceType,omitempty" xml:"supportedInstanceType,omitempty" type:"Repeated"`
}

func (s DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoInstanceTypes) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoInstanceTypes) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoInstanceTypes) GetSupportedInstanceType() []*string {
	return s.SupportedInstanceType
}

func (s *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoInstanceTypes) SetSupportedInstanceType(v []*string) *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoInstanceTypes {
	s.SupportedInstanceType = v
	return s
}

func (s *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoInstanceTypes) Validate() error {
	return dara.Validate(s)
}

type DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoNetworkTypes struct {
	SupportedNetworkCategory []*string `json:"supportedNetworkCategory,omitempty" xml:"supportedNetworkCategory,omitempty" type:"Repeated"`
}

func (s DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoNetworkTypes) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoNetworkTypes) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoNetworkTypes) GetSupportedNetworkCategory() []*string {
	return s.SupportedNetworkCategory
}

func (s *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoNetworkTypes) SetSupportedNetworkCategory(v []*string) *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoNetworkTypes {
	s.SupportedNetworkCategory = v
	return s
}

func (s *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoNetworkTypes) Validate() error {
	return dara.Validate(s)
}

type DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoSystemDiskCategories struct {
	SupportedSystemDiskCategory []*string `json:"supportedSystemDiskCategory,omitempty" xml:"supportedSystemDiskCategory,omitempty" type:"Repeated"`
}

func (s DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoSystemDiskCategories) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoSystemDiskCategories) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoSystemDiskCategories) GetSupportedSystemDiskCategory() []*string {
	return s.SupportedSystemDiskCategory
}

func (s *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoSystemDiskCategories) SetSupportedSystemDiskCategory(v []*string) *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoSystemDiskCategories {
	s.SupportedSystemDiskCategory = v
	return s
}

func (s *DescribeZonesResponseBodyZonesZoneAvailableResourcesResourcesInfoSystemDiskCategories) Validate() error {
	return dara.Validate(s)
}

type DescribeZonesResponseBodyZonesZoneAvailableVolumeCategories struct {
	VolumeCategories []*string `json:"VolumeCategories,omitempty" xml:"VolumeCategories,omitempty" type:"Repeated"`
}

func (s DescribeZonesResponseBodyZonesZoneAvailableVolumeCategories) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesResponseBodyZonesZoneAvailableVolumeCategories) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZonesZoneAvailableVolumeCategories) GetVolumeCategories() []*string {
	return s.VolumeCategories
}

func (s *DescribeZonesResponseBodyZonesZoneAvailableVolumeCategories) SetVolumeCategories(v []*string) *DescribeZonesResponseBodyZonesZoneAvailableVolumeCategories {
	s.VolumeCategories = v
	return s
}

func (s *DescribeZonesResponseBodyZonesZoneAvailableVolumeCategories) Validate() error {
	return dara.Validate(s)
}

type DescribeZonesResponseBodyZonesZoneDedicatedHostGenerations struct {
	DedicatedHostGeneration []*string `json:"DedicatedHostGeneration,omitempty" xml:"DedicatedHostGeneration,omitempty" type:"Repeated"`
}

func (s DescribeZonesResponseBodyZonesZoneDedicatedHostGenerations) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesResponseBodyZonesZoneDedicatedHostGenerations) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZonesZoneDedicatedHostGenerations) GetDedicatedHostGeneration() []*string {
	return s.DedicatedHostGeneration
}

func (s *DescribeZonesResponseBodyZonesZoneDedicatedHostGenerations) SetDedicatedHostGeneration(v []*string) *DescribeZonesResponseBodyZonesZoneDedicatedHostGenerations {
	s.DedicatedHostGeneration = v
	return s
}

func (s *DescribeZonesResponseBodyZonesZoneDedicatedHostGenerations) Validate() error {
	return dara.Validate(s)
}
