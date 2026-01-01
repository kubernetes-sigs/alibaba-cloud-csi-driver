// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataFlowSubTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDataFlowSubTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDataFlowSubTasksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDataFlowSubTasksResponseBody) *DescribeDataFlowSubTasksResponse
	GetBody() *DescribeDataFlowSubTasksResponseBody
}

type DescribeDataFlowSubTasksResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDataFlowSubTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDataFlowSubTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataFlowSubTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataFlowSubTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDataFlowSubTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDataFlowSubTasksResponse) GetBody() *DescribeDataFlowSubTasksResponseBody {
	return s.Body
}

func (s *DescribeDataFlowSubTasksResponse) SetHeaders(v map[string]*string) *DescribeDataFlowSubTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataFlowSubTasksResponse) SetStatusCode(v int32) *DescribeDataFlowSubTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataFlowSubTasksResponse) SetBody(v *DescribeDataFlowSubTasksResponseBody) *DescribeDataFlowSubTasksResponse {
	s.Body = v
	return s
}

func (s *DescribeDataFlowSubTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
