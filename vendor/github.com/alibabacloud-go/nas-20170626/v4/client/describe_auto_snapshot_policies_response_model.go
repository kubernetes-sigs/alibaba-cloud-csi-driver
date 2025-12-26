// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoSnapshotPoliciesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAutoSnapshotPoliciesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAutoSnapshotPoliciesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAutoSnapshotPoliciesResponseBody) *DescribeAutoSnapshotPoliciesResponse
	GetBody() *DescribeAutoSnapshotPoliciesResponseBody
}

type DescribeAutoSnapshotPoliciesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAutoSnapshotPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAutoSnapshotPoliciesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoSnapshotPoliciesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAutoSnapshotPoliciesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAutoSnapshotPoliciesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAutoSnapshotPoliciesResponse) GetBody() *DescribeAutoSnapshotPoliciesResponseBody {
	return s.Body
}

func (s *DescribeAutoSnapshotPoliciesResponse) SetHeaders(v map[string]*string) *DescribeAutoSnapshotPoliciesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAutoSnapshotPoliciesResponse) SetStatusCode(v int32) *DescribeAutoSnapshotPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesResponse) SetBody(v *DescribeAutoSnapshotPoliciesResponseBody) *DescribeAutoSnapshotPoliciesResponse {
	s.Body = v
	return s
}

func (s *DescribeAutoSnapshotPoliciesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
