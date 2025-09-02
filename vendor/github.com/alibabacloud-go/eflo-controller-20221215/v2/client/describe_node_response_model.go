// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNodeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNodeResponseBody) *DescribeNodeResponse
	GetBody() *DescribeNodeResponseBody
}

type DescribeNodeResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNodeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNodeResponse) GoString() string {
	return s.String()
}

func (s *DescribeNodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNodeResponse) GetBody() *DescribeNodeResponseBody {
	return s.Body
}

func (s *DescribeNodeResponse) SetHeaders(v map[string]*string) *DescribeNodeResponse {
	s.Headers = v
	return s
}

func (s *DescribeNodeResponse) SetStatusCode(v int32) *DescribeNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNodeResponse) SetBody(v *DescribeNodeResponseBody) *DescribeNodeResponse {
	s.Body = v
	return s
}

func (s *DescribeNodeResponse) Validate() error {
	return dara.Validate(s)
}
