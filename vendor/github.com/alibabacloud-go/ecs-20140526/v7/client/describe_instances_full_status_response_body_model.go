// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstancesFullStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceFullStatusSet(v *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSet) *DescribeInstancesFullStatusResponseBody
	GetInstanceFullStatusSet() *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSet
	SetPageNumber(v int32) *DescribeInstancesFullStatusResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeInstancesFullStatusResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeInstancesFullStatusResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeInstancesFullStatusResponseBody
	GetTotalCount() *int32
}

type DescribeInstancesFullStatusResponseBody struct {
	// The queried instances.
	//
	// >  If no instances exist, this parameter is empty.
	InstanceFullStatusSet *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSet `json:"InstanceFullStatusSet,omitempty" xml:"InstanceFullStatusSet,omitempty" type:"Struct"`
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
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeInstancesFullStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesFullStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstancesFullStatusResponseBody) GetInstanceFullStatusSet() *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSet {
	return s.InstanceFullStatusSet
}

func (s *DescribeInstancesFullStatusResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeInstancesFullStatusResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInstancesFullStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstancesFullStatusResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeInstancesFullStatusResponseBody) SetInstanceFullStatusSet(v *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSet) *DescribeInstancesFullStatusResponseBody {
	s.InstanceFullStatusSet = v
	return s
}

