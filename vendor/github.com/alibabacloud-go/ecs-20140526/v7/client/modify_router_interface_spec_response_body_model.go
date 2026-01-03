// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRouterInterfaceSpecResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyRouterInterfaceSpecResponseBody
	GetRequestId() *string
	SetSpec(v string) *ModifyRouterInterfaceSpecResponseBody
	GetSpec() *string
}

type ModifyRouterInterfaceSpecResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Spec      *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s ModifyRouterInterfaceSpecResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyRouterInterfaceSpecResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyRouterInterfaceSpecResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyRouterInterfaceSpecResponseBody) GetSpec() *string {
	return s.Spec
}

func (s *ModifyRouterInterfaceSpecResponseBody) SetRequestId(v string) *ModifyRouterInterfaceSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyRouterInterfaceSpecResponseBody) SetSpec(v string) *ModifyRouterInterfaceSpecResponseBody {
	s.Spec = &v
	return s
}

func (s *ModifyRouterInterfaceSpecResponseBody) Validate() error {
	return dara.Validate(s)
}
