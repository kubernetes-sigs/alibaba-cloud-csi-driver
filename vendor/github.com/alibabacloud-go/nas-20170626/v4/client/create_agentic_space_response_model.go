// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgenticSpaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAgenticSpaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAgenticSpaceResponse
	GetStatusCode() *int32
	SetBody(v *CreateAgenticSpaceResponseBody) *CreateAgenticSpaceResponse
	GetBody() *CreateAgenticSpaceResponseBody
}

type CreateAgenticSpaceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAgenticSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAgenticSpaceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAgenticSpaceResponse) GoString() string {
	return s.String()
}

func (s *CreateAgenticSpaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAgenticSpaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAgenticSpaceResponse) GetBody() *CreateAgenticSpaceResponseBody {
	return s.Body
}

func (s *CreateAgenticSpaceResponse) SetHeaders(v map[string]*string) *CreateAgenticSpaceResponse {
	s.Headers = v
	return s
}

func (s *CreateAgenticSpaceResponse) SetStatusCode(v int32) *CreateAgenticSpaceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAgenticSpaceResponse) SetBody(v *CreateAgenticSpaceResponseBody) *CreateAgenticSpaceResponse {
	s.Body = v
	return s
}

func (s *CreateAgenticSpaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
