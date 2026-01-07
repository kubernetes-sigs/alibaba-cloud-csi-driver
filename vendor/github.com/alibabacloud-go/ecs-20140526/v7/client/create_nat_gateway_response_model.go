// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNatGatewayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateNatGatewayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateNatGatewayResponse
	GetStatusCode() *int32
	SetBody(v *CreateNatGatewayResponseBody) *CreateNatGatewayResponse
	GetBody() *CreateNatGatewayResponseBody
}

type CreateNatGatewayResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNatGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNatGatewayResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateNatGatewayResponse) GoString() string {
	return s.String()
}

func (s *CreateNatGatewayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateNatGatewayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateNatGatewayResponse) GetBody() *CreateNatGatewayResponseBody {
	return s.Body
}

func (s *CreateNatGatewayResponse) SetHeaders(v map[string]*string) *CreateNatGatewayResponse {
	s.Headers = v
	return s
}

func (s *CreateNatGatewayResponse) SetStatusCode(v int32) *CreateNatGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNatGatewayResponse) SetBody(v *CreateNatGatewayResponseBody) *CreateNatGatewayResponse {
	s.Body = v
	return s
}

func (s *CreateNatGatewayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
