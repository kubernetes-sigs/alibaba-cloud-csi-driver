// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFileSystemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyFileSystemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyFileSystemResponse
	GetStatusCode() *int32
	SetBody(v *ModifyFileSystemResponseBody) *ModifyFileSystemResponse
	GetBody() *ModifyFileSystemResponseBody
}

type ModifyFileSystemResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyFileSystemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyFileSystemResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyFileSystemResponse) GoString() string {
	return s.String()
}

func (s *ModifyFileSystemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyFileSystemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyFileSystemResponse) GetBody() *ModifyFileSystemResponseBody {
	return s.Body
}

func (s *ModifyFileSystemResponse) SetHeaders(v map[string]*string) *ModifyFileSystemResponse {
	s.Headers = v
	return s
}

func (s *ModifyFileSystemResponse) SetStatusCode(v int32) *ModifyFileSystemResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyFileSystemResponse) SetBody(v *ModifyFileSystemResponseBody) *ModifyFileSystemResponse {
	s.Body = v
	return s
}

func (s *ModifyFileSystemResponse) Validate() error {
	return dara.Validate(s)
}
