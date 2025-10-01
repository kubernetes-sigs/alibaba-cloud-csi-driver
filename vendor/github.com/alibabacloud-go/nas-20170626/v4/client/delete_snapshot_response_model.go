// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSnapshotResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSnapshotResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSnapshotResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSnapshotResponseBody) *DeleteSnapshotResponse
	GetBody() *DeleteSnapshotResponseBody
}

type DeleteSnapshotResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSnapshotResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSnapshotResponse) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSnapshotResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSnapshotResponse) GetBody() *DeleteSnapshotResponseBody {
	return s.Body
}

func (s *DeleteSnapshotResponse) SetHeaders(v map[string]*string) *DeleteSnapshotResponse {
	s.Headers = v
	return s
}

func (s *DeleteSnapshotResponse) SetStatusCode(v int32) *DeleteSnapshotResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSnapshotResponse) SetBody(v *DeleteSnapshotResponseBody) *DeleteSnapshotResponse {
	s.Body = v
	return s
}

func (s *DeleteSnapshotResponse) Validate() error {
	return dara.Validate(s)
}
