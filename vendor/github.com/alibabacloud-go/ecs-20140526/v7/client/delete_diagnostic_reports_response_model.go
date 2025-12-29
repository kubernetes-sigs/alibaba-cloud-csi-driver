// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDiagnosticReportsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDiagnosticReportsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDiagnosticReportsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDiagnosticReportsResponseBody) *DeleteDiagnosticReportsResponse
	GetBody() *DeleteDiagnosticReportsResponseBody
}

type DeleteDiagnosticReportsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDiagnosticReportsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDiagnosticReportsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDiagnosticReportsResponse) GoString() string {
	return s.String()
}

func (s *DeleteDiagnosticReportsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDiagnosticReportsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDiagnosticReportsResponse) GetBody() *DeleteDiagnosticReportsResponseBody {
	return s.Body
}

func (s *DeleteDiagnosticReportsResponse) SetHeaders(v map[string]*string) *DeleteDiagnosticReportsResponse {
	s.Headers = v
	return s
}

func (s *DeleteDiagnosticReportsResponse) SetStatusCode(v int32) *DeleteDiagnosticReportsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDiagnosticReportsResponse) SetBody(v *DeleteDiagnosticReportsResponseBody) *DeleteDiagnosticReportsResponse {
	s.Body = v
	return s
}

func (s *DeleteDiagnosticReportsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
