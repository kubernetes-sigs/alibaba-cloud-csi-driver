// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInvocationAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCommandId(v string) *ModifyInvocationAttributeResponseBody
	GetCommandId() *string
	SetRequestId(v string) *ModifyInvocationAttributeResponseBody
	GetRequestId() *string
}

type ModifyInvocationAttributeResponseBody struct {
	// The command ID.
	//
	// 	- A new command is added and the `CommandId` value of the new command is returned only when `CommandContent` is changed.
	//
	// 	- No new command is added and the `CommandId` value of the command that is running is returned if `CommandContent` is not changed.
	//
	// 	- If you set `KeepCommand` to `true` when you called the [InvokeCommand](https://help.aliyun.com/document_detail/64841.html) or [RunCommand](https://help.aliyun.com/document_detail/141751.html) operation, a new command is added and retained. Otherwise, commands related to the task are deleted after all executions of the task are complete or the task is manually stopped.
	//
	// example:
	//
	// c-hz01272yr52****
	CommandId *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInvocationAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInvocationAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInvocationAttributeResponseBody) GetCommandId() *string {
	return s.CommandId
}

func (s *ModifyInvocationAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInvocationAttributeResponseBody) SetCommandId(v string) *ModifyInvocationAttributeResponseBody {
	s.CommandId = &v
	return s
}

func (s *ModifyInvocationAttributeResponseBody) SetRequestId(v string) *ModifyInvocationAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInvocationAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
