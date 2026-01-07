// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiagnosticReportsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDiagnosticReportsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDiagnosticReportsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDiagnosticReportsResponseBody) *DescribeDiagnosticReportsResponse
	GetBody() *DescribeDiagnosticReportsResponseBody
}

type DescribeDiagnosticReportsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDiagnosticReportsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDiagnosticReportsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosticReportsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosticReportsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDiagnosticReportsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDiagnosticReportsResponse) GetBody() *DescribeDiagnosticReportsResponseBody {
	return s.Body
}

func (s *DescribeDiagnosticReportsResponse) SetHeaders(v map[string]*string) *DescribeDiagnosticReportsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiagnosticReportsResponse) SetStatusCode(v int32) *DescribeDiagnosticReportsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDiagnosticReportsResponse) SetBody(v *DescribeDiagnosticReportsResponseBody) *DescribeDiagnosticReportsResponse {
	s.Body = v
	return s
}

func (s *DescribeDiagnosticReportsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
