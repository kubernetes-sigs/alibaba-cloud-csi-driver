// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePrefixListAssociationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePrefixListAssociationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePrefixListAssociationsResponse
	GetStatusCode() *int32
	SetBody(v *DescribePrefixListAssociationsResponseBody) *DescribePrefixListAssociationsResponse
	GetBody() *DescribePrefixListAssociationsResponseBody
}

type DescribePrefixListAssociationsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePrefixListAssociationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePrefixListAssociationsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePrefixListAssociationsResponse) GoString() string {
	return s.String()
}

func (s *DescribePrefixListAssociationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePrefixListAssociationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePrefixListAssociationsResponse) GetBody() *DescribePrefixListAssociationsResponseBody {
	return s.Body
}

func (s *DescribePrefixListAssociationsResponse) SetHeaders(v map[string]*string) *DescribePrefixListAssociationsResponse {
	s.Headers = v
	return s
}

func (s *DescribePrefixListAssociationsResponse) SetStatusCode(v int32) *DescribePrefixListAssociationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePrefixListAssociationsResponse) SetBody(v *DescribePrefixListAssociationsResponseBody) *DescribePrefixListAssociationsResponse {
	s.Body = v
	return s
}

func (s *DescribePrefixListAssociationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
