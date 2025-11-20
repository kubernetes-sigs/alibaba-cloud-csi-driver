// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInvocationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInvocationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInvocationsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInvocationsResponseBody) *DescribeInvocationsResponse
	GetBody() *DescribeInvocationsResponseBody
}

type DescribeInvocationsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInvocationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInvocationsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvocationsResponse) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInvocationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInvocationsResponse) GetBody() *DescribeInvocationsResponseBody {
	return s.Body
}

func (s *DescribeInvocationsResponse) SetHeaders(v map[string]*string) *DescribeInvocationsResponse {
	s.Headers = v
	return s
}

func (s *DescribeInvocationsResponse) SetStatusCode(v int32) *DescribeInvocationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInvocationsResponse) SetBody(v *DescribeInvocationsResponseBody) *DescribeInvocationsResponse {
	s.Body = v
	return s
}

func (s *DescribeInvocationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
