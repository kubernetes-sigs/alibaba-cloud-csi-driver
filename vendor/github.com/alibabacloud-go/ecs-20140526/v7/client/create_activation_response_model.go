// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateActivationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateActivationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateActivationResponse
	GetStatusCode() *int32
	SetBody(v *CreateActivationResponseBody) *CreateActivationResponse
	GetBody() *CreateActivationResponseBody
}

type CreateActivationResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateActivationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateActivationResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateActivationResponse) GoString() string {
	return s.String()
}

func (s *CreateActivationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateActivationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateActivationResponse) GetBody() *CreateActivationResponseBody {
	return s.Body
}

func (s *CreateActivationResponse) SetHeaders(v map[string]*string) *CreateActivationResponse {
	s.Headers = v
	return s
}

func (s *CreateActivationResponse) SetStatusCode(v int32) *CreateActivationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateActivationResponse) SetBody(v *CreateActivationResponseBody) *CreateActivationResponse {
	s.Body = v
	return s
}

func (s *CreateActivationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
