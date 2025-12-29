// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachDiskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachDiskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachDiskResponse
	GetStatusCode() *int32
	SetBody(v *AttachDiskResponseBody) *AttachDiskResponse
	GetBody() *AttachDiskResponseBody
}

type AttachDiskResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachDiskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachDiskResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachDiskResponse) GoString() string {
	return s.String()
}

func (s *AttachDiskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachDiskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachDiskResponse) GetBody() *AttachDiskResponseBody {
	return s.Body
}

func (s *AttachDiskResponse) SetHeaders(v map[string]*string) *AttachDiskResponse {
	s.Headers = v
	return s
}

func (s *AttachDiskResponse) SetStatusCode(v int32) *AttachDiskResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachDiskResponse) SetBody(v *AttachDiskResponseBody) *AttachDiskResponse {
	s.Body = v
	return s
}

func (s *AttachDiskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
