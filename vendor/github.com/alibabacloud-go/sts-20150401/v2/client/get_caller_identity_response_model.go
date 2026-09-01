// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCallerIdentityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCallerIdentityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCallerIdentityResponse
	GetStatusCode() *int32
	SetBody(v *GetCallerIdentityResponseBody) *GetCallerIdentityResponse
	GetBody() *GetCallerIdentityResponseBody
}

type GetCallerIdentityResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCallerIdentityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCallerIdentityResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCallerIdentityResponse) GoString() string {
	return s.String()
}

func (s *GetCallerIdentityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCallerIdentityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCallerIdentityResponse) GetBody() *GetCallerIdentityResponseBody {
	return s.Body
}

func (s *GetCallerIdentityResponse) SetHeaders(v map[string]*string) *GetCallerIdentityResponse {
	s.Headers = v
	return s
}

func (s *GetCallerIdentityResponse) SetStatusCode(v int32) *GetCallerIdentityResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCallerIdentityResponse) SetBody(v *GetCallerIdentityResponseBody) *GetCallerIdentityResponse {
	s.Body = v
	return s
}

func (s *GetCallerIdentityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
