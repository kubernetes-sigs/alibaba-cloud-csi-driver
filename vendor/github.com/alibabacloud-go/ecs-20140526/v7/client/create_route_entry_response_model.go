// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRouteEntryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRouteEntryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRouteEntryResponse
	GetStatusCode() *int32
	SetBody(v *CreateRouteEntryResponseBody) *CreateRouteEntryResponse
	GetBody() *CreateRouteEntryResponseBody
}

type CreateRouteEntryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRouteEntryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRouteEntryResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRouteEntryResponse) GoString() string {
	return s.String()
}

func (s *CreateRouteEntryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRouteEntryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRouteEntryResponse) GetBody() *CreateRouteEntryResponseBody {
	return s.Body
}

func (s *CreateRouteEntryResponse) SetHeaders(v map[string]*string) *CreateRouteEntryResponse {
	s.Headers = v
	return s
}

func (s *CreateRouteEntryResponse) SetStatusCode(v int32) *CreateRouteEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRouteEntryResponse) SetBody(v *CreateRouteEntryResponseBody) *CreateRouteEntryResponse {
	s.Body = v
	return s
}

func (s *CreateRouteEntryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
