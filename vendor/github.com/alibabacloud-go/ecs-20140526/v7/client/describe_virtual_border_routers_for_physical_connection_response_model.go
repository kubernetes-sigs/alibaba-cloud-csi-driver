// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVirtualBorderRoutersForPhysicalConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVirtualBorderRoutersForPhysicalConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVirtualBorderRoutersForPhysicalConnectionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBody) *DescribeVirtualBorderRoutersForPhysicalConnectionResponse
	GetBody() *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBody
}

type DescribeVirtualBorderRoutersForPhysicalConnectionResponse struct {
	Headers    map[string]*string                                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVirtualBorderRoutersForPhysicalConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVirtualBorderRoutersForPhysicalConnectionResponse) GoString() string {
	return s.String()
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponse) GetBody() *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBody {
	return s.Body
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponse) SetHeaders(v map[string]*string) *DescribeVirtualBorderRoutersForPhysicalConnectionResponse {
	s.Headers = v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponse) SetStatusCode(v int32) *DescribeVirtualBorderRoutersForPhysicalConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponse) SetBody(v *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBody) *DescribeVirtualBorderRoutersForPhysicalConnectionResponse {
	s.Body = v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
