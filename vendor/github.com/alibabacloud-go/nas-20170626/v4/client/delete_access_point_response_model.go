// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccessPointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAccessPointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAccessPointResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAccessPointResponseBody) *DeleteAccessPointResponse
	GetBody() *DeleteAccessPointResponseBody
}

type DeleteAccessPointResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAccessPointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAccessPointResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccessPointResponse) GoString() string {
	return s.String()
}

func (s *DeleteAccessPointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAccessPointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAccessPointResponse) GetBody() *DeleteAccessPointResponseBody {
	return s.Body
}

func (s *DeleteAccessPointResponse) SetHeaders(v map[string]*string) *DeleteAccessPointResponse {
	s.Headers = v
	return s
}

func (s *DeleteAccessPointResponse) SetStatusCode(v int32) *DeleteAccessPointResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAccessPointResponse) SetBody(v *DeleteAccessPointResponseBody) *DeleteAccessPointResponse {
	s.Body = v
	return s
}

func (s *DeleteAccessPointResponse) Validate() error {
	return dara.Validate(s)
}
