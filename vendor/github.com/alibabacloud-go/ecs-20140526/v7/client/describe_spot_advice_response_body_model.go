// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSpotAdviceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAvailableSpotZones(v *DescribeSpotAdviceResponseBodyAvailableSpotZones) *DescribeSpotAdviceResponseBody
	GetAvailableSpotZones() *DescribeSpotAdviceResponseBodyAvailableSpotZones
	SetRegionId(v string) *DescribeSpotAdviceResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeSpotAdviceResponseBody
	GetRequestId() *string
}

type DescribeSpotAdviceResponseBody struct {
	AvailableSpotZones *DescribeSpotAdviceResponseBodyAvailableSpotZones `json:"AvailableSpotZones,omitempty" xml:"AvailableSpotZones,omitempty" type:"Struct"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSpotAdviceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSpotAdviceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSpotAdviceResponseBody) GetAvailableSpotZones() *DescribeSpotAdviceResponseBodyAvailableSpotZones {
	return s.AvailableSpotZones
}

func (s *DescribeSpotAdviceResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSpotAdviceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSpotAdviceResponseBody) SetAvailableSpotZones(v *DescribeSpotAdviceResponseBodyAvailableSpotZones) *DescribeSpotAdviceResponseBody {
	s.AvailableSpotZones = v
	return s
}

func (s *DescribeSpotAdviceResponseBody) SetRegionId(v string) *DescribeSpotAdviceResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeSpotAdviceResponseBody) SetRequestId(v string) *DescribeSpotAdviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSpotAdviceResponseBody) Validate() error {
	if s.AvailableSpotZones != nil {
		if err := s.AvailableSpotZones.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSpotAdviceResponseBodyAvailableSpotZones struct {
	AvailableSpotZone []*DescribeSpotAdviceResponseBodyAvailableSpotZonesAvailableSpotZone `json:"AvailableSpotZone,omitempty" xml:"AvailableSpotZone,omitempty" type:"Repeated"`
}

func (s DescribeSpotAdviceResponseBodyAvailableSpotZones) String() string {
	return dara.Prettify(s)
}

func (s DescribeSpotAdviceResponseBodyAvailableSpotZones) GoString() string {
	return s.String()
}

func (s *DescribeSpotAdviceResponseBodyAvailableSpotZones) GetAvailableSpotZone() []*DescribeSpotAdviceResponseBodyAvailableSpotZonesAvailableSpotZone {
	return s.AvailableSpotZone
}

func (s *DescribeSpotAdviceResponseBodyAvailableSpotZones) SetAvailableSpotZone(v []*DescribeSpotAdviceResponseBodyAvailableSpotZonesAvailableSpotZone) *DescribeSpotAdviceResponseBodyAvailableSpotZones {
	s.AvailableSpotZone = v
	return s
}

func (s *DescribeSpotAdviceResponseBodyAvailableSpotZones) Validate() error {
	if s.AvailableSpotZone != nil {
		for _, item := range s.AvailableSpotZone {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSpotAdviceResponseBodyAvailableSpotZonesAvailableSpotZone struct {
	AvailableSpotResources *DescribeSpotAdviceResponseBodyAvailableSpotZonesAvailableSpotZoneAvailableSpotResources `json:"AvailableSpotResources,omitempty" xml:"AvailableSpotResources,omitempty" type:"Struct"`
	ZoneId                 *string                                                                                  `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeSpotAdviceResponseBodyAvailableSpotZonesAvailableSpotZone) String() string {
	return dara.Prettify(s)
}

func (s DescribeSpotAdviceResponseBodyAvailableSpotZonesAvailableSpotZone) GoString() string {
	return s.String()
}

func (s *DescribeSpotAdviceResponseBodyAvailableSpotZonesAvailableSpotZone) GetAvailableSpotResources() *DescribeSpotAdviceResponseBodyAvailableSpotZonesAvailableSpotZoneAvailableSpotResources {
	return s.AvailableSpotResources
}

func (s *DescribeSpotAdviceResponseBodyAvailableSpotZonesAvailableSpotZone) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeSpotAdviceResponseBodyAvailableSpotZonesAvailableSpotZone) SetAvailableSpotResources(v *DescribeSpotAdviceResponseBodyAvailableSpotZonesAvailableSpotZoneAvailableSpotResources) *DescribeSpotAdviceResponseBodyAvailableSpotZonesAvailableSpotZone {
	s.AvailableSpotResources = v
	return s
}

func (s *DescribeSpotAdviceResponseBodyAvailableSpotZonesAvailableSpotZone) SetZoneId(v string) *DescribeSpotAdviceResponseBodyAvailableSpotZonesAvailableSpotZone {
	s.ZoneId = &v
	return s
}

