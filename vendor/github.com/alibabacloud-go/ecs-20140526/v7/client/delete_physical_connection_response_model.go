// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePhysicalConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePhysicalConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePhysicalConnectionResponse
	GetStatusCode() *int32
	SetBody(v *DeletePhysicalConnectionResponseBody) *DeletePhysicalConnectionResponse
	GetBody() *DeletePhysicalConnectionResponseBody
}

type DeletePhysicalConnectionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePhysicalConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePhysicalConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePhysicalConnectionResponse) GoString() string {
	return s.String()
}

func (s *DeletePhysicalConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePhysicalConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePhysicalConnectionResponse) GetBody() *DeletePhysicalConnectionResponseBody {
	return s.Body
}

func (s *DeletePhysicalConnectionResponse) SetHeaders(v map[string]*string) *DeletePhysicalConnectionResponse {
	s.Headers = v
	return s
}

func (s *DeletePhysicalConnectionResponse) SetStatusCode(v int32) *DeletePhysicalConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePhysicalConnectionResponse) SetBody(v *DeletePhysicalConnectionResponseBody) *DeletePhysicalConnectionResponse {
	s.Body = v
	return s
}

func (s *DeletePhysicalConnectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
