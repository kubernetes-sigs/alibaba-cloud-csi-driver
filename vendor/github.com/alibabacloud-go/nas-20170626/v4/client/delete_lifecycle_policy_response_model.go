// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLifecyclePolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLifecyclePolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLifecyclePolicyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLifecyclePolicyResponseBody) *DeleteLifecyclePolicyResponse
	GetBody() *DeleteLifecyclePolicyResponseBody
}

type DeleteLifecyclePolicyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLifecyclePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLifecyclePolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLifecyclePolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteLifecyclePolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLifecyclePolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLifecyclePolicyResponse) GetBody() *DeleteLifecyclePolicyResponseBody {
	return s.Body
}

func (s *DeleteLifecyclePolicyResponse) SetHeaders(v map[string]*string) *DeleteLifecyclePolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteLifecyclePolicyResponse) SetStatusCode(v int32) *DeleteLifecyclePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLifecyclePolicyResponse) SetBody(v *DeleteLifecyclePolicyResponseBody) *DeleteLifecyclePolicyResponse {
	s.Body = v
	return s
}

func (s *DeleteLifecyclePolicyResponse) Validate() error {
	return dara.Validate(s)
}
