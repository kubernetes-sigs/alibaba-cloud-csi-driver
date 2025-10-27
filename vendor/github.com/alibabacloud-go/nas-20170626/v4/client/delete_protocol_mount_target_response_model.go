// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProtocolMountTargetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteProtocolMountTargetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteProtocolMountTargetResponse
	GetStatusCode() *int32
	SetBody(v *DeleteProtocolMountTargetResponseBody) *DeleteProtocolMountTargetResponse
	GetBody() *DeleteProtocolMountTargetResponseBody
}

type DeleteProtocolMountTargetResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteProtocolMountTargetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteProtocolMountTargetResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteProtocolMountTargetResponse) GoString() string {
	return s.String()
}

func (s *DeleteProtocolMountTargetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteProtocolMountTargetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteProtocolMountTargetResponse) GetBody() *DeleteProtocolMountTargetResponseBody {
	return s.Body
}

func (s *DeleteProtocolMountTargetResponse) SetHeaders(v map[string]*string) *DeleteProtocolMountTargetResponse {
	s.Headers = v
	return s
}

func (s *DeleteProtocolMountTargetResponse) SetStatusCode(v int32) *DeleteProtocolMountTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteProtocolMountTargetResponse) SetBody(v *DeleteProtocolMountTargetResponseBody) *DeleteProtocolMountTargetResponse {
	s.Body = v
	return s
}

func (s *DeleteProtocolMountTargetResponse) Validate() error {
	return dara.Validate(s)
}
