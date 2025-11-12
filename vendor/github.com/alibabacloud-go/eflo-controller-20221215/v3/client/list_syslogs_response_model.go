// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSyslogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSyslogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSyslogsResponse
	GetStatusCode() *int32
	SetBody(v *ListSyslogsResponseBody) *ListSyslogsResponse
	GetBody() *ListSyslogsResponseBody
}

type ListSyslogsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSyslogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSyslogsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSyslogsResponse) GoString() string {
	return s.String()
}

func (s *ListSyslogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSyslogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSyslogsResponse) GetBody() *ListSyslogsResponseBody {
	return s.Body
}

func (s *ListSyslogsResponse) SetHeaders(v map[string]*string) *ListSyslogsResponse {
	s.Headers = v
	return s
}

func (s *ListSyslogsResponse) SetStatusCode(v int32) *ListSyslogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSyslogsResponse) SetBody(v *ListSyslogsResponseBody) *ListSyslogsResponse {
	s.Body = v
	return s
}

func (s *ListSyslogsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
