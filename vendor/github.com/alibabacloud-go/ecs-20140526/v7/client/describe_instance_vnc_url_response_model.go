// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceVncUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceVncUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceVncUrlResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceVncUrlResponseBody) *DescribeInstanceVncUrlResponse
	GetBody() *DescribeInstanceVncUrlResponseBody
}

type DescribeInstanceVncUrlResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceVncUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceVncUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceVncUrlResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceVncUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceVncUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceVncUrlResponse) GetBody() *DescribeInstanceVncUrlResponseBody {
	return s.Body
}

func (s *DescribeInstanceVncUrlResponse) SetHeaders(v map[string]*string) *DescribeInstanceVncUrlResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceVncUrlResponse) SetStatusCode(v int32) *DescribeInstanceVncUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceVncUrlResponse) SetBody(v *DescribeInstanceVncUrlResponseBody) *DescribeInstanceVncUrlResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceVncUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
