// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVirtualBorderRouterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteVirtualBorderRouterResponseBody
	GetRequestId() *string
}

type DeleteVirtualBorderRouterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVirtualBorderRouterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVirtualBorderRouterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVirtualBorderRouterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVirtualBorderRouterResponseBody) SetRequestId(v string) *DeleteVirtualBorderRouterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVirtualBorderRouterResponseBody) Validate() error {
	return dara.Validate(s)
}
