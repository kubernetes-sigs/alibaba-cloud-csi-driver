// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPrepayInstanceSpecResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *ModifyPrepayInstanceSpecResponseBody
	GetOrderId() *string
	SetRequestId(v string) *ModifyPrepayInstanceSpecResponseBody
	GetRequestId() *string
}

type ModifyPrepayInstanceSpecResponseBody struct {
	// The ID of the order.
	//
	// example:
	//
	// 1234567890
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyPrepayInstanceSpecResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyPrepayInstanceSpecResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPrepayInstanceSpecResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *ModifyPrepayInstanceSpecResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyPrepayInstanceSpecResponseBody) SetOrderId(v string) *ModifyPrepayInstanceSpecResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyPrepayInstanceSpecResponseBody) SetRequestId(v string) *ModifyPrepayInstanceSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyPrepayInstanceSpecResponseBody) Validate() error {
	return dara.Validate(s)
}
