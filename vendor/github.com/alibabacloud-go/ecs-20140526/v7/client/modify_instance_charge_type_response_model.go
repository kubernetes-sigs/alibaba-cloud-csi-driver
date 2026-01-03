// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceChargeTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyInstanceChargeTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyInstanceChargeTypeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyInstanceChargeTypeResponseBody) *ModifyInstanceChargeTypeResponse
	GetBody() *ModifyInstanceChargeTypeResponseBody
}

type ModifyInstanceChargeTypeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceChargeTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceChargeTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceChargeTypeResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceChargeTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyInstanceChargeTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyInstanceChargeTypeResponse) GetBody() *ModifyInstanceChargeTypeResponseBody {
	return s.Body
}

func (s *ModifyInstanceChargeTypeResponse) SetHeaders(v map[string]*string) *ModifyInstanceChargeTypeResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceChargeTypeResponse) SetStatusCode(v int32) *ModifyInstanceChargeTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceChargeTypeResponse) SetBody(v *ModifyInstanceChargeTypeResponseBody) *ModifyInstanceChargeTypeResponse {
	s.Body = v
	return s
}

func (s *ModifyInstanceChargeTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
