// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReInitDiskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReInitDiskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReInitDiskResponse
	GetStatusCode() *int32
	SetBody(v *ReInitDiskResponseBody) *ReInitDiskResponse
	GetBody() *ReInitDiskResponseBody
}

type ReInitDiskResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReInitDiskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReInitDiskResponse) String() string {
	return dara.Prettify(s)
}

func (s ReInitDiskResponse) GoString() string {
	return s.String()
}

func (s *ReInitDiskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReInitDiskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReInitDiskResponse) GetBody() *ReInitDiskResponseBody {
	return s.Body
}

func (s *ReInitDiskResponse) SetHeaders(v map[string]*string) *ReInitDiskResponse {
	s.Headers = v
	return s
}

func (s *ReInitDiskResponse) SetStatusCode(v int32) *ReInitDiskResponse {
	s.StatusCode = &v
	return s
}

func (s *ReInitDiskResponse) SetBody(v *ReInitDiskResponseBody) *ReInitDiskResponse {
	s.Body = v
	return s
}

func (s *ReInitDiskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
