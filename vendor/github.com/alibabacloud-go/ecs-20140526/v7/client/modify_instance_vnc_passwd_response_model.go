// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceVncPasswdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyInstanceVncPasswdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyInstanceVncPasswdResponse
	GetStatusCode() *int32
	SetBody(v *ModifyInstanceVncPasswdResponseBody) *ModifyInstanceVncPasswdResponse
	GetBody() *ModifyInstanceVncPasswdResponseBody
}

type ModifyInstanceVncPasswdResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceVncPasswdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceVncPasswdResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceVncPasswdResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceVncPasswdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyInstanceVncPasswdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyInstanceVncPasswdResponse) GetBody() *ModifyInstanceVncPasswdResponseBody {
	return s.Body
}

func (s *ModifyInstanceVncPasswdResponse) SetHeaders(v map[string]*string) *ModifyInstanceVncPasswdResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceVncPasswdResponse) SetStatusCode(v int32) *ModifyInstanceVncPasswdResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceVncPasswdResponse) SetBody(v *ModifyInstanceVncPasswdResponseBody) *ModifyInstanceVncPasswdResponse {
	s.Body = v
	return s
}

func (s *ModifyInstanceVncPasswdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
