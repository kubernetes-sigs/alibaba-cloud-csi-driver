// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCommandResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCommandResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCommandResponse
	GetStatusCode() *int32
	SetBody(v *CreateCommandResponseBody) *CreateCommandResponse
	GetBody() *CreateCommandResponseBody
}

type CreateCommandResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCommandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCommandResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCommandResponse) GoString() string {
	return s.String()
}

func (s *CreateCommandResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCommandResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCommandResponse) GetBody() *CreateCommandResponseBody {
	return s.Body
}

func (s *CreateCommandResponse) SetHeaders(v map[string]*string) *CreateCommandResponse {
	s.Headers = v
	return s
}

func (s *CreateCommandResponse) SetStatusCode(v int32) *CreateCommandResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCommandResponse) SetBody(v *CreateCommandResponseBody) *CreateCommandResponse {
	s.Body = v
	return s
}

func (s *CreateCommandResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
