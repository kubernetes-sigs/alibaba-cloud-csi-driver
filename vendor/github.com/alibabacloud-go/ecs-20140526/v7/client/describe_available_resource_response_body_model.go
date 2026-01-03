// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAvailableResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAvailableZones(v *DescribeAvailableResourceResponseBodyAvailableZones) *DescribeAvailableResourceResponseBody
	GetAvailableZones() *DescribeAvailableResourceResponseBodyAvailableZones
	SetRequestId(v string) *DescribeAvailableResourceResponseBody
	GetRequestId() *string
}

type DescribeAvailableResourceResponseBody struct {
	// The information about the availability of resources in the zones.
	AvailableZones *DescribeAvailableResourceResponseBodyAvailableZones `json:"AvailableZones,omitempty" xml:"AvailableZones,omitempty" type:"Struct"`
	// example:
	//
	// 0041D94C-FB92-4C49-B115-259DA1C*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAvailableResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBody) GetAvailableZones() *DescribeAvailableResourceResponseBodyAvailableZones {
	return s.AvailableZones
}

func (s *DescribeAvailableResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAvailableResourceResponseBody) SetAvailableZones(v *DescribeAvailableResourceResponseBodyAvailableZones) *DescribeAvailableResourceResponseBody {
	s.AvailableZones = v
	return s
}

func (s *DescribeAvailableResourceResponseBody) SetRequestId(v string) *DescribeAvailableResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAvailableResourceResponseBody) Validate() error {
	if s.AvailableZones != nil {
		if err := s.AvailableZones.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAvailableResourceResponseBodyAvailableZones struct {
	AvailableZone []*DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone `json:"AvailableZone,omitempty" xml:"AvailableZone,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZones) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZones) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZones) GetAvailableZone() []*DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone {
	return s.AvailableZone
}

func (s *DescribeAvailableResourceResponseBodyAvailableZones) SetAvailableZone(v []*DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone) *DescribeAvailableResourceResponseBodyAvailableZones {
	s.AvailableZone = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZones) Validate() error {
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

type DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone struct {
	AvailableResources *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResources `json:"AvailableResources,omitempty" xml:"AvailableResources,omitempty" type:"Struct"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The resource status based on the stock level in the zone. Valid value:
	//
	// 	- WithStock: The resources are available and can be continuously replenished.
	//
	// 	- ClosedWithStock: Inventory is available, but resources will not be replenished. The ability to guarantee the supply of inventory is low. We recommend selecting a product specification in the WithStock state.
	//
	// 	- WithoutStock: The resource is out of stock and will be replenished. We recommend using other resources that are in stock.
	//
	// 	- ClosedWithoutStock: The resource is out of stock and will no longer be replenished. We recommend using other resources that are in stock.
	//
	// example:
	//
	// WithStock
	StatusCategory *string `json:"StatusCategory,omitempty" xml:"StatusCategory,omitempty"`
	// example:
	//
	// cn-hangzhou-e
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone) GetAvailableResources() *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResources {
	return s.AvailableResources
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone) GetStatus() *string {
	return s.Status
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone) GetStatusCategory() *string {
	return s.StatusCategory
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone) SetAvailableResources(v *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResources) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone {
	s.AvailableResources = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone) SetRegionId(v string) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone {
	s.RegionId = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone) SetStatus(v string) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone {
	s.Status = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone) SetStatusCategory(v string) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone {
	s.StatusCategory = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone) SetZoneId(v string) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone {
	s.ZoneId = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone) Validate() error {
	if s.AvailableResources != nil {
		if err := s.AvailableResources.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResources struct {
	AvailableResource []*DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource `json:"AvailableResource,omitempty" xml:"AvailableResource,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResources) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResources) GetAvailableResource() []*DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource {
	return s.AvailableResource
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResources) SetAvailableResource(v []*DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResources {
	s.AvailableResource = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResources) Validate() error {
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

type DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource struct {
	SupportedResources *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResources `json:"SupportedResources,omitempty" xml:"SupportedResources,omitempty" type:"Struct"`
	// example:
	//
	// InstanceType
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource) GetSupportedResources() *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResources {
	return s.SupportedResources
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource) GetType() *string {
	return s.Type
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource) SetSupportedResources(v *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResources) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource {
	s.SupportedResources = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource) SetType(v string) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource {
	s.Type = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource) Validate() error {
	if s.SupportedResources != nil {
		if err := s.SupportedResources.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResources struct {
	SupportedResource []*DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource `json:"SupportedResource,omitempty" xml:"SupportedResource,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResources) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResources) GetSupportedResource() []*DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource {
	return s.SupportedResource
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResources) SetSupportedResource(v []*DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResources {
	s.SupportedResource = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResources) Validate() error {
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

type DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource struct {
	// example:
	//
	// 2048
	Max *int32 `json:"Max,omitempty" xml:"Max,omitempty"`
	// example:
	//
	// 20
	Min *int32 `json:"Min,omitempty" xml:"Min,omitempty"`
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// WithStock
	StatusCategory *string `json:"StatusCategory,omitempty" xml:"StatusCategory,omitempty"`
	// example:
	//
	// GiB
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// example:
	//
	// ecs.d1ne.xlarge
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource) GetMax() *int32 {
	return s.Max
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource) GetMin() *int32 {
	return s.Min
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource) GetStatus() *string {
	return s.Status
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource) GetStatusCategory() *string {
	return s.StatusCategory
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource) GetUnit() *string {
	return s.Unit
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource) GetValue() *string {
	return s.Value
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource) SetMax(v int32) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource {
	s.Max = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource) SetMin(v int32) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource {
	s.Min = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource) SetStatus(v string) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource {
	s.Status = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource) SetStatusCategory(v string) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource {
	s.StatusCategory = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource) SetUnit(v string) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource {
	s.Unit = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource) SetValue(v string) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource {
	s.Value = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource) Validate() error {
	return dara.Validate(s)
}
