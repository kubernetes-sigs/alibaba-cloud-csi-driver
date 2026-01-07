// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetDiskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetDiskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetDiskResponse
	GetStatusCode() *int32
	SetBody(v *ResetDiskResponseBody) *ResetDiskResponse
	GetBody() *ResetDiskResponseBody
}

type ResetDiskResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetDiskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetDiskResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetDiskResponse) GoString() string {
	return s.String()
}

func (s *ResetDiskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetDiskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetDiskResponse) GetBody() *ResetDiskResponseBody {
	return s.Body
}

func (s *ResetDiskResponse) SetHeaders(v map[string]*string) *ResetDiskResponse {
	s.Headers = v
	return s
}

func (s *ResetDiskResponse) SetStatusCode(v int32) *ResetDiskResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetDiskResponse) SetBody(v *ResetDiskResponseBody) *ResetDiskResponse {
	s.Body = v
	return s
}

func (s *ResetDiskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
