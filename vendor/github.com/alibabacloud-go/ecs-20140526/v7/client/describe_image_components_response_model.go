// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageComponentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeImageComponentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeImageComponentsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeImageComponentsResponseBody) *DescribeImageComponentsResponse
	GetBody() *DescribeImageComponentsResponseBody
}

type DescribeImageComponentsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeImageComponentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeImageComponentsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageComponentsResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageComponentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeImageComponentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeImageComponentsResponse) GetBody() *DescribeImageComponentsResponseBody {
	return s.Body
}

func (s *DescribeImageComponentsResponse) SetHeaders(v map[string]*string) *DescribeImageComponentsResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageComponentsResponse) SetStatusCode(v int32) *DescribeImageComponentsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImageComponentsResponse) SetBody(v *DescribeImageComponentsResponseBody) *DescribeImageComponentsResponse {
	s.Body = v
	return s
}

func (s *DescribeImageComponentsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
