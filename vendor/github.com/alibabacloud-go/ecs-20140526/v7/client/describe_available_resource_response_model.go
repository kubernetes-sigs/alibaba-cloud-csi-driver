// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAvailableResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAvailableResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAvailableResourceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAvailableResourceResponseBody) *DescribeAvailableResourceResponse
	GetBody() *DescribeAvailableResourceResponseBody
}

type DescribeAvailableResourceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAvailableResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAvailableResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAvailableResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAvailableResourceResponse) GetBody() *DescribeAvailableResourceResponseBody {
	return s.Body
}

func (s *DescribeAvailableResourceResponse) SetHeaders(v map[string]*string) *DescribeAvailableResourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeAvailableResourceResponse) SetStatusCode(v int32) *DescribeAvailableResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAvailableResourceResponse) SetBody(v *DescribeAvailableResourceResponseBody) *DescribeAvailableResourceResponse {
	s.Body = v
	return s
}

func (s *DescribeAvailableResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
