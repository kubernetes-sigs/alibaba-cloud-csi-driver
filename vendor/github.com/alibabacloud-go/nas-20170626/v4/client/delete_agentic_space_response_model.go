// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgenticSpaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAgenticSpaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAgenticSpaceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAgenticSpaceResponseBody) *DeleteAgenticSpaceResponse
	GetBody() *DeleteAgenticSpaceResponseBody
}

type DeleteAgenticSpaceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAgenticSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAgenticSpaceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgenticSpaceResponse) GoString() string {
	return s.String()
}

func (s *DeleteAgenticSpaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAgenticSpaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAgenticSpaceResponse) GetBody() *DeleteAgenticSpaceResponseBody {
	return s.Body
}

func (s *DeleteAgenticSpaceResponse) SetHeaders(v map[string]*string) *DeleteAgenticSpaceResponse {
	s.Headers = v
	return s
}

func (s *DeleteAgenticSpaceResponse) SetStatusCode(v int32) *DeleteAgenticSpaceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAgenticSpaceResponse) SetBody(v *DeleteAgenticSpaceResponseBody) *DeleteAgenticSpaceResponse {
	s.Body = v
	return s
}

func (s *DeleteAgenticSpaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
