// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSnapshotLinksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSnapshotLinksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSnapshotLinksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSnapshotLinksResponseBody) *DescribeSnapshotLinksResponse
	GetBody() *DescribeSnapshotLinksResponseBody
}

type DescribeSnapshotLinksResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSnapshotLinksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSnapshotLinksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnapshotLinksResponse) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotLinksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSnapshotLinksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSnapshotLinksResponse) GetBody() *DescribeSnapshotLinksResponseBody {
	return s.Body
}

func (s *DescribeSnapshotLinksResponse) SetHeaders(v map[string]*string) *DescribeSnapshotLinksResponse {
	s.Headers = v
	return s
}

func (s *DescribeSnapshotLinksResponse) SetStatusCode(v int32) *DescribeSnapshotLinksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSnapshotLinksResponse) SetBody(v *DescribeSnapshotLinksResponseBody) *DescribeSnapshotLinksResponse {
	s.Body = v
	return s
}

func (s *DescribeSnapshotLinksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
