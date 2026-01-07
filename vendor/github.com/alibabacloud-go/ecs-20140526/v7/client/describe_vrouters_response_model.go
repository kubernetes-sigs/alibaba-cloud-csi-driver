// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVRoutersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVRoutersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVRoutersResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVRoutersResponseBody) *DescribeVRoutersResponse
	GetBody() *DescribeVRoutersResponseBody
}

type DescribeVRoutersResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVRoutersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVRoutersResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVRoutersResponse) GoString() string {
	return s.String()
}

func (s *DescribeVRoutersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVRoutersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVRoutersResponse) GetBody() *DescribeVRoutersResponseBody {
	return s.Body
}

func (s *DescribeVRoutersResponse) SetHeaders(v map[string]*string) *DescribeVRoutersResponse {
	s.Headers = v
	return s
}

func (s *DescribeVRoutersResponse) SetStatusCode(v int32) *DescribeVRoutersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVRoutersResponse) SetBody(v *DescribeVRoutersResponseBody) *DescribeVRoutersResponse {
	s.Body = v
	return s
}

func (s *DescribeVRoutersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
