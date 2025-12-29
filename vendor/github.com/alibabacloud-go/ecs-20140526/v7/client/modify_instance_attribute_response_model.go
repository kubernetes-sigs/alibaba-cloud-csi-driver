// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyInstanceAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyInstanceAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyInstanceAttributeResponseBody) *ModifyInstanceAttributeResponse
	GetBody() *ModifyInstanceAttributeResponseBody
}

type ModifyInstanceAttributeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyInstanceAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyInstanceAttributeResponse) GetBody() *ModifyInstanceAttributeResponseBody {
	return s.Body
}

func (s *ModifyInstanceAttributeResponse) SetHeaders(v map[string]*string) *ModifyInstanceAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceAttributeResponse) SetStatusCode(v int32) *ModifyInstanceAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceAttributeResponse) SetBody(v *ModifyInstanceAttributeResponseBody) *ModifyInstanceAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyInstanceAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
