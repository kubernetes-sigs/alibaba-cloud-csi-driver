// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableNetworkInterfaceQoSResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableNetworkInterfaceQoSResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableNetworkInterfaceQoSResponse
	GetStatusCode() *int32
	SetBody(v *DisableNetworkInterfaceQoSResponseBody) *DisableNetworkInterfaceQoSResponse
	GetBody() *DisableNetworkInterfaceQoSResponseBody
}

type DisableNetworkInterfaceQoSResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableNetworkInterfaceQoSResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableNetworkInterfaceQoSResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableNetworkInterfaceQoSResponse) GoString() string {
	return s.String()
}

func (s *DisableNetworkInterfaceQoSResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableNetworkInterfaceQoSResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableNetworkInterfaceQoSResponse) GetBody() *DisableNetworkInterfaceQoSResponseBody {
	return s.Body
}

func (s *DisableNetworkInterfaceQoSResponse) SetHeaders(v map[string]*string) *DisableNetworkInterfaceQoSResponse {
	s.Headers = v
	return s
}

func (s *DisableNetworkInterfaceQoSResponse) SetStatusCode(v int32) *DisableNetworkInterfaceQoSResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableNetworkInterfaceQoSResponse) SetBody(v *DisableNetworkInterfaceQoSResponseBody) *DisableNetworkInterfaceQoSResponse {
	s.Body = v
	return s
}

func (s *DisableNetworkInterfaceQoSResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
