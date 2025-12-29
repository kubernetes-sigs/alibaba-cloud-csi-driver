// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDiskChargeTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *ModifyDiskChargeTypeResponseBody
	GetOrderId() *string
	SetRequestId(v string) *ModifyDiskChargeTypeResponseBody
	GetRequestId() *string
}

type ModifyDiskChargeTypeResponseBody struct {
	// The ID of the order.
	//
	// example:
	//
	// 1234567890
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDiskChargeTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDiskChargeTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDiskChargeTypeResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *ModifyDiskChargeTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDiskChargeTypeResponseBody) SetOrderId(v string) *ModifyDiskChargeTypeResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyDiskChargeTypeResponseBody) SetRequestId(v string) *ModifyDiskChargeTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDiskChargeTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
