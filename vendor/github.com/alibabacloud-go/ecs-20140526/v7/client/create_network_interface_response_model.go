// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetworkInterfaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateNetworkInterfaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateNetworkInterfaceResponse
	GetStatusCode() *int32
	SetBody(v *CreateNetworkInterfaceResponseBody) *CreateNetworkInterfaceResponse
	GetBody() *CreateNetworkInterfaceResponseBody
}

type CreateNetworkInterfaceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNetworkInterfaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNetworkInterfaceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkInterfaceResponse) GoString() string {
	return s.String()
}

func (s *CreateNetworkInterfaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateNetworkInterfaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateNetworkInterfaceResponse) GetBody() *CreateNetworkInterfaceResponseBody {
	return s.Body
}

func (s *CreateNetworkInterfaceResponse) SetHeaders(v map[string]*string) *CreateNetworkInterfaceResponse {
	s.Headers = v
	return s
}

func (s *CreateNetworkInterfaceResponse) SetStatusCode(v int32) *CreateNetworkInterfaceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNetworkInterfaceResponse) SetBody(v *CreateNetworkInterfaceResponseBody) *CreateNetworkInterfaceResponse {
	s.Body = v
	return s
}

func (s *CreateNetworkInterfaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
