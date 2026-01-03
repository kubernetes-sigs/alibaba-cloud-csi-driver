// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyImageAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyImageAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyImageAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyImageAttributeResponseBody) *ModifyImageAttributeResponse
	GetBody() *ModifyImageAttributeResponseBody
}

type ModifyImageAttributeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyImageAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyImageAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyImageAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyImageAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyImageAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyImageAttributeResponse) GetBody() *ModifyImageAttributeResponseBody {
	return s.Body
}

func (s *ModifyImageAttributeResponse) SetHeaders(v map[string]*string) *ModifyImageAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyImageAttributeResponse) SetStatusCode(v int32) *ModifyImageAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyImageAttributeResponse) SetBody(v *ModifyImageAttributeResponseBody) *ModifyImageAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyImageAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
