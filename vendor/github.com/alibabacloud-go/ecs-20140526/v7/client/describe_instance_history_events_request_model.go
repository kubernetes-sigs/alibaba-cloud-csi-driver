// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceHistoryEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventPublishTime(v *DescribeInstanceHistoryEventsRequestEventPublishTime) *DescribeInstanceHistoryEventsRequest
	GetEventPublishTime() *DescribeInstanceHistoryEventsRequestEventPublishTime
	SetNotBefore(v *DescribeInstanceHistoryEventsRequestNotBefore) *DescribeInstanceHistoryEventsRequest
	GetNotBefore() *DescribeInstanceHistoryEventsRequestNotBefore
	SetEventCycleStatus(v string) *DescribeInstanceHistoryEventsRequest
	GetEventCycleStatus() *string
	SetEventId(v []*string) *DescribeInstanceHistoryEventsRequest
	GetEventId() []*string
	SetEventType(v string) *DescribeInstanceHistoryEventsRequest
	GetEventType() *string
	SetImpactLevel(v string) *DescribeInstanceHistoryEventsRequest
	GetImpactLevel() *string
	SetInstanceEventCycleStatus(v []*string) *DescribeInstanceHistoryEventsRequest
	GetInstanceEventCycleStatus() []*string
	SetInstanceEventType(v []*string) *DescribeInstanceHistoryEventsRequest
	GetInstanceEventType() []*string
	SetInstanceId(v string) *DescribeInstanceHistoryEventsRequest
	GetInstanceId() *string
	SetMaxResults(v int64) *DescribeInstanceHistoryEventsRequest
	GetMaxResults() *int64
	SetNextToken(v string) *DescribeInstanceHistoryEventsRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *DescribeInstanceHistoryEventsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeInstanceHistoryEventsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeInstanceHistoryEventsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeInstanceHistoryEventsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeInstanceHistoryEventsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeInstanceHistoryEventsRequest
	GetResourceGroupId() *string
	SetResourceId(v []*string) *DescribeInstanceHistoryEventsRequest
	GetResourceId() []*string
	SetResourceOwnerAccount(v string) *DescribeInstanceHistoryEventsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeInstanceHistoryEventsRequest
	GetResourceOwnerId() *int64
	SetResourceType(v string) *DescribeInstanceHistoryEventsRequest
	GetResourceType() *string
	SetTag(v []*DescribeInstanceHistoryEventsRequestTag) *DescribeInstanceHistoryEventsRequest
	GetTag() []*DescribeInstanceHistoryEventsRequestTag
}

