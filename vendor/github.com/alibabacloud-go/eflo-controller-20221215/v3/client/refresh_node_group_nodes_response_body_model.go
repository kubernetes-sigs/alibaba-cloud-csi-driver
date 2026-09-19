// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshNodeGroupNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNodeGroupRefreshTaskId(v string) *RefreshNodeGroupNodesResponseBody
	GetNodeGroupRefreshTaskId() *string
	SetRequestId(v string) *RefreshNodeGroupNodesResponseBody
	GetRequestId() *string
}

type RefreshNodeGroupNodesResponseBody struct {
	// The task ID. Query the task progress by calling DescribeNodeGroupRefreshTask. If no nodes have configuration drift, no task is created and this field returns an empty string.
	//
	// example:
	//
	// task-159136551662516768776
	NodeGroupRefreshTaskId *string `json:"NodeGroupRefreshTaskId,omitempty" xml:"NodeGroupRefreshTaskId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 4FD06DF0-9167-5C6F-A145-F30CA4A15D54
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RefreshNodeGroupNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RefreshNodeGroupNodesResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshNodeGroupNodesResponseBody) GetNodeGroupRefreshTaskId() *string {
	return s.NodeGroupRefreshTaskId
}

func (s *RefreshNodeGroupNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RefreshNodeGroupNodesResponseBody) SetNodeGroupRefreshTaskId(v string) *RefreshNodeGroupNodesResponseBody {
	s.NodeGroupRefreshTaskId = &v
	return s
}

func (s *RefreshNodeGroupNodesResponseBody) SetRequestId(v string) *RefreshNodeGroupNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefreshNodeGroupNodesResponseBody) Validate() error {
	return dara.Validate(s)
}
