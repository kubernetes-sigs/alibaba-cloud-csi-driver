// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVscsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVscsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVscsResponse
	GetStatusCode() *int32
	SetBody(v *ListVscsResponseBody) *ListVscsResponse
	GetBody() *ListVscsResponseBody
}

type ListVscsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVscsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVscsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVscsResponse) GoString() string {
	return s.String()
}

func (s *ListVscsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVscsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVscsResponse) GetBody() *ListVscsResponseBody {
	return s.Body
}

func (s *ListVscsResponse) SetHeaders(v map[string]*string) *ListVscsResponse {
	s.Headers = v
	return s
}

func (s *ListVscsResponse) SetStatusCode(v int32) *ListVscsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVscsResponse) SetBody(v *ListVscsResponseBody) *ListVscsResponse {
	s.Body = v
	return s
}

func (s *ListVscsResponse) Validate() error {
	return dara.Validate(s)
}
