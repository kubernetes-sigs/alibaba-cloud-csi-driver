// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewDedicatedHostsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *RenewDedicatedHostsResponseBody
	GetOrderId() *string
	SetRequestId(v string) *RenewDedicatedHostsResponseBody
	GetRequestId() *string
}

type RenewDedicatedHostsResponseBody struct {
	// The ID of the order.
	//
	// example:
	//
	// 23841229****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2A4EA075-CB5B-41B7-B0EB-70D339F6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RenewDedicatedHostsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RenewDedicatedHostsResponseBody) GoString() string {
	return s.String()
}

func (s *RenewDedicatedHostsResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *RenewDedicatedHostsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RenewDedicatedHostsResponseBody) SetOrderId(v string) *RenewDedicatedHostsResponseBody {
	s.OrderId = &v
	return s
}

func (s *RenewDedicatedHostsResponseBody) SetRequestId(v string) *RenewDedicatedHostsResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewDedicatedHostsResponseBody) Validate() error {
	return dara.Validate(s)
}
