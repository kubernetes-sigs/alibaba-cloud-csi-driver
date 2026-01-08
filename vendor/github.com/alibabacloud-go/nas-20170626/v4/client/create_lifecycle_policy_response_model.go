// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLifecyclePolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLifecyclePolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLifecyclePolicyResponse
	GetStatusCode() *int32
	SetBody(v *CreateLifecyclePolicyResponseBody) *CreateLifecyclePolicyResponse
	GetBody() *CreateLifecyclePolicyResponseBody
}

type CreateLifecyclePolicyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLifecyclePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLifecyclePolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLifecyclePolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateLifecyclePolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLifecyclePolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLifecyclePolicyResponse) GetBody() *CreateLifecyclePolicyResponseBody {
	return s.Body
}

func (s *CreateLifecyclePolicyResponse) SetHeaders(v map[string]*string) *CreateLifecyclePolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateLifecyclePolicyResponse) SetStatusCode(v int32) *CreateLifecyclePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLifecyclePolicyResponse) SetBody(v *CreateLifecyclePolicyResponseBody) *CreateLifecyclePolicyResponse {
	s.Body = v
	return s
}

func (s *CreateLifecyclePolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
