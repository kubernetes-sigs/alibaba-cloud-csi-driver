// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAutoSnapshotPolicyExResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAutoSnapshotPolicyExResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAutoSnapshotPolicyExResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAutoSnapshotPolicyExResponseBody) *ModifyAutoSnapshotPolicyExResponse
	GetBody() *ModifyAutoSnapshotPolicyExResponseBody
}

type ModifyAutoSnapshotPolicyExResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAutoSnapshotPolicyExResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAutoSnapshotPolicyExResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAutoSnapshotPolicyExResponse) GoString() string {
	return s.String()
}

func (s *ModifyAutoSnapshotPolicyExResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAutoSnapshotPolicyExResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAutoSnapshotPolicyExResponse) GetBody() *ModifyAutoSnapshotPolicyExResponseBody {
	return s.Body
}

func (s *ModifyAutoSnapshotPolicyExResponse) SetHeaders(v map[string]*string) *ModifyAutoSnapshotPolicyExResponse {
	s.Headers = v
	return s
}

func (s *ModifyAutoSnapshotPolicyExResponse) SetStatusCode(v int32) *ModifyAutoSnapshotPolicyExResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAutoSnapshotPolicyExResponse) SetBody(v *ModifyAutoSnapshotPolicyExResponseBody) *ModifyAutoSnapshotPolicyExResponse {
	s.Body = v
	return s
}

func (s *ModifyAutoSnapshotPolicyExResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
