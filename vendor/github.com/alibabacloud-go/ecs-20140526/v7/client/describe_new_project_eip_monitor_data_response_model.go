// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNewProjectEipMonitorDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNewProjectEipMonitorDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNewProjectEipMonitorDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNewProjectEipMonitorDataResponseBody) *DescribeNewProjectEipMonitorDataResponse
	GetBody() *DescribeNewProjectEipMonitorDataResponseBody
}

type DescribeNewProjectEipMonitorDataResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNewProjectEipMonitorDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNewProjectEipMonitorDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNewProjectEipMonitorDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeNewProjectEipMonitorDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNewProjectEipMonitorDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNewProjectEipMonitorDataResponse) GetBody() *DescribeNewProjectEipMonitorDataResponseBody {
	return s.Body
}

func (s *DescribeNewProjectEipMonitorDataResponse) SetHeaders(v map[string]*string) *DescribeNewProjectEipMonitorDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeNewProjectEipMonitorDataResponse) SetStatusCode(v int32) *DescribeNewProjectEipMonitorDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNewProjectEipMonitorDataResponse) SetBody(v *DescribeNewProjectEipMonitorDataResponseBody) *DescribeNewProjectEipMonitorDataResponse {
	s.Body = v
	return s
}

func (s *DescribeNewProjectEipMonitorDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
