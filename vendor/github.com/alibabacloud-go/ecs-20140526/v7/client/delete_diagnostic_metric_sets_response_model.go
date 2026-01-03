// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDiagnosticMetricSetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDiagnosticMetricSetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDiagnosticMetricSetsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDiagnosticMetricSetsResponseBody) *DeleteDiagnosticMetricSetsResponse
	GetBody() *DeleteDiagnosticMetricSetsResponseBody
}

type DeleteDiagnosticMetricSetsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDiagnosticMetricSetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDiagnosticMetricSetsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDiagnosticMetricSetsResponse) GoString() string {
	return s.String()
}

func (s *DeleteDiagnosticMetricSetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDiagnosticMetricSetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDiagnosticMetricSetsResponse) GetBody() *DeleteDiagnosticMetricSetsResponseBody {
	return s.Body
}

func (s *DeleteDiagnosticMetricSetsResponse) SetHeaders(v map[string]*string) *DeleteDiagnosticMetricSetsResponse {
	s.Headers = v
	return s
}

func (s *DeleteDiagnosticMetricSetsResponse) SetStatusCode(v int32) *DeleteDiagnosticMetricSetsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDiagnosticMetricSetsResponse) SetBody(v *DeleteDiagnosticMetricSetsResponseBody) *DeleteDiagnosticMetricSetsResponse {
	s.Body = v
	return s
}

func (s *DeleteDiagnosticMetricSetsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
