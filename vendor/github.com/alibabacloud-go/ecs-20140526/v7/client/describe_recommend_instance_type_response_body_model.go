// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecommendInstanceTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeRecommendInstanceTypeResponseBodyData) *DescribeRecommendInstanceTypeResponseBody
	GetData() *DescribeRecommendInstanceTypeResponseBodyData
	SetRequestId(v string) *DescribeRecommendInstanceTypeResponseBody
	GetRequestId() *string
}

type DescribeRecommendInstanceTypeResponseBody struct {
	// The details of the recommended instance types.
	Data *DescribeRecommendInstanceTypeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRecommendInstanceTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecommendInstanceTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRecommendInstanceTypeResponseBody) GetData() *DescribeRecommendInstanceTypeResponseBodyData {
	return s.Data
}

func (s *DescribeRecommendInstanceTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRecommendInstanceTypeResponseBody) SetData(v *DescribeRecommendInstanceTypeResponseBodyData) *DescribeRecommendInstanceTypeResponseBody {
	s.Data = v
	return s
}

func (s *DescribeRecommendInstanceTypeResponseBody) SetRequestId(v string) *DescribeRecommendInstanceTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRecommendInstanceTypeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRecommendInstanceTypeResponseBodyData struct {
	RecommendInstanceType []*DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceType `json:"RecommendInstanceType,omitempty" xml:"RecommendInstanceType,omitempty" type:"Repeated"`
}

func (s DescribeRecommendInstanceTypeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecommendInstanceTypeResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeRecommendInstanceTypeResponseBodyData) GetRecommendInstanceType() []*DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceType {
	return s.RecommendInstanceType
}

func (s *DescribeRecommendInstanceTypeResponseBodyData) SetRecommendInstanceType(v []*DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceType) *DescribeRecommendInstanceTypeResponseBodyData {
	s.RecommendInstanceType = v
	return s
}

func (s *DescribeRecommendInstanceTypeResponseBodyData) Validate() error {
	if s.RecommendInstanceType != nil {
		for _, item := range s.RecommendInstanceType {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceType struct {
	// The commodity code of the instance type.
	//
	// example:
	//
	// ecs
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The billing method of the instances.
	//
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The details of the instance type.
	InstanceType *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceTypeInstanceType `json:"InstanceType,omitempty" xml:"InstanceType,omitempty" type:"Struct"`
	// The network type of the ECS instances.
	//
	// example:
	//
	// vpc
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The priority based on which the system sorts the instance types.
	//
	// example:
	//
	// 2
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ID of the region in which the instance type is available.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The scenario in which the instance type is recommended.
	//
	// example:
	//
	// CREATE
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// The bidding policy for the spot instances.
	//
	// example:
	//
	// NoSpot
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// The ID of the zone in which the instance type is available.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// The details of the zones in which the instance type is available.
	Zones *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceTypeZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Struct"`
}

func (s DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceType) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceType) GoString() string {
	return s.String()
}

func (s *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceType) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceType) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceType) GetInstanceType() *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceTypeInstanceType {
	return s.InstanceType
}

func (s *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceType) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceType) GetPriority() *int32 {
	return s.Priority
}

func (s *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceType) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceType) GetScene() *string {
	return s.Scene
}

func (s *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceType) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceType) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceType) GetZones() *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceTypeZones {
	return s.Zones
}

func (s *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceType) SetCommodityCode(v string) *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceType {
	s.CommodityCode = &v
	return s
}

func (s *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceType) SetInstanceChargeType(v string) *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceType {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceType) SetInstanceType(v *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceTypeInstanceType) *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceType {
	s.InstanceType = v
	return s
}

func (s *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceType) SetNetworkType(v string) *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceType {
	s.NetworkType = &v
	return s
}

func (s *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceType) SetPriority(v int32) *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceType {
	s.Priority = &v
	return s
}

func (s *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceType) SetRegionId(v string) *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceType {
	s.RegionId = &v
	return s
}

func (s *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceType) SetScene(v string) *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceType {
	s.Scene = &v
	return s
}

func (s *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceType) SetSpotStrategy(v string) *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceType {
	s.SpotStrategy = &v
	return s
}

func (s *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceType) SetZoneId(v string) *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceType {
	s.ZoneId = &v
	return s
}

func (s *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceType) SetZones(v *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceTypeZones) *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceType {
	s.Zones = v
	return s
}

