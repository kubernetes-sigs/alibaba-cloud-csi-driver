// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActivateRouterInterfaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ActivateRouterInterfaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ActivateRouterInterfaceResponse
	GetStatusCode() *int32
	SetBody(v *ActivateRouterInterfaceResponseBody) *ActivateRouterInterfaceResponse
	GetBody() *ActivateRouterInterfaceResponseBody
}

type ActivateRouterInterfaceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ActivateRouterInterfaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ActivateRouterInterfaceResponse) String() string {
	return dara.Prettify(s)
}

func (s ActivateRouterInterfaceResponse) GoString() string {
	return s.String()
}

func (s *ActivateRouterInterfaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ActivateRouterInterfaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ActivateRouterInterfaceResponse) GetBody() *ActivateRouterInterfaceResponseBody {
	return s.Body
}

func (s *ActivateRouterInterfaceResponse) SetHeaders(v map[string]*string) *ActivateRouterInterfaceResponse {
	s.Headers = v
	return s
}

func (s *ActivateRouterInterfaceResponse) SetStatusCode(v int32) *ActivateRouterInterfaceResponse {
	s.StatusCode = &v
	return s
}

func (s *ActivateRouterInterfaceResponse) SetBody(v *ActivateRouterInterfaceResponseBody) *ActivateRouterInterfaceResponse {
	s.Body = v
	return s
}

func (s *ActivateRouterInterfaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
