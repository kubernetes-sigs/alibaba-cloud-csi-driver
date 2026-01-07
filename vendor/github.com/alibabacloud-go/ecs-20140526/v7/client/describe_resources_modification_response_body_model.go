// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourcesModificationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAvailableZones(v *DescribeResourcesModificationResponseBodyAvailableZones) *DescribeResourcesModificationResponseBody
	GetAvailableZones() *DescribeResourcesModificationResponseBodyAvailableZones
	SetRequestId(v string) *DescribeResourcesModificationResponseBody
	GetRequestId() *string
}

type DescribeResourcesModificationResponseBody struct {
	// The information about the queried zones.
	AvailableZones *DescribeResourcesModificationResponseBodyAvailableZones `json:"AvailableZones,omitempty" xml:"AvailableZones,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeResourcesModificationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourcesModificationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourcesModificationResponseBody) GetAvailableZones() *DescribeResourcesModificationResponseBodyAvailableZones {
	return s.AvailableZones
}

func (s *DescribeResourcesModificationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeResourcesModificationResponseBody) SetAvailableZones(v *DescribeResourcesModificationResponseBodyAvailableZones) *DescribeResourcesModificationResponseBody {
	s.AvailableZones = v
	return s
}

func (s *DescribeResourcesModificationResponseBody) SetRequestId(v string) *DescribeResourcesModificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResourcesModificationResponseBody) Validate() error {
	if s.AvailableZones != nil {
		if err := s.AvailableZones.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeResourcesModificationResponseBodyAvailableZones struct {
	AvailableZone []*DescribeResourcesModificationResponseBodyAvailableZonesAvailableZone `json:"AvailableZone,omitempty" xml:"AvailableZone,omitempty" type:"Repeated"`
}

func (s DescribeResourcesModificationResponseBodyAvailableZones) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourcesModificationResponseBodyAvailableZones) GoString() string {
	return s.String()
}

func (s *DescribeResourcesModificationResponseBodyAvailableZones) GetAvailableZone() []*DescribeResourcesModificationResponseBodyAvailableZonesAvailableZone {
	return s.AvailableZone
}

func (s *DescribeResourcesModificationResponseBodyAvailableZones) SetAvailableZone(v []*DescribeResourcesModificationResponseBodyAvailableZonesAvailableZone) *DescribeResourcesModificationResponseBodyAvailableZones {
	s.AvailableZone = v
	return s
}

func (s *DescribeResourcesModificationResponseBodyAvailableZones) Validate() error {
	if s.AvailableZone != nil {
		for _, item := range s.AvailableZone {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeResourcesModificationResponseBodyAvailableZonesAvailableZone struct {
	// The resources that are available in the zone.
	AvailableResources *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResources `json:"AvailableResources,omitempty" xml:"AvailableResources,omitempty" type:"Struct"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The state of the resource. Valid values:
	//
	// 	- Available
	//
	// 	- SoldOut
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The category of the resource based on stock status. Valid values:
	//
	// 	- WithStock: resources that are in sufficient stock
	//
	// 	- ClosedWithStock: resources that are in insufficient stock
	//
	// 	- WithoutStock: resources that are out of stock
	//
	// example:
	//
	// WithStock
	StatusCategory *string `json:"StatusCategory,omitempty" xml:"StatusCategory,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-e
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeResourcesModificationResponseBodyAvailableZonesAvailableZone) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourcesModificationResponseBodyAvailableZonesAvailableZone) GoString() string {
	return s.String()
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZone) GetAvailableResources() *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResources {
	return s.AvailableResources
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZone) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZone) GetStatus() *string {
	return s.Status
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZone) GetStatusCategory() *string {
	return s.StatusCategory
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZone) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZone) SetAvailableResources(v *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResources) *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZone {
	s.AvailableResources = v
	return s
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZone) SetRegionId(v string) *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZone {
	s.RegionId = &v
	return s
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZone) SetStatus(v string) *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZone {
	s.Status = &v
	return s
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZone) SetStatusCategory(v string) *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZone {
	s.StatusCategory = &v
	return s
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZone) SetZoneId(v string) *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZone {
	s.ZoneId = &v
	return s
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZone) Validate() error {
	if s.AvailableResources != nil {
		if err := s.AvailableResources.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResources struct {
	AvailableResource []*DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource `json:"AvailableResource,omitempty" xml:"AvailableResource,omitempty" type:"Repeated"`
}

func (s DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResources) GoString() string {
	return s.String()
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResources) GetAvailableResource() []*DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource {
	return s.AvailableResource
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResources) SetAvailableResource(v []*DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource) *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResources {
	s.AvailableResource = v
	return s
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResources) Validate() error {
	if s.AvailableResource != nil {
		for _, item := range s.AvailableResource {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource struct {
	// The resource types that resources can be changed to after the resources meet specified conditions. If the conditions are met, you can change the current resource to a resource in the list.
	ConditionSupportedResources *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceConditionSupportedResources `json:"ConditionSupportedResources,omitempty" xml:"ConditionSupportedResources,omitempty" type:"Struct"`
	// The information about the supported resources.
	SupportedResources *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResources `json:"SupportedResources,omitempty" xml:"SupportedResources,omitempty" type:"Struct"`
	// The resource type. Valid values:
	//
	// 	- InstanceType
	//
	// 	- SystemDisk
	//
	// example:
	//
	// InstanceType
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource) GoString() string {
	return s.String()
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource) GetConditionSupportedResources() *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceConditionSupportedResources {
	return s.ConditionSupportedResources
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource) GetSupportedResources() *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResources {
	return s.SupportedResources
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource) GetType() *string {
	return s.Type
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource) SetConditionSupportedResources(v *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceConditionSupportedResources) *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource {
	s.ConditionSupportedResources = v
	return s
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource) SetSupportedResources(v *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResources) *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource {
	s.SupportedResources = v
	return s
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource) SetType(v string) *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource {
	s.Type = &v
	return s
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource) Validate() error {
	if s.ConditionSupportedResources != nil {
		if err := s.ConditionSupportedResources.Validate(); err != nil {
			return err
		}
	}
	if s.SupportedResources != nil {
		if err := s.SupportedResources.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceConditionSupportedResources struct {
	ConditionSupportedResource []*DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceConditionSupportedResourcesConditionSupportedResource `json:"ConditionSupportedResource,omitempty" xml:"ConditionSupportedResource,omitempty" type:"Repeated"`
}

func (s DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceConditionSupportedResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceConditionSupportedResources) GoString() string {
	return s.String()
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceConditionSupportedResources) GetConditionSupportedResource() []*DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceConditionSupportedResourcesConditionSupportedResource {
	return s.ConditionSupportedResource
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceConditionSupportedResources) SetConditionSupportedResource(v []*DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceConditionSupportedResourcesConditionSupportedResource) *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceConditionSupportedResources {
	s.ConditionSupportedResource = v
	return s
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceConditionSupportedResources) Validate() error {
	if s.ConditionSupportedResource != nil {
		for _, item := range s.ConditionSupportedResource {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceConditionSupportedResourcesConditionSupportedResource struct {
	// The conditions.
	Conditions *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceConditionSupportedResourcesConditionSupportedResourceConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Struct"`
	// The maximum disk capacity.
	//
	// This parameter takes effect only when the DestinationResource request parameter is set to SystemDisk.
	//
	// example:
	//
	// 2
	Max *int32 `json:"Max,omitempty" xml:"Max,omitempty"`
	// The minimum disk capacity.
	//
	// This parameter takes effect only when the DestinationResource request parameter is set to SystemDisk.
	//
	// example:
	//
	// 1
	Min *int32 `json:"Min,omitempty" xml:"Min,omitempty"`
	// The stock state of the resource. Valid values:
	//
	// 	- Available
	//
	// 	- SoldOut
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The category of the resource based on stock status. Valid values:
	//
	// 	- WithStock: resources that are in sufficient stock
	//
	// 	- ClosedWithStock: resources that are in insufficient stock
	//
	// 	- WithoutStock: resources that are out of stock
	//
	// example:
	//
	// WithStock
	StatusCategory *string `json:"StatusCategory,omitempty" xml:"StatusCategory,omitempty"`
	// The unit of the disk capacity.
	//
	// This parameter takes effect only when the DestinationResource request parameter is set to SystemDisk.
	//
	// example:
	//
	// null
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The resource type.
	//
	// example:
	//
	// ecs.g5.large
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceConditionSupportedResourcesConditionSupportedResource) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceConditionSupportedResourcesConditionSupportedResource) GoString() string {
	return s.String()
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceConditionSupportedResourcesConditionSupportedResource) GetConditions() *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceConditionSupportedResourcesConditionSupportedResourceConditions {
	return s.Conditions
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceConditionSupportedResourcesConditionSupportedResource) GetMax() *int32 {
	return s.Max
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceConditionSupportedResourcesConditionSupportedResource) GetMin() *int32 {
	return s.Min
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceConditionSupportedResourcesConditionSupportedResource) GetStatus() *string {
	return s.Status
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceConditionSupportedResourcesConditionSupportedResource) GetStatusCategory() *string {
	return s.StatusCategory
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceConditionSupportedResourcesConditionSupportedResource) GetUnit() *string {
	return s.Unit
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceConditionSupportedResourcesConditionSupportedResource) GetValue() *string {
	return s.Value
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceConditionSupportedResourcesConditionSupportedResource) SetConditions(v *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceConditionSupportedResourcesConditionSupportedResourceConditions) *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceConditionSupportedResourcesConditionSupportedResource {
	s.Conditions = v
	return s
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceConditionSupportedResourcesConditionSupportedResource) SetMax(v int32) *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceConditionSupportedResourcesConditionSupportedResource {
	s.Max = &v
	return s
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceConditionSupportedResourcesConditionSupportedResource) SetMin(v int32) *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceConditionSupportedResourcesConditionSupportedResource {
	s.Min = &v
	return s
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceConditionSupportedResourcesConditionSupportedResource) SetStatus(v string) *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceConditionSupportedResourcesConditionSupportedResource {
	s.Status = &v
	return s
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceConditionSupportedResourcesConditionSupportedResource) SetStatusCategory(v string) *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceConditionSupportedResourcesConditionSupportedResource {
	s.StatusCategory = &v
	return s
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceConditionSupportedResourcesConditionSupportedResource) SetUnit(v string) *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceConditionSupportedResourcesConditionSupportedResource {
	s.Unit = &v
	return s
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceConditionSupportedResourcesConditionSupportedResource) SetValue(v string) *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceConditionSupportedResourcesConditionSupportedResource {
	s.Value = &v
	return s
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceConditionSupportedResourcesConditionSupportedResource) Validate() error {
	if s.Conditions != nil {
		if err := s.Conditions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceConditionSupportedResourcesConditionSupportedResourceConditions struct {
	Condition []*DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceConditionSupportedResourcesConditionSupportedResourceConditionsCondition `json:"Condition,omitempty" xml:"Condition,omitempty" type:"Repeated"`
}

func (s DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceConditionSupportedResourcesConditionSupportedResourceConditions) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceConditionSupportedResourcesConditionSupportedResourceConditions) GoString() string {
	return s.String()
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceConditionSupportedResourcesConditionSupportedResourceConditions) GetCondition() []*DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceConditionSupportedResourcesConditionSupportedResourceConditionsCondition {
	return s.Condition
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceConditionSupportedResourcesConditionSupportedResourceConditions) SetCondition(v []*DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceConditionSupportedResourcesConditionSupportedResourceConditionsCondition) *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceConditionSupportedResourcesConditionSupportedResourceConditions {
	s.Condition = v
	return s
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceConditionSupportedResourcesConditionSupportedResourceConditions) Validate() error {
	if s.Condition != nil {
		for _, item := range s.Condition {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceConditionSupportedResourcesConditionSupportedResourceConditionsCondition struct {
	// The condition name. Valid value:
	//
	// DiskCategory, which indicates a disk category change.
	//
	// example:
	//
	// DiskCategory
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceConditionSupportedResourcesConditionSupportedResourceConditionsCondition) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceConditionSupportedResourcesConditionSupportedResourceConditionsCondition) GoString() string {
	return s.String()
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceConditionSupportedResourcesConditionSupportedResourceConditionsCondition) GetKey() *string {
	return s.Key
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceConditionSupportedResourcesConditionSupportedResourceConditionsCondition) SetKey(v string) *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceConditionSupportedResourcesConditionSupportedResourceConditionsCondition {
	s.Key = &v
	return s
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceConditionSupportedResourcesConditionSupportedResourceConditionsCondition) Validate() error {
	return dara.Validate(s)
}

type DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResources struct {
	SupportedResource []*DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource `json:"SupportedResource,omitempty" xml:"SupportedResource,omitempty" type:"Repeated"`
}

func (s DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResources) GoString() string {
	return s.String()
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResources) GetSupportedResource() []*DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource {
	return s.SupportedResource
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResources) SetSupportedResource(v []*DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource) *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResources {
	s.SupportedResource = v
	return s
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResources) Validate() error {
	if s.SupportedResource != nil {
		for _, item := range s.SupportedResource {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource struct {
	// The maximum disk capacity.
	//
	// This parameter takes effect only when the DestinationResource request parameter is set to SystemDisk.
	//
	// example:
	//
	// 2
	Max *int32 `json:"Max,omitempty" xml:"Max,omitempty"`
	// The minimum disk capacity.
	//
	// This parameter takes effect only when the DestinationResource request parameter is set to SystemDisk.
	//
	// example:
	//
	// 1
	Min *int32 `json:"Min,omitempty" xml:"Min,omitempty"`
	// The state of the resource. Valid values:
	//
	// 	- Available
	//
	// 	- SoldOut
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The category of the resource based on stock status. Valid values:
	//
	// 	- WithStock: resources that are in sufficient stock
	//
	// 	- ClosedWithStock: resources that are in insufficient stock
	//
	// 	- WithoutStock: resources that are out of stock
	//
	// example:
	//
	// WithStock
	StatusCategory *string `json:"StatusCategory,omitempty" xml:"StatusCategory,omitempty"`
	// The unit of the disk capacity. This parameter takes effect only when the DestinationResource request parameter is set to SystemDisk.
	//
	// example:
	//
	// null
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The resource type.
	//
	// example:
	//
	// ecs.g5.large
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource) GoString() string {
	return s.String()
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource) GetMax() *int32 {
	return s.Max
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource) GetMin() *int32 {
	return s.Min
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource) GetStatus() *string {
	return s.Status
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource) GetStatusCategory() *string {
	return s.StatusCategory
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource) GetUnit() *string {
	return s.Unit
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource) GetValue() *string {
	return s.Value
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource) SetMax(v int32) *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource {
	s.Max = &v
	return s
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource) SetMin(v int32) *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource {
	s.Min = &v
	return s
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource) SetStatus(v string) *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource {
	s.Status = &v
	return s
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource) SetStatusCategory(v string) *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource {
	s.StatusCategory = &v
	return s
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource) SetUnit(v string) *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource {
	s.Unit = &v
	return s
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource) SetValue(v string) *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource {
	s.Value = &v
	return s
}

func (s *DescribeResourcesModificationResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource) Validate() error {
	return dara.Validate(s)
}
