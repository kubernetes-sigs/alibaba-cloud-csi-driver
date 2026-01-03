// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseCapacityReservationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReleaseCapacityReservationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReleaseCapacityReservationResponse
	GetStatusCode() *int32
	SetBody(v *ReleaseCapacityReservationResponseBody) *ReleaseCapacityReservationResponse
	GetBody() *ReleaseCapacityReservationResponseBody
}

type ReleaseCapacityReservationResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseCapacityReservationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseCapacityReservationResponse) String() string {
	return dara.Prettify(s)
}

func (s ReleaseCapacityReservationResponse) GoString() string {
	return s.String()
}

func (s *ReleaseCapacityReservationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReleaseCapacityReservationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReleaseCapacityReservationResponse) GetBody() *ReleaseCapacityReservationResponseBody {
	return s.Body
}

func (s *ReleaseCapacityReservationResponse) SetHeaders(v map[string]*string) *ReleaseCapacityReservationResponse {
	s.Headers = v
	return s
}

func (s *ReleaseCapacityReservationResponse) SetStatusCode(v int32) *ReleaseCapacityReservationResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseCapacityReservationResponse) SetBody(v *ReleaseCapacityReservationResponseBody) *ReleaseCapacityReservationResponse {
	s.Body = v
	return s
}

func (s *ReleaseCapacityReservationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
