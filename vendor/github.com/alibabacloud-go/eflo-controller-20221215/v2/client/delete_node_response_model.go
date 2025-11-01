// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteNodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteNodeResponse
	GetStatusCode() *int32
	SetBody(v *DeleteNodeResponseBody) *DeleteNodeResponse
	GetBody() *DeleteNodeResponseBody
}

type DeleteNodeResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNodeResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteNodeResponse) GoString() string {
	return s.String()
}

func (s *DeleteNodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteNodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteNodeResponse) GetBody() *DeleteNodeResponseBody {
	return s.Body
}

func (s *DeleteNodeResponse) SetHeaders(v map[string]*string) *DeleteNodeResponse {
	s.Headers = v
	return s
}

func (s *DeleteNodeResponse) SetStatusCode(v int32) *DeleteNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNodeResponse) SetBody(v *DeleteNodeResponseBody) *DeleteNodeResponse {
	s.Body = v
	return s
}

func (s *DeleteNodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
