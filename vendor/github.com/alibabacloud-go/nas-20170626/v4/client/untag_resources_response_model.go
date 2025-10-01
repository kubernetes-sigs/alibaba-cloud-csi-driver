// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUntagResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UntagResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UntagResourcesResponse
	GetStatusCode() *int32
	SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse
	GetBody() *UntagResourcesResponseBody
}

type UntagResourcesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UntagResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s UntagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UntagResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UntagResourcesResponse) GetBody() *UntagResourcesResponseBody {
	return s.Body
}

func (s *UntagResourcesResponse) SetHeaders(v map[string]*string) *UntagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UntagResourcesResponse) SetStatusCode(v int32) *UntagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
	s.Body = v
	return s
}

func (s *UntagResourcesResponse) Validate() error {
	return dara.Validate(s)
}
