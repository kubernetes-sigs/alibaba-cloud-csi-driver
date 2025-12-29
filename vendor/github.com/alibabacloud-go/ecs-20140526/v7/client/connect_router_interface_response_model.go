// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConnectRouterInterfaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConnectRouterInterfaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConnectRouterInterfaceResponse
	GetStatusCode() *int32
	SetBody(v *ConnectRouterInterfaceResponseBody) *ConnectRouterInterfaceResponse
	GetBody() *ConnectRouterInterfaceResponseBody
}

type ConnectRouterInterfaceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConnectRouterInterfaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConnectRouterInterfaceResponse) String() string {
	return dara.Prettify(s)
}

func (s ConnectRouterInterfaceResponse) GoString() string {
	return s.String()
}

func (s *ConnectRouterInterfaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConnectRouterInterfaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConnectRouterInterfaceResponse) GetBody() *ConnectRouterInterfaceResponseBody {
	return s.Body
}

func (s *ConnectRouterInterfaceResponse) SetHeaders(v map[string]*string) *ConnectRouterInterfaceResponse {
	s.Headers = v
	return s
}

func (s *ConnectRouterInterfaceResponse) SetStatusCode(v int32) *ConnectRouterInterfaceResponse {
	s.StatusCode = &v
	return s
}

func (s *ConnectRouterInterfaceResponse) SetBody(v *ConnectRouterInterfaceResponseBody) *ConnectRouterInterfaceResponse {
	s.Body = v
	return s
}

func (s *ConnectRouterInterfaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