type DescribeInstanceHistoryEventsRequest struct {
	EventPublishTime *DescribeInstanceHistoryEventsRequestEventPublishTime `json:"EventPublishTime,omitempty" xml:"EventPublishTime,omitempty" type:"Struct"`
	NotBefore        *DescribeInstanceHistoryEventsRequestNotBefore        `json:"NotBefore,omitempty" xml:"NotBefore,omitempty" type:"Struct"`
	// The lifecycle state of the system event. This parameter takes effect only when InstanceEventCycleStatus.N is not specified. Valid values:
	//
	// 	- Scheduled
	//
	// 	- Avoided
	//
	// 	- Executing
	//
	// 	- Executed
	//
	// 	- Canceled
	//
	// 	- Failed
	//
	// 	- Inquiring
	//
	// example:
	//
	// Executed
	EventCycleStatus *string `json:"EventCycleStatus,omitempty" xml:"EventCycleStatus,omitempty"`
	// The ID of system event N. Valid values of N: 1 to 100. You can repeat this parameter to pass multiple values.
	//
	// example:
	//
	// e-uf64yvznlao4jl2c****
	EventId []*string `json:"EventId,omitempty" xml:"EventId,omitempty" type:"Repeated"`
	// The type of the system event. This parameter takes effect only when InstanceEventType.N is not specified. Valid values:
	//
	// 	- SystemMaintenance.Reboot: The instance is restarted due to system maintenance.
	//
	// 	- SystemMaintenance.Redeploy: The instance is redeployed due to system maintenance.
	//
	// 	- SystemFailure.Reboot: The instance is restarted due to a system error.
	//
	// 	- SystemFailure.Redeploy: The instance is redeployed due to a system error.
	//
	// 	- SystemFailure.Delete: The instance is released due to an instance creation failure.
	//
	// 	- InstanceFailure.Reboot: The instance is restarted due to an instance error.
	//
	// 	- InstanceExpiration.Stop: The subscription instance is stopped due to expiration.
	//
	// 	- InstanceExpiration.Delete: The subscription instance is released due to expiration.
	//
	// 	- AccountUnbalanced.Stop: The pay-as-you-go instance is stopped due to an overdue payment.
	//
	// 	- AccountUnbalanced.Delete: The pay-as-you-go instance is released due to an overdue payment.
	//
	// >  For more information, see [Overview](https://help.aliyun.com/document_detail/66574.html). The values of this parameter are applicable only to instance system events, but not to disk system events.
	//
	// example:
	//
	// SystemMaintenance.Reboot
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// null
	ImpactLevel *string `json:"ImpactLevel,omitempty" xml:"ImpactLevel,omitempty"`
	// The lifecycle state of system event N. Valid values of N: 1 to 7. You can repeat this parameter to pass multiple values. Valid values:
	//
	// 	- Scheduled
	//
	// 	- Avoided
	//
	// 	- Executing
	//
	// 	- Executed
	//
	// 	- Canceled
	//
	// 	- Failed
	//
	// 	- Inquiring
	//
	// example:
	//
	// Executed
	InstanceEventCycleStatus []*string `json:"InstanceEventCycleStatus,omitempty" xml:"InstanceEventCycleStatus,omitempty" type:"Repeated"`
	// The type of system event N. Valid values of N: 1 to 30. You can repeat this parameter to pass multiple values. Valid values:
	//
	// 	- SystemMaintenance.Reboot: The instance is restarted due to system maintenance.
	//
	// 	- SystemMaintenance.Redeploy: The instance is redeployed due to system maintenance.
	//
	// 	- SystemFailure.Reboot: The instance is restarted due to a system error.
	//
	// 	- SystemFailure.Redeploy: The instance is redeployed due to a system error.
	//
	// 	- SystemFailure.Delete: The instance is released due to an instance creation failure.
	//
	// 	- InstanceFailure.Reboot: The instance is restarted due to an instance error.
	//
	// 	- InstanceExpiration.Stop: The subscription instance is stopped due to expiration.
	//
	// 	- InstanceExpiration.Delete: The subscription instance is released due to expiration.
	//
	// 	- AccountUnbalanced.Stop: The pay-as-you-go instance is stopped due to an overdue payment.
	//
	// 	- AccountUnbalanced.Delete: The pay-as-you-go instance is released due to an overdue payment.
	//
	// >  For more information, see [Overview](https://help.aliyun.com/document_detail/66574.html). The values of this parameter are applicable only to instance system events, but not to disk system events.
	//
	// example:
	//
	// SystemMaintenance.Reboot
	InstanceEventType []*string `json:"InstanceEventType,omitempty" xml:"InstanceEventType,omitempty" type:"Repeated"`
	// The ID of the instance. If this parameter is not specified, the system events of all instances in the specified region are queried.
	//
	// example:
	//
	// i-uf678mass4zvr9n1****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of entries to return on each page. Valid values: 10 to 100.
	//
	// Default values:
	//
	// 	- If you set a value greater than 0 and less than 10, the default value is 10.
	//
	// 	- If you set this parameter to a value that is greater than 100, the default value is 100.
	//
	// example:
	//
	// 100
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// f1c9fa9de5752***
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// >  This parameter is deprecated. We recommend that you specify MaxResults or NextToken for a paged query.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// >  This parameter is deprecated. We recommend that you specify MaxResults or NextToken for a paged query.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the resource. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the resource belongs.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of resource N. Valid values of N: 1 to 100. You can repeat this parameter to pass multiple values. Valid values:
	//
	// 	- When `ResourceType` is set to instance, ResourceId.N specifies the ID of instance N.
	//
	// 	- When `ResourceType` is set to ddh, ResourceId.N specifies the ID of dedicated host N.
	//
	// 	- When `ResourceType` is set to managedhost, ResourceId.N specifies the ID of physical machine N from a smart hosting pool.
	//
	// If this parameter is not specified, the system events of all resources of the type specified by `ResourceType` in the region specified by `RegionId` are queried.
	//
	// >  We recommend that you use `ResourceId.N` to specify one or more resource IDs. If you specify both `ResourceId.N` and `InstanceId`, `ResourceId.N` takes precedence by default.
	//
	// example:
	//
	// i-uf678mass4zvr9n1****
	ResourceId           []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of the resource. Valid values:
	//
	// 	- instance: ECS instance
	//
	// 	- ddh: dedicated host
	//
	// 	- managehost: physical machine in a smart hosting pool
	//
	// Default value: instance.
	//
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The list of tags.
	Tag []*DescribeInstanceHistoryEventsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeInstanceHistoryEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceHistoryEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceHistoryEventsRequest) GetEventPublishTime() *DescribeInstanceHistoryEventsRequestEventPublishTime {
	return s.EventPublishTime
}

