// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCapacityReservationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCapacityReservationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCapacityReservationsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCapacityReservationsResponseBody) *DescribeCapacityReservationsResponse
	GetBody() *DescribeCapacityReservationsResponseBody
}

type DescribeCapacityReservationsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCapacityReservationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCapacityReservationsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCapacityReservationsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCapacityReservationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCapacityReservationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCapacityReservationsResponse) GetBody() *DescribeCapacityReservationsResponseBody {
	return s.Body
}

func (s *DescribeCapacityReservationsResponse) SetHeaders(v map[string]*string) *DescribeCapacityReservationsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCapacityReservationsResponse) SetStatusCode(v int32) *DescribeCapacityReservationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCapacityReservationsResponse) SetBody(v *DescribeCapacityReservationsResponseBody) *DescribeCapacityReservationsResponse {
	s.Body = v
	return s
}

func (s *DescribeCapacityReservationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
