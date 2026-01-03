// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDiagnosticMetricSetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDiagnosticMetricSetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDiagnosticMetricSetResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDiagnosticMetricSetResponseBody) *ModifyDiagnosticMetricSetResponse
	GetBody() *ModifyDiagnosticMetricSetResponseBody
}

type ModifyDiagnosticMetricSetResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDiagnosticMetricSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDiagnosticMetricSetResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDiagnosticMetricSetResponse) GoString() string {
	return s.String()
}

func (s *ModifyDiagnosticMetricSetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDiagnosticMetricSetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDiagnosticMetricSetResponse) GetBody() *ModifyDiagnosticMetricSetResponseBody {
	return s.Body
}

func (s *ModifyDiagnosticMetricSetResponse) SetHeaders(v map[string]*string) *ModifyDiagnosticMetricSetResponse {
	s.Headers = v
	return s
}

func (s *ModifyDiagnosticMetricSetResponse) SetStatusCode(v int32) *ModifyDiagnosticMetricSetResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDiagnosticMetricSetResponse) SetBody(v *ModifyDiagnosticMetricSetResponseBody) *ModifyDiagnosticMetricSetResponse {
	s.Body = v
	return s
}

func (s *ModifyDiagnosticMetricSetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
