// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelSimulatedSystemEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelSimulatedSystemEventsResponseBody
	GetRequestId() *string
}

type CancelSimulatedSystemEventsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelSimulatedSystemEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelSimulatedSystemEventsResponseBody) GoString() string {
	return s.String()
}

func (s *CancelSimulatedSystemEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelSimulatedSystemEventsResponseBody) SetRequestId(v string) *CancelSimulatedSystemEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelSimulatedSystemEventsResponseBody) Validate() error {
	return dara.Validate(s)
}
