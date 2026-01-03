// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSpotPriceHistoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSpotPriceHistoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSpotPriceHistoryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSpotPriceHistoryResponseBody) *DescribeSpotPriceHistoryResponse
	GetBody() *DescribeSpotPriceHistoryResponseBody
}

type DescribeSpotPriceHistoryResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSpotPriceHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSpotPriceHistoryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSpotPriceHistoryResponse) GoString() string {
	return s.String()
}

func (s *DescribeSpotPriceHistoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSpotPriceHistoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSpotPriceHistoryResponse) GetBody() *DescribeSpotPriceHistoryResponseBody {
	return s.Body
}

func (s *DescribeSpotPriceHistoryResponse) SetHeaders(v map[string]*string) *DescribeSpotPriceHistoryResponse {
	s.Headers = v
	return s
}

func (s *DescribeSpotPriceHistoryResponse) SetStatusCode(v int32) *DescribeSpotPriceHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSpotPriceHistoryResponse) SetBody(v *DescribeSpotPriceHistoryResponseBody) *DescribeSpotPriceHistoryResponse {
	s.Body = v
	return s
}

func (s *DescribeSpotPriceHistoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
