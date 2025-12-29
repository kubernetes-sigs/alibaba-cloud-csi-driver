// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceHistoryEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceHistoryEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceHistoryEventsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceHistoryEventsResponseBody) *DescribeInstanceHistoryEventsResponse
	GetBody() *DescribeInstanceHistoryEventsResponseBody
}

type DescribeInstanceHistoryEventsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceHistoryEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceHistoryEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceHistoryEventsResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceHistoryEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceHistoryEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceHistoryEventsResponse) GetBody() *DescribeInstanceHistoryEventsResponseBody {
	return s.Body
}

func (s *DescribeInstanceHistoryEventsResponse) SetHeaders(v map[string]*string) *DescribeInstanceHistoryEventsResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceHistoryEventsResponse) SetStatusCode(v int32) *DescribeInstanceHistoryEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceHistoryEventsResponse) SetBody(v *DescribeInstanceHistoryEventsResponseBody) *DescribeInstanceHistoryEventsResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceHistoryEventsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
