// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCapacityReservationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCapacityReservationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCapacityReservationResponse
	GetStatusCode() *int32
	SetBody(v *CreateCapacityReservationResponseBody) *CreateCapacityReservationResponse
	GetBody() *CreateCapacityReservationResponseBody
}

type CreateCapacityReservationResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCapacityReservationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCapacityReservationResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCapacityReservationResponse) GoString() string {
	return s.String()
}

func (s *CreateCapacityReservationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCapacityReservationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCapacityReservationResponse) GetBody() *CreateCapacityReservationResponseBody {
	return s.Body
}

func (s *CreateCapacityReservationResponse) SetHeaders(v map[string]*string) *CreateCapacityReservationResponse {
	s.Headers = v
	return s
}

func (s *CreateCapacityReservationResponse) SetStatusCode(v int32) *CreateCapacityReservationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCapacityReservationResponse) SetBody(v *CreateCapacityReservationResponseBody) *CreateCapacityReservationResponse {
	s.Body = v
	return s
}

func (s *CreateCapacityReservationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
