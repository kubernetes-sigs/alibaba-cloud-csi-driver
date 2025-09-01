// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunCommandResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunCommandResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunCommandResponse
	GetStatusCode() *int32
	SetBody(v *RunCommandResponseBody) *RunCommandResponse
	GetBody() *RunCommandResponseBody
}

type RunCommandResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunCommandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunCommandResponse) String() string {
	return dara.Prettify(s)
}

func (s RunCommandResponse) GoString() string {
	return s.String()
}

func (s *RunCommandResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunCommandResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunCommandResponse) GetBody() *RunCommandResponseBody {
	return s.Body
}

func (s *RunCommandResponse) SetHeaders(v map[string]*string) *RunCommandResponse {
	s.Headers = v
	return s
}

func (s *RunCommandResponse) SetStatusCode(v int32) *RunCommandResponse {
	s.StatusCode = &v
	return s
}

func (s *RunCommandResponse) SetBody(v *RunCommandResponseBody) *RunCommandResponse {
	s.Body = v
	return s
}

func (s *RunCommandResponse) Validate() error {
	return dara.Validate(s)
}
