// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteImageComponentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteImageComponentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteImageComponentResponse
	GetStatusCode() *int32
	SetBody(v *DeleteImageComponentResponseBody) *DeleteImageComponentResponse
	GetBody() *DeleteImageComponentResponseBody
}

type DeleteImageComponentResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteImageComponentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteImageComponentResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteImageComponentResponse) GoString() string {
	return s.String()
}

func (s *DeleteImageComponentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteImageComponentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteImageComponentResponse) GetBody() *DeleteImageComponentResponseBody {
	return s.Body
}

func (s *DeleteImageComponentResponse) SetHeaders(v map[string]*string) *DeleteImageComponentResponse {
	s.Headers = v
	return s
}

func (s *DeleteImageComponentResponse) SetStatusCode(v int32) *DeleteImageComponentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteImageComponentResponse) SetBody(v *DeleteImageComponentResponseBody) *DeleteImageComponentResponse {
	s.Body = v
	return s
}

func (s *DeleteImageComponentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
