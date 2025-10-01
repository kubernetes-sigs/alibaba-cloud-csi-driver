// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyProtocolMountTargetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyProtocolMountTargetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyProtocolMountTargetResponse
	GetStatusCode() *int32
	SetBody(v *ModifyProtocolMountTargetResponseBody) *ModifyProtocolMountTargetResponse
	GetBody() *ModifyProtocolMountTargetResponseBody
}

type ModifyProtocolMountTargetResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyProtocolMountTargetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyProtocolMountTargetResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyProtocolMountTargetResponse) GoString() string {
	return s.String()
}

func (s *ModifyProtocolMountTargetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyProtocolMountTargetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyProtocolMountTargetResponse) GetBody() *ModifyProtocolMountTargetResponseBody {
	return s.Body
}

func (s *ModifyProtocolMountTargetResponse) SetHeaders(v map[string]*string) *ModifyProtocolMountTargetResponse {
	s.Headers = v
	return s
}

func (s *ModifyProtocolMountTargetResponse) SetStatusCode(v int32) *ModifyProtocolMountTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyProtocolMountTargetResponse) SetBody(v *ModifyProtocolMountTargetResponseBody) *ModifyProtocolMountTargetResponse {
	s.Body = v
	return s
}

func (s *ModifyProtocolMountTargetResponse) Validate() error {
	return dara.Validate(s)
}
