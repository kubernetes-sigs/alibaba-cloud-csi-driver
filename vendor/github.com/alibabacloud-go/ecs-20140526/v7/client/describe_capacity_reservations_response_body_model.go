// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCapacityReservationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCapacityReservationSet(v *DescribeCapacityReservationsResponseBodyCapacityReservationSet) *DescribeCapacityReservationsResponseBody
	GetCapacityReservationSet() *DescribeCapacityReservationsResponseBodyCapacityReservationSet
	SetMaxResults(v int32) *DescribeCapacityReservationsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeCapacityReservationsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeCapacityReservationsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeCapacityReservationsResponseBody
	GetTotalCount() *int32
}

type DescribeCapacityReservationsResponseBody struct {
	CapacityReservationSet *DescribeCapacityReservationsResponseBodyCapacityReservationSet `json:"CapacityReservationSet,omitempty" xml:"CapacityReservationSet,omitempty" type:"Struct"`
	// The maximum number of entries returned per page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.If NextToken is empty, no next page exists.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCapacityReservationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCapacityReservationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCapacityReservationsResponseBody) GetCapacityReservationSet() *DescribeCapacityReservationsResponseBodyCapacityReservationSet {
	return s.CapacityReservationSet
}

func (s *DescribeCapacityReservationsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeCapacityReservationsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeCapacityReservationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCapacityReservationsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeCapacityReservationsResponseBody) SetCapacityReservationSet(v *DescribeCapacityReservationsResponseBodyCapacityReservationSet) *DescribeCapacityReservationsResponseBody {
	s.CapacityReservationSet = v
	return s
}

