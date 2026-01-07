// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAcceptInquiredSystemEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AcceptInquiredSystemEventResponseBody
	GetRequestId() *string
}

type AcceptInquiredSystemEventResponseBody struct {
	// example:
	//
	// 4DD56CA6-6D75-4D33-BE34-E4A44EBE1C3D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AcceptInquiredSystemEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AcceptInquiredSystemEventResponseBody) GoString() string {
	return s.String()
}

func (s *AcceptInquiredSystemEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AcceptInquiredSystemEventResponseBody) SetRequestId(v string) *AcceptInquiredSystemEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *AcceptInquiredSystemEventResponseBody) Validate() error {
	return dara.Validate(s)
}
