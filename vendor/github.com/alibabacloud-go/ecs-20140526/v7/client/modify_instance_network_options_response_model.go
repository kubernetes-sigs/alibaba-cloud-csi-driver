// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceNetworkOptionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyInstanceNetworkOptionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyInstanceNetworkOptionsResponse
	GetStatusCode() *int32
	SetBody(v *ModifyInstanceNetworkOptionsResponseBody) *ModifyInstanceNetworkOptionsResponse
	GetBody() *ModifyInstanceNetworkOptionsResponseBody
}

type ModifyInstanceNetworkOptionsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceNetworkOptionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceNetworkOptionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceNetworkOptionsResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceNetworkOptionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyInstanceNetworkOptionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyInstanceNetworkOptionsResponse) GetBody() *ModifyInstanceNetworkOptionsResponseBody {
	return s.Body
}

func (s *ModifyInstanceNetworkOptionsResponse) SetHeaders(v map[string]*string) *ModifyInstanceNetworkOptionsResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceNetworkOptionsResponse) SetStatusCode(v int32) *ModifyInstanceNetworkOptionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceNetworkOptionsResponse) SetBody(v *ModifyInstanceNetworkOptionsResponseBody) *ModifyInstanceNetworkOptionsResponse {
	s.Body = v
	return s
}

func (s *ModifyInstanceNetworkOptionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
