// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRouteEntryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRouteEntryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRouteEntryResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRouteEntryResponseBody) *DeleteRouteEntryResponse
	GetBody() *DeleteRouteEntryResponseBody
}

type DeleteRouteEntryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRouteEntryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRouteEntryResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRouteEntryResponse) GoString() string {
	return s.String()
}

func (s *DeleteRouteEntryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRouteEntryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRouteEntryResponse) GetBody() *DeleteRouteEntryResponseBody {
	return s.Body
}

func (s *DeleteRouteEntryResponse) SetHeaders(v map[string]*string) *DeleteRouteEntryResponse {
	s.Headers = v
	return s
}

func (s *DeleteRouteEntryResponse) SetStatusCode(v int32) *DeleteRouteEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRouteEntryResponse) SetBody(v *DeleteRouteEntryResponseBody) *DeleteRouteEntryResponse {
	s.Body = v
	return s
}

func (s *DeleteRouteEntryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
