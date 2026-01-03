// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeploymentSetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeploymentSets(v *DescribeDeploymentSetsResponseBodyDeploymentSets) *DescribeDeploymentSetsResponseBody
	GetDeploymentSets() *DescribeDeploymentSetsResponseBodyDeploymentSets
	SetPageNumber(v int32) *DescribeDeploymentSetsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDeploymentSetsResponseBody
	GetPageSize() *int32
	SetRegionId(v string) *DescribeDeploymentSetsResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeDeploymentSetsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeDeploymentSetsResponseBody
	GetTotalCount() *int32
}

type DescribeDeploymentSetsResponseBody struct {
	// Details about the deployment sets.
	DeploymentSets *DescribeDeploymentSetsResponseBodyDeploymentSets `json:"DeploymentSets,omitempty" xml:"DeploymentSets,omitempty" type:"Struct"`
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
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region.
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
	// The total number of queried deployment sets.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDeploymentSetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeploymentSetsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeploymentSetsResponseBody) GetDeploymentSets() *DescribeDeploymentSetsResponseBodyDeploymentSets {
	return s.DeploymentSets
}

func (s *DescribeDeploymentSetsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDeploymentSetsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDeploymentSetsResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDeploymentSetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDeploymentSetsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDeploymentSetsResponseBody) SetDeploymentSets(v *DescribeDeploymentSetsResponseBodyDeploymentSets) *DescribeDeploymentSetsResponseBody {
	s.DeploymentSets = v
	return s
}

