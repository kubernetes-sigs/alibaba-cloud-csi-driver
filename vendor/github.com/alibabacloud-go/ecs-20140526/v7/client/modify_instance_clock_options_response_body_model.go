// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceClockOptionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyInstanceClockOptionsResponseBody
	GetRequestId() *string
	SetTaskId(v string) *ModifyInstanceClockOptionsResponseBody
	GetTaskId() *string
}

type ModifyInstanceClockOptionsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the modification task.
	//
	// You can call the [DescribeTasks](https://help.aliyun.com/document_detail/25622.html) operation to query the modification results.
	//
	// example:
	//
	// t-bp1hvgwromzv32iq****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ModifyInstanceClockOptionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceClockOptionsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceClockOptionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceClockOptionsResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *ModifyInstanceClockOptionsResponseBody) SetRequestId(v string) *ModifyInstanceClockOptionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceClockOptionsResponseBody) SetTaskId(v string) *ModifyInstanceClockOptionsResponseBody {
	s.TaskId = &v
	return s
}

func (s *ModifyInstanceClockOptionsResponseBody) Validate() error {
	return dara.Validate(s)
}
