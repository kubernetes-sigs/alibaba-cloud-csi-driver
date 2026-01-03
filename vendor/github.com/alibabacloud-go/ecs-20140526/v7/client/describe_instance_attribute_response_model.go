// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceAttributeResponseBody) *DescribeInstanceAttributeResponse
	GetBody() *DescribeInstanceAttributeResponseBody
}

type DescribeInstanceAttributeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceAttributeResponse) GetBody() *DescribeInstanceAttributeResponseBody {
	return s.Body
}

func (s *DescribeInstanceAttributeResponse) SetHeaders(v map[string]*string) *DescribeInstanceAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceAttributeResponse) SetStatusCode(v int32) *DescribeInstanceAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceAttributeResponse) SetBody(v *DescribeInstanceAttributeResponseBody) *DescribeInstanceAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
