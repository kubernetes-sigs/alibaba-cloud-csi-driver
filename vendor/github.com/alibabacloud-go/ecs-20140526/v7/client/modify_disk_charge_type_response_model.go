// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDiskChargeTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDiskChargeTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDiskChargeTypeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDiskChargeTypeResponseBody) *ModifyDiskChargeTypeResponse
	GetBody() *ModifyDiskChargeTypeResponseBody
}

type ModifyDiskChargeTypeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDiskChargeTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDiskChargeTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDiskChargeTypeResponse) GoString() string {
	return s.String()
}

func (s *ModifyDiskChargeTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDiskChargeTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDiskChargeTypeResponse) GetBody() *ModifyDiskChargeTypeResponseBody {
	return s.Body
}

func (s *ModifyDiskChargeTypeResponse) SetHeaders(v map[string]*string) *ModifyDiskChargeTypeResponse {
	s.Headers = v
	return s
}

func (s *ModifyDiskChargeTypeResponse) SetStatusCode(v int32) *ModifyDiskChargeTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDiskChargeTypeResponse) SetBody(v *ModifyDiskChargeTypeResponseBody) *ModifyDiskChargeTypeResponse {
	s.Body = v
	return s
}

func (s *ModifyDiskChargeTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