func (s *DescribeSpotAdviceResponseBodyAvailableSpotZonesAvailableSpotZone) Validate() error {
	if s.AvailableSpotResources != nil {
		if err := s.AvailableSpotResources.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSpotAdviceResponseBodyAvailableSpotZonesAvailableSpotZoneAvailableSpotResources struct {
	AvailableSpotResource []*DescribeSpotAdviceResponseBodyAvailableSpotZonesAvailableSpotZoneAvailableSpotResourcesAvailableSpotResource `json:"AvailableSpotResource,omitempty" xml:"AvailableSpotResource,omitempty" type:"Repeated"`
}

func (s DescribeSpotAdviceResponseBodyAvailableSpotZonesAvailableSpotZoneAvailableSpotResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeSpotAdviceResponseBodyAvailableSpotZonesAvailableSpotZoneAvailableSpotResources) GoString() string {
	return s.String()
}

func (s *DescribeSpotAdviceResponseBodyAvailableSpotZonesAvailableSpotZoneAvailableSpotResources) GetAvailableSpotResource() []*DescribeSpotAdviceResponseBodyAvailableSpotZonesAvailableSpotZoneAvailableSpotResourcesAvailableSpotResource {
	return s.AvailableSpotResource
}

func (s *DescribeSpotAdviceResponseBodyAvailableSpotZonesAvailableSpotZoneAvailableSpotResources) SetAvailableSpotResource(v []*DescribeSpotAdviceResponseBodyAvailableSpotZonesAvailableSpotZoneAvailableSpotResourcesAvailableSpotResource) *DescribeSpotAdviceResponseBodyAvailableSpotZonesAvailableSpotZoneAvailableSpotResources {
	s.AvailableSpotResource = v
	return s
}

func (s *DescribeSpotAdviceResponseBodyAvailableSpotZonesAvailableSpotZoneAvailableSpotResources) Validate() error {
	if s.AvailableSpotResource != nil {
		for _, item := range s.AvailableSpotResource {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSpotAdviceResponseBodyAvailableSpotZonesAvailableSpotZoneAvailableSpotResourcesAvailableSpotResource struct {
	AverageSpotDiscount *int32   `json:"AverageSpotDiscount,omitempty" xml:"AverageSpotDiscount,omitempty"`
	InstanceType        *string  `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InterruptRateDesc   *string  `json:"InterruptRateDesc,omitempty" xml:"InterruptRateDesc,omitempty"`
	InterruptionRate    *float32 `json:"InterruptionRate,omitempty" xml:"InterruptionRate,omitempty"`
}

func (s DescribeSpotAdviceResponseBodyAvailableSpotZonesAvailableSpotZoneAvailableSpotResourcesAvailableSpotResource) String() string {
	return dara.Prettify(s)
}

func (s DescribeSpotAdviceResponseBodyAvailableSpotZonesAvailableSpotZoneAvailableSpotResourcesAvailableSpotResource) GoString() string {
	return s.String()
}

func (s *DescribeSpotAdviceResponseBodyAvailableSpotZonesAvailableSpotZoneAvailableSpotResourcesAvailableSpotResource) GetAverageSpotDiscount() *int32 {
	return s.AverageSpotDiscount
}

func (s *DescribeSpotAdviceResponseBodyAvailableSpotZonesAvailableSpotZoneAvailableSpotResourcesAvailableSpotResource) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeSpotAdviceResponseBodyAvailableSpotZonesAvailableSpotZoneAvailableSpotResourcesAvailableSpotResource) GetInterruptRateDesc() *string {
	return s.InterruptRateDesc
}

func (s *DescribeSpotAdviceResponseBodyAvailableSpotZonesAvailableSpotZoneAvailableSpotResourcesAvailableSpotResource) GetInterruptionRate() *float32 {
	return s.InterruptionRate
}

func (s *DescribeSpotAdviceResponseBodyAvailableSpotZonesAvailableSpotZoneAvailableSpotResourcesAvailableSpotResource) SetAverageSpotDiscount(v int32) *DescribeSpotAdviceResponseBodyAvailableSpotZonesAvailableSpotZoneAvailableSpotResourcesAvailableSpotResource {
	s.AverageSpotDiscount = &v
	return s
}

func (s *DescribeSpotAdviceResponseBodyAvailableSpotZonesAvailableSpotZoneAvailableSpotResourcesAvailableSpotResource) SetInstanceType(v string) *DescribeSpotAdviceResponseBodyAvailableSpotZonesAvailableSpotZoneAvailableSpotResourcesAvailableSpotResource {
	s.InstanceType = &v
	return s
}

func (s *DescribeSpotAdviceResponseBodyAvailableSpotZonesAvailableSpotZoneAvailableSpotResourcesAvailableSpotResource) SetInterruptRateDesc(v string) *DescribeSpotAdviceResponseBodyAvailableSpotZonesAvailableSpotZoneAvailableSpotResourcesAvailableSpotResource {
	s.InterruptRateDesc = &v
	return s
}

func (s *DescribeSpotAdviceResponseBodyAvailableSpotZonesAvailableSpotZoneAvailableSpotResourcesAvailableSpotResource) SetInterruptionRate(v float32) *DescribeSpotAdviceResponseBodyAvailableSpotZonesAvailableSpotZoneAvailableSpotResourcesAvailableSpotResource {
	s.InterruptionRate = &v
	return s
}

func (s *DescribeSpotAdviceResponseBodyAvailableSpotZonesAvailableSpotZoneAvailableSpotResourcesAvailableSpotResource) Validate() error {
	return dara.Validate(s)
}
