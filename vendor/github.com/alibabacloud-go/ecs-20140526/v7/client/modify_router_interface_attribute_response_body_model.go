// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRouterInterfaceAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyRouterInterfaceAttributeResponseBody
	GetRequestId() *string
}

type ModifyRouterInterfaceAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyRouterInterfaceAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyRouterInterfaceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyRouterInterfaceAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyRouterInterfaceAttributeResponseBody) SetRequestId(v string) *ModifyRouterInterfaceAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyRouterInterfaceAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
