// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeactivateRouterInterfaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeactivateRouterInterfaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeactivateRouterInterfaceResponse
	GetStatusCode() *int32
	SetBody(v *DeactivateRouterInterfaceResponseBody) *DeactivateRouterInterfaceResponse
	GetBody() *DeactivateRouterInterfaceResponseBody
}

type DeactivateRouterInterfaceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeactivateRouterInterfaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeactivateRouterInterfaceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeactivateRouterInterfaceResponse) GoString() string {
	return s.String()
}

func (s *DeactivateRouterInterfaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeactivateRouterInterfaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeactivateRouterInterfaceResponse) GetBody() *DeactivateRouterInterfaceResponseBody {
	return s.Body
}

func (s *DeactivateRouterInterfaceResponse) SetHeaders(v map[string]*string) *DeactivateRouterInterfaceResponse {
	s.Headers = v
	return s
}

func (s *DeactivateRouterInterfaceResponse) SetStatusCode(v int32) *DeactivateRouterInterfaceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeactivateRouterInterfaceResponse) SetBody(v *DeactivateRouterInterfaceResponseBody) *DeactivateRouterInterfaceResponse {
	s.Body = v
	return s
}

func (s *DeactivateRouterInterfaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
