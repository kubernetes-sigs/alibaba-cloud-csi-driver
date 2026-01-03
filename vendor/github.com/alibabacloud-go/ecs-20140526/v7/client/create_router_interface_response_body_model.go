// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRouterInterfaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *CreateRouterInterfaceResponseBody
	GetOrderId() *int64
	SetRequestId(v string) *CreateRouterInterfaceResponseBody
	GetRequestId() *string
	SetRouterInterfaceId(v string) *CreateRouterInterfaceResponseBody
	GetRouterInterfaceId() *string
}

type CreateRouterInterfaceResponseBody struct {
	OrderId           *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RouterInterfaceId *string `json:"RouterInterfaceId,omitempty" xml:"RouterInterfaceId,omitempty"`
}

func (s CreateRouterInterfaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRouterInterfaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRouterInterfaceResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *CreateRouterInterfaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRouterInterfaceResponseBody) GetRouterInterfaceId() *string {
	return s.RouterInterfaceId
}

func (s *CreateRouterInterfaceResponseBody) SetOrderId(v int64) *CreateRouterInterfaceResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateRouterInterfaceResponseBody) SetRequestId(v string) *CreateRouterInterfaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRouterInterfaceResponseBody) SetRouterInterfaceId(v string) *CreateRouterInterfaceResponseBody {
	s.RouterInterfaceId = &v
	return s
}

func (s *CreateRouterInterfaceResponseBody) Validate() error {
	return dara.Validate(s)
}
