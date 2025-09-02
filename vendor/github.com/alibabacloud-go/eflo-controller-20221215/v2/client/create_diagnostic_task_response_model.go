// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDiagnosticTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDiagnosticTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDiagnosticTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateDiagnosticTaskResponseBody) *CreateDiagnosticTaskResponse
	GetBody() *CreateDiagnosticTaskResponseBody
}

type CreateDiagnosticTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDiagnosticTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDiagnosticTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDiagnosticTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateDiagnosticTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDiagnosticTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDiagnosticTaskResponse) GetBody() *CreateDiagnosticTaskResponseBody {
	return s.Body
}

func (s *CreateDiagnosticTaskResponse) SetHeaders(v map[string]*string) *CreateDiagnosticTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateDiagnosticTaskResponse) SetStatusCode(v int32) *CreateDiagnosticTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDiagnosticTaskResponse) SetBody(v *CreateDiagnosticTaskResponseBody) *CreateDiagnosticTaskResponse {
	s.Body = v
	return s
}

func (s *CreateDiagnosticTaskResponse) Validate() error {
	return dara.Validate(s)
}
