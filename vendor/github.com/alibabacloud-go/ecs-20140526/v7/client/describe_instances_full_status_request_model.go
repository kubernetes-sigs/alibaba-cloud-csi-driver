// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstancesFullStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventPublishTime(v *DescribeInstancesFullStatusRequestEventPublishTime) *DescribeInstancesFullStatusRequest
	GetEventPublishTime() *DescribeInstancesFullStatusRequestEventPublishTime
	SetNotBefore(v *DescribeInstancesFullStatusRequestNotBefore) *DescribeInstancesFullStatusRequest
	GetNotBefore() *DescribeInstancesFullStatusRequestNotBefore
	SetEventId(v []*string) *DescribeInstancesFullStatusRequest
	GetEventId() []*string
	SetEventType(v string) *DescribeInstancesFullStatusRequest
	GetEventType() *string
	SetHealthStatus(v string) *DescribeInstancesFullStatusRequest
	GetHealthStatus() *string
	SetInstanceEventType(v []*string) *DescribeInstancesFullStatusRequest
	GetInstanceEventType() []*string
	SetInstanceId(v []*string) *DescribeInstancesFullStatusRequest
	GetInstanceId() []*string
	SetOwnerAccount(v string) *DescribeInstancesFullStatusRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeInstancesFullStatusRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeInstancesFullStatusRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeInstancesFullStatusRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeInstancesFullStatusRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeInstancesFullStatusRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeInstancesFullStatusRequest
	GetResourceOwnerId() *int64
	SetStatus(v string) *DescribeInstancesFullStatusRequest
	GetStatus() *string
}

