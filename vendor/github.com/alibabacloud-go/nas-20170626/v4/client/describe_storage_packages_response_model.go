// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStoragePackagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeStoragePackagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeStoragePackagesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeStoragePackagesResponseBody) *DescribeStoragePackagesResponse
	GetBody() *DescribeStoragePackagesResponseBody
}

type DescribeStoragePackagesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeStoragePackagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeStoragePackagesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeStoragePackagesResponse) GoString() string {
	return s.String()
}

func (s *DescribeStoragePackagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeStoragePackagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeStoragePackagesResponse) GetBody() *DescribeStoragePackagesResponseBody {
	return s.Body
}

func (s *DescribeStoragePackagesResponse) SetHeaders(v map[string]*string) *DescribeStoragePackagesResponse {
	s.Headers = v
	return s
}

func (s *DescribeStoragePackagesResponse) SetStatusCode(v int32) *DescribeStoragePackagesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeStoragePackagesResponse) SetBody(v *DescribeStoragePackagesResponseBody) *DescribeStoragePackagesResponse {
	s.Body = v
	return s
}

func (s *DescribeStoragePackagesResponse) Validate() error {
	return dara.Validate(s)
}
