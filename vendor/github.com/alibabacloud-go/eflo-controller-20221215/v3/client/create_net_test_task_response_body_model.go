// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetTestTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateNetTestTaskResponseBody
	GetRequestId() *string
	SetTestId(v string) *CreateNetTestTaskResponseBody
	GetTestId() *string
}

type CreateNetTestTaskResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the test task. The unique identifier of a network test task.
	//
	// example:
	//
	// dr-uf6i0tv2refv8wz*****
	TestId *string `json:"TestId,omitempty" xml:"TestId,omitempty"`
}

func (s CreateNetTestTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateNetTestTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNetTestTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateNetTestTaskResponseBody) GetTestId() *string {
	return s.TestId
}

func (s *CreateNetTestTaskResponseBody) SetRequestId(v string) *CreateNetTestTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNetTestTaskResponseBody) SetTestId(v string) *CreateNetTestTaskResponseBody {
	s.TestId = &v
	return s
}

func (s *CreateNetTestTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
