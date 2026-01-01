// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFileSystemStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFileSystemStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFileSystemStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFileSystemStatisticsResponseBody) *DescribeFileSystemStatisticsResponse
	GetBody() *DescribeFileSystemStatisticsResponseBody
}

type DescribeFileSystemStatisticsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFileSystemStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFileSystemStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileSystemStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFileSystemStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFileSystemStatisticsResponse) GetBody() *DescribeFileSystemStatisticsResponseBody {
	return s.Body
}

func (s *DescribeFileSystemStatisticsResponse) SetHeaders(v map[string]*string) *DescribeFileSystemStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeFileSystemStatisticsResponse) SetStatusCode(v int32) *DescribeFileSystemStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponse) SetBody(v *DescribeFileSystemStatisticsResponseBody) *DescribeFileSystemStatisticsResponse {
	s.Body = v
	return s
}

func (s *DescribeFileSystemStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
