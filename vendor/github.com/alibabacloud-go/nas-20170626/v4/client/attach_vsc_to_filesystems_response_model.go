// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachVscToFilesystemsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachVscToFilesystemsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachVscToFilesystemsResponse
	GetStatusCode() *int32
	SetBody(v *AttachVscToFilesystemsResponseBody) *AttachVscToFilesystemsResponse
	GetBody() *AttachVscToFilesystemsResponseBody
}

type AttachVscToFilesystemsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachVscToFilesystemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachVscToFilesystemsResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachVscToFilesystemsResponse) GoString() string {
	return s.String()
}

func (s *AttachVscToFilesystemsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachVscToFilesystemsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachVscToFilesystemsResponse) GetBody() *AttachVscToFilesystemsResponseBody {
	return s.Body
}

func (s *AttachVscToFilesystemsResponse) SetHeaders(v map[string]*string) *AttachVscToFilesystemsResponse {
	s.Headers = v
	return s
}

func (s *AttachVscToFilesystemsResponse) SetStatusCode(v int32) *AttachVscToFilesystemsResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachVscToFilesystemsResponse) SetBody(v *AttachVscToFilesystemsResponseBody) *AttachVscToFilesystemsResponse {
	s.Body = v
	return s
}

func (s *AttachVscToFilesystemsResponse) Validate() error {
	return dara.Validate(s)
}
