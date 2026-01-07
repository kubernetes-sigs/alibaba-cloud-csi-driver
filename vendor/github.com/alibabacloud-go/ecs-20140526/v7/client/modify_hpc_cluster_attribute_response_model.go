// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHpcClusterAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyHpcClusterAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyHpcClusterAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyHpcClusterAttributeResponseBody) *ModifyHpcClusterAttributeResponse
	GetBody() *ModifyHpcClusterAttributeResponseBody
}

type ModifyHpcClusterAttributeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyHpcClusterAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyHpcClusterAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyHpcClusterAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyHpcClusterAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyHpcClusterAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyHpcClusterAttributeResponse) GetBody() *ModifyHpcClusterAttributeResponseBody {
	return s.Body
}

func (s *ModifyHpcClusterAttributeResponse) SetHeaders(v map[string]*string) *ModifyHpcClusterAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyHpcClusterAttributeResponse) SetStatusCode(v int32) *ModifyHpcClusterAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyHpcClusterAttributeResponse) SetBody(v *ModifyHpcClusterAttributeResponseBody) *ModifyHpcClusterAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyHpcClusterAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
