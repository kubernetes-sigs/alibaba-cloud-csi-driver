// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVirtualBorderRouterAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyVirtualBorderRouterAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyVirtualBorderRouterAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyVirtualBorderRouterAttributeResponseBody) *ModifyVirtualBorderRouterAttributeResponse
	GetBody() *ModifyVirtualBorderRouterAttributeResponseBody
}

type ModifyVirtualBorderRouterAttributeResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyVirtualBorderRouterAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyVirtualBorderRouterAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyVirtualBorderRouterAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyVirtualBorderRouterAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyVirtualBorderRouterAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyVirtualBorderRouterAttributeResponse) GetBody() *ModifyVirtualBorderRouterAttributeResponseBody {
	return s.Body
}

func (s *ModifyVirtualBorderRouterAttributeResponse) SetHeaders(v map[string]*string) *ModifyVirtualBorderRouterAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyVirtualBorderRouterAttributeResponse) SetStatusCode(v int32) *ModifyVirtualBorderRouterAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyVirtualBorderRouterAttributeResponse) SetBody(v *ModifyVirtualBorderRouterAttributeResponseBody) *ModifyVirtualBorderRouterAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyVirtualBorderRouterAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
