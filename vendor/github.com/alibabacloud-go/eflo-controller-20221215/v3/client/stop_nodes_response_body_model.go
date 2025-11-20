// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopNodesResponseBody
	GetRequestId() *string
	SetTaskId(v string) *StopNodesResponseBody
	GetTaskId() *string
}

type StopNodesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4FD06DF0-9167-5C6F-A145-F30CA4A15D54
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// i155847351716171893489
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StopNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopNodesResponseBody) GoString() string {
	return s.String()
}

func (s *StopNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopNodesResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *StopNodesResponseBody) SetRequestId(v string) *StopNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopNodesResponseBody) SetTaskId(v string) *StopNodesResponseBody {
	s.TaskId = &v
	return s
}

func (s *StopNodesResponseBody) Validate() error {
	return dara.Validate(s)
}
