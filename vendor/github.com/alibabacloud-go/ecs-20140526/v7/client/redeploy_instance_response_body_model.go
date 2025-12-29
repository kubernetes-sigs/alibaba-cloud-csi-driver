// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRedeployInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RedeployInstanceResponseBody
	GetRequestId() *string
	SetTaskId(v string) *RedeployInstanceResponseBody
	GetTaskId() *string
}

type RedeployInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the redeployment task.
	//
	// You can call the [DescribeTasks](https://help.aliyun.com/document_detail/25622.html) operation to query the redeployment result.
	//
	// example:
	//
	// t-bp10e8orkp8x****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s RedeployInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RedeployInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RedeployInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RedeployInstanceResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *RedeployInstanceResponseBody) SetRequestId(v string) *RedeployInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RedeployInstanceResponseBody) SetTaskId(v string) *RedeployInstanceResponseBody {
	s.TaskId = &v
	return s
}

func (s *RedeployInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
