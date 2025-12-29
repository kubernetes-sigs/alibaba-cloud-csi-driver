// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeElasticityAssurancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetElasticityAssuranceSet(v *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSet) *DescribeElasticityAssurancesResponseBody
	GetElasticityAssuranceSet() *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSet
	SetMaxResults(v int32) *DescribeElasticityAssurancesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeElasticityAssurancesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeElasticityAssurancesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeElasticityAssurancesResponseBody
	GetTotalCount() *int32
}

type DescribeElasticityAssurancesResponseBody struct {
	// Details about the elasticity assurances.
	ElasticityAssuranceSet *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSet `json:"ElasticityAssuranceSet,omitempty" xml:"ElasticityAssuranceSet,omitempty" type:"Struct"`
	// The maximum number of entries returned per page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeElasticityAssurancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticityAssurancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeElasticityAssurancesResponseBody) GetElasticityAssuranceSet() *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSet {
	return s.ElasticityAssuranceSet
}

func (s *DescribeElasticityAssurancesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeElasticityAssurancesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeElasticityAssurancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeElasticityAssurancesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeElasticityAssurancesResponseBody) SetElasticityAssuranceSet(v *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSet) *DescribeElasticityAssurancesResponseBody {
	s.ElasticityAssuranceSet = v
	return s
}

