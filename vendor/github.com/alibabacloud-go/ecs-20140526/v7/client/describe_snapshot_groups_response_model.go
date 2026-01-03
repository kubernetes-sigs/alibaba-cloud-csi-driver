// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSnapshotGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSnapshotGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSnapshotGroupsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSnapshotGroupsResponseBody) *DescribeSnapshotGroupsResponse
	GetBody() *DescribeSnapshotGroupsResponseBody
}

type DescribeSnapshotGroupsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSnapshotGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSnapshotGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnapshotGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSnapshotGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSnapshotGroupsResponse) GetBody() *DescribeSnapshotGroupsResponseBody {
	return s.Body
}

func (s *DescribeSnapshotGroupsResponse) SetHeaders(v map[string]*string) *DescribeSnapshotGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSnapshotGroupsResponse) SetStatusCode(v int32) *DescribeSnapshotGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSnapshotGroupsResponse) SetBody(v *DescribeSnapshotGroupsResponseBody) *DescribeSnapshotGroupsResponse {
	s.Body = v
	return s
}

func (s *DescribeSnapshotGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
