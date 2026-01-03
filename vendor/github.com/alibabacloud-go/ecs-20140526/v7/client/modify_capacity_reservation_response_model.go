// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCapacityReservationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCapacityReservationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCapacityReservationResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCapacityReservationResponseBody) *ModifyCapacityReservationResponse
	GetBody() *ModifyCapacityReservationResponseBody
}

type ModifyCapacityReservationResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCapacityReservationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCapacityReservationResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCapacityReservationResponse) GoString() string {
	return s.String()
}

func (s *ModifyCapacityReservationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCapacityReservationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCapacityReservationResponse) GetBody() *ModifyCapacityReservationResponseBody {
	return s.Body
}

func (s *ModifyCapacityReservationResponse) SetHeaders(v map[string]*string) *ModifyCapacityReservationResponse {
	s.Headers = v
	return s
}

func (s *ModifyCapacityReservationResponse) SetStatusCode(v int32) *ModifyCapacityReservationResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCapacityReservationResponse) SetBody(v *ModifyCapacityReservationResponseBody) *ModifyCapacityReservationResponse {
	s.Body = v
	return s
}

func (s *ModifyCapacityReservationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