type DescribeInstancesFullStatusRequest struct {
	EventPublishTime *DescribeInstancesFullStatusRequestEventPublishTime `json:"EventPublishTime,omitempty" xml:"EventPublishTime,omitempty" type:"Struct"`
	NotBefore        *DescribeInstancesFullStatusRequestNotBefore        `json:"NotBefore,omitempty" xml:"NotBefore,omitempty" type:"Struct"`
	// The IDs of the system events. You can specify up to 100 event IDs in a single request.
	//
	// example:
	//
	// e-bp1hygp5b04o56l0****
	EventId []*string `json:"EventId,omitempty" xml:"EventId,omitempty" type:"Repeated"`
	// The type of the system event. This parameter is valid only when InstanceEventType.N is not specified. Valid values:
	//
	// 	- SystemMaintenance.Reboot: The instance is restarted due to system maintenance.
	//
	// 	- SystemFailure.Reboot: The instance is restarted due to a system failure.
	//
	// 	- InstanceFailure.Reboot: The instance is restarted due to an instance failure.
	//
	// 	- InstanceExpiration.Stop: The subscription instance is stopped due to expiration.
	//
	// 	- InstanceExpiration.Delete: The subscription instance is released due to expiration.
	//
	// 	- AccountUnbalanced.Stop: The pay-as-you-go instance is stopped due to an overdue payment.
	//
	// 	- AccountUnbalanced.Delete: The pay-as-you-go instance is released due to an overdue payment.
	//
	// example:
	//
	// InstanceExpiration.Stop
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The health status of the instance. Valid values:
	//
	// 	- Impaired
	//
	// 	- Warning: The instance performance may be degraded due to maintenance or technical issues.
	//
	// 	- Maintaining
	//
	// 	- Initializing
	//
	// 	- InsufficientData
	//
	// 	- NotApplicable
	//
	// All the values are case-sensitive.
	//
	// example:
	//
	// Maintaining
	HealthStatus *string `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	// The types of system events. You can specify up to 30 event types in a single request.
	//
	// example:
	//
	// InstanceExpiration.Stop
	InstanceEventType []*string `json:"InstanceEventType,omitempty" xml:"InstanceEventType,omitempty" type:"Repeated"`
	// The IDs of the instances. You can specify up to 100 instance IDs in a single request.
	//
	// example:
	//
	// i-bp67acfmxazb4p****
	InstanceId   []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	OwnerAccount *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. The value must be a positive integer.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The lifecycle status of the instance. Valid values:
	//
	// 	- Starting
	//
	// 	- Running
	//
	// 	- Stopped
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeInstancesFullStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesFullStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstancesFullStatusRequest) GetEventPublishTime() *DescribeInstancesFullStatusRequestEventPublishTime {
	return s.EventPublishTime
}

func (s *DescribeInstancesFullStatusRequest) GetNotBefore() *DescribeInstancesFullStatusRequestNotBefore {
	return s.NotBefore
}

func (s *DescribeInstancesFullStatusRequest) GetEventId() []*string {
	return s.EventId
}

func (s *DescribeInstancesFullStatusRequest) GetEventType() *string {
	return s.EventType
}

func (s *DescribeInstancesFullStatusRequest) GetHealthStatus() *string {
	return s.HealthStatus
}

func (s *DescribeInstancesFullStatusRequest) GetInstanceEventType() []*string {
	return s.InstanceEventType
}

func (s *DescribeInstancesFullStatusRequest) GetInstanceId() []*string {
	return s.InstanceId
}

func (s *DescribeInstancesFullStatusRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeInstancesFullStatusRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeInstancesFullStatusRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeInstancesFullStatusRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInstancesFullStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstancesFullStatusRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeInstancesFullStatusRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeInstancesFullStatusRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeInstancesFullStatusRequest) SetEventPublishTime(v *DescribeInstancesFullStatusRequestEventPublishTime) *DescribeInstancesFullStatusRequest {
	s.EventPublishTime = v
	return s
}

func (s *DescribeInstancesFullStatusRequest) SetNotBefore(v *DescribeInstancesFullStatusRequestNotBefore) *DescribeInstancesFullStatusRequest {
	s.NotBefore = v
	return s
}

func (s *DescribeInstancesFullStatusRequest) SetEventId(v []*string) *DescribeInstancesFullStatusRequest {
	s.EventId = v
	return s
}

func (s *DescribeInstancesFullStatusRequest) SetEventType(v string) *DescribeInstancesFullStatusRequest {
	s.EventType = &v
	return s
}

func (s *DescribeInstancesFullStatusRequest) SetHealthStatus(v string) *DescribeInstancesFullStatusRequest {
	s.HealthStatus = &v
	return s
}

func (s *DescribeInstancesFullStatusRequest) SetInstanceEventType(v []*string) *DescribeInstancesFullStatusRequest {
	s.InstanceEventType = v
	return s
}

func (s *DescribeInstancesFullStatusRequest) SetInstanceId(v []*string) *DescribeInstancesFullStatusRequest {
	s.InstanceId = v
	return s
}

func (s *DescribeInstancesFullStatusRequest) SetOwnerAccount(v string) *DescribeInstancesFullStatusRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeInstancesFullStatusRequest) SetOwnerId(v int64) *DescribeInstancesFullStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeInstancesFullStatusRequest) SetPageNumber(v int32) *DescribeInstancesFullStatusRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstancesFullStatusRequest) SetPageSize(v int32) *DescribeInstancesFullStatusRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstancesFullStatusRequest) SetRegionId(v string) *DescribeInstancesFullStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstancesFullStatusRequest) SetResourceOwnerAccount(v string) *DescribeInstancesFullStatusRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeInstancesFullStatusRequest) SetResourceOwnerId(v int64) *DescribeInstancesFullStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeInstancesFullStatusRequest) SetStatus(v string) *DescribeInstancesFullStatusRequest {
	s.Status = &v
	return s
}

func (s *DescribeInstancesFullStatusRequest) Validate() error {
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
	return nil
}

type DescribeInstancesFullStatusRequestEventPublishTime struct {
	// The end of the time range during which system events are published. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2017-12-07T00:00:00Z
	End *string `json:"End,omitempty" xml:"End,omitempty"`
	// The beginning of the time range during which system events are published. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2017-11-30T00:00:00Z
	Start *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s DescribeInstancesFullStatusRequestEventPublishTime) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesFullStatusRequestEventPublishTime) GoString() string {
	return s.String()
}

func (s *DescribeInstancesFullStatusRequestEventPublishTime) GetEnd() *string {
	return s.End
}

func (s *DescribeInstancesFullStatusRequestEventPublishTime) GetStart() *string {
	return s.Start
}

func (s *DescribeInstancesFullStatusRequestEventPublishTime) SetEnd(v string) *DescribeInstancesFullStatusRequestEventPublishTime {
	s.End = &v
	return s
}

func (s *DescribeInstancesFullStatusRequestEventPublishTime) SetStart(v string) *DescribeInstancesFullStatusRequestEventPublishTime {
	s.Start = &v
	return s
}

func (s *DescribeInstancesFullStatusRequestEventPublishTime) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesFullStatusRequestNotBefore struct {
	// The end of the time range during which O\\&M tasks related to scheduled system events are executed. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2017-11-30T00:00:00Z
	End *string `json:"End,omitempty" xml:"End,omitempty"`
	// The beginning of the time range during which O\\&M tasks related to scheduled system events are executed. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2017-12-07T00:00:00Z
	Start *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s DescribeInstancesFullStatusRequestNotBefore) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesFullStatusRequestNotBefore) GoString() string {
	return s.String()
}

func (s *DescribeInstancesFullStatusRequestNotBefore) GetEnd() *string {
	return s.End
}

func (s *DescribeInstancesFullStatusRequestNotBefore) GetStart() *string {
	return s.Start
}

func (s *DescribeInstancesFullStatusRequestNotBefore) SetEnd(v string) *DescribeInstancesFullStatusRequestNotBefore {
	s.End = &v
	return s
}

func (s *DescribeInstancesFullStatusRequestNotBefore) SetStart(v string) *DescribeInstancesFullStatusRequestNotBefore {
	s.Start = &v
	return s
}

func (s *DescribeInstancesFullStatusRequestNotBefore) Validate() error {
	return dara.Validate(s)
}
