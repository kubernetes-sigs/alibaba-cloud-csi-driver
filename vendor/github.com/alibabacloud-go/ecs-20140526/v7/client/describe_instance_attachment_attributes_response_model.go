// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceAttachmentAttributesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceAttachmentAttributesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceAttachmentAttributesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceAttachmentAttributesResponseBody) *DescribeInstanceAttachmentAttributesResponse
	GetBody() *DescribeInstanceAttachmentAttributesResponseBody
}

type DescribeInstanceAttachmentAttributesResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceAttachmentAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceAttachmentAttributesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAttachmentAttributesResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttachmentAttributesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceAttachmentAttributesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceAttachmentAttributesResponse) GetBody() *DescribeInstanceAttachmentAttributesResponseBody {
	return s.Body
}

func (s *DescribeInstanceAttachmentAttributesResponse) SetHeaders(v map[string]*string) *DescribeInstanceAttachmentAttributesResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceAttachmentAttributesResponse) SetStatusCode(v int32) *DescribeInstanceAttachmentAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceAttachmentAttributesResponse) SetBody(v *DescribeInstanceAttachmentAttributesResponseBody) *DescribeInstanceAttachmentAttributesResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceAttachmentAttributesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
