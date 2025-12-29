// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopySnapshotResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CopySnapshotResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CopySnapshotResponse
	GetStatusCode() *int32
	SetBody(v *CopySnapshotResponseBody) *CopySnapshotResponse
	GetBody() *CopySnapshotResponseBody
}

type CopySnapshotResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CopySnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CopySnapshotResponse) String() string {
	return dara.Prettify(s)
}

func (s CopySnapshotResponse) GoString() string {
	return s.String()
}

func (s *CopySnapshotResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CopySnapshotResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CopySnapshotResponse) GetBody() *CopySnapshotResponseBody {
	return s.Body
}

func (s *CopySnapshotResponse) SetHeaders(v map[string]*string) *CopySnapshotResponse {
	s.Headers = v
	return s
}

func (s *CopySnapshotResponse) SetStatusCode(v int32) *CopySnapshotResponse {
	s.StatusCode = &v
	return s
}

func (s *CopySnapshotResponse) SetBody(v *CopySnapshotResponseBody) *CopySnapshotResponse {
	s.Body = v
	return s
}

func (s *CopySnapshotResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
