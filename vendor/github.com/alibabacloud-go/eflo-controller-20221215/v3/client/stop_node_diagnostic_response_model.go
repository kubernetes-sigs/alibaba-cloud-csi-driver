// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopNodeDiagnosticResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopNodeDiagnosticResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopNodeDiagnosticResponse
	GetStatusCode() *int32
	SetBody(v *StopNodeDiagnosticResponseBody) *StopNodeDiagnosticResponse
	GetBody() *StopNodeDiagnosticResponseBody
}

type StopNodeDiagnosticResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopNodeDiagnosticResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopNodeDiagnosticResponse) String() string {
	return dara.Prettify(s)
}

func (s StopNodeDiagnosticResponse) GoString() string {
	return s.String()
}

func (s *StopNodeDiagnosticResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopNodeDiagnosticResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopNodeDiagnosticResponse) GetBody() *StopNodeDiagnosticResponseBody {
	return s.Body
}

func (s *StopNodeDiagnosticResponse) SetHeaders(v map[string]*string) *StopNodeDiagnosticResponse {
	s.Headers = v
	return s
}

func (s *StopNodeDiagnosticResponse) SetStatusCode(v int32) *StopNodeDiagnosticResponse {
	s.StatusCode = &v
	return s
}

func (s *StopNodeDiagnosticResponse) SetBody(v *StopNodeDiagnosticResponseBody) *StopNodeDiagnosticResponse {
	s.Body = v
	return s
}

func (s *StopNodeDiagnosticResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
