// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoProvisioningGroupHistoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAutoProvisioningGroupHistories(v *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistories) *DescribeAutoProvisioningGroupHistoryResponseBody
	GetAutoProvisioningGroupHistories() *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistories
	SetPageNumber(v int32) *DescribeAutoProvisioningGroupHistoryResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAutoProvisioningGroupHistoryResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeAutoProvisioningGroupHistoryResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeAutoProvisioningGroupHistoryResponseBody
	GetTotalCount() *int32
}

type DescribeAutoProvisioningGroupHistoryResponseBody struct {
	// An array consisting of AutoProvisioningGroupHistory data.
	AutoProvisioningGroupHistories *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistories `json:"AutoProvisioningGroupHistories,omitempty" xml:"AutoProvisioningGroupHistories,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B48A12CD-1295-4A38-A8F0-0E92C937****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of queried scheduling tasks in the auto provisioning group.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAutoProvisioningGroupHistoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoProvisioningGroupHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBody) GetAutoProvisioningGroupHistories() *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistories {
	return s.AutoProvisioningGroupHistories
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBody) SetAutoProvisioningGroupHistories(v *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistories) *DescribeAutoProvisioningGroupHistoryResponseBody {
	s.AutoProvisioningGroupHistories = v
	return s
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBody) SetPageNumber(v int32) *DescribeAutoProvisioningGroupHistoryResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBody) SetPageSize(v int32) *DescribeAutoProvisioningGroupHistoryResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBody) SetRequestId(v string) *DescribeAutoProvisioningGroupHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBody) SetTotalCount(v int32) *DescribeAutoProvisioningGroupHistoryResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBody) Validate() error {
	if s.AutoProvisioningGroupHistories != nil {
		if err := s.AutoProvisioningGroupHistories.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistories struct {
	AutoProvisioningGroupHistory []*DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistory `json:"AutoProvisioningGroupHistory,omitempty" xml:"AutoProvisioningGroupHistory,omitempty" type:"Repeated"`
}

func (s DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistories) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistories) GoString() string {
	return s.String()
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistories) GetAutoProvisioningGroupHistory() []*DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistory {
	return s.AutoProvisioningGroupHistory
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistories) SetAutoProvisioningGroupHistory(v []*DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistory) *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistories {
	s.AutoProvisioningGroupHistory = v
	return s
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistories) Validate() error {
	if s.AutoProvisioningGroupHistory != nil {
		for _, item := range s.AutoProvisioningGroupHistory {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistory struct {
	// An array consisting of ActivityDetail data.
	ActivityDetails *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetails `json:"ActivityDetails,omitempty" xml:"ActivityDetails,omitempty" type:"Struct"`
	// The execution time of the last instance creation performed by the single scheduling task.
	//
	// example:
	//
	// 2019-04-01T15:10:20Z
	LastEventTime *string `json:"LastEventTime,omitempty" xml:"LastEventTime,omitempty"`
	// The start time of executing the single scheduling task.
	//
	// example:
	//
	// 2019-04-01T15:10:20Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The execution status of the single scheduling task. Valid values:
	//
	// 	- prepare: The scheduling task is being executed.
	//
	// 	- success: The scheduling task is executed.
	//
	// 	- failed: The scheduling task failed to be executed.
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the scheduling task.
	//
	// example:
	//
	// apg-task-bp67acfmxazb4p****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistory) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistory) GoString() string {
	return s.String()
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistory) GetActivityDetails() *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetails {
	return s.ActivityDetails
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistory) GetLastEventTime() *string {
	return s.LastEventTime
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistory) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistory) GetStatus() *string {
	return s.Status
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistory) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistory) SetActivityDetails(v *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetails) *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistory {
	s.ActivityDetails = v
	return s
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistory) SetLastEventTime(v string) *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistory {
	s.LastEventTime = &v
	return s
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistory) SetStartTime(v string) *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistory {
	s.StartTime = &v
	return s
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistory) SetStatus(v string) *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistory {
	s.Status = &v
	return s
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistory) SetTaskId(v string) *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistory {
	s.TaskId = &v
	return s
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistory) Validate() error {
	if s.ActivityDetails != nil {
		if err := s.ActivityDetails.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetails struct {
	ActivityDetail []*DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetail `json:"ActivityDetail,omitempty" xml:"ActivityDetail,omitempty" type:"Repeated"`
}

func (s DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetails) GoString() string {
	return s.String()
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetails) GetActivityDetail() []*DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetail {
	return s.ActivityDetail
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetails) SetActivityDetail(v []*DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetail) *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetails {
	s.ActivityDetail = v
	return s
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetails) Validate() error {
	if s.ActivityDetail != nil {
		for _, item := range s.ActivityDetail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetail struct {
	// The execution details of instance creation performed by the single scheduling task.
	//
	// example:
	//
	// New ECS instances "i-bp67acfmxazb4p****, i-bp67acfmxazb5p****" created.
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The execution status of instance creation performed by the single scheduling task. Valid values:
	//
	// 	- Successful: Instances are created.
	//
	// 	- Failed: Instances failed to be created.
	//
	// 	- InProgress: Instances are being created.
	//
	// 	- Warning: Some instances are created.
	//
	// example:
	//
	// Successful
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetail) GoString() string {
	return s.String()
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetail) GetDetail() *string {
	return s.Detail
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetail) GetStatus() *string {
	return s.Status
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetail) SetDetail(v string) *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetail {
	s.Detail = &v
	return s
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetail) SetStatus(v string) *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetail {
	s.Status = &v
	return s
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetail) Validate() error {
	return dara.Validate(s)
}
