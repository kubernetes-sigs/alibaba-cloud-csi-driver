// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDiagnosticReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDiagnosticReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDiagnosticReportResponse
	GetStatusCode() *int32
	SetBody(v *CreateDiagnosticReportResponseBody) *CreateDiagnosticReportResponse
	GetBody() *CreateDiagnosticReportResponseBody
}

type CreateDiagnosticReportResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDiagnosticReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDiagnosticReportResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDiagnosticReportResponse) GoString() string {
	return s.String()
}

func (s *CreateDiagnosticReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDiagnosticReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDiagnosticReportResponse) GetBody() *CreateDiagnosticReportResponseBody {
	return s.Body
}

func (s *CreateDiagnosticReportResponse) SetHeaders(v map[string]*string) *CreateDiagnosticReportResponse {
	s.Headers = v
	return s
}

func (s *CreateDiagnosticReportResponse) SetStatusCode(v int32) *CreateDiagnosticReportResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDiagnosticReportResponse) SetBody(v *CreateDiagnosticReportResponseBody) *CreateDiagnosticReportResponse {
	s.Body = v
	return s
}

func (s *CreateDiagnosticReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
