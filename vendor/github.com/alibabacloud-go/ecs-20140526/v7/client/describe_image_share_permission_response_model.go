// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageSharePermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeImageSharePermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeImageSharePermissionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeImageSharePermissionResponseBody) *DescribeImageSharePermissionResponse
	GetBody() *DescribeImageSharePermissionResponseBody
}

type DescribeImageSharePermissionResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeImageSharePermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeImageSharePermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageSharePermissionResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageSharePermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeImageSharePermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeImageSharePermissionResponse) GetBody() *DescribeImageSharePermissionResponseBody {
	return s.Body
}

func (s *DescribeImageSharePermissionResponse) SetHeaders(v map[string]*string) *DescribeImageSharePermissionResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageSharePermissionResponse) SetStatusCode(v int32) *DescribeImageSharePermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImageSharePermissionResponse) SetBody(v *DescribeImageSharePermissionResponseBody) *DescribeImageSharePermissionResponse {
	s.Body = v
	return s
}

func (s *DescribeImageSharePermissionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
