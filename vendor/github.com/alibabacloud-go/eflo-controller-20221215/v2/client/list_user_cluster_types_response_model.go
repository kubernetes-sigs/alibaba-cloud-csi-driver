// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserClusterTypesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserClusterTypesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserClusterTypesResponse
	GetStatusCode() *int32
	SetBody(v *ListUserClusterTypesResponseBody) *ListUserClusterTypesResponse
	GetBody() *ListUserClusterTypesResponseBody
}

type ListUserClusterTypesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserClusterTypesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserClusterTypesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserClusterTypesResponse) GoString() string {
	return s.String()
}

func (s *ListUserClusterTypesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserClusterTypesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserClusterTypesResponse) GetBody() *ListUserClusterTypesResponseBody {
	return s.Body
}

func (s *ListUserClusterTypesResponse) SetHeaders(v map[string]*string) *ListUserClusterTypesResponse {
	s.Headers = v
	return s
}

func (s *ListUserClusterTypesResponse) SetStatusCode(v int32) *ListUserClusterTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserClusterTypesResponse) SetBody(v *ListUserClusterTypesResponseBody) *ListUserClusterTypesResponse {
	s.Body = v
	return s
}

func (s *ListUserClusterTypesResponse) Validate() error {
	return dara.Validate(s)
}
