// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFileSystemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateFileSystemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateFileSystemResponse
	GetStatusCode() *int32
	SetBody(v *CreateFileSystemResponseBody) *CreateFileSystemResponse
	GetBody() *CreateFileSystemResponseBody
}

type CreateFileSystemResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFileSystemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFileSystemResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateFileSystemResponse) GoString() string {
	return s.String()
}

func (s *CreateFileSystemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateFileSystemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateFileSystemResponse) GetBody() *CreateFileSystemResponseBody {
	return s.Body
}

func (s *CreateFileSystemResponse) SetHeaders(v map[string]*string) *CreateFileSystemResponse {
	s.Headers = v
	return s
}

func (s *CreateFileSystemResponse) SetStatusCode(v int32) *CreateFileSystemResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFileSystemResponse) SetBody(v *CreateFileSystemResponseBody) *CreateFileSystemResponse {
	s.Body = v
	return s
}

func (s *CreateFileSystemResponse) Validate() error {
	return dara.Validate(s)
}
