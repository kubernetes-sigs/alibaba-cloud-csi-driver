// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCpfsAccessPointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCpfsAccessPointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCpfsAccessPointResponse
	GetStatusCode() *int32
	SetBody(v *CreateCpfsAccessPointResponseBody) *CreateCpfsAccessPointResponse
	GetBody() *CreateCpfsAccessPointResponseBody
}

type CreateCpfsAccessPointResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCpfsAccessPointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCpfsAccessPointResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCpfsAccessPointResponse) GoString() string {
	return s.String()
}

func (s *CreateCpfsAccessPointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCpfsAccessPointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCpfsAccessPointResponse) GetBody() *CreateCpfsAccessPointResponseBody {
	return s.Body
}

func (s *CreateCpfsAccessPointResponse) SetHeaders(v map[string]*string) *CreateCpfsAccessPointResponse {
	s.Headers = v
	return s
}

func (s *CreateCpfsAccessPointResponse) SetStatusCode(v int32) *CreateCpfsAccessPointResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCpfsAccessPointResponse) SetBody(v *CreateCpfsAccessPointResponseBody) *CreateCpfsAccessPointResponse {
	s.Body = v
	return s
}

func (s *CreateCpfsAccessPointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
