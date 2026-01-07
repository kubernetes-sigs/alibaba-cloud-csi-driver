// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccessRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAccessRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAccessRulesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAccessRulesResponseBody) *DescribeAccessRulesResponse
	GetBody() *DescribeAccessRulesResponseBody
}

type DescribeAccessRulesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAccessRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAccessRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccessRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAccessRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAccessRulesResponse) GetBody() *DescribeAccessRulesResponseBody {
	return s.Body
}

func (s *DescribeAccessRulesResponse) SetHeaders(v map[string]*string) *DescribeAccessRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccessRulesResponse) SetStatusCode(v int32) *DescribeAccessRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAccessRulesResponse) SetBody(v *DescribeAccessRulesResponseBody) *DescribeAccessRulesResponse {
	s.Body = v
	return s
}

func (s *DescribeAccessRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
