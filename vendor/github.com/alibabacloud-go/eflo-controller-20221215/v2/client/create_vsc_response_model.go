// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVscResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVscResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVscResponse
	GetStatusCode() *int32
	SetBody(v *CreateVscResponseBody) *CreateVscResponse
	GetBody() *CreateVscResponseBody
}

type CreateVscResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVscResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVscResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVscResponse) GoString() string {
	return s.String()
}

func (s *CreateVscResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVscResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVscResponse) GetBody() *CreateVscResponseBody {
	return s.Body
}

func (s *CreateVscResponse) SetHeaders(v map[string]*string) *CreateVscResponse {
	s.Headers = v
	return s
}

func (s *CreateVscResponse) SetStatusCode(v int32) *CreateVscResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVscResponse) SetBody(v *CreateVscResponseBody) *CreateVscResponse {
	s.Body = v
	return s
}

func (s *CreateVscResponse) Validate() error {
	return dara.Validate(s)
}