func (s *DescribeInstancesFullStatusResponseBody) SetPageNumber(v int32) *DescribeInstancesFullStatusResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstancesFullStatusResponseBody) SetPageSize(v int32) *DescribeInstancesFullStatusResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeInstancesFullStatusResponseBody) SetRequestId(v string) *DescribeInstancesFullStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstancesFullStatusResponseBody) SetTotalCount(v int32) *DescribeInstancesFullStatusResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeInstancesFullStatusResponseBody) Validate() error {
	if s.InstanceFullStatusSet != nil {
		if err := s.InstanceFullStatusSet.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstancesFullStatusResponseBodyInstanceFullStatusSet struct {
	InstanceFullStatusType []*DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusType `json:"InstanceFullStatusType,omitempty" xml:"InstanceFullStatusType,omitempty" type:"Repeated"`
}

func (s DescribeInstancesFullStatusResponseBodyInstanceFullStatusSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesFullStatusResponseBodyInstanceFullStatusSet) GoString() string {
	return s.String()
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSet) GetInstanceFullStatusType() []*DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusType {
	return s.InstanceFullStatusType
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSet) SetInstanceFullStatusType(v []*DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusType) *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSet {
	s.InstanceFullStatusType = v
	return s
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSet) Validate() error {
	if s.InstanceFullStatusType != nil {
		for _, item := range s.InstanceFullStatusType {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusType struct {
	// The health state of the instance.
	HealthStatus *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeHealthStatus `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty" type:"Struct"`
	// The instance ID.
	//
	// example:
	//
	// i-bp67acfmxazb4p****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The system events that are in the Scheduled or Inquiring state.
	ScheduledSystemEventSet *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSet `json:"ScheduledSystemEventSet,omitempty" xml:"ScheduledSystemEventSet,omitempty" type:"Struct"`
	// The lifecycle state of the instance.
	Status *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeStatus `json:"Status,omitempty" xml:"Status,omitempty" type:"Struct"`
}

func (s DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusType) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusType) GoString() string {
	return s.String()
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusType) GetHealthStatus() *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeHealthStatus {
	return s.HealthStatus
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusType) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusType) GetScheduledSystemEventSet() *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSet {
	return s.ScheduledSystemEventSet
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusType) GetStatus() *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeStatus {
	return s.Status
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusType) SetHealthStatus(v *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeHealthStatus) *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusType {
	s.HealthStatus = v
	return s
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusType) SetInstanceId(v string) *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusType {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusType) SetScheduledSystemEventSet(v *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSet) *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusType {
	s.ScheduledSystemEventSet = v
	return s
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusType) SetStatus(v *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeStatus) *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusType {
	s.Status = v
	return s
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusType) Validate() error {
	if s.HealthStatus != nil {
		if err := s.HealthStatus.Validate(); err != nil {
			return err
		}
	}
	if s.ScheduledSystemEventSet != nil {
		if err := s.ScheduledSystemEventSet.Validate(); err != nil {
			return err
		}
	}
	if s.Status != nil {
		if err := s.Status.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeHealthStatus struct {
	// The code of the health state.
	//
	// example:
	//
	// 64
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The name of the health state.
	//
	// example:
	//
	// Warning
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeHealthStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeHealthStatus) GoString() string {
	return s.String()
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeHealthStatus) GetCode() *int32 {
	return s.Code
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeHealthStatus) GetName() *string {
	return s.Name
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeHealthStatus) SetCode(v int32) *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeHealthStatus {
	s.Code = &v
	return s
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeHealthStatus) SetName(v string) *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeHealthStatus {
	s.Name = &v
	return s
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeHealthStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSet struct {
	ScheduledSystemEventType []*DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventType `json:"ScheduledSystemEventType,omitempty" xml:"ScheduledSystemEventType,omitempty" type:"Repeated"`
}

func (s DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSet) GoString() string {
	return s.String()
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSet) GetScheduledSystemEventType() []*DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventType {
	return s.ScheduledSystemEventType
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSet) SetScheduledSystemEventType(v []*DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventType) *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSet {
	s.ScheduledSystemEventType = v
	return s
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSet) Validate() error {
	if s.ScheduledSystemEventType != nil {
		for _, item := range s.ScheduledSystemEventType {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventType struct {
	// The state of the system event.
	EventCycleStatus *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeEventCycleStatus `json:"EventCycleStatus,omitempty" xml:"EventCycleStatus,omitempty" type:"Struct"`
	// The system event ID.
	//
	// example:
	//
	// e-bp1hygp5b04o56l0****
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The time when the system event was published. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-11-30T06:32:31Z
	EventPublishTime *string `json:"EventPublishTime,omitempty" xml:"EventPublishTime,omitempty"`
	// The type of the system event.
	EventType *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeEventType `json:"EventType,omitempty" xml:"EventType,omitempty" type:"Struct"`
	// The extended attributes of system events generated for instances that have local disks attached.
	//
	// The return values vary based on the system event type.
	//
	// If the system event type is not one of the following types, this parameter is empty:
	//
	// 	- SystemMaintenance.StopAndRepair
	//
	// 	- SystemMaintenance.CleanInactiveDisks
	//
	// 	- SecurityPunish.Locked
	//
	// 	- SecurityPunish.WebsiteBanned
	//
	// 	- SystemUpgrade.Migrate
	//
	// 	- SystemMaintenance.RebootAndIsolateErrorDisk
	//
	// 	- SystemMaintenance.RebootAndReInitErrorDisk
	//
	// 	- SystemMaintenance.ReInitErrorDisk
	//
	// 	- SystemMaintenance.IsolateErrorDisk
	ExtendedAttribute *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeExtendedAttribute `json:"ExtendedAttribute,omitempty" xml:"ExtendedAttribute,omitempty" type:"Struct"`
	// The impact level of the system event.
	//
	// >  If the user is not in a whitelist, this parameter is empty.
	//
	// example:
	//
	// 100
	ImpactLevel *string `json:"ImpactLevel,omitempty" xml:"ImpactLevel,omitempty"`
	// The scheduled time at which to execute the O\\&M task related to the system event. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-12-07T00:00:00Z
	NotBefore *string `json:"NotBefore,omitempty" xml:"NotBefore,omitempty"`
	// The reason why the system event was scheduled.
	//
	// >  If the exception cause is not detected, this parameter is empty.
	//
	// example:
	//
	// A simulated event.
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventType) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventType) GoString() string {
	return s.String()
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventType) GetEventCycleStatus() *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeEventCycleStatus {
	return s.EventCycleStatus
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventType) GetEventId() *string {
	return s.EventId
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventType) GetEventPublishTime() *string {
	return s.EventPublishTime
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventType) GetEventType() *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeEventType {
	return s.EventType
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventType) GetExtendedAttribute() *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeExtendedAttribute {
	return s.ExtendedAttribute
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventType) GetImpactLevel() *string {
	return s.ImpactLevel
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventType) GetNotBefore() *string {
	return s.NotBefore
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventType) GetReason() *string {
	return s.Reason
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventType) SetEventCycleStatus(v *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeEventCycleStatus) *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventType {
	s.EventCycleStatus = v
	return s
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventType) SetEventId(v string) *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventType {
	s.EventId = &v
	return s
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventType) SetEventPublishTime(v string) *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventType {
	s.EventPublishTime = &v
	return s
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventType) SetEventType(v *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeEventType) *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventType {
	s.EventType = v
	return s
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventType) SetExtendedAttribute(v *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeExtendedAttribute) *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventType {
	s.ExtendedAttribute = v
	return s
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventType) SetImpactLevel(v string) *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventType {
	s.ImpactLevel = &v
	return s
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventType) SetNotBefore(v string) *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventType {
	s.NotBefore = &v
	return s
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventType) SetReason(v string) *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventType {
	s.Reason = &v
	return s
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventType) Validate() error {
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

type DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeEventCycleStatus struct {
	// The code of the system event state.
	//
	// example:
	//
	// 24
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The name of the system event state.
	//
	// example:
	//
	// Scheduled
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeEventCycleStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeEventCycleStatus) GoString() string {
	return s.String()
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeEventCycleStatus) GetCode() *int32 {
	return s.Code
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeEventCycleStatus) GetName() *string {
	return s.Name
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeEventCycleStatus) SetCode(v int32) *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeEventCycleStatus {
	s.Code = &v
	return s
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeEventCycleStatus) SetName(v string) *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeEventCycleStatus {
	s.Name = &v
	return s
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeEventCycleStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeEventType struct {
	// The code of the system event type.
	//
	// example:
	//
	// 1
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The name of the system event type.
	//
	// example:
	//
	// SystemMaintenance.Reboot
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeEventType) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeEventType) GoString() string {
	return s.String()
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeEventType) GetCode() *int32 {
	return s.Code
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeEventType) GetName() *string {
	return s.Name
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeEventType) SetCode(v int32) *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeEventType {
	s.Code = &v
	return s
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeEventType) SetName(v string) *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeEventType {
	s.Name = &v
	return s
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeEventType) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeExtendedAttribute struct {
	// The device name of the local disk.
	//
	// example:
	//
	// /dev/vdb
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The ID of the local disk.
	//
	// example:
	//
	// d-bp67acfmxazb4p****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The inactive disks that have been released and must be cleared.
	InactiveDisks *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeExtendedAttributeInactiveDisks `json:"InactiveDisks,omitempty" xml:"InactiveDisks,omitempty" type:"Struct"`
}

func (s DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeExtendedAttribute) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeExtendedAttribute) GoString() string {
	return s.String()
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeExtendedAttribute) GetDevice() *string {
	return s.Device
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeExtendedAttribute) GetDiskId() *string {
	return s.DiskId
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeExtendedAttribute) GetInactiveDisks() *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeExtendedAttributeInactiveDisks {
	return s.InactiveDisks
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeExtendedAttribute) SetDevice(v string) *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeExtendedAttribute {
	s.Device = &v
	return s
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeExtendedAttribute) SetDiskId(v string) *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeExtendedAttribute {
	s.DiskId = &v
	return s
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeExtendedAttribute) SetInactiveDisks(v *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeExtendedAttributeInactiveDisks) *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeExtendedAttribute {
	s.InactiveDisks = v
	return s
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeExtendedAttribute) Validate() error {
	if s.InactiveDisks != nil {
		if err := s.InactiveDisks.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeExtendedAttributeInactiveDisks struct {
	InactiveDisk []*DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeExtendedAttributeInactiveDisksInactiveDisk `json:"InactiveDisk,omitempty" xml:"InactiveDisk,omitempty" type:"Repeated"`
}

func (s DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeExtendedAttributeInactiveDisks) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeExtendedAttributeInactiveDisks) GoString() string {
	return s.String()
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeExtendedAttributeInactiveDisks) GetInactiveDisk() []*DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeExtendedAttributeInactiveDisksInactiveDisk {
	return s.InactiveDisk
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeExtendedAttributeInactiveDisks) SetInactiveDisk(v []*DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeExtendedAttributeInactiveDisksInactiveDisk) *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeExtendedAttributeInactiveDisks {
	s.InactiveDisk = v
	return s
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeExtendedAttributeInactiveDisks) Validate() error {
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

type DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeExtendedAttributeInactiveDisksInactiveDisk struct {
	// The time when the disk was created. The time follows the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-07-27T13:53:25Z
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
	// cloud_ssd
	DeviceCategory *string `json:"DeviceCategory,omitempty" xml:"DeviceCategory,omitempty"`
	// The size of the disk. Unit: GiB.
	//
	// example:
	//
	// 80
	DeviceSize *string `json:"DeviceSize,omitempty" xml:"DeviceSize,omitempty"`
	// The type of the disk. Valid values:
	//
	// 	- system
	//
	// 	- data
	//
	// example:
	//
	// system
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// The time when the disk was released. The time follows the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-07-27T13:53:25Z
	ReleaseTime *string `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
}

func (s DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeExtendedAttributeInactiveDisksInactiveDisk) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeExtendedAttributeInactiveDisksInactiveDisk) GoString() string {
	return s.String()
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeExtendedAttributeInactiveDisksInactiveDisk) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeExtendedAttributeInactiveDisksInactiveDisk) GetDeviceCategory() *string {
	return s.DeviceCategory
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeExtendedAttributeInactiveDisksInactiveDisk) GetDeviceSize() *string {
	return s.DeviceSize
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeExtendedAttributeInactiveDisksInactiveDisk) GetDeviceType() *string {
	return s.DeviceType
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeExtendedAttributeInactiveDisksInactiveDisk) GetReleaseTime() *string {
	return s.ReleaseTime
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeExtendedAttributeInactiveDisksInactiveDisk) SetCreationTime(v string) *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeExtendedAttributeInactiveDisksInactiveDisk {
	s.CreationTime = &v
	return s
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeExtendedAttributeInactiveDisksInactiveDisk) SetDeviceCategory(v string) *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeExtendedAttributeInactiveDisksInactiveDisk {
	s.DeviceCategory = &v
	return s
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeExtendedAttributeInactiveDisksInactiveDisk) SetDeviceSize(v string) *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeExtendedAttributeInactiveDisksInactiveDisk {
	s.DeviceSize = &v
	return s
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeExtendedAttributeInactiveDisksInactiveDisk) SetDeviceType(v string) *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeExtendedAttributeInactiveDisksInactiveDisk {
	s.DeviceType = &v
	return s
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeExtendedAttributeInactiveDisksInactiveDisk) SetReleaseTime(v string) *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeExtendedAttributeInactiveDisksInactiveDisk {
	s.ReleaseTime = &v
	return s
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeScheduledSystemEventSetScheduledSystemEventTypeExtendedAttributeInactiveDisksInactiveDisk) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeStatus struct {
	// The code of the instance lifecycle state.
	//
	// example:
	//
	// 1
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The name of the instance lifecycle state.
	//
	// example:
	//
	// Running
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeStatus) GoString() string {
	return s.String()
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeStatus) GetCode() *int32 {
	return s.Code
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeStatus) GetName() *string {
	return s.Name
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeStatus) SetCode(v int32) *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeStatus {
	s.Code = &v
	return s
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeStatus) SetName(v string) *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeStatus {
	s.Name = &v
	return s
}

func (s *DescribeInstancesFullStatusResponseBodyInstanceFullStatusSetInstanceFullStatusTypeStatus) Validate() error {
	return dara.Validate(s)
}
