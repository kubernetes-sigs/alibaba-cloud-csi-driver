// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImagePipelineExecutionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeImagePipelineExecutionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeImagePipelineExecutionsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeImagePipelineExecutionsResponseBody) *DescribeImagePipelineExecutionsResponse
	GetBody() *DescribeImagePipelineExecutionsResponseBody
}

type DescribeImagePipelineExecutionsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeImagePipelineExecutionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeImagePipelineExecutionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeImagePipelineExecutionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeImagePipelineExecutionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeImagePipelineExecutionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeImagePipelineExecutionsResponse) GetBody() *DescribeImagePipelineExecutionsResponseBody {
	return s.Body
}

func (s *DescribeImagePipelineExecutionsResponse) SetHeaders(v map[string]*string) *DescribeImagePipelineExecutionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeImagePipelineExecutionsResponse) SetStatusCode(v int32) *DescribeImagePipelineExecutionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImagePipelineExecutionsResponse) SetBody(v *DescribeImagePipelineExecutionsResponseBody) *DescribeImagePipelineExecutionsResponse {
	s.Body = v
	return s
}

func (s *DescribeImagePipelineExecutionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
