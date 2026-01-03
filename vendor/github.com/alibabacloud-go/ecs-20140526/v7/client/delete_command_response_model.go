// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCommandResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCommandResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCommandResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCommandResponseBody) *DeleteCommandResponse
	GetBody() *DeleteCommandResponseBody
}

type DeleteCommandResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCommandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCommandResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCommandResponse) GoString() string {
	return s.String()
}

func (s *DeleteCommandResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCommandResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCommandResponse) GetBody() *DeleteCommandResponseBody {
	return s.Body
}

func (s *DeleteCommandResponse) SetHeaders(v map[string]*string) *DeleteCommandResponse {
	s.Headers = v
	return s
}

func (s *DeleteCommandResponse) SetStatusCode(v int32) *DeleteCommandResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCommandResponse) SetBody(v *DeleteCommandResponseBody) *DeleteCommandResponse {
	s.Body = v
	return s
}

func (s *DeleteCommandResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
