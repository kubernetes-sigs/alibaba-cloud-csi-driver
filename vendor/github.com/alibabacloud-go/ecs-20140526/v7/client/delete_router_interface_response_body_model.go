// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRouterInterfaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteRouterInterfaceResponseBody
	GetRequestId() *string
}

type DeleteRouterInterfaceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRouterInterfaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRouterInterfaceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRouterInterfaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRouterInterfaceResponseBody) SetRequestId(v string) *DeleteRouterInterfaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRouterInterfaceResponseBody) Validate() error {
	return dara.Validate(s)
}
