// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCpfsAccessPointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCpfsAccessPointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCpfsAccessPointResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCpfsAccessPointResponseBody) *DeleteCpfsAccessPointResponse
	GetBody() *DeleteCpfsAccessPointResponseBody
}

type DeleteCpfsAccessPointResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCpfsAccessPointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCpfsAccessPointResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCpfsAccessPointResponse) GoString() string {
	return s.String()
}

func (s *DeleteCpfsAccessPointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCpfsAccessPointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCpfsAccessPointResponse) GetBody() *DeleteCpfsAccessPointResponseBody {
	return s.Body
}

func (s *DeleteCpfsAccessPointResponse) SetHeaders(v map[string]*string) *DeleteCpfsAccessPointResponse {
	s.Headers = v
	return s
}

func (s *DeleteCpfsAccessPointResponse) SetStatusCode(v int32) *DeleteCpfsAccessPointResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCpfsAccessPointResponse) SetBody(v *DeleteCpfsAccessPointResponseBody) *DeleteCpfsAccessPointResponse {
	s.Body = v
	return s
}

func (s *DeleteCpfsAccessPointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
