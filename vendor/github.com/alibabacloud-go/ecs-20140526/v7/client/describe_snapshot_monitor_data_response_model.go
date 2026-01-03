// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSnapshotMonitorDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSnapshotMonitorDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSnapshotMonitorDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSnapshotMonitorDataResponseBody) *DescribeSnapshotMonitorDataResponse
	GetBody() *DescribeSnapshotMonitorDataResponseBody
}

type DescribeSnapshotMonitorDataResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSnapshotMonitorDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSnapshotMonitorDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnapshotMonitorDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotMonitorDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSnapshotMonitorDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSnapshotMonitorDataResponse) GetBody() *DescribeSnapshotMonitorDataResponseBody {
	return s.Body
}

func (s *DescribeSnapshotMonitorDataResponse) SetHeaders(v map[string]*string) *DescribeSnapshotMonitorDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeSnapshotMonitorDataResponse) SetStatusCode(v int32) *DescribeSnapshotMonitorDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSnapshotMonitorDataResponse) SetBody(v *DescribeSnapshotMonitorDataResponseBody) *DescribeSnapshotMonitorDataResponse {
	s.Body = v
	return s
}

func (s *DescribeSnapshotMonitorDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
