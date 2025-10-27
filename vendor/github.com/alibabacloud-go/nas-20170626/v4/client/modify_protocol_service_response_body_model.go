// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyProtocolServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyProtocolServiceResponseBody
	GetRequestId() *string
}

type ModifyProtocolServiceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 98696EF0-1607-4E9D-B01D-F20930B6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyProtocolServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyProtocolServiceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyProtocolServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyProtocolServiceResponseBody) SetRequestId(v string) *ModifyProtocolServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyProtocolServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
