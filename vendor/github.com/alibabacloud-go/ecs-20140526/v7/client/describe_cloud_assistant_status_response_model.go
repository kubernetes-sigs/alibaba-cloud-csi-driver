// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudAssistantStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCloudAssistantStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCloudAssistantStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCloudAssistantStatusResponseBody) *DescribeCloudAssistantStatusResponse
	GetBody() *DescribeCloudAssistantStatusResponseBody
}

type DescribeCloudAssistantStatusResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudAssistantStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudAssistantStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudAssistantStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudAssistantStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCloudAssistantStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCloudAssistantStatusResponse) GetBody() *DescribeCloudAssistantStatusResponseBody {
	return s.Body
}

func (s *DescribeCloudAssistantStatusResponse) SetHeaders(v map[string]*string) *DescribeCloudAssistantStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudAssistantStatusResponse) SetStatusCode(v int32) *DescribeCloudAssistantStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudAssistantStatusResponse) SetBody(v *DescribeCloudAssistantStatusResponseBody) *DescribeCloudAssistantStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeCloudAssistantStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
