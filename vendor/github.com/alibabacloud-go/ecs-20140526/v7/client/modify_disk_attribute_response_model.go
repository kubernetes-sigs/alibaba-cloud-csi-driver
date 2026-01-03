// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDiskAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDiskAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDiskAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDiskAttributeResponseBody) *ModifyDiskAttributeResponse
	GetBody() *ModifyDiskAttributeResponseBody
}

type ModifyDiskAttributeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDiskAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDiskAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDiskAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyDiskAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDiskAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDiskAttributeResponse) GetBody() *ModifyDiskAttributeResponseBody {
	return s.Body
}

func (s *ModifyDiskAttributeResponse) SetHeaders(v map[string]*string) *ModifyDiskAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyDiskAttributeResponse) SetStatusCode(v int32) *ModifyDiskAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDiskAttributeResponse) SetBody(v *ModifyDiskAttributeResponseBody) *ModifyDiskAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyDiskAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
