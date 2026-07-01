// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLifecyclePolicyLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLifecyclePolicyLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLifecyclePolicyLogsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLifecyclePolicyLogsResponseBody) *DescribeLifecyclePolicyLogsResponse
	GetBody() *DescribeLifecyclePolicyLogsResponseBody
}

type DescribeLifecyclePolicyLogsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLifecyclePolicyLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLifecyclePolicyLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLifecyclePolicyLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeLifecyclePolicyLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLifecyclePolicyLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLifecyclePolicyLogsResponse) GetBody() *DescribeLifecyclePolicyLogsResponseBody {
	return s.Body
}

func (s *DescribeLifecyclePolicyLogsResponse) SetHeaders(v map[string]*string) *DescribeLifecyclePolicyLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeLifecyclePolicyLogsResponse) SetStatusCode(v int32) *DescribeLifecyclePolicyLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLifecyclePolicyLogsResponse) SetBody(v *DescribeLifecyclePolicyLogsResponseBody) *DescribeLifecyclePolicyLogsResponse {
	s.Body = v
	return s
}

func (s *DescribeLifecyclePolicyLogsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