func (s *DescribeInstanceHistoryEventsRequest) GetNotBefore() *DescribeInstanceHistoryEventsRequestNotBefore {
	return s.NotBefore
}

func (s *DescribeInstanceHistoryEventsRequest) GetEventCycleStatus() *string {
	return s.EventCycleStatus
}

func (s *DescribeInstanceHistoryEventsRequest) GetEventId() []*string {
	return s.EventId
}

func (s *DescribeInstanceHistoryEventsRequest) GetEventType() *string {
	return s.EventType
}

func (s *DescribeInstanceHistoryEventsRequest) GetImpactLevel() *string {
	return s.ImpactLevel
}

func (s *DescribeInstanceHistoryEventsRequest) GetInstanceEventCycleStatus() []*string {
	return s.InstanceEventCycleStatus
}

func (s *DescribeInstanceHistoryEventsRequest) GetInstanceEventType() []*string {
	return s.InstanceEventType
}

func (s *DescribeInstanceHistoryEventsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceHistoryEventsRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *DescribeInstanceHistoryEventsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeInstanceHistoryEventsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeInstanceHistoryEventsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeInstanceHistoryEventsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeInstanceHistoryEventsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInstanceHistoryEventsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstanceHistoryEventsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeInstanceHistoryEventsRequest) GetResourceId() []*string {
	return s.ResourceId
}

func (s *DescribeInstanceHistoryEventsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeInstanceHistoryEventsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeInstanceHistoryEventsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeInstanceHistoryEventsRequest) GetTag() []*DescribeInstanceHistoryEventsRequestTag {
	return s.Tag
}

func (s *DescribeInstanceHistoryEventsRequest) SetEventPublishTime(v *DescribeInstanceHistoryEventsRequestEventPublishTime) *DescribeInstanceHistoryEventsRequest {
	s.EventPublishTime = v
	return s
}

func (s *DescribeInstanceHistoryEventsRequest) SetNotBefore(v *DescribeInstanceHistoryEventsRequestNotBefore) *DescribeInstanceHistoryEventsRequest {
	s.NotBefore = v
	return s
}

func (s *DescribeInstanceHistoryEventsRequest) SetEventCycleStatus(v string) *DescribeInstanceHistoryEventsRequest {
	s.EventCycleStatus = &v
	return s
}

func (s *DescribeInstanceHistoryEventsRequest) SetEventId(v []*string) *DescribeInstanceHistoryEventsRequest {
	s.EventId = v
	return s
}

func (s *DescribeInstanceHistoryEventsRequest) SetEventType(v string) *DescribeInstanceHistoryEventsRequest {
	s.EventType = &v
	return s
}

func (s *DescribeInstanceHistoryEventsRequest) SetImpactLevel(v string) *DescribeInstanceHistoryEventsRequest {
	s.ImpactLevel = &v
	return s
}

func (s *DescribeInstanceHistoryEventsRequest) SetInstanceEventCycleStatus(v []*string) *DescribeInstanceHistoryEventsRequest {
	s.InstanceEventCycleStatus = v
	return s
}

func (s *DescribeInstanceHistoryEventsRequest) SetInstanceEventType(v []*string) *DescribeInstanceHistoryEventsRequest {
	s.InstanceEventType = v
	return s
}

