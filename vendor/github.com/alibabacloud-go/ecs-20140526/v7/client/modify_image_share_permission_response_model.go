// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyImageSharePermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyImageSharePermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyImageSharePermissionResponse
	GetStatusCode() *int32
	SetBody(v *ModifyImageSharePermissionResponseBody) *ModifyImageSharePermissionResponse
	GetBody() *ModifyImageSharePermissionResponseBody
}

type ModifyImageSharePermissionResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyImageSharePermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyImageSharePermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyImageSharePermissionResponse) GoString() string {
	return s.String()
}

func (s *ModifyImageSharePermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyImageSharePermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyImageSharePermissionResponse) GetBody() *ModifyImageSharePermissionResponseBody {
	return s.Body
}

func (s *ModifyImageSharePermissionResponse) SetHeaders(v map[string]*string) *ModifyImageSharePermissionResponse {
	s.Headers = v
	return s
}

func (s *ModifyImageSharePermissionResponse) SetStatusCode(v int32) *ModifyImageSharePermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyImageSharePermissionResponse) SetBody(v *ModifyImageSharePermissionResponseBody) *ModifyImageSharePermissionResponse {
	s.Body = v
	return s
}

func (s *ModifyImageSharePermissionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
