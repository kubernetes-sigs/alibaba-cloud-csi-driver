// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLimitationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLimitationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLimitationResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLimitationResponseBody) *DescribeLimitationResponse
	GetBody() *DescribeLimitationResponseBody
}

type DescribeLimitationResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLimitationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLimitationResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLimitationResponse) GoString() string {
	return s.String()
}

func (s *DescribeLimitationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLimitationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLimitationResponse) GetBody() *DescribeLimitationResponseBody {
	return s.Body
}

func (s *DescribeLimitationResponse) SetHeaders(v map[string]*string) *DescribeLimitationResponse {
	s.Headers = v
	return s
}

func (s *DescribeLimitationResponse) SetStatusCode(v int32) *DescribeLimitationResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLimitationResponse) SetBody(v *DescribeLimitationResponseBody) *DescribeLimitationResponse {
	s.Body = v
	return s
}

func (s *DescribeLimitationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
