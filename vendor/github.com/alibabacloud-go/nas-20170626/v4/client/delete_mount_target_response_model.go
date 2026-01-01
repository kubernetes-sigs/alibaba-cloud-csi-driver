// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMountTargetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMountTargetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMountTargetResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMountTargetResponseBody) *DeleteMountTargetResponse
	GetBody() *DeleteMountTargetResponseBody
}

type DeleteMountTargetResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMountTargetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMountTargetResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMountTargetResponse) GoString() string {
	return s.String()
}

func (s *DeleteMountTargetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMountTargetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMountTargetResponse) GetBody() *DeleteMountTargetResponseBody {
	return s.Body
}

func (s *DeleteMountTargetResponse) SetHeaders(v map[string]*string) *DeleteMountTargetResponse {
	s.Headers = v
	return s
}

func (s *DeleteMountTargetResponse) SetStatusCode(v int32) *DeleteMountTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMountTargetResponse) SetBody(v *DeleteMountTargetResponseBody) *DeleteMountTargetResponse {
	s.Body = v
	return s
}

func (s *DeleteMountTargetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
