// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNodeGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateNodeGroupResponseBody
	GetRequestId() *string
	SetTaskId(v string) *UpdateNodeGroupResponseBody
	GetTaskId() *string
}

type UpdateNodeGroupResponseBody struct {
	// Request ID
	//
	// example:
	//
	// 8F065DDD-6996-5973-9691-9EC57BD0072E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Task ID
	//
	// example:
	//
	// i154273451745372344629
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpdateNodeGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateNodeGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNodeGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateNodeGroupResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateNodeGroupResponseBody) SetRequestId(v string) *UpdateNodeGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateNodeGroupResponseBody) SetTaskId(v string) *UpdateNodeGroupResponseBody {
	s.TaskId = &v
	return s
}

func (s *UpdateNodeGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
