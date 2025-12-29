// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCommandResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCommandResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCommandResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCommandResponseBody) *ModifyCommandResponse
	GetBody() *ModifyCommandResponseBody
}

type ModifyCommandResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCommandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCommandResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCommandResponse) GoString() string {
	return s.String()
}

func (s *ModifyCommandResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCommandResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCommandResponse) GetBody() *ModifyCommandResponseBody {
	return s.Body
}

func (s *ModifyCommandResponse) SetHeaders(v map[string]*string) *ModifyCommandResponse {
	s.Headers = v
	return s
}

func (s *ModifyCommandResponse) SetStatusCode(v int32) *ModifyCommandResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCommandResponse) SetBody(v *ModifyCommandResponseBody) *ModifyCommandResponse {
	s.Body = v
	return s
}

func (s *ModifyCommandResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
