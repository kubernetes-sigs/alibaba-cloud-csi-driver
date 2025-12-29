// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeTasksRequest
	GetEndTime() *string
	SetOwnerAccount(v string) *DescribeTasksRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeTasksRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeTasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeTasksRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeTasksRequest
	GetRegionId() *string
	SetResourceIds(v []*string) *DescribeTasksRequest
	GetResourceIds() []*string
	SetResourceOwnerAccount(v string) *DescribeTasksRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeTasksRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *DescribeTasksRequest
	GetStartTime() *string
	SetTaskAction(v string) *DescribeTasksRequest
	GetTaskAction() *string
	SetTaskIds(v string) *DescribeTasksRequest
	GetTaskIds() *string
	SetTaskStatus(v string) *DescribeTasksRequest
	GetTaskStatus() *string
}

type DescribeTasksRequest struct {
	// The end of the time range to query. The time range refers to the period of time during which the task is created. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2020-11-23T15:16:00Z
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number.
	//
	// Pages start from page 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of the resources associated with the task. Valid values of N: 1 to 100.
	ResourceIds          []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The beginning of the time range to query. The time range refers to the period of time during which the task is created. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2020-11-23T15:10:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the operation that generates the task. Valid values:
	//
	// 	- ImportImage
	//
	// 	- ExportImage
	//
	// 	- RedeployInstance
	//
	// 	- ModifyDiskSpec
	//
	// 	- ArchiveSnapshot
	//
	// example:
	//
	// ImportImage
	TaskAction *string `json:"TaskAction,omitempty" xml:"TaskAction,omitempty"`
	// The task IDs. You can specify up to 100 task IDs at a time. Separate the task IDs with commas (,).
	//
	// example:
	//
	// t-bp1hvgwromzv32iq****,t-bp179lofu2pv768w****
	TaskIds *string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty"`
	// The task status. Valid values:
	//
	// 	- Finished
	//
	// 	- Processing
	//
	// 	- Failed
	//
	// This parameter is left empty by default.
	//
	// >  The system only queries tasks in the Finished, Processing, and Failed states and ignores other values.
	//
	// example:
	//
	// Finished
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s DescribeTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeTasksRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeTasksRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeTasksRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeTasksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeTasksRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeTasksRequest) GetResourceIds() []*string {
	return s.ResourceIds
}

func (s *DescribeTasksRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeTasksRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeTasksRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeTasksRequest) GetTaskAction() *string {
	return s.TaskAction
}

func (s *DescribeTasksRequest) GetTaskIds() *string {
	return s.TaskIds
}

func (s *DescribeTasksRequest) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *DescribeTasksRequest) SetEndTime(v string) *DescribeTasksRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeTasksRequest) SetOwnerAccount(v string) *DescribeTasksRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeTasksRequest) SetOwnerId(v int64) *DescribeTasksRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeTasksRequest) SetPageNumber(v int32) *DescribeTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeTasksRequest) SetPageSize(v int32) *DescribeTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTasksRequest) SetRegionId(v string) *DescribeTasksRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTasksRequest) SetResourceIds(v []*string) *DescribeTasksRequest {
	s.ResourceIds = v
	return s
}

func (s *DescribeTasksRequest) SetResourceOwnerAccount(v string) *DescribeTasksRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeTasksRequest) SetResourceOwnerId(v int64) *DescribeTasksRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeTasksRequest) SetStartTime(v string) *DescribeTasksRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeTasksRequest) SetTaskAction(v string) *DescribeTasksRequest {
	s.TaskAction = &v
	return s
}

func (s *DescribeTasksRequest) SetTaskIds(v string) *DescribeTasksRequest {
	s.TaskIds = &v
	return s
}

func (s *DescribeTasksRequest) SetTaskStatus(v string) *DescribeTasksRequest {
	s.TaskStatus = &v
	return s
}

func (s *DescribeTasksRequest) Validate() error {
	return dara.Validate(s)
}
