// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPhysicalConnectionAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyPhysicalConnectionAttributeResponseBody
	GetRequestId() *string
}

type ModifyPhysicalConnectionAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyPhysicalConnectionAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyPhysicalConnectionAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPhysicalConnectionAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyPhysicalConnectionAttributeResponseBody) SetRequestId(v string) *ModifyPhysicalConnectionAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyPhysicalConnectionAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
