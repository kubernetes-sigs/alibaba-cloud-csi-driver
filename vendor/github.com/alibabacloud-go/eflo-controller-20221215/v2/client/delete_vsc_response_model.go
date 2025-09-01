// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVscResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVscResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVscResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVscResponseBody) *DeleteVscResponse
	GetBody() *DeleteVscResponseBody
}

type DeleteVscResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVscResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVscResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVscResponse) GoString() string {
	return s.String()
}

func (s *DeleteVscResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVscResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVscResponse) GetBody() *DeleteVscResponseBody {
	return s.Body
}

func (s *DeleteVscResponse) SetHeaders(v map[string]*string) *DeleteVscResponse {
	s.Headers = v
	return s
}

func (s *DeleteVscResponse) SetStatusCode(v int32) *DeleteVscResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVscResponse) SetBody(v *DeleteVscResponseBody) *DeleteVscResponse {
	s.Body = v
	return s
}

func (s *DeleteVscResponse) Validate() error {
	return dara.Validate(s)
}
