// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeNodeGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ChangeNodeGroupResponseBody
	GetRequestId() *string
	SetTaskId(v string) *ChangeNodeGroupResponseBody
	GetTaskId() *string
}

type ChangeNodeGroupResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 887FA855-89F4-5DB3-B305-C5879EC480E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// i159136551662516768776
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ChangeNodeGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeNodeGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeNodeGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeNodeGroupResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *ChangeNodeGroupResponseBody) SetRequestId(v string) *ChangeNodeGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeNodeGroupResponseBody) SetTaskId(v string) *ChangeNodeGroupResponseBody {
	s.TaskId = &v
	return s
}

func (s *ChangeNodeGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
