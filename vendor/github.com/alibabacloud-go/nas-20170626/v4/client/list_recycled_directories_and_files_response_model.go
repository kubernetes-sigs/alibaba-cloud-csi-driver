// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecycledDirectoriesAndFilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRecycledDirectoriesAndFilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRecycledDirectoriesAndFilesResponse
	GetStatusCode() *int32
	SetBody(v *ListRecycledDirectoriesAndFilesResponseBody) *ListRecycledDirectoriesAndFilesResponse
	GetBody() *ListRecycledDirectoriesAndFilesResponseBody
}

type ListRecycledDirectoriesAndFilesResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRecycledDirectoriesAndFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRecycledDirectoriesAndFilesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRecycledDirectoriesAndFilesResponse) GoString() string {
	return s.String()
}

func (s *ListRecycledDirectoriesAndFilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRecycledDirectoriesAndFilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRecycledDirectoriesAndFilesResponse) GetBody() *ListRecycledDirectoriesAndFilesResponseBody {
	return s.Body
}

func (s *ListRecycledDirectoriesAndFilesResponse) SetHeaders(v map[string]*string) *ListRecycledDirectoriesAndFilesResponse {
	s.Headers = v
	return s
}

func (s *ListRecycledDirectoriesAndFilesResponse) SetStatusCode(v int32) *ListRecycledDirectoriesAndFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRecycledDirectoriesAndFilesResponse) SetBody(v *ListRecycledDirectoriesAndFilesResponseBody) *ListRecycledDirectoriesAndFilesResponse {
	s.Body = v
	return s
}

func (s *ListRecycledDirectoriesAndFilesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
