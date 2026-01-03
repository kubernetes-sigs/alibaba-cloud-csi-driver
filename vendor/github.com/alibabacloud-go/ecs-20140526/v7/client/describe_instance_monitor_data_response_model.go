// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceMonitorDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceMonitorDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceMonitorDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceMonitorDataResponseBody) *DescribeInstanceMonitorDataResponse
	GetBody() *DescribeInstanceMonitorDataResponseBody
}

type DescribeInstanceMonitorDataResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceMonitorDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceMonitorDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceMonitorDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceMonitorDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceMonitorDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceMonitorDataResponse) GetBody() *DescribeInstanceMonitorDataResponseBody {
	return s.Body
}

func (s *DescribeInstanceMonitorDataResponse) SetHeaders(v map[string]*string) *DescribeInstanceMonitorDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceMonitorDataResponse) SetStatusCode(v int32) *DescribeInstanceMonitorDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceMonitorDataResponse) SetBody(v *DescribeInstanceMonitorDataResponseBody) *DescribeInstanceMonitorDataResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceMonitorDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
