// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDirectoriesAndFilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDirectoriesAndFilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDirectoriesAndFilesResponse
	GetStatusCode() *int32
	SetBody(v *ListDirectoriesAndFilesResponseBody) *ListDirectoriesAndFilesResponse
	GetBody() *ListDirectoriesAndFilesResponseBody
}

type ListDirectoriesAndFilesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDirectoriesAndFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDirectoriesAndFilesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDirectoriesAndFilesResponse) GoString() string {
	return s.String()
}

func (s *ListDirectoriesAndFilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDirectoriesAndFilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDirectoriesAndFilesResponse) GetBody() *ListDirectoriesAndFilesResponseBody {
	return s.Body
}

func (s *ListDirectoriesAndFilesResponse) SetHeaders(v map[string]*string) *ListDirectoriesAndFilesResponse {
	s.Headers = v
	return s
}

func (s *ListDirectoriesAndFilesResponse) SetStatusCode(v int32) *ListDirectoriesAndFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDirectoriesAndFilesResponse) SetBody(v *ListDirectoriesAndFilesResponseBody) *ListDirectoriesAndFilesResponse {
	s.Body = v
	return s
}

func (s *ListDirectoriesAndFilesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
