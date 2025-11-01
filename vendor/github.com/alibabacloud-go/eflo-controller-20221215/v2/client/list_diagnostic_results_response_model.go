// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDiagnosticResultsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDiagnosticResultsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDiagnosticResultsResponse
	GetStatusCode() *int32
	SetBody(v *ListDiagnosticResultsResponseBody) *ListDiagnosticResultsResponse
	GetBody() *ListDiagnosticResultsResponseBody
}

type ListDiagnosticResultsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDiagnosticResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDiagnosticResultsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDiagnosticResultsResponse) GoString() string {
	return s.String()
}

func (s *ListDiagnosticResultsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDiagnosticResultsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDiagnosticResultsResponse) GetBody() *ListDiagnosticResultsResponseBody {
	return s.Body
}

func (s *ListDiagnosticResultsResponse) SetHeaders(v map[string]*string) *ListDiagnosticResultsResponse {
	s.Headers = v
	return s
}

func (s *ListDiagnosticResultsResponse) SetStatusCode(v int32) *ListDiagnosticResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDiagnosticResultsResponse) SetBody(v *ListDiagnosticResultsResponseBody) *ListDiagnosticResultsResponse {
	s.Body = v
	return s
}

func (s *ListDiagnosticResultsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
