// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDiskSpecResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *ModifyDiskSpecResponseBody
	GetOrderId() *string
	SetRequestId(v string) *ModifyDiskSpecResponseBody
	GetRequestId() *string
	SetTaskId(v string) *ModifyDiskSpecResponseBody
	GetTaskId() *string
}

type ModifyDiskSpecResponseBody struct {
	// The order ID.
	//
	// >  This parameter is returned only when the category of a subscription disk or the performance level of a subscription ESSD is modified.
	//
	// example:
	//
	// 20413515388****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the disk category change task.
	//
	// >  If you only modify the performance level of an ESSD, this parameter is not returned.
	//
	// example:
	//
	// t-bp67acfmxazb4p****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ModifyDiskSpecResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDiskSpecResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDiskSpecResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *ModifyDiskSpecResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDiskSpecResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *ModifyDiskSpecResponseBody) SetOrderId(v string) *ModifyDiskSpecResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyDiskSpecResponseBody) SetRequestId(v string) *ModifyDiskSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDiskSpecResponseBody) SetTaskId(v string) *ModifyDiskSpecResponseBody {
	s.TaskId = &v
	return s
}

func (s *ModifyDiskSpecResponseBody) Validate() error {
	return dara.Validate(s)
}
