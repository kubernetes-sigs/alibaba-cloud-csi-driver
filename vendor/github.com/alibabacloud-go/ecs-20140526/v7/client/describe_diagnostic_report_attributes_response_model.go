// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiagnosticReportAttributesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDiagnosticReportAttributesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDiagnosticReportAttributesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDiagnosticReportAttributesResponseBody) *DescribeDiagnosticReportAttributesResponse
	GetBody() *DescribeDiagnosticReportAttributesResponseBody
}

type DescribeDiagnosticReportAttributesResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDiagnosticReportAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDiagnosticReportAttributesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosticReportAttributesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosticReportAttributesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDiagnosticReportAttributesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDiagnosticReportAttributesResponse) GetBody() *DescribeDiagnosticReportAttributesResponseBody {
	return s.Body
}

func (s *DescribeDiagnosticReportAttributesResponse) SetHeaders(v map[string]*string) *DescribeDiagnosticReportAttributesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiagnosticReportAttributesResponse) SetStatusCode(v int32) *DescribeDiagnosticReportAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDiagnosticReportAttributesResponse) SetBody(v *DescribeDiagnosticReportAttributesResponseBody) *DescribeDiagnosticReportAttributesResponse {
	s.Body = v
	return s
}

func (s *DescribeDiagnosticReportAttributesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
