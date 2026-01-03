// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDisksFullStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDiskFullStatusSet(v *DescribeDisksFullStatusResponseBodyDiskFullStatusSet) *DescribeDisksFullStatusResponseBody
	GetDiskFullStatusSet() *DescribeDisksFullStatusResponseBodyDiskFullStatusSet
	SetPageNumber(v int32) *DescribeDisksFullStatusResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDisksFullStatusResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeDisksFullStatusResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeDisksFullStatusResponseBody
	GetTotalCount() *int32
}

type DescribeDisksFullStatusResponseBody struct {
	// The collection of full status information of the EBS devices.
	DiskFullStatusSet *DescribeDisksFullStatusResponseBodyDiskFullStatusSet `json:"DiskFullStatusSet,omitempty" xml:"DiskFullStatusSet,omitempty" type:"Struct"`
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
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of EBS devices for which full status information is returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDisksFullStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDisksFullStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDisksFullStatusResponseBody) GetDiskFullStatusSet() *DescribeDisksFullStatusResponseBodyDiskFullStatusSet {
	return s.DiskFullStatusSet
}

func (s *DescribeDisksFullStatusResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDisksFullStatusResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDisksFullStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDisksFullStatusResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDisksFullStatusResponseBody) SetDiskFullStatusSet(v *DescribeDisksFullStatusResponseBodyDiskFullStatusSet) *DescribeDisksFullStatusResponseBody {
	s.DiskFullStatusSet = v
	return s
}

