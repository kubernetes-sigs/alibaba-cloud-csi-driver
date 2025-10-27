// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProtocolServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteProtocolServiceResponseBody
	GetRequestId() *string
}

type DeleteProtocolServiceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 98696EF0-1607-4E9D-B01D-F20930B6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteProtocolServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteProtocolServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProtocolServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteProtocolServiceResponseBody) SetRequestId(v string) *DeleteProtocolServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteProtocolServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
