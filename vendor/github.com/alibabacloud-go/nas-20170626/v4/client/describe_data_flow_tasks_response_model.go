// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataFlowTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDataFlowTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDataFlowTasksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDataFlowTasksResponseBody) *DescribeDataFlowTasksResponse
	GetBody() *DescribeDataFlowTasksResponseBody
}

type DescribeDataFlowTasksResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDataFlowTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDataFlowTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataFlowTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataFlowTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDataFlowTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDataFlowTasksResponse) GetBody() *DescribeDataFlowTasksResponseBody {
	return s.Body
}

func (s *DescribeDataFlowTasksResponse) SetHeaders(v map[string]*string) *DescribeDataFlowTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataFlowTasksResponse) SetStatusCode(v int32) *DescribeDataFlowTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataFlowTasksResponse) SetBody(v *DescribeDataFlowTasksResponseBody) *DescribeDataFlowTasksResponse {
	s.Body = v
	return s
}

func (s *DescribeDataFlowTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
