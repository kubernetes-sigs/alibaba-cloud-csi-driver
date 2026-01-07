// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSecurityGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSecurityGroupsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSecurityGroupsResponseBody) *DescribeSecurityGroupsResponse
	GetBody() *DescribeSecurityGroupsResponseBody
}

type DescribeSecurityGroupsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSecurityGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSecurityGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSecurityGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSecurityGroupsResponse) GetBody() *DescribeSecurityGroupsResponseBody {
	return s.Body
}

func (s *DescribeSecurityGroupsResponse) SetHeaders(v map[string]*string) *DescribeSecurityGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSecurityGroupsResponse) SetStatusCode(v int32) *DescribeSecurityGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSecurityGroupsResponse) SetBody(v *DescribeSecurityGroupsResponseBody) *DescribeSecurityGroupsResponse {
	s.Body = v
	return s
}

func (s *DescribeSecurityGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
