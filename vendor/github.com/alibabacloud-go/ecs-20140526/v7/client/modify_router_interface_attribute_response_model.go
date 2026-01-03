// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRouterInterfaceAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyRouterInterfaceAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyRouterInterfaceAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyRouterInterfaceAttributeResponseBody) *ModifyRouterInterfaceAttributeResponse
	GetBody() *ModifyRouterInterfaceAttributeResponseBody
}

type ModifyRouterInterfaceAttributeResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyRouterInterfaceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyRouterInterfaceAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyRouterInterfaceAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyRouterInterfaceAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyRouterInterfaceAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyRouterInterfaceAttributeResponse) GetBody() *ModifyRouterInterfaceAttributeResponseBody {
	return s.Body
}

func (s *ModifyRouterInterfaceAttributeResponse) SetHeaders(v map[string]*string) *ModifyRouterInterfaceAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyRouterInterfaceAttributeResponse) SetStatusCode(v int32) *ModifyRouterInterfaceAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyRouterInterfaceAttributeResponse) SetBody(v *ModifyRouterInterfaceAttributeResponseBody) *ModifyRouterInterfaceAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyRouterInterfaceAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
