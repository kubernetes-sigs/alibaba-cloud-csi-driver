// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDirResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDirResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDirResponse
	GetStatusCode() *int32
	SetBody(v *CreateDirResponseBody) *CreateDirResponse
	GetBody() *CreateDirResponseBody
}

type CreateDirResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDirResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDirResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDirResponse) GoString() string {
	return s.String()
}

func (s *CreateDirResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDirResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDirResponse) GetBody() *CreateDirResponseBody {
	return s.Body
}

func (s *CreateDirResponse) SetHeaders(v map[string]*string) *CreateDirResponse {
	s.Headers = v
	return s
}

func (s *CreateDirResponse) SetStatusCode(v int32) *CreateDirResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDirResponse) SetBody(v *CreateDirResponseBody) *CreateDirResponse {
	s.Body = v
	return s
}

func (s *CreateDirResponse) Validate() error {
	return dara.Validate(s)
}
