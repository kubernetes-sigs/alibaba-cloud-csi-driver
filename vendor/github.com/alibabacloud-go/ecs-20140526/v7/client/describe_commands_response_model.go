// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCommandsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCommandsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCommandsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCommandsResponseBody) *DescribeCommandsResponse
	GetBody() *DescribeCommandsResponseBody
}

type DescribeCommandsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCommandsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCommandsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommandsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCommandsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCommandsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCommandsResponse) GetBody() *DescribeCommandsResponseBody {
	return s.Body
}

func (s *DescribeCommandsResponse) SetHeaders(v map[string]*string) *DescribeCommandsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCommandsResponse) SetStatusCode(v int32) *DescribeCommandsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCommandsResponse) SetBody(v *DescribeCommandsResponseBody) *DescribeCommandsResponse {
	s.Body = v
	return s
}

func (s *DescribeCommandsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
