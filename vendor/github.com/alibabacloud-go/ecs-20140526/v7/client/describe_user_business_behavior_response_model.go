// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserBusinessBehaviorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserBusinessBehaviorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserBusinessBehaviorResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserBusinessBehaviorResponseBody) *DescribeUserBusinessBehaviorResponse
	GetBody() *DescribeUserBusinessBehaviorResponseBody
}

type DescribeUserBusinessBehaviorResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserBusinessBehaviorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserBusinessBehaviorResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserBusinessBehaviorResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserBusinessBehaviorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserBusinessBehaviorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserBusinessBehaviorResponse) GetBody() *DescribeUserBusinessBehaviorResponseBody {
	return s.Body
}

func (s *DescribeUserBusinessBehaviorResponse) SetHeaders(v map[string]*string) *DescribeUserBusinessBehaviorResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserBusinessBehaviorResponse) SetStatusCode(v int32) *DescribeUserBusinessBehaviorResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserBusinessBehaviorResponse) SetBody(v *DescribeUserBusinessBehaviorResponseBody) *DescribeUserBusinessBehaviorResponse {
	s.Body = v
	return s
}

func (s *DescribeUserBusinessBehaviorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
