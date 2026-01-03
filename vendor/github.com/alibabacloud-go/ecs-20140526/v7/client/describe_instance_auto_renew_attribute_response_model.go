// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceAutoRenewAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceAutoRenewAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceAutoRenewAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceAutoRenewAttributeResponseBody) *DescribeInstanceAutoRenewAttributeResponse
	GetBody() *DescribeInstanceAutoRenewAttributeResponseBody
}

type DescribeInstanceAutoRenewAttributeResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceAutoRenewAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceAutoRenewAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAutoRenewAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAutoRenewAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceAutoRenewAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceAutoRenewAttributeResponse) GetBody() *DescribeInstanceAutoRenewAttributeResponseBody {
	return s.Body
}

func (s *DescribeInstanceAutoRenewAttributeResponse) SetHeaders(v map[string]*string) *DescribeInstanceAutoRenewAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceAutoRenewAttributeResponse) SetStatusCode(v int32) *DescribeInstanceAutoRenewAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceAutoRenewAttributeResponse) SetBody(v *DescribeInstanceAutoRenewAttributeResponseBody) *DescribeInstanceAutoRenewAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceAutoRenewAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
