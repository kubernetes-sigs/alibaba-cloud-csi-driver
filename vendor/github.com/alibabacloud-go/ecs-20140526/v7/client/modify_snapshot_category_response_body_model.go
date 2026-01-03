// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySnapshotCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifySnapshotCategoryResponseBody
	GetRequestId() *string
	SetTaskId(v string) *ModifySnapshotCategoryResponseBody
	GetTaskId() *string
}

type ModifySnapshotCategoryResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B48A12CD-1295-4A38-A8F0-0E92C937****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the archive task. You can call the DescribeTasks operation to query the status and progress of the archive task.
	//
	// example:
	//
	// t-dxh34xds**d
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ModifySnapshotCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySnapshotCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySnapshotCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySnapshotCategoryResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *ModifySnapshotCategoryResponseBody) SetRequestId(v string) *ModifySnapshotCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySnapshotCategoryResponseBody) SetTaskId(v string) *ModifySnapshotCategoryResponseBody {
	s.TaskId = &v
	return s
}

func (s *ModifySnapshotCategoryResponseBody) Validate() error {
	return dara.Validate(s)
}
