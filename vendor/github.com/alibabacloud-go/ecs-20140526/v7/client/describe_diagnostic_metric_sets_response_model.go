// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiagnosticMetricSetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDiagnosticMetricSetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDiagnosticMetricSetsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDiagnosticMetricSetsResponseBody) *DescribeDiagnosticMetricSetsResponse
	GetBody() *DescribeDiagnosticMetricSetsResponseBody
}

type DescribeDiagnosticMetricSetsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDiagnosticMetricSetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDiagnosticMetricSetsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosticMetricSetsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosticMetricSetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDiagnosticMetricSetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDiagnosticMetricSetsResponse) GetBody() *DescribeDiagnosticMetricSetsResponseBody {
	return s.Body
}

func (s *DescribeDiagnosticMetricSetsResponse) SetHeaders(v map[string]*string) *DescribeDiagnosticMetricSetsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiagnosticMetricSetsResponse) SetStatusCode(v int32) *DescribeDiagnosticMetricSetsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDiagnosticMetricSetsResponse) SetBody(v *DescribeDiagnosticMetricSetsResponseBody) *DescribeDiagnosticMetricSetsResponse {
	s.Body = v
	return s
}

func (s *DescribeDiagnosticMetricSetsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
