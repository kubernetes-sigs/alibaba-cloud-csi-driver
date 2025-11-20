// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RebootNodesResponseBody
	GetRequestId() *string
	SetTaskId(v string) *RebootNodesResponseBody
	GetTaskId() *string
}

type RebootNodesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4FD06DF0-9167-5C6F-A145-F30CA4A15D54
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The job ID.
	//
	// example:
	//
	// i158475611663639202234
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s RebootNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RebootNodesResponseBody) GoString() string {
	return s.String()
}

func (s *RebootNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RebootNodesResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *RebootNodesResponseBody) SetRequestId(v string) *RebootNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RebootNodesResponseBody) SetTaskId(v string) *RebootNodesResponseBody {
	s.TaskId = &v
	return s
}

func (s *RebootNodesResponseBody) Validate() error {
	return dara.Validate(s)
}
