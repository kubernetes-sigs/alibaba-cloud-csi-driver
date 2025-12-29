// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoSnapshotPolicyExResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAutoSnapshotPolicyExResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAutoSnapshotPolicyExResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAutoSnapshotPolicyExResponseBody) *DescribeAutoSnapshotPolicyExResponse
	GetBody() *DescribeAutoSnapshotPolicyExResponseBody
}

type DescribeAutoSnapshotPolicyExResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAutoSnapshotPolicyExResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAutoSnapshotPolicyExResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoSnapshotPolicyExResponse) GoString() string {
	return s.String()
}

func (s *DescribeAutoSnapshotPolicyExResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAutoSnapshotPolicyExResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAutoSnapshotPolicyExResponse) GetBody() *DescribeAutoSnapshotPolicyExResponseBody {
	return s.Body
}

func (s *DescribeAutoSnapshotPolicyExResponse) SetHeaders(v map[string]*string) *DescribeAutoSnapshotPolicyExResponse {
	s.Headers = v
	return s
}

func (s *DescribeAutoSnapshotPolicyExResponse) SetStatusCode(v int32) *DescribeAutoSnapshotPolicyExResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyExResponse) SetBody(v *DescribeAutoSnapshotPolicyExResponseBody) *DescribeAutoSnapshotPolicyExResponse {
	s.Body = v
	return s
}

func (s *DescribeAutoSnapshotPolicyExResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
