// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReportInstancesStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReportInstancesStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReportInstancesStatusResponse
	GetStatusCode() *int32
	SetBody(v *ReportInstancesStatusResponseBody) *ReportInstancesStatusResponse
	GetBody() *ReportInstancesStatusResponseBody
}

type ReportInstancesStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReportInstancesStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReportInstancesStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ReportInstancesStatusResponse) GoString() string {
	return s.String()
}

func (s *ReportInstancesStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReportInstancesStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReportInstancesStatusResponse) GetBody() *ReportInstancesStatusResponseBody {
	return s.Body
}

func (s *ReportInstancesStatusResponse) SetHeaders(v map[string]*string) *ReportInstancesStatusResponse {
	s.Headers = v
	return s
}

func (s *ReportInstancesStatusResponse) SetStatusCode(v int32) *ReportInstancesStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ReportInstancesStatusResponse) SetBody(v *ReportInstancesStatusResponseBody) *ReportInstancesStatusResponse {
	s.Body = v
	return s
}

func (s *ReportInstancesStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
