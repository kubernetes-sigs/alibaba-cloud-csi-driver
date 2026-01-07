// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySnapshotAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySnapshotAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySnapshotAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifySnapshotAttributeResponseBody) *ModifySnapshotAttributeResponse
	GetBody() *ModifySnapshotAttributeResponseBody
}

type ModifySnapshotAttributeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySnapshotAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySnapshotAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySnapshotAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifySnapshotAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySnapshotAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySnapshotAttributeResponse) GetBody() *ModifySnapshotAttributeResponseBody {
	return s.Body
}

func (s *ModifySnapshotAttributeResponse) SetHeaders(v map[string]*string) *ModifySnapshotAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifySnapshotAttributeResponse) SetStatusCode(v int32) *ModifySnapshotAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySnapshotAttributeResponse) SetBody(v *ModifySnapshotAttributeResponseBody) *ModifySnapshotAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifySnapshotAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
