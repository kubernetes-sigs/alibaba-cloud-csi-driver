// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLockedSnapshotsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLockedSnapshotsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLockedSnapshotsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLockedSnapshotsResponseBody) *DescribeLockedSnapshotsResponse
	GetBody() *DescribeLockedSnapshotsResponseBody
}

type DescribeLockedSnapshotsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLockedSnapshotsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLockedSnapshotsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLockedSnapshotsResponse) GoString() string {
	return s.String()
}

func (s *DescribeLockedSnapshotsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLockedSnapshotsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLockedSnapshotsResponse) GetBody() *DescribeLockedSnapshotsResponseBody {
	return s.Body
}

func (s *DescribeLockedSnapshotsResponse) SetHeaders(v map[string]*string) *DescribeLockedSnapshotsResponse {
	s.Headers = v
	return s
}

func (s *DescribeLockedSnapshotsResponse) SetStatusCode(v int32) *DescribeLockedSnapshotsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLockedSnapshotsResponse) SetBody(v *DescribeLockedSnapshotsResponseBody) *DescribeLockedSnapshotsResponse {
	s.Body = v
	return s
}

func (s *DescribeLockedSnapshotsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
