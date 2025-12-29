// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetworkInterfacePermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateNetworkInterfacePermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateNetworkInterfacePermissionResponse
	GetStatusCode() *int32
	SetBody(v *CreateNetworkInterfacePermissionResponseBody) *CreateNetworkInterfacePermissionResponse
	GetBody() *CreateNetworkInterfacePermissionResponseBody
}

type CreateNetworkInterfacePermissionResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNetworkInterfacePermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNetworkInterfacePermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkInterfacePermissionResponse) GoString() string {
	return s.String()
}

func (s *CreateNetworkInterfacePermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateNetworkInterfacePermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateNetworkInterfacePermissionResponse) GetBody() *CreateNetworkInterfacePermissionResponseBody {
	return s.Body
}

func (s *CreateNetworkInterfacePermissionResponse) SetHeaders(v map[string]*string) *CreateNetworkInterfacePermissionResponse {
	s.Headers = v
	return s
}

func (s *CreateNetworkInterfacePermissionResponse) SetStatusCode(v int32) *CreateNetworkInterfacePermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNetworkInterfacePermissionResponse) SetBody(v *CreateNetworkInterfacePermissionResponseBody) *CreateNetworkInterfacePermissionResponse {
	s.Body = v
	return s
}

func (s *CreateNetworkInterfacePermissionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
