// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProtocolServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeProtocolServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeProtocolServiceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeProtocolServiceResponseBody) *DescribeProtocolServiceResponse
	GetBody() *DescribeProtocolServiceResponseBody
}

type DescribeProtocolServiceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeProtocolServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeProtocolServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeProtocolServiceResponse) GoString() string {
	return s.String()
}

func (s *DescribeProtocolServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeProtocolServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeProtocolServiceResponse) GetBody() *DescribeProtocolServiceResponseBody {
	return s.Body
}

func (s *DescribeProtocolServiceResponse) SetHeaders(v map[string]*string) *DescribeProtocolServiceResponse {
	s.Headers = v
	return s
}

func (s *DescribeProtocolServiceResponse) SetStatusCode(v int32) *DescribeProtocolServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeProtocolServiceResponse) SetBody(v *DescribeProtocolServiceResponseBody) *DescribeProtocolServiceResponse {
	s.Body = v
	return s
}

func (s *DescribeProtocolServiceResponse) Validate() error {
	return dara.Validate(s)
}
