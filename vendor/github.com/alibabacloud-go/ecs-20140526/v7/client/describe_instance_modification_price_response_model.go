// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceModificationPriceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceModificationPriceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceModificationPriceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceModificationPriceResponseBody) *DescribeInstanceModificationPriceResponse
	GetBody() *DescribeInstanceModificationPriceResponseBody
}

type DescribeInstanceModificationPriceResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceModificationPriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceModificationPriceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceModificationPriceResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceModificationPriceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceModificationPriceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceModificationPriceResponse) GetBody() *DescribeInstanceModificationPriceResponseBody {
	return s.Body
}

func (s *DescribeInstanceModificationPriceResponse) SetHeaders(v map[string]*string) *DescribeInstanceModificationPriceResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceModificationPriceResponse) SetStatusCode(v int32) *DescribeInstanceModificationPriceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceModificationPriceResponse) SetBody(v *DescribeInstanceModificationPriceResponseBody) *DescribeInstanceModificationPriceResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceModificationPriceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
