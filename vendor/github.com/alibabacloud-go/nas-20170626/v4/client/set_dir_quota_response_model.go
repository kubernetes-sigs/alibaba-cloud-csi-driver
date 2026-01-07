// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDirQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetDirQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetDirQuotaResponse
	GetStatusCode() *int32
	SetBody(v *SetDirQuotaResponseBody) *SetDirQuotaResponse
	GetBody() *SetDirQuotaResponseBody
}

type SetDirQuotaResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDirQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDirQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s SetDirQuotaResponse) GoString() string {
	return s.String()
}

func (s *SetDirQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetDirQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetDirQuotaResponse) GetBody() *SetDirQuotaResponseBody {
	return s.Body
}

func (s *SetDirQuotaResponse) SetHeaders(v map[string]*string) *SetDirQuotaResponse {
	s.Headers = v
	return s
}

func (s *SetDirQuotaResponse) SetStatusCode(v int32) *SetDirQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDirQuotaResponse) SetBody(v *SetDirQuotaResponseBody) *SetDirQuotaResponse {
	s.Body = v
	return s
}

func (s *SetDirQuotaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
