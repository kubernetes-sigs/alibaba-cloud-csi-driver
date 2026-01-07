// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEipAddressAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyEipAddressAttributeResponseBody
	GetRequestId() *string
}

type ModifyEipAddressAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyEipAddressAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyEipAddressAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyEipAddressAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyEipAddressAttributeResponseBody) SetRequestId(v string) *ModifyEipAddressAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyEipAddressAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
