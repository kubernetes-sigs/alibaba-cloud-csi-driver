// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHyperNodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteHyperNodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteHyperNodeResponse
	GetStatusCode() *int32
	SetBody(v *DeleteHyperNodeResponseBody) *DeleteHyperNodeResponse
	GetBody() *DeleteHyperNodeResponseBody
}

type DeleteHyperNodeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHyperNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHyperNodeResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteHyperNodeResponse) GoString() string {
	return s.String()
}

func (s *DeleteHyperNodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteHyperNodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteHyperNodeResponse) GetBody() *DeleteHyperNodeResponseBody {
	return s.Body
}

func (s *DeleteHyperNodeResponse) SetHeaders(v map[string]*string) *DeleteHyperNodeResponse {
	s.Headers = v
	return s
}

func (s *DeleteHyperNodeResponse) SetStatusCode(v int32) *DeleteHyperNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHyperNodeResponse) SetBody(v *DeleteHyperNodeResponseBody) *DeleteHyperNodeResponse {
	s.Body = v
	return s
}

func (s *DeleteHyperNodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
