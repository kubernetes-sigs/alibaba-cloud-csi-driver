// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSnapshotsUsageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSnapshotsUsageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSnapshotsUsageResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSnapshotsUsageResponseBody) *DescribeSnapshotsUsageResponse
	GetBody() *DescribeSnapshotsUsageResponseBody
}

type DescribeSnapshotsUsageResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSnapshotsUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSnapshotsUsageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnapshotsUsageResponse) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsUsageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSnapshotsUsageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSnapshotsUsageResponse) GetBody() *DescribeSnapshotsUsageResponseBody {
	return s.Body
}

func (s *DescribeSnapshotsUsageResponse) SetHeaders(v map[string]*string) *DescribeSnapshotsUsageResponse {
	s.Headers = v
	return s
}

func (s *DescribeSnapshotsUsageResponse) SetStatusCode(v int32) *DescribeSnapshotsUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSnapshotsUsageResponse) SetBody(v *DescribeSnapshotsUsageResponseBody) *DescribeSnapshotsUsageResponse {
	s.Body = v
	return s
}

func (s *DescribeSnapshotsUsageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
