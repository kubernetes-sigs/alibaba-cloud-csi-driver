// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceHistoryEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceSystemEventSet(v *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSet) *DescribeInstanceHistoryEventsResponseBody
	GetInstanceSystemEventSet() *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSet
	SetNextToken(v string) *DescribeInstanceHistoryEventsResponseBody
	GetNextToken() *string
	SetPageNumber(v int32) *DescribeInstanceHistoryEventsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeInstanceHistoryEventsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeInstanceHistoryEventsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeInstanceHistoryEventsResponseBody
	GetTotalCount() *int32
}

type DescribeInstanceHistoryEventsResponseBody struct {
	// Details about the instance system events.
	InstanceSystemEventSet *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSet `json:"InstanceSystemEventSet,omitempty" xml:"InstanceSystemEventSet,omitempty" type:"Struct"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// >  If the NextToken parameter is not returned when you use the MaxResults and NextToken parameters to perform a paged query, no more data is returned.
	//
	// example:
	//
	// f1c9fa9de5752***
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number.
	//
	// >
	//
	// 	- If MaxResults and NextToken are used to query results by page, ignore this parameter.
	//
	// 	- This parameter will be removed in the future. We recommend that you use the NextToken and MaxResults parameters for a paged query.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// >
	//
	// 	- If MaxResults and NextToken are used to query results by page, ignore this parameter.
	//
	// 	- This parameter will be removed in the future. We recommend that you use the NextToken and MaxResults parameters for a paged query.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of instances.
	//
	// >  If you specify the MaxResults and NextToken request parameters to perform a paged query, the value of the TotalCount response parameter is invalid.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeInstanceHistoryEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceHistoryEventsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceHistoryEventsResponseBody) GetInstanceSystemEventSet() *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSet {
	return s.InstanceSystemEventSet
}

func (s *DescribeInstanceHistoryEventsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeInstanceHistoryEventsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeInstanceHistoryEventsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInstanceHistoryEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceHistoryEventsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeInstanceHistoryEventsResponseBody) SetInstanceSystemEventSet(v *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSet) *DescribeInstanceHistoryEventsResponseBody {
	s.InstanceSystemEventSet = v
	return s
}

func (s *DescribeInstanceHistoryEventsResponseBody) SetNextToken(v string) *DescribeInstanceHistoryEventsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeInstanceHistoryEventsResponseBody) SetPageNumber(v int32) *DescribeInstanceHistoryEventsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstanceHistoryEventsResponseBody) SetPageSize(v int32) *DescribeInstanceHistoryEventsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeInstanceHistoryEventsResponseBody) SetRequestId(v string) *DescribeInstanceHistoryEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceHistoryEventsResponseBody) SetTotalCount(v int32) *DescribeInstanceHistoryEventsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeInstanceHistoryEventsResponseBody) Validate() error {
	if s.InstanceSystemEventSet != nil {
		if err := s.InstanceSystemEventSet.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSet struct {
	InstanceSystemEventType []*DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventType `json:"InstanceSystemEventType,omitempty" xml:"InstanceSystemEventType,omitempty" type:"Repeated"`
}

func (s DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSet) GoString() string {
	return s.String()
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSet) GetInstanceSystemEventType() []*DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventType {
	return s.InstanceSystemEventType
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSet) SetInstanceSystemEventType(v []*DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventType) *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSet {
	s.InstanceSystemEventType = v
	return s
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSet) Validate() error {
	if s.InstanceSystemEventType != nil {
		for _, item := range s.InstanceSystemEventType {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventType struct {
	// The lifecycle status of the system event.
	EventCycleStatus *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeEventCycleStatus `json:"EventCycleStatus,omitempty" xml:"EventCycleStatus,omitempty" type:"Struct"`
	// The time when the system event ended. The time follows the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-12-01T06:35:31Z
	EventFinishTime *string `json:"EventFinishTime,omitempty" xml:"EventFinishTime,omitempty"`
	// The ID of the system event.
	//
	// example:
	//
	// e-uf64yvznlao4jl2c****
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The time when the system event was published. The time follows the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-11-30T06:32:31Z
	EventPublishTime *string `json:"EventPublishTime,omitempty" xml:"EventPublishTime,omitempty"`
	// The type of the system event.
	EventType *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeEventType `json:"EventType,omitempty" xml:"EventType,omitempty" type:"Struct"`
	// The extended attribute of the system event.
	ExtendedAttribute *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttribute `json:"ExtendedAttribute,omitempty" xml:"ExtendedAttribute,omitempty" type:"Struct"`
	// The impact level of the system event.
	//
	// example:
	//
	// 100
	ImpactLevel *string `json:"ImpactLevel,omitempty" xml:"ImpactLevel,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// i-uf678mass4zvr9n1****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The scheduled start time of the system event. The time follows the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-12-06T00:00:00Z
	NotBefore *string `json:"NotBefore,omitempty" xml:"NotBefore,omitempty"`
	// The reason why the system event occurred.
	//
	// example:
	//
	// System maintenance is scheduled due to ***.
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The reason code category for the system event.
	//
	// example:
	//
	// VPCMigrationEcs
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// The type of the resource. Valid values:
	//
	// 	- instance: ECS instance
	//
	// 	- ddh: dedicated host
	//
	// 	- managehost: physical machine in a smart hosting pool
	//
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventType) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventType) GoString() string {
	return s.String()
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventType) GetEventCycleStatus() *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeEventCycleStatus {
	return s.EventCycleStatus
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventType) GetEventFinishTime() *string {
	return s.EventFinishTime
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventType) GetEventId() *string {
	return s.EventId
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventType) GetEventPublishTime() *string {
	return s.EventPublishTime
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventType) GetEventType() *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeEventType {
	return s.EventType
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventType) GetExtendedAttribute() *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttribute {
	return s.ExtendedAttribute
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventType) GetImpactLevel() *string {
	return s.ImpactLevel
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventType) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventType) GetNotBefore() *string {
	return s.NotBefore
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventType) GetReason() *string {
	return s.Reason
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventType) GetReasonCode() *string {
	return s.ReasonCode
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventType) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventType) SetEventCycleStatus(v *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeEventCycleStatus) *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventType {
	s.EventCycleStatus = v
	return s
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventType) SetEventFinishTime(v string) *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventType {
	s.EventFinishTime = &v
	return s
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventType) SetEventId(v string) *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventType {
	s.EventId = &v
	return s
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventType) SetEventPublishTime(v string) *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventType {
	s.EventPublishTime = &v
	return s
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventType) SetEventType(v *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeEventType) *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventType {
	s.EventType = v
	return s
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventType) SetExtendedAttribute(v *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttribute) *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventType {
	s.ExtendedAttribute = v
	return s
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventType) SetImpactLevel(v string) *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventType {
	s.ImpactLevel = &v
	return s
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventType) SetInstanceId(v string) *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventType {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventType) SetNotBefore(v string) *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventType {
	s.NotBefore = &v
	return s
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventType) SetReason(v string) *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventType {
	s.Reason = &v
	return s
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventType) SetReasonCode(v string) *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventType {
	s.ReasonCode = &v
	return s
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventType) SetResourceType(v string) *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventType {
	s.ResourceType = &v
	return s
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventType) Validate() error {
	if s.EventCycleStatus != nil {
		if err := s.EventCycleStatus.Validate(); err != nil {
			return err
		}
	}
	if s.EventType != nil {
		if err := s.EventType.Validate(); err != nil {
			return err
		}
	}
	if s.ExtendedAttribute != nil {
		if err := s.ExtendedAttribute.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeEventCycleStatus struct {
	// The state code of the system event.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The state name of the system event.
	//
	// example:
	//
	// Executed
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeEventCycleStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeEventCycleStatus) GoString() string {
	return s.String()
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeEventCycleStatus) GetCode() *int32 {
	return s.Code
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeEventCycleStatus) GetName() *string {
	return s.Name
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeEventCycleStatus) SetCode(v int32) *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeEventCycleStatus {
	s.Code = &v
	return s
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeEventCycleStatus) SetName(v string) *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeEventCycleStatus {
	s.Name = &v
	return s
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeEventCycleStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeEventType struct {
	// The code of the system event type.
	//
	// example:
	//
	// 34
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The name of the system event type.
	//
	// example:
	//
	// InstanceExpiration.Stop
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeEventType) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeEventType) GoString() string {
	return s.String()
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeEventType) GetCode() *int32 {
	return s.Code
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeEventType) GetName() *string {
	return s.Name
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeEventType) SetCode(v int32) *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeEventType {
	s.Code = &v
	return s
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeEventType) SetName(v string) *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeEventType {
	s.Name = &v
	return s
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeEventType) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttribute struct {
	// Indicates whether the event can be handled.
	//
	// example:
	//
	// true
	CanAccept *string `json:"CanAccept,omitempty" xml:"CanAccept,omitempty"`
	// The code of the security violation.
	//
	// example:
	//
	// PR111
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The device name of the local disk.
	//
	// example:
	//
	// /dev/vda
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The ID of the local disk.
	//
	// example:
	//
	// d-diskid1
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The ID of the host.
	//
	// example:
	//
	// dh-bp1ewce1gk3iwv2****
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// The type of the host. Valid values:
	//
	// 	- ddh: dedicated host
	//
	// 	- managehost: physical machine in a smart hosting pool
	//
	// example:
	//
	// ddh
	HostType *string `json:"HostType,omitempty" xml:"HostType,omitempty"`
	// The inactive disks that were released and whose data must be cleared.
	InactiveDisks *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttributeInactiveDisks `json:"InactiveDisks,omitempty" xml:"InactiveDisks,omitempty" type:"Struct"`
	MetricName    *string                                                                                                               `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	MetricValue   *string                                                                                                               `json:"MetricValue,omitempty" xml:"MetricValue,omitempty"`
	// The migration solution of the instance. Valid value: MigrationPlan. Instances can be migrated only by using migration plans.
	MigrationOptions *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttributeMigrationOptions `json:"MigrationOptions,omitempty" xml:"MigrationOptions,omitempty" type:"Struct"`
	// The online repair policy for the damaged disk. Valid value: IsolateOnly, which indicates that damaged disks are isolated but not repaired.
	//
	// example:
	//
	// IsolateOnly
	OnlineRepairPolicy *string `json:"OnlineRepairPolicy,omitempty" xml:"OnlineRepairPolicy,omitempty"`
	// The illegal domain name.
	//
	// example:
	//
	// 1228.test.com
	PunishDomain *string `json:"PunishDomain,omitempty" xml:"PunishDomain,omitempty"`
	// The type of the penalty.
	//
	// example:
	//
	// ecs_message_alert
	PunishType *string `json:"PunishType,omitempty" xml:"PunishType,omitempty"`
	// The illegal URL.
	//
	// example:
	//
	// http://1228.test.com/1
	PunishUrl *string `json:"PunishUrl,omitempty" xml:"PunishUrl,omitempty"`
	// The rack number of the cloud box.
	//
	// example:
	//
	// A01
	Rack *string `json:"Rack,omitempty" xml:"Rack,omitempty"`
	// The response result of the event. Valid values:
	//
	// 	- true: The event was handled.
	//
	// 	- false: The event failed to be handled.
	//
	// example:
	//
	// true
	ResponseResult *string `json:"ResponseResult,omitempty" xml:"ResponseResult,omitempty"`
}

func (s DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttribute) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttribute) GoString() string {
	return s.String()
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttribute) GetCanAccept() *string {
	return s.CanAccept
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttribute) GetCode() *string {
	return s.Code
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttribute) GetDevice() *string {
	return s.Device
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttribute) GetDiskId() *string {
	return s.DiskId
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttribute) GetHostId() *string {
	return s.HostId
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttribute) GetHostType() *string {
	return s.HostType
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttribute) GetInactiveDisks() *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttributeInactiveDisks {
	return s.InactiveDisks
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttribute) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttribute) GetMetricValue() *string {
	return s.MetricValue
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttribute) GetMigrationOptions() *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttributeMigrationOptions {
	return s.MigrationOptions
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttribute) GetOnlineRepairPolicy() *string {
	return s.OnlineRepairPolicy
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttribute) GetPunishDomain() *string {
	return s.PunishDomain
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttribute) GetPunishType() *string {
	return s.PunishType
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttribute) GetPunishUrl() *string {
	return s.PunishUrl
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttribute) GetRack() *string {
	return s.Rack
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttribute) GetResponseResult() *string {
	return s.ResponseResult
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttribute) SetCanAccept(v string) *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttribute {
	s.CanAccept = &v
	return s
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttribute) SetCode(v string) *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttribute {
	s.Code = &v
	return s
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttribute) SetDevice(v string) *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttribute {
	s.Device = &v
	return s
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttribute) SetDiskId(v string) *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttribute {
	s.DiskId = &v
	return s
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttribute) SetHostId(v string) *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttribute {
	s.HostId = &v
	return s
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttribute) SetHostType(v string) *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttribute {
	s.HostType = &v
	return s
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttribute) SetInactiveDisks(v *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttributeInactiveDisks) *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttribute {
	s.InactiveDisks = v
	return s
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttribute) SetMetricName(v string) *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttribute {
	s.MetricName = &v
	return s
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttribute) SetMetricValue(v string) *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttribute {
	s.MetricValue = &v
	return s
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttribute) SetMigrationOptions(v *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttributeMigrationOptions) *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttribute {
	s.MigrationOptions = v
	return s
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttribute) SetOnlineRepairPolicy(v string) *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttribute {
	s.OnlineRepairPolicy = &v
	return s
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttribute) SetPunishDomain(v string) *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttribute {
	s.PunishDomain = &v
	return s
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttribute) SetPunishType(v string) *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttribute {
	s.PunishType = &v
	return s
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttribute) SetPunishUrl(v string) *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttribute {
	s.PunishUrl = &v
	return s
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttribute) SetRack(v string) *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttribute {
	s.Rack = &v
	return s
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttribute) SetResponseResult(v string) *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttribute {
	s.ResponseResult = &v
	return s
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttribute) Validate() error {
	if s.InactiveDisks != nil {
		if err := s.InactiveDisks.Validate(); err != nil {
			return err
		}
	}
	if s.MigrationOptions != nil {
		if err := s.MigrationOptions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttributeInactiveDisks struct {
	InactiveDisk []*DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttributeInactiveDisksInactiveDisk `json:"InactiveDisk,omitempty" xml:"InactiveDisk,omitempty" type:"Repeated"`
}

func (s DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttributeInactiveDisks) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttributeInactiveDisks) GoString() string {
	return s.String()
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttributeInactiveDisks) GetInactiveDisk() []*DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttributeInactiveDisksInactiveDisk {
	return s.InactiveDisk
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttributeInactiveDisks) SetInactiveDisk(v []*DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttributeInactiveDisksInactiveDisk) *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttributeInactiveDisks {
	s.InactiveDisk = v
	return s
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttributeInactiveDisks) Validate() error {
	if s.InactiveDisk != nil {
		for _, item := range s.InactiveDisk {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttributeInactiveDisksInactiveDisk struct {
	// The time when the disk was created. The time follows the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-11-30T06:32:31Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The category of the disk. Valid values:
	//
	// 	- cloud: basic disk
	//
	// 	- cloud_efficiency: ultra disk
	//
	// 	- cloud_ssd: standard SSD
	//
	// 	- cloud_essd: Enterprise SSD (ESSD)
	//
	// 	- local_ssd_pro: I/O-intensive local disk
	//
	// 	- local_hdd_pro: throughput-intensive local disk
	//
	// 	- ephemeral: retired local disk
	//
	// 	- ephemeral_ssd: retired local SSD
	//
	// example:
	//
	// cloud_efficiency
	DeviceCategory *string `json:"DeviceCategory,omitempty" xml:"DeviceCategory,omitempty"`
	// The size of the disk. Unit: GiB.
	//
	// example:
	//
	// 80
	DeviceSize *string `json:"DeviceSize,omitempty" xml:"DeviceSize,omitempty"`
	// The type of the disk. Valid values:
	//
	// 	- system: system disk
	//
	// 	- data: data disk
	//
	// example:
	//
	// data
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// The time when the disk was released. The time follows the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-11-30T06:32:31Z
	ReleaseTime *string `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
}

func (s DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttributeInactiveDisksInactiveDisk) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttributeInactiveDisksInactiveDisk) GoString() string {
	return s.String()
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttributeInactiveDisksInactiveDisk) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttributeInactiveDisksInactiveDisk) GetDeviceCategory() *string {
	return s.DeviceCategory
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttributeInactiveDisksInactiveDisk) GetDeviceSize() *string {
	return s.DeviceSize
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttributeInactiveDisksInactiveDisk) GetDeviceType() *string {
	return s.DeviceType
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttributeInactiveDisksInactiveDisk) GetReleaseTime() *string {
	return s.ReleaseTime
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttributeInactiveDisksInactiveDisk) SetCreationTime(v string) *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttributeInactiveDisksInactiveDisk {
	s.CreationTime = &v
	return s
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttributeInactiveDisksInactiveDisk) SetDeviceCategory(v string) *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttributeInactiveDisksInactiveDisk {
	s.DeviceCategory = &v
	return s
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttributeInactiveDisksInactiveDisk) SetDeviceSize(v string) *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttributeInactiveDisksInactiveDisk {
	s.DeviceSize = &v
	return s
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttributeInactiveDisksInactiveDisk) SetDeviceType(v string) *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttributeInactiveDisksInactiveDisk {
	s.DeviceType = &v
	return s
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttributeInactiveDisksInactiveDisk) SetReleaseTime(v string) *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttributeInactiveDisksInactiveDisk {
	s.ReleaseTime = &v
	return s
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttributeInactiveDisksInactiveDisk) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttributeMigrationOptions struct {
	MigrationOption []*string `json:"MigrationOption,omitempty" xml:"MigrationOption,omitempty" type:"Repeated"`
}

func (s DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttributeMigrationOptions) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttributeMigrationOptions) GoString() string {
	return s.String()
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttributeMigrationOptions) GetMigrationOption() []*string {
	return s.MigrationOption
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttributeMigrationOptions) SetMigrationOption(v []*string) *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttributeMigrationOptions {
	s.MigrationOption = v
	return s
}

func (s *DescribeInstanceHistoryEventsResponseBodyInstanceSystemEventSetInstanceSystemEventTypeExtendedAttributeMigrationOptions) Validate() error {
	return dara.Validate(s)
}
