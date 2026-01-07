// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSavingsPlanPriceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSavingsPlanPriceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSavingsPlanPriceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSavingsPlanPriceResponseBody) *DescribeSavingsPlanPriceResponse
	GetBody() *DescribeSavingsPlanPriceResponseBody
}

type DescribeSavingsPlanPriceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSavingsPlanPriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSavingsPlanPriceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSavingsPlanPriceResponse) GoString() string {
	return s.String()
}

func (s *DescribeSavingsPlanPriceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSavingsPlanPriceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSavingsPlanPriceResponse) GetBody() *DescribeSavingsPlanPriceResponseBody {
	return s.Body
}

func (s *DescribeSavingsPlanPriceResponse) SetHeaders(v map[string]*string) *DescribeSavingsPlanPriceResponse {
	s.Headers = v
	return s
}

func (s *DescribeSavingsPlanPriceResponse) SetStatusCode(v int32) *DescribeSavingsPlanPriceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSavingsPlanPriceResponse) SetBody(v *DescribeSavingsPlanPriceResponseBody) *DescribeSavingsPlanPriceResponse {
	s.Body = v
	return s
}

func (s *DescribeSavingsPlanPriceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
