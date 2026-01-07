// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachNetworkInterfaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachNetworkInterfaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachNetworkInterfaceResponse
	GetStatusCode() *int32
	SetBody(v *DetachNetworkInterfaceResponseBody) *DetachNetworkInterfaceResponse
	GetBody() *DetachNetworkInterfaceResponseBody
}

type DetachNetworkInterfaceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachNetworkInterfaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachNetworkInterfaceResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachNetworkInterfaceResponse) GoString() string {
	return s.String()
}

func (s *DetachNetworkInterfaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachNetworkInterfaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachNetworkInterfaceResponse) GetBody() *DetachNetworkInterfaceResponseBody {
	return s.Body
}

func (s *DetachNetworkInterfaceResponse) SetHeaders(v map[string]*string) *DetachNetworkInterfaceResponse {
	s.Headers = v
	return s
}

func (s *DetachNetworkInterfaceResponse) SetStatusCode(v int32) *DetachNetworkInterfaceResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachNetworkInterfaceResponse) SetBody(v *DetachNetworkInterfaceResponseBody) *DetachNetworkInterfaceResponse {
	s.Body = v
	return s
}

func (s *DetachNetworkInterfaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
