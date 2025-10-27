// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAutoSnapshotPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAutoSnapshotPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAutoSnapshotPolicyResponse
	GetStatusCode() *int32
	SetBody(v *CreateAutoSnapshotPolicyResponseBody) *CreateAutoSnapshotPolicyResponse
	GetBody() *CreateAutoSnapshotPolicyResponseBody
}

type CreateAutoSnapshotPolicyResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAutoSnapshotPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAutoSnapshotPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoSnapshotPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateAutoSnapshotPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAutoSnapshotPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAutoSnapshotPolicyResponse) GetBody() *CreateAutoSnapshotPolicyResponseBody {
	return s.Body
}

func (s *CreateAutoSnapshotPolicyResponse) SetHeaders(v map[string]*string) *CreateAutoSnapshotPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateAutoSnapshotPolicyResponse) SetStatusCode(v int32) *CreateAutoSnapshotPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAutoSnapshotPolicyResponse) SetBody(v *CreateAutoSnapshotPolicyResponseBody) *CreateAutoSnapshotPolicyResponse {
	s.Body = v
	return s
}

func (s *CreateAutoSnapshotPolicyResponse) Validate() error {
	return dara.Validate(s)
}
