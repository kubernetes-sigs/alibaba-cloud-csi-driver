// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachNetworkInterfaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachNetworkInterfaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachNetworkInterfaceResponse
	GetStatusCode() *int32
	SetBody(v *AttachNetworkInterfaceResponseBody) *AttachNetworkInterfaceResponse
	GetBody() *AttachNetworkInterfaceResponseBody
}

type AttachNetworkInterfaceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachNetworkInterfaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachNetworkInterfaceResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachNetworkInterfaceResponse) GoString() string {
	return s.String()
}

func (s *AttachNetworkInterfaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachNetworkInterfaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachNetworkInterfaceResponse) GetBody() *AttachNetworkInterfaceResponseBody {
	return s.Body
}

func (s *AttachNetworkInterfaceResponse) SetHeaders(v map[string]*string) *AttachNetworkInterfaceResponse {
	s.Headers = v
	return s
}

func (s *AttachNetworkInterfaceResponse) SetStatusCode(v int32) *AttachNetworkInterfaceResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachNetworkInterfaceResponse) SetBody(v *AttachNetworkInterfaceResponseBody) *AttachNetworkInterfaceResponse {
	s.Body = v
	return s
}

func (s *AttachNetworkInterfaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
