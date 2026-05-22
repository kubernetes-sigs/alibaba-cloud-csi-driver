// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeReservedInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeReservedInstancesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeReservedInstancesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeReservedInstancesResponseBody
	GetRequestId() *string
	SetReservedInstances(v *DescribeReservedInstancesResponseBodyReservedInstances) *DescribeReservedInstancesResponseBody
	GetReservedInstances() *DescribeReservedInstancesResponseBodyReservedInstances
	SetTotalCount(v int32) *DescribeReservedInstancesResponseBody
	GetTotalCount() *int32
}

type DescribeReservedInstancesResponseBody struct {
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
	// The request ID.
	//
	// example:
	//
	// E572643C-6A29-49D6-9D4E-6CFA4E063A3E
	RequestId         *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ReservedInstances *DescribeReservedInstancesResponseBodyReservedInstances `json:"ReservedInstances,omitempty" xml:"ReservedInstances,omitempty" type:"Struct"`
	// The total number of reserved instances.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeReservedInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeReservedInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeReservedInstancesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeReservedInstancesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeReservedInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeReservedInstancesResponseBody) GetReservedInstances() *DescribeReservedInstancesResponseBodyReservedInstances {
	return s.ReservedInstances
}

func (s *DescribeReservedInstancesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeReservedInstancesResponseBody) SetPageNumber(v int32) *DescribeReservedInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeReservedInstancesResponseBody) SetPageSize(v int32) *DescribeReservedInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeReservedInstancesResponseBody) SetRequestId(v string) *DescribeReservedInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeReservedInstancesResponseBody) SetReservedInstances(v *DescribeReservedInstancesResponseBodyReservedInstances) *DescribeReservedInstancesResponseBody {
	s.ReservedInstances = v
	return s
}

