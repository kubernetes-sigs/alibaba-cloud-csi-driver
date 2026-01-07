// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNatGatewayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteNatGatewayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteNatGatewayResponse
	GetStatusCode() *int32
	SetBody(v *DeleteNatGatewayResponseBody) *DeleteNatGatewayResponse
	GetBody() *DeleteNatGatewayResponseBody
}

type DeleteNatGatewayResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNatGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNatGatewayResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteNatGatewayResponse) GoString() string {
	return s.String()
}

func (s *DeleteNatGatewayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteNatGatewayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteNatGatewayResponse) GetBody() *DeleteNatGatewayResponseBody {
	return s.Body
}

func (s *DeleteNatGatewayResponse) SetHeaders(v map[string]*string) *DeleteNatGatewayResponse {
	s.Headers = v
	return s
}

func (s *DeleteNatGatewayResponse) SetStatusCode(v int32) *DeleteNatGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNatGatewayResponse) SetBody(v *DeleteNatGatewayResponseBody) *DeleteNatGatewayResponse {
	s.Body = v
	return s
}

func (s *DeleteNatGatewayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
