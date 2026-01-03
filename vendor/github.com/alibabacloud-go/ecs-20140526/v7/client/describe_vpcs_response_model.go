// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVpcsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVpcsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVpcsResponseBody) *DescribeVpcsResponse
	GetBody() *DescribeVpcsResponseBody
}

type DescribeVpcsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVpcsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVpcsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcsResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpcsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVpcsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVpcsResponse) GetBody() *DescribeVpcsResponseBody {
	return s.Body
}

func (s *DescribeVpcsResponse) SetHeaders(v map[string]*string) *DescribeVpcsResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpcsResponse) SetStatusCode(v int32) *DescribeVpcsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpcsResponse) SetBody(v *DescribeVpcsResponseBody) *DescribeVpcsResponse {
	s.Body = v
	return s
}

func (s *DescribeVpcsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