func (s *DescribeDeploymentSetsResponseBody) SetPageNumber(v int32) *DescribeDeploymentSetsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDeploymentSetsResponseBody) SetPageSize(v int32) *DescribeDeploymentSetsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDeploymentSetsResponseBody) SetRegionId(v string) *DescribeDeploymentSetsResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeDeploymentSetsResponseBody) SetRequestId(v string) *DescribeDeploymentSetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDeploymentSetsResponseBody) SetTotalCount(v int32) *DescribeDeploymentSetsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDeploymentSetsResponseBody) Validate() error {
	if s.DeploymentSets != nil {
		if err := s.DeploymentSets.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDeploymentSetsResponseBodyDeploymentSets struct {
	DeploymentSet []*DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSet `json:"DeploymentSet,omitempty" xml:"DeploymentSet,omitempty" type:"Repeated"`
}

func (s DescribeDeploymentSetsResponseBodyDeploymentSets) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeploymentSetsResponseBodyDeploymentSets) GoString() string {
	return s.String()
}

func (s *DescribeDeploymentSetsResponseBodyDeploymentSets) GetDeploymentSet() []*DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSet {
	return s.DeploymentSet
}

func (s *DescribeDeploymentSetsResponseBodyDeploymentSets) SetDeploymentSet(v []*DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) *DescribeDeploymentSetsResponseBodyDeploymentSets {
	s.DeploymentSet = v
	return s
}

func (s *DescribeDeploymentSetsResponseBodyDeploymentSets) Validate() error {
	if s.DeploymentSet != nil {
		for _, item := range s.DeploymentSet {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSet struct {
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// Details of the capacities of the deployment set. This parameter is valid only when the deployment set contains ECS instances. The value contains information about the capacities of the deployment set in different zones.
	Capacities *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSetCapacities `json:"Capacities,omitempty" xml:"Capacities,omitempty" type:"Struct"`
	// The time when the deployment set was created.
	//
	// example:
	//
	// 2021-12-07T06:01:46Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the deployment set.
	//
	// example:
	//
	// testDeploymentSetDescription
	DeploymentSetDescription *string `json:"DeploymentSetDescription,omitempty" xml:"DeploymentSetDescription,omitempty"`
	// The ID of the deployment set.
	//
	// example:
	//
	// ds-bp67acfmxazb4ph****
	DeploymentSetId *string `json:"DeploymentSetId,omitempty" xml:"DeploymentSetId,omitempty"`
	// The name of the deployment set.
	//
	// example:
	//
	// testDeploymentSetName
	DeploymentSetName *string `json:"DeploymentSetName,omitempty" xml:"DeploymentSetName,omitempty"`
	// The deployment strategy. The return value of this parameter is the value of the `Strategy` request parameter.
	//
	// example:
	//
	// Availability
	DeploymentStrategy *string `json:"DeploymentStrategy,omitempty" xml:"DeploymentStrategy,omitempty"`
	// The deployment domain.
	//
	// example:
	//
	// default
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The deployment granularity.
	//
	// example:
	//
	// host
	Granularity *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	// The number of deployment set groups in the deployment set.
	//
	// >  This parameter is valid only when the Strategy request parameter is set to AvailabilityGroup.
	//
	// example:
	//
	// 3
	GroupCount *int32 `json:"GroupCount,omitempty" xml:"GroupCount,omitempty"`
	// The number of instances in the deployment set.
	//
	// example:
	//
	// 3
	InstanceAmount *int32 `json:"InstanceAmount,omitempty" xml:"InstanceAmount,omitempty"`
	// The IDs of the Elastic Compute Service (ECS) instances in the deployment set.
	InstanceIds *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSetInstanceIds `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Struct"`
	// The deployment strategy.
	//
	// example:
	//
	// LooseDispersion
	Strategy *string `json:"Strategy,omitempty" xml:"Strategy,omitempty"`
}

func (s DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) GoString() string {
	return s.String()
}

func (s *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) GetAccountId() *int64 {
	return s.AccountId
}

func (s *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) GetCapacities() *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSetCapacities {
	return s.Capacities
}

func (s *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) GetDeploymentSetDescription() *string {
	return s.DeploymentSetDescription
}

func (s *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) GetDeploymentSetId() *string {
	return s.DeploymentSetId
}

func (s *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) GetDeploymentSetName() *string {
	return s.DeploymentSetName
}

func (s *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) GetDeploymentStrategy() *string {
	return s.DeploymentStrategy
}

func (s *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) GetGranularity() *string {
	return s.Granularity
}

func (s *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) GetGroupCount() *int32 {
	return s.GroupCount
}

func (s *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) GetInstanceAmount() *int32 {
	return s.InstanceAmount
}

func (s *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) GetInstanceIds() *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSetInstanceIds {
	return s.InstanceIds
}

func (s *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) GetStrategy() *string {
	return s.Strategy
}

func (s *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) SetAccountId(v int64) *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSet {
	s.AccountId = &v
	return s
}

func (s *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) SetCapacities(v *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSetCapacities) *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSet {
	s.Capacities = v
	return s
}

func (s *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) SetCreationTime(v string) *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSet {
	s.CreationTime = &v
	return s
}

func (s *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) SetDeploymentSetDescription(v string) *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSet {
	s.DeploymentSetDescription = &v
	return s
}

func (s *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) SetDeploymentSetId(v string) *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSet {
	s.DeploymentSetId = &v
	return s
}

func (s *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) SetDeploymentSetName(v string) *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSet {
	s.DeploymentSetName = &v
	return s
}

func (s *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) SetDeploymentStrategy(v string) *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSet {
	s.DeploymentStrategy = &v
	return s
}

func (s *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) SetDomain(v string) *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSet {
	s.Domain = &v
	return s
}

func (s *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) SetGranularity(v string) *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSet {
	s.Granularity = &v
	return s
}

func (s *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) SetGroupCount(v int32) *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSet {
	s.GroupCount = &v
	return s
}

func (s *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) SetInstanceAmount(v int32) *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSet {
	s.InstanceAmount = &v
	return s
}

func (s *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) SetInstanceIds(v *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSetInstanceIds) *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSet {
	s.InstanceIds = v
	return s
}

func (s *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) SetStrategy(v string) *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSet {
	s.Strategy = &v
	return s
}

func (s *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) Validate() error {
	if s.Capacities != nil {
		if err := s.Capacities.Validate(); err != nil {
			return err
		}
	}
	if s.InstanceIds != nil {
		if err := s.InstanceIds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSetCapacities struct {
	Capacity []*DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSetCapacitiesCapacity `json:"Capacity,omitempty" xml:"Capacity,omitempty" type:"Repeated"`
}

func (s DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSetCapacities) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSetCapacities) GoString() string {
	return s.String()
}

func (s *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSetCapacities) GetCapacity() []*DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSetCapacitiesCapacity {
	return s.Capacity
}

func (s *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSetCapacities) SetCapacity(v []*DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSetCapacitiesCapacity) *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSetCapacities {
	s.Capacity = v
	return s
}

func (s *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSetCapacities) Validate() error {
	if s.Capacity != nil {
		for _, item := range s.Capacity {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSetCapacitiesCapacity struct {
	// The number of ECS instances that can be added to the deployment set within the zone.
	//
	// example:
	//
	// 18
	AvailableAmount *int32 `json:"AvailableAmount,omitempty" xml:"AvailableAmount,omitempty"`
	// The number of ECS instances that reside in the zone in the deployment set.
	//
	// example:
	//
	// 2
	UsedAmount *int32 `json:"UsedAmount,omitempty" xml:"UsedAmount,omitempty"`
	// The ID of the zone. Only the zone IDs of existing ECS instances in the deployment set are returned.
	//
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSetCapacitiesCapacity) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSetCapacitiesCapacity) GoString() string {
	return s.String()
}

func (s *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSetCapacitiesCapacity) GetAvailableAmount() *int32 {
	return s.AvailableAmount
}

func (s *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSetCapacitiesCapacity) GetUsedAmount() *int32 {
	return s.UsedAmount
}

func (s *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSetCapacitiesCapacity) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSetCapacitiesCapacity) SetAvailableAmount(v int32) *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSetCapacitiesCapacity {
	s.AvailableAmount = &v
	return s
}

func (s *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSetCapacitiesCapacity) SetUsedAmount(v int32) *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSetCapacitiesCapacity {
	s.UsedAmount = &v
	return s
}

func (s *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSetCapacitiesCapacity) SetZoneId(v string) *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSetCapacitiesCapacity {
	s.ZoneId = &v
	return s
}

func (s *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSetCapacitiesCapacity) Validate() error {
	return dara.Validate(s)
}

type DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSetInstanceIds struct {
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
}

func (s DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSetInstanceIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSetInstanceIds) GoString() string {
	return s.String()
}

func (s *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSetInstanceIds) GetInstanceId() []*string {
	return s.InstanceId
}

func (s *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSetInstanceIds) SetInstanceId(v []*string) *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSetInstanceIds {
	s.InstanceId = v
	return s
}

func (s *DescribeDeploymentSetsResponseBodyDeploymentSetsDeploymentSetInstanceIds) Validate() error {
	return dara.Validate(s)
}
