// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeForwardTableEntriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeForwardTableEntriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeForwardTableEntriesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeForwardTableEntriesResponseBody) *DescribeForwardTableEntriesResponse
	GetBody() *DescribeForwardTableEntriesResponseBody
}

type DescribeForwardTableEntriesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeForwardTableEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeForwardTableEntriesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeForwardTableEntriesResponse) GoString() string {
	return s.String()
}

func (s *DescribeForwardTableEntriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeForwardTableEntriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeForwardTableEntriesResponse) GetBody() *DescribeForwardTableEntriesResponseBody {
	return s.Body
}

func (s *DescribeForwardTableEntriesResponse) SetHeaders(v map[string]*string) *DescribeForwardTableEntriesResponse {
	s.Headers = v
	return s
}

func (s *DescribeForwardTableEntriesResponse) SetStatusCode(v int32) *DescribeForwardTableEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeForwardTableEntriesResponse) SetBody(v *DescribeForwardTableEntriesResponseBody) *DescribeForwardTableEntriesResponse {
	s.Body = v
	return s
}

func (s *DescribeForwardTableEntriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
