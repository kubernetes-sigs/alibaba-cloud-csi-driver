// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNodeGroupRefreshTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNodeGroupRefreshTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNodeGroupRefreshTaskResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNodeGroupRefreshTaskResponseBody) *DescribeNodeGroupRefreshTaskResponse
	GetBody() *DescribeNodeGroupRefreshTaskResponseBody
}

type DescribeNodeGroupRefreshTaskResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNodeGroupRefreshTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNodeGroupRefreshTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNodeGroupRefreshTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeNodeGroupRefreshTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNodeGroupRefreshTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNodeGroupRefreshTaskResponse) GetBody() *DescribeNodeGroupRefreshTaskResponseBody {
	return s.Body
}

func (s *DescribeNodeGroupRefreshTaskResponse) SetHeaders(v map[string]*string) *DescribeNodeGroupRefreshTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeNodeGroupRefreshTaskResponse) SetStatusCode(v int32) *DescribeNodeGroupRefreshTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNodeGroupRefreshTaskResponse) SetBody(v *DescribeNodeGroupRefreshTaskResponseBody) *DescribeNodeGroupRefreshTaskResponse {
	s.Body = v
	return s
}

func (s *DescribeNodeGroupRefreshTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
