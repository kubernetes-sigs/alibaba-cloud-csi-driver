// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVSwitchAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyVSwitchAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyVSwitchAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyVSwitchAttributeResponseBody) *ModifyVSwitchAttributeResponse
	GetBody() *ModifyVSwitchAttributeResponseBody
}

type ModifyVSwitchAttributeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyVSwitchAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyVSwitchAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyVSwitchAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyVSwitchAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyVSwitchAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyVSwitchAttributeResponse) GetBody() *ModifyVSwitchAttributeResponseBody {
	return s.Body
}

func (s *ModifyVSwitchAttributeResponse) SetHeaders(v map[string]*string) *ModifyVSwitchAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyVSwitchAttributeResponse) SetStatusCode(v int32) *ModifyVSwitchAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyVSwitchAttributeResponse) SetBody(v *ModifyVSwitchAttributeResponseBody) *ModifyVSwitchAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyVSwitchAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
