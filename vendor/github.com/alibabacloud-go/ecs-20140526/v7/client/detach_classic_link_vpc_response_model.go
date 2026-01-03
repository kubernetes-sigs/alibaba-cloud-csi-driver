// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachClassicLinkVpcResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachClassicLinkVpcResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachClassicLinkVpcResponse
	GetStatusCode() *int32
	SetBody(v *DetachClassicLinkVpcResponseBody) *DetachClassicLinkVpcResponse
	GetBody() *DetachClassicLinkVpcResponseBody
}

type DetachClassicLinkVpcResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachClassicLinkVpcResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachClassicLinkVpcResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachClassicLinkVpcResponse) GoString() string {
	return s.String()
}

func (s *DetachClassicLinkVpcResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachClassicLinkVpcResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachClassicLinkVpcResponse) GetBody() *DetachClassicLinkVpcResponseBody {
	return s.Body
}

func (s *DetachClassicLinkVpcResponse) SetHeaders(v map[string]*string) *DetachClassicLinkVpcResponse {
	s.Headers = v
	return s
}

func (s *DetachClassicLinkVpcResponse) SetStatusCode(v int32) *DetachClassicLinkVpcResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachClassicLinkVpcResponse) SetBody(v *DetachClassicLinkVpcResponseBody) *DetachClassicLinkVpcResponse {
	s.Body = v
	return s
}

func (s *DetachClassicLinkVpcResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
