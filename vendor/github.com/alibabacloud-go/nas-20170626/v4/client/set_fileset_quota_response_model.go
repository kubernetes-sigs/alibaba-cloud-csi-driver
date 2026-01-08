// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetFilesetQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetFilesetQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetFilesetQuotaResponse
	GetStatusCode() *int32
	SetBody(v *SetFilesetQuotaResponseBody) *SetFilesetQuotaResponse
	GetBody() *SetFilesetQuotaResponseBody
}

type SetFilesetQuotaResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetFilesetQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetFilesetQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s SetFilesetQuotaResponse) GoString() string {
	return s.String()
}

func (s *SetFilesetQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetFilesetQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetFilesetQuotaResponse) GetBody() *SetFilesetQuotaResponseBody {
	return s.Body
}

func (s *SetFilesetQuotaResponse) SetHeaders(v map[string]*string) *SetFilesetQuotaResponse {
	s.Headers = v
	return s
}

func (s *SetFilesetQuotaResponse) SetStatusCode(v int32) *SetFilesetQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *SetFilesetQuotaResponse) SetBody(v *SetFilesetQuotaResponseBody) *SetFilesetQuotaResponse {
	s.Body = v
	return s
}

func (s *SetFilesetQuotaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
