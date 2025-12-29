// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVRouterAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyVRouterAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyVRouterAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyVRouterAttributeResponseBody) *ModifyVRouterAttributeResponse
	GetBody() *ModifyVRouterAttributeResponseBody
}

type ModifyVRouterAttributeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyVRouterAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyVRouterAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyVRouterAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyVRouterAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyVRouterAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyVRouterAttributeResponse) GetBody() *ModifyVRouterAttributeResponseBody {
	return s.Body
}

func (s *ModifyVRouterAttributeResponse) SetHeaders(v map[string]*string) *ModifyVRouterAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyVRouterAttributeResponse) SetStatusCode(v int32) *ModifyVRouterAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyVRouterAttributeResponse) SetBody(v *ModifyVRouterAttributeResponseBody) *ModifyVRouterAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyVRouterAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
