// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CreateClusterResponseBody
	GetClusterId() *string
	SetRequestId(v string) *CreateClusterResponseBody
	GetRequestId() *string
	SetTaskId(v string) *CreateClusterResponseBody
	GetTaskId() *string
}

type CreateClusterResponseBody struct {
	// Cluster ID
	//
	// example:
	//
	// i116913051663373010974
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Request ID
	//
	// example:
	//
	// 3C683243-7915-57FB-9570-A2932C1C0F78
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Task ID
	//
	// example:
	//
	// i159809891662373011015
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClusterResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateClusterResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateClusterResponseBody) SetClusterId(v string) *CreateClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *CreateClusterResponseBody) SetRequestId(v string) *CreateClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateClusterResponseBody) SetTaskId(v string) *CreateClusterResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
