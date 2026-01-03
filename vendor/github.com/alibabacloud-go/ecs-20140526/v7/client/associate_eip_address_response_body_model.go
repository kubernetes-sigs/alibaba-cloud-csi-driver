// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateEipAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AssociateEipAddressResponseBody
	GetRequestId() *string
}

type AssociateEipAddressResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateEipAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssociateEipAddressResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateEipAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssociateEipAddressResponseBody) SetRequestId(v string) *AssociateEipAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssociateEipAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
