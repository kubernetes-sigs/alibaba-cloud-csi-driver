// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetFileSystemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetFileSystemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetFileSystemResponse
	GetStatusCode() *int32
	SetBody(v *ResetFileSystemResponseBody) *ResetFileSystemResponse
	GetBody() *ResetFileSystemResponseBody
}

type ResetFileSystemResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetFileSystemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetFileSystemResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetFileSystemResponse) GoString() string {
	return s.String()
}

func (s *ResetFileSystemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetFileSystemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetFileSystemResponse) GetBody() *ResetFileSystemResponseBody {
	return s.Body
}

func (s *ResetFileSystemResponse) SetHeaders(v map[string]*string) *ResetFileSystemResponse {
	s.Headers = v
	return s
}

func (s *ResetFileSystemResponse) SetStatusCode(v int32) *ResetFileSystemResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetFileSystemResponse) SetBody(v *ResetFileSystemResponseBody) *ResetFileSystemResponse {
	s.Body = v
	return s
}

func (s *ResetFileSystemResponse) Validate() error {
	return dara.Validate(s)
}