func (s *DescribeDisksFullStatusResponseBody) SetPageNumber(v int32) *DescribeDisksFullStatusResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDisksFullStatusResponseBody) SetPageSize(v int32) *DescribeDisksFullStatusResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDisksFullStatusResponseBody) SetRequestId(v string) *DescribeDisksFullStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDisksFullStatusResponseBody) SetTotalCount(v int32) *DescribeDisksFullStatusResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDisksFullStatusResponseBody) Validate() error {
	if s.DiskFullStatusSet != nil {
		if err := s.DiskFullStatusSet.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDisksFullStatusResponseBodyDiskFullStatusSet struct {
	DiskFullStatusType []*DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusType `json:"DiskFullStatusType,omitempty" xml:"DiskFullStatusType,omitempty" type:"Repeated"`
}

func (s DescribeDisksFullStatusResponseBodyDiskFullStatusSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeDisksFullStatusResponseBodyDiskFullStatusSet) GoString() string {
	return s.String()
}

func (s *DescribeDisksFullStatusResponseBodyDiskFullStatusSet) GetDiskFullStatusType() []*DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusType {
	return s.DiskFullStatusType
}

func (s *DescribeDisksFullStatusResponseBodyDiskFullStatusSet) SetDiskFullStatusType(v []*DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusType) *DescribeDisksFullStatusResponseBodyDiskFullStatusSet {
	s.DiskFullStatusType = v
	return s
}

func (s *DescribeDisksFullStatusResponseBodyDiskFullStatusSet) Validate() error {
	if s.DiskFullStatusType != nil {
		for _, item := range s.DiskFullStatusType {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusType struct {
	// The name of the EBS device that is attached to an instance. Example: /dev/xvdb.
	//
	// This parameter has a value only when the value of `Status` is `In_use`.
	//
	// > This parameter will be deprecated in the future. To ensure future compatibility, we recommend that you do not use this parameter.
	//
	// example:
	//
	// null
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The events about the EBS device.
	DiskEventSet *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeDiskEventSet `json:"DiskEventSet,omitempty" xml:"DiskEventSet,omitempty" type:"Struct"`
	// The EBS device ID.
	//
	// example:
	//
	// d-bp67acfmxazb4p****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The health status of the EBS device.
	HealthStatus *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeHealthStatus `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty" type:"Struct"`
	// The instance ID.
	//
	// example:
	//
	// i-bp67acfmxazb4p****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The lifecycle status of the EBS device.
	Status *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeStatus `json:"Status,omitempty" xml:"Status,omitempty" type:"Struct"`
}

func (s DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusType) String() string {
	return dara.Prettify(s)
}

func (s DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusType) GoString() string {
	return s.String()
}

func (s *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusType) GetDevice() *string {
	return s.Device
}

func (s *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusType) GetDiskEventSet() *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeDiskEventSet {
	return s.DiskEventSet
}

func (s *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusType) GetDiskId() *string {
	return s.DiskId
}

func (s *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusType) GetHealthStatus() *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeHealthStatus {
	return s.HealthStatus
}

func (s *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusType) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusType) GetStatus() *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeStatus {
	return s.Status
}

func (s *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusType) SetDevice(v string) *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusType {
	s.Device = &v
	return s
}

func (s *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusType) SetDiskEventSet(v *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeDiskEventSet) *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusType {
	s.DiskEventSet = v
	return s
}

func (s *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusType) SetDiskId(v string) *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusType {
	s.DiskId = &v
	return s
}

func (s *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusType) SetHealthStatus(v *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeHealthStatus) *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusType {
	s.HealthStatus = v
	return s
}

func (s *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusType) SetInstanceId(v string) *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusType {
	s.InstanceId = &v
	return s
}

func (s *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusType) SetStatus(v *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeStatus) *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusType {
	s.Status = v
	return s
}

func (s *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusType) Validate() error {
	if s.DiskEventSet != nil {
		if err := s.DiskEventSet.Validate(); err != nil {
			return err
		}
	}
	if s.HealthStatus != nil {
		if err := s.HealthStatus.Validate(); err != nil {
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

type DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeDiskEventSet struct {
	DiskEventType []*DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeDiskEventSetDiskEventType `json:"DiskEventType,omitempty" xml:"DiskEventType,omitempty" type:"Repeated"`
}

func (s DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeDiskEventSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeDiskEventSet) GoString() string {
	return s.String()
}

func (s *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeDiskEventSet) GetDiskEventType() []*DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeDiskEventSetDiskEventType {
	return s.DiskEventType
}

func (s *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeDiskEventSet) SetDiskEventType(v []*DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeDiskEventSetDiskEventType) *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeDiskEventSet {
	s.DiskEventType = v
	return s
}

func (s *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeDiskEventSet) Validate() error {
	if s.DiskEventType != nil {
		for _, item := range s.DiskEventType {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeDiskEventSetDiskEventType struct {
	// The time when the event ended.
	//
	// example:
	//
	// 2018-05-06T02:48:52Z
	EventEndTime *string `json:"EventEndTime,omitempty" xml:"EventEndTime,omitempty"`
	// The ID of the event.
	//
	// example:
	//
	// e-bp67acfmxazb4p****
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The time when the event occurred.
	//
	// example:
	//
	// 2018-05-08T02:43:10Z
	EventTime *string `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
	// The type of the event.
	EventType *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeDiskEventSetDiskEventTypeEventType `json:"EventType,omitempty" xml:"EventType,omitempty" type:"Struct"`
	// The impact level of the event.
	//
	// example:
	//
	// 100
	ImpactLevel *string `json:"ImpactLevel,omitempty" xml:"ImpactLevel,omitempty"`
}

func (s DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeDiskEventSetDiskEventType) String() string {
	return dara.Prettify(s)
}

func (s DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeDiskEventSetDiskEventType) GoString() string {
	return s.String()
}

func (s *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeDiskEventSetDiskEventType) GetEventEndTime() *string {
	return s.EventEndTime
}

func (s *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeDiskEventSetDiskEventType) GetEventId() *string {
	return s.EventId
}

func (s *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeDiskEventSetDiskEventType) GetEventTime() *string {
	return s.EventTime
}

func (s *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeDiskEventSetDiskEventType) GetEventType() *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeDiskEventSetDiskEventTypeEventType {
	return s.EventType
}

func (s *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeDiskEventSetDiskEventType) GetImpactLevel() *string {
	return s.ImpactLevel
}

func (s *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeDiskEventSetDiskEventType) SetEventEndTime(v string) *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeDiskEventSetDiskEventType {
	s.EventEndTime = &v
	return s
}

func (s *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeDiskEventSetDiskEventType) SetEventId(v string) *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeDiskEventSetDiskEventType {
	s.EventId = &v
	return s
}

func (s *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeDiskEventSetDiskEventType) SetEventTime(v string) *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeDiskEventSetDiskEventType {
	s.EventTime = &v
	return s
}

func (s *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeDiskEventSetDiskEventType) SetEventType(v *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeDiskEventSetDiskEventTypeEventType) *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeDiskEventSetDiskEventType {
	s.EventType = v
	return s
}

func (s *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeDiskEventSetDiskEventType) SetImpactLevel(v string) *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeDiskEventSetDiskEventType {
	s.ImpactLevel = &v
	return s
}

func (s *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeDiskEventSetDiskEventType) Validate() error {
	if s.EventType != nil {
		if err := s.EventType.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeDiskEventSetDiskEventTypeEventType struct {
	// The code of the event type.
	//
	// example:
	//
	// 7
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The name of the event type. Valid values:
	//
	// 	- Degraded: The performance of the EBS device is degraded.
	//
	// 	- SeverelyDegraded: The performance of the EBS device is severely degraded.
	//
	// 	- Stalled: The performance of the EBS device is severely affected.
	//
	// 	- ErrorDetected: The local disk is damaged.
	//
	// example:
	//
	// Stalled
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeDiskEventSetDiskEventTypeEventType) String() string {
	return dara.Prettify(s)
}

func (s DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeDiskEventSetDiskEventTypeEventType) GoString() string {
	return s.String()
}

func (s *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeDiskEventSetDiskEventTypeEventType) GetCode() *int32 {
	return s.Code
}

func (s *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeDiskEventSetDiskEventTypeEventType) GetName() *string {
	return s.Name
}

func (s *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeDiskEventSetDiskEventTypeEventType) SetCode(v int32) *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeDiskEventSetDiskEventTypeEventType {
	s.Code = &v
	return s
}

func (s *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeDiskEventSetDiskEventTypeEventType) SetName(v string) *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeDiskEventSetDiskEventTypeEventType {
	s.Name = &v
	return s
}

func (s *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeDiskEventSetDiskEventTypeEventType) Validate() error {
	return dara.Validate(s)
}

type DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeHealthStatus struct {
	// The code of the health status of the EBS device.
	//
	// example:
	//
	// 128
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The name of the health status of the EBS device.
	//
	// example:
	//
	// Impaired
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeHealthStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeHealthStatus) GoString() string {
	return s.String()
}

func (s *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeHealthStatus) GetCode() *int32 {
	return s.Code
}

func (s *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeHealthStatus) GetName() *string {
	return s.Name
}

func (s *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeHealthStatus) SetCode(v int32) *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeHealthStatus {
	s.Code = &v
	return s
}

func (s *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeHealthStatus) SetName(v string) *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeHealthStatus {
	s.Name = &v
	return s
}

func (s *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeHealthStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeStatus struct {
	// The code of the lifecycle status of the EBS device.
	//
	// example:
	//
	// 129
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The name of the lifecycle status of the EBS device.
	//
	// example:
	//
	// Available
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeStatus) GoString() string {
	return s.String()
}

func (s *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeStatus) GetCode() *int32 {
	return s.Code
}

func (s *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeStatus) GetName() *string {
	return s.Name
}

func (s *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeStatus) SetCode(v int32) *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeStatus {
	s.Code = &v
	return s
}

func (s *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeStatus) SetName(v string) *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeStatus {
	s.Name = &v
	return s
}

func (s *DescribeDisksFullStatusResponseBodyDiskFullStatusSetDiskFullStatusTypeStatus) Validate() error {
	return dara.Validate(s)
}
