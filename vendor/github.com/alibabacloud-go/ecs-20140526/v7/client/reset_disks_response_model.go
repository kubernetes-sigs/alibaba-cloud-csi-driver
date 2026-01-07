// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetDisksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetDisksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetDisksResponse
	GetStatusCode() *int32
	SetBody(v *ResetDisksResponseBody) *ResetDisksResponse
	GetBody() *ResetDisksResponseBody
}

type ResetDisksResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetDisksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetDisksResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetDisksResponse) GoString() string {
	return s.String()
}

func (s *ResetDisksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetDisksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetDisksResponse) GetBody() *ResetDisksResponseBody {
	return s.Body
}

func (s *ResetDisksResponse) SetHeaders(v map[string]*string) *ResetDisksResponse {
	s.Headers = v
	return s
}

func (s *ResetDisksResponse) SetStatusCode(v int32) *ResetDisksResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetDisksResponse) SetBody(v *ResetDisksResponseBody) *ResetDisksResponse {
	s.Body = v
	return s
}

func (s *ResetDisksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
