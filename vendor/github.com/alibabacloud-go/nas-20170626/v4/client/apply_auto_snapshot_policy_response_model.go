// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyAutoSnapshotPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApplyAutoSnapshotPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApplyAutoSnapshotPolicyResponse
	GetStatusCode() *int32
	SetBody(v *ApplyAutoSnapshotPolicyResponseBody) *ApplyAutoSnapshotPolicyResponse
	GetBody() *ApplyAutoSnapshotPolicyResponseBody
}

type ApplyAutoSnapshotPolicyResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyAutoSnapshotPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyAutoSnapshotPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s ApplyAutoSnapshotPolicyResponse) GoString() string {
	return s.String()
}

func (s *ApplyAutoSnapshotPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApplyAutoSnapshotPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApplyAutoSnapshotPolicyResponse) GetBody() *ApplyAutoSnapshotPolicyResponseBody {
	return s.Body
}

func (s *ApplyAutoSnapshotPolicyResponse) SetHeaders(v map[string]*string) *ApplyAutoSnapshotPolicyResponse {
	s.Headers = v
	return s
}

func (s *ApplyAutoSnapshotPolicyResponse) SetStatusCode(v int32) *ApplyAutoSnapshotPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyAutoSnapshotPolicyResponse) SetBody(v *ApplyAutoSnapshotPolicyResponseBody) *ApplyAutoSnapshotPolicyResponse {
	s.Body = v
	return s
}

func (s *ApplyAutoSnapshotPolicyResponse) Validate() error {
	return dara.Validate(s)
}
