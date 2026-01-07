// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoSnapshotTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAutoSnapshotTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAutoSnapshotTasksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAutoSnapshotTasksResponseBody) *DescribeAutoSnapshotTasksResponse
	GetBody() *DescribeAutoSnapshotTasksResponseBody
}

type DescribeAutoSnapshotTasksResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAutoSnapshotTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAutoSnapshotTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoSnapshotTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeAutoSnapshotTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAutoSnapshotTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAutoSnapshotTasksResponse) GetBody() *DescribeAutoSnapshotTasksResponseBody {
	return s.Body
}

func (s *DescribeAutoSnapshotTasksResponse) SetHeaders(v map[string]*string) *DescribeAutoSnapshotTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeAutoSnapshotTasksResponse) SetStatusCode(v int32) *DescribeAutoSnapshotTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAutoSnapshotTasksResponse) SetBody(v *DescribeAutoSnapshotTasksResponseBody) *DescribeAutoSnapshotTasksResponse {
	s.Body = v
	return s
}

func (s *DescribeAutoSnapshotTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
