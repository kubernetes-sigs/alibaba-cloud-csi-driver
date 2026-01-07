// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTaskAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTaskAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTaskAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTaskAttributeResponseBody) *DescribeTaskAttributeResponse
	GetBody() *DescribeTaskAttributeResponseBody
}

type DescribeTaskAttributeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTaskAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTaskAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTaskAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeTaskAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTaskAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTaskAttributeResponse) GetBody() *DescribeTaskAttributeResponseBody {
	return s.Body
}

func (s *DescribeTaskAttributeResponse) SetHeaders(v map[string]*string) *DescribeTaskAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeTaskAttributeResponse) SetStatusCode(v int32) *DescribeTaskAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTaskAttributeResponse) SetBody(v *DescribeTaskAttributeResponseBody) *DescribeTaskAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeTaskAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
