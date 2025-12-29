// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePortRangeListEntriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePortRangeListEntriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePortRangeListEntriesResponse
	GetStatusCode() *int32
	SetBody(v *DescribePortRangeListEntriesResponseBody) *DescribePortRangeListEntriesResponse
	GetBody() *DescribePortRangeListEntriesResponseBody
}

type DescribePortRangeListEntriesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePortRangeListEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePortRangeListEntriesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePortRangeListEntriesResponse) GoString() string {
	return s.String()
}

func (s *DescribePortRangeListEntriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePortRangeListEntriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePortRangeListEntriesResponse) GetBody() *DescribePortRangeListEntriesResponseBody {
	return s.Body
}

func (s *DescribePortRangeListEntriesResponse) SetHeaders(v map[string]*string) *DescribePortRangeListEntriesResponse {
	s.Headers = v
	return s
}

func (s *DescribePortRangeListEntriesResponse) SetStatusCode(v int32) *DescribePortRangeListEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePortRangeListEntriesResponse) SetBody(v *DescribePortRangeListEntriesResponseBody) *DescribePortRangeListEntriesResponse {
	s.Body = v
	return s
}

func (s *DescribePortRangeListEntriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
