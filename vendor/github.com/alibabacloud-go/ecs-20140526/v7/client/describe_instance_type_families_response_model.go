// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceTypeFamiliesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceTypeFamiliesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceTypeFamiliesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceTypeFamiliesResponseBody) *DescribeInstanceTypeFamiliesResponse
	GetBody() *DescribeInstanceTypeFamiliesResponseBody
}

type DescribeInstanceTypeFamiliesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceTypeFamiliesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceTypeFamiliesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceTypeFamiliesResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypeFamiliesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceTypeFamiliesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceTypeFamiliesResponse) GetBody() *DescribeInstanceTypeFamiliesResponseBody {
	return s.Body
}

func (s *DescribeInstanceTypeFamiliesResponse) SetHeaders(v map[string]*string) *DescribeInstanceTypeFamiliesResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceTypeFamiliesResponse) SetStatusCode(v int32) *DescribeInstanceTypeFamiliesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceTypeFamiliesResponse) SetBody(v *DescribeInstanceTypeFamiliesResponseBody) *DescribeInstanceTypeFamiliesResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceTypeFamiliesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
