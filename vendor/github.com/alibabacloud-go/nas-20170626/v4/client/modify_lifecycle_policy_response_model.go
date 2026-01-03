// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLifecyclePolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyLifecyclePolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyLifecyclePolicyResponse
	GetStatusCode() *int32
	SetBody(v *ModifyLifecyclePolicyResponseBody) *ModifyLifecyclePolicyResponse
	GetBody() *ModifyLifecyclePolicyResponseBody
}

type ModifyLifecyclePolicyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyLifecyclePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyLifecyclePolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyLifecyclePolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyLifecyclePolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyLifecyclePolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyLifecyclePolicyResponse) GetBody() *ModifyLifecyclePolicyResponseBody {
	return s.Body
}

func (s *ModifyLifecyclePolicyResponse) SetHeaders(v map[string]*string) *ModifyLifecyclePolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyLifecyclePolicyResponse) SetStatusCode(v int32) *ModifyLifecyclePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyLifecyclePolicyResponse) SetBody(v *ModifyLifecyclePolicyResponseBody) *ModifyLifecyclePolicyResponse {
	s.Body = v
	return s
}

func (s *ModifyLifecyclePolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
