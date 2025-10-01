// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProtocolMountTargetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateProtocolMountTargetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateProtocolMountTargetResponse
	GetStatusCode() *int32
	SetBody(v *CreateProtocolMountTargetResponseBody) *CreateProtocolMountTargetResponse
	GetBody() *CreateProtocolMountTargetResponseBody
}

type CreateProtocolMountTargetResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateProtocolMountTargetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateProtocolMountTargetResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateProtocolMountTargetResponse) GoString() string {
	return s.String()
}

func (s *CreateProtocolMountTargetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateProtocolMountTargetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateProtocolMountTargetResponse) GetBody() *CreateProtocolMountTargetResponseBody {
	return s.Body
}

func (s *CreateProtocolMountTargetResponse) SetHeaders(v map[string]*string) *CreateProtocolMountTargetResponse {
	s.Headers = v
	return s
}

func (s *CreateProtocolMountTargetResponse) SetStatusCode(v int32) *CreateProtocolMountTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProtocolMountTargetResponse) SetBody(v *CreateProtocolMountTargetResponseBody) *CreateProtocolMountTargetResponse {
	s.Body = v
	return s
}

func (s *CreateProtocolMountTargetResponse) Validate() error {
	return dara.Validate(s)
}
