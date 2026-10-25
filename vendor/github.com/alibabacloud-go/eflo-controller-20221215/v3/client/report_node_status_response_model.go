// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReportNodeStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReportNodeStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReportNodeStatusResponse
	GetStatusCode() *int32
	SetBody(v *ReportNodeStatusResponseBody) *ReportNodeStatusResponse
	GetBody() *ReportNodeStatusResponseBody
}

type ReportNodeStatusResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReportNodeStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReportNodeStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ReportNodeStatusResponse) GoString() string {
	return s.String()
}

func (s *ReportNodeStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReportNodeStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReportNodeStatusResponse) GetBody() *ReportNodeStatusResponseBody {
	return s.Body
}

func (s *ReportNodeStatusResponse) SetHeaders(v map[string]*string) *ReportNodeStatusResponse {
	s.Headers = v
	return s
}

func (s *ReportNodeStatusResponse) SetStatusCode(v int32) *ReportNodeStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ReportNodeStatusResponse) SetBody(v *ReportNodeStatusResponseBody) *ReportNodeStatusResponse {
	s.Body = v
	return s
}

func (s *ReportNodeStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
