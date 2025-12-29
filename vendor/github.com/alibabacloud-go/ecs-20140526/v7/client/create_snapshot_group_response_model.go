// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSnapshotGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSnapshotGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSnapshotGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateSnapshotGroupResponseBody) *CreateSnapshotGroupResponse
	GetBody() *CreateSnapshotGroupResponseBody
}

type CreateSnapshotGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSnapshotGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSnapshotGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSnapshotGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateSnapshotGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSnapshotGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSnapshotGroupResponse) GetBody() *CreateSnapshotGroupResponseBody {
	return s.Body
}

func (s *CreateSnapshotGroupResponse) SetHeaders(v map[string]*string) *CreateSnapshotGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateSnapshotGroupResponse) SetStatusCode(v int32) *CreateSnapshotGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSnapshotGroupResponse) SetBody(v *CreateSnapshotGroupResponseBody) *CreateSnapshotGroupResponse {
	s.Body = v
	return s
}

func (s *CreateSnapshotGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
