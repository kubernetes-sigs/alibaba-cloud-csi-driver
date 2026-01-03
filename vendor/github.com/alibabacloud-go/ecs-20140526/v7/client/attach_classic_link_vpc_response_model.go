// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachClassicLinkVpcResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachClassicLinkVpcResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachClassicLinkVpcResponse
	GetStatusCode() *int32
	SetBody(v *AttachClassicLinkVpcResponseBody) *AttachClassicLinkVpcResponse
	GetBody() *AttachClassicLinkVpcResponseBody
}

type AttachClassicLinkVpcResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachClassicLinkVpcResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachClassicLinkVpcResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachClassicLinkVpcResponse) GoString() string {
	return s.String()
}

func (s *AttachClassicLinkVpcResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachClassicLinkVpcResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachClassicLinkVpcResponse) GetBody() *AttachClassicLinkVpcResponseBody {
	return s.Body
}

func (s *AttachClassicLinkVpcResponse) SetHeaders(v map[string]*string) *AttachClassicLinkVpcResponse {
	s.Headers = v
	return s
}

func (s *AttachClassicLinkVpcResponse) SetStatusCode(v int32) *AttachClassicLinkVpcResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachClassicLinkVpcResponse) SetBody(v *AttachClassicLinkVpcResponseBody) *AttachClassicLinkVpcResponse {
	s.Body = v
	return s
}

func (s *AttachClassicLinkVpcResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
