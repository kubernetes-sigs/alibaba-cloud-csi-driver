// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyImageShareGroupPermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyImageShareGroupPermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyImageShareGroupPermissionResponse
	GetStatusCode() *int32
	SetBody(v *ModifyImageShareGroupPermissionResponseBody) *ModifyImageShareGroupPermissionResponse
	GetBody() *ModifyImageShareGroupPermissionResponseBody
}

type ModifyImageShareGroupPermissionResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyImageShareGroupPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyImageShareGroupPermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyImageShareGroupPermissionResponse) GoString() string {
	return s.String()
}

func (s *ModifyImageShareGroupPermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyImageShareGroupPermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyImageShareGroupPermissionResponse) GetBody() *ModifyImageShareGroupPermissionResponseBody {
	return s.Body
}

func (s *ModifyImageShareGroupPermissionResponse) SetHeaders(v map[string]*string) *ModifyImageShareGroupPermissionResponse {
	s.Headers = v
	return s
}

func (s *ModifyImageShareGroupPermissionResponse) SetStatusCode(v int32) *ModifyImageShareGroupPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyImageShareGroupPermissionResponse) SetBody(v *ModifyImageShareGroupPermissionResponseBody) *ModifyImageShareGroupPermissionResponse {
	s.Body = v
	return s
}

func (s *ModifyImageShareGroupPermissionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
