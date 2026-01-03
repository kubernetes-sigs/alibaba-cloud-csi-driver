// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNatGatewaysResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNatGatewaysResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNatGatewaysResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNatGatewaysResponseBody) *DescribeNatGatewaysResponse
	GetBody() *DescribeNatGatewaysResponseBody
}

type DescribeNatGatewaysResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNatGatewaysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNatGatewaysResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatGatewaysResponse) GoString() string {
	return s.String()
}

func (s *DescribeNatGatewaysResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNatGatewaysResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNatGatewaysResponse) GetBody() *DescribeNatGatewaysResponseBody {
	return s.Body
}

func (s *DescribeNatGatewaysResponse) SetHeaders(v map[string]*string) *DescribeNatGatewaysResponse {
	s.Headers = v
	return s
}

func (s *DescribeNatGatewaysResponse) SetStatusCode(v int32) *DescribeNatGatewaysResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNatGatewaysResponse) SetBody(v *DescribeNatGatewaysResponseBody) *DescribeNatGatewaysResponse {
	s.Body = v
	return s
}

func (s *DescribeNatGatewaysResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
