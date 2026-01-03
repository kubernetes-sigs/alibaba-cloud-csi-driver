// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartTerminalSessionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartTerminalSessionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartTerminalSessionResponse
	GetStatusCode() *int32
	SetBody(v *StartTerminalSessionResponseBody) *StartTerminalSessionResponse
	GetBody() *StartTerminalSessionResponseBody
}

type StartTerminalSessionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartTerminalSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartTerminalSessionResponse) String() string {
	return dara.Prettify(s)
}

func (s StartTerminalSessionResponse) GoString() string {
	return s.String()
}

func (s *StartTerminalSessionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartTerminalSessionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartTerminalSessionResponse) GetBody() *StartTerminalSessionResponseBody {
	return s.Body
}

func (s *StartTerminalSessionResponse) SetHeaders(v map[string]*string) *StartTerminalSessionResponse {
	s.Headers = v
	return s
}

func (s *StartTerminalSessionResponse) SetStatusCode(v int32) *StartTerminalSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *StartTerminalSessionResponse) SetBody(v *StartTerminalSessionResponseBody) *StartTerminalSessionResponse {
	s.Body = v
	return s
}

func (s *StartTerminalSessionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
