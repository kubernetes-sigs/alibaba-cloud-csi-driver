// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceStatuses(v *DescribeInstanceStatusResponseBodyInstanceStatuses) *DescribeInstanceStatusResponseBody
	GetInstanceStatuses() *DescribeInstanceStatusResponseBodyInstanceStatuses
	SetPageNumber(v int32) *DescribeInstanceStatusResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeInstanceStatusResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeInstanceStatusResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeInstanceStatusResponseBody
	GetTotalCount() *int32
}

type DescribeInstanceStatusResponseBody struct {
	// The IDs and status of the ECS instances.
	InstanceStatuses *DescribeInstanceStatusResponseBodyInstanceStatuses `json:"InstanceStatuses,omitempty" xml:"InstanceStatuses,omitempty" type:"Struct"`
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
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of instances.
	//
	// example:
	//
	// 58
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeInstanceStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceStatusResponseBody) GetInstanceStatuses() *DescribeInstanceStatusResponseBodyInstanceStatuses {
	return s.InstanceStatuses
}

func (s *DescribeInstanceStatusResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeInstanceStatusResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInstanceStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceStatusResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeInstanceStatusResponseBody) SetInstanceStatuses(v *DescribeInstanceStatusResponseBodyInstanceStatuses) *DescribeInstanceStatusResponseBody {
	s.InstanceStatuses = v
	return s
}

func (s *DescribeInstanceStatusResponseBody) SetPageNumber(v int32) *DescribeInstanceStatusResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstanceStatusResponseBody) SetPageSize(v int32) *DescribeInstanceStatusResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeInstanceStatusResponseBody) SetRequestId(v string) *DescribeInstanceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceStatusResponseBody) SetTotalCount(v int32) *DescribeInstanceStatusResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeInstanceStatusResponseBody) Validate() error {
	if s.InstanceStatuses != nil {
		if err := s.InstanceStatuses.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstanceStatusResponseBodyInstanceStatuses struct {
	InstanceStatus []*DescribeInstanceStatusResponseBodyInstanceStatusesInstanceStatus `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty" type:"Repeated"`
}

func (s DescribeInstanceStatusResponseBodyInstanceStatuses) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceStatusResponseBodyInstanceStatuses) GoString() string {
	return s.String()
}

func (s *DescribeInstanceStatusResponseBodyInstanceStatuses) GetInstanceStatus() []*DescribeInstanceStatusResponseBodyInstanceStatusesInstanceStatus {
	return s.InstanceStatus
}

func (s *DescribeInstanceStatusResponseBodyInstanceStatuses) SetInstanceStatus(v []*DescribeInstanceStatusResponseBodyInstanceStatusesInstanceStatus) *DescribeInstanceStatusResponseBodyInstanceStatuses {
	s.InstanceStatus = v
	return s
}

func (s *DescribeInstanceStatusResponseBodyInstanceStatuses) Validate() error {
	if s.InstanceStatus != nil {
		for _, item := range s.InstanceStatus {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstanceStatusResponseBodyInstanceStatusesInstanceStatus struct {
	// The ID of the instance.
	//
	// example:
	//
	// i-bp1j4i2jdf3owlhe****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The status of the instance. Valid values:
	//
	// 	- Pending: The instance is being created.
	//
	// 	- Running: The instance is running.
	//
	// 	- Starting: The instance is being started.
	//
	// 	- Stopping: The instance is being stopped.
	//
	// 	- Stopped: The instance is stopped.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeInstanceStatusResponseBodyInstanceStatusesInstanceStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceStatusResponseBodyInstanceStatusesInstanceStatus) GoString() string {
	return s.String()
}

func (s *DescribeInstanceStatusResponseBodyInstanceStatusesInstanceStatus) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceStatusResponseBodyInstanceStatusesInstanceStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeInstanceStatusResponseBodyInstanceStatusesInstanceStatus) SetInstanceId(v string) *DescribeInstanceStatusResponseBodyInstanceStatusesInstanceStatus {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceStatusResponseBodyInstanceStatusesInstanceStatus) SetStatus(v string) *DescribeInstanceStatusResponseBodyInstanceStatusesInstanceStatus {
	s.Status = &v
	return s
}

func (s *DescribeInstanceStatusResponseBodyInstanceStatusesInstanceStatus) Validate() error {
	return dara.Validate(s)
}
