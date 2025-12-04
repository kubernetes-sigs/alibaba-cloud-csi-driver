// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProtocolMountTargetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeProtocolMountTargetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeProtocolMountTargetResponse
	GetStatusCode() *int32
	SetBody(v *DescribeProtocolMountTargetResponseBody) *DescribeProtocolMountTargetResponse
	GetBody() *DescribeProtocolMountTargetResponseBody
}

type DescribeProtocolMountTargetResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeProtocolMountTargetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeProtocolMountTargetResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeProtocolMountTargetResponse) GoString() string {
	return s.String()
}

func (s *DescribeProtocolMountTargetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeProtocolMountTargetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeProtocolMountTargetResponse) GetBody() *DescribeProtocolMountTargetResponseBody {
	return s.Body
}

func (s *DescribeProtocolMountTargetResponse) SetHeaders(v map[string]*string) *DescribeProtocolMountTargetResponse {
	s.Headers = v
	return s
}

func (s *DescribeProtocolMountTargetResponse) SetStatusCode(v int32) *DescribeProtocolMountTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeProtocolMountTargetResponse) SetBody(v *DescribeProtocolMountTargetResponseBody) *DescribeProtocolMountTargetResponse {
	s.Body = v
	return s
}

func (s *DescribeProtocolMountTargetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
