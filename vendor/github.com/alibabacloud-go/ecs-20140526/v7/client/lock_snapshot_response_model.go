// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLockSnapshotResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *LockSnapshotResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *LockSnapshotResponse
	GetStatusCode() *int32
	SetBody(v *LockSnapshotResponseBody) *LockSnapshotResponse
	GetBody() *LockSnapshotResponseBody
}

type LockSnapshotResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LockSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LockSnapshotResponse) String() string {
	return dara.Prettify(s)
}

func (s LockSnapshotResponse) GoString() string {
	return s.String()
}

func (s *LockSnapshotResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *LockSnapshotResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *LockSnapshotResponse) GetBody() *LockSnapshotResponseBody {
	return s.Body
}

func (s *LockSnapshotResponse) SetHeaders(v map[string]*string) *LockSnapshotResponse {
	s.Headers = v
	return s
}

func (s *LockSnapshotResponse) SetStatusCode(v int32) *LockSnapshotResponse {
	s.StatusCode = &v
	return s
}

func (s *LockSnapshotResponse) SetBody(v *LockSnapshotResponseBody) *LockSnapshotResponse {
	s.Body = v
	return s
}

func (s *LockSnapshotResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
