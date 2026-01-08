// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataFlowsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDataFlowsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDataFlowsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDataFlowsResponseBody) *DescribeDataFlowsResponse
	GetBody() *DescribeDataFlowsResponseBody
}

type DescribeDataFlowsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDataFlowsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDataFlowsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataFlowsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataFlowsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDataFlowsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDataFlowsResponse) GetBody() *DescribeDataFlowsResponseBody {
	return s.Body
}

func (s *DescribeDataFlowsResponse) SetHeaders(v map[string]*string) *DescribeDataFlowsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataFlowsResponse) SetStatusCode(v int32) *DescribeDataFlowsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataFlowsResponse) SetBody(v *DescribeDataFlowsResponseBody) *DescribeDataFlowsResponse {
	s.Body = v
	return s
}

func (s *DescribeDataFlowsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
