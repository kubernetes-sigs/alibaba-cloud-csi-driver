// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceTypesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceTypesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceTypesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceTypesResponseBody) *DescribeInstanceTypesResponse
	GetBody() *DescribeInstanceTypesResponseBody
}

type DescribeInstanceTypesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceTypesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceTypesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceTypesResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceTypesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceTypesResponse) GetBody() *DescribeInstanceTypesResponseBody {
	return s.Body
}

func (s *DescribeInstanceTypesResponse) SetHeaders(v map[string]*string) *DescribeInstanceTypesResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceTypesResponse) SetStatusCode(v int32) *DescribeInstanceTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceTypesResponse) SetBody(v *DescribeInstanceTypesResponseBody) *DescribeInstanceTypesResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceTypesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
