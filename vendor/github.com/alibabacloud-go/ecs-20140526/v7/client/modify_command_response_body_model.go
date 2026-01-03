// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCommandResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyCommandResponseBody
	GetRequestId() *string
}

type ModifyCommandResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0DE9B41E-EF0D-40A0-BB43-37749C5BDA9C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCommandResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCommandResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCommandResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCommandResponseBody) SetRequestId(v string) *ModifyCommandResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCommandResponseBody) Validate() error {
	return dara.Validate(s)
}
