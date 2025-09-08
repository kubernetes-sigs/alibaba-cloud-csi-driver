// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNodeTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNodeTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNodeTypeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNodeTypeResponseBody) *DescribeNodeTypeResponse
	GetBody() *DescribeNodeTypeResponseBody
}

type DescribeNodeTypeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNodeTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNodeTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNodeTypeResponse) GoString() string {
	return s.String()
}

func (s *DescribeNodeTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNodeTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNodeTypeResponse) GetBody() *DescribeNodeTypeResponseBody {
	return s.Body
}

func (s *DescribeNodeTypeResponse) SetHeaders(v map[string]*string) *DescribeNodeTypeResponse {
	s.Headers = v
	return s
}

func (s *DescribeNodeTypeResponse) SetStatusCode(v int32) *DescribeNodeTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNodeTypeResponse) SetBody(v *DescribeNodeTypeResponseBody) *DescribeNodeTypeResponse {
	s.Body = v
	return s
}

func (s *DescribeNodeTypeResponse) Validate() error {
	return dara.Validate(s)
}
