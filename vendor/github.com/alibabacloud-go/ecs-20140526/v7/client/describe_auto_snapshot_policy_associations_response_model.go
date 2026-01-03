// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoSnapshotPolicyAssociationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAutoSnapshotPolicyAssociationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAutoSnapshotPolicyAssociationsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAutoSnapshotPolicyAssociationsResponseBody) *DescribeAutoSnapshotPolicyAssociationsResponse
	GetBody() *DescribeAutoSnapshotPolicyAssociationsResponseBody
}

type DescribeAutoSnapshotPolicyAssociationsResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAutoSnapshotPolicyAssociationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAutoSnapshotPolicyAssociationsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoSnapshotPolicyAssociationsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAutoSnapshotPolicyAssociationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAutoSnapshotPolicyAssociationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAutoSnapshotPolicyAssociationsResponse) GetBody() *DescribeAutoSnapshotPolicyAssociationsResponseBody {
	return s.Body
}

func (s *DescribeAutoSnapshotPolicyAssociationsResponse) SetHeaders(v map[string]*string) *DescribeAutoSnapshotPolicyAssociationsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAutoSnapshotPolicyAssociationsResponse) SetStatusCode(v int32) *DescribeAutoSnapshotPolicyAssociationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyAssociationsResponse) SetBody(v *DescribeAutoSnapshotPolicyAssociationsResponseBody) *DescribeAutoSnapshotPolicyAssociationsResponse {
	s.Body = v
	return s
}

func (s *DescribeAutoSnapshotPolicyAssociationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
