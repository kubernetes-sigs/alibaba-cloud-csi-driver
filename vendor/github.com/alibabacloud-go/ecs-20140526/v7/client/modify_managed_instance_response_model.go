// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyManagedInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyManagedInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyManagedInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ModifyManagedInstanceResponseBody) *ModifyManagedInstanceResponse
	GetBody() *ModifyManagedInstanceResponseBody
}

type ModifyManagedInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyManagedInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyManagedInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyManagedInstanceResponse) GoString() string {
	return s.String()
}

func (s *ModifyManagedInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyManagedInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyManagedInstanceResponse) GetBody() *ModifyManagedInstanceResponseBody {
	return s.Body
}

func (s *ModifyManagedInstanceResponse) SetHeaders(v map[string]*string) *ModifyManagedInstanceResponse {
	s.Headers = v
	return s
}

func (s *ModifyManagedInstanceResponse) SetStatusCode(v int32) *ModifyManagedInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyManagedInstanceResponse) SetBody(v *ModifyManagedInstanceResponseBody) *ModifyManagedInstanceResponse {
	s.Body = v
	return s
}

func (s *ModifyManagedInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
