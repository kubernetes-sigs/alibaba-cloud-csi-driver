// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnlockSnapshotResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnlockSnapshotResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnlockSnapshotResponse
	GetStatusCode() *int32
	SetBody(v *UnlockSnapshotResponseBody) *UnlockSnapshotResponse
	GetBody() *UnlockSnapshotResponseBody
}

type UnlockSnapshotResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnlockSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnlockSnapshotResponse) String() string {
	return dara.Prettify(s)
}

func (s UnlockSnapshotResponse) GoString() string {
	return s.String()
}

func (s *UnlockSnapshotResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnlockSnapshotResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnlockSnapshotResponse) GetBody() *UnlockSnapshotResponseBody {
	return s.Body
}

func (s *UnlockSnapshotResponse) SetHeaders(v map[string]*string) *UnlockSnapshotResponse {
	s.Headers = v
	return s
}

func (s *UnlockSnapshotResponse) SetStatusCode(v int32) *UnlockSnapshotResponse {
	s.StatusCode = &v
	return s
}

func (s *UnlockSnapshotResponse) SetBody(v *UnlockSnapshotResponseBody) *UnlockSnapshotResponse {
	s.Body = v
	return s
}

func (s *UnlockSnapshotResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
