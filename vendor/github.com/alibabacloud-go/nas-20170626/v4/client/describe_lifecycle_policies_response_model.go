// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLifecyclePoliciesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLifecyclePoliciesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLifecyclePoliciesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLifecyclePoliciesResponseBody) *DescribeLifecyclePoliciesResponse
	GetBody() *DescribeLifecyclePoliciesResponseBody
}

type DescribeLifecyclePoliciesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLifecyclePoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLifecyclePoliciesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLifecyclePoliciesResponse) GoString() string {
	return s.String()
}

func (s *DescribeLifecyclePoliciesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLifecyclePoliciesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLifecyclePoliciesResponse) GetBody() *DescribeLifecyclePoliciesResponseBody {
	return s.Body
}

func (s *DescribeLifecyclePoliciesResponse) SetHeaders(v map[string]*string) *DescribeLifecyclePoliciesResponse {
	s.Headers = v
	return s
}

func (s *DescribeLifecyclePoliciesResponse) SetStatusCode(v int32) *DescribeLifecyclePoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLifecyclePoliciesResponse) SetBody(v *DescribeLifecyclePoliciesResponseBody) *DescribeLifecyclePoliciesResponse {
	s.Body = v
	return s
}

func (s *DescribeLifecyclePoliciesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