func (s *DescribeElasticityAssurancesResponseBody) SetMaxResults(v int32) *DescribeElasticityAssurancesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeElasticityAssurancesResponseBody) SetNextToken(v string) *DescribeElasticityAssurancesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeElasticityAssurancesResponseBody) SetRequestId(v string) *DescribeElasticityAssurancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeElasticityAssurancesResponseBody) SetTotalCount(v int32) *DescribeElasticityAssurancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeElasticityAssurancesResponseBody) Validate() error {
	if s.ElasticityAssuranceSet != nil {
		if err := s.ElasticityAssuranceSet.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeElasticityAssurancesResponseBodyElasticityAssuranceSet struct {
	ElasticityAssuranceItem []*DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem `json:"ElasticityAssuranceItem,omitempty" xml:"ElasticityAssuranceItem,omitempty" type:"Repeated"`
}

func (s DescribeElasticityAssurancesResponseBodyElasticityAssuranceSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticityAssurancesResponseBodyElasticityAssuranceSet) GoString() string {
	return s.String()
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSet) GetElasticityAssuranceItem() []*DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem {
	return s.ElasticityAssuranceItem
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSet) SetElasticityAssuranceItem(v []*DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem) *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSet {
	s.ElasticityAssuranceItem = v
	return s
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSet) Validate() error {
	if s.ElasticityAssuranceItem != nil {
		for _, item := range s.ElasticityAssuranceItem {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem struct {
	// Details of the allocated resources.
	AllocatedResources *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResources `json:"AllocatedResources,omitempty" xml:"AllocatedResources,omitempty" type:"Struct"`
	// The description of the elasticity assurance.
	//
	// example:
	//
	// This is description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// e114477596854834
	ElasticityAssuranceOwnerId *string `json:"ElasticityAssuranceOwnerId,omitempty" xml:"ElasticityAssuranceOwnerId,omitempty"`
	// The time when the elasticity assurance expires.
	//
	// example:
	//
	// 2021-12-03T16:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The billing method of the instance. The value can be only PostPaid. Only pay-as-you-go instances can be created by using elasticity assurances.
	//
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// > This parameter is not publicly available.
	//
	// example:
	//
	// null
	LatestStartTime *string `json:"LatestStartTime,omitempty" xml:"LatestStartTime,omitempty"`
	// The type of the elasticity assurance. Valid values:
	//
	// 	- ElasticityAssurance: the general-purpose elasticity assurance.
	//
	// 	- TimeDivisionElasticityAssurance: the time-segmented assurance of the elasticity assurance.
	//
	// example:
	//
	// ElasticityAssurance
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// The ID of the elasticity assurance.
	//
	// example:
	//
	// eap-bp1ir35kqtn8ywci****
	PrivatePoolOptionsId *string `json:"PrivatePoolOptionsId,omitempty" xml:"PrivatePoolOptionsId,omitempty"`
	// The type of the private pool associated with the elasticity assurance. Valid values:
	//
	// 	- Open: open private pool
	//
	// 	- Target: specific private pool
	//
	// example:
	//
	// Target
	PrivatePoolOptionsMatchCriteria *string `json:"PrivatePoolOptionsMatchCriteria,omitempty" xml:"PrivatePoolOptionsMatchCriteria,omitempty"`
	// The name of the elasticity assurance.
	//
	// example:
	//
	// eapTestName
	PrivatePoolOptionsName *string `json:"PrivatePoolOptionsName,omitempty" xml:"PrivatePoolOptionsName,omitempty"`
	// The recurrence rules of the time-segmented assurances.
	RecurrenceRules *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemRecurrenceRules `json:"RecurrenceRules,omitempty" xml:"RecurrenceRules,omitempty" type:"Struct"`
	// The region ID of the elasticity assurance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The time when the elasticity assurance takes effect.
	//
	// example:
	//
	// 2020-12-03T05:25Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Indicates when the elasticity assurance takes effect. Valid values:
	//
	// 	- Now: The elasticity assurance takes effect immediately after it is created.
	//
	// 	- Later: The elasticity assurance takes effect at a specified time.
	//
	// example:
	//
	// Now
	StartTimeType *string `json:"StartTimeType,omitempty" xml:"StartTimeType,omitempty"`
	// The status of the elasticity assurance. Valid values:
	//
	// 	- Preparing
	//
	// 	- Prepared
	//
	// 	- Active
	//
	// 	- Released
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags of the elasticity assurance.
	Tags *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The total number of times that the elasticity assurance is applied.
	//
	// example:
	//
	// Unlimited
	TotalAssuranceTimes *string `json:"TotalAssuranceTimes,omitempty" xml:"TotalAssuranceTimes,omitempty"`
	// > This parameter is not publicly available.
	//
	// example:
	//
	// 0
	UsedAssuranceTimes *int32 `json:"UsedAssuranceTimes,omitempty" xml:"UsedAssuranceTimes,omitempty"`
}

func (s DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem) GoString() string {
	return s.String()
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem) GetAllocatedResources() *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResources {
	return s.AllocatedResources
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem) GetDescription() *string {
	return s.Description
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem) GetElasticityAssuranceOwnerId() *string {
	return s.ElasticityAssuranceOwnerId
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem) GetLatestStartTime() *string {
	return s.LatestStartTime
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem) GetPackageType() *string {
	return s.PackageType
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem) GetPrivatePoolOptionsId() *string {
	return s.PrivatePoolOptionsId
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem) GetPrivatePoolOptionsMatchCriteria() *string {
	return s.PrivatePoolOptionsMatchCriteria
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem) GetPrivatePoolOptionsName() *string {
	return s.PrivatePoolOptionsName
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem) GetRecurrenceRules() *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemRecurrenceRules {
	return s.RecurrenceRules
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem) GetStartTimeType() *string {
	return s.StartTimeType
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem) GetStatus() *string {
	return s.Status
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem) GetTags() *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemTags {
	return s.Tags
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem) GetTotalAssuranceTimes() *string {
	return s.TotalAssuranceTimes
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem) GetUsedAssuranceTimes() *int32 {
	return s.UsedAssuranceTimes
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem) SetAllocatedResources(v *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResources) *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem {
	s.AllocatedResources = v
	return s
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem) SetDescription(v string) *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem {
	s.Description = &v
	return s
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem) SetElasticityAssuranceOwnerId(v string) *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem {
	s.ElasticityAssuranceOwnerId = &v
	return s
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem) SetEndTime(v string) *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem {
	s.EndTime = &v
	return s
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem) SetInstanceChargeType(v string) *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem) SetLatestStartTime(v string) *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem {
	s.LatestStartTime = &v
	return s
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem) SetPackageType(v string) *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem {
	s.PackageType = &v
	return s
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem) SetPrivatePoolOptionsId(v string) *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem {
	s.PrivatePoolOptionsId = &v
	return s
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem) SetPrivatePoolOptionsMatchCriteria(v string) *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem {
	s.PrivatePoolOptionsMatchCriteria = &v
	return s
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem) SetPrivatePoolOptionsName(v string) *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem {
	s.PrivatePoolOptionsName = &v
	return s
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem) SetRecurrenceRules(v *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemRecurrenceRules) *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem {
	s.RecurrenceRules = v
	return s
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem) SetRegionId(v string) *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem {
	s.RegionId = &v
	return s
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem) SetResourceGroupId(v string) *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem) SetStartTime(v string) *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem {
	s.StartTime = &v
	return s
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem) SetStartTimeType(v string) *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem {
	s.StartTimeType = &v
	return s
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem) SetStatus(v string) *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem {
	s.Status = &v
	return s
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem) SetTags(v *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemTags) *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem {
	s.Tags = v
	return s
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem) SetTotalAssuranceTimes(v string) *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem {
	s.TotalAssuranceTimes = &v
	return s
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem) SetUsedAssuranceTimes(v int32) *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem {
	s.UsedAssuranceTimes = &v
	return s
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItem) Validate() error {
	if s.AllocatedResources != nil {
		if err := s.AllocatedResources.Validate(); err != nil {
			return err
		}
	}
	if s.RecurrenceRules != nil {
		if err := s.RecurrenceRules.Validate(); err != nil {
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

type DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResources struct {
	AllocatedResource []*DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResourcesAllocatedResource `json:"AllocatedResource,omitempty" xml:"AllocatedResource,omitempty" type:"Repeated"`
}

func (s DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResources) GoString() string {
	return s.String()
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResources) GetAllocatedResource() []*DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResourcesAllocatedResource {
	return s.AllocatedResource
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResources) SetAllocatedResource(v []*DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResourcesAllocatedResource) *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResources {
	s.AllocatedResource = v
	return s
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResources) Validate() error {
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

type DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResourcesAllocatedResource struct {
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// 5
	AvailableAmount *int32 `json:"AvailableAmount,omitempty" xml:"AvailableAmount,omitempty"`
	// >  This parameter is not publicly available.
	ElasticityAssuranceUsages *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResourcesAllocatedResourceElasticityAssuranceUsages `json:"ElasticityAssuranceUsages,omitempty" xml:"ElasticityAssuranceUsages,omitempty" type:"Struct"`
	// The instance type.
	//
	// example:
	//
	// ecs.c6.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The total number of instances for which capacity of an instance type is reserved.
	//
	// example:
	//
	// 2
	TotalAmount *int32 `json:"TotalAmount,omitempty" xml:"TotalAmount,omitempty"`
	// The number of instances that have used the elasticity assurance.
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

func (s DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResourcesAllocatedResource) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResourcesAllocatedResource) GoString() string {
	return s.String()
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResourcesAllocatedResource) GetAvailableAmount() *int32 {
	return s.AvailableAmount
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResourcesAllocatedResource) GetElasticityAssuranceUsages() *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResourcesAllocatedResourceElasticityAssuranceUsages {
	return s.ElasticityAssuranceUsages
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResourcesAllocatedResource) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResourcesAllocatedResource) GetTotalAmount() *int32 {
	return s.TotalAmount
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResourcesAllocatedResource) GetUsedAmount() *int32 {
	return s.UsedAmount
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResourcesAllocatedResource) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResourcesAllocatedResource) SetAvailableAmount(v int32) *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResourcesAllocatedResource {
	s.AvailableAmount = &v
	return s
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResourcesAllocatedResource) SetElasticityAssuranceUsages(v *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResourcesAllocatedResourceElasticityAssuranceUsages) *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResourcesAllocatedResource {
	s.ElasticityAssuranceUsages = v
	return s
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResourcesAllocatedResource) SetInstanceType(v string) *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResourcesAllocatedResource {
	s.InstanceType = &v
	return s
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResourcesAllocatedResource) SetTotalAmount(v int32) *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResourcesAllocatedResource {
	s.TotalAmount = &v
	return s
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResourcesAllocatedResource) SetUsedAmount(v int32) *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResourcesAllocatedResource {
	s.UsedAmount = &v
	return s
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResourcesAllocatedResource) SetZoneId(v string) *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResourcesAllocatedResource {
	s.ZoneId = &v
	return s
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResourcesAllocatedResource) Validate() error {
	if s.ElasticityAssuranceUsages != nil {
		if err := s.ElasticityAssuranceUsages.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResourcesAllocatedResourceElasticityAssuranceUsages struct {
	ElasticityAssuranceUsage []*DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResourcesAllocatedResourceElasticityAssuranceUsagesElasticityAssuranceUsage `json:"ElasticityAssuranceUsage,omitempty" xml:"ElasticityAssuranceUsage,omitempty" type:"Repeated"`
}

func (s DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResourcesAllocatedResourceElasticityAssuranceUsages) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResourcesAllocatedResourceElasticityAssuranceUsages) GoString() string {
	return s.String()
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResourcesAllocatedResourceElasticityAssuranceUsages) GetElasticityAssuranceUsage() []*DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResourcesAllocatedResourceElasticityAssuranceUsagesElasticityAssuranceUsage {
	return s.ElasticityAssuranceUsage
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResourcesAllocatedResourceElasticityAssuranceUsages) SetElasticityAssuranceUsage(v []*DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResourcesAllocatedResourceElasticityAssuranceUsagesElasticityAssuranceUsage) *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResourcesAllocatedResourceElasticityAssuranceUsages {
	s.ElasticityAssuranceUsage = v
	return s
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResourcesAllocatedResourceElasticityAssuranceUsages) Validate() error {
	if s.ElasticityAssuranceUsage != nil {
		for _, item := range s.ElasticityAssuranceUsage {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResourcesAllocatedResourceElasticityAssuranceUsagesElasticityAssuranceUsage struct {
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// 1144775968548340
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// name
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// 1
	UsedAmount *int32 `json:"UsedAmount,omitempty" xml:"UsedAmount,omitempty"`
}

func (s DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResourcesAllocatedResourceElasticityAssuranceUsagesElasticityAssuranceUsage) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResourcesAllocatedResourceElasticityAssuranceUsagesElasticityAssuranceUsage) GoString() string {
	return s.String()
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResourcesAllocatedResourceElasticityAssuranceUsagesElasticityAssuranceUsage) GetAccountId() *string {
	return s.AccountId
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResourcesAllocatedResourceElasticityAssuranceUsagesElasticityAssuranceUsage) GetServiceName() *string {
	return s.ServiceName
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResourcesAllocatedResourceElasticityAssuranceUsagesElasticityAssuranceUsage) GetUsedAmount() *int32 {
	return s.UsedAmount
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResourcesAllocatedResourceElasticityAssuranceUsagesElasticityAssuranceUsage) SetAccountId(v string) *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResourcesAllocatedResourceElasticityAssuranceUsagesElasticityAssuranceUsage {
	s.AccountId = &v
	return s
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResourcesAllocatedResourceElasticityAssuranceUsagesElasticityAssuranceUsage) SetServiceName(v string) *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResourcesAllocatedResourceElasticityAssuranceUsagesElasticityAssuranceUsage {
	s.ServiceName = &v
	return s
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResourcesAllocatedResourceElasticityAssuranceUsagesElasticityAssuranceUsage) SetUsedAmount(v int32) *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResourcesAllocatedResourceElasticityAssuranceUsagesElasticityAssuranceUsage {
	s.UsedAmount = &v
	return s
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemAllocatedResourcesAllocatedResourceElasticityAssuranceUsagesElasticityAssuranceUsage) Validate() error {
	return dara.Validate(s)
}

type DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemRecurrenceRules struct {
	RecurrenceRule []*DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemRecurrenceRulesRecurrenceRule `json:"RecurrenceRule,omitempty" xml:"RecurrenceRule,omitempty" type:"Repeated"`
}

func (s DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemRecurrenceRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemRecurrenceRules) GoString() string {
	return s.String()
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemRecurrenceRules) GetRecurrenceRule() []*DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemRecurrenceRulesRecurrenceRule {
	return s.RecurrenceRule
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemRecurrenceRules) SetRecurrenceRule(v []*DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemRecurrenceRulesRecurrenceRule) *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemRecurrenceRules {
	s.RecurrenceRule = v
	return s
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemRecurrenceRules) Validate() error {
	if s.RecurrenceRule != nil {
		for _, item := range s.RecurrenceRule {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemRecurrenceRulesRecurrenceRule struct {
	// The time when the time-segmented assurance ends.
	//
	// example:
	//
	// 10
	EndHour *int32 `json:"EndHour,omitempty" xml:"EndHour,omitempty"`
	// The type of the recurrence rule. Valid values:
	//
	// 	- Daily
	//
	// 	- Weekly
	//
	// 	- Monthly
	//
	// example:
	//
	// Daily
	RecurrenceType *string `json:"RecurrenceType,omitempty" xml:"RecurrenceType,omitempty"`
	// The recurrence value of the time-segmented assurance. Valid values:
	//
	// 	- If you set `RecurrenceType` to `Daily`, you can set RecurrenceValue to only one value. Valid values: 1 to 31. The time-segmented assurance is performed every few days.
	//
	// 	- If you set `RecurrenceType` to `Weekly`, you can set RecurrenceValue to one or more values. Separate the values with commas (,). The values that correspond to Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, and Saturday are 0, 1, 2, 3, 4, 5, and 6. For example, `1,2` indicates that the time-segmented assurance is performed on Monday and Tuesday of every week.
	//
	// 	- If you set `RecurrenceType` to `Monthly`, you can set RecurrenceValue to two values in the `A-B` format. Valid values of A and B: 1 to 31. B must be greater than or equal to A. For example, `1-5` indicates that the time-segmented assurance is performed from the 1st to the 5th of each month.
	//
	// example:
	//
	// 1
	RecurrenceValue *string `json:"RecurrenceValue,omitempty" xml:"RecurrenceValue,omitempty"`
	// The time when the time-segmented assurance takes effect.
	//
	// example:
	//
	// 4
	StartHour *int32 `json:"StartHour,omitempty" xml:"StartHour,omitempty"`
}

func (s DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemRecurrenceRulesRecurrenceRule) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemRecurrenceRulesRecurrenceRule) GoString() string {
	return s.String()
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemRecurrenceRulesRecurrenceRule) GetEndHour() *int32 {
	return s.EndHour
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemRecurrenceRulesRecurrenceRule) GetRecurrenceType() *string {
	return s.RecurrenceType
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemRecurrenceRulesRecurrenceRule) GetRecurrenceValue() *string {
	return s.RecurrenceValue
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemRecurrenceRulesRecurrenceRule) GetStartHour() *int32 {
	return s.StartHour
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemRecurrenceRulesRecurrenceRule) SetEndHour(v int32) *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemRecurrenceRulesRecurrenceRule {
	s.EndHour = &v
	return s
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemRecurrenceRulesRecurrenceRule) SetRecurrenceType(v string) *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemRecurrenceRulesRecurrenceRule {
	s.RecurrenceType = &v
	return s
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemRecurrenceRulesRecurrenceRule) SetRecurrenceValue(v string) *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemRecurrenceRulesRecurrenceRule {
	s.RecurrenceValue = &v
	return s
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemRecurrenceRulesRecurrenceRule) SetStartHour(v int32) *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemRecurrenceRulesRecurrenceRule {
	s.StartHour = &v
	return s
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemRecurrenceRulesRecurrenceRule) Validate() error {
	return dara.Validate(s)
}

type DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemTags struct {
	Tag []*DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemTags) GoString() string {
	return s.String()
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemTags) GetTag() []*DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemTagsTag {
	return s.Tag
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemTags) SetTag(v []*DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemTagsTag) *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemTags {
	s.Tag = v
	return s
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemTags) Validate() error {
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

type DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemTagsTag struct {
	// The key of the tag.
	//
	// example:
	//
	// TestKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// TestValue
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemTagsTag) SetTagKey(v string) *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemTagsTag {
	s.TagKey = &v
	return s
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemTagsTag) SetTagValue(v string) *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemTagsTag {
	s.TagValue = &v
	return s
}

func (s *DescribeElasticityAssurancesResponseBodyElasticityAssuranceSetElasticityAssuranceItemTagsTag) Validate() error {
	return dara.Validate(s)
}