func (s *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceType) Validate() error {
	if s.InstanceType != nil {
		if err := s.InstanceType.Validate(); err != nil {
			return err
		}
	}
	if s.Zones != nil {
		if err := s.Zones.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceTypeInstanceType struct {
	// The number of vCPUs of the instance type.
	//
	// example:
	//
	// 1
	Cores *int32 `json:"Cores,omitempty" xml:"Cores,omitempty"`
	// The generation of the instance family.
	//
	// example:
	//
	// ecs-4
	Generation *string `json:"Generation,omitempty" xml:"Generation,omitempty"`
	// The name of the instance type.
	//
	// example:
	//
	// ecs.hfg6.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The instance family.
	//
	// example:
	//
	// ecs.hfg6
	InstanceTypeFamily *string `json:"InstanceTypeFamily,omitempty" xml:"InstanceTypeFamily,omitempty"`
	// The memory size of the instance type. Unit: MB.
	//
	// example:
	//
	// 8192
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// Indicates whether the instance type supports I/O optimization.
	//
	// example:
	//
	// optimized
	SupportIoOptimized *string `json:"SupportIoOptimized,omitempty" xml:"SupportIoOptimized,omitempty"`
}

func (s DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceTypeInstanceType) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceTypeInstanceType) GoString() string {
	return s.String()
}

func (s *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceTypeInstanceType) GetCores() *int32 {
	return s.Cores
}

func (s *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceTypeInstanceType) GetGeneration() *string {
	return s.Generation
}

func (s *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceTypeInstanceType) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceTypeInstanceType) GetInstanceTypeFamily() *string {
	return s.InstanceTypeFamily
}

func (s *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceTypeInstanceType) GetMemory() *int32 {
	return s.Memory
}

func (s *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceTypeInstanceType) GetSupportIoOptimized() *string {
	return s.SupportIoOptimized
}

func (s *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceTypeInstanceType) SetCores(v int32) *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceTypeInstanceType {
	s.Cores = &v
	return s
}

func (s *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceTypeInstanceType) SetGeneration(v string) *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceTypeInstanceType {
	s.Generation = &v
	return s
}

func (s *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceTypeInstanceType) SetInstanceType(v string) *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceTypeInstanceType {
	s.InstanceType = &v
	return s
}

func (s *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceTypeInstanceType) SetInstanceTypeFamily(v string) *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceTypeInstanceType {
	s.InstanceTypeFamily = &v
	return s
}

func (s *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceTypeInstanceType) SetMemory(v int32) *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceTypeInstanceType {
	s.Memory = &v
	return s
}

func (s *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceTypeInstanceType) SetSupportIoOptimized(v string) *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceTypeInstanceType {
	s.SupportIoOptimized = &v
	return s
}

func (s *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceTypeInstanceType) Validate() error {
	return dara.Validate(s)
}

type DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceTypeZones struct {
	Zone []*DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceTypeZonesZone `json:"zone,omitempty" xml:"zone,omitempty" type:"Repeated"`
}

func (s DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceTypeZones) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceTypeZones) GoString() string {
	return s.String()
}

func (s *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceTypeZones) GetZone() []*DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceTypeZonesZone {
	return s.Zone
}

func (s *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceTypeZones) SetZone(v []*DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceTypeZonesZone) *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceTypeZones {
	s.Zone = v
	return s
}

func (s *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceTypeZones) Validate() error {
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

type DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceTypeZonesZone struct {
	// The details of the network types of the instance type.
	NetworkTypes *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceTypeZonesZoneNetworkTypes `json:"NetworkTypes,omitempty" xml:"NetworkTypes,omitempty" type:"Struct"`
	// The ID of the zone in which the instance type is available.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneNo *string `json:"ZoneNo,omitempty" xml:"ZoneNo,omitempty"`
}

func (s DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceTypeZonesZone) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceTypeZonesZone) GoString() string {
	return s.String()
}

func (s *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceTypeZonesZone) GetNetworkTypes() *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceTypeZonesZoneNetworkTypes {
	return s.NetworkTypes
}

func (s *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceTypeZonesZone) GetZoneNo() *string {
	return s.ZoneNo
}

func (s *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceTypeZonesZone) SetNetworkTypes(v *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceTypeZonesZoneNetworkTypes) *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceTypeZonesZone {
	s.NetworkTypes = v
	return s
}

func (s *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceTypeZonesZone) SetZoneNo(v string) *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceTypeZonesZone {
	s.ZoneNo = &v
	return s
}

func (s *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceTypeZonesZone) Validate() error {
	if s.NetworkTypes != nil {
		if err := s.NetworkTypes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceTypeZonesZoneNetworkTypes struct {
	NetworkType []*string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty" type:"Repeated"`
}

func (s DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceTypeZonesZoneNetworkTypes) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceTypeZonesZoneNetworkTypes) GoString() string {
	return s.String()
}

func (s *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceTypeZonesZoneNetworkTypes) GetNetworkType() []*string {
	return s.NetworkType
}

func (s *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceTypeZonesZoneNetworkTypes) SetNetworkType(v []*string) *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceTypeZonesZoneNetworkTypes {
	s.NetworkType = v
	return s
}

func (s *DescribeRecommendInstanceTypeResponseBodyDataRecommendInstanceTypeZonesZoneNetworkTypes) Validate() error {
	return dara.Validate(s)
}
