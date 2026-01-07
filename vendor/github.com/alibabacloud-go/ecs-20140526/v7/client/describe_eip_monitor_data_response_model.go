// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEipMonitorDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEipMonitorDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEipMonitorDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEipMonitorDataResponseBody) *DescribeEipMonitorDataResponse
	GetBody() *DescribeEipMonitorDataResponseBody
}

type DescribeEipMonitorDataResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEipMonitorDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEipMonitorDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEipMonitorDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeEipMonitorDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEipMonitorDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEipMonitorDataResponse) GetBody() *DescribeEipMonitorDataResponseBody {
	return s.Body
}

func (s *DescribeEipMonitorDataResponse) SetHeaders(v map[string]*string) *DescribeEipMonitorDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeEipMonitorDataResponse) SetStatusCode(v int32) *DescribeEipMonitorDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEipMonitorDataResponse) SetBody(v *DescribeEipMonitorDataResponseBody) *DescribeEipMonitorDataResponse {
	s.Body = v
	return s
}

func (s *DescribeEipMonitorDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
