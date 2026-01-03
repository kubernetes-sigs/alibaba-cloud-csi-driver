// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSnapshotPackageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSnapshotPackageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSnapshotPackageResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSnapshotPackageResponseBody) *DescribeSnapshotPackageResponse
	GetBody() *DescribeSnapshotPackageResponseBody
}

type DescribeSnapshotPackageResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSnapshotPackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSnapshotPackageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnapshotPackageResponse) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotPackageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSnapshotPackageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSnapshotPackageResponse) GetBody() *DescribeSnapshotPackageResponseBody {
	return s.Body
}

func (s *DescribeSnapshotPackageResponse) SetHeaders(v map[string]*string) *DescribeSnapshotPackageResponse {
	s.Headers = v
	return s
}

func (s *DescribeSnapshotPackageResponse) SetStatusCode(v int32) *DescribeSnapshotPackageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSnapshotPackageResponse) SetBody(v *DescribeSnapshotPackageResponseBody) *DescribeSnapshotPackageResponse {
	s.Body = v
	return s
}

func (s *DescribeSnapshotPackageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
