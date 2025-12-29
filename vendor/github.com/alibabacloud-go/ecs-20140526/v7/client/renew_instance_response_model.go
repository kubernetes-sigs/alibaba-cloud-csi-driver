// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RenewInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RenewInstanceResponse
	GetStatusCode() *int32
	SetBody(v *RenewInstanceResponseBody) *RenewInstanceResponse
	GetBody() *RenewInstanceResponseBody
}

type RenewInstanceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenewInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenewInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s RenewInstanceResponse) GoString() string {
	return s.String()
}

func (s *RenewInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RenewInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RenewInstanceResponse) GetBody() *RenewInstanceResponseBody {
	return s.Body
}

func (s *RenewInstanceResponse) SetHeaders(v map[string]*string) *RenewInstanceResponse {
	s.Headers = v
	return s
}

func (s *RenewInstanceResponse) SetStatusCode(v int32) *RenewInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RenewInstanceResponse) SetBody(v *RenewInstanceResponseBody) *RenewInstanceResponse {
	s.Body = v
	return s
}

func (s *RenewInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