func (s *DescribeCapacityReservationsResponseBody) SetMaxResults(v int32) *DescribeCapacityReservationsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeCapacityReservationsResponseBody) SetNextToken(v string) *DescribeCapacityReservationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeCapacityReservationsResponseBody) SetRequestId(v string) *DescribeCapacityReservationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCapacityReservationsResponseBody) SetTotalCount(v int32) *DescribeCapacityReservationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCapacityReservationsResponseBody) Validate() error {
	if s.CapacityReservationSet != nil {
		if err := s.CapacityReservationSet.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCapacityReservationsResponseBodyCapacityReservationSet struct {
	CapacityReservationItem []*DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem `json:"CapacityReservationItem,omitempty" xml:"CapacityReservationItem,omitempty" type:"Repeated"`
}

func (s DescribeCapacityReservationsResponseBodyCapacityReservationSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeCapacityReservationsResponseBodyCapacityReservationSet) GoString() string {
	return s.String()
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSet) GetCapacityReservationItem() []*DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem {
	return s.CapacityReservationItem
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSet) SetCapacityReservationItem(v []*DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem) *DescribeCapacityReservationsResponseBodyCapacityReservationSet {
	s.CapacityReservationItem = v
	return s
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSet) Validate() error {
	if s.CapacityReservationItem != nil {
		for _, item := range s.CapacityReservationItem {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem struct {
	AllocatedResources              *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResources `json:"AllocatedResources,omitempty" xml:"AllocatedResources,omitempty" type:"Struct"`
	CapacityReservationOwnerId      *string                                                                                                  `json:"CapacityReservationOwnerId,omitempty" xml:"CapacityReservationOwnerId,omitempty"`
	Description                     *string                                                                                                  `json:"Description,omitempty" xml:"Description,omitempty"`
	EndTime                         *string                                                                                                  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EndTimeType                     *string                                                                                                  `json:"EndTimeType,omitempty" xml:"EndTimeType,omitempty"`
	InstanceChargeType              *string                                                                                                  `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	Platform                        *string                                                                                                  `json:"Platform,omitempty" xml:"Platform,omitempty"`
	PrivatePoolOptionsId            *string                                                                                                  `json:"PrivatePoolOptionsId,omitempty" xml:"PrivatePoolOptionsId,omitempty"`
	PrivatePoolOptionsMatchCriteria *string                                                                                                  `json:"PrivatePoolOptionsMatchCriteria,omitempty" xml:"PrivatePoolOptionsMatchCriteria,omitempty"`
	PrivatePoolOptionsName          *string                                                                                                  `json:"PrivatePoolOptionsName,omitempty" xml:"PrivatePoolOptionsName,omitempty"`
	RegionId                        *string                                                                                                  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ReservedInstanceId              *string                                                                                                  `json:"ReservedInstanceId,omitempty" xml:"ReservedInstanceId,omitempty"`
	ResourceGroupId                 *string                                                                                                  `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SavingPlanId                    *string                                                                                                  `json:"SavingPlanId,omitempty" xml:"SavingPlanId,omitempty"`
	StartTime                       *string                                                                                                  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StartTimeType                   *string                                                                                                  `json:"StartTimeType,omitempty" xml:"StartTimeType,omitempty"`
	Status                          *string                                                                                                  `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags                            *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemTags               `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	TimeSlot                        *string                                                                                                  `json:"TimeSlot,omitempty" xml:"TimeSlot,omitempty"`
}

func (s DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem) GoString() string {
	return s.String()
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem) GetAllocatedResources() *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResources {
	return s.AllocatedResources
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem) GetCapacityReservationOwnerId() *string {
	return s.CapacityReservationOwnerId
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem) GetDescription() *string {
	return s.Description
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem) GetEndTimeType() *string {
	return s.EndTimeType
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem) GetPlatform() *string {
	return s.Platform
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem) GetPrivatePoolOptionsId() *string {
	return s.PrivatePoolOptionsId
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem) GetPrivatePoolOptionsMatchCriteria() *string {
	return s.PrivatePoolOptionsMatchCriteria
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem) GetPrivatePoolOptionsName() *string {
	return s.PrivatePoolOptionsName
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem) GetReservedInstanceId() *string {
	return s.ReservedInstanceId
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem) GetSavingPlanId() *string {
	return s.SavingPlanId
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem) GetStartTimeType() *string {
	return s.StartTimeType
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem) GetStatus() *string {
	return s.Status
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem) GetTags() *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemTags {
	return s.Tags
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem) GetTimeSlot() *string {
	return s.TimeSlot
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem) SetAllocatedResources(v *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResources) *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem {
	s.AllocatedResources = v
	return s
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem) SetCapacityReservationOwnerId(v string) *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem {
	s.CapacityReservationOwnerId = &v
	return s
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem) SetDescription(v string) *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem {
	s.Description = &v
	return s
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem) SetEndTime(v string) *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem {
	s.EndTime = &v
	return s
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem) SetEndTimeType(v string) *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem {
	s.EndTimeType = &v
	return s
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem) SetInstanceChargeType(v string) *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem) SetPlatform(v string) *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem {
	s.Platform = &v
	return s
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem) SetPrivatePoolOptionsId(v string) *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem {
	s.PrivatePoolOptionsId = &v
	return s
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem) SetPrivatePoolOptionsMatchCriteria(v string) *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem {
	s.PrivatePoolOptionsMatchCriteria = &v
	return s
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem) SetPrivatePoolOptionsName(v string) *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem {
	s.PrivatePoolOptionsName = &v
	return s
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem) SetRegionId(v string) *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem {
	s.RegionId = &v
	return s
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem) SetReservedInstanceId(v string) *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem {
	s.ReservedInstanceId = &v
	return s
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem) SetResourceGroupId(v string) *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem) SetSavingPlanId(v string) *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem {
	s.SavingPlanId = &v
	return s
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem) SetStartTime(v string) *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem {
	s.StartTime = &v
	return s
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem) SetStartTimeType(v string) *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem {
	s.StartTimeType = &v
	return s
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem) SetStatus(v string) *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem {
	s.Status = &v
	return s
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem) SetTags(v *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemTags) *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem {
	s.Tags = v
	return s
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem) SetTimeSlot(v string) *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem {
	s.TimeSlot = &v
	return s
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItem) Validate() error {
	if s.AllocatedResources != nil {
		if err := s.AllocatedResources.Validate(); err != nil {
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

type DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResources struct {
	AllocatedResource []*DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResource `json:"AllocatedResource,omitempty" xml:"AllocatedResource,omitempty" type:"Repeated"`
}

func (s DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResources) GoString() string {
	return s.String()
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResources) GetAllocatedResource() []*DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResource {
	return s.AllocatedResource
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResources) SetAllocatedResource(v []*DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResource) *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResources {
	s.AllocatedResource = v
	return s
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResources) Validate() error {
	if s.AllocatedResource != nil {
		for _, item := range s.AllocatedResource {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResource struct {
	AvailableAmount           *int32                                                                                                                                             `json:"AvailableAmount,omitempty" xml:"AvailableAmount,omitempty"`
	CapacityReservationUsages *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResourceCapacityReservationUsages `json:"CapacityReservationUsages,omitempty" xml:"CapacityReservationUsages,omitempty" type:"Struct"`
	// example:
	//
	// 1
	FailedAmount *int32  `json:"FailedAmount,omitempty" xml:"FailedAmount,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// example:
	//
	// 1
	LockedAmount *int32  `json:"LockedAmount,omitempty" xml:"LockedAmount,omitempty"`
	TotalAmount  *int32  `json:"TotalAmount,omitempty" xml:"TotalAmount,omitempty"`
	UsedAmount   *int32  `json:"UsedAmount,omitempty" xml:"UsedAmount,omitempty"`
	ZoneId       *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResource) String() string {
	return dara.Prettify(s)
}

func (s DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResource) GoString() string {
	return s.String()
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResource) GetAvailableAmount() *int32 {
	return s.AvailableAmount
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResource) GetCapacityReservationUsages() *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResourceCapacityReservationUsages {
	return s.CapacityReservationUsages
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResource) GetFailedAmount() *int32 {
	return s.FailedAmount
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResource) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResource) GetLockedAmount() *int32 {
	return s.LockedAmount
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResource) GetTotalAmount() *int32 {
	return s.TotalAmount
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResource) GetUsedAmount() *int32 {
	return s.UsedAmount
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResource) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResource) SetAvailableAmount(v int32) *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResource {
	s.AvailableAmount = &v
	return s
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResource) SetCapacityReservationUsages(v *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResourceCapacityReservationUsages) *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResource {
	s.CapacityReservationUsages = v
	return s
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResource) SetFailedAmount(v int32) *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResource {
	s.FailedAmount = &v
	return s
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResource) SetInstanceType(v string) *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResource {
	s.InstanceType = &v
	return s
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResource) SetLockedAmount(v int32) *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResource {
	s.LockedAmount = &v
	return s
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResource) SetTotalAmount(v int32) *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResource {
	s.TotalAmount = &v
	return s
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResource) SetUsedAmount(v int32) *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResource {
	s.UsedAmount = &v
	return s
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResource) SetZoneId(v string) *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResource {
	s.ZoneId = &v
	return s
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResource) Validate() error {
	if s.CapacityReservationUsages != nil {
		if err := s.CapacityReservationUsages.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResourceCapacityReservationUsages struct {
	CapacityReservationUsage []*DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResourceCapacityReservationUsagesCapacityReservationUsage `json:"CapacityReservationUsage,omitempty" xml:"CapacityReservationUsage,omitempty" type:"Repeated"`
}

func (s DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResourceCapacityReservationUsages) String() string {
	return dara.Prettify(s)
}

func (s DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResourceCapacityReservationUsages) GoString() string {
	return s.String()
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResourceCapacityReservationUsages) GetCapacityReservationUsage() []*DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResourceCapacityReservationUsagesCapacityReservationUsage {
	return s.CapacityReservationUsage
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResourceCapacityReservationUsages) SetCapacityReservationUsage(v []*DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResourceCapacityReservationUsagesCapacityReservationUsage) *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResourceCapacityReservationUsages {
	s.CapacityReservationUsage = v
	return s
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResourceCapacityReservationUsages) Validate() error {
	if s.CapacityReservationUsage != nil {
		for _, item := range s.CapacityReservationUsage {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResourceCapacityReservationUsagesCapacityReservationUsage struct {
	AccountId   *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	UsedAmount  *int32  `json:"UsedAmount,omitempty" xml:"UsedAmount,omitempty"`
}

func (s DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResourceCapacityReservationUsagesCapacityReservationUsage) String() string {
	return dara.Prettify(s)
}

func (s DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResourceCapacityReservationUsagesCapacityReservationUsage) GoString() string {
	return s.String()
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResourceCapacityReservationUsagesCapacityReservationUsage) GetAccountId() *string {
	return s.AccountId
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResourceCapacityReservationUsagesCapacityReservationUsage) GetServiceName() *string {
	return s.ServiceName
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResourceCapacityReservationUsagesCapacityReservationUsage) GetUsedAmount() *int32 {
	return s.UsedAmount
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResourceCapacityReservationUsagesCapacityReservationUsage) SetAccountId(v string) *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResourceCapacityReservationUsagesCapacityReservationUsage {
	s.AccountId = &v
	return s
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResourceCapacityReservationUsagesCapacityReservationUsage) SetServiceName(v string) *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResourceCapacityReservationUsagesCapacityReservationUsage {
	s.ServiceName = &v
	return s
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResourceCapacityReservationUsagesCapacityReservationUsage) SetUsedAmount(v int32) *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResourceCapacityReservationUsagesCapacityReservationUsage {
	s.UsedAmount = &v
	return s
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResourceCapacityReservationUsagesCapacityReservationUsage) Validate() error {
	return dara.Validate(s)
}

type DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemTags struct {
	Tag []*DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemTags) GoString() string {
	return s.String()
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemTags) GetTag() []*DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemTagsTag {
	return s.Tag
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemTags) SetTag(v []*DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemTagsTag) *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemTags {
	s.Tag = v
	return s
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemTags) Validate() error {
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

type DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemTagsTag struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemTagsTag) SetTagKey(v string) *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemTagsTag {
	s.TagKey = &v
	return s
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemTagsTag) SetTagValue(v string) *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemTagsTag {
	s.TagValue = &v
	return s
}

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemTagsTag) Validate() error {
	return dara.Validate(s)
}
