// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpcResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVpcResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVpcResponse
	GetStatusCode() *int32
	SetBody(v *CreateVpcResponseBody) *CreateVpcResponse
	GetBody() *CreateVpcResponseBody
}

type CreateVpcResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVpcResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVpcResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVpcResponse) GoString() string {
	return s.String()
}

func (s *CreateVpcResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVpcResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVpcResponse) GetBody() *CreateVpcResponseBody {
	return s.Body
}

func (s *CreateVpcResponse) SetHeaders(v map[string]*string) *CreateVpcResponse {
	s.Headers = v
	return s
}

func (s *CreateVpcResponse) SetStatusCode(v int32) *CreateVpcResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVpcResponse) SetBody(v *CreateVpcResponseBody) *CreateVpcResponse {
	s.Body = v
	return s
}

func (s *CreateVpcResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
