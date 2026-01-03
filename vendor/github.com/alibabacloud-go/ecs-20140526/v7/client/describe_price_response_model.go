// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePriceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePriceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePriceResponse
	GetStatusCode() *int32
	SetBody(v *DescribePriceResponseBody) *DescribePriceResponse
	GetBody() *DescribePriceResponseBody
}

type DescribePriceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePriceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponse) GoString() string {
	return s.String()
}

func (s *DescribePriceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePriceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePriceResponse) GetBody() *DescribePriceResponseBody {
	return s.Body
}

func (s *DescribePriceResponse) SetHeaders(v map[string]*string) *DescribePriceResponse {
	s.Headers = v
	return s
}

func (s *DescribePriceResponse) SetStatusCode(v int32) *DescribePriceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePriceResponse) SetBody(v *DescribePriceResponseBody) *DescribePriceResponse {
	s.Body = v
	return s
}

func (s *DescribePriceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
