// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTerminatePhysicalConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *TerminatePhysicalConnectionResponseBody
	GetRequestId() *string
}

type TerminatePhysicalConnectionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TerminatePhysicalConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TerminatePhysicalConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *TerminatePhysicalConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TerminatePhysicalConnectionResponseBody) SetRequestId(v string) *TerminatePhysicalConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *TerminatePhysicalConnectionResponseBody) Validate() error {
	return dara.Validate(s)
}
