// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNodeGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNodeGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNodeGroupResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNodeGroupResponseBody) *DescribeNodeGroupResponse
	GetBody() *DescribeNodeGroupResponseBody
}

type DescribeNodeGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNodeGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNodeGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNodeGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeNodeGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNodeGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNodeGroupResponse) GetBody() *DescribeNodeGroupResponseBody {
	return s.Body
}

func (s *DescribeNodeGroupResponse) SetHeaders(v map[string]*string) *DescribeNodeGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeNodeGroupResponse) SetStatusCode(v int32) *DescribeNodeGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNodeGroupResponse) SetBody(v *DescribeNodeGroupResponseBody) *DescribeNodeGroupResponse {
	s.Body = v
	return s
}

func (s *DescribeNodeGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
