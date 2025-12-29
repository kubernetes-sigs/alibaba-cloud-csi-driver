// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiagnosticMetricsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDiagnosticMetricsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDiagnosticMetricsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDiagnosticMetricsResponseBody) *DescribeDiagnosticMetricsResponse
	GetBody() *DescribeDiagnosticMetricsResponseBody
}

type DescribeDiagnosticMetricsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDiagnosticMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDiagnosticMetricsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosticMetricsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosticMetricsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDiagnosticMetricsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDiagnosticMetricsResponse) GetBody() *DescribeDiagnosticMetricsResponseBody {
	return s.Body
}

func (s *DescribeDiagnosticMetricsResponse) SetHeaders(v map[string]*string) *DescribeDiagnosticMetricsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiagnosticMetricsResponse) SetStatusCode(v int32) *DescribeDiagnosticMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDiagnosticMetricsResponse) SetBody(v *DescribeDiagnosticMetricsResponseBody) *DescribeDiagnosticMetricsResponse {
	s.Body = v
	return s
}

func (s *DescribeDiagnosticMetricsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
