// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNetworkInterfacePermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteNetworkInterfacePermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteNetworkInterfacePermissionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteNetworkInterfacePermissionResponseBody) *DeleteNetworkInterfacePermissionResponse
	GetBody() *DeleteNetworkInterfacePermissionResponseBody
}

type DeleteNetworkInterfacePermissionResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNetworkInterfacePermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNetworkInterfacePermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteNetworkInterfacePermissionResponse) GoString() string {
	return s.String()
}

func (s *DeleteNetworkInterfacePermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteNetworkInterfacePermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteNetworkInterfacePermissionResponse) GetBody() *DeleteNetworkInterfacePermissionResponseBody {
	return s.Body
}

func (s *DeleteNetworkInterfacePermissionResponse) SetHeaders(v map[string]*string) *DeleteNetworkInterfacePermissionResponse {
	s.Headers = v
	return s
}

func (s *DeleteNetworkInterfacePermissionResponse) SetStatusCode(v int32) *DeleteNetworkInterfacePermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNetworkInterfacePermissionResponse) SetBody(v *DeleteNetworkInterfacePermissionResponseBody) *DeleteNetworkInterfacePermissionResponse {
	s.Body = v
	return s
}

func (s *DeleteNetworkInterfacePermissionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
