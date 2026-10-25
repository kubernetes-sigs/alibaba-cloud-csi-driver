// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFaultReportsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFaultReportsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFaultReportsResponse
	GetStatusCode() *int32
	SetBody(v *ListFaultReportsResponseBody) *ListFaultReportsResponse
	GetBody() *ListFaultReportsResponseBody
}

type ListFaultReportsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFaultReportsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFaultReportsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFaultReportsResponse) GoString() string {
	return s.String()
}

func (s *ListFaultReportsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFaultReportsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFaultReportsResponse) GetBody() *ListFaultReportsResponseBody {
	return s.Body
}

func (s *ListFaultReportsResponse) SetHeaders(v map[string]*string) *ListFaultReportsResponse {
	s.Headers = v
	return s
}

func (s *ListFaultReportsResponse) SetStatusCode(v int32) *ListFaultReportsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFaultReportsResponse) SetBody(v *ListFaultReportsResponseBody) *ListFaultReportsResponse {
	s.Body = v
	return s
}

func (s *ListFaultReportsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
