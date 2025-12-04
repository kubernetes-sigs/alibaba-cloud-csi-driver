// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoSnapshotTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAutoSnapshotTasks(v *DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasks) *DescribeAutoSnapshotTasksResponseBody
	GetAutoSnapshotTasks() *DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasks
	SetPageNumber(v int32) *DescribeAutoSnapshotTasksResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAutoSnapshotTasksResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeAutoSnapshotTasksResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeAutoSnapshotTasksResponseBody
	GetTotalCount() *int32
}

type DescribeAutoSnapshotTasksResponseBody struct {
	// The queried automatic snapshot tasks.
	AutoSnapshotTasks *DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasks `json:"AutoSnapshotTasks,omitempty" xml:"AutoSnapshotTasks,omitempty" type:"Struct"`
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
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of automatic snapshot tasks.
	//
	// example:
	//
	// 9
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAutoSnapshotTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoSnapshotTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAutoSnapshotTasksResponseBody) GetAutoSnapshotTasks() *DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasks {
	return s.AutoSnapshotTasks
}

func (s *DescribeAutoSnapshotTasksResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAutoSnapshotTasksResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAutoSnapshotTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAutoSnapshotTasksResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeAutoSnapshotTasksResponseBody) SetAutoSnapshotTasks(v *DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasks) *DescribeAutoSnapshotTasksResponseBody {
	s.AutoSnapshotTasks = v
	return s
}

func (s *DescribeAutoSnapshotTasksResponseBody) SetPageNumber(v int32) *DescribeAutoSnapshotTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAutoSnapshotTasksResponseBody) SetPageSize(v int32) *DescribeAutoSnapshotTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAutoSnapshotTasksResponseBody) SetRequestId(v string) *DescribeAutoSnapshotTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAutoSnapshotTasksResponseBody) SetTotalCount(v int32) *DescribeAutoSnapshotTasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAutoSnapshotTasksResponseBody) Validate() error {
	if s.AutoSnapshotTasks != nil {
		if err := s.AutoSnapshotTasks.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasks struct {
	AutoSnapshotTask []*DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasksAutoSnapshotTask `json:"AutoSnapshotTask,omitempty" xml:"AutoSnapshotTask,omitempty" type:"Repeated"`
}

func (s DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasks) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasks) GoString() string {
	return s.String()
}

func (s *DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasks) GetAutoSnapshotTask() []*DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasksAutoSnapshotTask {
	return s.AutoSnapshotTask
}

func (s *DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasks) SetAutoSnapshotTask(v []*DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasksAutoSnapshotTask) *DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasks {
	s.AutoSnapshotTask = v
	return s
}

func (s *DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasks) Validate() error {
	if s.AutoSnapshotTask != nil {
		for _, item := range s.AutoSnapshotTask {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasksAutoSnapshotTask struct {
	// The ID of the automatic snapshot policy.
	//
	// example:
	//
	// sp-extreme-233e6****
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	// The ID of the file system.
	//
	// example:
	//
	// extreme-233e6****
	SourceFileSystemId *string `json:"SourceFileSystemId,omitempty" xml:"SourceFileSystemId,omitempty"`
}

func (s DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasksAutoSnapshotTask) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasksAutoSnapshotTask) GoString() string {
	return s.String()
}

func (s *DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasksAutoSnapshotTask) GetAutoSnapshotPolicyId() *string {
	return s.AutoSnapshotPolicyId
}

func (s *DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasksAutoSnapshotTask) GetSourceFileSystemId() *string {
	return s.SourceFileSystemId
}

func (s *DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasksAutoSnapshotTask) SetAutoSnapshotPolicyId(v string) *DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasksAutoSnapshotTask {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasksAutoSnapshotTask) SetSourceFileSystemId(v string) *DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasksAutoSnapshotTask {
	s.SourceFileSystemId = &v
	return s
}

func (s *DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasksAutoSnapshotTask) Validate() error {
	return dara.Validate(s)
}
