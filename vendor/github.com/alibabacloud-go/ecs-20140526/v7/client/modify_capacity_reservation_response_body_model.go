// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCapacityReservationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyCapacityReservationResponseBody
	GetRequestId() *string
}

type ModifyCapacityReservationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 8455DD10-84F8-43C9-8365-5F448EB169B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCapacityReservationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCapacityReservationResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCapacityReservationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCapacityReservationResponseBody) SetRequestId(v string) *ModifyCapacityReservationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCapacityReservationResponseBody) Validate() error {
	return dara.Validate(s)
}
