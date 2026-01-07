// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySnapshotGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySnapshotGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySnapshotGroupResponse
	GetStatusCode() *int32
	SetBody(v *ModifySnapshotGroupResponseBody) *ModifySnapshotGroupResponse
	GetBody() *ModifySnapshotGroupResponseBody
}

type ModifySnapshotGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySnapshotGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySnapshotGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySnapshotGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifySnapshotGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySnapshotGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySnapshotGroupResponse) GetBody() *ModifySnapshotGroupResponseBody {
	return s.Body
}

func (s *ModifySnapshotGroupResponse) SetHeaders(v map[string]*string) *ModifySnapshotGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifySnapshotGroupResponse) SetStatusCode(v int32) *ModifySnapshotGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySnapshotGroupResponse) SetBody(v *ModifySnapshotGroupResponseBody) *ModifySnapshotGroupResponse {
	s.Body = v
	return s
}

func (s *ModifySnapshotGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
