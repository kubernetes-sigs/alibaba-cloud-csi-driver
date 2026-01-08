// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFilesetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFilesetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFilesetResponse
	GetStatusCode() *int32
	SetBody(v *GetFilesetResponseBody) *GetFilesetResponse
	GetBody() *GetFilesetResponseBody
}

type GetFilesetResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFilesetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFilesetResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFilesetResponse) GoString() string {
	return s.String()
}

func (s *GetFilesetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFilesetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFilesetResponse) GetBody() *GetFilesetResponseBody {
	return s.Body
}

func (s *GetFilesetResponse) SetHeaders(v map[string]*string) *GetFilesetResponse {
	s.Headers = v
	return s
}

func (s *GetFilesetResponse) SetStatusCode(v int32) *GetFilesetResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFilesetResponse) SetBody(v *GetFilesetResponseBody) *GetFilesetResponse {
	s.Body = v
	return s
}

func (s *GetFilesetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
