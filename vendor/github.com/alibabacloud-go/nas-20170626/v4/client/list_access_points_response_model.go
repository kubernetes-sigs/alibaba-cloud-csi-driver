// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccessPointsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAccessPointsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAccessPointsResponse
	GetStatusCode() *int32
	SetBody(v *ListAccessPointsResponseBody) *ListAccessPointsResponse
	GetBody() *ListAccessPointsResponseBody
}

type ListAccessPointsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAccessPointsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAccessPointsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAccessPointsResponse) GoString() string {
	return s.String()
}

func (s *ListAccessPointsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAccessPointsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAccessPointsResponse) GetBody() *ListAccessPointsResponseBody {
	return s.Body
}

func (s *ListAccessPointsResponse) SetHeaders(v map[string]*string) *ListAccessPointsResponse {
	s.Headers = v
	return s
}

func (s *ListAccessPointsResponse) SetStatusCode(v int32) *ListAccessPointsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAccessPointsResponse) SetBody(v *ListAccessPointsResponseBody) *ListAccessPointsResponse {
	s.Body = v
	return s
}

func (s *ListAccessPointsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
