// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActivateRouterInterfaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ActivateRouterInterfaceResponseBody
	GetRequestId() *string
}

type ActivateRouterInterfaceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ActivateRouterInterfaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ActivateRouterInterfaceResponseBody) GoString() string {
	return s.String()
}

func (s *ActivateRouterInterfaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ActivateRouterInterfaceResponseBody) SetRequestId(v string) *ActivateRouterInterfaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ActivateRouterInterfaceResponseBody) Validate() error {
	return dara.Validate(s)
}
