// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceSystemDiskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReplaceSystemDiskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReplaceSystemDiskResponse
	GetStatusCode() *int32
	SetBody(v *ReplaceSystemDiskResponseBody) *ReplaceSystemDiskResponse
	GetBody() *ReplaceSystemDiskResponseBody
}

type ReplaceSystemDiskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReplaceSystemDiskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReplaceSystemDiskResponse) String() string {
	return dara.Prettify(s)
}

func (s ReplaceSystemDiskResponse) GoString() string {
	return s.String()
}

func (s *ReplaceSystemDiskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReplaceSystemDiskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReplaceSystemDiskResponse) GetBody() *ReplaceSystemDiskResponseBody {
	return s.Body
}

func (s *ReplaceSystemDiskResponse) SetHeaders(v map[string]*string) *ReplaceSystemDiskResponse {
	s.Headers = v
	return s
}

func (s *ReplaceSystemDiskResponse) SetStatusCode(v int32) *ReplaceSystemDiskResponse {
	s.StatusCode = &v
	return s
}

func (s *ReplaceSystemDiskResponse) SetBody(v *ReplaceSystemDiskResponseBody) *ReplaceSystemDiskResponse {
	s.Body = v
	return s
}

func (s *ReplaceSystemDiskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
