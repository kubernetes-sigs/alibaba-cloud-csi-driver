// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFilesystemsVscAttachInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFilesystemsVscAttachInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFilesystemsVscAttachInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFilesystemsVscAttachInfoResponseBody) *DescribeFilesystemsVscAttachInfoResponse
	GetBody() *DescribeFilesystemsVscAttachInfoResponseBody
}

type DescribeFilesystemsVscAttachInfoResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFilesystemsVscAttachInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFilesystemsVscAttachInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFilesystemsVscAttachInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeFilesystemsVscAttachInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFilesystemsVscAttachInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFilesystemsVscAttachInfoResponse) GetBody() *DescribeFilesystemsVscAttachInfoResponseBody {
	return s.Body
}

func (s *DescribeFilesystemsVscAttachInfoResponse) SetHeaders(v map[string]*string) *DescribeFilesystemsVscAttachInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeFilesystemsVscAttachInfoResponse) SetStatusCode(v int32) *DescribeFilesystemsVscAttachInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFilesystemsVscAttachInfoResponse) SetBody(v *DescribeFilesystemsVscAttachInfoResponseBody) *DescribeFilesystemsVscAttachInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeFilesystemsVscAttachInfoResponse) Validate() error {
	return dara.Validate(s)
}
