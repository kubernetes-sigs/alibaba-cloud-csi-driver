// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRouterInterfaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRouterInterfaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRouterInterfaceResponse
	GetStatusCode() *int32
	SetBody(v *CreateRouterInterfaceResponseBody) *CreateRouterInterfaceResponse
	GetBody() *CreateRouterInterfaceResponseBody
}

type CreateRouterInterfaceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRouterInterfaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRouterInterfaceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRouterInterfaceResponse) GoString() string {
	return s.String()
}

func (s *CreateRouterInterfaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRouterInterfaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRouterInterfaceResponse) GetBody() *CreateRouterInterfaceResponseBody {
	return s.Body
}

func (s *CreateRouterInterfaceResponse) SetHeaders(v map[string]*string) *CreateRouterInterfaceResponse {
	s.Headers = v
	return s
}

func (s *CreateRouterInterfaceResponse) SetStatusCode(v int32) *CreateRouterInterfaceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRouterInterfaceResponse) SetBody(v *CreateRouterInterfaceResponseBody) *CreateRouterInterfaceResponse {
	s.Body = v
	return s
}

func (s *CreateRouterInterfaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
