// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReportNodesStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReportNodesStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReportNodesStatusResponse
	GetStatusCode() *int32
	SetBody(v *ReportNodesStatusResponseBody) *ReportNodesStatusResponse
	GetBody() *ReportNodesStatusResponseBody
}

type ReportNodesStatusResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReportNodesStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReportNodesStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ReportNodesStatusResponse) GoString() string {
	return s.String()
}

func (s *ReportNodesStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReportNodesStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReportNodesStatusResponse) GetBody() *ReportNodesStatusResponseBody {
	return s.Body
}

func (s *ReportNodesStatusResponse) SetHeaders(v map[string]*string) *ReportNodesStatusResponse {
	s.Headers = v
	return s
}

func (s *ReportNodesStatusResponse) SetStatusCode(v int32) *ReportNodesStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ReportNodesStatusResponse) SetBody(v *ReportNodesStatusResponseBody) *ReportNodesStatusResponse {
	s.Body = v
	return s
}

func (s *ReportNodesStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
