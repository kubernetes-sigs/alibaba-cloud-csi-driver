// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachVscFromFilesystemsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachVscFromFilesystemsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachVscFromFilesystemsResponse
	GetStatusCode() *int32
	SetBody(v *DetachVscFromFilesystemsResponseBody) *DetachVscFromFilesystemsResponse
	GetBody() *DetachVscFromFilesystemsResponseBody
}

type DetachVscFromFilesystemsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachVscFromFilesystemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachVscFromFilesystemsResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachVscFromFilesystemsResponse) GoString() string {
	return s.String()
}

func (s *DetachVscFromFilesystemsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachVscFromFilesystemsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachVscFromFilesystemsResponse) GetBody() *DetachVscFromFilesystemsResponseBody {
	return s.Body
}

func (s *DetachVscFromFilesystemsResponse) SetHeaders(v map[string]*string) *DetachVscFromFilesystemsResponse {
	s.Headers = v
	return s
}

func (s *DetachVscFromFilesystemsResponse) SetStatusCode(v int32) *DetachVscFromFilesystemsResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachVscFromFilesystemsResponse) SetBody(v *DetachVscFromFilesystemsResponseBody) *DetachVscFromFilesystemsResponse {
	s.Body = v
	return s
}

func (s *DetachVscFromFilesystemsResponse) Validate() error {
	return dara.Validate(s)
}
