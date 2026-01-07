// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeReservedInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeReservedInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeReservedInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeReservedInstancesResponseBody) *DescribeReservedInstancesResponse
	GetBody() *DescribeReservedInstancesResponseBody
}

type DescribeReservedInstancesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeReservedInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeReservedInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeReservedInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeReservedInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeReservedInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeReservedInstancesResponse) GetBody() *DescribeReservedInstancesResponseBody {
	return s.Body
}

func (s *DescribeReservedInstancesResponse) SetHeaders(v map[string]*string) *DescribeReservedInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeReservedInstancesResponse) SetStatusCode(v int32) *DescribeReservedInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeReservedInstancesResponse) SetBody(v *DescribeReservedInstancesResponseBody) *DescribeReservedInstancesResponse {
	s.Body = v
	return s
}

func (s *DescribeReservedInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
