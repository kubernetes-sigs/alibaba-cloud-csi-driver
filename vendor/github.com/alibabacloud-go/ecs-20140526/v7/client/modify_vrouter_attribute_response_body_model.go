// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVRouterAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyVRouterAttributeResponseBody
	GetRequestId() *string
}

type ModifyVRouterAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyVRouterAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyVRouterAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVRouterAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyVRouterAttributeResponseBody) SetRequestId(v string) *ModifyVRouterAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyVRouterAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
