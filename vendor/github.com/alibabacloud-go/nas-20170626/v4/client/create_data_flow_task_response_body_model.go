// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataFlowTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateDataFlowTaskResponseBody
	GetRequestId() *string
	SetTaskId(v string) *CreateDataFlowTaskResponseBody
	GetTaskId() *string
}

type CreateDataFlowTaskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2D69A58F-345C-4FDE-88E4-BF518948****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the dataflow task.
	//
	// example:
	//
	// task-38aa8e890f45****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateDataFlowTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDataFlowTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataFlowTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDataFlowTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateDataFlowTaskResponseBody) SetRequestId(v string) *CreateDataFlowTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataFlowTaskResponseBody) SetTaskId(v string) *CreateDataFlowTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateDataFlowTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
