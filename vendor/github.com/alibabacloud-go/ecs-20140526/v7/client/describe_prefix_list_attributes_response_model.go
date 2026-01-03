// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePrefixListAttributesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePrefixListAttributesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePrefixListAttributesResponse
	GetStatusCode() *int32
	SetBody(v *DescribePrefixListAttributesResponseBody) *DescribePrefixListAttributesResponse
	GetBody() *DescribePrefixListAttributesResponseBody
}

type DescribePrefixListAttributesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePrefixListAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePrefixListAttributesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePrefixListAttributesResponse) GoString() string {
	return s.String()
}

func (s *DescribePrefixListAttributesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePrefixListAttributesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePrefixListAttributesResponse) GetBody() *DescribePrefixListAttributesResponseBody {
	return s.Body
}

func (s *DescribePrefixListAttributesResponse) SetHeaders(v map[string]*string) *DescribePrefixListAttributesResponse {
	s.Headers = v
	return s
}

func (s *DescribePrefixListAttributesResponse) SetStatusCode(v int32) *DescribePrefixListAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePrefixListAttributesResponse) SetBody(v *DescribePrefixListAttributesResponseBody) *DescribePrefixListAttributesResponse {
	s.Body = v
	return s
}

func (s *DescribePrefixListAttributesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
