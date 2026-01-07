// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCapacityReservationInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCapacityReservationInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCapacityReservationInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCapacityReservationInstancesResponseBody) *DescribeCapacityReservationInstancesResponse
	GetBody() *DescribeCapacityReservationInstancesResponseBody
}

type DescribeCapacityReservationInstancesResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCapacityReservationInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCapacityReservationInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCapacityReservationInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeCapacityReservationInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCapacityReservationInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCapacityReservationInstancesResponse) GetBody() *DescribeCapacityReservationInstancesResponseBody {
	return s.Body
}

func (s *DescribeCapacityReservationInstancesResponse) SetHeaders(v map[string]*string) *DescribeCapacityReservationInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeCapacityReservationInstancesResponse) SetStatusCode(v int32) *DescribeCapacityReservationInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCapacityReservationInstancesResponse) SetBody(v *DescribeCapacityReservationInstancesResponseBody) *DescribeCapacityReservationInstancesResponse {
	s.Body = v
	return s
}

func (s *DescribeCapacityReservationInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
