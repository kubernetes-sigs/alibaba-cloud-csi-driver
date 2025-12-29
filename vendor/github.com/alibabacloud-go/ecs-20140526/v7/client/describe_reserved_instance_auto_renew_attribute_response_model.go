// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeReservedInstanceAutoRenewAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeReservedInstanceAutoRenewAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeReservedInstanceAutoRenewAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeReservedInstanceAutoRenewAttributeResponseBody) *DescribeReservedInstanceAutoRenewAttributeResponse
	GetBody() *DescribeReservedInstanceAutoRenewAttributeResponseBody
}

type DescribeReservedInstanceAutoRenewAttributeResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeReservedInstanceAutoRenewAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeReservedInstanceAutoRenewAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeReservedInstanceAutoRenewAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeReservedInstanceAutoRenewAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeReservedInstanceAutoRenewAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeReservedInstanceAutoRenewAttributeResponse) GetBody() *DescribeReservedInstanceAutoRenewAttributeResponseBody {
	return s.Body
}

func (s *DescribeReservedInstanceAutoRenewAttributeResponse) SetHeaders(v map[string]*string) *DescribeReservedInstanceAutoRenewAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeReservedInstanceAutoRenewAttributeResponse) SetStatusCode(v int32) *DescribeReservedInstanceAutoRenewAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeReservedInstanceAutoRenewAttributeResponse) SetBody(v *DescribeReservedInstanceAutoRenewAttributeResponseBody) *DescribeReservedInstanceAutoRenewAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeReservedInstanceAutoRenewAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
