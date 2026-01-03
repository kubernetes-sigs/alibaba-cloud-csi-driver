// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokeCommandResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InvokeCommandResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InvokeCommandResponse
	GetStatusCode() *int32
	SetBody(v *InvokeCommandResponseBody) *InvokeCommandResponse
	GetBody() *InvokeCommandResponseBody
}

type InvokeCommandResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InvokeCommandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InvokeCommandResponse) String() string {
	return dara.Prettify(s)
}

func (s InvokeCommandResponse) GoString() string {
	return s.String()
}

func (s *InvokeCommandResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InvokeCommandResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InvokeCommandResponse) GetBody() *InvokeCommandResponseBody {
	return s.Body
}

func (s *InvokeCommandResponse) SetHeaders(v map[string]*string) *InvokeCommandResponse {
	s.Headers = v
	return s
}

func (s *InvokeCommandResponse) SetStatusCode(v int32) *InvokeCommandResponse {
	s.StatusCode = &v
	return s
}

func (s *InvokeCommandResponse) SetBody(v *InvokeCommandResponseBody) *InvokeCommandResponse {
	s.Body = v
	return s
}

func (s *InvokeCommandResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
