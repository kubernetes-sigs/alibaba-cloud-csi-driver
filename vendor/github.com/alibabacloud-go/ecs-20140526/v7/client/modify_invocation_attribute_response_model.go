// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInvocationAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyInvocationAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyInvocationAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyInvocationAttributeResponseBody) *ModifyInvocationAttributeResponse
	GetBody() *ModifyInvocationAttributeResponseBody
}

type ModifyInvocationAttributeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInvocationAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInvocationAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyInvocationAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyInvocationAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyInvocationAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyInvocationAttributeResponse) GetBody() *ModifyInvocationAttributeResponseBody {
	return s.Body
}

func (s *ModifyInvocationAttributeResponse) SetHeaders(v map[string]*string) *ModifyInvocationAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyInvocationAttributeResponse) SetStatusCode(v int32) *ModifyInvocationAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInvocationAttributeResponse) SetBody(v *ModifyInvocationAttributeResponseBody) *ModifyInvocationAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyInvocationAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
