// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccessPointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAccessPointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAccessPointResponse
	GetStatusCode() *int32
	SetBody(v *CreateAccessPointResponseBody) *CreateAccessPointResponse
	GetBody() *CreateAccessPointResponseBody
}

type CreateAccessPointResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAccessPointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAccessPointResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAccessPointResponse) GoString() string {
	return s.String()
}

func (s *CreateAccessPointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAccessPointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAccessPointResponse) GetBody() *CreateAccessPointResponseBody {
	return s.Body
}

func (s *CreateAccessPointResponse) SetHeaders(v map[string]*string) *CreateAccessPointResponse {
	s.Headers = v
	return s
}

func (s *CreateAccessPointResponse) SetStatusCode(v int32) *CreateAccessPointResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAccessPointResponse) SetBody(v *CreateAccessPointResponseBody) *CreateAccessPointResponse {
	s.Body = v
	return s
}

func (s *CreateAccessPointResponse) Validate() error {
	return dara.Validate(s)
}
