// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBandwidthPackagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBandwidthPackagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBandwidthPackagesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBandwidthPackagesResponseBody) *DescribeBandwidthPackagesResponse
	GetBody() *DescribeBandwidthPackagesResponseBody
}

type DescribeBandwidthPackagesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBandwidthPackagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBandwidthPackagesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBandwidthPackagesResponse) GoString() string {
	return s.String()
}

func (s *DescribeBandwidthPackagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBandwidthPackagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBandwidthPackagesResponse) GetBody() *DescribeBandwidthPackagesResponseBody {
	return s.Body
}

func (s *DescribeBandwidthPackagesResponse) SetHeaders(v map[string]*string) *DescribeBandwidthPackagesResponse {
	s.Headers = v
	return s
}

func (s *DescribeBandwidthPackagesResponse) SetStatusCode(v int32) *DescribeBandwidthPackagesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBandwidthPackagesResponse) SetBody(v *DescribeBandwidthPackagesResponseBody) *DescribeBandwidthPackagesResponse {
	s.Body = v
	return s
}

func (s *DescribeBandwidthPackagesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
