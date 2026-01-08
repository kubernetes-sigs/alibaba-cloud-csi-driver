// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogAnalysisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLogAnalysisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLogAnalysisResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLogAnalysisResponseBody) *DescribeLogAnalysisResponse
	GetBody() *DescribeLogAnalysisResponseBody
}

type DescribeLogAnalysisResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLogAnalysisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLogAnalysisResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogAnalysisResponse) GoString() string {
	return s.String()
}

func (s *DescribeLogAnalysisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLogAnalysisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLogAnalysisResponse) GetBody() *DescribeLogAnalysisResponseBody {
	return s.Body
}

func (s *DescribeLogAnalysisResponse) SetHeaders(v map[string]*string) *DescribeLogAnalysisResponse {
	s.Headers = v
	return s
}

func (s *DescribeLogAnalysisResponse) SetStatusCode(v int32) *DescribeLogAnalysisResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLogAnalysisResponse) SetBody(v *DescribeLogAnalysisResponseBody) *DescribeLogAnalysisResponse {
	s.Body = v
	return s
}

func (s *DescribeLogAnalysisResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
