// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFaultReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFaultReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFaultReportResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFaultReportResponseBody) *DescribeFaultReportResponse
	GetBody() *DescribeFaultReportResponseBody
}

type DescribeFaultReportResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFaultReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFaultReportResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFaultReportResponse) GoString() string {
	return s.String()
}

func (s *DescribeFaultReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFaultReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFaultReportResponse) GetBody() *DescribeFaultReportResponseBody {
	return s.Body
}

func (s *DescribeFaultReportResponse) SetHeaders(v map[string]*string) *DescribeFaultReportResponse {
	s.Headers = v
	return s
}

func (s *DescribeFaultReportResponse) SetStatusCode(v int32) *DescribeFaultReportResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFaultReportResponse) SetBody(v *DescribeFaultReportResponseBody) *DescribeFaultReportResponse {
	s.Body = v
	return s
}

func (s *DescribeFaultReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
