// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBandwidthLimitationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBandwidthLimitationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBandwidthLimitationResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBandwidthLimitationResponseBody) *DescribeBandwidthLimitationResponse
	GetBody() *DescribeBandwidthLimitationResponseBody
}

type DescribeBandwidthLimitationResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBandwidthLimitationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBandwidthLimitationResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBandwidthLimitationResponse) GoString() string {
	return s.String()
}

func (s *DescribeBandwidthLimitationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBandwidthLimitationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBandwidthLimitationResponse) GetBody() *DescribeBandwidthLimitationResponseBody {
	return s.Body
}

func (s *DescribeBandwidthLimitationResponse) SetHeaders(v map[string]*string) *DescribeBandwidthLimitationResponse {
	s.Headers = v
	return s
}

func (s *DescribeBandwidthLimitationResponse) SetStatusCode(v int32) *DescribeBandwidthLimitationResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBandwidthLimitationResponse) SetBody(v *DescribeBandwidthLimitationResponseBody) *DescribeBandwidthLimitationResponse {
	s.Body = v
	return s
}

func (s *DescribeBandwidthLimitationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
