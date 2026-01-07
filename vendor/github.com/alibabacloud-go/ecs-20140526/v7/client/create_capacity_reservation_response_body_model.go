// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCapacityReservationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPrivatePoolOptionsId(v string) *CreateCapacityReservationResponseBody
	GetPrivatePoolOptionsId() *string
	SetRequestId(v string) *CreateCapacityReservationResponseBody
	GetRequestId() *string
}

type CreateCapacityReservationResponseBody struct {
	// The capacity reservation ID.
	//
	// example:
	//
	// crp-bp67acfmxazb4****
	PrivatePoolOptionsId *string `json:"PrivatePoolOptionsId,omitempty" xml:"PrivatePoolOptionsId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCapacityReservationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCapacityReservationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCapacityReservationResponseBody) GetPrivatePoolOptionsId() *string {
	return s.PrivatePoolOptionsId
}

func (s *CreateCapacityReservationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCapacityReservationResponseBody) SetPrivatePoolOptionsId(v string) *CreateCapacityReservationResponseBody {
	s.PrivatePoolOptionsId = &v
	return s
}

func (s *CreateCapacityReservationResponseBody) SetRequestId(v string) *CreateCapacityReservationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCapacityReservationResponseBody) Validate() error {
	return dara.Validate(s)
}
