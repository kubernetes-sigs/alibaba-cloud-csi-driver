// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiagnosticResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDiagnosticResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDiagnosticResultResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDiagnosticResultResponseBody) *DescribeDiagnosticResultResponse
	GetBody() *DescribeDiagnosticResultResponseBody
}

type DescribeDiagnosticResultResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDiagnosticResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDiagnosticResultResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosticResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosticResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDiagnosticResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDiagnosticResultResponse) GetBody() *DescribeDiagnosticResultResponseBody {
	return s.Body
}

func (s *DescribeDiagnosticResultResponse) SetHeaders(v map[string]*string) *DescribeDiagnosticResultResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiagnosticResultResponse) SetStatusCode(v int32) *DescribeDiagnosticResultResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDiagnosticResultResponse) SetBody(v *DescribeDiagnosticResultResponseBody) *DescribeDiagnosticResultResponse {
	s.Body = v
	return s
}

func (s *DescribeDiagnosticResultResponse) Validate() error {
	return dara.Validate(s)
}
