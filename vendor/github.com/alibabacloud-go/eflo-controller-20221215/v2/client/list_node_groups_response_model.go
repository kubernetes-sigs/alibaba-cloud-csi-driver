// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodeGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNodeGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNodeGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ListNodeGroupsResponseBody) *ListNodeGroupsResponse
	GetBody() *ListNodeGroupsResponseBody
}

type ListNodeGroupsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNodeGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNodeGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNodeGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListNodeGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNodeGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNodeGroupsResponse) GetBody() *ListNodeGroupsResponseBody {
	return s.Body
}

func (s *ListNodeGroupsResponse) SetHeaders(v map[string]*string) *ListNodeGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListNodeGroupsResponse) SetStatusCode(v int32) *ListNodeGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNodeGroupsResponse) SetBody(v *ListNodeGroupsResponseBody) *ListNodeGroupsResponse {
	s.Body = v
	return s
}

func (s *ListNodeGroupsResponse) Validate() error {
	return dara.Validate(s)
}
