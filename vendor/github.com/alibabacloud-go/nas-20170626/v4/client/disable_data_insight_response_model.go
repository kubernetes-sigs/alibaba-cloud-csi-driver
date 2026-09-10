// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableDataInsightResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableDataInsightResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableDataInsightResponse
	GetStatusCode() *int32
	SetBody(v *DisableDataInsightResponseBody) *DisableDataInsightResponse
	GetBody() *DisableDataInsightResponseBody
}

type DisableDataInsightResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableDataInsightResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableDataInsightResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableDataInsightResponse) GoString() string {
	return s.String()
}

func (s *DisableDataInsightResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableDataInsightResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableDataInsightResponse) GetBody() *DisableDataInsightResponseBody {
	return s.Body
}

func (s *DisableDataInsightResponse) SetHeaders(v map[string]*string) *DisableDataInsightResponse {
	s.Headers = v
	return s
}

func (s *DisableDataInsightResponse) SetStatusCode(v int32) *DisableDataInsightResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableDataInsightResponse) SetBody(v *DisableDataInsightResponseBody) *DisableDataInsightResponse {
	s.Body = v
	return s
}

func (s *DisableDataInsightResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
