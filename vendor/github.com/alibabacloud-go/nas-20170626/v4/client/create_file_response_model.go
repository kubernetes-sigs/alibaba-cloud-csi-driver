// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateFileResponse
	GetStatusCode() *int32
	SetBody(v *CreateFileResponseBody) *CreateFileResponse
	GetBody() *CreateFileResponseBody
}

type CreateFileResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFileResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateFileResponse) GoString() string {
	return s.String()
}

func (s *CreateFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateFileResponse) GetBody() *CreateFileResponseBody {
	return s.Body
}

func (s *CreateFileResponse) SetHeaders(v map[string]*string) *CreateFileResponse {
	s.Headers = v
	return s
}

func (s *CreateFileResponse) SetStatusCode(v int32) *CreateFileResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFileResponse) SetBody(v *CreateFileResponseBody) *CreateFileResponse {
	s.Body = v
	return s
}

func (s *CreateFileResponse) Validate() error {
	return dara.Validate(s)
}
