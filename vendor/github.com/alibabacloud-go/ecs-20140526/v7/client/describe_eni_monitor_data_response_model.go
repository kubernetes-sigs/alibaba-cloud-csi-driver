// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEniMonitorDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEniMonitorDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEniMonitorDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEniMonitorDataResponseBody) *DescribeEniMonitorDataResponse
	GetBody() *DescribeEniMonitorDataResponseBody
}

type DescribeEniMonitorDataResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEniMonitorDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEniMonitorDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEniMonitorDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeEniMonitorDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEniMonitorDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEniMonitorDataResponse) GetBody() *DescribeEniMonitorDataResponseBody {
	return s.Body
}

func (s *DescribeEniMonitorDataResponse) SetHeaders(v map[string]*string) *DescribeEniMonitorDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeEniMonitorDataResponse) SetStatusCode(v int32) *DescribeEniMonitorDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEniMonitorDataResponse) SetBody(v *DescribeEniMonitorDataResponseBody) *DescribeEniMonitorDataResponse {
	s.Body = v
	return s
}

func (s *DescribeEniMonitorDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
