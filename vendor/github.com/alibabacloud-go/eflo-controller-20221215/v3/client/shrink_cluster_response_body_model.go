// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iShrinkClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ShrinkClusterResponseBody
	GetRequestId() *string
	SetTaskId(v string) *ShrinkClusterResponseBody
	GetTaskId() *string
}

type ShrinkClusterResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CC9FEF89-9BE5-5E03-845E-238B48D7599B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The job ID.
	//
	// example:
	//
	// i159136551662516768776
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ShrinkClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ShrinkClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ShrinkClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ShrinkClusterResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *ShrinkClusterResponseBody) SetRequestId(v string) *ShrinkClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ShrinkClusterResponseBody) SetTaskId(v string) *ShrinkClusterResponseBody {
	s.TaskId = &v
	return s
}

func (s *ShrinkClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
