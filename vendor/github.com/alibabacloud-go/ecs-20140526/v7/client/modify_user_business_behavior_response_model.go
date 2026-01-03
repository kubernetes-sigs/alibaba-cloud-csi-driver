// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUserBusinessBehaviorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyUserBusinessBehaviorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyUserBusinessBehaviorResponse
	GetStatusCode() *int32
	SetBody(v *ModifyUserBusinessBehaviorResponseBody) *ModifyUserBusinessBehaviorResponse
	GetBody() *ModifyUserBusinessBehaviorResponseBody
}

type ModifyUserBusinessBehaviorResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyUserBusinessBehaviorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyUserBusinessBehaviorResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserBusinessBehaviorResponse) GoString() string {
	return s.String()
}

func (s *ModifyUserBusinessBehaviorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyUserBusinessBehaviorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyUserBusinessBehaviorResponse) GetBody() *ModifyUserBusinessBehaviorResponseBody {
	return s.Body
}

func (s *ModifyUserBusinessBehaviorResponse) SetHeaders(v map[string]*string) *ModifyUserBusinessBehaviorResponse {
	s.Headers = v
	return s
}

func (s *ModifyUserBusinessBehaviorResponse) SetStatusCode(v int32) *ModifyUserBusinessBehaviorResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyUserBusinessBehaviorResponse) SetBody(v *ModifyUserBusinessBehaviorResponseBody) *ModifyUserBusinessBehaviorResponse {
	s.Body = v
	return s
}

func (s *ModifyUserBusinessBehaviorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
