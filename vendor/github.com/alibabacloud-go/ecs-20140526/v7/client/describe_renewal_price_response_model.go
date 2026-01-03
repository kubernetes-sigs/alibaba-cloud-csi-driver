// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRenewalPriceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRenewalPriceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRenewalPriceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRenewalPriceResponseBody) *DescribeRenewalPriceResponse
	GetBody() *DescribeRenewalPriceResponseBody
}

type DescribeRenewalPriceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRenewalPriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRenewalPriceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenewalPriceResponse) GoString() string {
	return s.String()
}

func (s *DescribeRenewalPriceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRenewalPriceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRenewalPriceResponse) GetBody() *DescribeRenewalPriceResponseBody {
	return s.Body
}

func (s *DescribeRenewalPriceResponse) SetHeaders(v map[string]*string) *DescribeRenewalPriceResponse {
	s.Headers = v
	return s
}

func (s *DescribeRenewalPriceResponse) SetStatusCode(v int32) *DescribeRenewalPriceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRenewalPriceResponse) SetBody(v *DescribeRenewalPriceResponseBody) *DescribeRenewalPriceResponse {
	s.Body = v
	return s
}

func (s *DescribeRenewalPriceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
