// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProtocolServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteProtocolServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteProtocolServiceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteProtocolServiceResponseBody) *DeleteProtocolServiceResponse
	GetBody() *DeleteProtocolServiceResponseBody
}

type DeleteProtocolServiceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteProtocolServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteProtocolServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteProtocolServiceResponse) GoString() string {
	return s.String()
}

func (s *DeleteProtocolServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteProtocolServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteProtocolServiceResponse) GetBody() *DeleteProtocolServiceResponseBody {
	return s.Body
}

func (s *DeleteProtocolServiceResponse) SetHeaders(v map[string]*string) *DeleteProtocolServiceResponse {
	s.Headers = v
	return s
}

func (s *DeleteProtocolServiceResponse) SetStatusCode(v int32) *DeleteProtocolServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteProtocolServiceResponse) SetBody(v *DeleteProtocolServiceResponseBody) *DeleteProtocolServiceResponse {
	s.Body = v
	return s
}

func (s *DeleteProtocolServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
