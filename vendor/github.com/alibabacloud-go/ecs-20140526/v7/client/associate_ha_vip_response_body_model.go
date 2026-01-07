// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateHaVipResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AssociateHaVipResponseBody
	GetRequestId() *string
}

type AssociateHaVipResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateHaVipResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssociateHaVipResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateHaVipResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssociateHaVipResponseBody) SetRequestId(v string) *AssociateHaVipResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssociateHaVipResponseBody) Validate() error {
	return dara.Validate(s)
}
