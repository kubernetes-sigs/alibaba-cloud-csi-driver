// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecentlyRecycledDirectoriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRecentlyRecycledDirectoriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRecentlyRecycledDirectoriesResponse
	GetStatusCode() *int32
	SetBody(v *ListRecentlyRecycledDirectoriesResponseBody) *ListRecentlyRecycledDirectoriesResponse
	GetBody() *ListRecentlyRecycledDirectoriesResponseBody
}

type ListRecentlyRecycledDirectoriesResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRecentlyRecycledDirectoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRecentlyRecycledDirectoriesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRecentlyRecycledDirectoriesResponse) GoString() string {
	return s.String()
}

func (s *ListRecentlyRecycledDirectoriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRecentlyRecycledDirectoriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRecentlyRecycledDirectoriesResponse) GetBody() *ListRecentlyRecycledDirectoriesResponseBody {
	return s.Body
}

func (s *ListRecentlyRecycledDirectoriesResponse) SetHeaders(v map[string]*string) *ListRecentlyRecycledDirectoriesResponse {
	s.Headers = v
	return s
}

func (s *ListRecentlyRecycledDirectoriesResponse) SetStatusCode(v int32) *ListRecentlyRecycledDirectoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRecentlyRecycledDirectoriesResponse) SetBody(v *ListRecentlyRecycledDirectoriesResponseBody) *ListRecentlyRecycledDirectoriesResponse {
	s.Body = v
	return s
}

func (s *ListRecentlyRecycledDirectoriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
