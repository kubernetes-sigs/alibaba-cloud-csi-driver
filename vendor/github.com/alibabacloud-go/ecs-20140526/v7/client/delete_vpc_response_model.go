// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVpcResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVpcResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVpcResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVpcResponseBody) *DeleteVpcResponse
	GetBody() *DeleteVpcResponseBody
}

type DeleteVpcResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVpcResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVpcResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVpcResponse) GoString() string {
	return s.String()
}

func (s *DeleteVpcResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVpcResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVpcResponse) GetBody() *DeleteVpcResponseBody {
	return s.Body
}

func (s *DeleteVpcResponse) SetHeaders(v map[string]*string) *DeleteVpcResponse {
	s.Headers = v
	return s
}

func (s *DeleteVpcResponse) SetStatusCode(v int32) *DeleteVpcResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVpcResponse) SetBody(v *DeleteVpcResponseBody) *DeleteVpcResponse {
	s.Body = v
	return s
}

func (s *DeleteVpcResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
