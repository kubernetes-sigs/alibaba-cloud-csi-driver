// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeActivationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeActivationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeActivationsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeActivationsResponseBody) *DescribeActivationsResponse
	GetBody() *DescribeActivationsResponseBody
}

type DescribeActivationsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeActivationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeActivationsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeActivationsResponse) GoString() string {
	return s.String()
}

func (s *DescribeActivationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeActivationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeActivationsResponse) GetBody() *DescribeActivationsResponseBody {
	return s.Body
}

func (s *DescribeActivationsResponse) SetHeaders(v map[string]*string) *DescribeActivationsResponse {
	s.Headers = v
	return s
}

func (s *DescribeActivationsResponse) SetStatusCode(v int32) *DescribeActivationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeActivationsResponse) SetBody(v *DescribeActivationsResponseBody) *DescribeActivationsResponse {
	s.Body = v
	return s
}

func (s *DescribeActivationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
