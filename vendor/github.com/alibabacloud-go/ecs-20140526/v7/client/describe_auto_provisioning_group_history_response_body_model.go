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
	ActivityDetails *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetails `json:"ActivityDetails,omitempty" xml:"ActivityDetails,omitempty" type:"Struct"`
	LastEventTime   *string                                                                                                                    `json:"LastEventTime,omitempty" xml:"LastEventTime,omitempty"`
	StartTime       *string                                                                                                                    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status          *string                                                                                                                    `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId          *string                                                                                                                    `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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
	CreatedInstanceIds   *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetailCreatedInstanceIds   `json:"CreatedInstanceIds,omitempty" xml:"CreatedInstanceIds,omitempty" type:"Struct"`
	DestroyedInstanceIds *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetailDestroyedInstanceIds `json:"DestroyedInstanceIds,omitempty" xml:"DestroyedInstanceIds,omitempty" type:"Struct"`
	Detail               *string                                                                                                                                                      `json:"Detail,omitempty" xml:"Detail,omitempty"`
	ErrorMessages        *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetailErrorMessages        `json:"ErrorMessages,omitempty" xml:"ErrorMessages,omitempty" type:"Struct"`
	Status               *string                                                                                                                                                      `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetail) GoString() string {
	return s.String()
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetail) GetCreatedInstanceIds() *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetailCreatedInstanceIds {
	return s.CreatedInstanceIds
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetail) GetDestroyedInstanceIds() *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetailDestroyedInstanceIds {
	return s.DestroyedInstanceIds
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetail) GetDetail() *string {
	return s.Detail
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetail) GetErrorMessages() *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetailErrorMessages {
	return s.ErrorMessages
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetail) GetStatus() *string {
	return s.Status
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetail) SetCreatedInstanceIds(v *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetailCreatedInstanceIds) *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetail {
	s.CreatedInstanceIds = v
	return s
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetail) SetDestroyedInstanceIds(v *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetailDestroyedInstanceIds) *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetail {
	s.DestroyedInstanceIds = v
	return s
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetail) SetDetail(v string) *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetail {
	s.Detail = &v
	return s
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetail) SetErrorMessages(v *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetailErrorMessages) *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetail {
	s.ErrorMessages = v
	return s
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetail) SetStatus(v string) *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetail {
	s.Status = &v
	return s
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetail) Validate() error {
	if s.CreatedInstanceIds != nil {
		if err := s.CreatedInstanceIds.Validate(); err != nil {
			return err
		}
	}
	if s.DestroyedInstanceIds != nil {
		if err := s.DestroyedInstanceIds.Validate(); err != nil {
			return err
		}
	}
	if s.ErrorMessages != nil {
		if err := s.ErrorMessages.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetailCreatedInstanceIds struct {
	CreatedInstanceId []*string `json:"CreatedInstanceId,omitempty" xml:"CreatedInstanceId,omitempty" type:"Repeated"`
}

func (s DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetailCreatedInstanceIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetailCreatedInstanceIds) GoString() string {
	return s.String()
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetailCreatedInstanceIds) GetCreatedInstanceId() []*string {
	return s.CreatedInstanceId
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetailCreatedInstanceIds) SetCreatedInstanceId(v []*string) *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetailCreatedInstanceIds {
	s.CreatedInstanceId = v
	return s
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetailCreatedInstanceIds) Validate() error {
	return dara.Validate(s)
}

type DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetailDestroyedInstanceIds struct {
	DestroyedInstanceId []*string `json:"DestroyedInstanceId,omitempty" xml:"DestroyedInstanceId,omitempty" type:"Repeated"`
}

func (s DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetailDestroyedInstanceIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetailDestroyedInstanceIds) GoString() string {
	return s.String()
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetailDestroyedInstanceIds) GetDestroyedInstanceId() []*string {
	return s.DestroyedInstanceId
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetailDestroyedInstanceIds) SetDestroyedInstanceId(v []*string) *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetailDestroyedInstanceIds {
	s.DestroyedInstanceId = v
	return s
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetailDestroyedInstanceIds) Validate() error {
	return dara.Validate(s)
}

type DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetailErrorMessages struct {
	ErrorMessage []*DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetailErrorMessagesErrorMessage `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty" type:"Repeated"`
}

func (s DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetailErrorMessages) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetailErrorMessages) GoString() string {
	return s.String()
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetailErrorMessages) GetErrorMessage() []*DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetailErrorMessagesErrorMessage {
	return s.ErrorMessage
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetailErrorMessages) SetErrorMessage(v []*DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetailErrorMessagesErrorMessage) *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetailErrorMessages {
	s.ErrorMessage = v
	return s
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetailErrorMessages) Validate() error {
	if s.ErrorMessage != nil {
		for _, item := range s.ErrorMessage {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetailErrorMessagesErrorMessage struct {
	// example:
	//
	// InvalidSecurityGroupId.NotFound
	Code              *string                                                                                                                                                                            `json:"Code,omitempty" xml:"Code,omitempty"`
	FailedInstanceIds *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetailErrorMessagesErrorMessageFailedInstanceIds `json:"FailedInstanceIds,omitempty" xml:"FailedInstanceIds,omitempty" type:"Struct"`
	// example:
	//
	// The specified SecurityGroupId "sg-bp1d8modxxxxx" is not found.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetailErrorMessagesErrorMessage) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetailErrorMessagesErrorMessage) GoString() string {
	return s.String()
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetailErrorMessagesErrorMessage) GetCode() *string {
	return s.Code
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetailErrorMessagesErrorMessage) GetFailedInstanceIds() *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetailErrorMessagesErrorMessageFailedInstanceIds {
	return s.FailedInstanceIds
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetailErrorMessagesErrorMessage) GetMessage() *string {
	return s.Message
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetailErrorMessagesErrorMessage) SetCode(v string) *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetailErrorMessagesErrorMessage {
	s.Code = &v
	return s
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetailErrorMessagesErrorMessage) SetFailedInstanceIds(v *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetailErrorMessagesErrorMessageFailedInstanceIds) *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetailErrorMessagesErrorMessage {
	s.FailedInstanceIds = v
	return s
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetailErrorMessagesErrorMessage) SetMessage(v string) *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetailErrorMessagesErrorMessage {
	s.Message = &v
	return s
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetailErrorMessagesErrorMessage) Validate() error {
	if s.FailedInstanceIds != nil {
		if err := s.FailedInstanceIds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetailErrorMessagesErrorMessageFailedInstanceIds struct {
	FailedInstanceId []*string `json:"FailedInstanceId,omitempty" xml:"FailedInstanceId,omitempty" type:"Repeated"`
}

func (s DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetailErrorMessagesErrorMessageFailedInstanceIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetailErrorMessagesErrorMessageFailedInstanceIds) GoString() string {
	return s.String()
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetailErrorMessagesErrorMessageFailedInstanceIds) GetFailedInstanceId() []*string {
	return s.FailedInstanceId
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetailErrorMessagesErrorMessageFailedInstanceIds) SetFailedInstanceId(v []*string) *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetailErrorMessagesErrorMessageFailedInstanceIds {
	s.FailedInstanceId = v
	return s
}

func (s *DescribeAutoProvisioningGroupHistoryResponseBodyAutoProvisioningGroupHistoriesAutoProvisioningGroupHistoryActivityDetailsActivityDetailErrorMessagesErrorMessageFailedInstanceIds) Validate() error {
	return dara.Validate(s)
}
