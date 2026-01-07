// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRouterInterfaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRouterInterfaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRouterInterfaceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRouterInterfaceResponseBody) *DeleteRouterInterfaceResponse
	GetBody() *DeleteRouterInterfaceResponseBody
}

type DeleteRouterInterfaceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRouterInterfaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRouterInterfaceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRouterInterfaceResponse) GoString() string {
	return s.String()
}

func (s *DeleteRouterInterfaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRouterInterfaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRouterInterfaceResponse) GetBody() *DeleteRouterInterfaceResponseBody {
	return s.Body
}

func (s *DeleteRouterInterfaceResponse) SetHeaders(v map[string]*string) *DeleteRouterInterfaceResponse {
	s.Headers = v
	return s
}

func (s *DeleteRouterInterfaceResponse) SetStatusCode(v int32) *DeleteRouterInterfaceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRouterInterfaceResponse) SetBody(v *DeleteRouterInterfaceResponseBody) *DeleteRouterInterfaceResponse {
	s.Body = v
	return s
}

func (s *DeleteRouterInterfaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
