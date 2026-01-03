// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnassociateEipAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnassociateEipAddressResponseBody
	GetRequestId() *string
}

type UnassociateEipAddressResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnassociateEipAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnassociateEipAddressResponseBody) GoString() string {
	return s.String()
}

func (s *UnassociateEipAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnassociateEipAddressResponseBody) SetRequestId(v string) *UnassociateEipAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnassociateEipAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
