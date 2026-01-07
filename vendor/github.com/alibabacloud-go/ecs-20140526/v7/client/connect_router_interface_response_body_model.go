// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConnectRouterInterfaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ConnectRouterInterfaceResponseBody
	GetRequestId() *string
}

type ConnectRouterInterfaceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConnectRouterInterfaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConnectRouterInterfaceResponseBody) GoString() string {
	return s.String()
}

func (s *ConnectRouterInterfaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConnectRouterInterfaceResponseBody) SetRequestId(v string) *ConnectRouterInterfaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConnectRouterInterfaceResponseBody) Validate() error {
	return dara.Validate(s)
}
