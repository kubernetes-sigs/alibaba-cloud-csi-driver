// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSnapshotGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSnapshotGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSnapshotGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSnapshotGroupResponseBody) *DeleteSnapshotGroupResponse
	GetBody() *DeleteSnapshotGroupResponseBody
}

type DeleteSnapshotGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSnapshotGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSnapshotGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSnapshotGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSnapshotGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSnapshotGroupResponse) GetBody() *DeleteSnapshotGroupResponseBody {
	return s.Body
}

func (s *DeleteSnapshotGroupResponse) SetHeaders(v map[string]*string) *DeleteSnapshotGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteSnapshotGroupResponse) SetStatusCode(v int32) *DeleteSnapshotGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSnapshotGroupResponse) SetBody(v *DeleteSnapshotGroupResponseBody) *DeleteSnapshotGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteSnapshotGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
