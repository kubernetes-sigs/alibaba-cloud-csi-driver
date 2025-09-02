// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *DescribeTaskRequest
	GetTaskId() *string
}

type DescribeTaskRequest struct {
	// The task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// i156331731670384438138
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeTaskRequest) SetTaskId(v string) *DescribeTaskRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeTaskRequest) Validate() error {
	return dara.Validate(s)
}
