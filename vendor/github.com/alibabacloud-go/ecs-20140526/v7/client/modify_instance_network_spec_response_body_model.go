// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceNetworkSpecResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *ModifyInstanceNetworkSpecResponseBody
	GetOrderId() *string
	SetRequestId(v string) *ModifyInstanceNetworkSpecResponseBody
	GetRequestId() *string
}

type ModifyInstanceNetworkSpecResponseBody struct {
	// The order ID.
	//
	// example:
	//
	// 123457890
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceNetworkSpecResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceNetworkSpecResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceNetworkSpecResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *ModifyInstanceNetworkSpecResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceNetworkSpecResponseBody) SetOrderId(v string) *ModifyInstanceNetworkSpecResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyInstanceNetworkSpecResponseBody) SetRequestId(v string) *ModifyInstanceNetworkSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceNetworkSpecResponseBody) Validate() error {
	return dara.Validate(s)
}
