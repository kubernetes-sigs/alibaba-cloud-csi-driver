// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataInsightDirectoriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataInsightDirectoriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataInsightDirectoriesResponse
	GetStatusCode() *int32
	SetBody(v *ListDataInsightDirectoriesResponseBody) *ListDataInsightDirectoriesResponse
	GetBody() *ListDataInsightDirectoriesResponseBody
}

type ListDataInsightDirectoriesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataInsightDirectoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataInsightDirectoriesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataInsightDirectoriesResponse) GoString() string {
	return s.String()
}

func (s *ListDataInsightDirectoriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataInsightDirectoriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataInsightDirectoriesResponse) GetBody() *ListDataInsightDirectoriesResponseBody {
	return s.Body
}

func (s *ListDataInsightDirectoriesResponse) SetHeaders(v map[string]*string) *ListDataInsightDirectoriesResponse {
	s.Headers = v
	return s
}

func (s *ListDataInsightDirectoriesResponse) SetStatusCode(v int32) *ListDataInsightDirectoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataInsightDirectoriesResponse) SetBody(v *ListDataInsightDirectoriesResponseBody) *ListDataInsightDirectoriesResponse {
	s.Body = v
	return s
}

func (s *ListDataInsightDirectoriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
