// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseCapacityReservationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ReleaseCapacityReservationResponseBody
	GetRequestId() *string
}

type ReleaseCapacityReservationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseCapacityReservationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReleaseCapacityReservationResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseCapacityReservationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReleaseCapacityReservationResponseBody) SetRequestId(v string) *ReleaseCapacityReservationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseCapacityReservationResponseBody) Validate() error {
	return dara.Validate(s)
}
