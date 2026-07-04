// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLifecyclePolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLifecyclePolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLifecyclePolicyResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLifecyclePolicyResponseBody) *UpdateLifecyclePolicyResponse
	GetBody() *UpdateLifecyclePolicyResponseBody
}

type UpdateLifecyclePolicyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLifecyclePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLifecyclePolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLifecyclePolicyResponse) GoString() string {
	return s.String()
}

func (s *UpdateLifecyclePolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLifecyclePolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLifecyclePolicyResponse) GetBody() *UpdateLifecyclePolicyResponseBody {
	return s.Body
}

func (s *UpdateLifecyclePolicyResponse) SetHeaders(v map[string]*string) *UpdateLifecyclePolicyResponse {
	s.Headers = v
	return s
}

func (s *UpdateLifecyclePolicyResponse) SetStatusCode(v int32) *UpdateLifecyclePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLifecyclePolicyResponse) SetBody(v *UpdateLifecyclePolicyResponseBody) *UpdateLifecyclePolicyResponse {
	s.Body = v
	return s
}

func (s *UpdateLifecyclePolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
