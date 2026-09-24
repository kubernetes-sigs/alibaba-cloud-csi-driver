// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodeGroupRefreshTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNodeGroupRefreshTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNodeGroupRefreshTasksResponse
	GetStatusCode() *int32
	SetBody(v *ListNodeGroupRefreshTasksResponseBody) *ListNodeGroupRefreshTasksResponse
	GetBody() *ListNodeGroupRefreshTasksResponseBody
}

type ListNodeGroupRefreshTasksResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNodeGroupRefreshTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNodeGroupRefreshTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNodeGroupRefreshTasksResponse) GoString() string {
	return s.String()
}

func (s *ListNodeGroupRefreshTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNodeGroupRefreshTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNodeGroupRefreshTasksResponse) GetBody() *ListNodeGroupRefreshTasksResponseBody {
	return s.Body
}

func (s *ListNodeGroupRefreshTasksResponse) SetHeaders(v map[string]*string) *ListNodeGroupRefreshTasksResponse {
	s.Headers = v
	return s
}

func (s *ListNodeGroupRefreshTasksResponse) SetStatusCode(v int32) *ListNodeGroupRefreshTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNodeGroupRefreshTasksResponse) SetBody(v *ListNodeGroupRefreshTasksResponseBody) *ListNodeGroupRefreshTasksResponse {
	s.Body = v
	return s
}

func (s *ListNodeGroupRefreshTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
