// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableActivationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableActivationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableActivationResponse
	GetStatusCode() *int32
	SetBody(v *DisableActivationResponseBody) *DisableActivationResponse
	GetBody() *DisableActivationResponseBody
}

type DisableActivationResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableActivationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableActivationResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableActivationResponse) GoString() string {
	return s.String()
}

func (s *DisableActivationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableActivationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableActivationResponse) GetBody() *DisableActivationResponseBody {
	return s.Body
}

func (s *DisableActivationResponse) SetHeaders(v map[string]*string) *DisableActivationResponse {
	s.Headers = v
	return s
}

func (s *DisableActivationResponse) SetStatusCode(v int32) *DisableActivationResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableActivationResponse) SetBody(v *DisableActivationResponseBody) *DisableActivationResponse {
	s.Body = v
	return s
}

func (s *DisableActivationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
