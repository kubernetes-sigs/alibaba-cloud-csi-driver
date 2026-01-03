// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnassignPrivateIpAddressesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnassignPrivateIpAddressesResponseBody
	GetRequestId() *string
}

type UnassignPrivateIpAddressesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnassignPrivateIpAddressesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnassignPrivateIpAddressesResponseBody) GoString() string {
	return s.String()
}

func (s *UnassignPrivateIpAddressesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnassignPrivateIpAddressesResponseBody) SetRequestId(v string) *UnassignPrivateIpAddressesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnassignPrivateIpAddressesResponseBody) Validate() error {
	return dara.Validate(s)
}
