// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNetworkInterfaceAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyNetworkInterfaceAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyNetworkInterfaceAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyNetworkInterfaceAttributeResponseBody) *ModifyNetworkInterfaceAttributeResponse
	GetBody() *ModifyNetworkInterfaceAttributeResponseBody
}

type ModifyNetworkInterfaceAttributeResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyNetworkInterfaceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyNetworkInterfaceAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyNetworkInterfaceAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyNetworkInterfaceAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyNetworkInterfaceAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyNetworkInterfaceAttributeResponse) GetBody() *ModifyNetworkInterfaceAttributeResponseBody {
	return s.Body
}

func (s *ModifyNetworkInterfaceAttributeResponse) SetHeaders(v map[string]*string) *ModifyNetworkInterfaceAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyNetworkInterfaceAttributeResponse) SetStatusCode(v int32) *ModifyNetworkInterfaceAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyNetworkInterfaceAttributeResponse) SetBody(v *ModifyNetworkInterfaceAttributeResponseBody) *ModifyNetworkInterfaceAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyNetworkInterfaceAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
