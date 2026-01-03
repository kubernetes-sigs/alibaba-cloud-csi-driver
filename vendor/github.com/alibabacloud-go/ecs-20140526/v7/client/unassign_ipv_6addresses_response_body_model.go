// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnassignIpv6AddressesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnassignIpv6AddressesResponseBody
	GetRequestId() *string
}

type UnassignIpv6AddressesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnassignIpv6AddressesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnassignIpv6AddressesResponseBody) GoString() string {
	return s.String()
}

func (s *UnassignIpv6AddressesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnassignIpv6AddressesResponseBody) SetRequestId(v string) *UnassignIpv6AddressesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnassignIpv6AddressesResponseBody) Validate() error {
	return dara.Validate(s)
}
