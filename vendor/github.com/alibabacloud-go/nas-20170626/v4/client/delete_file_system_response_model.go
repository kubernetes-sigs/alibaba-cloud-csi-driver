// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFileSystemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFileSystemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFileSystemResponse
	GetStatusCode() *int32
	SetBody(v *DeleteFileSystemResponseBody) *DeleteFileSystemResponse
	GetBody() *DeleteFileSystemResponseBody
}

type DeleteFileSystemResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFileSystemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFileSystemResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFileSystemResponse) GoString() string {
	return s.String()
}

func (s *DeleteFileSystemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFileSystemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFileSystemResponse) GetBody() *DeleteFileSystemResponseBody {
	return s.Body
}

func (s *DeleteFileSystemResponse) SetHeaders(v map[string]*string) *DeleteFileSystemResponse {
	s.Headers = v
	return s
}

func (s *DeleteFileSystemResponse) SetStatusCode(v int32) *DeleteFileSystemResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFileSystemResponse) SetBody(v *DeleteFileSystemResponseBody) *DeleteFileSystemResponse {
	s.Body = v
	return s
}

func (s *DeleteFileSystemResponse) Validate() error {
	return dara.Validate(s)
}
