// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachDiskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachDiskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachDiskResponse
	GetStatusCode() *int32
	SetBody(v *DetachDiskResponseBody) *DetachDiskResponse
	GetBody() *DetachDiskResponseBody
}

type DetachDiskResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachDiskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachDiskResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachDiskResponse) GoString() string {
	return s.String()
}

func (s *DetachDiskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachDiskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachDiskResponse) GetBody() *DetachDiskResponseBody {
	return s.Body
}

func (s *DetachDiskResponse) SetHeaders(v map[string]*string) *DetachDiskResponse {
	s.Headers = v
	return s
}

func (s *DetachDiskResponse) SetStatusCode(v int32) *DetachDiskResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachDiskResponse) SetBody(v *DetachDiskResponseBody) *DetachDiskResponse {
	s.Body = v
	return s
}

func (s *DetachDiskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
