// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProtocolServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProtocolServiceId(v string) *CreateProtocolServiceResponseBody
	GetProtocolServiceId() *string
	SetRequestId(v string) *CreateProtocolServiceResponseBody
	GetRequestId() *string
}

type CreateProtocolServiceResponseBody struct {
	// The ID of the protocol service.
	//
	// example:
	//
	// ptc-123****
	ProtocolServiceId *string `json:"ProtocolServiceId,omitempty" xml:"ProtocolServiceId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 98696EF0-1607-4E9D-B01D-F20930B6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateProtocolServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateProtocolServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProtocolServiceResponseBody) GetProtocolServiceId() *string {
	return s.ProtocolServiceId
}

func (s *CreateProtocolServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateProtocolServiceResponseBody) SetProtocolServiceId(v string) *CreateProtocolServiceResponseBody {
	s.ProtocolServiceId = &v
	return s
}

func (s *CreateProtocolServiceResponseBody) SetRequestId(v string) *CreateProtocolServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProtocolServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