func (s *DescribeInstanceHistoryEventsRequest) SetInstanceId(v string) *DescribeInstanceHistoryEventsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceHistoryEventsRequest) SetMaxResults(v int64) *DescribeInstanceHistoryEventsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeInstanceHistoryEventsRequest) SetNextToken(v string) *DescribeInstanceHistoryEventsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeInstanceHistoryEventsRequest) SetOwnerAccount(v string) *DescribeInstanceHistoryEventsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeInstanceHistoryEventsRequest) SetOwnerId(v int64) *DescribeInstanceHistoryEventsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeInstanceHistoryEventsRequest) SetPageNumber(v int32) *DescribeInstanceHistoryEventsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstanceHistoryEventsRequest) SetPageSize(v int32) *DescribeInstanceHistoryEventsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstanceHistoryEventsRequest) SetRegionId(v string) *DescribeInstanceHistoryEventsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceHistoryEventsRequest) SetResourceGroupId(v string) *DescribeInstanceHistoryEventsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstanceHistoryEventsRequest) SetResourceId(v []*string) *DescribeInstanceHistoryEventsRequest {
	s.ResourceId = v
	return s
}

func (s *DescribeInstanceHistoryEventsRequest) SetResourceOwnerAccount(v string) *DescribeInstanceHistoryEventsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeInstanceHistoryEventsRequest) SetResourceOwnerId(v int64) *DescribeInstanceHistoryEventsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeInstanceHistoryEventsRequest) SetResourceType(v string) *DescribeInstanceHistoryEventsRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeInstanceHistoryEventsRequest) SetTag(v []*DescribeInstanceHistoryEventsRequestTag) *DescribeInstanceHistoryEventsRequest {
	s.Tag = v
	return s
}

func (s *DescribeInstanceHistoryEventsRequest) Validate() error {
	if s.EventPublishTime != nil {
		if err := s.EventPublishTime.Validate(); err != nil {
			return err
		}
	}
	if s.NotBefore != nil {
		if err := s.NotBefore.Validate(); err != nil {
			return err
		}
	}
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

type DescribeInstanceHistoryEventsRequestEventPublishTime struct {
	// The end of the time range in which to query published system events. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2017-12-01T06:32:31Z
	End *string `json:"End,omitempty" xml:"End,omitempty"`
	// The beginning of the time range in which to query published system events. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2017-11-30T06:32:31Z
	Start *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s DescribeInstanceHistoryEventsRequestEventPublishTime) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceHistoryEventsRequestEventPublishTime) GoString() string {
	return s.String()
}

func (s *DescribeInstanceHistoryEventsRequestEventPublishTime) GetEnd() *string {
	return s.End
}

func (s *DescribeInstanceHistoryEventsRequestEventPublishTime) GetStart() *string {
	return s.Start
}

func (s *DescribeInstanceHistoryEventsRequestEventPublishTime) SetEnd(v string) *DescribeInstanceHistoryEventsRequestEventPublishTime {
	s.End = &v
	return s
}

func (s *DescribeInstanceHistoryEventsRequestEventPublishTime) SetStart(v string) *DescribeInstanceHistoryEventsRequestEventPublishTime {
	s.Start = &v
	return s
}

func (s *DescribeInstanceHistoryEventsRequestEventPublishTime) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceHistoryEventsRequestNotBefore struct {
	// The latest scheduled end time for the system event. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2017-12-01T06:32:31Z
	End *string `json:"End,omitempty" xml:"End,omitempty"`
	// The earliest scheduled start time for the system event. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2017-11-30T06:32:31Z
	Start *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s DescribeInstanceHistoryEventsRequestNotBefore) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceHistoryEventsRequestNotBefore) GoString() string {
	return s.String()
}

func (s *DescribeInstanceHistoryEventsRequestNotBefore) GetEnd() *string {
	return s.End
}

func (s *DescribeInstanceHistoryEventsRequestNotBefore) GetStart() *string {
	return s.Start
}

func (s *DescribeInstanceHistoryEventsRequestNotBefore) SetEnd(v string) *DescribeInstanceHistoryEventsRequestNotBefore {
	s.End = &v
	return s
}

func (s *DescribeInstanceHistoryEventsRequestNotBefore) SetStart(v string) *DescribeInstanceHistoryEventsRequestNotBefore {
	s.Start = &v
	return s
}

func (s *DescribeInstanceHistoryEventsRequestNotBefore) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceHistoryEventsRequestTag struct {
	// The key of tag N of the resource.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N of the resource.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeInstanceHistoryEventsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceHistoryEventsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeInstanceHistoryEventsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeInstanceHistoryEventsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeInstanceHistoryEventsRequestTag) SetKey(v string) *DescribeInstanceHistoryEventsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeInstanceHistoryEventsRequestTag) SetValue(v string) *DescribeInstanceHistoryEventsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeInstanceHistoryEventsRequestTag) Validate() error {
	return dara.Validate(s)
}
