// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityGroupAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSecurityGroupAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSecurityGroupAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSecurityGroupAttributeResponseBody) *DescribeSecurityGroupAttributeResponse
	GetBody() *DescribeSecurityGroupAttributeResponseBody
}

type DescribeSecurityGroupAttributeResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSecurityGroupAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSecurityGroupAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityGroupAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSecurityGroupAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSecurityGroupAttributeResponse) GetBody() *DescribeSecurityGroupAttributeResponseBody {
	return s.Body
}

func (s *DescribeSecurityGroupAttributeResponse) SetHeaders(v map[string]*string) *DescribeSecurityGroupAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeSecurityGroupAttributeResponse) SetStatusCode(v int32) *DescribeSecurityGroupAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponse) SetBody(v *DescribeSecurityGroupAttributeResponseBody) *DescribeSecurityGroupAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeSecurityGroupAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
