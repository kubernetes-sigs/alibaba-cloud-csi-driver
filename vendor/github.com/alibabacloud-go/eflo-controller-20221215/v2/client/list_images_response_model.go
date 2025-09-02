// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListImagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListImagesResponse
	GetStatusCode() *int32
	SetBody(v *ListImagesResponseBody) *ListImagesResponse
	GetBody() *ListImagesResponseBody
}

type ListImagesResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListImagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListImagesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListImagesResponse) GoString() string {
	return s.String()
}

func (s *ListImagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListImagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListImagesResponse) GetBody() *ListImagesResponseBody {
	return s.Body
}

func (s *ListImagesResponse) SetHeaders(v map[string]*string) *ListImagesResponse {
	s.Headers = v
	return s
}

func (s *ListImagesResponse) SetStatusCode(v int32) *ListImagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListImagesResponse) SetBody(v *ListImagesResponseBody) *ListImagesResponse {
	s.Body = v
	return s
}

func (s *ListImagesResponse) Validate() error {
	return dara.Validate(s)
}