func (s *DescribeReservedInstancesResponseBody) SetTotalCount(v int32) *DescribeReservedInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeReservedInstancesResponseBody) Validate() error {
	if s.ReservedInstances != nil {
		if err := s.ReservedInstances.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeReservedInstancesResponseBodyReservedInstances struct {
	ReservedInstance []*DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance `json:"ReservedInstance,omitempty" xml:"ReservedInstance,omitempty" type:"Repeated"`
}

func (s DescribeReservedInstancesResponseBodyReservedInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeReservedInstancesResponseBodyReservedInstances) GoString() string {
	return s.String()
}

func (s *DescribeReservedInstancesResponseBodyReservedInstances) GetReservedInstance() []*DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance {
	return s.ReservedInstance
}

func (s *DescribeReservedInstancesResponseBodyReservedInstances) SetReservedInstance(v []*DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance) *DescribeReservedInstancesResponseBodyReservedInstances {
	s.ReservedInstance = v
	return s
}

func (s *DescribeReservedInstancesResponseBodyReservedInstances) Validate() error {
	if s.ReservedInstance != nil {
		for _, item := range s.ReservedInstance {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance struct {
	AllocationStatus     *string                                                                               `json:"AllocationStatus,omitempty" xml:"AllocationStatus,omitempty"`
	CreationTime         *string                                                                               `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description          *string                                                                               `json:"Description,omitempty" xml:"Description,omitempty"`
	ExpiredTime          *string                                                                               `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	InstanceAmount       *int32                                                                                `json:"InstanceAmount,omitempty" xml:"InstanceAmount,omitempty"`
	InstanceType         *string                                                                               `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	OfferingType         *string                                                                               `json:"OfferingType,omitempty" xml:"OfferingType,omitempty"`
	OperationLocks       *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstanceOperationLocks `json:"OperationLocks,omitempty" xml:"OperationLocks,omitempty" type:"Struct"`
	Platform             *string                                                                               `json:"Platform,omitempty" xml:"Platform,omitempty"`
	RegionId             *string                                                                               `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ReservedInstanceId   *string                                                                               `json:"ReservedInstanceId,omitempty" xml:"ReservedInstanceId,omitempty"`
	ReservedInstanceName *string                                                                               `json:"ReservedInstanceName,omitempty" xml:"ReservedInstanceName,omitempty"`
	ResourceGroupId      *string                                                                               `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Scope                *string                                                                               `json:"Scope,omitempty" xml:"Scope,omitempty"`
	StartTime            *string                                                                               `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status               *string                                                                               `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags                 *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstanceTags           `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	ZoneId               *string                                                                               `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance) String() string {
	return dara.Prettify(s)
}

func (s DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance) GoString() string {
	return s.String()
}

func (s *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance) GetAllocationStatus() *string {
	return s.AllocationStatus
}

func (s *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance) GetDescription() *string {
	return s.Description
}

func (s *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance) GetInstanceAmount() *int32 {
	return s.InstanceAmount
}

func (s *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance) GetOfferingType() *string {
	return s.OfferingType
}

func (s *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance) GetOperationLocks() *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstanceOperationLocks {
	return s.OperationLocks
}

func (s *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance) GetPlatform() *string {
	return s.Platform
}

func (s *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance) GetReservedInstanceId() *string {
	return s.ReservedInstanceId
}

func (s *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance) GetReservedInstanceName() *string {
	return s.ReservedInstanceName
}

func (s *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance) GetScope() *string {
	return s.Scope
}

func (s *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance) GetStatus() *string {
	return s.Status
}

func (s *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance) GetTags() *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstanceTags {
	return s.Tags
}

func (s *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance) SetAllocationStatus(v string) *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance {
	s.AllocationStatus = &v
	return s
}

func (s *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance) SetCreationTime(v string) *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance {
	s.CreationTime = &v
	return s
}

func (s *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance) SetDescription(v string) *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance {
	s.Description = &v
	return s
}

func (s *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance) SetExpiredTime(v string) *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance) SetInstanceAmount(v int32) *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance {
	s.InstanceAmount = &v
	return s
}

func (s *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance) SetInstanceType(v string) *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance {
	s.InstanceType = &v
	return s
}

func (s *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance) SetOfferingType(v string) *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance {
	s.OfferingType = &v
	return s
}

func (s *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance) SetOperationLocks(v *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstanceOperationLocks) *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance {
	s.OperationLocks = v
	return s
}

func (s *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance) SetPlatform(v string) *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance {
	s.Platform = &v
	return s
}

func (s *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance) SetRegionId(v string) *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance {
	s.RegionId = &v
	return s
}

func (s *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance) SetReservedInstanceId(v string) *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance {
	s.ReservedInstanceId = &v
	return s
}

func (s *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance) SetReservedInstanceName(v string) *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance {
	s.ReservedInstanceName = &v
	return s
}

func (s *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance) SetResourceGroupId(v string) *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance) SetScope(v string) *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance {
	s.Scope = &v
	return s
}

func (s *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance) SetStartTime(v string) *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance {
	s.StartTime = &v
	return s
}

func (s *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance) SetStatus(v string) *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance {
	s.Status = &v
	return s
}

func (s *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance) SetTags(v *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstanceTags) *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance {
	s.Tags = v
	return s
}

func (s *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance) SetZoneId(v string) *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance {
	s.ZoneId = &v
	return s
}

func (s *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstance) Validate() error {
	if s.OperationLocks != nil {
		if err := s.OperationLocks.Validate(); err != nil {
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

type DescribeReservedInstancesResponseBodyReservedInstancesReservedInstanceOperationLocks struct {
	OperationLock []*DescribeReservedInstancesResponseBodyReservedInstancesReservedInstanceOperationLocksOperationLock `json:"OperationLock,omitempty" xml:"OperationLock,omitempty" type:"Repeated"`
}

func (s DescribeReservedInstancesResponseBodyReservedInstancesReservedInstanceOperationLocks) String() string {
	return dara.Prettify(s)
}

func (s DescribeReservedInstancesResponseBodyReservedInstancesReservedInstanceOperationLocks) GoString() string {
	return s.String()
}

func (s *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstanceOperationLocks) GetOperationLock() []*DescribeReservedInstancesResponseBodyReservedInstancesReservedInstanceOperationLocksOperationLock {
	return s.OperationLock
}

func (s *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstanceOperationLocks) SetOperationLock(v []*DescribeReservedInstancesResponseBodyReservedInstancesReservedInstanceOperationLocksOperationLock) *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstanceOperationLocks {
	s.OperationLock = v
	return s
}

func (s *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstanceOperationLocks) Validate() error {
	if s.OperationLock != nil {
		for _, item := range s.OperationLock {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeReservedInstancesResponseBodyReservedInstancesReservedInstanceOperationLocksOperationLock struct {
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
}

func (s DescribeReservedInstancesResponseBodyReservedInstancesReservedInstanceOperationLocksOperationLock) String() string {
	return dara.Prettify(s)
}

func (s DescribeReservedInstancesResponseBodyReservedInstancesReservedInstanceOperationLocksOperationLock) GoString() string {
	return s.String()
}

func (s *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstanceOperationLocksOperationLock) GetLockReason() *string {
	return s.LockReason
}

func (s *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstanceOperationLocksOperationLock) SetLockReason(v string) *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstanceOperationLocksOperationLock {
	s.LockReason = &v
	return s
}

func (s *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstanceOperationLocksOperationLock) Validate() error {
	return dara.Validate(s)
}

type DescribeReservedInstancesResponseBodyReservedInstancesReservedInstanceTags struct {
	Tag []*DescribeReservedInstancesResponseBodyReservedInstancesReservedInstanceTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeReservedInstancesResponseBodyReservedInstancesReservedInstanceTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeReservedInstancesResponseBodyReservedInstancesReservedInstanceTags) GoString() string {
	return s.String()
}

func (s *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstanceTags) GetTag() []*DescribeReservedInstancesResponseBodyReservedInstancesReservedInstanceTagsTag {
	return s.Tag
}

func (s *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstanceTags) SetTag(v []*DescribeReservedInstancesResponseBodyReservedInstancesReservedInstanceTagsTag) *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstanceTags {
	s.Tag = v
	return s
}

func (s *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstanceTags) Validate() error {
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

type DescribeReservedInstancesResponseBodyReservedInstancesReservedInstanceTagsTag struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeReservedInstancesResponseBodyReservedInstancesReservedInstanceTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeReservedInstancesResponseBodyReservedInstancesReservedInstanceTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstanceTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstanceTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstanceTagsTag) SetTagKey(v string) *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstanceTagsTag {
	s.TagKey = &v
	return s
}

func (s *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstanceTagsTag) SetTagValue(v string) *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstanceTagsTag {
	s.TagValue = &v
	return s
}

func (s *DescribeReservedInstancesResponseBodyReservedInstancesReservedInstanceTagsTag) Validate() error {
	return dara.Validate(s)
}
