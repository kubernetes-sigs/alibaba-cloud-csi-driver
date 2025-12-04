// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelFilesetQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelFilesetQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelFilesetQuotaResponse
	GetStatusCode() *int32
	SetBody(v *CancelFilesetQuotaResponseBody) *CancelFilesetQuotaResponse
	GetBody() *CancelFilesetQuotaResponseBody
}

type CancelFilesetQuotaResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelFilesetQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelFilesetQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelFilesetQuotaResponse) GoString() string {
	return s.String()
}

func (s *CancelFilesetQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelFilesetQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelFilesetQuotaResponse) GetBody() *CancelFilesetQuotaResponseBody {
	return s.Body
}

func (s *CancelFilesetQuotaResponse) SetHeaders(v map[string]*string) *CancelFilesetQuotaResponse {
	s.Headers = v
	return s
}

func (s *CancelFilesetQuotaResponse) SetStatusCode(v int32) *CancelFilesetQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelFilesetQuotaResponse) SetBody(v *CancelFilesetQuotaResponseBody) *CancelFilesetQuotaResponse {
	s.Body = v
	return s
}

func (s *CancelFilesetQuotaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
