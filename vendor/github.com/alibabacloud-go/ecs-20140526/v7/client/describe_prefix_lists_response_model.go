// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePrefixListsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePrefixListsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePrefixListsResponse
	GetStatusCode() *int32
	SetBody(v *DescribePrefixListsResponseBody) *DescribePrefixListsResponse
	GetBody() *DescribePrefixListsResponseBody
}

type DescribePrefixListsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePrefixListsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePrefixListsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePrefixListsResponse) GoString() string {
	return s.String()
}

func (s *DescribePrefixListsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePrefixListsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePrefixListsResponse) GetBody() *DescribePrefixListsResponseBody {
	return s.Body
}

func (s *DescribePrefixListsResponse) SetHeaders(v map[string]*string) *DescribePrefixListsResponse {
	s.Headers = v
	return s
}

func (s *DescribePrefixListsResponse) SetStatusCode(v int32) *DescribePrefixListsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePrefixListsResponse) SetBody(v *DescribePrefixListsResponseBody) *DescribePrefixListsResponse {
	s.Body = v
	return s
}

func (s *DescribePrefixListsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
