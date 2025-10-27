// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProtocolServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateProtocolServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateProtocolServiceResponse
	GetStatusCode() *int32
	SetBody(v *CreateProtocolServiceResponseBody) *CreateProtocolServiceResponse
	GetBody() *CreateProtocolServiceResponseBody
}

type CreateProtocolServiceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateProtocolServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateProtocolServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateProtocolServiceResponse) GoString() string {
	return s.String()
}

func (s *CreateProtocolServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateProtocolServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateProtocolServiceResponse) GetBody() *CreateProtocolServiceResponseBody {
	return s.Body
}

func (s *CreateProtocolServiceResponse) SetHeaders(v map[string]*string) *CreateProtocolServiceResponse {
	s.Headers = v
	return s
}

func (s *CreateProtocolServiceResponse) SetStatusCode(v int32) *CreateProtocolServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProtocolServiceResponse) SetBody(v *CreateProtocolServiceResponseBody) *CreateProtocolServiceResponse {
	s.Body = v
	return s
}

func (s *CreateProtocolServiceResponse) Validate() error {
	return dara.Validate(s)
}
