// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMountTargetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMountTargetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMountTargetResponse
	GetStatusCode() *int32
	SetBody(v *CreateMountTargetResponseBody) *CreateMountTargetResponse
	GetBody() *CreateMountTargetResponseBody
}

type CreateMountTargetResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMountTargetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMountTargetResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMountTargetResponse) GoString() string {
	return s.String()
}

func (s *CreateMountTargetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMountTargetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMountTargetResponse) GetBody() *CreateMountTargetResponseBody {
	return s.Body
}

func (s *CreateMountTargetResponse) SetHeaders(v map[string]*string) *CreateMountTargetResponse {
	s.Headers = v
	return s
}

func (s *CreateMountTargetResponse) SetStatusCode(v int32) *CreateMountTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMountTargetResponse) SetBody(v *CreateMountTargetResponseBody) *CreateMountTargetResponse {
	s.Body = v
	return s
}

func (s *CreateMountTargetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
