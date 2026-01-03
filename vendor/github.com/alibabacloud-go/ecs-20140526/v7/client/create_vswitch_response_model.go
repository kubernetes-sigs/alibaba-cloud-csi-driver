// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVSwitchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVSwitchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVSwitchResponse
	GetStatusCode() *int32
	SetBody(v *CreateVSwitchResponseBody) *CreateVSwitchResponse
	GetBody() *CreateVSwitchResponseBody
}

type CreateVSwitchResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVSwitchResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVSwitchResponse) GoString() string {
	return s.String()
}

func (s *CreateVSwitchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVSwitchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVSwitchResponse) GetBody() *CreateVSwitchResponseBody {
	return s.Body
}

func (s *CreateVSwitchResponse) SetHeaders(v map[string]*string) *CreateVSwitchResponse {
	s.Headers = v
	return s
}

func (s *CreateVSwitchResponse) SetStatusCode(v int32) *CreateVSwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVSwitchResponse) SetBody(v *CreateVSwitchResponseBody) *CreateVSwitchResponse {
	s.Body = v
	return s
}

func (s *CreateVSwitchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
