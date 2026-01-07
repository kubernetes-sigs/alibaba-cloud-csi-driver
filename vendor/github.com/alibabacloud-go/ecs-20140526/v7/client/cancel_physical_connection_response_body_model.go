// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelPhysicalConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelPhysicalConnectionResponseBody
	GetRequestId() *string
}

type CancelPhysicalConnectionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelPhysicalConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelPhysicalConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *CancelPhysicalConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelPhysicalConnectionResponseBody) SetRequestId(v string) *CancelPhysicalConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelPhysicalConnectionResponseBody) Validate() error {
	return dara.Validate(s)
}
