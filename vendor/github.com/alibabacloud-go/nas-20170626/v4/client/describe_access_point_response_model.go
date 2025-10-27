// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccessPointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAccessPointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAccessPointResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAccessPointResponseBody) *DescribeAccessPointResponse
	GetBody() *DescribeAccessPointResponseBody
}

type DescribeAccessPointResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAccessPointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAccessPointResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessPointResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccessPointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAccessPointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAccessPointResponse) GetBody() *DescribeAccessPointResponseBody {
	return s.Body
}

func (s *DescribeAccessPointResponse) SetHeaders(v map[string]*string) *DescribeAccessPointResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccessPointResponse) SetStatusCode(v int32) *DescribeAccessPointResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAccessPointResponse) SetBody(v *DescribeAccessPointResponseBody) *DescribeAccessPointResponse {
	s.Body = v
	return s
}

func (s *DescribeAccessPointResponse) Validate() error {
	return dara.Validate(s)
}
