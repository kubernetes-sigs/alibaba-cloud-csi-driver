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
	// Details of the capacity reservations.
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
	// Details of the allocated resources.
	AllocatedResources *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResources `json:"AllocatedResources,omitempty" xml:"AllocatedResources,omitempty" type:"Struct"`
	// The ID of the capacity reservation owner.
	//
	// example:
	//
	// 100************7
	CapacityReservationOwnerId *string `json:"CapacityReservationOwnerId,omitempty" xml:"CapacityReservationOwnerId,omitempty"`
	// The description of the capacity reservation.
	//
	// example:
	//
	// This is description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the capacity reservation expires.
	//
	// example:
	//
	// 2021-02-19T03:02Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The release mode of the capacity reservation. Valid values:
	//
	// 	- Limited: The capacity reservation is automatically released at a specified time.
	//
	// 	- Unlimited: The capacity reservation is manually released. You can release the capacity reservation anytime.
	//
	// example:
	//
	// Unlimited
	EndTimeType *string `json:"EndTimeType,omitempty" xml:"EndTimeType,omitempty"`
	// The billing method of the instances created by using the capacity reservation. Valid values:
	//
	// 	- PostPaid: pay-as-you-go.
	//
	// 	- PrePaid: subscription.
	//
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The operating system type of the instances created by using the capacity reservation. Valid values:
	//
	// 	- windows
	//
	// 	- linux
	//
	// example:
	//
	// linux
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The ID of the capacity reservation.
	//
	// example:
	//
	// crp-bp1gubrkqutenqdd****
	PrivatePoolOptionsId *string `json:"PrivatePoolOptionsId,omitempty" xml:"PrivatePoolOptionsId,omitempty"`
	// The type of the private pool generated after the capacity reservation takes effect. Valid values:
	//
	// 	- Open: open private pool. If you use the capacity reservation to create Elastic Compute Service (ECS) instances, the open private pool that is associated with the capacity reservation is automatically matched. If no capacity is available in the open private pool, resources in the public pool are automatically used to create the instances.
	//
	// 	- Target: targeted private pool. If you use the capacity reservation to create ECS instances, the targeted private pool that is associated with the capacity reservation is automatically matched. If no capacity is available in the private pool, the instances fail to be created.
	//
	// example:
	//
	// Open
	PrivatePoolOptionsMatchCriteria *string `json:"PrivatePoolOptionsMatchCriteria,omitempty" xml:"PrivatePoolOptionsMatchCriteria,omitempty"`
	// The name of the capacity reservation.
	//
	// example:
	//
	// crpTestName
	PrivatePoolOptionsName *string `json:"PrivatePoolOptionsName,omitempty" xml:"PrivatePoolOptionsName,omitempty"`
	// The region ID of the capacity reservation.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the reserved instance used with the capacity reservation.
	//
	// example:
	//
	// ri-bpzhex2ulpzf53****
	ReservedInstanceId *string `json:"ReservedInstanceId,omitempty" xml:"ReservedInstanceId,omitempty"`
	// The ID of the resource group to which the capacity reservation belongs.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the savings plan used with the capacity reservation.
	//
	// example:
	//
	// spn-c29b5e18pJMT****
	SavingPlanId *string `json:"SavingPlanId,omitempty" xml:"SavingPlanId,omitempty"`
	// The time when the capacity reservation takes effect.
	//
	// example:
	//
	// 2021-02-19T02:01Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The mode in which the capacity reservation takes effect. Valid values:
	//
	// 	- Now: The capacity reservation takes effect immediately after it is created.
	//
	// 	- Later: The capacity reservation takes effect at a specified time.
	//
	// example:
	//
	// Now
	StartTimeType *string `json:"StartTimeType,omitempty" xml:"StartTimeType,omitempty"`
	// The status of the capacity reservation. Valid values:
	//
	// 	- Pending: The capacity reservation is being initialized.
	//
	// 	- Preparing: The capacity reservation is being prepared.
	//
	// 	- Prepared: The capacity reservation is to take effect.
	//
	// 	- Active: The capacity reservation is in effect.
	//
	// 	- Released: The capacity reservation has been released manually or automatically when it expired.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags that are added to the capacity reservation.
	Tags *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// >  This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// null
	TimeSlot *string `json:"TimeSlot,omitempty" xml:"TimeSlot,omitempty"`
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
	// The number of available instances.
	//
	// example:
	//
	// 2
	AvailableAmount *int32 `json:"AvailableAmount,omitempty" xml:"AvailableAmount,omitempty"`
	// Details of instance usage.
	CapacityReservationUsages *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResourceCapacityReservationUsages `json:"CapacityReservationUsages,omitempty" xml:"CapacityReservationUsages,omitempty" type:"Struct"`
	// The instance type of the instances.
	//
	// example:
	//
	// ecs.c6.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The total number of instances for which the capacity of an instance type is reserved.
	//
	// example:
	//
	// 2
	TotalAmount *int32 `json:"TotalAmount,omitempty" xml:"TotalAmount,omitempty"`
	// The number of instances that have used the capacity reservation.
	//
	// example:
	//
	// 2
	UsedAmount *int32 `json:"UsedAmount,omitempty" xml:"UsedAmount,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
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

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResource) GetInstanceType() *string {
	return s.InstanceType
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

func (s *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResource) SetInstanceType(v string) *DescribeCapacityReservationsResponseBodyCapacityReservationSetCapacityReservationItemAllocatedResourcesAllocatedResource {
	s.InstanceType = &v
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
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 105909559088****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The name of the Alibaba Cloud service.
	//
	// example:
	//
	// maxcompute.aliyuncs.com
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The number of instances that are used by the Alibaba Cloud account or service.
	//
	// example:
	//
	// 20
	UsedAmount *int32 `json:"UsedAmount,omitempty" xml:"UsedAmount,omitempty"`
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
	// The tag key.
	//
	// example:
	//
	// TestKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// TestValue
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
