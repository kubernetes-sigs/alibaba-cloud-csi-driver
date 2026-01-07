// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNetworkInterfaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteNetworkInterfaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteNetworkInterfaceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteNetworkInterfaceResponseBody) *DeleteNetworkInterfaceResponse
	GetBody() *DeleteNetworkInterfaceResponseBody
}

type DeleteNetworkInterfaceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNetworkInterfaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNetworkInterfaceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteNetworkInterfaceResponse) GoString() string {
	return s.String()
}

func (s *DeleteNetworkInterfaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteNetworkInterfaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteNetworkInterfaceResponse) GetBody() *DeleteNetworkInterfaceResponseBody {
	return s.Body
}

func (s *DeleteNetworkInterfaceResponse) SetHeaders(v map[string]*string) *DeleteNetworkInterfaceResponse {
	s.Headers = v
	return s
}

func (s *DeleteNetworkInterfaceResponse) SetStatusCode(v int32) *DeleteNetworkInterfaceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNetworkInterfaceResponse) SetBody(v *DeleteNetworkInterfaceResponseBody) *DeleteNetworkInterfaceResponse {
	s.Body = v
	return s
}

func (s *DeleteNetworkInterfaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
