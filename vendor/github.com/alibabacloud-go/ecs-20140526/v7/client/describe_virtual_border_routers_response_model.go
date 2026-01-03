// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVirtualBorderRoutersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVirtualBorderRoutersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVirtualBorderRoutersResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVirtualBorderRoutersResponseBody) *DescribeVirtualBorderRoutersResponse
	GetBody() *DescribeVirtualBorderRoutersResponseBody
}

type DescribeVirtualBorderRoutersResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVirtualBorderRoutersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVirtualBorderRoutersResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVirtualBorderRoutersResponse) GoString() string {
	return s.String()
}

func (s *DescribeVirtualBorderRoutersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVirtualBorderRoutersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVirtualBorderRoutersResponse) GetBody() *DescribeVirtualBorderRoutersResponseBody {
	return s.Body
}

func (s *DescribeVirtualBorderRoutersResponse) SetHeaders(v map[string]*string) *DescribeVirtualBorderRoutersResponse {
	s.Headers = v
	return s
}

func (s *DescribeVirtualBorderRoutersResponse) SetStatusCode(v int32) *DescribeVirtualBorderRoutersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponse) SetBody(v *DescribeVirtualBorderRoutersResponseBody) *DescribeVirtualBorderRoutersResponse {
	s.Body = v
	return s
}

func (s *DescribeVirtualBorderRoutersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
