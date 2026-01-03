// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDiskSpecResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDiskSpecResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDiskSpecResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDiskSpecResponseBody) *ModifyDiskSpecResponse
	GetBody() *ModifyDiskSpecResponseBody
}

type ModifyDiskSpecResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDiskSpecResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDiskSpecResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDiskSpecResponse) GoString() string {
	return s.String()
}

func (s *ModifyDiskSpecResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDiskSpecResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDiskSpecResponse) GetBody() *ModifyDiskSpecResponseBody {
	return s.Body
}

func (s *ModifyDiskSpecResponse) SetHeaders(v map[string]*string) *ModifyDiskSpecResponse {
	s.Headers = v
	return s
}

func (s *ModifyDiskSpecResponse) SetStatusCode(v int32) *ModifyDiskSpecResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDiskSpecResponse) SetBody(v *ModifyDiskSpecResponseBody) *ModifyDiskSpecResponse {
	s.Body = v
	return s
}

func (s *ModifyDiskSpecResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
