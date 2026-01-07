// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSnapshotsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSnapshotsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSnapshotsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSnapshotsResponseBody) *DescribeSnapshotsResponse
	GetBody() *DescribeSnapshotsResponseBody
}

type DescribeSnapshotsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSnapshotsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSnapshotsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnapshotsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSnapshotsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSnapshotsResponse) GetBody() *DescribeSnapshotsResponseBody {
	return s.Body
}

func (s *DescribeSnapshotsResponse) SetHeaders(v map[string]*string) *DescribeSnapshotsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSnapshotsResponse) SetStatusCode(v int32) *DescribeSnapshotsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSnapshotsResponse) SetBody(v *DescribeSnapshotsResponseBody) *DescribeSnapshotsResponse {
	s.Body = v
	return s
}

func (s *DescribeSnapshotsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
