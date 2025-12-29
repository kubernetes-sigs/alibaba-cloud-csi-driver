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
	// Details about spot instances in the zones of the specified region.
	//
	// >  The return values are sorted based on the historical percentages of average spot instance prices relative to pay-as-you-go instance prices for instance types.
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
	// Details about spot instances in the previous 30 days, including the release rate of spot instances and percentages of average spot instance prices relative to pay-as-you-go instance prices.
	AvailableSpotResources *DescribeSpotAdviceResponseBodyAvailableSpotZonesAvailableSpotZoneAvailableSpotResources `json:"AvailableSpotResources,omitempty" xml:"AvailableSpotResources,omitempty" type:"Struct"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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
	// The percentage of the average spot instance price relative to the pay-as-you-go instance price in the previous 30 days. Unit: %. Valid values: 1 to 100.
	//
	// You can calculate the average spot instance price based on the return value. For example, if the pay-as-you-go instance price is 1 and the return value of this parameter is 20, the average spot instance price in the previous 30 days is 0.2.
	//
	// example:
	//
	// 20
	AverageSpotDiscount *int32 `json:"AverageSpotDiscount,omitempty" xml:"AverageSpotDiscount,omitempty"`
	// The instance type.
	//
	// example:
	//
	// ecs.c5.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The release rate range of spot instances in the previous 30 days, which corresponds to the `InterruptionRate` value. Valid values:
	//
	// 	- 0-3%
	//
	// 	- 3-5%
	//
	// 	- 5-10%
	//
	// 	- 10-100%
	//
	// example:
	//
	// 0-3%
	InterruptRateDesc *string `json:"InterruptRateDesc,omitempty" xml:"InterruptRateDesc,omitempty"`
	// The average release rate of spot instances in the previous 30 days. Unit: %.
	//
	// example:
	//
	// 0
	InterruptionRate *float32 `json:"InterruptionRate,omitempty" xml:"InterruptionRate,omitempty"`
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
