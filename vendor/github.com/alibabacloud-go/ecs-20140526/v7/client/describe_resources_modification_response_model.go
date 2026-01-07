// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourcesModificationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeResourcesModificationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeResourcesModificationResponse
	GetStatusCode() *int32
	SetBody(v *DescribeResourcesModificationResponseBody) *DescribeResourcesModificationResponse
	GetBody() *DescribeResourcesModificationResponseBody
}

type DescribeResourcesModificationResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeResourcesModificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeResourcesModificationResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourcesModificationResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourcesModificationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeResourcesModificationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeResourcesModificationResponse) GetBody() *DescribeResourcesModificationResponseBody {
	return s.Body
}

func (s *DescribeResourcesModificationResponse) SetHeaders(v map[string]*string) *DescribeResourcesModificationResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourcesModificationResponse) SetStatusCode(v int32) *DescribeResourcesModificationResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResourcesModificationResponse) SetBody(v *DescribeResourcesModificationResponseBody) *DescribeResourcesModificationResponse {
	s.Body = v
	return s
}

func (s *DescribeResourcesModificationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
