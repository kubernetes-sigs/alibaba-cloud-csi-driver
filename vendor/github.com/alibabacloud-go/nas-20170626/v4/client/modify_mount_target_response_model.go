// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMountTargetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyMountTargetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyMountTargetResponse
	GetStatusCode() *int32
	SetBody(v *ModifyMountTargetResponseBody) *ModifyMountTargetResponse
	GetBody() *ModifyMountTargetResponseBody
}

type ModifyMountTargetResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyMountTargetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyMountTargetResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyMountTargetResponse) GoString() string {
	return s.String()
}

func (s *ModifyMountTargetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyMountTargetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyMountTargetResponse) GetBody() *ModifyMountTargetResponseBody {
	return s.Body
}

func (s *ModifyMountTargetResponse) SetHeaders(v map[string]*string) *ModifyMountTargetResponse {
	s.Headers = v
	return s
}

func (s *ModifyMountTargetResponse) SetStatusCode(v int32) *ModifyMountTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyMountTargetResponse) SetBody(v *ModifyMountTargetResponseBody) *ModifyMountTargetResponse {
	s.Body = v
	return s
}

func (s *ModifyMountTargetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
