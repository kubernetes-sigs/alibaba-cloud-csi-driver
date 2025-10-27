// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSnapshotResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSnapshotResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSnapshotResponse
	GetStatusCode() *int32
	SetBody(v *CreateSnapshotResponseBody) *CreateSnapshotResponse
	GetBody() *CreateSnapshotResponseBody
}

type CreateSnapshotResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSnapshotResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSnapshotResponse) GoString() string {
	return s.String()
}

func (s *CreateSnapshotResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSnapshotResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSnapshotResponse) GetBody() *CreateSnapshotResponseBody {
	return s.Body
}

func (s *CreateSnapshotResponse) SetHeaders(v map[string]*string) *CreateSnapshotResponse {
	s.Headers = v
	return s
}

func (s *CreateSnapshotResponse) SetStatusCode(v int32) *CreateSnapshotResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSnapshotResponse) SetBody(v *CreateSnapshotResponseBody) *CreateSnapshotResponse {
	s.Body = v
	return s
}

func (s *CreateSnapshotResponse) Validate() error {
	return dara.Validate(s)
}
