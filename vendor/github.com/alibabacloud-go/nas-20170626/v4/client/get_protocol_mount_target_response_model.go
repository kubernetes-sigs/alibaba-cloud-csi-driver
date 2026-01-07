// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProtocolMountTargetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetProtocolMountTargetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetProtocolMountTargetResponse
	GetStatusCode() *int32
	SetBody(v *GetProtocolMountTargetResponseBody) *GetProtocolMountTargetResponse
	GetBody() *GetProtocolMountTargetResponseBody
}

type GetProtocolMountTargetResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProtocolMountTargetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProtocolMountTargetResponse) String() string {
	return dara.Prettify(s)
}

func (s GetProtocolMountTargetResponse) GoString() string {
	return s.String()
}

func (s *GetProtocolMountTargetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetProtocolMountTargetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetProtocolMountTargetResponse) GetBody() *GetProtocolMountTargetResponseBody {
	return s.Body
}

func (s *GetProtocolMountTargetResponse) SetHeaders(v map[string]*string) *GetProtocolMountTargetResponse {
	s.Headers = v
	return s
}

func (s *GetProtocolMountTargetResponse) SetStatusCode(v int32) *GetProtocolMountTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProtocolMountTargetResponse) SetBody(v *GetProtocolMountTargetResponseBody) *GetProtocolMountTargetResponse {
	s.Body = v
	return s
}

func (s *GetProtocolMountTargetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
