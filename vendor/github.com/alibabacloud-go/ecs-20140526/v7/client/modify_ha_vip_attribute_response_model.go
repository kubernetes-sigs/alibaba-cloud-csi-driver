// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHaVipAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyHaVipAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyHaVipAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyHaVipAttributeResponseBody) *ModifyHaVipAttributeResponse
	GetBody() *ModifyHaVipAttributeResponseBody
}

type ModifyHaVipAttributeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyHaVipAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyHaVipAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyHaVipAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyHaVipAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyHaVipAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyHaVipAttributeResponse) GetBody() *ModifyHaVipAttributeResponseBody {
	return s.Body
}

func (s *ModifyHaVipAttributeResponse) SetHeaders(v map[string]*string) *ModifyHaVipAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyHaVipAttributeResponse) SetStatusCode(v int32) *ModifyHaVipAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyHaVipAttributeResponse) SetBody(v *ModifyHaVipAttributeResponseBody) *ModifyHaVipAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyHaVipAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
