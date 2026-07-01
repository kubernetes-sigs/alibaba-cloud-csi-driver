// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgenticSpaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAgenticSpaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAgenticSpaceResponse
	GetStatusCode() *int32
	SetBody(v *GetAgenticSpaceResponseBody) *GetAgenticSpaceResponse
	GetBody() *GetAgenticSpaceResponseBody
}

type GetAgenticSpaceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgenticSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgenticSpaceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgenticSpaceResponse) GoString() string {
	return s.String()
}

func (s *GetAgenticSpaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAgenticSpaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAgenticSpaceResponse) GetBody() *GetAgenticSpaceResponseBody {
	return s.Body
}

func (s *GetAgenticSpaceResponse) SetHeaders(v map[string]*string) *GetAgenticSpaceResponse {
	s.Headers = v
	return s
}

func (s *GetAgenticSpaceResponse) SetStatusCode(v int32) *GetAgenticSpaceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgenticSpaceResponse) SetBody(v *GetAgenticSpaceResponseBody) *GetAgenticSpaceResponse {
	s.Body = v
	return s
}

func (s *GetAgenticSpaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
