// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecoverVirtualBorderRouterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RecoverVirtualBorderRouterResponseBody
	GetRequestId() *string
}

type RecoverVirtualBorderRouterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecoverVirtualBorderRouterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RecoverVirtualBorderRouterResponseBody) GoString() string {
	return s.String()
}

func (s *RecoverVirtualBorderRouterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RecoverVirtualBorderRouterResponseBody) SetRequestId(v string) *RecoverVirtualBorderRouterResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecoverVirtualBorderRouterResponseBody) Validate() error {
	return dara.Validate(s)
}
