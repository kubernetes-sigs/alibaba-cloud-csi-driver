// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSpotAdviceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSpotAdviceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSpotAdviceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSpotAdviceResponseBody) *DescribeSpotAdviceResponse
	GetBody() *DescribeSpotAdviceResponseBody
}

type DescribeSpotAdviceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSpotAdviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSpotAdviceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSpotAdviceResponse) GoString() string {
	return s.String()
}

func (s *DescribeSpotAdviceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSpotAdviceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSpotAdviceResponse) GetBody() *DescribeSpotAdviceResponseBody {
	return s.Body
}

func (s *DescribeSpotAdviceResponse) SetHeaders(v map[string]*string) *DescribeSpotAdviceResponse {
	s.Headers = v
	return s
}

func (s *DescribeSpotAdviceResponse) SetStatusCode(v int32) *DescribeSpotAdviceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSpotAdviceResponse) SetBody(v *DescribeSpotAdviceResponseBody) *DescribeSpotAdviceResponse {
	s.Body = v
	return s
}

func (s *DescribeSpotAdviceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
