// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelDirQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelDirQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelDirQuotaResponse
	GetStatusCode() *int32
	SetBody(v *CancelDirQuotaResponseBody) *CancelDirQuotaResponse
	GetBody() *CancelDirQuotaResponseBody
}

type CancelDirQuotaResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelDirQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelDirQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelDirQuotaResponse) GoString() string {
	return s.String()
}

func (s *CancelDirQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelDirQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelDirQuotaResponse) GetBody() *CancelDirQuotaResponseBody {
	return s.Body
}

func (s *CancelDirQuotaResponse) SetHeaders(v map[string]*string) *CancelDirQuotaResponse {
	s.Headers = v
	return s
}

func (s *CancelDirQuotaResponse) SetStatusCode(v int32) *CancelDirQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelDirQuotaResponse) SetBody(v *CancelDirQuotaResponseBody) *CancelDirQuotaResponse {
	s.Body = v
	return s
}

func (s *CancelDirQuotaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
