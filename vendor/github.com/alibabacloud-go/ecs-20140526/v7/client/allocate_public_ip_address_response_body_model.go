// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocatePublicIpAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIpAddress(v string) *AllocatePublicIpAddressResponseBody
	GetIpAddress() *string
	SetRequestId(v string) *AllocatePublicIpAddressResponseBody
	GetRequestId() *string
}

type AllocatePublicIpAddressResponseBody struct {
	// The public IP address.
	//
	// example:
	//
	// ``112.124.**.**``
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AllocatePublicIpAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AllocatePublicIpAddressResponseBody) GoString() string {
	return s.String()
}

func (s *AllocatePublicIpAddressResponseBody) GetIpAddress() *string {
	return s.IpAddress
}

func (s *AllocatePublicIpAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AllocatePublicIpAddressResponseBody) SetIpAddress(v string) *AllocatePublicIpAddressResponseBody {
	s.IpAddress = &v
	return s
}

func (s *AllocatePublicIpAddressResponseBody) SetRequestId(v string) *AllocatePublicIpAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *AllocatePublicIpAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
