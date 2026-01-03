// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceNetworkOptionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyInstanceNetworkOptionsResponseBody
	GetRequestId() *string
	SetTaskId(v string) *ModifyInstanceNetworkOptionsResponseBody
	GetTaskId() *string
}

type ModifyInstanceNetworkOptionsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the task for which the bandwidth weight is modified.
	//
	// You can use the [DescribeTaskAttribute](https://help.aliyun.com/document_detail/2679968.html) interface to modify the bandwidth weight result.
	//
	// example:
	//
	// t-bp198jigq7l0h5ac****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ModifyInstanceNetworkOptionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceNetworkOptionsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceNetworkOptionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceNetworkOptionsResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *ModifyInstanceNetworkOptionsResponseBody) SetRequestId(v string) *ModifyInstanceNetworkOptionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceNetworkOptionsResponseBody) SetTaskId(v string) *ModifyInstanceNetworkOptionsResponseBody {
	s.TaskId = &v
	return s
}

func (s *ModifyInstanceNetworkOptionsResponseBody) Validate() error {
	return dara.Validate(s)
}
