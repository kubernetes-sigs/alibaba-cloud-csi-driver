// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDiagnosticMetricSetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDiagnosticMetricSetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDiagnosticMetricSetResponse
	GetStatusCode() *int32
	SetBody(v *CreateDiagnosticMetricSetResponseBody) *CreateDiagnosticMetricSetResponse
	GetBody() *CreateDiagnosticMetricSetResponseBody
}

type CreateDiagnosticMetricSetResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDiagnosticMetricSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDiagnosticMetricSetResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDiagnosticMetricSetResponse) GoString() string {
	return s.String()
}

func (s *CreateDiagnosticMetricSetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDiagnosticMetricSetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDiagnosticMetricSetResponse) GetBody() *CreateDiagnosticMetricSetResponseBody {
	return s.Body
}

func (s *CreateDiagnosticMetricSetResponse) SetHeaders(v map[string]*string) *CreateDiagnosticMetricSetResponse {
	s.Headers = v
	return s
}

func (s *CreateDiagnosticMetricSetResponse) SetStatusCode(v int32) *CreateDiagnosticMetricSetResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDiagnosticMetricSetResponse) SetBody(v *CreateDiagnosticMetricSetResponseBody) *CreateDiagnosticMetricSetResponse {
	s.Body = v
	return s
}

func (s *CreateDiagnosticMetricSetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
